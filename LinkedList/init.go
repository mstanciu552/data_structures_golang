package LinkedList

import "fmt"

type LinkedList struct {
	data interface{}
	next *LinkedList
}

func Init_ll(data interface{}) *LinkedList {
	return &LinkedList{data: data}
}

func (ll *LinkedList) Get_Data() interface{} {
	return ll.data
}

func (ll *LinkedList) Set_data(data interface{}) {
	ll.data = data
}

func (ll *LinkedList) Set_next(next *LinkedList) {
	ll.next = next
}

func Convert(data []interface{}) LinkedList {
	ll := LinkedList{}
	head := &ll

	for _, dt := range data {
		head.next = &LinkedList{data: dt}
		head = head.next
	}

	return ll
}

func (ll *LinkedList) Printll() {
	for ll != nil {
		fmt.Println(ll)
		ll = ll.next
	}
}

func (ll *LinkedList) Find(data interface{}) (bool, *LinkedList) {
	for ll != nil {
		if ll.Get_Data() == data {
			return true, ll
		}
		ll = ll.next
	}
	return false, nil
}
