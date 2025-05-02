package core_test

import (
	"testing"

	"github.com/Parab-Mishra/Matchmaking_Engine/internal/core"
	"github.com/Parab-Mishra/Matchmaking_Engine/internal/data"
	"github.com/mmcloughlin/geohash"
)

// TestGetGeohash tests the geohashing functionality
func TestGetGeohash(t *testing.T) {
	testCases := []struct {
		name     string
		profile  data.Profile
		expected string
	}{
		{
			name: "Mumbai City Center",
			profile: data.Profile{
				ID:       "user1",
				Location: data.Location{Lat: 19.076, Lon: 72.8777},
			},
			expected: "te7ud", // Precomputed geohash with precision 5
		},
		{
			name: "New Delhi",
			profile: data.Profile{
				ID:       "user2",
				Location: data.Location{Lat: 28.6139, Lon: 77.2090},
			},
			expected: "ttnfu", // Precomputed geohash with precision 5
		},
		{
			name: "Near Boundary Edge Case",
			profile: data.Profile{
				ID:       "user3",
				Location: data.Location{Lat: 19.0000, Lon: 73.0000},
			},
			expected: geohash.EncodeWithPrecision(19.0000, 73.0000, 5),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := core.GetGeohash(tc.profile)
			if result != tc.expected {
				t.Errorf("Expected geohash %s, got %s", tc.expected, result)
			}
		})
	}
}
