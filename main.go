package main

import (
	"math/rand"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/leafoliage/rollin-cli/models"
)

func main() {

	rand.Seed(time.Now().UnixNano())

	args := os.Args[1:]

	for _, arg := range args {

		var diceset models.DiceSet

		requests := parseArgument(arg)

		for _, request := range requests {
			diceset.Import(request)
		}

		showCursor(false)
		diceset.ScoreAnimation()
		showCursor(true)
	}
}

func parseArgument(arg string) []string {
	request := strings.Split(arg, "+")
	return request
}

func showCursor(show bool) {

	var arg string
	if show {
		arg = "cvvis"
	} else {
		arg = "civis"
	}

	cmd := exec.Command("tput", arg)
	cmd.Stdout = os.Stdout
	cmd.Run()
}
