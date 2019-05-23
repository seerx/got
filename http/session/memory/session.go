package memory

import (
	"time"
)

// Session 内存 session
type Session struct {
	Sid string
	// data sync.Map
	time time.Time
	// data map[interface{}]interface{}
}

//SID 获取 session id
func (ss *Session) SID() string {
	return ss.Sid
}

// Get 获取值
func (ss *Session) Get(key interface{}) interface{} {
	return nil
}

// Remove 删除值
func (ss *Session) Remove(key interface{}) error {
	return nil
}

//Set 设置 session 值
func (ss *Session) Set(key, value interface{}) error {

	return nil
}
