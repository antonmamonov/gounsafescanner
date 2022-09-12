package dummyapp

import (
	"fmt"
	"log"
	"time"
)

type DummyApp struct {
	DummyInt int
}

func CreateNewDefaultDummyApp() DummyApp {
	return DummyApp{
		DummyInt: 13371337,
	}
}

func printDummyApp(dummyApp *DummyApp) error {
	log.Println("current DummyInt: ", dummyApp.DummyInt)
	log.Println("\tcurrent DummyInt Pointer Value: ", &dummyApp.DummyInt)
	fmt.Println("")
	return nil
}

func printLoop(dummyApp *DummyApp) error {

	printError := printDummyApp(dummyApp)

	if printError != nil {
		return printError
	}

	time.Sleep(1 * time.Second)

	return printLoop(dummyApp)
}

func MainRunDummyApp() error {
	dummyApp := CreateNewDefaultDummyApp()

	fmt.Println("Running Dummy App")
	fmt.Println(dummyApp)
	return printLoop(&dummyApp)
}
