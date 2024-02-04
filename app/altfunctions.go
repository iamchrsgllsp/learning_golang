package app

import (
	"fmt"
)

type Movies struct {
	name     string
	director string
	year     int
	score    int
}

func CreateMovie(mname string, mdirector string, myear int, mscore int) {
	m := Movies{name: mname, director: mdirector, year: myear, score: mscore}
	fmt.Println(m)

}

// when building a func out if it is return you need to specify type
func HiThere() string {
	var hello string
	hello = "hello there"
	fmt.Println(hello)
	return hello

}

func ShowMovies() any {
	movies := make(map[string]int)
	fmt.Println(movies)
	return nil
}
