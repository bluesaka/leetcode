package main

import (
	"container/list"
	"fmt"
)

type CQueue struct {
	// 使用双向链表实现栈和队列，或者使用slice实现
	// stackIn做插入栈
	// stackOut做删除辅助栈
	// container/list库的双向链表
	stackIn, stackOut *list.List
}

func Constructor() CQueue {
	return CQueue{
		stackIn:  list.New(),
		stackOut: list.New(),
	}
}

func (this *CQueue) AppendTail(value int) {
	this.stackIn.PushBack(value)
}

func (this *CQueue) DeleteHead() int {
	// 如果删除栈为空，则从插入栈中弹出并插入到删除栈
	if this.stackOut.Len() == 0 {
		for this.stackIn.Len() > 0 {
			this.stackOut.PushBack(this.stackIn.Remove(this.stackIn.Back()))
		}
	}

	// 弹出删除栈的最后一个元素
	if this.stackOut.Len() != 0 {
		e := this.stackOut.Back()
		this.stackOut.Remove(e)
		return e.Value.(int)
	}

	// 如果获取后的删除栈还未空，则说明没有数据，返回-1
	return -1
}

/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */
func main() {
	obj := Constructor()
	obj.AppendTail(1)
	obj.AppendTail(5)
	obj.AppendTail(3)
	obj.AppendTail(2)
	fmt.Println(obj.DeleteHead())
	fmt.Println(obj.DeleteHead())
	fmt.Println(obj.DeleteHead())
	fmt.Println(obj.DeleteHead())
	fmt.Println(obj.DeleteHead())

}
