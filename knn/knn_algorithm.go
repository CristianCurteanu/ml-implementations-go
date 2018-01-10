package knn

import (
	"math"
	"sort"
)

type Point struct {
	x float64
	y float64
}

type Label struct {
	name     string
	location Point
}

type Labels []Label

type Pair struct {
	Key   string
	Value int
}

type PairList []Pair

func (p PairList) Len() int           { return len(p) }
func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p PairList) Less(i, j int) bool { return p[i].Value > p[j].Value }

func Distance(v1, v2 Point) (result float64) {
	s := Point{x: (v1.x - v2.x), y: (v1.y - v2.y)}
	result = math.Sqrt((s.x * s.x) + (s.y * s.y))
	return
}

func Classify(k int, data []Label, pt Point) (result string) {
	counts := map[string]int{}
	var distance float64
	for _, label := range data {
		distance = Distance(pt, label.location)
		if distance < float64(k) {
			if _, ok := counts[label.name]; !ok {
				counts[label.name] = 1
			} else {
				counts[label.name] += 1
			}
		}
	}
	result = HighestRate(counts)
	return
}

func HighestRate(collection map[string]int) (result string) {
	if len(collection) == 0 {
		result = ""
	} else {
		p := make(PairList, len(collection))
		i := 0
		for k, v := range collection {
			p[i] = Pair{k, v}
			i++
		}
		sort.Sort(p)
		result = p[0].Key
	}
	return
}
