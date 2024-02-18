package main

import (
	"fmt"
	"os"
	"sort"
	"strings"

	"log"

	strutil "github.com/adrg/strutil"
	"github.com/adrg/strutil/metrics"
)

func main() {
	dictPath := "/Users/phil/Downloads/task_1/dict.txt"
	queriesPath := "/Users/phil/Downloads/task_1/queries.txt"
	resultPath := "/Users/phil/Downloads/task_1/result.txt"

	dictSet := readDict(dictPath)
	dictList := makeList(dictSet)
	dictLen := len(dictSet)

	queriesList, err := readQueries(queriesPath)
	if err != nil {
		log.Fatal(err)
	}

	var result_map = make(map[int]string)

	for i, query := range queriesList {
		if _, found := dictSet[query]; found {
			result_map[i] = query + " 0\n"
		} else {
			myfunc(&result_map, i, query, &dictList, dictLen)
		}

		if i%100 == 0 {
			fmt.Println(i)
		}
	}

	err = writeResults(result_map, resultPath)
	if err != nil {
		log.Fatal(err)
	}
}

func myfunc(result_map *map[int]string, i int, query string, dictList *[]string, dictLen int) {
	var similarTuple struct {
		index int
		ratio float64
	}

	for j := 0; j < dictLen; j++ {
		similarity := strutil.Similarity(strings.ToLower(query), strings.ToLower((*dictList)[j]), metrics.NewLevenshtein())
		if similarTuple.ratio < similarity {
			similarTuple.index = j
			similarTuple.ratio = similarity
		}
	}

	word := (*dictList)[similarTuple.index]
	wordLst := strings.Split(word, " ")
	queryLst := strings.Split(query, " ")

	errorCount := 0
	var words string
	if len(wordLst) == len(queryLst) {
		for i := 0; i < len(wordLst); i++ {
			if queryLst[i] != wordLst[i] {
				queryLst[i] = wordLst[i]
				errorCount++
				words += strings.Join(queryLst, " ")
			}
		}

		if errorCount >= 3 {
			(*result_map)[i] = fmt.Sprintf("%s %d+\n", query,
				errorCount)
		} else {
			(*result_map)[i] = fmt.Sprintf("%s %d %s\n",
				query, errorCount, words)
		}

	} else {
		(*result_map)[i] = "\n"
	}
}

// ---------------------------------------------------------

func readDict(path string) map[string]struct{} {
	dictSet := make(map[string]struct{})
	contents, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(contents), "\n")
	for _, line := range lines {
		dictSet[strings.TrimSpace(line)] = struct{}{}
	}

	return dictSet
}

func makeList(dictSet map[string]struct{}) []string {
	dictList := make([]string, 0)
	for word := range dictSet {
		dictList = append(dictList, word)
	}

	return dictList
}

func readQueries(path string) ([]string, error) {
	contents, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(contents), "\n")
	queriesList := make([]string, 0)
	for _, line := range lines {
		queriesList = append(queriesList, strings.TrimSpace(line))
	}

	return queriesList, nil
}

func writeResults(result_map map[int]string, path string) error {
	keys := make([]int, 0, len(result_map))
	for key := range result_map {
		keys = append(keys, key)
	}
	sort.Ints(keys)

	resultFile, err := os.Create(path)
	if err != nil {
		return err
	}
	defer resultFile.Close()

	for i := range keys {
		resultFile.WriteString(result_map[i])
	}
	return nil
}
