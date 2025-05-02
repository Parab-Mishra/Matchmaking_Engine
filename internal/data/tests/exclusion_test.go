package data_test

import (
	"testing"

	"github.com/Parab-Mishra/Matchmaking_Engine/internal/data"
)

// Setup helper function to initialize test profiles
func setupProfiles() {
	data.InitStore()
	
	// Create test profiles with different combinations of genders and preferences
	profiles := []data.Profile{
		{
			ID: "male_seeking_female",
			Gender: "M",
			GenderSeeking: "F",
			Age: 25,
			Location: data.Location{Lat: 19.0, Lon: 72.0},
		},
		{
			ID: "male_seeking_male",
			Gender: "M",
			GenderSeeking: "M",
			Age: 25,
			Location: data.Location{Lat: 19.0, Lon: 72.0},
		},
		{
			ID: "male_seeking_both",
			Gender: "M",
			GenderSeeking: "B",
			Age: 25,
			Location: data.Location{Lat: 19.0, Lon: 72.0},
		},
		{
			ID: "female_seeking_male",
			Gender: "F",
			GenderSeeking: "M",
			Age: 25,
			Location: data.Location{Lat: 19.0, Lon: 72.0},
		},
		{
			ID: "female_seeking_female",
			Gender: "F",
			GenderSeeking: "F",
			Age: 25,
			Location: data.Location{Lat: 19.0, Lon: 72.0},
		},
		{
			ID: "female_seeking_both",
			Gender: "F",
			GenderSeeking: "B",
			Age: 25,
			Location: data.Location{Lat: 19.0, Lon: 72.0},
		},
	}
	
	// Store the profiles
	for _, p := range profiles {
		data.StoreProfile(p, "dummy")
	}
}

// TestIsEligible_GenderPreferences tests all combinations of gender preferences
func TestIsEligible_GenderPreferences(t *testing.T) {
	setupProfiles()
	
	testCases := []struct {
		name     string
		user1ID  string
		user2ID  string
		expected bool
	}{
		// Test matching gender preferences
		{
			name:     "M seeking F + F seeking M = Match",
			user1ID:  "male_seeking_female",
			user2ID:  "female_seeking_male",
			expected: true,
		},
		{
			name:     "M seeking B + F seeking M = Match",
			user1ID:  "male_seeking_both",
			user2ID:  "female_seeking_male",
			expected: true,
		},
		{
			name:     "F seeking B + M seeking F = Match",
			user1ID:  "female_seeking_both",
			user2ID:  "male_seeking_female",
			expected: true,
		},
		
		// Test non-matching gender preferences
		{
			name:     "M seeking F + M seeking M = No Match",
			user1ID:  "male_seeking_female",
			user2ID:  "male_seeking_male",
			expected: false,
		},
		{
			name:     "F seeking M + F seeking F = No Match",
			user1ID:  "female_seeking_male",
			user2ID:  "female_seeking_female",
			expected: false,
		},
		{
			name:     "M seeking M + F seeking F = No Match",
			user1ID:  "male_seeking_male",
			user2ID:  "female_seeking_female",
			expected: false,
		},
		{
			name:     "M seeking F + F seeking F = No Match",
			user1ID:  "male_seeking_female",
			user2ID:  "female_seeking_female",
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := data.IsEligible(tc.user1ID, tc.user2ID)
			if result != tc.expected {
				t.Errorf("Expected eligibility to be %v, got %v", tc.expected, result)
			}
		})
	}
}

// TestIsEligible_SameUser verifies that a user cannot match with themselves
func TestIsEligible_SameUser(t *testing.T) {
	setupProfiles()
	
	testCases := []struct {
		name string
		userID string
	}{
		{name: "Male user cannot match self", userID: "male_seeking_both"},
		{name: "Female user cannot match self", userID: "female_seeking_both"},
	}
	
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := data.IsEligible(tc.userID, tc.userID)
			if result != false {
				t.Errorf("Expected user to not be eligible to match with self, got eligible")
			}
		})
	}
}

// TestIsEligible_NonExistentUser checks that non-existent users are not eligible
func TestIsEligible_NonExistentUser(t *testing.T) {
	setupProfiles()
	
	testCases := []struct {
		name string
		user1ID string
		user2ID string
	}{
		{name: "First user doesn't exist", user1ID: "nonexistent_user", user2ID: "male_seeking_both"},
		{name: "Second user doesn't exist", user1ID: "female_seeking_both", user2ID: "nonexistent_user"},
		{name: "Both users don't exist", user1ID: "nonexistent_user1", user2ID: "nonexistent_user2"},
	}
	
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := data.IsEligible(tc.user1ID, tc.user2ID)
			if result != false {
				t.Errorf("Expected users to not be eligible when one or both don't exist, got eligible")
			}
		})
	}
}

// TestIsInterestedInGender_AllCombinations tests all combinations of gender interest
func TestIsInterestedInGender_AllCombinations(t *testing.T) {
	testCases := []struct {
		name     string
		seeking  string
		target   string
		expected bool
	}{
		// "B" = interested in both
		{name: "B is interested in M", seeking: "B", target: "M", expected: true},
		{name: "B is interested in F", seeking: "B", target: "F", expected: true},
		
		// Direct matches
		{name: "M is interested in M", seeking: "M", target: "M", expected: true},
		{name: "F is interested in F", seeking: "F", target: "F", expected: true},
		
		// Mismatches
		{name: "M is not interested in F", seeking: "M", target: "F", expected: false},
		{name: "F is not interested in M", seeking: "F", target: "M", expected: false},
		
		// Invalid values (testing robustness)
		{name: "Invalid seeking X, target M", seeking: "X", target: "M", expected: false},
		{name: "Seeking M, invalid target X", seeking: "M", target: "X", expected: false},
		{name: "Both invalid", seeking: "X", target: "Y", expected: false},
	}
	
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// The isInterestedInGender func is private, but we test it through the public IsEligible
			// by creating profiles with these preferences
			setupProfiles()
			
			// Create test profiles with custom preferences
			user1 := data.Profile{
				ID: "test_user1",
				Gender: "M", // Gender doesn't matter for this test
				GenderSeeking: tc.seeking,
				Age: 25,
				Location: data.Location{Lat: 19.0, Lon: 72.0},
			}
			
			user2 := data.Profile{
				ID: "test_user2",
				Gender: tc.target,
				GenderSeeking: "B", // Set to "B" so this direction always passes
				Age: 25,
				Location: data.Location{Lat: 19.0, Lon: 72.0},
			}
			
			data.StoreProfile(user1, "dummy")
			data.StoreProfile(user2, "dummy")
			
			// The eligibility is affected by both directions, so we're specifically
			// testing just the "seeking" direction here
			expected := tc.expected
			
			// For invalid values, IsEligible would return false regardless
			// of the isInterestedInGender result, so we need to adjust our expectation
			if tc.seeking != "M" && tc.seeking != "F" && tc.seeking != "B" {
				expected = false
			}
			if tc.target != "M" && tc.target != "F" {
				expected = false
			}
			
			result := data.IsEligible("test_user1", "test_user2")
			if result != expected {
				t.Errorf("Expected interest check to be %v, got %v", expected, result)
			}
		})
	}
}