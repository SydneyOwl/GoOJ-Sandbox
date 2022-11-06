package main

import (
	"github.com/SydneyOwl/GoOJ-Sandbox/entity"
	"github.com/SydneyOwl/GoOJ-Sandbox/judge"
	"github.com/SydneyOwl/GoOJ-Sandbox/utils"
	"github.com/tidwall/gjson"
	"log"
	"os"
	"time"
)

func LoadProblemFromFile(ppath string, apath string) (problemer entity.Problem, judger entity.Judge) {
	content, err := os.ReadFile(ppath)
	if err != nil {
		panic(err)
	}
	problem := string(content)
	tmp := gjson.Parse(problem)
	problemer = entity.Problem{
		ProblemId:   "nzxTestOnly-",
		Type:        utils.DEFAULT,
		TimeLimit:   tmp.Get("timeLimit").Int(),
		MemoryLimit: tmp.Get("memoryLimit").Int(),
		StackLimit:  tmp.Get("stackLimit").Int(),
		TestCases:   tmp.Get("cases").Array(),
	}
	content1, err := os.ReadFile(apath)
	if err != nil {
		panic(err)
	}
	judge := string(content1)
	tmp = gjson.Parse(judge)
	judger = entity.Judge{
		SubmitId:       "nzxTestOnly-",
		ProblemId:      "nzxTestOnly-",
		SubmitTime:     time.Now().String(),
		Status:         0,
		ErrorMessage:   "",
		TimeConsumed:   0,
		MemoryConsumed: 0,
		Score:          0,
		Code:           tmp.Get("code").String(),
		Language:       tmp.Get("language").String(),
	}
	return
}
func main() {
	a, b := LoadProblemFromFile(os.Args[1], os.Args[2])
	log.Println(judge.Judge(a, b))
}
