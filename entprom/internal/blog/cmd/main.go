package main

import (
	"context"
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

func createEntClientAndMigrate() *ent.Client {
	c, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	ctx := context.Background()
	// Run the auto migration tool.
	if err := c.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	return c
}

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	// Run operations.
	_, err := client.User.Create().SetName("a8m").Save(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func main() {
	// Create Ent client and migrate
	client = createEntClientAndMigrate()
	// Define const labels with our metrics
	constLabels := prometheus.Labels{"environment": "blog"}
	// Use the hook
	client.Use(entprom.Hook(constLabels))
	// Simple handler to run actions on our DB.
	http.HandleFunc("/", handler)
	// This endpoint sends metrics to the prometheus to collect
	http.Handle("/metrics", promhttp.Handler())
	log.Println("server starting on port 8080")
	// Run the server
	log.Fatal(http.ListenAndServe(":8080", nil))
}
