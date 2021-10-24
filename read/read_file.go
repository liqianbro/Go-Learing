package read

import (
	"encoding/json"
	"os"
)

const dataFile = "file/data.json"

type Feed struct {
	Name string `json:"site"`
	URI  string `json:"link"`
	Type string `json:"type"`
}

func (f *Feed) ReadFeedByFile() ([]*Feed, error) {
	file, err := os.Open(dataFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	var feedList []*Feed
	err = json.NewDecoder(file).Decode(&feedList)
	return feedList, err
}
