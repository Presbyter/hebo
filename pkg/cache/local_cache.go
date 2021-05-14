package cache

import (
	"bytes"
	"context"
	"encoding/gob"
	"sync"
	"time"
)

type valueEx struct {
	Value     interface{}
	Expire    time.Duration
	StartTime int64
}

type LocalCache struct {
	m         sync.Map
	closeChan chan string
	cancelMap sync.Map
}

func (l *LocalCache) Set(key string, value interface{}, expire time.Duration) error {
	var t = valueEx{
		Expire:    expire,
		StartTime: time.Now().UnixNano() / 1e6,
		Value:     value,
	}

	buf := new(bytes.Buffer)
	gobEncoder := gob.NewEncoder(buf)
	if err := gobEncoder.Encode(&t); err != nil {
		return err
	}

	// 查找该key是否已经存在
	if _, ok := l.m.Load(key); ok {
		// 更新定时器
		if v, ok := l.cancelMap.Load(key); ok {
			v.(context.CancelFunc)()
			l.cancelMap.Delete(key)
		}
	}

	// 设置定时器
	if t.Expire != 0 {
		ctx, cancel := context.WithCancel(context.Background())
		go func(ctx context.Context, s string, exp time.Duration) {
			timer := time.NewTimer(exp)
			select {
			case <-ctx.Done():
				return
			case <-timer.C:
				l.closeChan <- key
			}
			l.closeChan <- s
		}(ctx, key, expire)
		l.cancelMap.Store(key, cancel)
	}

	// 更新数据
	l.m.Store(key, buf.Bytes())

	return nil
}

func (l *LocalCache) Get(key string) (value interface{}, err error) {
	if v, ok := l.m.Load(key); ok {
		if vv, ok := v.([]byte); ok {
			t := valueEx{}
			gobDecoder := gob.NewDecoder(bytes.NewReader(vv))
			if err = gobDecoder.Decode(&t); err != nil {
				return nil, err
			}
			return t.Value, nil
		} else {
			return nil, DataTypeErr
		}
	} else {
		return nil, NotFoundKeyErr
	}
}

func (l *LocalCache) Delete(key string) error {
	if _, ok := l.m.Load(key); !ok {
		return NotFoundKeyErr
	}
	l.m.Delete(key)
	if v, ok := l.cancelMap.Load(key); ok {
		v.(context.CancelFunc)()
		l.cancelMap.Delete(key)
	}
	return nil
}

func New() cache {
	entity := new(LocalCache)
	entity.m = sync.Map{}
	entity.cancelMap = sync.Map{}
	entity.closeChan = make(chan string, 1<<32)
	go entity.delExpireKey()
	return entity
}

func (l *LocalCache) delExpireKey() {
	for v := range l.closeChan {
		l.m.Delete(v)
	}
}
