package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"icode.baidu.com/baidu/netdisk/clouddisk-golib/utils"
	"net"
	"net/url"
	"strings"
	"time"
)

func main() {
	t := time.Now().UnixNano() / 1e6
	fmt.Println(t)
	params := url.Values{}
	params.Add("user","1")
	params.Add("pass",fmt.Sprintf("%d",1))
	paramsStr := params.Encode()
	sign := utils.Md5(paramsStr)

	 var postParams = make(map[string]string)
	postParams["user"] = params.Get("user")
	fmt.Println(sign)
	v := 31004
	var couponTypes  []int
	couponTypes = append(couponTypes, v)
	couponTypesStr , _ := json.Marshal(couponTypes)
	str := url.QueryEscape(string(couponTypesStr))
	fmt.Println(str)

	/*var ips []string
	local := "127.0.0.1:8080"
	//local := "fe80:9528:4249:57b9:c2eb"
	ips =  append(ips,local)
	ips =  append(ips,local)
	addr := ClientIP(ips)
	fmt.Println(addr)
	ipv4 := "127.0.0"
	// ParseIP 这个方法 可以用来检查 ip 地址是否正确，如果不正确，该方法返回 nil
	address := CheckIp(ipv4)
	if address == false {
		fmt.Println("ip地址格式不正确")
	}else {
		fmt.Println("正确的ip地址", address)
	}
	return
	from := "search_box12"
	if  from == "search_box"{
		fmt.Println(from)
	}
	token := "509b1532-e333-20f9-a768-7429e9db0f5b"
	traceId := "dd"
	antiInfo := `{"c":"14DBF19E6ED5748286CCA5CA36013466|0","z":"","ip":"61.135.169.84","app":"android","ts":"1577268742275","ver":"10.0.140.0","os_version":"","model":"","brand":"","ua":"x+;android-android;7.1.1;","product_id":15}`
	param := url.Values{}
	param.Add("antiInfo", antiInfo)
	param.Add("reqId" , traceId)
	param.Add("token", token)
	paramsStr := param.Encode()
	fmt.Println(paramsStr)

	preSignStr := fmt.Sprintf("reqId=%s&token=%s",)
	fmt.Println(preSignStr)

	sign :=HmacSha256(preSignStr,token)
	fmt.Println(sign)

	fmt.Println(paramsStr)*/
}

func HmacSha256(str string, secret string) string {
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}
func CheckIp(ip string) bool {
	address := net.ParseIP(ip)
	if address == nil {
		return false
	}else{
		return true
	}
}
func  ClientIP( ips []string) string {
	len_ip := len(ips)
	if len(ips) > 0 && ips[len_ip-1] != "" {
		if CheckIp(ips[len_ip-1]) {
		}
		pos :=  utils.Strripos(ips[len_ip-1],":",-1)
		rip := Substr(ips[len_ip-1],0,pos)
		return rip
	}
	ip := strings.Split("", ",")
	if len(ip) > 0 {
		if ip[0] != "[" {
			return ip[0]
		}
	}
	return "127.0.0.1"
}
func Substr(str string, start int, length int) string {
	if start < 0 || length < -1 {
		return str
	}
	switch {
	case length == -1:
		return str[start:]
	case length == 0:
		return ""
	}
	end := int(start) + length
	if end > len(str) {
		end = len(str)
	}
	return str[start:end]
}
func GenOrderId(uid uint64) uint64 {
	date := time.Now().Format("20060102")
	r := GetRand(999)
	u := (uid + uint64(time.Now().Unix())) % 10000000

	id := fmt.Sprintf("%s%07d%03d", date, u, r)
	oid, _ := strconv.ParseUint(id, 10, 64)

	return oid
}