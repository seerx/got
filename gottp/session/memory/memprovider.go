package memory

import (
	"fmt"
	"time"

	"github.com/seerx/got/gottp/session"
)

//PROVIDER 提供者名称
const PROVIDER = "memory"

func init() {
	provider := memProvider{sessionMap: make(map[string]memSession)}
	session.RegisterProvider(PROVIDER, &provider)
}

//memoryProvider 内存session 提供者
type memProvider struct {
	sessionMap map[string]memSession
	// sessionMap sync.Map
}

func (p *memProvider) SessionInit(sid string) (session.Session, error) {
	ss := newMemSession(sid)
	// p.sessionMap.Store(sid, ss)
	p.sessionMap[sid] = *ss
	return ss, nil
}

func (p *memProvider) SessionRead(sid string) (session.Session, error) {
	ss, ok := p.sessionMap[sid]
	if ok {
		// ss := val.(*memSession)
		ss.time = time.Now()
		return &ss, nil
	}

	return nil, fmt.Errorf("No session found by id: %s", sid)
}

func (p *memProvider) SessionDestroy(sid string) error {
	// p.sessionMap.Delete(sid)
	delete(p.sessionMap, sid)
	return nil
}

func (p *memProvider) SessionGC(maxLifeTime int64) {
	now := time.Now()
	for _, ss := range p.sessionMap {
		if now.Unix()-ss.time.Unix() > maxLifeTime {
			// 销毁超时的 Session
			p.SessionDestroy(ss.SID())
		}
	}
}
