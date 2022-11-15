package main

import (
	"flag"
	"fmt"
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
	help := flag.Bool("h", false, "help page")

	flag.Parse()

	if *help || len(os.Args) <= 1 {
		showHelpPage()
		return
	}

	for _, arg := range flag.Args() {

		var diceset models.DiceSet

		requests := parseArgument(arg)

		for _, request := range requests {
			diceset.Import(request)
		}

		if diceset.Empty() {
			continue
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

func showHelpPage() {
	fmt.Printf("\nUsage: roll [amount]d[side](+[amount]d[side])\n")
	fmt.Printf("Example:\n  'roll 1d6' means roll 1 die with 6 face\n")
	fmt.Printf("  'roll 1d6+2d8' means roll 1 die with 6 face and 2 dice with 8 face\n")
	fmt.Printf("Flags:\n  -n: No dice rolling animation\n")
	fmt.Printf("  -h: Help page\n")
	fmt.Println()
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
