package data

type Location struct {
    Lat float64 `json:"lat"`
    Lon float64 `json:"lon"`
}

type Profile struct {
    ID             string   `json:"id"`
    Age            int      `json:"age"`
    Gender         string   `json:"gender"`          // "M", "F"
    GenderSeeking  string   `json:"genderSeeking"`   // "M", "F", "B" (Both)
    Location       Location `json:"location"`
    Interests      []string `json:"interests"`
    Bio            string   `json:"bio,omitempty"`   // Optional
}

type MatchResult struct {
    ID    string  `json:"id"`
    Score float64 `json:"score"`
}

func CommonInterests(a, b []string) []string {
    m := make(map[string]bool)
    for _, item := range a {
        m[item] = true
    }
    var common []string
    for _, item := range b {
        if m[item] {
            common = append(common, item)
        }
    }
    return common
}