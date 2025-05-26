package main

func CheckEquality[T comparable](a1, a2 []T) bool {
	return CheckEqualityWithFunc(a1, a2, func(s1, s2 T) bool { return s1 == s2 })
}

func CheckEqualityWithFunc[T any](a1, a2 []T, isEqual func(_, _ T) bool) bool {
	if len(a1) != len(a2) {
		return false
	}
	checked := []int{}
l1:
	for _, s1 := range a1 {
	l2:
		for i, s2 := range a2 {
			if isEqual(s1, s2) {
				for _, num := range checked {
					if i == num {
						continue l2
					}
				}
				checked = append(checked, i)
				continue l1
			}
		}
		return false
	}
	return true
}
