package shortest

import "reflect"

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
			unvisited := wp.GetUnvisited(destinations[wp.Location])
			finished, possible := Partition(unvisited, func(wp Waypoint) bool {
				return wp.Location == finish
			})
			shortest = getCurrentShortest(shortest, finished)
			possible = Filter(possible, func(wp Waypoint) bool {
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