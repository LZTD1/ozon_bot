package course

import "fmt"

type Course struct {
	ID          uint64
	Title       string
	Description string
}

func (c Course) String() string {
	return fmt.Sprintf("Course(ID: %d, Title: %s, Description: %s)", c.ID, c.Title, c.Description)
}
