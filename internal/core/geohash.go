package core

import (
    "github.com/mmcloughlin/geohash"
    "github.com/Parab-Mishra/Matchmaking_Engine/internal/data"
)

func GetGeohash(p data.Profile) string {
    return geohash.EncodeWithPrecision(p.Location.Lat, p.Location.Lon, 5)
}
