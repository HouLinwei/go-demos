package main

import (
	"github.com/emirpasic/gods/utils"
	"testing"
	"unicode"
	"fmt"
	"net"
	"time"
	"encoding/json"
)

type Type uint
type Rule struct {

	// 任务配置
	// 不同的任务类型都需要限制发放的金币数目。对于随机类型的金币下发，此值为最大数；固定数目发放类型的任务，此值为应发放的金币数目。
	MaxCoin int `json:"max_coin" bson:"max_coin"`

	// 对于签到类型的任务，连续签到2、7、14、21、30天
	Period int `json:"period,omitempty" bson:"period,omitempty"`

	// 评论奖励任务：需要限制每天（5）和总共（15）的奖励次数。
	// 阅读 push 奖励任务：每天的最大奖励次数等于当日push个数，且不需要最大限制。
	// 观看视频奖励任务：限制总共奖励次数20次。
	DailyLimit int `json:"daily_limit,omitempty" bson:"daily_limit,omitempty"`
	TotalLimit int `json:"total_limit ,omitempty" bson:"total_limit,omitempty"`
	// 对于视频播放奖励，需要确定需要每个视频的观看市场比例
	WatchTimeRate float64 `json:"watch_time_rate,omitempty" bson:"watch_time_rate,omitempty"`
}

type Detail struct {
	ID       string    `json:"id" bson:"_id"`
	Status   int       `json:"status" bson:"status"` // 0 disable 1 enable
	Type     Type      `json:"type" bson:"type"`
	Desc     string    `json:"desc" bson:"desc"`
	Rule     Rule      `json:"rule" bson:"rule"`
	UpdateAt time.Time `json:"update_at" bson:"update_at"`
}

type UserChannel string
type UserKind string //用户种类
type UserID struct {
	Channel UserChannel
	ID      string
	Kind    UserKind
}




type UserAction struct {
	User UserID `json:"user" bson:"user"`
	Ts   int64        `json:"ts" bson:"ts"`
	Type string       `json:"type" bson:"type"`
	// vary from different user action.
	//ItemID string `json:"item_id" bson:"item_id"`
	Vid       string `json:"vid,omitempty" bson:"vid,omitempty"`
	PushID    string `json:"push_id,omitempty" bson:"push_id,omitempty"`
	CommentID string `json:"comment_id,omitempty" bson:"comment_id,omitempty"`
}

func main(){
	ts1 := time.Now().Unix()
	ts2 := time.Now().Unix()

	ttt2 := time.Unix(ts2, 0)

	ttt1 := time.Unix(ts1, 0)
	fmt.Println(ttt1.Date())
	fmt.Println(ttt2.Date())
	fmt.Println(HasTheSameDate(ts1, ts2))

	a :=[]int{1,2,3,4}
	l := 4
	fmt.Println(a[0:l-1])

	F1()
}


func F1()error{
	defer func() {
		fmt.Println("In f1")
	}()
	return F2()
}

func F2()error{
	fmt.Println("IN F2")
	return nil
}



func HasTheSameDate(ts1, ts2 int64)bool{
	t1 := time.Unix(ts1, 0)
	t2 := time.Unix(ts2, 0)
	y1, m1, d1 := t1.Date()
	y2, m2, d2 := t2.Date()
	if y1 == y2 && m1 == m2 && d1== d2{
		return true
	}
	return false
}

func main1()  {
	//a := "12"
	//for _, v := range a{
	//	fmt.Println(v, reflect.TypeOf(v))
	//	fmt.Println(v-48, reflect.TypeOf(v-48))
	//	fmt.Println(string(v))
	//}
	d := Detail{
		Status  :1,
		Type    :0,
		Desc    :"test task",
		Rule   : Rule  {
			MaxCoin:19,

			// 对于签到类型的任务，连续签到2、7、14、21、30天
			Period:7,

			// 评论奖励任务：需要限制每天（5）和总共（15）的奖励次数。
			// 阅读 push 奖励任务：每天的最大奖励次数等于当日push个数，且不需要最大限制。
			// 观看视频奖励任务：限制总共奖励次数20次。
			DailyLimit :23,
			TotalLimit :100,
			// 对于视频播放奖励，需要确定需要每个视频的观看市场比例
			WatchTimeRate :20.0,
		},
		UpdateAt :time.Now(),
	}
	data,_ := json.Marshal(d)
	fmt.Println(string(data))

	ua := UserAction{
		User:UserID {
			Channel:"1",
			ID     :"1",
			Kind   :"1",
		},
		Ts :1121111111,
		Type :"1",
		// vary from different user action.
		//ItemID string `json:"item_id" bson:"item_id"`
		Vid   :"j-fsxxx",
		PushID:"xxx",
		CommentID :"xxxx",
	}
	data2,_ := json.Marshal(ua)
	fmt.Println(string(data2))
}

