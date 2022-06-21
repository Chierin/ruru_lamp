package model

import "gorm.io/gorm"

// 添加监控
func AddMonitor(tx *gorm.DB, UID, RoomID int64, UName string) {
	m := &Monitor{
		RoomID:   RoomID,
		UID:      UID,
		UName:    UName,
		LiveTime: 0,
		IsLiving: false,
	}
	tx.Save(m)
}

// 移除监控
func RemoveMonitor(tx *gorm.DB, UID int64) {
	tx.Where("uid = ?", UID).Delete(&Monitor{})
}

// 添加live记录
func AddLive(tx *gorm.DB, monitorID, startTime int64) {
	l := &Live{
		MonitorID: monitorID,
		StartTime: startTime,
		StopTime:  0,
	}
	tx.Save(l)
}

// 设置直播结束时间
func SetLiveStopTime(tx *gorm.DB, liveID, stopTime int64) {
	tx.Model(&Live{}).Where("id = ?", liveID).Update("stop_time", stopTime)
}

// 点灯
func AddLamp(tx *gorm.DB, liveID int64, nickname, content string, timestamp int64) {
	l := &Lamp{
		LiveID:    liveID,
		Content:   content,
		Nickname:  nickname,
		Timestamp: timestamp,
	}
	tx.Save(l)
}

// 更新灯
func UpdateLamp(tx *gorm.DB, lampID int64, lamp *Lamp) {
	tx.Where("id = ?", lamp.ID).Updates(lamp)
}

func AddComment()
