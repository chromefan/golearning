package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"math/rand"
	"strconv"
	"time"
)

var (
	RedisClient     *redis.Pool
	RedisHost            = "127.0.0.1:6379"
)

func main() {

	var t []string
	t = append(t,"")
	fmt.Println(t)
	fmt.Println(t[0])
	fmt.Println(len(t))
	return
	NewRedisPool()
	DubbleKey()

}
func DubbleKey()  {
	key := "user_feed_hash"
	var fields []string
	str := ""
	for i:=4 ;i>=0;i--{
		str = fmt.Sprintf("test%d",i)
		fields = append(fields,str)
	}
	resHash,_:= HGetPipline(key,fields)
	fmt.Println(resHash)
	var keys []string
	for i:=0 ;i<2;i++{
		key = fmt.Sprintf("suser_feed_set_%d",i)
		keys = append(keys,key)
	}
	res,_ := ZRangeByScorePipline(keys,0)

	var vals []string
	var scores []int64
	for i:=0;i<len(res);i++  {
		for val, score := range res[i] {
			fmt.Println(val, ":", score)
			scoreNum ,_:= strconv.ParseInt(score,10,64)
			vals = append(vals,val)
			scores = append(scores,scoreNum)
		}
	}
	fmt.Println(vals,scores)
	useKey := "user_feed_set"
	err := ZAddPipline(useKey,vals,scores)
	fmt.Println(err)
}
func planeSet() string {
	rc := RedisClient.Get()
	k := "planePool"
	num,err := rc.Do("llen",k)
	if err != nil{
		fmt.Println("mylist get len err",err.Error())
	}else{
		fmt.Println("mylist's len is ",num)
	}

	defer  rc.Close()
	l := fmt.Sprintf("%s",num)
	return l
}
func planeLen() string {
	rc := RedisClient.Get()
	k := "planePool"
	num,err := rc.Do("llen",k)
	if err != nil{
		fmt.Println("mylist get len err",err.Error())
	}else{
		fmt.Println("mylist's len is ",num)
	}

	defer  rc.Close()
	l := fmt.Sprintf("%s",num)
	return l
}
func getPlane()  {
	rc := RedisClient.Get()
	k := "planePool"
	planID, err := rc.Do("lpop", k)
	if err != nil {
		fmt.Println("rpop failed", err.Error())
	}
	fmt.Printf("get plane id : %s\n", planID)
	defer  rc.Close()
}
func planePool(){
	rc := RedisClient.Get()
	k := "planePool"
	for  i:=0 ;i<100;i++{
		fmt.Println(i)
		_, err := rc.Do("lpush",k,i)
		if err != nil {
			fmt.Println("LPUSH error", err.Error())
		}
	}
}
func planePoolLoop(rc redis.Conn){

	k := "planePool"
	start := time.Now()
	l := planeLen()
	rl, _:=strconv.Atoi(l)
	num :=0
	for  {
		planID, err := rc.Do("RPOPLPUSH", k,k)
		planID = fmt.Sprintf("%s", planID)
		//_, err = rc.Do("lpush",k,planID)
		if err != nil {
			fmt.Println("LPUSH error", err.Error())
		}
		num ++
		if num >= rl{

			break
		}
		//fmt.Println("LPUSH : ", planID)
	}
	end := time.Now()
	fmt.Println("loop total time:",end.Sub(start).Seconds())
}
func planeRandom(k string)  {
	rc := RedisClient.Get()
	start := random()
	fmt.Printf("start : %d",start)
	values,err := redis.Values(rc.Do("lrange",k,start,start))
	if err != nil{
		fmt.Println("lrange err",err.Error())
	}
	fmt.Printf("mylist is:")
	for _,v := range values{
		fmt.Printf(" %s ",v.([]byte))
	}
	fmt.Println()
}
func random() int {
	rand.Seed(time.Now().UnixNano())
	num := rand.Intn(100)
	return num
}
func test(k string)  {
	rc := RedisClient.Get()
	defer  rc.Close()
	currentTimeStart := time.Now()
	for  i:=0 ;i<100000;i++{
		fmt.Println(i)
		rc.Do("INCR",k)
	}
	currentTimeEnd := time.Now()
	fmt.Println(currentTimeStart)
	fmt.Println(currentTimeEnd)
}
func TestPipline(k string)  {

	rc := RedisClient.Get()
	defer  rc.Close()
	currentTimeStart := time.Now()
	for  i:=0 ;i<100000;i++{
		fmt.Println(i)
		rc.Send("INCR",k)
	}
	rc.Flush()
	rc.Receive()
	currentTimeEnd := time.Now()
	fmt.Println(currentTimeStart)
	fmt.Println(currentTimeEnd)
}
// NewRedisPool 返回redis连接池
func NewRedisPool()  {
	RedisClient = &redis.Pool{
		// 从配置文件获取maxidle以及maxactive，取不到则用后面的默认值
		MaxIdle:     1,
		MaxActive:   10,
		IdleTimeout: 180 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", RedisHost)
			if err != nil {
				return nil, err
			}
			/*// Auth 有密码的话需要做认证哦～
			c.Do("AUTH","密码")

			// 选择db
			c.Do("SELECT", REDIS_DB)*/
			return c, nil
		},
	}
}

