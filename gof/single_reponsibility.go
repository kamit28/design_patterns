package gof

import (
	"fmt"
)

var entryCount = 0

type Journal struct {
	entries []string
}

func (j *Journal) AddEntry(text string) int {
	entryCount++
	entry := fmt.Sprintf("%d: %s",
		entryCount, text)
	j.entries = append(j.entries, entry)
	return entryCount
}

func (j *Journal) RemoveEntry(index int) (int, error) {
	if index < 0 || index > len(j.entries)-1 {
		return entryCount, fmt.Errorf("index out of bounds: %d", index)
	}

	j.entries[index] = j.entries[len(j.entries)-1]
	j.entries = j.entries[:len(j.entries)-1]
	entryCount--
	return entryCount, nil
}
