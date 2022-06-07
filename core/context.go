package core

import "ruru_lamp/core/dto"

type Context struct {
	Event     string
	Timestamp int64
	Data      interface{}
}

func (c *Context) AsLiveStart() *dto.LiveStart {
	return c.Data.(*dto.LiveStart)
}

func (c *Context) AsLiveStop() *dto.LiveStop {
	return c.Data.(*dto.LiveStop)
}

func (c *Context) AsLampUp() *dto.LampUp {
	return c.Data.(*dto.LampUp)
}

func (c *Context) AsTitleChange() *dto.TitleChange {
	return c.Data.(*dto.TitleChange)
}
