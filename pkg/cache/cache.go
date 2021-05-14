package cache

import (
	"errors"
	"time"
)

type cache interface {
	Set(key string, value interface{}, expire time.Duration) error
	Get(key string) (value interface{}, err error)
	Delete(key string) error
}

var (
	NotFoundKeyErr = errors.New("not found the key")
	DataTypeErr    = errors.New("data type error")
)
