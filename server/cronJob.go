package server

import (
	"gobang/variable"
	"log"
	"time"
)

type Fn func()

type MyTicker struct {
	MyTick *time.Ticker
	Runner Fn
}

func NewMyTick(interval int, f Fn) *MyTicker {
	return &MyTicker{
		MyTick: time.NewTicker(time.Duration(interval) * time.Second),
		Runner: f,
	}
}

func (t *MyTicker) Start() {
	for {
		<-t.MyTick.C
		t.Runner()
	}
}

func GarbageCollection() {
	// sessionMap 回收已经结束的对局信息
	log.Default().Println("GCing, length of sessionMap:", len(variable.SessionMap))
	for key, session := range variable.SessionMap {
		if session.Finish {
			delete(variable.SessionMap, key)
		}
	}
}
