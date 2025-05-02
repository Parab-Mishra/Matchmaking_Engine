package core

import (
    "github.com/Parab-Mishra/Matchmaking_Engine/internal/data"
)

func GetTopMatches(userID string) []data.MatchResult {
    return data.GetMatchResults(userID)
}

// PrecomputeMatches calculates and stores top matches for a profile
func PrecomputeMatches(p data.Profile) {
    gh := GetGeohash(p)
    candidates := data.GetProfilesInGeohash(gh)

    var results []data.MatchResult
    for _, other := range candidates {
        if other.ID == p.ID || !data.IsEligible(p.ID, other.ID) {
            continue
        }
        score := ScoreProfiles(p, other)
        results = append(results, data.MatchResult{ID: other.ID, Score: score})
    }

    data.StoreMatchResults(p.ID, results)
}