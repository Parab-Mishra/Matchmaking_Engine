package core_test

import (
	"testing"

	"github.com/Parab-Mishra/Matchmaking_Engine/internal/core"
	"github.com/Parab-Mishra/Matchmaking_Engine/internal/data"
)

// TestScoreProfiles_ZeroAge tests scoring when age is zero
func TestScoreProfiles_ZeroAge(t *testing.T) {
	p1 := data.Profile{
		ID:        "user1",
		Age:       0,
		Interests: []string{"music", "art"},
	}
	
	p2 := data.Profile{
		ID:        "user2",
		Age:       25,
		Interests: []string{"music", "travel"},
	}
	
	// The expected score would be:
	// - Age score: 100 - 25 = 75 (since |0-25| = 25)
	// - Interest score: 1 * 10 = 10 (for "music")
	// Total: 75 + 10 = 85
	expected := 85.0
	
	score := core.ScoreProfiles(p1, p2)
	delta := 0.0001
	if score < expected-delta || score > expected+delta {
		t.Errorf("Expected score %.2f, got %.2f", expected, score)
	}
	
	// Test reverse order to ensure consistency
	scoreReverse := core.ScoreProfiles(p2, p1)
	if scoreReverse < expected-delta || scoreReverse > expected+delta {
		t.Errorf("Reverse order: Expected score %.2f, got %.2f", expected, scoreReverse)
	}
}

// TestScoreProfiles_ExtremeAgeGap tests scoring with extreme age differences
func TestScoreProfiles_ExtremeAgeGap(t *testing.T) {
	p1 := data.Profile{
		ID:        "user1",
		Age:       18,
		Interests: []string{"music", "art"},
	}
	
	p2 := data.Profile{
		ID:        "user2",
		Age:       99,
		Interests: []string{"music", "art"},
	}
	
	// The expected score would be:
	// - Age score: 100 - 81 = 19 (since |18-99| = 81)
	// - Interest score: 2 * 10 = 20 (for "music" and "art")
	// Total: 19 + 20 = 39
	expected := 39.0
	
	score := core.ScoreProfiles(p1, p2)
	delta := 0.0001
	if score < expected-delta || score > expected+delta {
		t.Errorf("Expected score %.2f, got %.2f", expected, score)
	}
}

// TestScoreProfiles_EmptyInterests tests scoring when interest lists are empty
func TestScoreProfiles_EmptyInterests(t *testing.T) {
	testCases := []struct {
		name     string
		p1       data.Profile
		p2       data.Profile
		expected float64
	}{
		{
			name: "Both empty interests",
			p1: data.Profile{
				ID:        "user1",
				Age:       30,
				Interests: []string{},
			},
			p2: data.Profile{
				ID:        "user2",
				Age:       30,
				Interests: []string{},
			},
			expected: 100.0, // Perfect age match, no interests to compare
		},
		{
			name: "First empty interests",
			p1: data.Profile{
				ID:        "user1",
				Age:       30,
				Interests: []string{},
			},
			p2: data.Profile{
				ID:        "user2",
				Age:       30,
				Interests: []string{"music", "art"},
			},
			expected: 100.0, // Perfect age match, no common interests
		},
		{
			name: "Second empty interests",
			p1: data.Profile{
				ID:        "user1",
				Age:       30,
				Interests: []string{"music", "art"},
			},
			p2: data.Profile{
				ID:        "user2",
				Age:       30,
				Interests: []string{},
			},
			expected: 100.0, // Perfect age match, no common interests
		},
	}
	
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			score := core.ScoreProfiles(tc.p1, tc.p2)
			delta := 0.0001
			if score < tc.expected-delta || score > tc.expected+delta {
				t.Errorf("Expected score %.2f, got %.2f", tc.expected, score)
			}
		})
	}
}

// TestScoreProfiles_DuplicateInterests tests scoring when there are duplicate interests
func TestScoreProfiles_DuplicateInterests(t *testing.T) {
	p1 := data.Profile{
		ID:        "user1",
		Age:       30,
		Interests: []string{"music", "art", "music", "travel"},
	}
	
	p2 := data.Profile{
		ID:        "user2",
		Age:       30,
		Interests: []string{"music", "music", "travel", "cooking"},
	}
	
	// The expected score would be:
	// - Age score: 100 (perfect age match)
	// - Interest score: 3 * 10 = 30 (for "music", "music", "travel")
	// Total: 100 + 30 = 130
	expected := 130.0
	
	score := core.ScoreProfiles(p1, p2)
	delta := 0.0001
	if score < expected-delta || score > expected+delta {
		t.Errorf("Expected score %.2f, got %.2f", expected, score)
	}
}

