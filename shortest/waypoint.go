package shortest

type Waypoint struct {
	Location      string
	Route         []string
	TotalDistance uint16
}

// func exists(items []string, predicate func(string) bool) bool {
// 	for _, item := range items {
// 		if predicate(item) {
// 			return true
// 		}
// 	}
// 	return false
// }

func exists[T any](items []T, predicate func(T) bool) bool {
	for _, item := range items {
		if predicate(item) {
			return true
		}
	}
	return false
}

func (wp Waypoint) Unvisited(possible []Connection) []Waypoint {
	var unvisited []Waypoint
	for _, cn := range possible {
		visited := exists(wp.Route, func(v string) bool {
			return v == cn.Finish
		})
		if !visited {
			newWp := Waypoint{
				Location:      cn.Finish,
				Route:         append(wp.Route, cn.Start),
				TotalDistance: wp.TotalDistance + cn.Distance}
			unvisited = append(unvisited, newWp)
		}
	}
	return unvisited
}
