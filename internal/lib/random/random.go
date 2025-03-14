package random

import (
	"math/rand"
	"time"
)

func NewRandomString(length int) string {
	rndm := rand.New(rand.NewSource(time.Now().UnixNano()))

	chairs := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ" + "abcdefghijklmnopqrstuvwxyz" + "0123456789")

	b := make([]rune, length)
	for i := range b {
		b[i] = chairs[rndm.Intn(len(chairs))]
	}
	return string(b)
}
