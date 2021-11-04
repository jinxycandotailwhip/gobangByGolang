package test

import (
	"gobang/negative_max"
	"log"
	"testing"
)

func TestAI(t *testing.T) {
	pos := negative_max.AI()
	log.Default().Println("pos:", pos)
}
