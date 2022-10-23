package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"
)

// item struct represents a TODO item
type item struct {
	Task        string
	Done        bool
	CreatedAt   time.Time
	CompletedAt time.Time
}

// List represents a list of TODO items
type List []item

// Add method creates a new TODO item and append to the list
func (l *List) Add(task string) {
	t := item{
		Task:        task,
		Done:        false,
		CreatedAt:   time.Now(),
		CompletedAt: time.Time{},
	}
	*l = append(*l, t)
}

// Complete method marks a TODO item as completed by
// setting Done=true and CompletedAt to the current time
func (l *List) Complete(i int) error {
	list := *l
	if i <= 0 || i > len(list) {
		return fmt.Errorf("item %d doesn't exist", i)
	}
	// Adjust index for 0 based index
	list[i-1].Done = true
	list[i-1].CompletedAt = time.Now()
	return nil
}

// Delete method deletes a TODO item from the list
func (l *List) Delete(i int) error {
	list := *l
	if i <= 0 || i > len(list) {
		return fmt.Errorf("item %d doesn't exist", i)
	}
	// Adjusting to 0 based index
	*l = append(list[:i-1], list[i:]...)
	return nil
}

// Save method parse the List into JSON format and saves to file, provided
// the file name.
func (l *List) Save(fileName string) error {
	js, err := json.Marshal(l)
	if err != nil {
		return err
	}
	return os.WriteFile(fileName, js, 06444)
}

// Get method open a file provided the file name, decodes the JSON data, and parse
// it to the list.
func (l *List) Get(fileName string) error {
	file, err := os.ReadFile(fileName)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return err
	}
	if len(file) == 0 {
		return nil
	}
	return json.Unmarshal(file, l)
}
