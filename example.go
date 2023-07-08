package main

import (
	"fmt"
	"math/rand"
	"time"
)

var array []string

type Player struct {
	ID    int
	Name  string
	IsSpy bool
}

func choseWord(locations []string) (string, error) {
	rand.Seed(time.Now().UnixNano())
	if len(locations) == 0 {
		return "", fmt.Errorf("")
	}
	randomIndex := rand.Intn(len(locations))
	return locations[randomIndex], nil
}

func choseSpy(players []Player) error {
	rand.Seed(time.Now().UnixNano())
	if len(players) == 0 {
		return fmt.Errorf("")
	}
	randomIndex := rand.Intn(len(players))
	players[randomIndex].IsSpy = true
	return nil
}
