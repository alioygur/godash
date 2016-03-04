package godash

// Abs returns absolute value of number
func Abs(value float64) float64 {
	return value * Sign(value)
}

// Sign returns signum of number: 1 in case of value > 0, -1 in case of value < 0, 0 otherwise
func Sign(value float64) float64 {
	if value > 0 {
		return 1
	} else if value < 0 {
		return -1
	} else {
		return 0
	}
}
