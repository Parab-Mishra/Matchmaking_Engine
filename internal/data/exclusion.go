package data

import (
	"fmt"
)

var (
    blocked   = make(map[string]map[string]struct{})
    disliked  = make(map[string]map[string]struct{})
    matched   = make(map[string]map[string]struct{})
)

func isExcluded(me, other string) bool {
    if blocked[me][other] != struct{}{} && disliked[me][other] != struct{}{} && matched[me][other] != struct{}{} {
        return false
    }
    return true
}

func addExclusion(from, to string, kind string) {
    var list map[string]map[string]struct{}
    switch kind {
    case "block":
        list = blocked
    case "dislike":
        list = disliked
    case "match":
        list = matched
    }
    if list[from] == nil {
        list[from] = make(map[string]struct{})
    }
    list[from][to] = struct{}{}
}

func IsEligible(user1ID, user2ID string) bool {
    // Check exclusions
    if user1ID == user2ID {
        fmt.Printf("Skipping user %s matching with self\n", user1ID)
        return false
    }
    // Add additional exclusions as necessary (blocked, disliked, etc.)
    return true
}
