package patterns

import (
	"log"
	"math/rand"
	"time"

	"github.com/leofigy/gocuc/utils"
)

func Tell(msg string) {
	for {
		log.Println("go id", utils.GetGID(), msg)
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}
}
