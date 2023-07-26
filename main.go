package main

import (
	"fmt"
	"github.com/nathan-fiscaletti/consolesize-go"
	"math/rand"
	"os"
	"os/user"
	"strings"
	"time"
)

func print(sentence string) {
	cols, _ := consolesize.GetConsoleSize()
	longestLine := 0
	lines := strings.Split(sentence, "\n")
	for _, line := range lines {
		if len(line) > longestLine {
			longestLine = len(line)
		}
	}
	longestLine += 4
	if longestLine > cols {
		longestLine = cols
	}
	for i := 0; i < longestLine; i++ {
		fmt.Print("#")
	}
	fmt.Print("\n")
	for _, line := range lines {
		fmt.Print("-- ")
		fmt.Println(line)
	}
	for i := 0; i < longestLine; i++ {
		fmt.Print("#")
	}
	fmt.Print("\n")
}

func main() {
	rand.Seed(time.Now().UnixMilli())
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	data, err := os.ReadFile(user.HomeDir + "/.config/motd/motd.txt")
	if err != nil {
		panic(err)
	}
	content := string(data)
	parts := strings.Split(content, "---")
	sentenceID := rand.Intn(len(parts))
	print(parts[sentenceID])
}
