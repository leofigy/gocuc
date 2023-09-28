package patterns

import (
	"math/rand"
	"time"
)

func MagicBox(msg []string) <-chan string {
	seed := len(msg)
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	c := make(chan string)
	go func() {
		defer close(c)
		for i := 0; i < seed*seed; i++ {
			c <- string(msg[r1.Intn(seed)])
		}
	}()
	return c
}
