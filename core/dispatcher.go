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

type dispatcher struct {
	dispatchMap map[string][]Handler
	processChan chan *processMeta
}

func newDispatcher() *dispatcher {
	return &dispatcher{
		dispatchMap: map[string][]Handler{},
		processChan: make(chan *processMeta),
	}
}

func (d *dispatcher) runProcesser() {
	go func() {
		for {
			select {
			case meta := <-d.processChan:
				meta.handler(meta.ctx)
			}
		}
	}()
}

func (d *dispatcher) addTask(evt string, handler Handler) {
	d.dispatchMap[evt] = append(d.dispatchMap[evt], handler)
}

func (d *dispatcher) dispatch(evt string, data interface{}) {
	ctx := d.makeContext(evt, data)
	handlers, ok := d.dispatchMap["evt"]
	if !ok {
		fmt.Println("创建分发任务失败，该事件不存在")
		return
	}
	for _, handler := range handlers {
		d.processChan <- &processMeta{
			handler: handler,
			ctx:     ctx,
		}
	}

}

func (d *dispatcher) makeContext(evt string, data interface{}) *Context {
	return &Context{
		Event:     evt,
		Timestamp: time.Now().UnixMilli(),
		Data:      data,
	}
}
