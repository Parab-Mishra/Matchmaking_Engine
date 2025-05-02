# Matchmaking_Engine
A high-performance in-memory matchmaking engine for a dating app.

HOW TO RUN
//Install dependencies
go mod tidy

//Run the app
go run ./cmd/server/main.go 

//Feed the dummy data programmatically
go run ./cmd/tools/bulk-insert.go 500

//Run API test using shell script
chmod +x test-api.sh
./test-api.sh

//Run unit test for geohash logic, exclusions, and match scoring
go test -v ./internal/core/tests/
go test -v ./internal/data/tests/

===================================================================================

DESCRIPTION
This application is a Go-based matchmaking service which:

-> Stores user profiles in memory
-> Uses geohash for location-based matching
-> Has API to create profiles and find matches
-> Contains a bulk-insert tool for inserting data for testing


ARCHITECTURE DECISIONS
* Modular Approach
* API layer
* data layer
* core logic layer

PRE-COMPUTATION DESIGN
-> Find candidates in the same geohash region
-> Checks gender preferences, interests, age etc
-> Calculates a compatibility score for each eligible candidate
-> Store the results in a match index for fast retrieval

UPGRADES TO CONSIDER IN PRODUCTION
* Use microservices (profile/match/auth)
* Implement authentication
* Data storage in persistent database
* Introduce more matching criteria
