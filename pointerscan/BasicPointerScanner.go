package pointerscan

import (
	"fmt"
	"unsafe"
)

type BasicPointerScanner struct {
}

func BeginBasicPointerScannerForInt(intValueToScanFor int) error {
	fmt.Println("BEGIN scanning for", intValueToScanFor)

	testInt := 0
	ptrStart := unsafe.Pointer(&testInt)

	fmt.Println("ptrStart", ptrStart)

	// fmt.Println(unsafe.Pointer(0))

	maxCount := 9999

	hitPointers := []unsafe.Pointer{}

	for i := 0; i < maxCount; i++ {

		// check both pointers under and over us

		checkPointerNegative := unsafe.Add(ptrStart, -i)
		checkPointerItemNegative := *(*int)(checkPointerNegative)

		checkPointerPositive := unsafe.Add(ptrStart, i)
		checkPointerItemPositive := *(*int)(checkPointerPositive)

		if checkPointerItemNegative != 0 || checkPointerItemPositive != 0 {

			if checkPointerItemNegative == intValueToScanFor {
				println("")
				println("HIT VALUE")
				println("checkPointerHIT", checkPointerNegative)
				println("checkPointerItemHIT", checkPointerItemNegative)
				println("")
			}

			if checkPointerItemPositive == intValueToScanFor {
				println("")
				println("HIT VALUE")
				println("checkPointerHIT", checkPointerPositive)
				println("checkPointerItemHIT", checkPointerItemPositive)
				println("")

				hitPointers = append(hitPointers, checkPointerPositive)
			}

			fmt.Println(checkPointerNegative, checkPointerItemNegative)
			fmt.Println(checkPointerPositive, checkPointerItemPositive)
		}
	}

	if len(hitPointers) == 0 {
		println("No pointers found!")
		return nil
	}

	// fmt.Println(hitPointers)

	// for hitPointer := range hitPointers {
	// 	// hitPointerValue := *(*int)(hitPointer)
	// 	println(hitPointer)
	// 	// println(hitPointerValue)
	// }

	return nil
}
