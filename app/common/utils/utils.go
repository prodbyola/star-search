package utils

import (
	"fmt"
	"strconv"
	"strings"
)

type Comparable[T any] interface {
	GreaterThan(RHS *T) bool
}

func sortPartition[T Comparable[T]](arr []T, low, high int, desc bool) ([]T, int) {
	pivot := arr[high]
	i := low

	for j := low; j < high; j++ {
		selected := arr[j]
		dir := !selected.GreaterThan(&pivot)
		if desc {
			dir = !dir
		}

		if dir {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}

	arr[i], arr[high] = arr[high], arr[i]

	return arr, i
}

func Sort[T Comparable[T]](arr []T, low, high int, desc bool) []T {
	if low < high {
		var p int
		arr, p = sortPartition(arr, low, high, desc)
		arr = Sort(arr, low, p-1, desc)
		arr = Sort(arr, p+1, high, desc)
	}

	return arr
}

func CmToFt(cm int64) string {
	constant := 2.54
	inches := float64(cm) / constant
	feet := inches / 12

	ft_str := strconv.FormatFloat(feet, 'f', 8, 64)
	ft_spl := strings.Split(ft_str, ".")
	ft_str = ft_spl[0]

	in_str := "0." + ft_spl[1]
	feet, _ = strconv.ParseFloat(in_str, 64)
	inches = feet * 12
	in_str = strconv.FormatFloat(inches, 'f', 2, 64)

	return fmt.Sprintf("%sft %sin", ft_str, in_str)
}
