package main

import "fmt"

const NEWSPAPER = "Newspaper"

type Newspaper struct {
	Publication
}

// Define Stringer interface
func (n Newspaper) String() string {
	return fmt.Sprintf("This is newspaper named %s", n.name)
}

func createNewspaper(name, publisher string, pages int) (iPublication, error) {
	if pages >= MIN_PAGE_LENGTH {
		return nil, fmt.Errorf("%s '%s' should not have more than %d pages.", NEWSPAPER, name, MIN_PAGE_LENGTH)
	}
	return &Newspaper{
		Publication: Publication{name, publisher, pages},
	}, nil
}
