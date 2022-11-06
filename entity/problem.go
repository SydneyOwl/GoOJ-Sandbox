/**
 * @Author: SydneyOwl
 * @Description:
 * @File: problem
 * @Date: 2022/11/2 13:44
 */

package entity

import (
	"github.com/SydneyOwl/GoOJ-Sandbox/utils"
	"github.com/tidwall/gjson"
)

// No details sir!
type Problem struct {
	ProblemId   string
	Type        utils.JudgeMode
	TimeLimit   int64
	MemoryLimit int64
	StackLimit  int64
	TestCases   []gjson.Result
}
