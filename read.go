package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"sync"
	"time"
)

type Reader interface {
	read(RC chan []byte, wg *sync.WaitGroup)
}

type ReadFromFile struct {
	Path string
}

func (r *ReadFromFile) read(RC chan []byte, wg *sync.WaitGroup) {
	defer wg.Done()

	f, err := os.Open(r.Path)
	if err != nil {
		log.Fatal("open file err", err.Error())
	}
	defer f.Close()

	f.Seek(0, 2)
	rd := bufio.NewReader(f)

	for {
		line, err := rd.ReadBytes('\n')
		if err == io.EOF {
			time.Sleep(100 * time.Millisecond)
			continue
		} else if err != nil {
			log.Fatalf("ReadBytes error:%s", err.Error())
		}
		RC <- line
	}

}
