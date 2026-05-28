package util

import (
	"time"
	"math/rand"
	"github.com/oklog/ulid/v2"
);

func GenerateRandomULID() ulid.ULID {
	entropy := ulid.Monotonic(rand.New(rand.NewSource(time.Now().UnixNano())), 0)
    leagueID := ulid.MustNew(ulid.Timestamp(time.Now()), entropy)
	return leagueID;
}