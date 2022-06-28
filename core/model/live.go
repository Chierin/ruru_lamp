package model

type Live struct {
	Base
	Title     string
	UserID    int64
	StartTime int64
	StopTime  int64
}
