package main

type Data int

type Source interface {
	Start()
	Stop()
	SetOutput(Sink)
}

type Sink interface {
	Input() chan<- Data
}

type Element interface {
	Sink
	SetOutput(Sink)
}
