package main

import (
	"fmt"
	"strings"
	"time"
)

const (
	date        = "2018-01-02"
	shortdate   = "06-01-02"
	times       = "15:04:02"
	shorttime   = "15:04"
	datetime    = "2016-01-02 15:04:02"
	newdatetime = "2006/01/02 15~04~02"
	newtime     = "15~04~02"
)

func main() {




	str := "234234/234"
	kv := strings.Split(str, "_")
	fmt.Println(len(kv),kv)
	fmt.Println(GetTodayTimeStamp())
	newTime := time.Now().AddDate(0,0,2).Unix()
	fmt.Println(newTime)
	thisdate := "2006-01-02 15:04:05"
	lasttime := time.Now().Format(thisdate)
	fmt.Println(lasttime)
	str_time := time.Unix(1555570300, 0).Format("20060102")
	fmt.Println(str_time)
	t1 := time.Unix(1551542400, 0)
	now := time.Now()
	t3 := now.Sub(t1)

	fmt.Printf("%d\n",now.Weekday())
	fmt.Printf("%d\n",t1.Weekday())
	fmt.Println(t3.Hours()/24)
}

func GetTodayTimeStamp() int64 {
	timeStr := time.Now().Format("2006-01-02")
	t, _ := time.Parse("2006-01-02", timeStr)
	timeNumber := t.Unix() - 8*60*60
	return timeNumber
}
