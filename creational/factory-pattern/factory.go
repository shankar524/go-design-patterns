package main

import "fmt"

func newPublication(pubType, name, publisher string, pages int) (pub iPublication, err error) {
	// TODO: Create right kind of publicationbased on given type
	switch pubType {
	case MAGAZINE:
		pub, err = createMagazine(name, publisher, pages)
	case NEWSPAPER:
		pub, err = createNewspaper(name, publisher, pages)
	default:
		err = fmt.Errorf("No such publication type: '%s'", pubType)
	}
	return
}
