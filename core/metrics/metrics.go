package metrics

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/rbobrovnikov/ejabberd-prometheus-metrics/core/httprequest"
)

// Define metrics
var (
	EjabberdConnectedUsersNumber = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: "ejabberd",
			Name:      "connected_users_number",
			Help:      "The number of established sessions",
		})

	EjabberdIncommingS2SNumber = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: "ejabberd",
			Name:      "incoming_s2s_number",
			Help:      "The number of incoming s2s connections on the node",
		})

	EjabberdOutgoingS2SNumber = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: "ejabberd",
			Name:      "outgoing_s2s_number",
			Help:      "The number of outgoing s2s connections on the node",
		})

	EjabberdRegisteredUsers = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: "ejabberd",
			Name:      "stats_registered_users",
			Help:      "The number of registered users",
		})

	EjabberdOnlineUsers = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: "ejabberd",
			Name:      "stats_online_users",
			Help:      "The number of online users",
		})

	// ejabberdOnlineUsers = prometheus.NewGauge(
	// 	prometheus.GaugeOpts{
	// 		Namespace: "ejabberd",
	// 		Name:      "stats_online_users",
	// 		Help:      "The number of online users",
	// 	})
)

// RecordMetrics generates metrics
func RecordMetrics(schema string, host string, port string, token string) {
	reqBodyJSONEmpty := `{}`
	scrapeInterval := (time.Duration(5) * time.Second)
	ticker := time.NewTicker(scrapeInterval)
	go func() {
		for range ticker.C {
			ejabberdMetricValue, _ := httprequest.EjabberAPICommonRequest(httprequest.HTTPBaseParams{schema, host, port, token, "connected_users_number", reqBodyJSONEmpty, "num_sessions"})
			EjabberdConnectedUsersNumber.Set(ejabberdMetricValue)
		}
	}()

	go func() {
		for range ticker.C {
			ejabberdMetricValue, _ := httprequest.EjabberAPICommonRequest(httprequest.HTTPBaseParams{schema, host, port, token, "incoming_s2s_number", reqBodyJSONEmpty, "s2s_incoming"})
			EjabberdIncommingS2SNumber.Set(ejabberdMetricValue)
		}
	}()

	go func() {
		for range ticker.C {
			ejabberdMetricValue, _ := httprequest.EjabberAPICommonRequest(httprequest.HTTPBaseParams{schema, host, port, token, "outgoing_s2s_number", reqBodyJSONEmpty, "s2s_outgoing"})
			EjabberdOutgoingS2SNumber.Set(ejabberdMetricValue)
		}
	}()

	go func() {
		reqBodyJSON := `{"name": "registeredusers"}`
		for range ticker.C {
			ejabberdMetricValue, _ := httprequest.EjabberAPICommonRequest(httprequest.HTTPBaseParams{schema, host, port, token, "stats", reqBodyJSON, "stat"})
			EjabberdRegisteredUsers.Set(ejabberdMetricValue)
		}
	}()

	go func() {
		reqBodyJSON := `{"name": "onlineusers"}`
		for range ticker.C {
			ejabberdMetricValue, _ := httprequest.EjabberAPICommonRequest(httprequest.HTTPBaseParams{schema, host, port, token, "stats", reqBodyJSON, "stat"})
			EjabberdOnlineUsers.Set(ejabberdMetricValue)
		}
	}()
}

// RegisterMetrics sets up configured metrics
func RegisterMetrics() {
	prometheus.MustRegister(EjabberdConnectedUsersNumber)
	prometheus.MustRegister(EjabberdIncommingS2SNumber)
	prometheus.MustRegister(EjabberdOutgoingS2SNumber)
	prometheus.MustRegister(EjabberdRegisteredUsers)
	prometheus.MustRegister(EjabberdOnlineUsers)
}
