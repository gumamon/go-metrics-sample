package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/prometheus/client_golang/prometheus"
	process "github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	// Prometheus用のレジストリを作成
	reg := prometheus.NewRegistry()

	// プロセスのメトリクスを追加
	reg.MustRegister(
		process.NewProcessCollector(process.ProcessCollectorOpts{}),
	)

	// /ping handler（ポート8080）
	go func() {
		http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintln(w, "pong")
		})
		fmt.Println("Starting ping server on :8080")
		if err := http.ListenAndServe(":8080", nil); err != nil {
			fmt.Println("ping server error:", err)
			os.Exit(1)
		}
	}()

	// /metrics handler（ポート8000）
	http.Handle("/metrics", promhttp.HandlerFor(reg, promhttp.HandlerOpts{}))
	fmt.Println("Starting metrics server on :8000")
	if err := http.ListenAndServe(":8000", nil); err != nil {
		fmt.Println("metrics server error:", err)
		os.Exit(1)
	}
}
