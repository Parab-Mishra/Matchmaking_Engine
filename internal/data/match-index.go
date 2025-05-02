package data

import (
    "sort"
)

var matchIndex = make(map[string][]MatchResult)

func StoreMatchResults(userID string, matches []MatchResult) {
    sort.Slice(matches, func(i, j int) bool {
        return matches[i].Score > matches[j].Score
    })
    // if len(matches) > 5 {
    //     matches = matches[:5]
    // }
    matchIndex[userID] = matches
}

func AppendMatchResults(userID string, matches []MatchResult) {
    // sort.Slice(matches, func(i, j int) bool {
    //     return matches[i].Score > matches[j].Score
    // })
    var finalMatches []MatchResult
    finalMatches = append(matchIndex[userID], matches...)
    sort.Slice(finalMatches, func(i, j int) bool {
        return finalMatches[i].Score > finalMatches[j].Score
    })
    // if len(finalMatches) > 5 {
    //     finalMatches = finalMatches[:5]
    // }
    matchIndex[userID] = finalMatches;
}

func GetMatchResults(userID string) []MatchResult {
    return matchIndex[userID]
}
