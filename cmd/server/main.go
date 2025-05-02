package main

import (
    "log"
    "net/http"

    "github.com/Parab-Mishra/Matchmaking_Engine/internal/api"
    "github.com/Parab-Mishra/Matchmaking_Engine/internal/data"
)

func main() {
    data.InitStore() // initialize memory store
    router := api.SetupRouter()

    log.Println("🚀 Server starting on port 8080...")
    if err := http.ListenAndServe(":8080", router); err != nil {
        log.Fatalf("Failed to start server: %v", err)
    }
}
