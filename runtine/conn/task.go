package main

import (
	"fmt"
	"net/http"
	"strconv"
	"sync"
	"time"
)

type Task struct {
	id int
}

var	mu sync.Mutex


func  RunTask(syncFlag chan bool) {
	mu.Lock()
	fmt.Println("locked")
	flag := <-syncFlag
	for _, task := range taskList {
		fmt.Println("run task ", task.id, flag)
		time.Sleep(1 * time.Second)
	}
	mu.Unlock()
	fmt.Println("unlocked")
}

var taskList []*Task
func addTask(id int) {
	syncFlag := make(chan bool)

	go func() {
		for i := 0; i < 2; i++ {
			task := &Task{}
			task.id = id
			taskList = append(taskList, task)
		}
		syncFlag <- true
	}()
	var 	mu sync.Mutex
	go runTask(syncFlag,mu)
}

func runTask(syncFlag chan bool,mu sync.Mutex) {
	mu.Lock()
	fmt.Println("locked")
	flag := <-syncFlag
	for _, task := range taskList {
		fmt.Println("run task ", task.id, flag)
		time.Sleep(1 * time.Second)

	}
	mu.Unlock()
	fmt.Println("unlocked")
}
func hiHandler(w http.ResponseWriter, r *http.Request) {
	//time.Sleep(200 * time.Millisecond)、
	query := r.URL.Query()
	fmt.Println(time.Now(),query)
	str := query.Get("id")
	id, _ := strconv.Atoi(str)

	go addTask(id)
	fmt.Println("w :",id)
	w.Write([]byte("hi"))
}
func main() {

	for i:=0 ; i<=10; i++ {
		for j := 0; j< 10; j++ {
			fmt.Println("j",j)
			break
		}
		fmt.Println("i",i)
		//go addTask(i)
	}
	time.Sleep(100*time.Second)
	/*http.HandleFunc("/", hiHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}*/
}
