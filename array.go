package go_basic_extension

type Array []any

func (a *Array) Contains(v any) bool {
	for _, r := range *a {
		if r == v {
			return true
		}
	}

	return false
}

func (a *Array) NotContains(v any) bool {
	return !a.Contains(v)
}

func Intersection[T comparable](f, s []T) []T {
	result := make([]T, 0, len(f))
	hm := make(map[T]interface{}, len(f))
	for _, fV := range f {
		hm[fV] = nil
	}

	for _, sV := range s {
		if _, e := hm[sV]; e {
			result = append(result, sV)
		}
	}

	return result
}

func NonIntersection[T comparable](f, s []T) []T {
	result := make([]T, 0, len(f))
	hm := make(map[T]interface{}, len(f))
	for _, sV := range s {
		hm[sV] = nil
	}

	for _, fV := range f {
		if _, e := hm[fV]; !e {
			result = append(result, fV)
		}
	}

	return result
}
