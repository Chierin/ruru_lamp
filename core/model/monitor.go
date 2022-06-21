package model

type Monitor struct {
	Base
	RoomID   int64
	UID      int64
	UName    string
	LiveTime int64
	IsLiving bool
}
