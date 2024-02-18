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

var wg sync.WaitGroup

func main() {
	universitiesFile, err := os.Open("/Users/phil/Downloads/task_2/universities.txt")
	if err != nil {
		fmt.Println("Error opening universities file:", err)
		return
	}
	defer universitiesFile.Close()

	queriesFile, err := os.Open("/Users/phil/Downloads/task_2/queries.txt")
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
	var counter atomic.Uint64
	for i := 0; i < len(queries); i++ {
		_func(&counter, result_map, i, queries[i], &universities)
	}
	wg.Wait()

	keys := make([]int, 0, len(result_map))
	for key := range result_map {
		keys = append(keys, key)
	}
	sort.Ints(keys)

	answerFile, err := os.Create("/Users/phil/Downloads/task_2/answerGo.txt")
	if err != nil {
		fmt.Println("Error opening answer file:", err)
		return
	}
	defer answerFile.Close()

	for i := range keys {
		answerFile.WriteString(result_map[i] + "\n")
	}
}

func _func(counter *atomic.Uint64, result map[int]string, i int, query string, universities *[]string) {
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

	counter.Add(1) // increment atomic counter
	if counter.Load()%100 == 0 {
		fmt.Println(counter.Load())
	}
}
