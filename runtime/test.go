package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"icode.baidu.com/baidu/netdisk/nd-go-taskscore/utils"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)
var SecretKey = "ae82c240578eb391de93c2f4c3dfc3ba"

//总是返回大的那个参数序号
func compareVersionNum(arr1 []string,arr2 []string) int {

	len1  := len(arr1)
	len2 := len(arr2)
	var arr_short []string
	var arr_long []string

	short_name := int(0)
	long_name := int(0)
	if len1 < len2{
		arr_short = arr1
		arr_long = arr2
		short_name = 1
		long_name = 2
	}else{
		arr_short = arr2
		arr_long = arr1
		short_name = 2
		long_name = 1
	}
	for key,v :=range arr_long {
		if key > len(arr_short)-1 {
			arr_short = append(arr_short,"0")
		}
		v_long ,_ := strconv.Atoi(v)
		v_short ,_ := strconv.Atoi(arr_short[key])

		if v_long == v_short {
			continue
		}else if v_long > v_short {
			return long_name
		}else {
			return short_name
		}
	}
	return 0
}
func removeUids(uid uint64,uids []uint64) []uint64 {

	var newUids  []uint64
	for _,diffUid := range uids {
		if uid == diffUid{
			continue
		}
		newUids = append(newUids,diffUid)
	}
	return newUids
}
func main() {


	GetHtml()

	params := url.Values{}
	params.Set("size","444")
	params.Set("transaction_id","44424434343434")
	params.Set("reason_code","444")
	//body, _ := json.Marshal(params)
	body := params.Encode()
	fmt.Println(body)
	str := Krand(4,1)
	fmt.Println(string(str))
	str4 := GetRandomString(4)
	fmt.Println(str4)
	return
	uid := uint64(33334)
	var uids []uint64
	uids = append(uids,uint64(1233),uint64(33334),uint64(333334))
	fmt.Println(uids)
	newuids := removeUids(uid,uids)
	fmt.Println(newuids)
	return
	n := int(0)
	m := int32(23444)
	n = int(m)
	fmt.Println(n)

	s1 := "139169912"
	s2 := string([]byte(s1)[:4])
	fmt.Println(s2) //得到 "abc"
	s3 := "746851555331"
	//s4 := string([]byte(s3)[:8])
	//fmt.Println(s4) //得到 "abc"
	length := len(s3)
	s4 := string([]byte(s3)[2:length])
	fmt.Println(s4)
	s5 := s2+s3

	fmt.Println(s5)
	return
	str1 := "10.0.60"
	str2 := "9.0.9"

	arr1 := strings.Split(str1,".")
	arr2 := strings.Split(str2,".")
	s := compareVersionNum(arr1,arr2)

	fmt.Println(s)
	type Fav struct {
		Content string
		Title string
	}
	var fav Fav
	fav.Content = "teswt"
	fav.Title = "<h1>sdf</h1>"
	content := fav.Title+fav.Content
	fmt.Println(content)
	noteIds :=[]int64{1,234,55}
	urls := []string{"123","234","sdf"}
	var extraList []map[string]string
	extraMap :=make(map[string]string)
	for key,noteId := range noteIds{
		extraMap["uid"] = ""
		extraMap["note_id"] = strconv.FormatInt(noteId,10)
		extraMap["url"] = urls[key]
		extraList = append(extraList,extraMap)
	}
	fmt.Println(extraList)
	extraJson, _:= json.Marshal(extraList)
	fmt.Println(string(extraJson))

	// just one second
	t1 :="2017-01-02"
	t2 := "2017-02-03"
	res := utils.TimeSubDay(t1, t2)

	fmt.Println(res)

	TaskID := 26864741842090
	uk := "12"
	rand := "133"
	time := "113"

	paramsStr := fmt.Sprintf("%d_%s_%s_%s_%s", TaskID, uk, rand, time,SecretKey)
	fmt.Println(paramsStr)
	rc4Token := utils.Md5(paramsStr)
	fmt.Println(rc4Token)
	//url :="http://yq01-luohj.epc.baidu.com:8086/note/cutcallback"
	//params := `{


	//HttpPost(url,params)
}

func timeSub(tstr1 string, tstr2 string) int {
	layout := "2006-01-02"
	t1, _ := time.Parse(layout, tstr1)
	t2, _ := time.Parse(layout, tstr2)

	t1 = time.Date(t1.Year(), t1.Month(), t1.Day(), 0, 0, 0, 0, time.Local)
	t2 = time.Date(t2.Year(), t2.Month(), t2.Day(), 0, 0, 0, 0, time.Local)

	return int(t1.Sub(t2).Hours() / 24)
}

func HttpPost(url string,params string)  {
	json := bytes.NewBuffer([]byte(params))
	req, err := http.NewRequest("POST", url, json)
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)

}

//size 长度，
// kind 为种类
//    0  // 纯数字
//    1  // 小写字母
//    2  // 大写字母
//    3  // 数字、大小写字母
func Krand(size int, kind int) []byte {
	ikind, kinds, result := kind, [][]int{[]int{10, 48}, []int{26, 97}, []int{26, 65}}, make([]byte, size)
	is_all := kind > 2 || kind < 0
	rand.Seed(time.Now().UnixNano())
	for i :=0; i < size; i++ {
		if is_all {
			ikind = rand.Intn(3)
		}
		scope, base := kinds[ikind][0], kinds[ikind][1]
		result[i] = uint8(base+rand.Intn(scope))
	}
	return result
}


func  GetRandomString(l int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyz"
	randBytes := []byte(str)
	var result []byte
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, randBytes[r.Intn(len(randBytes))])
	}
	return string(result)
}
func GetHtml()  {
	text := `Hello 世界！123 Go.<div class="custom_files"><div data-fsid="123123"></div><div data-fsid="1231233"></div><div data-fsid="12312443"></div></div>`
	reg := regexp.MustCompile(`data-fsid="([\d]+)"`)  // 查找连续的小写字母

	var fsIdList []int64
	//fmt.Printf("%q\n", reg.FindAllString(text,-1))
	result := reg.FindAllStringSubmatch(text, -1)
	for _,res := range result {
		fsId, _ := strconv.ParseInt(res[1], 10, 64)
		if fsId < 1{
			continue
		}
		fsIdList = append(fsIdList,fsId)
	}
	fmt.Println(result[0][1])
	fmt.Println(result[1][1])
	fmt.Println(fsIdList)
}

