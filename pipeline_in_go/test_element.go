package main

import "fmt"

type TestElement struct {
	out_ch chan<- Data
	in_ch  chan Data
}

func NewTestElement() *TestElement {
	tmp := &TestElement{
		in_ch: make(chan Data),
	}

	go tmp.loop()

	return tmp
}

func (s *TestElement) Input() chan<- Data {
	return s.in_ch
}

func (s *TestElement) SetOutput(sink Sink) {
	s.out_ch = sink.Input()
}

func (s *TestElement) loop() {
	defer fmt.Println("TestElement exit")

	for {
		select {
		//BEGIN OMIT
		case d, ok := <-s.in_ch: // HL1
			if !ok {
				close(s.out_ch)
				return
			}
			//END OMIT

			s.out_ch <- d // HL1
		}
	}
}
