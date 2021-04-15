package problems

func IsDivisible(value, divider int) bool {
	if value%divider == 0 {
		return true
	}
	return false
}
