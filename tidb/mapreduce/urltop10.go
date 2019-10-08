package main

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

// URLTop10 .
func URLTop10(nWorkers int) RoundsArgs {
	// YOUR CODE HERE :)
	// And don't forget to document your idea.
	var args RoundsArgs
	// round 1: do url count
	args = append(args, RoundArgs{
		MapFunc:    URLCountMap,
		ReduceFunc: URLCountReduce,
		NReduce:    nWorkers,
	})

	args = append(args, RoundArgs{
		MapFunc:    URLTop10Map,
		ReduceFunc: Top10Reduce,
		NReduce:    1,
	})
	return args
}

// URLCountReduce is the reduce function in the first round
func URLCountReduce(key string, values []string) string {
	t := 0
	for _, v := range values {
		num, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}
		t += num
	}
	return fmt.Sprintf("%s %s\n", key, strconv.Itoa(t))
}

// URLCountMap is the map function in the first round
func URLCountMap(filename string, contents string) []KeyValue {
	lines := strings.Split(string(contents), "\n")
	kvs := make([]KeyValue, 0, len(lines))

	cnts := make(map[string]int)
	for _, l := range lines {
		l = strings.TrimSpace(l)
		if len(l) == 0 {
			continue
		}

		if _, ok := cnts[l]; !ok {
			cnts[l] = 1
		} else {
			cnts[l]++
		}
	}

	for k, v := range cnts {
		kvs = append(kvs, KeyValue{Key: k, Value: strconv.Itoa(v)})
	}
	return kvs
}


// URLTop10Map is the map function in the second round
func URLTop10Map(filename string, contents string) []KeyValue {
	lines := strings.Split(contents, "\n")
	kvs := make([]KeyValue, 0, len(lines))
	cnts := make(map[string]int)
	for _, l := range lines {
		if len(l) == 0 {
			continue
		}
 		tmp := strings.Split(l, " ")
		n, err := strconv.Atoi(tmp[1])
		if err != nil {
			panic(err)
		}
		cnts[tmp[0]] = n
	}

	us, cs := TopN(cnts, 10)
	for i := range us {
		kvs = append(kvs, KeyValue{"", us[i] + " " + strconv.Itoa(cs[i])})
	}
	return kvs
}

// Top10Reduce is the reduce function in the second round
func Top10Reduce(key string, values []string) string {
	cnts := make(map[string]int, len(values))
	for _, v := range values {
		v := strings.TrimSpace(v)
		if len(v) == 0 {
			continue
		}
		tmp := strings.Split(v, " ")
		n, err := strconv.Atoi(tmp[1])
		if err != nil {
			panic(err)
		}
		cnts[tmp[0]] = n
	}

	us, cs := TopN(cnts, 10)
	buf := new(bytes.Buffer)
	for i := range us {
		fmt.Fprintf(buf, "%s: %d\n", us[i], cs[i])
	}
	return buf.String()
}