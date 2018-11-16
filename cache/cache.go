package cache

import (
	"time"
	)

//Cache is the base interface of all cache datasource.
type Cache interface {
	//StartConnection will let this cache connect to the source.
	StartConnection()(error)
	//Get takes in the key and return the value. All data stores by string.
	Get(key string) (string, error)
	//Store takes in the key-value pair and the expire time and save or update the data in the datasource.
	Set(key string , value interface{} , expire time.Duration)(error)
	//Delete takes in the key and delete the data in the datasource.
	Delete(key string)(error)
	//TestConnection test if the connection is normal.
	TestConnection()(error)
}
