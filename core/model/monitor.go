package model

type Monitor struct {
	RoomID    int64
	UID       int64
	UName     string
	LiveTime  int64
	IsLiving  bool
	CreatedAt int64 `gorm:"autoCreateTime:milli"`
	UpdatedAt int64 `gorm:"autoUpdateTime:milli"`
}
