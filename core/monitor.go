package core

import (
	"ruru_lamp/core/model"
	"strconv"
	"time"
)

type State struct {
	UserID     int64
	RoomID     int64
	Username   string
	StartTime  int64
	EndTime    int64
	IsLive     bool
	Title      string
	Cover      string
	UpdateTime int64
	GetTime    int64
}

func (s *State) LiveDuration() int64 {
	if !s.IsLive {
		return 0
	}
	return s.GetTime - s.StartTime
}

type StateCache struct {
	states map[int64]State
}

func NewStateCache() *StateCache {
	return &StateCache{states: map[int64]State{}}
}

func (c *StateCache) Get(userID int64) *State {
	s, ok := c.states[userID]
	if !ok {
		Logger.Warnf("获取用户【%d】状态失败", userID)
		return nil
	}
	s.GetTime = time.Now().UnixMilli()
	return &s
}

func (c *StateCache) UIDs() []int64 {
	keys := make([]int64, 0, len(c.states))
	for k := range c.states {
		keys = append(keys, k)
	}
	return keys
}

func (c *StateCache) Update(state State) {
	c.states[state.UserID] = state
}

type Monitor struct {
	cache  *StateCache
	pubber *Pubber
}

func NewMonitor() *Monitor {
	monitor := Monitor{
		cache:  NewStateCache(),
		pubber: NewPubber(),
	}
	return &monitor
}

func (m *Monitor) Handle(evt string, handler Handler) {
	m.pubber.AddSub(evt, handler)
}

func (m *Monitor) AddUser(userID int64) bool {
	liveStates, _ := GetLiveStateBatch([]int64{userID})
	liveState, ok := liveStates[strconv.FormatInt(userID, 10)]
	if !ok {
		Logger.Warnf("在添加用户【%d】时获取用户信息失败", userID)
		return false
	}
	// 纯添加用户记录，后面统一获取用户直播状态
	m.cache.Update(State{
		UserID:     userID,
		RoomID:     int64(liveState.RoomID),
		Username:   liveState.Uname,
		StartTime:  0,
		EndTime:    0,
		IsLive:     false,
		Title:      "",
		Cover:      "",
		UpdateTime: 0,
		GetTime:    0,
	})
	return true
}

func (m *Monitor) Start() {
	m.pubber.Start()
	go func() {
		for range time.Tick(time.Second * 5) {
			m.UpdateState()
		}
	}()
}

func (m *Monitor) UpdateState() {
	liveStates, err := GetLiveStateBatch(m.cache.UIDs())
	if err != nil {
		Logger.Warn("获取live信息失败")
	}
	for _, liveState := range liveStates {
		oldState := m.cache.Get(int64(liveState.UID))
		newState := tranLiveStateToState(&liveState)
		if oldState.IsLive == false && newState.IsLive == true {
			// 开播
			m.onLiveStart(oldState, newState)
		} else if oldState.IsLive == true && newState.IsLive == false {
			// 下播
			m.onLiveStop(oldState, newState)
		} else if oldState.IsLive == newState.IsLive == true && oldState.StartTime != newState.StartTime {
			// 反复横跳，光速下播又开播
			m.onLiveStop(oldState, newState)
			m.onLiveStart(oldState, newState)
		}

	}
}

func (m *Monitor) onLiveStart(oldState *State, newState *State) {
	model.AddLive(db, newState.UserID, newState.StartTime, newState.Title)
	m.cache.Update(*newState)
	m.pubber.Pub(EventLiveStart, newState)
}

func (m *Monitor) onLiveStop(oldState *State, newState *State) {
	oldState.EndTime = time.Now().Unix()
	oldState.IsLive = false
	model.UpdateLive(db, oldState.UserID, oldState.EndTime)
	m.cache.Update(*oldState)
	m.pubber.Pub(EventLiveStop, oldState)
}

func tranLiveStateToState(liveState *LiveState) *State {
	var isLive bool
	if liveState.LiveStatus == 1 {
		isLive = true
	} else {
		isLive = false
	}
	state := State{
		UserID:     int64(liveState.UID),
		RoomID:     int64(liveState.RoomID),
		Username:   liveState.Uname,
		StartTime:  int64(liveState.LiveTime),
		EndTime:    0,
		IsLive:     isLive,
		Title:      liveState.Title,
		Cover:      liveState.CoverFromUser,
		UpdateTime: time.Now().UnixMilli(),
		GetTime:    time.Now().UnixMilli(),
	}
	return &state
}
