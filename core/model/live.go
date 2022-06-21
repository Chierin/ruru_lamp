package model

type Live struct {
	Base
	MonitorID int64
	StartTime int64
	StopTime  int64
}
