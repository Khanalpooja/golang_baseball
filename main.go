package main

// add preamble in the code
// add your name, A#, seciton, details to follow in the progran
import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Entry point of the program // main function
func main() {
	var fileName string
	fmt.Println("Enter filename")
	fmt.Scan(&fileName)
	players, errors := getPlayers(fileName)
	printPlayerReport(players)
	printErrorReport(errors)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// Read from the file and get each player
func getPlayers(fileName string) ([]Player, []error) {
	var players []Player
	var errors []error

	file, err := os.Open(fileName)
	check(err)

	scanner := bufio.NewScanner(file)
	lineCount := 0
	for scanner.Scan() {
		lineCount++
		s := strings.Split(scanner.Text(), " ")
		player, err := getPlayer(s, lineCount)
		if err != nil {
			errors = append(errors, err)
		} else {
			players = append(players, player)
		}
	}
	file.Close()

	return players, errors
}
