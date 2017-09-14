package al

import (
	"math/rand"
	"time"
)

func Shuffler(r []int) {
	for i:= len(r)-1; i >=0; i--{
		rand.Seed(time.Now().Unix())
		j := rand.Intn(i+1)
		r[j], r[i] = r[i], r[j]
	}
}