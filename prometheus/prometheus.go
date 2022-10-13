package prometheus

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/push"
	jww "github.com/spf13/jwalterweatherman"
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

func Background(serverIP, pushgatewayUrl, pushgatewayUser, pushgatewayPassword, job, groupingName, groupingValue string) {
	jww.INFO.Print("Starting background process")

	ticker := time.NewTicker(time.Minute)
	done := make(chan bool)
	go func() {
		for {
			select {
			case <-done:
				jww.INFO.Print("Received done, stopping background progress")
				return
			case <-ticker.C:
				jww.INFO.Print("Updating value")
				UpdateValues(serverIP, pushgatewayUrl, pushgatewayUser, pushgatewayPassword, job, groupingName, groupingValue)
			}
		}
	}()

	select {}
}

func UpdateValues(serverIP, pushgatewayUrl, pushgatewayUser, pushgatewayPassword, job, groupingName, groupingValue string) {
	jww.INFO.Printf("Querying server %s", serverIP)

	resp, err := QueryYouless(serverIP)

	if err != nil {
		jww.ERROR.Println(err)
	} else {
		totalEnergyUsage.Set(float64(*resp.Net))
		totalGasUsage.Set(float64(*resp.Gas))
		currentEnergyUsage.Set(float64(*resp.Pwr))
		PushToGateway(pushgatewayUrl, pushgatewayUser, pushgatewayPassword, job, groupingName, groupingValue)
	}

}

func PushToGateway(pushgatewayUrl, pushgatewayUser, pushgatewayPassword, job, groupingName, groupingValue string) {
	jww.INFO.Printf("Pushing to gateway %s", pushgatewayUrl)
	if err := push.New(pushgatewayUrl, job).
		Collector(currentEnergyUsage).
		Collector(totalEnergyUsage).
		Collector(totalGasUsage).
		Grouping(groupingName, groupingValue).
		BasicAuth(pushgatewayUser, pushgatewayPassword).
		Push(); err != nil {
		jww.ERROR.Println("Could not push completion time to Pushgateway:", err)
	}

	jww.INFO.Printf("Push to job %s completed with grouping %s:%s", job, groupingName, groupingValue)
}
