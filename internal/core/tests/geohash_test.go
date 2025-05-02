package core_test

import (
	"testing"

	"github.com/Parab-Mishra/Matchmaking_Engine/internal/core"
	"github.com/Parab-Mishra/Matchmaking_Engine/internal/data"
	"github.com/mmcloughlin/geohash"
)

// TestGetGeohash_Precision verifies that the geohash precision is consistent
func TestGetGeohash_Precision(t *testing.T) {
	testCases := []struct {
		name        string
		profile     data.Profile
		expPrecision int
	}{
		{
			name: "Standard Location",
			profile: data.Profile{
				ID:       "user1",
				Location: data.Location{Lat: 19.076, Lon: 72.8777},
			},
			expPrecision: 5,
		},
		{
			name: "Extreme Northern Location",
			profile: data.Profile{
				ID:       "user2",
				Location: data.Location{Lat: 89.9, Lon: 45.0},
			},
			expPrecision: 5,
		},
		{
			name: "Extreme Southern Location",
			profile: data.Profile{
				ID:       "user3",
				Location: data.Location{Lat: -89.9, Lon: 45.0},
			},
			expPrecision: 5,
		},
		{
			name: "Zero Coordinates",
			profile: data.Profile{
				ID:       "user4",
				Location: data.Location{Lat: 0, Lon: 0},
			},
			expPrecision: 5,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := core.GetGeohash(tc.profile)
			if len(result) != tc.expPrecision {
				t.Errorf("Expected geohash with precision %d, got %s with length %d", 
					tc.expPrecision, result, len(result))
			}
		})
	}
}

// TestGetGeohash_CompareWithLibrary directly compares our implementation with the library function
func TestGetGeohash_CompareWithLibrary(t *testing.T) {
	testLocations := []data.Location{
		{Lat: 19.076, Lon: 72.8777},  // Mumbai
		{Lat: 28.6139, Lon: 77.2090}, // New Delhi
		{Lat: 12.9716, Lon: 77.5946}, // Bangalore
		{Lat: 22.5726, Lon: 88.3639}, // Kolkata
		{Lat: 17.3850, Lon: 78.4867}, // Hyderabad
	}
	
	for i, loc := range testLocations {
		profile := data.Profile{
			ID:       "test_user",
			Location: loc,
		}
		
		ourHash := core.GetGeohash(profile)
		libHash := geohash.EncodeWithPrecision(loc.Lat, loc.Lon, 5)
		
		if ourHash != libHash {
			t.Errorf("Test case %d: Expected %s (library), got %s (our implementation)", 
				i, libHash, ourHash)
		}
	}
}

// TestGetGeohash_EdgeCases tests geohash behavior with edge case coordinates
func TestGetGeohash_EdgeCases(t *testing.T) {
	testCases := []struct {
		name    string
		profile data.Profile
	}{
		{
			name: "Max Latitude",
			profile: data.Profile{
				ID:       "edge1",
				Location: data.Location{Lat: 90.0, Lon: 45.0},
			},
		},
		{
			name: "Min Latitude",
			profile: data.Profile{
				ID:       "edge2",
				Location: data.Location{Lat: -90.0, Lon: 45.0},
			},
		},
		{
			name: "Max Longitude",
			profile: data.Profile{
				ID:       "edge3",
				Location: data.Location{Lat: 45.0, Lon: 180.0},
			},
		},
		{
			name: "Min Longitude",
			profile: data.Profile{
				ID:       "edge4",
				Location: data.Location{Lat: 45.0, Lon: -180.0},
			},
		},
		{
			name: "Date Line Crossing East",
			profile: data.Profile{
				ID:       "edge5",
				Location: data.Location{Lat: 45.0, Lon: 179.9999},
			},
		},
		{
			name: "Date Line Crossing West",
			profile: data.Profile{
				ID:       "edge6",
				Location: data.Location{Lat: 45.0, Lon: -179.9999},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// We're just testing that these don't panic
			hash := core.GetGeohash(tc.profile)
			expected := geohash.EncodeWithPrecision(tc.profile.Location.Lat, tc.profile.Location.Lon, 5)
			if hash != expected {
				t.Errorf("Expected %s, got %s", expected, hash)
			}
		})
	}
}