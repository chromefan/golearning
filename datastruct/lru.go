package main

import (
	"fmt"
	"github.com/isdamir/gotype"
)

type LRU struct {
	cacheSize int
	queue *gotype.SliceQueue
	hashSet *gotype.Set
}

func (l *LRU) IsQueueFull()  bool {
	return l.queue.Size() == l.cacheSize
}
func (l *LRU) EnQueue(pageNum int)  {
	//如果队列已满则移除队尾元素
	if l.IsQueueFull() {
		l.hashSet.Remove(l.queue.PopBack())
	}

	l.queue.EnQueueFirst(pageNum)
	l.hashSet.Add(pageNum)
}
func (l *LRU) AccessPage(pageNum int)  {
	if !l.hashSet.Contains(pageNum) {
		l.EnQueue(pageNum)
	}else if pageNum != l.queue.GetFront(){
		//如果元素存在并且不在队头则移除已有元素，并重新加入到队头
		l.queue.Remove(pageNum)
		l.queue.EnQueueFirst(pageNum)
	}
}

func (l *LRU) PrintQueue()  {
	for !l.queue.IsEmpty() {
		fmt.Println(l.queue.DeQueue())
	}
}
func main()  {
	lru := &LRU{
		cacheSize: 10,
		queue: gotype.NewSliceQueue(),
		hashSet: gotype.NewSet(),
	}
	lru.AccessPage(1)
	lru.AccessPage(3)
	lru.AccessPage(5)
	lru.AccessPage(7)
	lru.AccessPage(1)
	lru.AccessPage(1)
	lru.AccessPage(6)
	lru.AccessPage(3)
	lru.AccessPage(10)
	lru.AccessPage(8)
	fmt.Println(lru.hashSet)
	lru.PrintQueue()
}
