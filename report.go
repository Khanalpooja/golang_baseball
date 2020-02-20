package main

import (
	"fmt"
	"io"
	"log"
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

//write the player report to the file
func printReportToFile(players []Player) {
	var overalAvg, sum float64

	sum = 0.0
	var outFile string
	var option string
	fmt.Printf("\nDo you want to write your output to a file? Type Y/N")
	fmt.Scan(&option)
	if option == "Y" || option == "y" {
		fmt.Println("Give a name for your output file")
		fmt.Scan(&outFile)
		outputf, _ := os.Create(outFile)
		os.Chmod(outFile, 0777)
		fw := io.Writer(outputf)

		log.SetOutput(fw)
		sum = 0.0
		fmt.Printf("BASEBALL TEAM REPORT --- %d PLAYERS FOUND IN FILE\n", len(players))

		for _, player := range players {
			sum = sum + player.getBattingAverage()
		}

		overalAvg = sum / float64(len(players))
		fmt.Printf("OVERALL BATTING AVERAGE is %.3f\n", overalAvg)
		fmt.Fprintf(fw, "\n      %-16s %-5s %8s %10s %10s", "PLAYER NAME", ":", "AVERAGE", "SLUGGING", "ONBASE%")

		fmt.Fprintf(fw, "\n %s", "-------------------------------------------------------------------")
		for _, player := range players {
			fmt.Fprintf(fw, "\n %21s %-7s %-10.3f %-10.3f %.3f", player.lastName+","+player.firstName, ":", player.getBattingAverage(), player.getSluggingPercentage(), player.getOnBasePercentage())
		}
		fmt.Println("Output written to", outFile)
		fmt.Println("End of program, GoodBye!")
	} else if option == "N" || option == "n" {
		fmt.Println("End of Program, GoodBye!")
	} else {
		fmt.Println("End of Program, GoodBye!")
	}

}
