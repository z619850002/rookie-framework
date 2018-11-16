package cache

import (
	"testing"
	)

func TestRedisCache_TestConnection(t *testing.T) {
	config := RedisConfig{RedisAddr:"172.26.163.124:6379" , RedisPassword:"ztjztj120" , RedisDB:0}
	connector := NewRedisCache(config)
	err := connector.StartConnection()
	if (err != nil){
		t.Error(err)
	}
	err = connector.TestConnection()
	if (err != nil){
		t.Error(err)
	}
}


func TestOperation(t *testing.T){
	config := RedisConfig{RedisAddr:"172.26.163.124:6379" , RedisPassword:"ztjztj120" , RedisDB:0}
	connector := NewRedisCache(config)
	err := connector.StartConnection()
	if (err != nil){
		t.Error(err)
	}
	//Test connection
	err = connector.TestConnection()
	if (err != nil){
		t.Error(err)
	}
	//Set
	err = connector.Set("test1" , 1 , 0)
	if (err != nil){
		t.Error(err)
	}
	err = connector.Set("test2" , "2" , 0)
	if (err != nil){
		t.Error(err)
	}

	err = connector.Delete("test1")
	if (err != nil){
		t.Error(err)
	}
	err = connector.Delete("test2")
	if (err != nil){
		t.Error(err)
	}
	err = connector.Delete("test3")
	if (err != nil){
		t.Error(err)
	}




}