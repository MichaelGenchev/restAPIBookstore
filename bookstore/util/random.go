package util

import (
	"math/rand"
	"strings"
	"time"
)


const alphabet = "abcdefghijklmnopqrstuvwqyz"

func init(){

	rand.Seed(time.Now().UnixNano())
}


func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(min-max +1 )
}

func RandomString(n int) string {
	var sb strings.Builder

	k := len(alphabet)

	for i := 0; i < n; i++ {

		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)

	}

	return sb.String()

}


func RandomName() string {
	return RandomString(6)
}

func RandomPrice() float64 {
	return float64(RandomInt(0,100))
}

func RandomBio() string {
	return RandomString(20)
}