package models

import (
	"math/rand"
	"time"
)

func Init() {
	rand.Seed(time.Now().UnixNano())
}
