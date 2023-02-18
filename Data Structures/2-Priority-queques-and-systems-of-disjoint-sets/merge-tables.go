package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
)

type DisjointSet struct {
	tables     []int
	records    []int
	maxRecords int
}

func (ds *DisjointSet) MakeTable(number, records int) {
	ds.tables[number] = number
	ds.records[number] = records
	if records > ds.maxRecords {
		ds.maxRecords = records
	}
}

func (ds *DisjointSet) Find(i int) int {
	if i != ds.tables[i] {
		ds.tables[i] = ds.Find(ds.tables[i])
	}
	return ds.tables[i]
}

func (ds *DisjointSet) Union(destination, source int) int {
	destinationId := ds.Find(destination)
	sourceId := ds.Find(source)

	if destinationId != sourceId {
		ds.records[destinationId] += ds.records[sourceId]
		ds.records[sourceId] = 0

		if ds.records[destinationId] > ds.maxRecords {
			ds.maxRecords = ds.records[destinationId]
		}
		ds.tables[sourceId] = destinationId
	}
	return ds.maxRecords
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	m, _ := strconv.Atoi(scanner.Text())
	set := DisjointSet{
		tables:  make([]int, n),
		records: make([]int, n),
	}
	for i := range set.tables {
		scanner.Scan()
		r, _ := strconv.Atoi(scanner.Text())
		set.MakeTable(i, r)
	}

	builder := new(bytes.Buffer)
	for i := 0; i < m; i++ {
		scanner.Scan()
		destination, _ := strconv.Atoi(scanner.Text())
		scanner.Scan()
		source, _ := strconv.Atoi(scanner.Text())
		size := set.Union(destination-1, source-1)
		builder.WriteString(strconv.Itoa(size) + "\n")
	}
	fmt.Print(builder.String())
}
