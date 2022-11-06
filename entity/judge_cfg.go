/**
 * @Author: SydneyOwl
 * @Description:
 * @File: judge_cfg
 * @Date: 2022/11/2 16:00
 */

package entity

import (
	"github.com/SydneyOwl/GoOJ-Sandbox/utils"
)

type JudgeGlobalCfg struct {
	ProblemId       string
	JudgeMode       utils.JudgeMode
	UserFileId      string
	UserFileContent string
	//RunDir             string
	SandboxMaxTime int64 //sandbox
	ProblemMaxTime int64 //problem
	MaxMemory      int64
	MaxStack       int64
	RunConfig      utils.RunConfig
}
