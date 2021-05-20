package main

import (
	"fmt"
	"time"
)

func main()  {
	type UserLoginLog struct {
		uid uint64
		loginTime int64
		logoutTime int64
	}
	type CountGroup struct {
		Count int
		startTime int64
		endTime int64
	}
	var userLog UserLoginLog
	userLog.uid = 11233
	userLog.loginTime = 1618935935
	userLog.logoutTime = 234234234
	var logList []UserLoginLog
	logList = append(logList,userLog)
	thisdate := "2006-01-02 15:04:05"
	lasttime := time.Now().Format(thisdate)
	fmt.Println(lasttime)
	groupCountMap := make(map[int]CountGroup)
	for i:=0 ; i < 24; i++ {
		str := ""
		if i < 10 {
			str = fmt.Sprintf("0%d",i)
		}else{
			str = fmt.Sprintf("%d",i)
		}
		dateStr := fmt.Sprintf("2021-04-21 %s:59:59",str)
		t := getHourByStr(dateStr)
		groupCountMap[i] = CountGroup{
			Count: 0,
			startTime: t,
			endTime: t,
		}
	}
	for _,logInfo := range logList{
		fmt.Println(logInfo.logoutTime)
		logoutHour := getHourByTime(logInfo.logoutTime)
		countGroup := groupCountMap[logoutHour]
		if logInfo.loginTime < countGroup.startTime && logInfo.logoutTime >= countGroup.endTime  {
			countGroup.Count++
			groupCountMap[logoutHour] = countGroup
		}
	}
}
func getHourByTime(t int64) int{
	tm := time.Unix(t,0)
	fmt.Println(tm.Hour())
	return tm.Hour()
}
func getHourByStr(datestr string) int64{
	thisdate := "2006-01-02 15:04:05"
	//str1 := "2021-04-21 01:00:00"
	loc, _ := time.LoadLocation("Asia/Shanghai")
	tt,_ := time.ParseInLocation(thisdate,datestr,loc)
	fmt.Println(tt)
	fmt.Println(tt.Hour())
	return tt.Unix()
}
