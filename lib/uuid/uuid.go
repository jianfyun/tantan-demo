package uuid

import (
	"math/rand"
	"strconv"
	"time"

	"github.com/satori/go.uuid"
)

// New returns a new Universally Unique Identifier based on MAC address, timestamp and random number.
func New() string {
	ns, _ := uuid.NewV1()
	rand.Seed(time.Now().UnixNano())
	name := strconv.FormatInt(rand.Int63(), 10)
	return uuid.NewV5(ns, name).String()
}
