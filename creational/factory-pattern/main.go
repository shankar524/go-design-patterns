package main

import "fmt"

func main() {
	mag1, _ := newPublication(MAGAZINE, "Time", "Andrew Garfield", 50)
	mag2, _ := newPublication(MAGAZINE, "Life", "Marget Thacher", 90)
	news1, _ := newPublication(NEWSPAPER, "The Herald", "Mark Noble", 6)
	news2, _ := newPublication(NEWSPAPER, "The Standard", "Anim Shah", 9)

	printPublicationDetails(mag1)
	printPublicationDetails(mag2)
	printPublicationDetails(news1)
	printPublicationDetails(news2)

	// factory returning error
	_, err := newPublication("Book", "Science", "Mark Twin", 205)
	fmt.Println(err)

	_, err = newPublication(NEWSPAPER, "Element", "Anul Shakya", MIN_PAGE_LENGTH)
	fmt.Println(err)

	_, err = newPublication(MAGAZINE, "Vertex", "Alex Noman", MIN_PAGE_LENGTH-1)
	fmt.Println(err)
}

func printPublicationDetails(pub iPublication) {
	fmt.Println("-----------------------------")
	fmt.Printf("%s\n", pub)
	fmt.Printf("Type: %T\n", pub)
	fmt.Printf("Name: %s\n", pub.GetName())
	fmt.Printf("Pages: %d\n", pub.GetPages())
	fmt.Printf("Publisher: %s\n", pub.GetPublisher())
	fmt.Println("-----------------------------")
}
