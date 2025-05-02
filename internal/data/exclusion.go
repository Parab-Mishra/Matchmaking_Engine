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

// IsEligible determines if two users are eligible to match based on various criteria
func IsEligible(user1ID, user2ID string) bool {
    // Check exclusions
    if user1ID == user2ID {
        fmt.Printf("Skipping user %s matching with self\n", user1ID)
        return false
    }

    // Get profiles
    profile1, ok1 := GetProfile(user1ID)
    profile2, ok2 := GetProfile(user2ID)
    
    if !ok1 || !ok2 {
        fmt.Printf("Profile not found for user %s or %s\n", user1ID, user2ID)
        return false
    }
    
    // Check gender preferences
    // First check if user1 is interested in user2's gender
    if !isInterestedInGender(profile1.GenderSeeking, profile2.Gender) {
        fmt.Printf("User %s not interested in %s's gender (%s)\n", user1ID, user2ID, profile2.Gender)
        return false
    }
    
    // Then check if user2 is interested in user1's gender
    if !isInterestedInGender(profile2.GenderSeeking, profile1.Gender) {
        fmt.Printf("User %s not interested in %s's gender (%s)\n", user2ID, user1ID, profile1.Gender)
        return false
    }
    
    // Add additional exclusions as necessary (blocked, disliked, etc.)
    return true
}

// isInterestedInGender checks if a user with preference 'seeking' is interested in gender 'target'
func isInterestedInGender(seeking, target string) bool {
    if seeking == "B" {
        // "B" means interested in both M and F
        return true
    }
    // Otherwise, the seeking preference must match the target gender
    return seeking == target
}