package helper

import "strconv"

type Integer interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

type Number interface {
	Integer | float32 | float64
}

func GreatestCommonDivisor(a, b int64) int64 {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func LeastCommonMultiple(vals ...int64) int64 {
	result := vals[0] * vals[1] / GreatestCommonDivisor(vals[0], vals[1])
	for i := 2; i < len(vals); i++ {
		result = LeastCommonMultiple(result, vals[i])
	}
	return result
}

func Abs[T Number](val T) T {
	if val < 0 {
		return -val
	}
	return val
}

func Sign[T Number](val T) int {
	if val < 0 {
		return -1
	}
	if val > 0 {
		return 1
	}
	return 0
}

func SignBit[T Number](val T) bool {
	return val >= 0
}

func Digits[T Integer](val T) int {
	if val <= 0 {
		panic("Digits(T) not defined for negative and zero values")
	}

	n := T(10)
	for i := 1; ; i++ {
		if val < n {
			return i
		}
		n *= 10
	}
}

func Min[T Ordered](values ...T) T {
	min := values[0]
	for i := 1; i < len(values); i++ {
		if values[i] < min {
			min = values[i]
		}
	}
	return min
}

func Max[T Ordered](values ...T) T {
	max := values[0]
	for i := 1; i < len(values); i++ {
		if values[i] > max {
			max = values[i]
		}
	}
	return max
}

func Mod[T Integer](d, m T) T {
	var res T = d % m
	if res < 0 && m > 0 {
		return res + m
	}
	return res
}

func SumAll[OUT, T Number](values []T) OUT {
	var sum OUT
	for _, val := range values {
		sum += OUT(val)
	}
	return sum
}

func ParseInt[T Integer](str string) T {
	val, err := strconv.ParseInt(str, 10, 64)
	ExitOnError(err, "failed to parse int from %q", str)
	return T(val)
}
