package main

import "fmt"

const MAGAZINE = "Magazine"

type Magazine struct {
	Publication
}

// Define Stringer interface
func (m Magazine) String() string {
	return fmt.Sprintf("This is magazine named %s", m.name)
}

func createMagazine(name, publisher string, pages int) (iPublication, error) {
	if pages < MIN_PAGE_LENGTH {
		return nil, fmt.Errorf("%s '%s' should have at least %d pages.", MAGAZINE, name, MIN_PAGE_LENGTH)
	}
	return &Magazine{
		Publication: Publication{name, publisher, pages},
	}, nil
}
