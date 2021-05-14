package cache

import (
	"reflect"
	"testing"
	"time"
)

var (
	c = New()
)

func TestLocalCache_Set(t *testing.T) {
	type args struct {
		key    string
		value  interface{}
		expire time.Duration
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "第一组 Set 测试",
			args: args{
				key:    "key_test_1",
				value:  "false",
				expire: 0,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		if err := c.Set(tt.args.key, tt.args.value, tt.args.expire); (err != nil) != tt.wantErr {
			t.Errorf("Set() error = %v, wantErr %v", err, tt.wantErr)
		}
	}
}

func TestLocalCache_Get(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name      string
		args      args
		wantValue interface{}
		wantErr   bool
	}{
		// TODO: Add test cases.
		{
			name: "第一组 Get 测试",
			args: args{
				key: "key_test_1",
			},
			wantValue: "false",
			wantErr:   false,
		},
	}
	for _, tt := range tests {
		gotValue, err := c.Get(tt.args.key)
		if (err != nil) != tt.wantErr {
			t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
		if !reflect.DeepEqual(gotValue, tt.wantValue) {
			t.Errorf("Get() gotValue = %v, want %v", gotValue, tt.wantValue)
		}
	}
}

func TestLocalCache_Delete(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "第一组 Del 测试",
			args: args{
				key: "key_test_1",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		if err := c.Delete(tt.args.key); (err != nil) != tt.wantErr {
			t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
		}
	}
}

func TestAll(t *testing.T) {
	t.Run("TestLocalCache_Set", TestLocalCache_Set)
	t.Run("TestLocalCache_Get", TestLocalCache_Get)
	t.Run("TestLocalCache_Delete", TestLocalCache_Delete)
}
