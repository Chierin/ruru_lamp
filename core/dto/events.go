package dto

type LiveStart struct {
	RoomID    int64
	UserID    int64
	Username  string
	Title     string
	StartTime int64
}

type LiveStop struct {
	RoomID    int64
	UserID    int64
	Username  string
	Title     string
	StartTime int64
	StopTime  int64
}

type LampUp struct {
	RoomID     int64
	UserID     int64
	Username   string
	Title      string
	StartTime  int64
	LampUpTime int64
}

func (l *LampUp) LiveDuration() int64 {
	return l.LampUpTime - l.StartTime
}

type TitleChange struct {
	RoomID     int64
	UserID     int64
	Username   string
	Title      string
	StartTime  int64
	ChangeTime int64
}
