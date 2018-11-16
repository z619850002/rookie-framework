package cache

import (
	"github.com/go-redis/redis"
	"time"
	"sync"
	"hub000.xindong.com/rookie/rookie-framework/syserr"
)


type RedisConfig struct {
	RedisAddr string
	RedisPassword	string
	RedisDB		int
}

var (
	//RedisAddr is the address of the redis instance.
	RedisAddr = "172.26.163.124:6379"
	//RedisPassword is the password to login the redis.
	RedisPassword = "ztjztj120"
	//RedisDB is the database redis will use.
	RedisDB = 0
)

//Cache is the implement of cache with redis.
type RedisCache struct {
	//client can operate the database.
	client	* redis.Client
	config 		RedisConfig
	mutex 	*sync.Mutex
}

//NewRedisCache will return a new RedisCache instance.
func NewRedisCache(config RedisConfig) *RedisCache{
	cache := &RedisCache{}
	cache.mutex = &sync.Mutex{}
	cache.config = config
	return cache
}

//StartConnection will create a new client connected to the redis.
func (h *RedisCache) StartConnection()(err error){
	config := h.config
	h.client = redis.NewClient(&redis.Options{
		Addr:config.RedisAddr,
		Password:config.RedisPassword,
		DB:config.RedisDB,
	})
	//Generate the error
	if h.client == nil{
		err = syserr.CacheConnectionError{Name:config.RedisAddr}
	}else{
		err = h.TestConnection()
	}
	return
}

//TestConnection will check the connection of the client.
func (h *RedisCache) TestConnection() (err error){
	_, err = h.client.Ping().Result()
	return
}

//Get takes in the key and return the value. Type string.
func (h *RedisCache) Get(key string) (string , error){
	value , err := h.client.Get(key).Result()
	return value , err
}

//GetInt takes in the key and return the value. Type int.
func (h *RedisCache) GetInt(key string) (int , error){
	value , err := h.client.Get(key).Int()
	return value , err
}

//GetFloat takes in the key and return the value. Type float64.
func (h *RedisCache) GetFloat(key string) (float64 , error){
	value , err := h.client.Get(key).Float64()
	return value , err
}

//Set takes in the key-value pair and save it in the database.
func (h *RedisCache) Set(key string , value interface{} , expire time.Duration) (error){
	h.mutex.Lock()
	defer func() {
		h.mutex.Unlock()
	}()
	err := h.client.Set(key , value , expire).Err()
	return err
}

//Delete takes the key and delete the key-value pair in the database.
func (h *RedisCache) Delete(key string) (error){
	h.mutex.Lock()
	defer func() {
		h.mutex.Unlock()
	}()
	err := h.client.Del(key).Err()
	return err
}
