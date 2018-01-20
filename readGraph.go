package search

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Graph interface {
	V() int
	E() int
	AddEdge(v int, w int, weight int)
	HasEdge(v int, w int) bool
	Iterator(v int) Iterator
}

func ReadGraph(g Graph, name string) error {
	file, err := os.Open(name)
	if err != nil {
		return err
	}
	defer file.Close()
	rd := bufio.NewReader(file)
	line, err := rd.ReadString('\n')
	if err != nil {
		return err
	}
	datas := strings.Fields(line)
	e, _ := strconv.Atoi(datas[1])

	for i := 0; i < e; i++ {
		line, err = rd.ReadString('\n')
		datas := strings.Fields(line)
		if err != nil {
			return err
		}
		v, _ := strconv.Atoi(datas[0])
		w, _ := strconv.Atoi(datas[1])
		weight, _ := strconv.Atoi(datas[2])
		g.AddEdge(v, w, weight)
	}
	return nil
}
