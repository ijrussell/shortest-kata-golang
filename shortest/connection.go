package shortest

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Connection struct {
	Start    string
	Finish   string
	Distance uint16
}

func LoadConnections(path string) ([]Connection, error) {
	readFile, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}
	defer readFile.Close()
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var connections []Connection
	// skip first row
	fileScanner.Scan()
	for fileScanner.Scan() {
		row := strings.Split(fileScanner.Text(), ",")
		distance, err := strconv.Atoi(row[2])
		if err != nil {
			return nil, err
		}
		items := []Connection{
			{Start: row[0], Finish: row[1], Distance: uint16(distance)},
			{Start: row[1], Finish: row[0], Distance: uint16(distance)},
		}
		connections = append(connections, items...)
	}
	return connections, nil
}
