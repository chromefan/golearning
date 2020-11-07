package main

import (
	"fmt"
	. "github.com/isdamir/gotype"
)

type bigTree struct {
	data       interface{}
	leftChild  *bigTree
	rightChild *bigTree
}

func arrayToTree(arr []int, start int, end int) *BNode {
	var root *BNode
	if end >= start {
		root = NewBNode()
		mid := (start + end + 1) / 2
		root.Data = arr[mid]
		root.LeftChild = arrayToTree(arr, start, mid-1)
		root.RightChild = arrayToTree(arr, mid+1, end)
	}
	return root
}
func main()  {
	data := []int{1,2,3,4,5,6,7,8,9,10}
	fmt.Println("数组",data)
	root := arrayToTree(data,0,len(data)-1)
	fmt.Println("转换成树的中序遍历")
	PrintTreeMidOrder(root)
}

var max int
func FindMaxSubTree(root *BNode,maxRoot *BNode) int {
	if root == nil {
		return 0
	}
	//求出root所有左子树节点的和
	lmax := FindMaxSubTree(root.LeftChild,maxRoot)
	//求出root所有右子树节点的和
	rmax := FindMaxSubTree(root.RightChild,maxRoot)
	sum := lmax + rmax + root.Data.(int)
	if sum > max {
		max = sum
		maxRoot.Data = root.Data
	}
	return sum
}
