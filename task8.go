package main

import (
	"fmt"
	"strings"
	"sync"
)

type letter struct {
	hitung map[string]int
	mute   sync.Mutex
}

func (s *letter) letFreq(jumlahHuruf chan int, huruf string, group *sync.WaitGroup) {
	// startmute
	s.mute.Lock()

	s.hitung[huruf]++
	jumlahHuruf <- s.hitung[huruf]
	fmt.Println(huruf, " : ", <-jumlahHuruf)
	group.Done()

	// endmute
	s.mute.Unlock()
}

func main() {
	// vardeclare
	var sync sync.WaitGroup
	var mapping = letter{hitung: make(map[string]int)}

	word := "lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua"
	resString := strings.ReplaceAll(word, " ", "")
	count := make(chan int, len(resString))
	sync.Add(len(resString))

	for _, value := range resString {
		go mapping.letFreq(count, string(value), &sync)
	}
	sync.Wait()

	delete(mapping.hitung, ",")
	fmt.Print("result: ", mapping.hitung)
}
