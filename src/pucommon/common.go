package pucommon

import (
	"fmt"
	"hash/fnv"
	"log"
	"runtime"
	"strconv"
)

const isDubeg = false

// PULOG DEBUG func
func PULOG(args ...interface{}) {
	if !isDubeg {
		return
	}
	var funcname string
	pc, _, line, ok := runtime.Caller(1)
	if ok {
		funcname = runtime.FuncForPC(pc).Name()
	}
	log.Println("PULOG:", funcname, line, "|", args)
}

// KeyValue the url store in this way
type KeyValue struct {
	Key   string
	Value uint64
}

func Ihash(s string) int {
	h := fnv.New32a()
	h.Write([]byte(s))
	return int(h.Sum32() & 0x7fffffff)
}

func MergeFileName(baseDir string, jobName string, taskNo int) string {
	return baseDir + jobName + "_" + strconv.Itoa(taskNo)
}

type TaskArgs struct {
	TaskFileName string
	TopN         uint32
	MapTaskNum   int
}

func NewTaskArgs(args []string) (*TaskArgs, bool) {
	usage := func() {
		fmt.Println("3 args, 1. filename 2. topn 3. map task num")
	}
	if len(args) != 3 {
		usage()
		return nil, false
	}
	taskargs := &TaskArgs{}
	taskargs.TaskFileName = args[0]
	var err error
	var num int
	num, err = strconv.Atoi(args[1])
	if err != nil {
		usage()
		return nil, false
	}
	taskargs.TopN = uint32(num)
	taskargs.MapTaskNum, err = strconv.Atoi(args[2])
	if err != nil {
		usage()
		return nil, false
	}
	return taskargs, true
}
