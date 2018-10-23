package mapreduce

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func doReduce(
	jobName string, // the name of the whole MapReduce job
	reduceTask int, // which reduce task this is
	outFile string, // write the output here
	nMap int, // the number of map tasks that were run ("M" in the paper)
	reduceF func(key string, values []string) string,
) {
	//
	// doReduce manages one reduce task: it should read the intermediate
	// files for the task, sort the intermediate key/value pairs by key,
	// call the user-defined reduce function (reduceF) for each key, and
	// write reduceF's output to disk.
	//
	// You'll need to read one intermediate file from each map task;
	// reduceName(jobName, m, reduceTask) yields the file
	// name from map task m.
	//
	// Your doMap() encoded the key/value pairs in the intermediate
	// files, so you will need to decode them. If you used JSON, you can
	// read and decode by creating a decoder and repeatedly calling
	// .Decode(&kv) on it until it returns an error.
	//
	// You may find the first example in the golang sort package
	// documentation useful.
	//
	// reduceF() is the application's reduce function. You should
	// call it once per distinct key, with a slice of all the values
	// for that key. reduceF() returns the reduced value for that key.
	//
	// You should write the reduce output as JSON encoded KeyValue
	// objects to the file named outFile. We require you to use JSON
	// because that is what the merger than combines the output
	// from all the reduce tasks expects. There is nothing special about
	// JSON -- it is just the marshalling format we chose to use. Your
	// output code will look something like this:
	//
	// enc := json.NewEncoder(file)
	// for key := ... {
	// 	enc.Encode(KeyValue{key, reduceF(...)})
	// }
	// file.Close()
	//
	// Your code here (Part I).
	//
	resultMap := make(map[string][]string)
	for m := 0; m < nMap; m++ {
		inFile := reduceName(jobName, m, reduceTask)
		infile, err := os.Open(inFile)
		if err != nil {
			fmt.Printf("doReduce:%v", err)
		}
		br := bufio.NewReader(infile)
		dec := json.NewDecoder(br)
		var kv KeyValue
		for {
			err = dec.Decode(&kv)
			if err != nil {
				break
			}
			resultMap[kv.Key] = append(resultMap[kv.Key], kv.Value)
		}

		rf, err := os.OpenFile(outFile, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0660)
		if err != nil {
			log.Println(err)
		}
		enc := json.NewEncoder(rf)
		for k, v := range resultMap {
			enc.Encode(KeyValue{k, reduceF(k, v)})
		}
		infile.Close()
	}
}
