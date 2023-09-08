package main

import (
	"sync"
	"time"

	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetLevel(logrus.InfoLevel)
	var wg sync.WaitGroup
	wg.Add(1)
	go pump()
	wg.Wait()
}

func pump() {
	ticker := time.NewTicker(time.Second * 5)
	for range ticker.C {
		logrus.Infof("current time is %s", time.Now().UTC())
	}
}
