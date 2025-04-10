package todo

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type item struct {
	Task        string
	Done        bool
	CreatedAt   time.Time
	CompletedAt time.Time
}

type List []item

func (l *List) Add(task string) {
	t := item{
		Task:        task,
		Done:        false,
		CreatedAt:   time.Now(),
		CompletedAt: time.Time{},
	}
	*l = append(*l, t)

}

func (l *List) Complete(i int) error {
	ls := *l
	if i <= 0 || i > len(ls) {
		return fmt.Errorf("Item %d does not exist", i)
	}

	ls[i-1].Done = true
	ls[i-1].CompletedAt = time.Now()

	return nil
}

func (l *List) Delete(i int) error {

	ls := *l

	if i <= 0 || i > len(ls) {
		return fmt.Errorf("Item %d does not exist", i)
	}

	*l = append(ls[:i-1], ls[i:]...)

	return nil
}

func (l *List) Save(filename string) error {
	b, err := json.Marshal(l)
	if err != nil {
		return err
	}

	if err = os.WriteFile(filename, b, 0600); err != nil {
		return err
	}

	return nil
}

func (l *List) Get(filename string) error {
	r, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	if len(r) == 0 {
		return fmt.Errorf("file %s is empty", filename)
	}

	if err = json.Unmarshal(r, l); err != nil {
		return err
	}

	return nil
}
