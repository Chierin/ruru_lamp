package model

type Comment struct {
	Base
	Nickname string
	LampID   int64
	Content  string
}
