package sliceup

const size = 100_000

func UpdateWithAppend() []int {
	s := make([]int, size)
	for i := range s {
		s = append(s, i)
	}
	return s
}

func UpdateWithIndex() []int {
	s := make([]int, size)
	for i := range size {
		s[i] = i
	}
	return s
}

func UpdateArrayWithIndex() [size]int {
	var arr [size]int
	for i := range arr {
		arr[i] = i
	}
	return arr
}
