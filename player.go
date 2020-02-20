package main

// add preamble in the code
// add your name, A#, seciton, details to follow in the progran
import (
	"fmt"
	"strconv"
)

// Player structure
type Player struct {
	firstName, lastName                                                              string
	plateAppearances, atBats, singles, doubles, triples, homeRuns, walks, hitByPitch int
}

// function to get the batting average of a player
func (p Player) getBattingAverage() float64 {
	return float64(p.singles+p.doubles+p.triples+p.homeRuns) / float64(p.atBats)
}

//function to get he slugging of a player
func (p Player) getSluggingPercentage() float64 {
	return float64(p.singles+2*p.doubles+3*p.triples+4*p.homeRuns) / float64(p.atBats)
}

//function to get Onbase percentage
func (p Player) getOnBasePercentage() float64 {
	return float64(p.singles+p.doubles+p.triples+p.hitByPitch+p.walks) / float64(p.plateAppearances)
}

//get the details of the player and create an instance of each Player Structure
func getPlayer(s []string, lineCount int) (Player, error) {
	err := checkErrors(s, lineCount)
	var player Player
	if err == nil {
		player = Player{
			firstName: s[0],
			lastName:  s[1],
		}
		player.plateAppearances, _ = strconv.Atoi(s[2])
		player.atBats, _ = strconv.Atoi(s[3])
		player.singles, _ = strconv.Atoi(s[4])
		player.doubles, _ = strconv.Atoi(s[5])
		player.triples, _ = strconv.Atoi(s[6])
		player.homeRuns, _ = strconv.Atoi(s[7])
		player.walks, _ = strconv.Atoi(s[8])
		player.hitByPitch, _ = strconv.Atoi(s[9])
	}
	return player, err
}

//check for errors in the line
func checkErrors(s []string, lineCount int) error {
	if len(s) < 10 {
		return fmt.Errorf("line %d: Line contains not enough data", lineCount)
	}
	for i := 2; i < 10; i++ {
		_, err := strconv.Atoi(s[i])
		if err != nil {
			return fmt.Errorf("line %d: Line contains invalid numeric data", lineCount)
		}
	}
	return nil
}
