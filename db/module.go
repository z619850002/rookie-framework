package db

import (
	"database/sql"
	"sync"

	"github.com/pkg/errors"
	"hub000.xindong.com/rookie/rookie-framework/module"
)

//DBModule is a module which provides a datasource pool
type DBModule struct {
	module.Module
	DataSourcePool map[string] DataSource
	mutex          *sync.RWMutex
}

//NewDBModule returns an DBModule instance
func NewDBModule() *DBModule {
	return &DBModule{Module: *module.NewModule(), mutex: &sync.RWMutex{}, DataSourcePool: make(map[string]DataSource)}
}

//RegisterDB is used to register a database source
func (h *DBModule) RegisterDataSource(name string, ds DataSource) error {
	h.mutex.Lock()
	defer func() {
		h.mutex.Unlock()
	}()

	if _, ok := h.DataSourcePool[name]; !ok {
		h.DataSourcePool[name] = ds
		return nil
	}
	return errors.New("the datasource name has been registered")
}

//Exec is for updating, deleting, inserting data into database
func (h *DBModule) Exec(dsName string, strSQL string, args ...interface{}) error {
	err := h.LogicBlock.Call0(func(args2 ...interface{}) error {
		return h.DataSourcePool[dsName].Exec(strSQL, args...)
	}, dsName, strSQL, args)
	return err
}

//Query is for querying data from database
func (h *DBModule) Query(dsName string, strSQL string, args ...interface{}) (*sql.Rows, error) {
	ret, err := h.LogicBlock.Call1(func(args2 ...interface{}) (interface{}, error) {
		return h.DataSourcePool[dsName].Query(strSQL, args...)
	}, dsName, strSQL, args)
	result := ret.(*sql.Rows)
	return result, err
}

//QueryRow is for querying one data from database
func (h *DBModule) QueryRow(dsName string, strSQL string, args ...interface{}) (*sql.Row, error) {
	ret, err := h.LogicBlock.Call1(func(args2 ...interface{}) (interface{}, error) {
		return h.DataSourcePool[dsName].QueryRow(strSQL, args...)
	}, dsName, strSQL, args)
	result := ret.(*sql.Row)
	return result, err
}
