package shortest

// func Exists[T any](items []T, predicate func(T) bool) bool {
// 	for _, item := range items {
// 		if predicate(item) {
// 			return true
// 		}
// 	}
// 	return false
// }

// func Filter[T any](items []T, predicate func(T) bool) []T {
// 	var output []T
// 	for _, item := range items {
// 		if predicate(item) {
// 			output = append(output, item)
// 		}
// 	}
// 	return output
// }

// func Partition[T any](items []T, predicate func(T) bool) (matched []T, unmatched []T) {
// 	for _, item := range items {
// 		if predicate(item) {
// 			matched = append(matched, item)
// 		} else {
// 			unmatched = append(unmatched, item)
// 		}
// 	}
// 	return matched, unmatched
// }