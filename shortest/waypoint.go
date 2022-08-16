package shortest

type Waypoint struct {
	Location      string
	Route         []string
	TotalDistance uint16
}

func (wp Waypoint) GetUnvisited(possible []Connection) []Waypoint {
	var unvisited []Waypoint
	for _, cn := range possible {
		if !Exists(wp.Route, func(v string) bool {
			return v == cn.Finish
		}) {
			newWp := Waypoint{Location: cn.Finish, Route: append(wp.Route, cn.Start), TotalDistance: wp.TotalDistance + cn.Distance}
			unvisited = append(unvisited, newWp)
		}
	}
	return unvisited
}
