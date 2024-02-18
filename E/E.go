package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"

	"sync"
	"sync/atomic"

	strutil "github.com/adrg/strutil"
	"github.com/adrg/strutil/metrics"
)

type Container struct {
	num atomic.Uint64
}

var wg sync.WaitGroup

func _func(c *Container, result map[int]string, i int, query string, universities *[]string) {
	defer wg.Done()

	maxSimilarity := 0.0
	bestMatch := -1
	for j, university := range *universities {

		similarity := strutil.Similarity(strings.ToLower(query),
			strings.ToLower(university), metrics.NewLevenshtein())

		if similarity > maxSimilarity {
			maxSimilarity = similarity
			bestMatch = j
		}
	}

	result[i] = (*universities)[bestMatch]

	c.num.Add(1) // increment atomic counter
	if c.num.Load()%100 == 0 {
		fmt.Println(c.num.Load())
	}
}

func main() {
	universitiesFile, err :=
		os.Open("/Users/phil/Downloads/task_2/universities.txt")
	if err != nil {
		fmt.Println("Error opening universities file:", err)
		return
	}
	defer universitiesFile.Close()

	queriesFile, err :=
		os.Open("/Users/phil/Downloads/task_2/queries.txt")
	if err != nil {
		fmt.Println("Error opening queries file:", err)
		return
	}
	defer queriesFile.Close()

	universitiesScanner := bufio.NewScanner(universitiesFile)
	queriesScanner := bufio.NewScanner(queriesFile)

	var universities []string
	for universitiesScanner.Scan() {
		universities = append(universities,
			universitiesScanner.Text())
	}

	var queries []string
	for queriesScanner.Scan() {
		queries = append(queries, queriesScanner.Text())
	}

	wg.Add(len(queries))
	var result_map = make(map[int]string)
	c := Container{}
	for i := 0; i < len(queries); i++ {
		_func(&c, result_map, i, queries[i], &universities)
	}
	wg.Wait()

	keys := make([]int, 0, len(result_map))
	for key := range result_map {
		keys = append(keys, key)
	}
	sort.Ints(keys)

	answerFile, err :=
		os.Create("/Users/phil/Downloads/task_2/answerGo.txt")
	if err != nil {
		fmt.Println("Error opening answer file:", err)
		return
	}
	defer answerFile.Close()

	for i := range keys {
		answerFile.WriteString(result_map[i] + "\n")
	}
}
