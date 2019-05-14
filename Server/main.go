package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/handlers"
)

const (
	//VERSION of the software
	VERSION = "0.1.0"
	//BUILD of the software
	BUILD = 100
)

var (
	localSettings settings
	debugChan     = make(chan string, 100)
	webHandler    WebHandler
	mw            io.Writer
	logFile       io.WriteCloser
	hub           *Hub
	gameManager   GameManager
)

func main() {

	var err error

	err = localSettings.getSettings("Settings.yaml")
	if err != nil {
		log.Printf("unable to get settings: %v\n", err)
		os.Exit(2)
		return
	}

	if localSettings.UpdateInterval < 1 {
		localSettings.UpdateInterval = 1
	}

	run()

}

func run() {
	var err error

	logFile, err = os.OpenFile("testlogfile.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening log file: %v", err)
	}
	defer logFile.Close()

	mw = io.MultiWriter(os.Stdout, logFile)
	log.SetOutput(mw)

	fmt.Print("\n\n\n")
	fmt.Print("NOT Jeopardy!\n")
	fmt.Print("(C)2019 ThisIsNotJeopardy.com\n")
	fmt.Printf("Version: %s Build: %d\n\n", VERSION, BUILD)

	if localSettings.Production {
		log.Print("PRODUCTION MODE\n\n")
	}

	appCtx, appCancel := context.WithCancel(context.Background())
	defer appCancel()

	errors := make(chan appError)

	go errorHandler(appCtx, appCancel, errors)

	router := NewRouter()

	httpServer, httpErrorChan := startHTTPServer(localSettings.Listener, router)

	httpShutdownCtx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	defer httpServer.Shutdown(httpShutdownCtx)

	hub = newHub()
	go hub.run()

	go gameManager.Run(appCtx)

	sigs := make(chan os.Signal, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		select {
		case <-sigs:
			log.Println("Received the TERM or INTERUPT signal. Quitting")
			appCancel()
			return
		}
	}()

	for {
		select {
		case err := <-httpErrorChan:
			log.Printf("%v", err)
		case <-appCtx.Done():
			return
		default:
		}
	}
}

func startHTTPServer(listener string, handler http.Handler) (*http.Server, <-chan error) {
	/*cer, err := tls.LoadX509KeyPair("server.crt", "server.key")
	if err != nil {
		log.Println(err)
		return nil, nil
	}

	tlsConfig := &tls.Config{Certificates: []tls.Certificate{cer}}
	*/
	srv := &http.Server{
		Addr:    listener,
		Handler: handler,
		//TLSConfig: tlsConfig,
	}
	errorChan := make(chan error)

	headersOk := handlers.AllowedHeaders([]string{"Content-Type", "X-Requested-With"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "DELETE", "OPTIONS"})

	srv.Handler = handlers.CORS(headersOk, originsOk, methodsOk)(srv.Handler)

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			errorChan <- fmt.Errorf("http server error: %v", err)
		}
	}()

	// returning reference so caller can call Shutdown()
	return srv, errorChan
}

func errorHandler(ctx context.Context, cancel context.CancelFunc, errorChan <-chan appError) {
	select {
	case thisError := <-errorChan:
		switch thisError.Severity {
		case "MINOR":
			log.Printf("MINOR: %v", thisError.Error())

		case "CRITICAL":
			log.Printf("CRITICAL: %v", thisError.Error())

		case "FATAL":
			log.Printf("FATAL: %v", thisError.Error())
			cancel()
			return
		}
	case <-ctx.Done():
		return
	}

}

type appError struct {
	error
	Severity string
}
