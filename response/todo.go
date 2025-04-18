package response

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"
)

type TodoDecoder struct {
}

type Todo struct {
	UserID    int
	ID        int
	Title     string
	Completed bool
}

func (t Todo) Print() {
	fmt.Println(strings.Repeat("*", 100))
	fmt.Println("UserID:", t.UserID)
	fmt.Println("ID:", t.ID)
	fmt.Println("Title:", t.Title)
	fmt.Println("Completed:", t.Completed)
	fmt.Println(strings.Repeat("*", 100))
}

func (d TodoDecoder) ToList(r io.Reader) ([]Todo, error) {
	var todos []Todo
	err := json.NewDecoder(r).Decode(&todos)
	if err != nil {
		return nil, err
	}
	return todos, nil
}

func (d TodoDecoder) ToSingle(r io.Reader) (Todo, error) {
	body, err := io.ReadAll(r)
	if err != nil {
		return Todo{}, err
	}
	var todo Todo
	err = json.Unmarshal(body, &todo)
	if err != nil {
		return todo, err
	}
	return todo, nil
}
