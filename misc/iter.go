package misc

func Iter(count int) []int {
	arr := make([]int, count)
	i := 0
	for i < count {
		arr[i] = i
		i++
	}
	return arr
}