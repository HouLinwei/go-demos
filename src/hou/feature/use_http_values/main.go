package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"time"
)

type Item struct {
	StockID string `json:"stock_id"`
	Time    int64  `json:"time"`
}

func makeItem() Item {
	return Item{
		StockID: "test",
		Time:    time.Now().Unix(),
	}
}

type Report struct {
	APPID    string `json:"appid"`
	RC_ITEMS []Item `json:"rc_items"`
}

func makeReport() Report {
	return Report{
		APPID: "appid",
	}
}

const REPORT_API = "http://www.baidu.com/"

func (r Report) transfer() url.Values {
	req, _ := http.NewRequest("GET", REPORT_API, nil)
	q := req.URL.Query()
	t := reflect.TypeOf(r)
	v := reflect.ValueOf(r)
	for i := 0; i < t.NumField(); i++ {
		value := v.Field(i).Interface()
		switch value.(type) {
		case string:
			v := value.(string)
			tag := t.Field(i).Tag.Get("http")
			necessary := t.Field(i).Tag.Get("necessary")
			if v == "" && necessary == "" {
				continue
			}
			if _, ok := q[tag]; ok {
				q.Add(tag, v)
			} else {
				q.Set(tag, v)
			}
		case []Item:
			fmt.Println("")
			v := value.([]Item)
			fmt.Println(v)
			data, _ := json.Marshal(v)
			fmt.Println(data)
			fmt.Println(string(data))
		}
	}
	return q
}

func fullReport(report Report, items ...Item) Report {
	for _, v := range items {
		report.RC_ITEMS = append(report.RC_ITEMS, v)
	}
	return report
}

func main() {
	r := makeReport()
	i1 := makeItem()
	i2 := makeItem()
	r0 := fullReport(r, i1, i2)
	fmt.Println(r0)
	r0.transfer()
}
