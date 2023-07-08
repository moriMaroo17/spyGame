package main

import (
	"bufio"
	"embed"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

//go:embed locations.txt
var locations embed.FS

func main() {
	fmt.Print("Введите количество игроков: ")
	reader := bufio.NewReader(os.Stdin)
	// ReadString will block until the delimiter is entered
	playersNumber, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal("Возникла ошибка чтения. Попробуйте еще раз", err)
	}

	// remove the delimeter from the string
	playersNumber = strings.TrimSuffix(playersNumber, "\n")
	playersNumberInt, err := strconv.Atoi(playersNumber)
	if err != nil {
		log.Fatal("Введите корректное количество игроков")
	}
	players, err := playersGenerator(playersNumberInt)
	if err != nil {
		log.Fatal("Возникла ошибка чтения. Попробуйте еще раз", err)
	}

	if err := choseSpy(players); err != nil {
		log.Fatal("Возникла ошибка выбора роли. Попробуйте еще раз")
	}

	fmt.Printf("%v", players)

	file, err := locations.Open("locations.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	locationsArray := make([]string, 0)

	for scanner.Scan() {
		locationsArray = append(locationsArray, scanner.Text())
		//fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	location, err := choseWord(locationsArray)
	if err != nil {
		log.Fatal("Ошибка выбора локации. Попробуйте еще раз")
	}

	for _, player := range players {
		CallClear()
		fmt.Printf("Следующий игрок: %v\n", player.Name)
		reader.ReadString('\n')
		if player.IsSpy {
			fmt.Println("Вы шпион!")
		} else {
			fmt.Println(location)
		}
		reader.ReadString('\n')
	}
	CallClear()
}

func playersGenerator(playersNumber int) ([]Player, error) {
	players := make([]Player, playersNumber)
	for i := 0; i < playersNumber; i++ {
		fmt.Print("Введите имя игрока: ")
		reader := bufio.NewReader(os.Stdin)
		name, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Возникла ошибка чтения. Попробуйте еще раз", err)
			return nil, err
		}
		players[i] = Player{
			ID:    i,
			Name:  strings.TrimSuffix(name, "\n"),
			IsSpy: false,
		}
	}
	return players, nil
}
