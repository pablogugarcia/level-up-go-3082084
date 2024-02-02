package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"os"
	"text/tabwriter"
	"time"
)



const path = "songs.json"

// Song stores all the song related information
type Song struct {
	Name      string `json:"name"`
	Album     string `json:"album"`
	PlayCount int64  `json:"play_count"`
}

func totalSongsInAlbums(albums [][]Song) int {
	var total int

	for _, a := range albums {
		total += len(a)
	}
	return total
}

func removeAlbum(a [][]Song, s int) [][]Song {
    return append(a[:s], a[s+1:]...)
}

func removeSong(song []Song, s int) []Song {
    return append(song[:s], song[s+1:]...)
}

// makePlaylist makes the merged sorted list of songs
func makePlaylist(albums [][]Song) []Song {
	rand.Seed(time.Now().UnixNano())
	var totalSongs = totalSongsInAlbums(albums)
	result := make([]Song, 0)
	for totalSongs > 0 {
		a := rand.Intn(len(albums))
		s := rand.Intn(len(albums[a]))
		result = append(result, albums[a][s])
		if len(albums[a]) == 1 {
			albums = removeAlbum(albums, a)
		}else {
			albums[a] = removeSong(albums[a], s)
		}
		totalSongs--
 	}
	return result
}

func main() {
	albums := importData()
	printTable(makePlaylist(albums))
}

// printTable prints merged playlist as a table
func printTable(songs []Song) {
	w := tabwriter.NewWriter(os.Stdout, 3, 3, 3, ' ', tabwriter.TabIndent)
	fmt.Fprintln(w, "####\tSong\tAlbum\tPlay count")
	for i, s := range songs {
		fmt.Fprintf(w, "[%d]:\t%s\t%s\t%d\n", i+1, s.Name, s.Album, s.PlayCount)
	}
	w.Flush()

}

// importData reads the input data from file and creates the friends map
func importData() [][]Song {
	file, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	var data [][]Song
	err = json.Unmarshal(file, &data)
	if err != nil {
		log.Fatal(err)
	}

	return data
}
