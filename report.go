// Prepared by: Pooja Khanal, A25288740
// Prepared for : Assignment 1, CS524, Spring 2020
// Submitted on : 02/20/2020

//Program Description:
// Main tasks covered by this program are
// 1. This program calculates the player statistics, batting avaerage, slugging and base percentage.
// 2. Checks if inconsisent data are present in any player line.
// 3. Calculates overall batting average for players without any error in their input data
// 4. main.go is the starting point for the program

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// print the error report
func printErrorReport(errors []error, stream io.Writer) {
	fmt.Fprintf(stream, "\n\n----- %d ERROR LINES FOUND IN INPUT DATA ----\n\n", len(errors))
	for _, err := range errors {
		fmt.Fprintf(stream, "%s \n", err)
	}
}

// print the player report
func printPlayerReport(players []Player, stream io.Writer) {
	var overalAvg, sum float64

	sum = 0.0
	fmt.Fprintf(stream, "\nBASEBALL TEAM REPORT --- %d PLAYERS FOUND IN FILE\n", len(players))

	for _, player := range players {
		sum = sum + player.getBattingAverage()
	}

	overalAvg = sum / float64(len(players))
	fmt.Fprintf(stream, "OVERALL BATTING AVERAGE is %.3f\n", overalAvg)

	fmt.Fprintf(stream, "\n      %-16s %-5s %8s %10s %10s", "PLAYER NAME", ":", "AVERAGE", "SLUGGING", "ONBASE%")

	fmt.Fprintf(stream, "\n %s", "-------------------------------------------------------------------")
	for _, player := range players {
		fmt.Fprintf(stream, "\n %21s %-7s %-10.3f %-10.3f %.3f", player.lastName+","+player.firstName, ":", player.getBattingAverage(), player.getSluggingPercentage(), player.getOnBasePercentage())
	}

}

//Output stream handler to write to file
func handlePrintToFile(players []Player, errors []error) {
	var option string
	fmt.Printf("\nDo you want to write your output to a file? Type Y/N\t")
	fmt.Scan(&option)
	if option == "Y" || option == "y" {
		var outFile string
		fmt.Println("Give a name for your output file and press ENTER")
		fmt.Scan(&outFile)
		outputf, _ := os.Create(outFile)
		os.Chmod(outFile, 0777)
		fileStream := bufio.NewWriter(outputf)

		printPlayerReport(players, fileStream)
		printErrorReport(errors, fileStream)
		fileStream.Flush()
		fmt.Println("Report Written to file", outFile)

	}
}
