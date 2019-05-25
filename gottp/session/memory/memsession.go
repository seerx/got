package memory

import (
	"sync"
	"time"
)

// memSession 内存 session
type memSession struct {
	Sid string
	// data sync.Map
	time time.Time
	lock *sync.RWMutex
	data map[interface{}]interface{}
}

func newMemSession(sid string) *memSession {
	ss := memSession{
		Sid:  sid,
		time: time.Now(),
		lock: new(sync.RWMutex)}
	ss.data = make(map[interface{}]interface{})
	return &ss
}

//SID 获取 session id
func (ss *memSession) SID() string {
	return ss.Sid
}

// Get 获取值
func (ss *memSession) Get(key interface{}) interface{} {
	ss.lock.RLock()
	defer ss.lock.RUnlock()
	return ss.data[key]
}

// Remove 删除值
func (ss *memSession) Remove(key interface{}) error {
	ss.lock.Lock()
	defer ss.lock.Unlock()
	delete(ss.data, key)
	return nil
}

//Set 设置 session 值
func (ss *memSession) Set(key, value interface{}) error {
	ss.lock.Lock()
	defer ss.lock.Unlock()
	ss.data[key] = value

	return nil
}
