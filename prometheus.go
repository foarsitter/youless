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
		Help: "Current energy usage in watts",
	})
	totalEnergyUsage = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "total_energy_usage",
		Help: "Total energy usage in watts",
	})
	totalGasUsage = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "total_gas_usage",
		Help: "Total gas usage in m^3",
	})
)

func Background() {
	fmt.Println("Doei")

	UpdateValues()

	ticker := time.NewTicker(time.Minute)
	done := make(chan bool)
	go func() {
		for {
			select {
			case <-done:
				return
			case <-ticker.C:
				UpdateValues()
			}
		}
	}()

	select {}
}

func UpdateValues() {
	fmt.Println("Hoi")
	resp, err := QueryYouless()

	if err != nil {
		return
	} else {
		totalEnergyUsage.Set(float64(*resp.Net))
		totalGasUsage.Set(float64(*resp.Gas))
		currentEnergyUsage.Set(float64(*resp.Pwr))
		PushToGateway()
	}

}

func PushToGateway() {
	if err := push.New(os.Getenv("PUSHGATEWAY_URL"), "raspcuterie").
		Collector(currentEnergyUsage).
		Collector(totalEnergyUsage).
		Collector(totalGasUsage).
		Grouping("youless", "rivierenhof").
		BasicAuth(os.Getenv("PUSHGATEWAY_USERNAME"), os.Getenv("PUSHGATEWAY_PASSWORD")).
		Push(); err != nil {
		fmt.Println("Could not push completion time to Pushgateway:", err)
	}
}
