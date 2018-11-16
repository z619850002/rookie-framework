package cache

import (
		"time"
	"hub000.xindong.com/rookie/rookie-framework/module"
)





//Module in cache package can be considered as the adapter of cache.
type CacheModule struct {
	module.Module
	cachePool map[string]Cache
}

//NewCacheModule will take in the cache and wrap it as a module.
func NewCacheModule() *CacheModule{
	cacheModule := &CacheModule{
		Module:*module.NewModule(),
	}
	cacheModule.cachePool = make(map[string] Cache)
	return cacheModule
}

//RegistCache takes in the name and the cache, then store the pair in the pool.
func (h *CacheModule) RegistCache(name string, c Cache){
	h.cachePool[name] = c
}



//Get takes in the key and return the value.
func (h *CacheModule)Get(name string , key string) (string, error){
	ret , err := h.LogicBlock.Call1(func(args ...interface{}) (interface{}, error) {
		name = args[0].(string)
		key = args[1].(string)
		//TODO:check!
		result , err := h.cachePool[name].Get(key)
		return result , err
	} ,name, key)
	result := ret.(string)
	return result, err
}

//Set takes in the key-value pair and the expire time, then save the pair in the database.
func (h *CacheModule)Set(name string , key string , value interface{} , expire time.Duration)(error){
	err := h.LogicBlock.Call0(func(args ...interface{}) (error) {
		name = args[0].(string)
		key = args[1].(string)
		value = args[2]
		expire = args[3].(time.Duration)
		//TODO:check!
		err := h.cachePool[name].Set(key , value , expire)
		return err
	} , name, key , value , expire)
	return err
}

//Delete takes in the key and delete the key-value pair in the cache source.
func (h *CacheModule)Delete(name string , key string) error{
	err := h.LogicBlock.Call0(func(args ...interface{}) error {
		name = args[0].(string)
		key = args[1].(string)
		//TODO:check!
		err := h.cachePool[name].Delete(key)
		return err
	} , name , key)
	return err
}

//TestConnection will check if the connection to the cache is normal.
func (h *CacheModule) TestConnection(name string)(error){
	err := h.LogicBlock.Call0(func(args ...interface{}) error {
		name = args[0].(string)
		//TODO:check!
		err := h.cachePool[name].TestConnection()
		return err
	} , name)
	return err
}
