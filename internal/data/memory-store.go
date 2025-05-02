package data

import (
    "sync"
)

var (
    profiles     = make(map[string]Profile)
    geoBuckets   = make(map[string][]Profile)
    profileMutex sync.RWMutex
)

func InitStore() {
    profiles = make(map[string]Profile)
    geoBuckets = make(map[string][]Profile)
}

func StoreProfile(p Profile, geohash string) {
    profileMutex.Lock()
    defer profileMutex.Unlock()
    profiles[p.ID] = p
    geoBuckets[geohash] = append(geoBuckets[geohash], p)
}

func GetProfilesInGeohash(gh string) []Profile {
    profileMutex.RLock()
    defer profileMutex.RUnlock()
    return geoBuckets[gh]
}

func GetProfile(id string) (Profile, bool) {
    profileMutex.RLock()
    defer profileMutex.RUnlock()
    p, ok := profiles[id]
    return p, ok
}

// GetAllProfiles retrieves all user profiles stored in memory.
func GetAllProfiles() []Profile {
    var allProfiles []Profile
    for _, profile := range profiles {
        allProfiles = append(allProfiles, profile)
    }
    return allProfiles
}
