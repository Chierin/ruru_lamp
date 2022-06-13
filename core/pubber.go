package core

import (
	"time"
)

type Handler func(ctx *Context)

type notifyMeta struct {
	handler Handler
	ctx     *Context
}

type Pubber struct {
	pubberMap  map[string][]Handler
	notifyChan chan *notifyMeta
}

func NewPubber() *Pubber {
	return &Pubber{
		pubberMap:  map[string][]Handler{},
		notifyChan: make(chan *notifyMeta),
	}
}

func (d *Pubber) Start() {
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

func (d *Pubber) AddSub(evt string, handler Handler) {
	d.pubberMap[evt] = append(d.pubberMap[evt], handler)
}

func (d *Pubber) Pub(evt string, state *State) {
	ctx := d.makeContext(evt, *state)
	handlers, ok := d.pubberMap[evt]
	if !ok {
		Logger.Warnf("发布失败，事件【%s】不存在", evt)
		return
	}
	for _, handler := range handlers {
		d.notifyChan <- &notifyMeta{
			handler: handler,
			ctx:     ctx,
		}
	}

}

func (d *Pubber) makeContext(evt string, state State) *Context {
	return &Context{
		Event:     evt,
		Timestamp: time.Now().UnixMilli(),
		State:     state,
	}
}
