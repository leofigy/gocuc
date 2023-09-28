package patterns

import (
	"fmt"
	"log"
	"testing"
)

var dic = []string{
	"aitana",
	"melanie",
	"beatriz",
	"angel",
	"salvador",
	"nestor",
	"miguel",
	"jonas",
	"maria",
	"dracula",
	"guadalupe",
	"dakota",
	"werewolf",
	"witch",
}

var fruits = []string{
	"apple",
	"banana",
	"blueberry",
	"strawberry",
	"lemon",
	"orange",
}

func TestMagicBox(t *testing.T) {

	gen := MagicBox(dic)

	for i := range gen {
		log.Println(i)
	}
}

func TestMerge(t *testing.T) {
	names := MagicBox(dic)
	fruit := MagicBox(fruits)

	merged := Merge(names, fruit)

	for i := 0; i < 10; i++ {
		fmt.Println("from merged channel", <-merged)
	}

}
