package main

import (
    "bytes"
    "encoding/json"
    "fmt"
    "math/rand"
    "net/http"
    "strconv"
    "time"
	"os"

    "github.com/bxcodec/faker/v4"
)

type Location struct {
    Lat float64 `json:"lat"`
    Lon float64 `json:"lon"`
}

type Profile struct {
    ID        string   `json:"id"`
    Age       int      `json:"age"`
    Gender    string   `json:"gender"`
    Location  Location `json:"location"`
    Interests []string `json:"interests"`
	Bio       string   `json:"bio,omitempty"` // Optional: add in your profile.go if supported
}

var sampleInterests = []string{
    "music", "art", "travel", "reading", "yoga", "movies", "fitness", "cooking", "tech", "fashion",
}

func randomLatLonIndia() (float64, float64) {
    // Rough bounding box around Mumbai/India
    lat := randFloat(18.9, 19.3) // Latitude between 18.9 and 19.3
    lon := randFloat(72.7, 73.1) // Longitude between 72.7 and 73.1
    return lat, lon
}

func randomInterests() []string {
    n := rand.Intn(3) + 2 // 2 to 4 interests
    rand.Shuffle(len(sampleInterests), func(i, j int) {
        sampleInterests[i], sampleInterests[j] = sampleInterests[j], sampleInterests[i]
    })
    return sampleInterests[:n]
}

func randFloat(min, max float64) float64 {
    return min + rand.Float64()*(max-min)
}

func generateProfile(i int) Profile {
    gender := []string{"M", "F"}[rand.Intn(2)]
    lat, lon := randomLatLonIndia()
    return Profile{
        ID:        "user" + strconv.Itoa(i),
        Age:       rand.Intn(20) + 20, // 20-39
        Gender:    gender,
        Location:  Location{Lat: lat, Lon: lon},
        Interests: randomInterests(),
		Bio:       faker.Sentence(),
    }
}

func main() {
    rand.Seed(time.Now().UnixNano())
    total := 10 // Change this or use flag for CLI arg
	if len(os.Args) > 1 {
		if n, err := strconv.Atoi(os.Args[1]); err == nil {
			total = n
		}
	}

    for i := 1; i <= total; i++ {
        profile := generateProfile(i)
        payload, _ := json.Marshal(profile)

        resp, err := http.Post("http://localhost:8080/profiles", "application/json", bytes.NewBuffer(payload))
        if err != nil {
            fmt.Printf("Error inserting user%d: %v\n", i, err)
            continue
        }
        fmt.Printf("Inserted %s (status: %d)\n", profile.ID, resp.StatusCode)
        resp.Body.Close()
    }
}
