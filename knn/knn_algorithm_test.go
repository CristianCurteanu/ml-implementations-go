package knn

import (
	"testing"
)

func TestKnnAlgorithm(test *testing.T) {
	datas := []Label{Label{name: "y", location: Point{1.0, 2.0}}}
	datas = append(datas, Label{name: "y", location: Point{3.0, 5.0}})
	datas = append(datas, Label{name: "y", location: Point{4.0, 2.0}})
	datas = append(datas, Label{name: "z", location: Point{1.0, 5.0}})
	datas = append(datas, Label{name: "z", location: Point{4.0, 7.0}})

	label1 := Classify(2, datas, Point{4, 7.2})
	if label1 != "z" {
		test.Fatal("Error")
	}

	label2 := Classify(1, datas, Point{0.75, 1})
	if label2 != "" {
		test.Fatal("Error")
	}
}