// TestScoreProfiles_ConsistencyCheck tests if scoring is consistent regardless of the order
func TestScoreProfiles_ConsistencyCheck(t *testing.T) {
	testCases := []struct {
		name string
		p1   data.Profile
		p2   data.Profile
	}{
		{
			name: "Standard profiles",
			p1: data.Profile{
				ID:        "user1",
				Age:       25,
				Interests: []string{"music", "art", "travel"},
			},
			p2: data.Profile{
				ID:        "user2",
				Age:       30,
				Interests: []string{"music", "cooking", "travel"},
			},
		},
		{
			name: "Edge case: large age gap",
			p1: data.Profile{
				ID:        "user3",
				Age:       20,
				Interests: []string{"music", "sports"},
			},
			p2: data.Profile{
				ID:        "user4",
				Age:       60,
				Interests: []string{"music", "reading"},
			},
		},
		{
			name: "Edge case: completely different interests",
			p1: data.Profile{
				ID:        "user5",
				Age:       35,
				Interests: []string{"hiking", "swimming", "cycling"},
			},
			p2: data.Profile{
				ID:        "user6",
				Age:       35,
				Interests: []string{"gaming", "movies", "cooking"},
			},
		},
	}
	
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			score1 := core.ScoreProfiles(tc.p1, tc.p2)
			score2 := core.ScoreProfiles(tc.p2, tc.p1)
			
			delta := 0.0001
			if score1 < score2-delta || score1 > score2+delta {
				t.Errorf("Scores not consistent: %.2f vs %.2f", score1, score2)
			}
		})
	}
}

// TestScoreProfiles_ManyInterests tests scoring with a large number of interests
func TestScoreProfiles_ManyInterests(t *testing.T) {
	// Create profiles with many interests
	p1 := data.Profile{
		ID:        "user1",
		Age:       30,
		Interests: []string{
			"music", "art", "travel", "reading", "cooking", 
			"hiking", "gaming", "movies", "fitness", "photography",
			"dancing", "swimming", "cycling", "programming", "painting",
		},
	}
	
	p2 := data.Profile{
		ID:        "user2",
		Age:       30,
		Interests: []string{
			"music", "art", "travel", "reading", "cooking", 
			"baking", "gardening", "yoga", "meditation", "pottery",
			"knitting", "sewing", "singing", "writing", "chess",
		},
	}
	
	// The expected score would be:
	// - Age score: 100 (perfect age match)
	// - Interest score: 5 * 10 = 50 (for music, art, travel, reading, cooking)
	// Total: 100 + 50 = 150
	expected := 150.0
	
	score := core.ScoreProfiles(p1, p2)
	delta := 0.0001
	if score < expected-delta || score > expected+delta {
		t.Errorf("Expected score %.2f, got %.2f", expected, score)
	}
}

// TestScoreProfiles_VerifyMathFormula tests that the scoring follows the expected formula
func TestScoreProfiles_VerifyMathFormula(t *testing.T) {
	testCases := []struct {
		name               string
		age1               int
		age2               int
		commonInterestCount int
	}{
		{name: "Same age, no interests", age1: 25, age2: 25, commonInterestCount: 0},
		{name: "10 year gap, 2 interests", age1: 25, age2: 35, commonInterestCount: 2},
		{name: "20 year gap, 5 interests", age1: 20, age2: 40, commonInterestCount: 5},
		{name: "No gap, 10 interests", age1: 30, age2: 30, commonInterestCount: 10},
		{name: "Extreme gap, many interests", age1: 18, age2: 99, commonInterestCount: 20},
	}
	
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Create interests arrays
			interests1 := make([]string, tc.commonInterestCount)
			interests2 := make([]string, tc.commonInterestCount)
			for i := 0; i < tc.commonInterestCount; i++ {
				interests1[i] = "interest" + string(rune('a'+i))
				interests2[i] = "interest" + string(rune('a'+i))
			}
			
			p1 := data.Profile{
				ID:        "user1",
				Age:       tc.age1,
				Interests: interests1,
			}
			
			p2 := data.Profile{
				ID:        "user2",
				Age:       tc.age2,
				Interests: interests2,
			}
			
			// Calculate expected score using the formula:
			// score = (100 - |age1 - age2|) + (commonInterestCount * 10)
			ageGap := tc.age1 - tc.age2
			if ageGap < 0 {
				ageGap = -ageGap
			}
			
			ageScore := 100 - float64(ageGap)
			interestScore := float64(tc.commonInterestCount) * 10
			expected := ageScore + interestScore
			
			score := core.ScoreProfiles(p1, p2)
			delta := 0.0001
			if score < expected-delta || score > expected+delta {
				t.Errorf("Expected score %.2f, got %.2f", expected, score)
			}
		})
	}
}