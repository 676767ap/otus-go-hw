package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

type kv struct {
	Key   string
	Value int
}

func Top10(inptStr string) []string {
	strSlice := strings.Fields(inptStr)
	resSlice := make([]string, 0, 10)
	dup := make(map[string]int)
	dupKv := make([]kv, 0, len(strSlice))
	for _, word := range strSlice {
		_, exist := dup[word]
		if exist {
			dup[word]++
		} else {
			dup[word] = 1
		}
	}
	for key, val := range dup {
		dupKv = append(dupKv, kv{key, val})
	}
	sort.Slice(dupKv, func(i, j int) bool {
		return dupKv[i].Value > dupKv[j].Value
	})
	var topKv []kv
	if len(dupKv) > 10 {
		topKv = dupKv[0:10]
	} else {
		topKv = dupKv
	}
	var lastVal int
	var curNumCounterPos int
	for idx, kvTemp := range topKv {
		if lastVal != kvTemp.Value {
			sliceKvSort(topKv[curNumCounterPos:idx])
			curNumCounterPos = idx
		} else if lastVal == kvTemp.Value && idx == len(topKv)-1 {
			sliceKvSort(topKv[curNumCounterPos:])
		}
		lastVal = kvTemp.Value
	}
	for _, kvt := range topKv {
		resSlice = append(resSlice, kvt.Key)
	}
	return resSlice
}

func sliceKvSort(slice []kv) {
	sort.Slice(slice, func(i, j int) bool {
		return slice[i].Key < slice[j].Key
	})
}
