#!/bin/bash

# Test script for Matchmaking Engine API
# This script requires curl and jq

set -e

echo "Testing Matchmaking Engine API..."
SERVER="http://localhost:8080"

echo -e "\nServer is up! Creating test profiles..."

# Create test profile 1
echo "Creating profile: user1"
curl -s -X POST ${SERVER}/profiles \
  -H "Content-Type: application/json" \
  -d '{
    "id": "user1",
    "age": 28,
    "gender": "M",
    "genderSeeking": "F",
    "location": {
      "lat": 19.076,
      "lon": 72.8777
    },
    "interests": ["music", "travel", "cooking", "reading"],
    "bio": "Software engineer who loves to travel"
  }'
echo

# Create test profile 2
echo "Creating profile: user2"
curl -s -X POST ${SERVER}/profiles \
  -H "Content-Type: application/json" \
  -d '{
    "id": "user2",
    "age": 26,
    "gender": "F",
    "genderSeeking": "M",
    "location": {
      "lat": 19.077,
      "lon": 72.8775
    },
    "interests": ["music", "art", "cooking", "yoga"],
    "bio": "Artist who enjoys yoga and cooking"
  }'
echo

# Create test profile 3 (not compatible with user1)
echo "Creating profile: user3"
curl -s -X POST ${SERVER}/profiles \
  -H "Content-Type: application/json" \
  -d '{
    "id": "user3",
    "age": 30,
    "gender": "M",
    "genderSeeking": "M",
    "location": {
      "lat": 19.075,
      "lon": 72.8780
    },
    "interests": ["music", "travel", "fitness"],
    "bio": "Fitness trainer who loves music"
  }'
echo

# Create test profile 4
echo "Creating profile: user4"
curl -s -X POST ${SERVER}/profiles \
  -H "Content-Type: application/json" \
  -d '{
    "id": "user4",
    "age": 29,
    "gender": "F",
    "genderSeeking": "B",
    "location": {
      "lat": 19.080,
      "lon": 72.8790
    },
    "interests": ["reading", "travel", "movies"],
    "bio": "Book lover who enjoys traveling"
  }'
echo

# Let the server process the profiles
sleep 1

echo -e "\nFetching all profiles..."
curl -s ${SERVER}/profiles | jq . || echo "Failed to get profiles"

echo -e "\nFetching matches for user1..."
curl -s "${SERVER}/match/user1" | jq . || echo "Failed to get matches for user1"

echo -e "\nFetching matches for user2..."
curl -s "${SERVER}/match/user2" | jq . || echo "Failed to get matches for user2"

echo -e "\nFetching matches for user3..."
curl -s "${SERVER}/match/user3" | jq . || echo "Failed to get matches for user3"

echo -e "\nFetching matches for user4..."
curl -s "${SERVER}/match/user4" | jq . || echo "Failed to get matches for user4"

echo -e "\nTesting pagination and filtering..."
echo "Matches for user1 (page 1, pageSize 2):"
curl -s "${SERVER}/match/user1?page=1&pageSize=2" | jq . || echo "Failed"

echo -e "\nMatches for user4 (filtered by gender=M):"
curl -s "${SERVER}/match/user4?gender=M" | jq . || echo "Failed"

echo -e "\nTests completed!"