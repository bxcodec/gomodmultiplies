package gomodmultiplies

func Multiply(items ...int64) int64 {
	res := int64(1)
	for _, item := range items {
		res *= item
	}
	return res
}
