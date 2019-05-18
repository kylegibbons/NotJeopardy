package main

import (
	"context"
	"fmt"
	"log"
	"sync"

	gocb "gopkg.in/couchbase/gocb.v1"
)

type DatabaseManager struct {
	//SaveChan chan SaveItem
	Cluster *gocb.Cluster
	Bucket  *gocb.Bucket
	//*sync.RWMutex
}

// SaveItem is the metadata used by the Saver to save an object to the DB
/*type SaveItem struct {
	ID       string
	Delete   bool
	Type     string
	Item     interface{}
	Metadata DocumentMetadata
}*/

type DocumentMetadata struct {
}

func (dm *DatabaseManager) Run(ctx context.Context, wg *sync.WaitGroup) {
	var err error
	//dm.SaveChan = make(chan SaveItem, 100)

	dm.Cluster, err = gocb.Connect(fmt.Sprintf("couchbase://%s", localSettings.DBHost))

	if err != nil {
		log.Printf("Unable to connect to DB: %v", err)
		return
	}

	dm.Cluster.Authenticate(gocb.PasswordAuthenticator{
		Username: localSettings.DBUser,
		Password: localSettings.DBPassword,
	})

	dm.Bucket, err = dm.Cluster.OpenBucket(localSettings.DBBucket, "")

	if err != nil {
		log.Printf("Unable to connect to DB bucket: %v", err)
		return
	}

	dm.Bucket.Manager("", "").CreatePrimaryIndex("", true, false)

	// TODO: Potential race condition when an item is deleted. There could be an update in the queue
	// which could potentially recreate the item in the DB

	wg.Done()

	/*
		for {
			select {
			case saveItem := <-dm.SaveChan:
				//dm.Lock()
				if saveItem.Delete {

					/*err = dm.db.Update(func(txn *badger.Txn) error {
						err := txn.Delete([]byte(fmt.Sprintf("%s:%s", saveItem.Type, saveItem.ID)))
						return err
					})

					if err != nil {
						log.Printf("Unable to delete item from database: %v", err)
						//dm.Unlock()
						continue
					}

					log.Printf("DELETING: %s", saveItem.Type+":"+saveItem.ID)
					//dm.Unlock()
					continue
				}

				itemJSON, err := json.MarshalIndent(saveItem.Item, "", "    ")

				if err != nil {
					log.Printf("can't encode item JSON: %v", err)
					//dm.Unlock()
					continue
				}

				//Write to DB
				_, err = dm.bucket.Upsert(fmt.Sprintf("%s:%s", saveItem.Type, saveItem.ID), itemJSON, 0)

				if err != nil {
					log.Printf("Unable to update database: %v", err)
					//dm.Unlock()
					continue
				} else {
					log.Printf("SAVED: %s", saveItem.Type+":"+saveItem.ID)
				}

			case <-ctx.Done():
				return

			}
		}
	*/
}

/*
func (dm *DatabaseManager) Get(Type string, ID string, value interface{}) error {

	//Ensure value is a pointer
	rv := reflect.ValueOf(value)
	if rv.Kind() != reflect.Ptr || rv.IsNil() {
		return fmt.Errorf("value is not a pointer")
	}

	// Get the value back
	//var item []byte

	_, err := dm.bucket.Get(fmt.Sprintf("%s:%s", Type, ID), &value)

	if err != nil {
		log.Printf("Unable to get item: %s:%s: %v", Type, ID, err)
		return nil, err
	}

	return nil

}
*/
