package main

import (
	"github.com/codestoke/oracle_exporter/collector"
	"github.com/prometheus/client_golang/prometheus"
	"log"
	"os"

	"database/sql"
	_ "github.com/mattn/go-oci8"
)

const (
	namespace = "oracle"
)

var (
	oracleSessionCount = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "session", "count"),
		"the number of sessions in the instance",
		nil,
		nil,
	)
	oracleSessionCountMax = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "session", "count_max"),
		"the maximum number of sessions configured for this instance",
		nil,
		nil,
	)
)

type Exporter struct {
	sessionClient *collector.OracleSessionClient
}

func Describe(ch chan<- *prometheus.Desc) {
	ch <- oracleSessionCount
	ch <- oracleSessionCountMax
}

func main() {
	db, err := sql.Open("oci8", "user:")
	if err != nil {
		log.Fatal("could not open db")
		os.Exit(1)
	}
	defer db.Close()

	db.Ping()
}
