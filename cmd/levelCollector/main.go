package main

import (
	"fmt"
	"github.com/kislikjeka/sdp-collector/collector"
	"log"
	"os"
	"time"
)

func main() {
	f, err := os.OpenFile("level_collector", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()

	logger := log.Logger{}
	logger.SetOutput(f)

	ctr := collector.Collector{
		Name:            "level_collector",
		MaxTimeDuration: 86400 * time.Second,
	}

	lastTime, nextTime, err := ctr.GetTimeProceed()
	if err != nil {
		logger.Println(fmt.Sprintf("Error while taking last time %s", err))
		return
	}

	ctr.WriteError("Start Collector")

	logger.Println(fmt.Sprintf("LastTime: %s, NextTime: %s", lastTime.String(), nextTime.String()))
}
