package gottp

import (
	"crypto/rand"
	"encoding/base64"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/seerx/got/cache"
)

// Manager 会话对象
type Manager struct {
	cookieName string
	cache      *cache.Manager
	// maxLifeTime int
}

// SSManager 全局变量
var SSManager *Manager

// Init 初始化
// 要使用 Session 功能必须在程序初始化时调用此函数
func InitSession(cookieName string, cache *cache.Manager) {
	// memcache.RegisterProvider()
	SSManager = &Manager{
		cookieName: cookieName,
		// maxLifeTime: int(expiredTime),
		cache: cache,
	}
}

// GenerateSID 产生唯一的Session ID
func (ss *Manager) GenerateSID() string {
	b := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return base64.URLEncoding.EncodeToString(b)
}

// func (ss *SessionObject) sessionGet(sid string) cache.Entity {
// 	return ss.cache.GetEntity(sid)
// }

// SessionStart 启动Session功能
func (ss *Manager) SessionStart(w http.ResponseWriter, r *http.Request) cache.Entity {
	cookie, err := r.Cookie(ss.cookieName)

	if err == nil && cookie.Value != "" {
		// 读取 Session
		sid, _ := url.QueryUnescape(cookie.Value)
		entity := ss.cache.GetEntity(sid)
		if entity != nil {
			// 成功读取
			return entity
		}
	}

	// 没有 session 创建
	sid := ss.GenerateSID()
	entity := ss.cache.NewEntity(sid)

	newCookie := http.Cookie{
		Name:     ss.cookieName,
		Value:    url.QueryEscape(entity.Name()),
		Path:     "/",
		HttpOnly: true,
		MaxAge:   int(ss.cache.MaxLifeTime),
	}
	http.SetCookie(w, &newCookie)

	return entity
}

// SessionDestory 注销Session
func (ss *Manager) SessionDestory(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie(ss.cookieName)
	if err != nil || cookie.Value == "" {
		return
	}

	ss.cache.DestroyEntity(cookie.Value)

	expiredTime := time.Now()
	newCookie := http.Cookie{
		Name:     ss.cookieName,
		Path:     "/",
		HttpOnly: true,
		Expires:  expiredTime,
		MaxAge:   -1,
	}
	http.SetCookie(w, &newCookie)
}
