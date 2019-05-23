package memory

import (
	"fmt"
	"sync"
	"time"

	"github.com/seerx/got/http/session"
)

func init() {
	session.RegisterProvider("memory", &memoryProvider{})
}

//memoryProvider 内存session 提供者
type memoryProvider struct {
	// sessionMap map[string]session.Session
	sessionMap sync.Map
}

func (p *memoryProvider) SessionInit(sid string) (session.Session, error) {
	ss := &Session{Sid: sid, time: time.Now()}
	p.sessionMap.Store(sid, ss)
	// p.sessionMap[sid] = ss
	return ss, nil
}

func (p *memoryProvider) SessionRead(sid string) (session.Session, error) {
	val, ok := p.sessionMap.Load(sid)
	if ok {
		ss := val.(Session)
		ss.time = time.Now()
		return &ss, nil
	}

	return nil, fmt.Errorf("No session found by id: %s", sid)
}

func (p *memoryProvider) SessionDestroy(sid string) error {
	p.sessionMap.Delete(sid)
	return nil
}

func (p *memoryProvider) SessionGC(maxLifeTime int64) {
	now := time.Now()
	p.sessionMap.Range(func(key, val interface{}) bool {
		ss := val.(Session)

		if now.Unix()-ss.time.Unix() > maxLifeTime {
			// 销毁超时的 Session
			p.SessionDestroy(ss.SID())
		}

		return true
	})
}
