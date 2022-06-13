package core

import (
	"bytes"
	"io/ioutil"
	"net/http"

	jsoniter "github.com/json-iterator/go"
)

type Res struct {
	Code    int                  `json:"code"`
	Msg     string               `json:"msg"`
	Message string               `json:"message"`
	Data    map[string]LiveState `json:"data"`
}
type LiveState struct {
	Title            string `json:"title"`
	RoomID           int    `json:"room_id"`
	UID              int    `json:"uid"`
	Online           int    `json:"online"`
	LiveTime         int    `json:"live_time"`
	LiveStatus       int    `json:"live_status"`
	ShortID          int    `json:"short_id"`
	Area             int    `json:"area"`
	AreaName         string `json:"area_name"`
	AreaV2ID         int    `json:"area_v2_id"`
	AreaV2Name       string `json:"area_v2_name"`
	AreaV2ParentName string `json:"area_v2_parent_name"`
	AreaV2ParentID   int    `json:"area_v2_parent_id"`
	Uname            string `json:"uname"`
	Face             string `json:"face"`
	TagName          string `json:"tag_name"`
	Tags             string `json:"tags"`
	CoverFromUser    string `json:"cover_from_user"`
	Keyframe         string `json:"keyframe"`
	LockTill         string `json:"lock_till"`
	HiddenTill       string `json:"hidden_till"`
	BroadcastType    int    `json:"broadcast_type"`
}

func GetLiveStateBatch(uids []int64) (map[string]LiveState, error) {
	url := "https://api.live.bilibili.com/room/v1/Room/get_status_info_by_uids"
	payload, _ := jsoniter.Marshal(map[string]interface{}{"uids": uids})
	r, err := http.Post(url, "application/json", bytes.NewReader(payload))
	if err != nil {
		return nil, ErrFetchLiveState
	}
	b, _ := ioutil.ReadAll(r.Body)
	var res Res
	err = jsoniter.Unmarshal(b, &res)
	if err != nil {
		return nil, ErrFetchLiveState
	}
	if res.Code != 0 {
		return nil, ErrFetchLiveState
	}
	return res.Data, nil
}
