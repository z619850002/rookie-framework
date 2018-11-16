package config

import (
	"fmt"
	"database/sql"
	"reflect"
	"unsafe"
	"strconv"
	"errors"
	"strings"
)

type ConfigDB struct {
	DriverName 	string
	Dbhost     string
	Dbport     string
	Dbuser     string
	Dbpassword string
	Dbname     string
	Tblname    string
}

func (c *ConfigDB) GetParameters(ic Parameter) error {
	//首先获得map
	paramsStrmap := map[string]string{}
	if c.Dbhost == "" {
		c.Dbhost = "127.0.0.1"
	}
	if c.Dbport == "" {
		c.Dbport = "3306"
	}
	if c.Dbuser == "" {
		c.Dbuser = "root"
	}
	if c.Tblname == "" {
		c.Tblname = "PARAMS"
	}
	dburl := c.Dbuser + ":" + c.Dbpassword + "@tcp(" + c.Dbhost + ":" + c.Dbport + ")/" + c.Dbname + "?charset=utf8"
	db, err := sql.Open(c.DriverName, dburl)
	if err != nil {
		fmt.Println(err)
		return err
	}
	rows, err := db.Query(`SELECT NAME,VALUE FROM ` + c.Tblname)
	if err != nil {
		fmt.Println(err)
		return err
	}
	for rows.Next() {
		var name, value string
		err = rows.Scan(&name, &value)
		if err != nil {
			fmt.Println(err)
			return err
		}
		//Change the name into lower case.
		name = strings.ToLower(name)
		paramsStrmap[name] = value
	}

	//根据map进行反射，为ic结构体赋值
	val := reflect.ValueOf(ic)
	kd := val.Kind()
	if kd != reflect.Ptr && val.Elem().Kind() == reflect.Struct {
		fmt.Println("expect struct")
		return errors.New("No right input")
	}
	numFields := val.Elem().NumField()
	for i := 0; i < numFields; i++ {
		fieldName := fmt.Sprintf("%v", val.Elem().Type().Field(i).Name)
		lowerFieldName := strings.ToLower(fieldName)
		if _, ok := paramsStrmap[lowerFieldName]; !ok {
			//FIXME:this shouldn`t return error. just skip it.
			continue
			//return errors.New("No value in map")
		}

		field := val.Elem().FieldByName(fieldName)
		switch field.Kind() {
		case reflect.String:
			*(*string)(unsafe.Pointer(field.Addr().Pointer())) = paramsStrmap[lowerFieldName]
		case reflect.Int:
			item, err := strconv.Atoi(paramsStrmap[lowerFieldName])
			if err != nil {
				fmt.Println(err)
				return err
			}
			*(*int)(unsafe.Pointer(field.Addr().Pointer())) = item
		case reflect.Int8:
			item, err := strconv.ParseInt(paramsStrmap[lowerFieldName], 10, 8)
			if err != nil {
				fmt.Println(err)
				return err
			}
			*(*int8)(unsafe.Pointer(field.Addr().Pointer())) = int8(item)
		case reflect.Int16:
			item, err := strconv.ParseInt(paramsStrmap[lowerFieldName], 10, 16)
			if err != nil {
				fmt.Println(err)
				return err
			}
			*(*int16)(unsafe.Pointer(field.Addr().Pointer())) = int16(item)
		case reflect.Int32:
			item, err := strconv.ParseInt(paramsStrmap[lowerFieldName], 10, 32)
			if err != nil {
				fmt.Println(err)
				return err
			}
			*(*int32)(unsafe.Pointer(field.Addr().Pointer())) = int32(item)
		case reflect.Int64:
			item, err := strconv.ParseInt(paramsStrmap[lowerFieldName], 10, 64)
			if err != nil {
				fmt.Println(err)
				return err
			}
			*(*int64)(unsafe.Pointer(field.Addr().Pointer())) = int64(item)
		case reflect.Bool:
			item, err := strconv.ParseBool(paramsStrmap[lowerFieldName])
			if err != nil {
				fmt.Println(err)
				return err
			}
			*(*bool)(unsafe.Pointer(field.Addr().Pointer())) = item
		case reflect.Float32:
			item, err := strconv.ParseFloat(paramsStrmap[lowerFieldName], 32)
			if err != nil {
				fmt.Println(err)
				return err
			}
			*(*float32)(unsafe.Pointer(field.Addr().Pointer())) = float32(item)
		case reflect.Float64:
			item, err := strconv.ParseFloat(paramsStrmap[lowerFieldName], 64)
			if err != nil {
				fmt.Println(err)
				return err
			}
			*(*float64)(unsafe.Pointer(field.Addr().Pointer())) = item
		}
	}
	return nil
}
