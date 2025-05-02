package api

import (
    "encoding/json"
    "net/http"
    "fmt"

    "github.com/gorilla/mux"
    "github.com/Parab-Mishra/Matchmaking_Engine/internal/core"
    "github.com/Parab-Mishra/Matchmaking_Engine/internal/data"
)

func CreateProfileHandler(w http.ResponseWriter, r *http.Request) {
    var profile data.Profile
    if err := json.NewDecoder(r.Body).Decode(&profile); err != nil {
        http.Error(w, "Invalid input", http.StatusBadRequest)
        return
    }

    geohash := core.GetGeohash(profile)
    data.StoreProfile(profile, geohash)
    core.PrecomputeMatches(profile)

    w.WriteHeader(http.StatusCreated)
}

func MatchHandler(w http.ResponseWriter, r *http.Request) {
    
    id := mux.Vars(r)["id"]
    fmt.Printf("Matching for user: %s\n", id)
    matches := core.GetTopMatches(id)


    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(matches)
}
