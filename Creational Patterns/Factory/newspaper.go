package main

import "fmt"

type newspaper struct {
	publication
}

func (m newspaper) String() string {
	return fmt.Sprintf("This is magazine name %s", m.name)
}
func createNewspaper(name string, pages int, publisher string) iPublication {
	return &newspaper{
		publication: publication{
			name:      name,
			pages:     pages,
			publisher: publisher,
		},
	}
}
