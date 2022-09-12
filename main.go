package main

import (
	"fmt"
	"os"

	"github.com/antonmamonov/gounsafescanner/dummyapp"
	"github.com/antonmamonov/gounsafescanner/pointerscan"
	"github.com/antonmamonov/gounsafescanner/pointertest"
)

func help() {
	fmt.Println("-----------------------")
	fmt.Println("go run main.go dummyapp")
	fmt.Println("go run main.go pointertest")
	fmt.Println("go run main.go pointerscan")
	fmt.Println("-----------------------")
}

func main() {

	if len(os.Args) < 2 {
		panic("Please give an argument to select the submodule. If confused type 'go run main.go help'")
	}

	subModule := os.Args[1]

	if subModule == "help" {
		help()
		return
	}

	if subModule == "pointertest" {
		pointertest.TestUnsafePointers()
		return
	}

	if subModule == "dummyapp" {
		dummyapp.MainRunDummyApp()
		return
	}

	if subModule == "pointerscan" {
		pointerscan.BeginBasicPointerScannerForInt(13371337)
		return
	}

	fmt.Println("Not a valid command. If confused type 'go run main.go help'")
}
