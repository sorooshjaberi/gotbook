package notes

import (
	"encoding/json"
	"errors"
	"os"
	"strings"
	"time"
)

var NoteSaveFileName = "notes.json"

type Note struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	CreatedAt   int    `json:"createdAt"`
}

func New(title, description string) (*Note, error) {
	if title == "" || description == "" {
		return nil, errors.New("no empty values allowed")
	}

	createdAt := int(time.Now().UnixNano() / int64(time.Millisecond))

	note := Note{
		Title:       title,
		Description: description,
		CreatedAt:   createdAt,
	}

	return &note, nil

}

func InjectFromFile(fileName string) *Note {

	fileName = convertTitle2Filename(fileName)

	if !strings.HasSuffix(fileName, ".json") {
		fileName = fileName + ".json"
	}

	fileData, err := os.ReadFile(fileName)

	note := Note{}

	if err != nil {
		println("something went wrong reading the file")
		return &note
	}

	json.Unmarshal(fileData, &note)

	return &note

}

func (note *Note) Save() {

	noteJsonBytes, err := json.Marshal(note)

	if err != nil {
		println("something went wrong writing to file")
		return
	}
	fileName := convertTitle2Filename(note.Title)

	os.WriteFile(fileName, (noteJsonBytes), 0644)
}

func convertTitle2Filename(title string) string {
	return strings.ToLower(strings.ReplaceAll(title, " ", "_")) + ".json"
}
