package core

import (
    "math"
    "github.com/Parab-Mishra/Matchmaking_Engine/internal/data"
)

func ScoreProfiles(p1, p2 data.Profile) float64 {
    ageScore := 100 - math.Abs(float64(p1.Age-p2.Age))
    interestScore := float64(len(data.CommonInterests(p1.Interests, p2.Interests))) * 10
    return ageScore + interestScore
}
