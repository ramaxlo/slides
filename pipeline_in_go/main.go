package main

import (
	"fmt"
	"time"
)

func main() {
	//BEGIN OMIT
	source := NewTestSource()
	element := NewTestElement()
	sink := NewTestSink()

	source.SetOutput(element)
	element.SetOutput(sink)

	fmt.Println("Start pipeline")
	source.Start()
	//END OMIT

	time.Sleep(5 * time.Second)
	source.Stop()
	time.Sleep(1 * time.Second)
}
