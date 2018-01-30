package main

import "fmt"

type TestElementQ struct {
	out_ch chan<- Data
	in_ch  chan Data
}

func NewTestElementQ() *TestElementQ {
	tmp := &TestElementQ{
		in_ch: make(chan Data),
	}

	go tmp.loop()

	return tmp
}

func (s *TestElementQ) Input() chan<- Data {
	return s.in_ch
}

func (s *TestElementQ) SetOutput(sink Sink) {
	s.out_ch = sink.Input()
}

func (s *TestElementQ) loop() {
	defer fmt.Println("TestElementQ exit")

	//BEGIN OMIT
	queue := make([]Data, 0)
	var out_ch chan<- Data
	var first Data

	for {
		if len(queue) > 0 && out_ch == nil {
			first = queue[0]
			queue = queue[1:]
			out_ch = s.out_ch
		}

		select {
		case d, ok := <-s.in_ch:
			if !ok {
				close(s.out_ch)
				return
			}

			queue = append(queue, d)
		case out_ch <- first:
			out_ch = nil
		}
	}
	//END OMIT
}
