package main

import (
	"bufio"
	"embed"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

type game struct {
	answer     string
	trialsLeft int
	dictionary []string
}

const (
	RESET_COLOR  = "\033[0m"
	RED_COLOR    = "\033[31m"
	GREEN_COLOR  = "\033[32m"
	YELLOW_COLOR = "\033[33m"
	BLUE_COLOR   = "\033[34m"
	PURPLE_COLOR = "\033[35m"
	CYAN_COLOR   = "\033[36m"
	GRAY_COLOR   = "\033[37m"
	WHITE_COLOR  = "\033[97m"

	TRIALS = 5
)

//go:embed words.txt
var f embed.FS

func main() {
	rand.Seed(time.Now().Unix())
	wordle, err := buildGame()
	if err != nil {
		log.Fatal(err)
	}
	wordle.play()
}

func (wordle *game) play() {
	fmt.Printf("Guess a %d letter word\n", len(wordle.answer))
	for {
		fmt.Printf("------lives left %d------\n", wordle.trialsLeft)
		if wordle.trialsLeft < 1 {
			fmt.Printf("You lose! The word was: %s\n", wordle.answer)
			os.Exit(0)
		}
		var ans string
		ans, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			panic(err)
		}
		ans = strings.TrimSpace(ans)
		if !inDictionary(wordle, ans) {
			fmt.Println("Your word is not in our dictionary!")
			continue
		}
		if ans == wordle.answer {
			fmt.Println("You got it right!")
			break
		} else {
			for i, ch := range ans {
				if strings.Contains(wordle.answer, string(ch)) {
					if wordle.answer[i] == byte(ch) {
						fmt.Print(GREEN_COLOR + string(ch) + RESET_COLOR + "\t")
					} else {
						fmt.Print(YELLOW_COLOR + string(ch) + RESET_COLOR + "\t")
					}
				} else {
					fmt.Print(RED_COLOR + string(ch) + RESET_COLOR + "\t")
				}
			}
		}
		fmt.Println()
		wordle.trialsLeft--
	}
}

func inDictionary(src *game, ans string) bool {
	for _, word := range src.dictionary {
		if word == ans {
			return true
		}
	}
	return false
}

func buildGame() (*game, error) {
	fcontent, err := f.ReadFile("words.txt")
	if err != nil {
		return nil, err
	}
	words := strings.Split(string(fcontent), "\n")
	wordle := game{
		answer:     words[rand.Intn(len(words))],
		trialsLeft: TRIALS,
		dictionary: words,
	}
	return &wordle, nil
}
