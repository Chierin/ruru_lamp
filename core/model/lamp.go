package model

type Lamp struct {
	Base
	LiveID    int64
	Content   string
	Nickname  string
	Timestamp int64
	No        int64
}
