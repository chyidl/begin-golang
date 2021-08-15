package series

func GetFibonacciSeries(n int) []int {
	ret := []int{1, 1}
	for i := 2; i < n; i++ {
		ret = append(ret, ret[i-2]+ret[i-1])
	}
	return ret
}

// 小写声明不允许在包外部访问
func Square(n int) int {
	return n * n
}
