package main

import (
	"data_structures/HashMap"
	"fmt"
)

func main() {
	h := HashMap.HashMap{}
	h.Init()
	h.Add("Sam")
	h.Add("Nico")
	h.Add("Jo")
	h.Add("Meredith")

	h.Print()

	fmt.Println(h.Find("Sam"))
}
