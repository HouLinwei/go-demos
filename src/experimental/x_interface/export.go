package x_interface

import (
	"fmt"
	"net/http"
)

type KuaiSortedSet interface {
	GetMember() string
	GetScore() float64
	SetMember(string) (bool, error)
	SetScore(float64)(bool, error)
}

// 多设备同步播放记录
type VideoWithTimestamp struct {
	Vid       string
	Timestamp int64
}

type WatchRecordAddReq struct {
	Records []VideoWithTimestamp `json:"records"`
}

func(v *VideoWithTimestamp)GetMember()(ret string){
	ret = v.Vid
	return
}

func(v *VideoWithTimestamp)GetScore()(ret float64){
	ret = float64(v.Timestamp)
	return
}

func(v *VideoWithTimestamp)SetMember(s string)(bool, error){
	v.Vid = s
	return true, nil
}

func(v *VideoWithTimestamp)SetScore(f float64)(bool, error){
	v.Timestamp = int64(f)
	return true, nil
}


func Run(){
	res := []string{"1","2","3"}
	var ret []KuaiSortedSet
	for idx, r := range res {
		var t KuaiSortedSet
		t.SetMember(r)
		t.SetScore(float64(idx))
		ret = append(ret, t)
	}

	for _, v:= range ret{
		fmt.Println(v.GetScore())
		fmt.Println(v.GetMember())
	}
}