func set(k, v string) {
	c := RedisClient.Get()
	defer c.Close()
	_, err := c.Do("SET", k, v)
	if err != nil {
		fmt.Println("set error", err.Error())
	}
}
func setEx(k, v string,ex int) {
	c := RedisClient.Get()
	defer c.Close()
	_, err := c.Do("SET", k, v,"EX",ex)
	if err != nil {
		fmt.Println("set error", err.Error())
	}
}

func getStringValue(k string) string {
	c := RedisClient.Get()
	defer c.Close()
	username, err := redis.String(c.Do("GET", k))
	if err != nil {
		//fmt.Println("Get Error: ", err.Error())
		return ""
	}
	return username
}

func  ZRangeByScorePipline(keys []string,score int64) ([]map[string]string,error) {
	conn := RedisClient.Get()
	defer conn.Close()
	if err := conn.Err(); err != nil {
		return nil, fmt.Errorf("[get connection failed] [error:%s] [active nums:%d]",
			err.Error())
	}
	var resList  []map[string]string
	for _,key := range keys{
		fmt.Println(key)
		conn.Send("ZREVRANGEBYSCORE", key, "+inf", score,"WITHSCORES")
	}
	err := conn.Flush()
	if err !=nil {
		return nil, fmt.Errorf("[ZRANGEBYSCORE key %v failed]  [error:%s] [active nums:%d]",
			keys, err.Error())
	}
	for  i:=0 ; i<len(keys) ; i++ {
		res, _:= redis.StringMap(conn.Receive())
		resList = append(resList,res)
	}
	return resList,nil
}
func  ZAddPipline(key string,vals []string , scores []int64) error {
	conn := RedisClient.Get()
	defer conn.Close()
	if err := conn.Err(); err != nil {
		return  fmt.Errorf("[get connection failed] [error:%s] [active nums:%d]",
			err.Error())
	}
	for i,val := range vals{
		fmt.Println(val)
		conn.Send("ZADD" ,key, scores[i],val)
	}
	err := conn.Flush()
	res,err := conn.Receive()
	if err !=nil {
		return fmt.Errorf("[ZADD key %v failed]  [error:%s] [active nums:%d] [res] %v" ,
			key, err.Error(),res)
	}

	return nil
}
func  HGetPipline(key string, fields []string) ([]string, error) {
	conn := RedisClient.Get()
	defer conn.Close()
	if err := conn.Err(); err != nil {
		return nil, fmt.Errorf("[get connection failed] [error:%s] [active nums:%d]",
			err.Error())
	}
	var resList []string
	for _,field := range fields{
		fmt.Println(field)
		conn.Send("hget", key, field)
	}
	err := conn.Flush()
	if err !=nil {
		return nil, fmt.Errorf("[ZREVRANGEBYSCORE key %v failed]  [error:%s] [active nums:%d]",
			fields, err.Error())
	}
	for  i:=0 ;i<len(fields) ; i++ {
		res, _:= redis.String(conn.Receive())
		resList = append(resList,res)
	}
	return resList,nil
}