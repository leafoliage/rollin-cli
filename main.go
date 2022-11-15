package main

import (
	"flag"
	"math/rand"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/leafoliage/rollin-cli/models"
)

func main() {

	rand.Seed(time.Now().UnixNano())

	noAnime := flag.Bool("n", false, "no dice rolling animation")
	flag.Parse()

	for _, arg := range flag.Args() {

		var diceset models.DiceSet

		requests := parseArgument(arg)

		for _, request := range requests {
			diceset.Import(request)
		}

		if *noAnime {
			diceset.Roll()
			diceset.Print(true)
		} else {
			showCursor(false)
			diceset.ScoreAnimation()
			showCursor(true)
		}
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
