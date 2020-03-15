package service

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Session struct {
	store    Store
	ctx      *gin.Context
	Name     string
	liftTime int // 单位是s
}

func NewSession(ctx *gin.Context, name string, liftTime int) *Session {
	// 这里查看配置，看目前配置的什么类型的store
	return &Session{nil, ctx, name, liftTime}
}

func (session *Session) GetSessionId() string {
	var (
		sid string
	)
	if sid, _ = session.ctx.Cookie(session.Name); sid == "" {
		sid = session.CreateId()
	}

	return sid
}

func (session *Session) CreateId() string {
	return uuid.New().String()
}

func (session *Session) Get() interface{} {
	return session.store.Get()
}

func (session *Session) Put(data interface{}) bool {
	return session.store.Set(data)
}

func (session *Session) CreateDriver(driver string) error {
	var err error
	switch driver {
	case "redis":
		session.store = NewRedisStore(session)
	default:
		err = errors.New("not support file driver")
	}

	return err
}

func (session *Session) SyncSessionIdToCookie() {
	session.ctx.SetCookie(session.Name, session.GetSessionId(), session.liftTime, "/", "*", false, false)
}

type Store interface {
	Get() interface{}
	Set(data interface{}) bool
	SetSession(session *Session)
}

type RedisStore struct {
	session *Session
}

func NewRedisStore(session *Session) *RedisStore {
	return &RedisStore{
		session: session,
	}
}

func (store *RedisStore) Get() interface{} {
	return nil
}

func (store *RedisStore) Set(data interface{}) bool {
	return true
}

func (store *RedisStore) SetSession(session *Session) {
	store.session = session
}
