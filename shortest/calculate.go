package shortest

import "reflect"

// func filter(items []Waypoint, predicate func(Waypoint) bool) []Waypoint {
// 	var output []Waypoint
// 	for _, item := range items {
// 		if predicate(item) {
// 			output = append(output, item)
// 		}
// 	}
// 	return output
// }

// func filter(items []Waypoint, predicate func(Waypoint) bool) (output []Waypoint) {
// 	for _, item := range items {
// 		if predicate(item) {
// 			output = append(output, item)
// 		}
// 	}
// 	return output
// }

func filter[T any](items []T, predicate func(T) bool) (output []T) {
	for _, item := range items {
		if predicate(item) {
			output = append(output, item)
		}
	}
	return output
}

// func partition(items []Waypoint, predicate func(Waypoint) bool) (matched []Waypoint, unmatched []Waypoint) {
// 	for _, item := range items {
// 		if predicate(item) {
// 			matched = append(matched, item)
// 		} else {
// 			unmatched = append(unmatched, item)
// 		}
// 	}
// 	return matched, unmatched
// }

func partition[T any](items []T, predicate func(T) bool) (matched []T, unmatched []T) {
	for _, item := range items {
		if predicate(item) {
			matched = append(matched, item)
		} else {
			unmatched = append(unmatched, item)
		}
	}
	return matched, unmatched
}

func getCurrentShortest(current Waypoint, candidates []Waypoint) Waypoint {
	shortest := current
	for _, wp := range candidates {
		if reflect.DeepEqual(shortest, Waypoint{}) || wp.TotalDistance < shortest.TotalDistance {
			shortest = wp
		}
	}
	return shortest
}

func findShortestRoute(start string, finish string, connections []Connection) Waypoint {
	var shortest Waypoint
	destinations := make(map[string][]Connection)
	for _, item := range connections {
		destinations[item.Start] = append(destinations[item.Start], item)
	}
	waypoints := []Waypoint{{Location: start, Route: []string{}, TotalDistance: 0}}
	for {
		if len(waypoints) == 0 {
			return shortest
		}
		var candidates []Waypoint
		for _, wp := range waypoints {
			unvisited := wp.Unvisited(destinations[wp.Location])
			finished, possible := partition(unvisited, func(wp Waypoint) bool {
				return wp.Location == finish
			})
			shortest = getCurrentShortest(shortest, finished)
			possible = filter(possible, func(wp Waypoint) bool {
				return reflect.DeepEqual(shortest, Waypoint{}) || wp.TotalDistance < shortest.TotalDistance
			})
			candidates = append(candidates, possible...)
		}
		waypoints = candidates
	}
}

func GetShortestRoute(filePath string, start string, finish string) (Waypoint, error) {
	connections, err := LoadConnections(filePath)
	if err != nil {
		return Waypoint{}, err
	}
	shortest := findShortestRoute(start, finish, connections)
	return shortest, nil
}
