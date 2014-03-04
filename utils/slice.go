package utils

// AppendStr appends string to slice with no duplicates.
func AppendStr(strs []string, str string) []string {
	for _, s := range strs {
		if s == str {
			return strs
		}
	}
	return append(strs, str)
}

// CompareSliceStr compares two 'string' type slices.
// It returns true if elements and order are both the same.
func CompareSliceStr(s1, s2 []string) bool {
	if len(s1) != len(s2) {
		return false
	}

	for i := range s1 {
		if s1[i] != s2[i] {
			return false
		}
	}

	return true
}

// CompareSliceStr compares two 'string' type slices.
// It returns true if elements are the same, and ignores the order.
func CompareSliceStrU(s1, s2 []string) bool {
	if len(s1) != len(s2) {
		return false
	}

	for i := range s1 {
		for j := len(s2) - 1; j >= 0; j-- {
			if s1[i] == s2[j] {
				s2 = append(s2[:j], s2[j+1:]...)
				break
			}
		}
	}
	if len(s2) > 0 {
		return false
	}
	return true
}
