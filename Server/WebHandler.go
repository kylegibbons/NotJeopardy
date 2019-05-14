package main

import (
	"fmt"
	"log"
	"net/http"
)

type WebHandler struct {
}

func (wh *WebHandler) SetAccessHeaders(w http.ResponseWriter) {
	if !localSettings.Production {
		w.Header().Set("Access-Control-Allow-Headers", "*")
		w.Header().Set("Access-Control-Allow-Methods", "DELETE, GET, HEAD, POST, PUT, OPTIONS, TRACE")
		w.Header().Set("Access-Control-Allow-Origin", "*")
	}
}

func (wh *WebHandler) Index(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/app", 301)
}

func (wh *WebHandler) notFound(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/JSON; charset=UTF-8")
	w.WriteHeader(http.StatusNotFound)

	fmt.Fprint(w, "404 Not Found!\r\r")
	fmt.Fprintf(w, "%v", r.RequestURI)

	log.Printf(
		"%s\t%s NOT FOUND",
		r.Method,
		r.RequestURI,
		//name,
		//time.Since(start),
	)

}

/*func (wh *WebHandler) GetLocations(w http.ResponseWriter, r *http.Request) {

	wh.SetAccessHeaders(w)

	print, _ := locationManager.ToJSON()

	fmt.Fprintf(w, "%s", print)

}*/

/*func PushRegistration(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/JSON; charset=UTF-8")

	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error reading body: %v", err)
		http.Error(w, "Can't read body", http.StatusBadRequest)
		return
	}

	//subJSON := `{<YOUR SUBSCRIPTION JSON>}`

	// Decode subscription
	sub := webpush.Subscription{}
	if err := json.Unmarshal(body, &sub); err != nil {
		log.Fatal(err)
	}

	// Send Notification
	_, err = webpush.SendNotification([]byte("Test"), &sub, &webpush.Options{
		Subscriber:      "<EMAIL@EXAMPLE.COM>",
		VAPIDPrivateKey: settings.VapidPrivateKey,
	})
	if err != nil {
		log.Fatal(err)
	}

	//vars := mux.Vars(r)
	//if incident, ok := incidents[GUID]; ok {
}*/
