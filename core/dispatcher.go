package core

type Handler func(ctx *Context)

type Dispatcher struct {
	dispatchMap map[string]Handler
}

func (d *Dispatcher) addTask() {

}

func (d *Dispatcher) dispatch() {

}
