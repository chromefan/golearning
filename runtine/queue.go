package main


import "fmt"

type Element interface{}

type Queue interface {
	Push(e Element) //向队列中添加元素
	Pop() Element   //移除队列中最前面的元素
	Clear() bool     //清空队列
	Size() int       //获取队列的元素个数
	IsEmpty() bool   //判断队列是否是空
}

type sliceEntry struct {
	element []Element
}

func NewQueue() *sliceEntry {
	return &sliceEntry{}
}

//向队列中添加元素
func (entry *sliceEntry) Push(e Element) {
	entry.element = append(entry.element, e)
}

//移除队列中最前面的额元素
func (entry *sliceEntry) Pop() Element {
	if entry.IsEmpty() {
		fmt.Println("queue is empty!")
		return nil
	}

	firstElement := entry.element[0]
	entry.element = entry.element[1:]
	return firstElement
}

func (entry *sliceEntry) Clear() bool {
	if entry.IsEmpty() {
		fmt.Println("queue is empty!")
		return false
	}
	for i := 0; i < entry.Size(); i++ {
		entry.element[i] = nil
	}
	entry.element = nil
	return true
}

func (entry *sliceEntry) Size() int {
	return len(entry.element)
}

func (entry *sliceEntry) IsEmpty() bool {
	if len(entry.element) == 0 {
		return true
	}
	return false
}

func main() {
	queue := NewQueue()
	for i := 0; i < 50; i++ {
		queue.Push(i)
	}
	fmt.Println("size:", queue.Size())
	fmt.Println("移除最前面的元素：", queue.Pop())
	fmt.Println("size:", queue.Size())
	fmt.Println("清空：", queue.Clear())
	for i := 0; i < 50; i++ {
		queue.Push(i)
	}
	fmt.Println(queue.Pop())
	fmt.Println(queue.Pop())
	fmt.Println(queue.Size())
	fmt.Println(queue.element)
}
