package patterns

import(
	"testing"
	"time"
)

func TestTell(t *testing.T){
	go Tell("hola")
	go Tell("mundo")

	time.Sleep(5 * time.Second)
}