package main

import (
	"context"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"log"
	"strconv"
	"sync"
	"time"
)

type Writer interface {
	write(WC chan *GinLog, wg *sync.WaitGroup)
}

type WriteToInfluxDb struct {
	InfluxDBDsn string
}

func (w *WriteToInfluxDb) write(WC chan *GinLog, wg *sync.WaitGroup) {
	defer wg.Done()

	// Create a new client using an InfluxDB server base URL and an authentication token
	client := influxdb2.NewClient("http://localhost:8086", "admin_token")
	// Use blocking write client for writes to desired bucket
	writeAPI := client.WriteAPIBlocking("primary", "primary")
	defer client.Close()

	for v := range WC {

		parse, err := time.Parse("2006-01-02T15:04:05.000Z0700", v.Timestamp)

		if err != nil {
			log.Println("InfluxDB parsing time error:", err.Error())
		}

		p := influxdb2.NewPointWithMeasurement("GinLog").
			AddTag("Path", v.Path).
			AddTag("Method", v.Method).
			AddTag("Status", strconv.Itoa(int(v.Status))).
			AddField("Duration", v.Duration).
			AddField("ByteSent", 1/v.Duration).
			SetTime(parse)

		err = writeAPI.WritePoint(context.Background(), p)
		if err != nil {
			log.Println("InfluxDB write error", err.Error())
		}

		log.Println("InfluxDB write success")

	}
}
