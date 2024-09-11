package utils

func MakeUnique[T int64 | int](arr []T) (resp []T) {
	unique := make(map[T]struct{})
	for _, v := range arr {
		if _, ok := unique[v]; !ok {
			unique[v] = struct{}{}
			resp = append(resp, v)
		}
	}
	return
}
