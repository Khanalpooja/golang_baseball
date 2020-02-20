package main

import (
	"fmt"
	"os"
	"text/tabwriter"
)

// print the error report
func printErrorReport(errors []error) {
	fmt.Printf("\n\n----- %d ERROR LINES FOUND IN INPUT DATA ----\n\n", len(errors))
	for _, err := range errors {
		fmt.Println(err)
	}
}

// print the player report
func printPlayerReport(players []Player) {
	var overalAvg, sum float64
	sum = 0.0
	fmt.Printf("BASEBALL TEAM REPORT --- %d PLAYERS FOUND IN FILE\n", len(players))

	for _, player := range players {
		sum = sum + player.getBattingAverage()
	}

	overalAvg = sum / float64(len(players))
	fmt.Printf("OVERALL BATTING AVERAGE is %.3f\n", overalAvg)

	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 10, 10, 4, '\t', 0)
	defer w.Flush()
	fmt.Fprintf(w, "\n      %-16s %-5s %8s %10s %10s", "PLAYER NAME", ":", "AVERAGE", "SLUGGING", "ONBASE%")
	fmt.Fprintf(w, "\n %s", "-------------------------------------------------------------------")
	for _, player := range players {
		fmt.Fprintf(w, "\n %21s %-7s %-10.3f %-10.3f %.3f", player.lastName+","+player.firstName, ":", player.getBattingAverage(), player.getSluggingPercentage(), player.getOnBasePercentage())
	}
}
