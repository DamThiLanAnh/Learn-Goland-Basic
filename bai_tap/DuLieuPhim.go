package main

import "fmt"

type Displayable interface {
	Display()
}
type MovieData struct {
	Title        string
	Director     string
	YearReleased int
	RunningTime  int
}

func (c MovieData) Display() {
	fmt.Println("Title:", c.Title)
	fmt.Println("Director:", c.Director)
	fmt.Println("YearReleased:", c.YearReleased)
	fmt.Println("RunningTime:", c.RunningTime)
	fmt.Println()
}
func main() {
	movie1 := MovieData{
		Title:        "Movie One",
		Director:     "Luna",
		YearReleased: 2021,
		RunningTime:  120,
	}
	movie2 := MovieData{
		Title:        "Movie Two",
		Director:     "Alex",
		YearReleased: 2022,
		RunningTime:  90,
	}
	var dis Displayable

	dis = movie1
	dis.Display()

	dis = movie2
	dis.Display()
}
