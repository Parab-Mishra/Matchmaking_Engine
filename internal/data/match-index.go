package data

import (
    "sort"
)

var matchIndex = make(map[string][]MatchResult)

func StoreMatchResults(userID string, matches []MatchResult) {
    sort.Slice(matches, func(i, j int) bool {
        return matches[i].Score > matches[j].Score
    })
    if len(matches) > 5 {
        matches = matches[:5]
    }
    matchIndex[userID] = matches
}

func GetMatchResults(userID string) []MatchResult {
    return matchIndex[userID]
}
