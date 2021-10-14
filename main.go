package main

import (
	"sync"
)

type LogProcess struct {
	RC    chan []byte
	WC    chan *GinLog
	Read  Reader
	Write Writer
}

func main() {

	read := &ReadFromFile{
		Path: "C:\\Users\\zaqxs\\GolandProjects\\redis_prac\\info.log",
		//Path: "./access.log",
	}

	write := &WriteToInfluxDb{
		InfluxDBDsn: "localhost:8086@admin@admin@db0",
	}

	lp := &LogProcess{
		RC:    make(chan []byte),
		WC:    make(chan *GinLog),
		Read:  read,
		Write: write,
	}

	wg := sync.WaitGroup{}
	wg.Add(3)
	go lp.Read.read(lp.RC, &wg)
	go lp.process(&wg)
	go lp.Write.write(lp.WC, &wg)
	wg.Wait()
}
