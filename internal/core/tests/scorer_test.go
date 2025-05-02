package core_test

import (
	"testing"

	"github.com/Parab-Mishra/Matchmaking_Engine/internal/core"
	"github.com/Parab-Mishra/Matchmaking_Engine/internal/data"
)

// TestScoreProfiles tests the profile matching scoring algorithm
func TestScoreProfiles(t *testing.T) {
	testCases := []struct {
		name     string
		profile1 data.Profile
		profile2 data.Profile
		expected float64
	}{
		{
			name: "Identical Age and Interests",
			profile1: data.Profile{
				ID:        "user1",
				Age:       30,
				Interests: []string{"music", "art", "travel"},
			},
			profile2: data.Profile{
				ID:        "user2",
				Age:       30,
				Interests: []string{"music", "art", "travel"},
			},
			expected: 100 + (3 * 10), // 100 (perfect age match) + 30 (3 common interests)
		},
		{
			name: "Different Age, Same Interests",
			profile1: data.Profile{
				ID:        "user1",
				Age:       25,
				Interests: []string{"music", "art", "travel"},
			},
			profile2: data.Profile{
				ID:        "user2",
				Age:       35,
				Interests: []string{"music", "art", "travel"},
			},
			expected: 90 + (3 * 10), // 90 (10 years age difference) + 30 (3 common interests)
		},
		{
			name: "Same Age, No Common Interests",
			profile1: data.Profile{
				ID:        "user1",
				Age:       30,
				Interests: []string{"music", "art", "travel"},
			},
			profile2: data.Profile{
				ID:        "user2",
				Age:       30,
				Interests: []string{"sports", "gaming", "cooking"},
			},
			expected: 100 + (0 * 10), // 100 (perfect age match) + 0 (0 common interests)
		},
		{
			name: "Different Age, Some Common Interests",
			profile1: data.Profile{
				ID:        "user1",
				Age:       25,
				Interests: []string{"music", "art", "travel", "tech"},
			},
			profile2: data.Profile{
				ID:        "user2",
				Age:       30,
				Interests: []string{"music", "tech", "fitness"},
			},
			expected: 95 + (2 * 10), // 95 (5 years age difference) + 20 (2 common interests)
		},
		{
			name: "Large Age Gap, Many Common Interests",
			profile1: data.Profile{
				ID:        "user1",
				Age:       20,
				Interests: []string{"music", "art", "travel", "tech", "fitness"},
			},
			profile2: data.Profile{
				ID:        "user2",
				Age:       40,
				Interests: []string{"music", "art", "travel", "tech", "fitness"},
			},
			expected: 80 + (5 * 10), // 80 (20 years age difference) + 50 (5 common interests)
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := core.ScoreProfiles(tc.profile1, tc.profile2)
			
			// Check for floating point precision within a small delta
			delta := 0.0001
			if result < tc.expected-delta || result > tc.expected+delta {
				t.Errorf("Expected score %.2f, got %.2f", tc.expected, result)
			}
		})
	}
}

// TestMatchResultOrdering tests storing and sorting match results
func TestStoreAndRetrieveMatchResults(t *testing.T) {
	// Initialize match index (normally done by InitStore)
	data.InitStore()

	// Create test data
	userID := "test-user"
	matches := []data.MatchResult{
		{ID: "user1", Score: 80.0},
		{ID: "user2", Score: 120.0},
		{ID: "user3", Score: 90.0},
		{ID: "user4", Score: 110.0},
		{ID: "user5", Score: 70.0},
		{ID: "user6", Score: 100.0},
	}

	// Store match results
	data.StoreMatchResults(userID, matches)

	// Get match results
	results := data.GetMatchResults(userID)

	// Verify results are sorted by score (descending) and limited to 5
	if len(results) > 5 {
		t.Errorf("Expected max 5 results, got %d", len(results))
	}

	// Check ordering (should be descending by score)
	for i := 0; i < len(results)-1; i++ {
		if results[i].Score < results[i+1].Score {
			t.Errorf("Results not properly sorted at position %d: %.2f should be >= %.2f", 
				i, results[i].Score, results[i+1].Score)
		}
	}

	// Check highest score is first
	if len(results) > 0 && results[0].ID != "user2" {
		t.Errorf("Expected top match ID to be user2, got %s", results[0].ID)
	}
}