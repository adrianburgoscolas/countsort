//countsort package provide CountSort function to sort a slice of integers.
package countsort

//CountSort return a sorted slice of integers.
func CountSort(slc []int) []int {
	var rang int
	for _, val := range slc {
		if rang < val {
			rang = val
		}
	}
	freqSlc := make([]int, rang+1)
	for _, val := range slc {
		freqSlc[val]++
	}
	offset := 0
	outSlc := make([]int, len(slc))
	for num, amount := range freqSlc {
		if amount == 0 {
			continue
		}
		for i := offset; i < offset+amount; i++ {
			outSlc[i] = num
		}
		offset += amount
	}
	return outSlc
}
