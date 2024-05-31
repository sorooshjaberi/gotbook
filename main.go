package main

import (
	"bufio"
	"fmt"
	"os"
	"soroushjb/entities/notes"
	"strings"
)

func main() {
	title, desc := readFromUser("Give a title:"), readFromUser("Give a description:")

	note := notes.InjectFromFile(title)

	fmt.Println(note)

	note, err := notes.New(title, desc)

	if err != nil {
		panic(err)
	}

	note.Save()
}

func readFromUser(label string) string {

	println(label)

	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}
