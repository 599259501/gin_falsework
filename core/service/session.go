package service

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Session struct {
	store Store
	ctx *gin.Context
}

func NewSession(ctx *gin.Context) *Session{
	// 这里查看配置，看目前配置的什么类型的store
	return &Session{nil, ctx}
}

func (session *Session)GetSessionId() string  {
	var (
		sid string
	)
	if sid,_ = session.ctx.Cookie("sessionId");sid == ""{
		sid = session.CreateId()
	}

	return sid
}

func (session *Session) CreateId() string  {
	return uuid.New().String()
}

func (session *Session)Get() interface{}{
	return session.store.Get()
}

func (session *Session)Put(data interface{}) bool{
	return session.store.Set(data)
}

func (session *Session)CreateDriver(driver string) error{
	var err error
	switch driver {
	case "redis":
		session.store = NewRedisStore()
	default:
		err = errors.New("not support file driver")
	}

	return err
}


type Store interface {
	Get() interface{}
	Set(data interface{})bool
}

type RedisStore struct {
}

func NewRedisStore() *RedisStore {
	return &RedisStore{
	}
}

func (store *RedisStore) Get()interface{} {
	return nil
}

func (store *RedisStore) Set(data interface{}) bool {
	return true
}
