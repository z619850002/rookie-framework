package cache

import (
	"testing"
	"fmt"
	"hub000.xindong.com/rookie/rookie-framework/protobuf"
)

func TestCacheHandler(t *testing.T) {
	OnInit()
	cache := NewRedisCache(RedisConfig{RedisAddr:"172.26.163.124:6379" , RedisPassword:"ztjztj120" , RedisDB:0})
	cache.StartConnection()
	Module.RegistCache("base" , cache)

	//Test the test connection.
	resp , err := Module.Call(Communicator , protobuf.CacheTestConnectionReq{"base"})
	if (err != nil){
		t.Error(err)
	}
	if resp.GetProtocolNum() == protobuf.CacheTestConnectionRespNum{
		fmt.Println("resp 1:")
		fmt.Println(resp.GetProtocolNum())
		resp := resp.(protobuf.CacheTestConnectionResp)
		fmt.Println(resp.Status)
	}


	//Test the set.
	resp , err = Module.Call(Communicator , protobuf.CacheSetReq{"base" , "test" , "test_value" , 0})
	if (err != nil){
		t.Error(err)
	}
	if resp.GetProtocolNum() == protobuf.CacheSetRespNum{
		fmt.Println("resp 2:")
		fmt.Println(resp.GetProtocolNum())
		resp := resp.(protobuf.CacheSetResp)
		fmt.Println(resp.Status)
	}

	//Test the get.
	resp , err = Module.Call(Communicator , protobuf.CacheGetReq{"base" , "test" })
	if (err != nil){
		t.Error(err)
	}
	if resp.GetProtocolNum() == protobuf.CacheGetRespNum{
		fmt.Println("resp 3:")
		fmt.Println(resp.GetProtocolNum())
		resp := resp.(protobuf.CacheGetResp)
		fmt.Println(resp.Value)
		fmt.Println(resp.Status)
	}

	//Test the delete.
	resp , err = Module.Call(Communicator , protobuf.CacheDeleteReq{"base" , "test" })
	if (err != nil){
		t.Error(err)
	}
	if resp.GetProtocolNum() == protobuf.CacheDeleteRespNum{
		fmt.Println("resp 4:")
		fmt.Println(resp.GetProtocolNum())
		resp := resp.(protobuf.CacheDeleteResp)
		fmt.Println(resp.Status)
	}


	resp , err = Module.Call(Communicator , protobuf.CacheGetReq{"base" , "test" })
	if (err != nil){
		t.Error(err)
	}
	if resp.GetProtocolNum() == protobuf.CacheGetRespNum{
		fmt.Println("resp 5:")
		fmt.Println(resp.GetProtocolNum())
		resp := resp.(protobuf.CacheGetResp)
		fmt.Println(resp.Value)
		fmt.Println(resp.Status)
	}

}