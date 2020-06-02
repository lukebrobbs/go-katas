package splitandadd

func splitAndAdd(numbers []int, n int) []int {
	if len(numbers) == 0 || n == 0 {
		return numbers
	}

	ret := []int{}
	m := len(numbers) / 2
	f, s := numbers[0:m], numbers[m:]

	if st := len(s) - len(f); st == 1 {
		ret = append(ret, s[0])
		s = numbers[m+1:]
	}

	for i := 0; i < len(s); i++ {
		t := f[i] + s[i]
		ret = append(ret, t)
	}

	return splitAndAdd(ret, n-1)

}
