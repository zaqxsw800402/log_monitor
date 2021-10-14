package main

import (
	"encoding/json"
	"sync"
)

type GinLog struct {
	Level     string  `json:"level"`
	Timestamp string  `json:"timestamp"`
	Caller    string  `json:"caller"`
	Message   string  `json:"msg"`
	Method    string  `json:"method"`
	Status    int32   `json:"status"`
	Path      string  `json:"path"`
	Query     string  `json:"query"`
	Ip        string  `json:"ip"`
	UserAgent string  `json:"user-agent"`
	Error     string  `json:"errors"`
	Duration  float64 `json:"cost"`
}

func (l *LogProcess) process(wg *sync.WaitGroup) {
	defer wg.Done()
	ginlog := &GinLog{}

	for data := range l.RC {

		_ = json.Unmarshal(data, ginlog)
		l.WC <- ginlog
		//l.WC <- strings.ToUpper(string(data))
	}
}
