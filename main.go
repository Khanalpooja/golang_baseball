// Prepared by: Pooja Khanal, A25288740
// Prepared for : Assignment 1, CS524, Spring 2020
// Submitted on : 02/20/2020

//Program Description:
// Main tasks covered by this program are
// 1. This program calculates the player statistics, batting avaerage, slugging and base percentage.
// 2. Checks if inconsisent data are present in any player line.
// 3. Calculates overall batting average for players without any error in their input data

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
)

// Entry point of the program // main function
func main() {
	color.Cyan(`Welcome to the player statistics calculator test program. I am going to
read players from an input data file.  You will tell me the name of
your input file.  I will store all of the players in a list,
compute each player's averages and then write the resulting team report 
to your output file.`)

	var fileName, outFile, option string
	fmt.Println("Enter the name of your input file")
	fmt.Scan(&fileName)
	players, errors := getPlayers(fileName)
	printPlayerReport(players)
	printErrorReport(errors)
	fmt.Println("Fo you want to write your output to a file? Type Y/N")
	fmt.Scan(&option)
	if option == "Y" || option == "y" {
		fmt.Println("Give a name for your output file")
		fmt.Scan(&outFile)
		// outputf, _ := os.Create(outFile)
		// fmt.Println()
		// writer := bufio.NewWriter(outputf)
		// defer writer.Flush()

		fmt.Println("Output written to", outFile)
		fmt.Println("End of program, GoodBye!")
	} else if option == "N" || option == "n" {
		fmt.Println("End of Program, GoodBye!")
	} else {
		fmt.Println("End of Program, GoodBye!")
	}

	// printPlayerReport(players)
	// printErrorReport(errors)

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