func main0() {
	//// todo
	////s := `eyJIIjp7Im1vZGVsIjoiTUkgNExURSIsImFuZHJvaWQiOiI2LjAuMSIsIm1pdWkiOiI4LjEuMjUiLCJibiI6IkQiLCJwcm9kdWN0IjoiY2FuY3JvX3djX2x0ZSIsImRldmljZSI6ImNhbmNybyIsImN2IjoiMi40LjEiLCJsYW5nIjoiemhfQ04iLCJyZWdpb24iOiJDTiIsIm1pIjoiMCIsInNlbmRlciI6ImNvbS5taXVpLmFuYWx5dGljcyIsInVzZXJpZCI6MCwiYWFpZCI6IjhmYjFkM2E5LTE3ZWEtNDcxZC04MWM3LWI1MGRlMTAwNDg0YyIsIm4iOiIxMCIsIm1hYyI6IjRhZjcwZjM3ZWM4ZWZkYmE0OWNjYzNhY2RhNzI4NTE0IiwiaW1laSI6IjFlYjJlYjQ3YzhhMjIzMDJmZGVjYWFlZGI4N2VmYzU4Iiwic3QiOiIxNTE3NTYwODAzMTA5In0sIkIiOlt7IkgiOnsic2lkIjoiIiwicGsiOiJjb20ua3VhaWVzdC52aWRlbyIsImtleSI6Imt1YWllc3RfYmlzbG9nc3RhZ2luZyIsImV2ZW50VGltZSI6IjE1MTc1NjA4MDMwNTAiLCJzbiI6Ijk1Nzc2MjgzMjI0MTEiLCJyZXRyeUNudCI6MH0sIkIiOnsiX2NhdGVnb3J5XyI6ImFkIiwiX2FjdGlvbl8iOiJzZWFyY2hfaG90X2l0ZW1fY2xpY2siLCJjb21tb24iOnsiX3JlcyI6ImhkMTA4MCIsIl90b2tlbiI6Im8xU2RiSFlWelF6Y0o4MEVEM3k3dktMTGxGV05yTlc3SENJYjVZZVRMN2hkYlJvU1NmNXBrU3JPMGlMVVozM0VLcUt1TXhvUmZtakJPS3JYVTNXc3ZTemx6NmZUazVZQ3ZpaWlEMzg3eEFHNWNWMXAycEZFaTdhUVA1aEJSWlAwU1NxUUotX3VKR2cxY2RCckU0U0gtbEhlNks0XzZscW5Nd19MRnRaRWpIR3VybWs5TmxhVk5uY0JhMzdjcjBGLUpsV01lQ2s0MTJaT0ctMlVMVG1EZ0E9PSIsIl9sYXRpdHVkZSI6Ik5EQXVNREkxTURRelxuIiwiX25ldHR5cGUiOiIxIiwiX2xvY2FsIjoiemhfQ04iLCJfdmVyc2lvbiI6ImRldiIsIl9taXVpIjoiVjkiLCJfdHMiOiIxNTE3NTYwODAzIiwiX3ZlciI6IjEuOS4wIiwiX2FuZHZlciI6IjIzIiwiX21vZGVsIjoiTUkgNExURSIsIl9jaGFubmVsIjoic3RhbmRhcmQiLCJfZGV2dHlwZSI6IjEiLCJfYXBwdmVyIjoiMjcxMzYwIiwiX21pdWl2ZXIiOiI4LjEuMjUiLCJpbWVpIjoiMWViMmViNDdjOGEyMjMwMmZkZWNhYWVkYjg3ZWZjNTgiLCJfcmVmIjoibWluaXZpZGVvIiwiX2RldmlkIjoiOTlmNzJjOTczYTVkNmQ1NjFiODY2ODU3ZTgxOWQ2NmMiLCJfaXAiOiIxMC4yMzUuMjAyLjE4MCIsIl9sb25naXR1ZGUiOiJNVEUyTGpNeU56YzFOQT09XG4iLCJfcGx5dmVyIjoiMjAxNjEyMTIifSwiYWN0aW9uIjoic2VhcmNoX2hvdF9pdGVtX2NsaWNrIiwiY2F0ZWdvcnkiOiJzZWFyY2giLCJleHQiOiJ7XCJuYW1lXCI6XCLlsI_kvLbnjqnlhbdcIixcInRhZ1wiOlwiXCIsXCJwb3NpdGlvblwiOjcsXCJ0aW1lc3RhbXBcIjpcIjE1MTc1NjA4MDMwNTBcIn0ifX1dfQ==`
	//s := `eyJIIjp7Im1vZGVsIjoiTUkgNExURSIsImFuZHJvaWQiOiI2LjAuMSIsIm1pdWkiOiI4LjEuMjUiLCJibiI6IkQiLCJwcm9kdWN0IjoiY2FuY3JvX3djX2x0ZSIsImRldmljZSI6ImNhbmNybyIsImN2IjoiMi40LjEiLCJsYW5nIjoiemhfQ04iLCJyZWdpb24iOiJDTiIsIm1pIjoiMCIsInNlbmRlciI6ImNvbS5taXVpLmFuYWx5dGljcyIsInVzZXJpZCI6MCwiYWFpZCI6IjhmYjFkM2E5LTE3ZWEtNDcxZC04MWM3LWI1MGRlMTAwNDg0YyIsIm4iOiIxMCIsIm1hYyI6IjRhZjcwZjM3ZWM4ZWZkYmE0OWNjYzNhY2RhNzI4NTE0IiwiaW1laSI6IjFlYjJlYjQ3YzhhMjIzMDJmZGVjYWFlZGI4N2VmYzU4Iiwic3QiOiIxNTE3NTYwODAzMTA4In0sIkIiOlt7IkgiOnsic2lkIjoiIiwicGsiOiJjb20ua3VhaWVzdC52aWRlbyIsImtleSI6Imt1YWllc3RfYmlzbG9nc3RhZ2luZyIsImV2ZW50VGltZSI6IjE1MTc1NjA4MDMwNjgiLCJzbiI6Ijk1Nzc2MjgzMjI0MTIiLCJyZXRyeUNudCI6MH0sIkIiOnsiX2NhdGVnb3J5XyI6ImFkIiwiX2FjdGlvbl8iOiJzZWFyY2giLCJjb21tb24iOnsiX3JlcyI6ImhkMTA4MCIsIl90b2tlbiI6Im8xU2RiSFlWelF6Y0o4MEVEM3k3dktMTGxGV05yTlc3SENJYjVZZVRMN2hkYlJvU1NmNXBrU3JPMGlMVVozM0VLcUt1TXhvUmZtakJPS3JYVTNXc3ZTemx6NmZUazVZQ3ZpaWlEMzg3eEFHNWNWMXAycEZFaTdhUVA1aEJSWlAwU1NxUUotX3VKR2cxY2RCckU0U0gtbEhlNks0XzZscW5Nd19MRnRaRWpIR3VybWs5TmxhVk5uY0JhMzdjcjBGLUpsV01lQ2s0MTJaT0ctMlVMVG1EZ0E9PSIsIl9sYXRpdHVkZSI6Ik5EQXVNREkxTURRelxuIiwiX25ldHR5cGUiOiIxIiwiX2xvY2FsIjoiemhfQ04iLCJfdmVyc2lvbiI6ImRldiIsIl9taXVpIjoiVjkiLCJfdHMiOiIxNTE3NTYwODAzIiwiX3ZlciI6IjEuOS4wIiwiX2FuZHZlciI6IjIzIiwiX21vZGVsIjoiTUkgNExURSIsIl9jaGFubmVsIjoic3RhbmRhcmQiLCJfZGV2dHlwZSI6IjEiLCJfYXBwdmVyIjoiMjcxMzYwIiwiX21pdWl2ZXIiOiI4LjEuMjUiLCJpbWVpIjoiMWViMmViNDdjOGEyMjMwMmZkZWNhYWVkYjg3ZWZjNTgiLCJfcmVmIjoibWluaXZpZGVvIiwiX2RldmlkIjoiOTlmNzJjOTczYTVkNmQ1NjFiODY2ODU3ZTgxOWQ2NmMiLCJfaXAiOiIxMC4yMzUuMjAyLjE4MCIsIl9sb25naXR1ZGUiOiJNVEUyTGpNeU56YzFOQT09XG4iLCJfcGx5dmVyIjoiMjAxNjEyMTIifSwiYWN0aW9uIjoic2VhcmNoIiwiY2F0ZWdvcnkiOiJ1c2VyIiwiZXh0Ijoie1wia2V5d29yZFwiOlwi5bCP5Ly2546p5YW3XCIsXCJob3RcIjpcInRydWVcIixcInVybFwiOlwiU2VhcmNoUGFnZVwiLFwidGltZXN0YW1wXCI6XCIxNTE3NTYwODAzMDY4XCJ9In19XX0=`
	//fmt.Printf("%x", s)
	//
	//n1, err := base64.StdEncoding.DecodeString(s)
	//if err != nil {
	//	fmt.Println("xxx1",err)
	//}
	//fmt.Println("--->1", string(n1))
	//
	//
	////n2, err := base64.RawStdEncoding.DecodeString(s)
	////if err != nil {
	////	fmt.Println("xxx2",err)
	////}
	////fmt.Println(string(n2))
	//
	//
	//n3, err := base64.URLEncoding.DecodeString(s)
	//if err != nil {
	//	fmt.Println("xxx3",err)
	//}
	//fmt.Println("--->3", string(n3))
	//
	//
	////n4, err := base64.RawURLEncoding.DecodeString(s)
	////if err != nil {
	////	fmt.Println("xxx4",err)
	////}
	////fmt.Println(string(n4))
	//
	//m := make(map[string]string)
	//m = nil
	//fmt.Println(m == nil)
	//fmt.Println(len(m)==0)

	s := []string{"你好123", "ただいま", "사랑해요"}
	r := []int{4, 4, 4}
	for i, v := range s {
		if LenTitle(v) != r[i] {
			fmt.Println(v, "want", r[i], "get", LenTitle(v))
		}
	}


	p := 7
	hh := Float64Ptr(float64(p))
	fmt.Println(hh)
	fmt.Println(*hh)

	ip := "127.0.0.1"
	iip := net.ParseIP(ip)
	fmt.Println([]byte(iip))
}


func Float64Ptr(v float64) *float64 { return &v }

func TestLenTitle(t *testing.T) {

}


type M struct {
	Age int
	Name string
	P int
}

func Comarator(a, b interface{})int{
	m1 := a.(M)
	m2 := b.(M)
	if m1.Age != m2.Age{
		return 0 - utils.IntComparator(m1.Age, m2.Age)
	}else{
		return utils.StringComparator(m1.Name, m2.Name)
	}
}

func LenTitle(s string) int {
	var l int
	// CJK 字符算一个。其他字符算半个。
	for _, v := range s {
		// 汉字字符
		if unicode.Is(unicode.Scripts["Han"], v) {
			// 汉字
			l += 2
			// 韩语字符
		} else if unicode.Is(unicode.Scripts["Hangul"], v) {
			l += 2
			// 日语字符
		} else if unicode.Is(unicode.Scripts["Hiragana"], v) {
			// 平假字符
			l += 2
		} else if unicode.Is(unicode.Scripts["Katakana"], v) {
			// 片假字符
			l += 2
		} else {
			l += 1
		}
	}
	if l%2 != 0{
		return l / 2 +1
	}else{
		return l/2
	}
}
