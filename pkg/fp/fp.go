package fp

// Split 集合分割
func Split[T any](predicate func(x T) bool) func(xs []T) ([]T, []T) {
	return func(xs []T) ([]T, []T) {
		left := make([]T, 0)
		right := make([]T, 0)
		if len(xs) == 0 {
			return left, right
		}

		for _, x := range xs {
			if predicate(x) {
				left = append(left, x)
			} else {
				right = append(right, x)
			}
		}
		return left, right
	}

}

// AnyMatch 匹配
func AnyMatch[T any](predicate func(x T) bool) func(xs []T) bool {
	return func(xs []T) bool {
		for _, x := range xs {
			if predicate(x) {
				return true
			}
		}
		return false
	}
}

// Filter 过滤
func Filter[T any](predicate func(x T) bool) func(xs []T) []T {
	return func(xs []T) []T {
		result := make([]T, 0)
		if len(xs) == 0 {
			return result
		}
		for _, x := range xs {
			if predicate(x) {
				result = append(result, x)
			}
		}
		return result
	}
}

// FindFirst 找到首个满足条件
func FindFirst[T any](predicate func(x T) bool) func(xs []T) (T, bool) {
	return func(xs []T) (T, bool) {
		for _, x := range xs {
			if predicate(x) {
				return x, true
			}
		}
		var x T
		return x, false
	}
}

// Map 集合映射
func Map[T, U any](mapTo func(x T) U) func(xs []T) []U {
	return func(ts []T) []U {
		us := make([]U, 0)
		for _, x := range ts {
			us = append(us, mapTo(x))
		}
		return us
	}
}

// GroupBy 分组
func GroupBy[T any, K comparable](keyFunc func(x T) K) func(xs []T) map[K][]T {
	return func(xs []T) map[K][]T {
		group := make(map[K][]T, 0)
		for _, x := range xs {
			group[keyFunc(x)] = append(group[keyFunc(x)], x)
		}
		return group
	}
}
