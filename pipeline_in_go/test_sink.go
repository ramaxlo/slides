package main

import "fmt"

type TestSink struct {
	in_ch chan Data
}

func NewTestSink() *TestSink {
	tmp := &TestSink{
		in_ch: make(chan Data),
	}

	go tmp.loop()

	return tmp
}

//IMP BEGIN OMIT
func (s *TestSink) Input() chan<- Data {
	return s.in_ch
}

func (s *TestSink) loop() {
	defer fmt.Println("TestSink exit")

	for {
		select {
		//BEGIN2 OMIT
		case d, ok := <-s.in_ch: // HL1
			if !ok {
				return
			}

			//END2 OMIT
			fmt.Printf("Got data: %v\n", d) // HL1
		}
	}
}

//IMP END OMIT
