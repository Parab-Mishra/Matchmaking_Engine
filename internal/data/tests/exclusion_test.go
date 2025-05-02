package data_test

import (
	"reflect"
	"testing"

	"github.com/Parab-Mishra/Matchmaking_Engine/internal/data"
)

// TestCommonInterests tests the logic for finding common interests
func TestCommonInterests(t *testing.T) {
	testCases := []struct {
		name     string
		list1    []string
		list2    []string
		expected []string
	}{
		{
			name:     "Identical Lists",
			list1:    []string{"music", "art", "travel"},
			list2:    []string{"music", "art", "travel"},
			expected: []string{"music", "art", "travel"},
		},
		{
			name:     "No Common Interests",
			list1:    []string{"music", "art", "travel"},
			list2:    []string{"sports", "gaming", "cooking"},
			expected: []string{},
		},
		{
			name:     "Some Common Interests",
			list1:    []string{"music", "art", "travel", "fitness"},
			list2:    []string{"music", "cooking", "travel", "tech"},
			expected: []string{"music", "travel"},
		},
		{
			name:     "Empty Lists",
			list1:    []string{},
			list2:    []string{},
			expected: []string{},
		},
		{
			name:     "One Empty List",
			list1:    []string{"music", "art"},
			list2:    []string{},
			expected: []string{},
		},
		{
			name:     "Duplicate Interests",
			list1:    []string{"music", "art", "music"},
			list2:    []string{"music", "tech", "music"},
			expected: []string{"music", "music"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := data.CommonInterests(tc.list1, tc.list2)
			
			// Compare slices
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected %v, got %v", tc.expected, result)
			}
		})
	}
}

// TestIsEligible tests the logic for determining match eligibility
func TestIsEligible(t *testing.T) {
	testCases := []struct {
		name     string
		user1ID  string
		user2ID  string
		expected bool
	}{
		{
			name:     "Different Users",
			user1ID:  "user1",
			user2ID:  "user2",
			expected: true,
		},
		{
			name:     "Same User",
			user1ID:  "user1",
			user2ID:  "user1",
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := data.IsEligible(tc.user1ID, tc.user2ID)
			if result != tc.expected {
				t.Errorf("Expected %v, got %v", tc.expected, result)
			}
		})
	}
}