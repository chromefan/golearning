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
var max []int
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
func createTree() *BNode {
	root := &BNode{}
	node1 := &BNode{}
	node2 := &BNode{}
	node3 := &BNode{}
	node4 := &BNode{}
	root.Data = 6
	node1.Data = 3
	node2.Data = 7
	node3.Data = -1
	node4.Data = 9
	root.LeftChild = node1
	root.RightChild = node2
	node1.LeftChild = node3
	node1.RightChild = node4
	return root
}
func main()  {
	/*data := []int{1,2,3,4,5,6,7,8,9,10}
	fmt.Println("数组",data)
	root := arrayToTree(data,0,len(data)-1)
	fmt.Println("转换成树的中序遍历")
	PrintTreeMidOrder(root)*/
	tree := createTree()
	fmt.Println("手动生成的树")
	PrintTreeMidOrder(tree)
	/*sum := FindMaxSubTree(tree)
	fmt.Println("所有子树和为：",sum)*/
	PreOrderToStr(tree)
	fmt.Println("子树和列表：",max)
	//fmt.Println("最大的和节点为：",maxRoot.Data)
	fmt.Println()
}


func FindMaxSubTree(root *BNode ) int  {
	if root == nil {
		return 0
	}
	//求出root所有左子树节点的和
	lmax := FindMaxSubTree(root.LeftChild)
	//求出root所有右子树节点的和
	rmax := FindMaxSubTree(root.RightChild)
	sum := lmax + rmax + root.Data.(int)
	fmt.Println("\n print ",sum,root.Data)
	max = append(max,sum)
	return sum
}
// 求最大连续子数组和
func MaxSubArray(arr []int) int {
	currSum := 0
	maxSum := arr[0]

	for _, v := range arr {
		if currSum > 0 {
			currSum += v
		} else {
			currSum = v
		}
		if maxSum < currSum {
			maxSum = currSum
		}
	}
	return maxSum
}
func PreOrderToStr(node *BNode) int {
	if node == nil {
		return 0
	}

	ret := node.Data.(int)
	fmt.Println(ret)
	ret += PreOrderToStr(node.LeftChild)
	ret += PreOrderToStr(node.RightChild)
	max = append(max,ret)
	return ret
}