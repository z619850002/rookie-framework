package db

import (
	"fmt"
	"hub000.xindong.com/rookie/rookie-framework/config"
	"testing"

	"hub000.xindong.com/rookie/rookie-framework/communication"
	"hub000.xindong.com/rookie/rookie-framework/protobuf"
)

func TestDBHandler(t *testing.T) {
	Module = NewDBModule()
	c := &config.ConfigDB{
		Dbhost:     "172.26.163.76",
		Dbport:     "3306",
		Dbuser:     "rookie2",
		Dbpassword: "12345678",
		Dbname:     "dbconfig",
		Tblname:    "dbinfo",
	}
	Communicator = communication.NewRPCCommunicator()
	Module.SetCommunicator(Communicator)
	Module.SetHandler(&Handler{module: Module})
	Module.Run()

	//测试Query
	resp, err := Module.Call(Communicator, protobuf.DBQueryReq{DsName: "mysql2", StrSQL: `select name,age from table1 where name="lihao"`})
	if err != nil {
		t.Error(err)
	}
	if resp.GetProtocolNum() == protobuf.DBQueryRespNum {
		fmt.Println("1. response for query:")
		fmt.Println(resp.GetProtocolNum())
		resp := resp.(protobuf.DBQueryResp)
		rows := resp.Rows
		fmt.Println(resp.Status)
		res := []User{}
		for rows.Next() {
			r := User{}
			rows.Scan(&r.Name, &r.Age)
			res = append(res, r)
		}
		fmt.Println(res[0])
		if res[0] != (User{Name: "lihao", Age: 25}) {
			t.Error("Exec函数测试错误！")
		}
	}

	//测试QueryRow
	resp3, err := Module.Call(Communicator, protobuf.DBQueryRowReq{DsName: "mysql2", StrSQL: `select name,age from table1 where name="lihao"`})
	if err != nil {
		t.Error(err)
	}
	if resp3.GetProtocolNum() == protobuf.DBQueryRowRespNum {
		fmt.Println("3. response for queryRow:")
		fmt.Println(resp3.GetProtocolNum())
		resp := resp3.(protobuf.DBQueryRowResp)
		rows := resp.Row
		fmt.Println(resp.Status)
		res := User{}
		rows.Scan(&res.Name, &res.Age)
		fmt.Println(res)
		if res != (User{Name: "lihao", Age: 25}) {
			t.Error("QueryRow函数测试错误！")
		}
	}

	//测试Exec
	resp2, err2 := Module.Call(Communicator, protobuf.DBExecReq{DsName: "mysql2", StrSQL: `update table1 set age=? where name=?`, Args: []interface{}{100, "lihao"}})
	if err2 != nil {
		t.Error(err)
	}
	if resp2.GetProtocolNum() == protobuf.DBExecRespNum {
		//更新lihao的age
		fmt.Println("2. response for exec:")
		fmt.Println(resp2.GetProtocolNum())
		resp2 := resp2.(protobuf.DBExecResp)
		fmt.Println(resp2.Status)

		//查询是否更新成功
		resp, err = Module.Call(Communicator, protobuf.DBQueryReq{DsName: "mysql2", StrSQL: `select name,age from table1 where name="lihao"`})
		if err != nil {
			t.Error(err)
		}
		if resp.GetProtocolNum() == protobuf.DBQueryRespNum {
			fmt.Println("1. response for query:")
			fmt.Println(resp.GetProtocolNum())
			resp := resp.(protobuf.DBQueryResp)
			rows := resp.Rows
			fmt.Println(resp.Status)
			res := []User{}
			for rows.Next() {
				r := User{}
				rows.Scan(&r.Name, &r.Age)
				res = append(res, r)
			}
			fmt.Println(res[0])
			if res[0] != (User{Name: "lihao", Age: 100}) {
				t.Error("Exec函数测试错误！")
			}
		}
	}
	//还原数据
	Module.Call(Communicator, protobuf.DBExecReq{DsName: "mysql2", StrSQL: `update table1 set age=? where name=?`, Args: []interface{}{25, "lihao"}})

}
