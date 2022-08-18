package prototype

import (
	"encoding/json"
	"time"
)

type Keyword struct {
	word     string
	visit    int
	UpdateAt *time.Time
}

func (K *Keyword) Clone() *Keyword {
	var newKeyword Keyword
	b, _ := json.Marshal(K)
	json.Unmarshal(b, &newKeyword)
	return &newKeyword
}

type Keywords map[string]*Keyword

func (words Keywords) Clone(updateWords []*Keyword) Keywords {
	newKeywords := Keywords{}

	for k, v := range words {
		newKeywords[k] = v
	}

	for _, word := range updateWords {
		newKeywords[word.word] = word.Clone()
	}

	return newKeywords
}
