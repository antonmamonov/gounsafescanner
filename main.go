package main

import (
	"unsafe"
)

func main() {
	vals := []int{7, 12, 9, 5, 2, 14, 3}

	ptrStart := unsafe.Pointer(&vals[0])
	itemSize := unsafe.Sizeof(vals[0])

	println(ptrStart)
	println(itemSize)

	for i := 0; i < len(vals); i++ {
		println("ptrStart", ptrStart)
		println("ptrAdded", unsafe.Add(ptrStart, uintptr(i)*itemSize))
		item := *(*int)(unsafe.Add(ptrStart, uintptr(i)*itemSize))
		println(item)
	}
}
