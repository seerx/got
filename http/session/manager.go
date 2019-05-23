package session

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"sync"
	"time"
)

//Manager Session 管理
type Manager struct {
	cookieName  string
	lock        sync.Mutex
	provider    Provider
	maxLifeTime int64
}

//AppSession 全局 session 变量
var AppSession *Manager

func initSession(provider string, cookieName string, expiredTime int64) {
	AppSession, _ = GetManager(provider, cookieName, expiredTime)
	go AppSession.SessionGC()
}

// GetManager 获取 session 管理器
func GetManager(providerName string, cookieName string, maxLifeTime int64) (*Manager, error) {
	provider, ok := providers[providerName]
	if !ok {
		return nil, fmt.Errorf("session: unknown provide %q (forgotten import?)", providerName)
	}

	return &Manager{
		cookieName:  cookieName,
		maxLifeTime: maxLifeTime,
		provider:    provider,
	}, nil
}

// GenerateSID 产生唯一的Session ID
func (m *Manager) GenerateSID() string {
	b := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return base64.URLEncoding.EncodeToString(b)
}

// SessionStart 启动Session功能
func (m *Manager) SessionStart(w http.ResponseWriter, r *http.Request) Session {
	m.lock.Lock()
	defer m.lock.Unlock()
	cookie, err := r.Cookie(m.cookieName)
	var session Session
	if err != nil || cookie.Value == "" {
		sid := m.GenerateSID()
		session, _ = m.provider.SessionInit(sid)
		newCookie := http.Cookie{
			Name:     m.cookieName,
			Value:    url.QueryEscape(sid),
			Path:     "/",
			HttpOnly: true,
			MaxAge:   int(m.maxLifeTime),
		}
		http.SetCookie(w, &newCookie)
	} else {
		sid, _ := url.QueryUnescape(cookie.Value)
		session, _ = m.provider.SessionRead(sid)
	}

	return session
}

// SessionDestory 注销Session
func (m *Manager) SessionDestory(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie(m.cookieName)
	if err != nil || cookie.Value == "" {
		return
	}

	m.lock.Lock()
	defer m.lock.Unlock()
	m.provider.SessionDestroy(cookie.Value)
	expiredTime := time.Now()
	newCookie := http.Cookie{
		Name:     m.cookieName,
		Path:     "/",
		HttpOnly: true,
		Expires:  expiredTime,
		MaxAge:   -1,
	}
	http.SetCookie(w, &newCookie)
}

// SessionGC Session 垃圾回收
func (m *Manager) SessionGC() {
	m.lock.Lock()
	defer m.lock.Unlock()
	m.provider.SessionGC(m.maxLifeTime)
	time.AfterFunc(time.Duration(m.maxLifeTime), func() {
		m.SessionGC()
	})
}
