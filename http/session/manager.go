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
	lock        sync.RWMutex
	provider    Provider
	maxLifeTime int64
}

//AppSession 全局 session 变量
var AppSession *Manager

//InitSession 初始化 Session
func InitSession(provider string, cookieName string, expiredTime int64) {
	var err error
	AppSession, err = newManager(provider, cookieName, expiredTime)
	if err != nil {
		panic(err)
	}
	go AppSession.SessionGC()
}

// newManager 获取 session 管理器
func newManager(providerName string, cookieName string, maxLifeTime int64) (*Manager, error) {
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

func (m *Manager) sessionGet(sid string) *Session {
	m.lock.RLock()
	defer m.lock.RUnlock()

	session, e := m.provider.SessionRead(sid)
	if e == nil {
		return &session
	}

	return nil
}

func (m *Manager) generateSession() *Session {
	m.lock.Lock()
	defer m.lock.Unlock()

	sid := m.GenerateSID()
	session, _ := m.provider.SessionInit(sid)
	return &session
}

// SessionStart 启动Session功能
func (m *Manager) SessionStart(w http.ResponseWriter, r *http.Request) Session {

	cookie, err := r.Cookie(m.cookieName)

	if err == nil && cookie.Value != "" {
		// 读取 Session
		sid, _ := url.QueryUnescape(cookie.Value)
		ss := m.sessionGet(sid)
		if ss != nil {
			// 成功读取
			return *ss
		}
	}

	// 没有 session 创建
	ss := m.generateSession()
	newCookie := http.Cookie{
		Name:     m.cookieName,
		Value:    url.QueryEscape((*ss).SID()),
		Path:     "/",
		HttpOnly: true,
		MaxAge:   int(m.maxLifeTime),
	}
	http.SetCookie(w, &newCookie)

	return *ss
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

	// tf := tools.NewTimeFormatter()
	// fmt.Printf("m.maxLifeTime:%d\n", m.maxLifeTime)
	// fmt.Println("gc ..." + tf.FormatDateTime(time.Now()))
	// a := int(10)

	m.provider.SessionGC(m.maxLifeTime)

	du := time.Second * time.Duration(m.maxLifeTime)
	time.AfterFunc(du, func() {
		m.SessionGC()
	})
}
