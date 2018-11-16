package config

import (
	_ "github.com/go-sql-driver/mysql"
)


type Config interface {
	GetParameters(ic Parameter) error
}

