package trade

import (
	"fmt"
	"math/rand"
)

const maxLevel = 32
const prob = 0.35

type SkipValue struct {
	Score int64
	Value interface{}
}

func (v *SkipValue) Compare(value *SkipValue) int {
	f1 := v.Score
	f2 := value.Score
	if f1 > f2 {
		return 1
	} else if f1 == f2 {
		return 0
	}
	return -1
}

type skipListNode struct {
	next  []*skipListNode
	prev  *skipListNode
	Value *SkipValue
}

type SkipList struct {
	header, tail *skipListNode
	findcount    int
	count        int
	level        int
}

type SkipListIterator struct {
	list *SkipList
	node *skipListNode
}

func (sli *SkipListIterator) First() *SkipValue {
	if sli.list.header.next[0] == nil {
		return nil
	}
	sli.node = sli.list.header.next[0]
	return sli.node.Value
}

func (sli *SkipListIterator) Last() *SkipValue {
	if sli.list.tail == nil {
		return nil
	}
	sli.node = sli.list.tail
	return sli.node.Value
}

func (sli *SkipListIterator) Current() *SkipValue {
	if sli.node == nil {
		return nil
	}
	return sli.node.Value
}

func (node *skipListNode) Prev() *skipListNode {
	if node == nil || node.prev == nil {
		return nil
	}
	return node.prev
}

func (node *skipListNode) Next() *skipListNode {
	if node == nil || node.next[0] == nil {
		return nil
	}
	return node.next[0]
}

func (sli *SkipListIterator) Seek(value *SkipValue) *SkipValue {
	x := sli.list.find(value)
	if x.next[0] == nil {
		return nil
	}
	sli.node = x.next[0]
	return sli.node.Value
}

func newskipListNode(level int, value *SkipValue) *skipListNode {
	node := &skipListNode{}
	node.next = make([]*skipListNode, level)
	node.Value = value
	return node
}

//构建一个value的最小值
func NewSkipList(min *SkipValue) *SkipList {
	sl := &SkipList{}
	sl.level = 1
	sl.header = newskipListNode(maxLevel, min)
	return sl
}

func randomLevel() int {
	level := 1
	t := prob * 0xFFFF
	for rand.Int()&0xFFFF < int(t) {
		level += 1
		if level == maxLevel {
			break
		}
	}
	return level
}

func (sl *SkipList) GetIterator() *SkipListIterator {
	it := &SkipListIterator{}
	it.list = sl
	it.First()
	return it
}

func (sl *SkipList) Len() int {
	return sl.count
}

func (sl *SkipList) Level() int {
	return sl.level
}

func (sl *SkipList) find(value *SkipValue) *skipListNode {
	x := sl.header
	for i := sl.level - 1; i >= 0; i-- {
		for x.next[i] != nil && x.next[i].Value.Compare(value) < 0 {
			sl.findcount++
			x = x.next[i]
		}
	}
	return x
}

func (sl *SkipList) FindCount() int {
	return sl.findcount
}

func (sl *SkipList) Find(value *SkipValue) *SkipValue {
	x := sl.find(value)
	if x.next[0] != nil && x.next[0].Value.Compare(value) == 0 {
		return x.next[0].Value
	}
	return nil
}

func (sl *SkipList) FindGreaterOrEqual(value *SkipValue) *SkipValue {
	x := sl.find(value)
	if x.next[0] != nil {
		return x.next[0].Value
	}
	return nil
}

func (sl *SkipList) Insert(value *SkipValue) int {
	var update [maxLevel]*skipListNode
	x := sl.header
	for i := sl.level - 1; i >= 0; i-- {
		for x.next[i] != nil && x.next[i].Value.Compare(value) <= 0 {
			x = x.next[i]
		}
		update[i] = x
	}
	//if x.next[0] != nil && x.next[0].Value.Compare(value) == 0 { //update
	//	x.next[0].Value = value
	//	return 0
	//}
	level := randomLevel()
	if level > sl.level {
		for i := sl.level; i < level; i++ {
			update[i] = sl.header
		}
		sl.level = level
	}
	x = newskipListNode(level, value)
	for i := 0; i < level; i++ {
		x.next[i] = update[i].next[i]
		update[i].next[i] = x
	}
	//形成一个双向链表
	if update[0] != sl.header {
		x.prev = update[0]
	}
	if x.next[0] != nil {
		x.next[0].prev = x
	} else {
		sl.tail = x
	}
	sl.count++
	return 1
}

func (sl *SkipList) Delete(value *SkipValue) int {
	var update [maxLevel]*skipListNode
	x := sl.header
	for i := sl.level - 1; i >= 0; i-- {
		for x.next[i] != nil && x.next[i].Value.Compare(value) < 0 {
			x = x.next[i]
		}
		update[i] = x
	}
	if x.next[0] == nil || x.next[0].Value.Compare(value) != 0 { //not find
		return 0
	}
	x = x.next[0]
	for i := 0; i < sl.level; i++ {
		if update[i].next[i] == x {
			update[i].next[i] = x.next[i]
		}
	}
	if x.next[0] != nil {
		x.next[0].prev = x.prev
	} else {
		sl.tail = x.prev
	}
	for sl.level > 1 && sl.header.next[sl.level-1] == nil {
		sl.level--
	}
	sl.count--
	return 1
}

//测试用的输出函数
func (l *SkipList) Print() {
	if l.count > 0 {
		r := l.header
		for i := l.level - 1; i >= 0; i-- {
			e := r.next[i]
			//fmt.Print(i)
			for e != nil {
				fmt.Print(e.Value)
				fmt.Print("    ")
				e = e.next[i]
			}
			fmt.Println()
		}
	} else {
		fmt.Println("空")
	}
}

//Walk 遍历整个结构，如果cb 返回false 那么停止遍历
func (lm *SkipList) Walk(cb func(value interface{}) bool) {
	for e := lm.header.Next(); e != nil; e = e.Next() {
		if cb == nil {
			return
		}
		if !cb(e.Value.Value) {
			return
		}
	}
}
