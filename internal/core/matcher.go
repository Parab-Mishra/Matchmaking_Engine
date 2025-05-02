package core

import (
	"fmt"

    "github.com/Parab-Mishra/Matchmaking_Engine/internal/data"
)

func GetAllMatches(userID string) []data.MatchResult {
    return data.GetMatchResults(userID)
}

// PrecomputeMatches calculates and stores top matches for a profile
func PrecomputeMatches(p data.Profile) {
    gh := GetGeohash(p)
    candidates := data.GetProfilesInGeohash(gh)

    var results []data.MatchResult
    
    fmt.Printf("User %s → %d candidates in geohash %s\n", p.ID, len(candidates), gh)  // Debugging log
    for _, other := range candidates {
        if other.ID == p.ID || !data.IsEligible(p.ID, other.ID) {
            continue
        }
        score := ScoreProfiles(p, other)
        fmt.Printf("Scoring %s and %s: %.2f\n", p.ID, other.ID, score)  // Debugging log
        results = append(results, data.MatchResult{ID: other.ID, Score: score})
        var otherMatchResults []data.MatchResult
        otherMatchResults = append(otherMatchResults, data.MatchResult{ID: p.ID, Score: score})
        data.AppendMatchResults(other.ID, otherMatchResults)
    }


    if len(results) > 0 {
        data.StoreMatchResults(p.ID, results)
    }
}
