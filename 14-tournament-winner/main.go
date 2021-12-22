package main

import "fmt"

func getWinner(pointsTable map[string]int) string {
	var winner string
	var maxPoints int
	for key, value := range pointsTable {
		if value > maxPoints {
			winner = key
			maxPoints = value
		}
	}
	return winner
}

func TournamentWinner(competitions [][]string, results []int) string {
	var pointsTable = make(map[string]int)

	if len(competitions) == 0 || len(results) == 0 {
		return ""
	}

	for match := 0; match < len(competitions); match++ {
		if results[match] == 0 {
			pointsTable[competitions[match][1]]++
		} else {
			pointsTable[competitions[match][0]]++
		}
	}

	return getWinner(pointsTable)
}

func main() {
	var competitions [][]string
	var results []int

	competitions = append(competitions, []string{"HTML", "C#"})
	competitions = append(competitions, []string{"C#", "Python"})
	competitions = append(competitions, []string{"Python", "HTML"})

	results = append(results, 0)
	results = append(results, 0)
	results = append(results, 1)

	fmt.Println(TournamentWinner(competitions, results))
}
