package api

import (
    "encoding/json"
    "net/http"
    "fmt"
    "strconv"

    "github.com/gorilla/mux"
    "github.com/Parab-Mishra/Matchmaking_Engine/internal/core"
    "github.com/Parab-Mishra/Matchmaking_Engine/internal/data"
)

// PaginatedMatches represents the response structure for paginated match results
type PaginatedMatches struct {
    Results     []data.MatchResult `json:"results"`
    Page        int                `json:"page"`
    PageSize    int                `json:"pageSize"`
    TotalCount  int                `json:"totalCount"`
    TotalPages  int                `json:"totalPages"`
}

func CreateProfileHandler(w http.ResponseWriter, r *http.Request) {
    var profile data.Profile
    if err := json.NewDecoder(r.Body).Decode(&profile); err != nil {
        http.Error(w, "Invalid input", http.StatusBadRequest)
        return
    }

    // Validate gender and genderSeeking fields
    if profile.Gender != "M" && profile.Gender != "F" {
        http.Error(w, "Gender must be 'M' or 'F'", http.StatusBadRequest)
        return
    }

    if profile.GenderSeeking != "M" && profile.GenderSeeking != "F" && profile.GenderSeeking != "B" {
        http.Error(w, "GenderSeeking must be 'M', 'F', or 'B'", http.StatusBadRequest)
        return
    }

    geohash := core.GetGeohash(profile)
    data.StoreProfile(profile, geohash)
    core.PrecomputeMatches(profile)

    w.WriteHeader(http.StatusCreated)
}

func MatchHandler(w http.ResponseWriter, r *http.Request) {
    
    // Get user ID from URL
    id := mux.Vars(r)["id"]
    fmt.Printf("Matching for user: %s\n", id)
    
    // Parse query parameters for pagination
    pageStr := r.URL.Query().Get("page")
    pageSizeStr := r.URL.Query().Get("pageSize")
    genderFilter := r.URL.Query().Get("gender")
    
    // Default values
    page := 1
    pageSize := 10
    
    // Try to parse page number
    if pageStr != "" {
        if p, err := strconv.Atoi(pageStr); err == nil && p > 0 {
            page = p
        }
    }
    
    // Try to parse page size
    if pageSizeStr != "" {
        if ps, err := strconv.Atoi(pageSizeStr); err == nil && ps > 0 && ps <= 100 {
            pageSize = ps
        }
    }
    
    // Validate gender filter if provided
    if genderFilter != "" && genderFilter != "M" && genderFilter != "F" && genderFilter != "B" {
        http.Error(w, "Gender filter must be 'M', 'F', or 'B'", http.StatusBadRequest)
        return
    }
    
    // Get all matches for the user
    allMatches := core.GetAllMatches(id)
    
    // Apply gender filter if specified
    if genderFilter != "" {
        filteredMatches := []data.MatchResult{}
        for _, match := range allMatches {
            // Get the profile for this match
            if matchProfile, exists := data.GetProfile(match.ID); exists {
                // Only include if it matches the gender filter
                if genderFilter == "B" || matchProfile.Gender == genderFilter {
                    filteredMatches = append(filteredMatches, match)
                }
            }
        }
        allMatches = filteredMatches
    }
    
    totalCount := len(allMatches)
    
    // Calculate pagination values
    startIndex := (page - 1) * pageSize
    endIndex := startIndex + pageSize
    if endIndex > totalCount {
        endIndex = totalCount
    }
    
    // Handle out-of-range page number
    if startIndex >= totalCount && totalCount > 0 {
        http.Error(w, "Page number out of range", http.StatusBadRequest)
        return
    }
    
    // Slice the results for the requested page
    var pageResults []data.MatchResult
    if startIndex < totalCount {
        pageResults = allMatches[startIndex:endIndex]
    } else {
        pageResults = []data.MatchResult{}
    }
    
    // Calculate total pages
    totalPages := (totalCount + pageSize - 1) / pageSize
    if totalPages == 0 {
        totalPages = 1
    }
    
    // Prepare the response
    response := PaginatedMatches{
        Results:    pageResults,
        Page:       page,
        PageSize:   pageSize,
        TotalCount: totalCount,
        TotalPages: totalPages,
    }
    
    fmt.Printf("Returning %d matches for user %s (page %d/%d)\n", 
        len(pageResults), id, page, totalPages)
    
    // Send the response
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
}

// GetAllProfiles handles the request to fetch all user profiles
func GetAllProfiles(w http.ResponseWriter, r *http.Request) {
    // Fetch all profiles from memory
    profiles := data.GetAllProfiles()

    // Set header as JSON
    w.Header().Set("Content-Type", "application/json")

    // Respond with JSON of profiles
    err := json.NewEncoder(w).Encode(profiles)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    fmt.Printf("Fetched %d profiles", len(profiles)) // For debugging
}