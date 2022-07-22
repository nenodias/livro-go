package main

import (
	"fmt"
	"os"
	"sort"
	"text/tabwriter"
	"time"
)

type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

var tracks = []*Track{
	{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
	{"Go", "Moby", "Moby", 1992, length("4m24s")},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
}

func length(value string) time.Duration {
	t, err := time.ParseDuration(value)
	if err != nil {
		panic(err)
	}
	return t
}

func printTracks(tracks []*Track) {
	const format = "%v\t%v\t%v\t%v\t%v\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length")
	fmt.Fprintf(tw, format, "-----", "------", "-----", "----", "------")
	for _, t := range tracks {
		fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year, t.Length)
	}
	tw.Flush()
}

type byArtist []*Track

func (x byArtist) Len() int {
	return len(x)
}
func (x byArtist) Less(i, j int) bool {
	return x[i].Artist < x[j].Artist
}
func (x byArtist) Swap(i, j int) {
	x[i], x[j] = x[j], x[i]
}

//Generic sort with Generics

type Sortable[T any] struct {
	Values []*T
	less   func(x, y *T) bool
}

func (x Sortable[T]) Len() int {
	return len(x.Values)
}
func (x Sortable[T]) Less(i, j int) bool {
	return x.less(x.Values[i], x.Values[j])
}
func (x Sortable[T]) Swap(i, j int) {
	x.Values[i], x.Values[j] = x.Values[j], x.Values[i]
}

func main() {
	tracksByArtist := byArtist(tracks)
	sort.Sort(tracksByArtist)
	printTracks(tracksByArtist)

	fmt.Println("")

	customSort := Sortable[Track]{tracksByArtist, func(x, y *Track) bool {
		return x.Year < y.Year
	}}
	sort.Sort(sort.Reverse(customSort))
	printTracks(tracksByArtist)
}
