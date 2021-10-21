package HashMap

import (
	"data_structures/LinkedList"
	"fmt"
)

type HashMap struct {
	array [10]*LinkedList.LinkedList
}

func Hash(data string, length int) int {
	var sum int = 0
	for _, val := range data {
		sum += int(val)
	}
	return sum % length
}

func (h *HashMap) Init() {
	for i := range h.array {
		h.array[i] = LinkedList.Init_ll(-1)
	}
}

func (h *HashMap) Print() {
	for i, ll := range h.array {
		fmt.Printf("%d -> ", i)
		ll.Printll()
	}
}

func (h *HashMap) Add(data string) {
	if h.array[Hash(data, 10)].Get_Data() == -1 {
		h.array[Hash(data, 10)].Set_data(data)
	} else {
		h.array[Hash(data, 10)].Set_next(LinkedList.Init_ll(data))
	}
}

func (h *HashMap) Find(data string) (bool, *LinkedList.LinkedList) {
	index := Hash(data, 10)

	return h.array[index].Find(data)
}
