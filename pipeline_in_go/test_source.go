package main

import (
	"fmt"
	"time"
)

type TestSource struct {
	out_ch  chan<- Data
	stop_ch chan bool
}

func NewTestSource() *TestSource {
	tmp := &TestSource{
		stop_ch: make(chan bool),
	}

	return tmp
}

//IMP BEGIN OMIT
func (s *TestSource) Start() {
	go s.loop()
}

func (s *TestSource) Stop() {
	close(s.stop_ch)
}

func (s *TestSource) SetOutput(sink Sink) {
	s.out_ch = sink.Input()
}

//IMP END OMIT

func (s *TestSource) readdata(data_ch chan<- Data) {
	timeout := time.After(1 * time.Second)
	count := 0

	for {
		select {
		case <-timeout:
			d := Data(count)
			data_ch <- d

			timeout = time.After(1 * time.Second)

		case <-s.stop_ch:
			return
		}

		count++
	}
}

func (s *TestSource) loop() {
	defer fmt.Println("TestSource exit")

	data_ch := make(chan Data)
	go s.readdata(data_ch)

	for {
		select {
		case d := <-data_ch:
			s.out_ch <- d // HL1

		//BEGIN2 OMIT
		case <-s.stop_ch:
			close(s.out_ch)
			return
		}
		//END2 OMIT

		count++
	}
}
