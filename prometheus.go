package main

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/push"
	"os"
	"time"
)

var (
	currentEnergyUsage = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "current_energy_usage",
		Help: "Returns the measured humidity in percent.",
	})
)

func Background() {
	fmt.Println("Doei")
	ticker := time.NewTicker(time.Minute)
	done := make(chan bool)
	go func() {
		for {
			select {
			case <-done:
				return
			case <-ticker.C:
				fmt.Println("Hoi")
				resp := QueryYouless()

				currentEnergyUsage.Set(float64(*resp.Pwr))
				PushToGateway()
			}
		}
	}()

	select {}
}

func PushToGateway() {
	if err := push.New(os.Getenv("PUSHGATEWAY_URL"), "raspcuterie").
		Collector(currentEnergyUsage).
		Grouping("youless", "rivierenhof").
		BasicAuth(os.Getenv("PUSHGATEWAY_USERNAME"), os.Getenv("PUSHGATEWAY_PASSWORD")).
		Push(); err != nil {
		fmt.Println("Could not push completion time to Pushgateway:", err)
	}
}
