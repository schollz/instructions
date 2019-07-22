package main

import (
	"io/ioutil"
	"os"
	"sort"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	f, err := os.Create("corpus.go")
	check(err)
	defer f.Close()

	f.WriteString("package instructions\n")

	var b []byte
	var pl pairList
	var i int

	// MAKE DIRECTIONS CORPUS
	b, err = ioutil.ReadFile("corpus/directions_pos.txt")
	corpusDirections := strings.Fields(string(b))
	corpusDirectionsMap := make(map[string]struct{})
	for _, c := range corpusDirections {
		corpusDirectionsMap[" "+strings.ToLower(c)+" "] = struct{}{}
	}
	pl = make(pairList, len(corpusDirectionsMap))
	i = 0
	for k := range corpusDirectionsMap {
		pl[i] = pair{k, len(k)}
		i++
	}
	sort.Slice(pl, func(i, j int) bool {
		return pl[i].Key < pl[j].Key
	})
	corpusDirections = make([]string, len(pl))
	for i, p := range pl {
		corpusDirections[i] = p.Key
	}
	f.WriteString(`var corpusDirections = []string{"` + strings.Join(corpusDirections, `"`+",\n"+`"`) + `"}` + "\n")
	f.Sync()

	// MAKE DIRECTIONS NEG CORPUS
	b, err = ioutil.ReadFile("corpus/directions_neg.txt")
	corpusDirections = strings.Fields(string(b))
	corpusDirectionsMap = make(map[string]struct{})
	for _, c := range corpusDirections {
		corpusDirectionsMap[" "+strings.ToLower(c)+" "] = struct{}{}
	}
	pl = make(pairList, len(corpusDirectionsMap))
	i = 0
	for k := range corpusDirectionsMap {
		pl[i] = pair{k, len(k)}
		i++
	}
	sort.Slice(pl, func(i, j int) bool {
		return pl[i].Key < pl[j].Key
	})
	corpusDirections = make([]string, len(pl))
	for i, p := range pl {
		corpusDirections[i] = p.Key
	}
	f.WriteString(`var corpusDirectionsNeg = []string{"` + strings.Join(corpusDirections, `"`+",\n"+`"`) + `"}` + "\n")
	f.Sync()
}

type pair struct {
	Key   string
	Value int
}

type pairList []pair

func (p pairList) Len() int           { return len(p) }
func (p pairList) Less(i, j int) bool { return p[i].Value < p[j].Value }
func (p pairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
