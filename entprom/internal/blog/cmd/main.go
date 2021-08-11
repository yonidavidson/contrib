package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"entgo.io/contrib/entprom/internal/blog/entprom"
	"entgo.io/contrib/entprom/internal/ent"
	_ "entgo.io/contrib/entprom/internal/ent/runtime"
	_ "github.com/mattn/go-sqlite3"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var client *ent.Client

func initClient() {
	client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	ctx := context.Background()
	// Run the auto migration tool.
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	constLabels := prometheus.Labels{"environment": "blog"}
	client.Use(entprom.Hook(constLabels))
}

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	// Run operations.
	a8m, err := client.User.Create().SetName("a8m").Save(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	root, err := client.File.Create().SetName("/").SetOwner(a8m).Save(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	_, err = client.File.Create().SetName("dev").SetParent(root).Save(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "check the metrics")
}

func main() {
	initClient()
	http.HandleFunc("/", handler)
	http.Handle("/metrics", promhttp.Handler())
	log.Println("server starting on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
