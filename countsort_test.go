package countsort

import (
	//"fmt"
	"sort"
	"testing"
)

func TestCountSort(t *testing.T) {
	for _, slc := range cases {
		if countSortResult := CountSort(slc); len(countSortResult) != len(slc) {
			t.Fatalf("CountSort output length: Expected: %v\n Got: %v\n", len(slc), len(countSortResult))
		}
	}
	for _, slc := range cases {
		//fmt.Printf("Init slice: %v\n", slc)
		countSortResult := CountSort(slc)
		//fmt.Printf("CountSort result: %v\n", countSortResult)
		sort.Ints(slc)
		//fmt.Printf("Std sort result: %v\n", slc)
		for i, val := range countSortResult {
			if val != slc[i] {
				t.Fatalf("CountSort output: Expected: %v\n Got: %v\n", slc, countSortResult)
			}
		}
	}

}

func BenchmarkStdSort(b *testing.B) {
	if testing.Short() {
		b.Skip("Skipping bechmark in short mode.")
	}
	sort.Ints(bigSlc)
}

func BenchmarkCountSort(b *testing.B) {
	if testing.Short() {
		b.Skip("Skipping bechmark in short mode.")
	}
	CountSort(bigSlc)
}
