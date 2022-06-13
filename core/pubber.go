package core

import (
	"fmt"
	"time"
)

type Handler func(ctx *Context)

type processMeta struct {
	handler Handler
	ctx     *Context
}

type pubber struct {
	pubberMap  map[string][]Handler
	notifyChan chan *processMeta
}

func NewPubber() *pubber {
	return &pubber{
		pubberMap:  map[string][]Handler{},
		notifyChan: make(chan *processMeta),
	}
}

func (d *pubber) Start() {
	go func() {
		for meta := range d.notifyChan {
			func() {
				// 防止程序寄了
				defer RecoverAndPrint()
				meta.handler(meta.ctx)
			}()
		}
	}()
}

func (d *pubber) AddSub(evt string, handler Handler) {
	d.pubberMap[evt] = append(d.pubberMap[evt], handler)
}

func (d *pubber) Pub(evt string, data interface{}) {
	ctx := d.makeContext(evt, data)
	handlers, ok := d.pubberMap[evt]
	if !ok {
		fmt.Printf("创建分发任务失败，事件【%s】不存在\n", evt)
		return
	}
	for _, handler := range handlers {
		d.notifyChan <- &processMeta{
			handler: handler,
			ctx:     ctx,
		}
	}

}

func (d *pubber) makeContext(evt string, data interface{}) *Context {
	return &Context{
		Event:     evt,
		Timestamp: time.Now().UnixMilli(),
		Data:      data,
	}
}
