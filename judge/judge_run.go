/**
 * @Author: SydneyOwl
 * @Description:
 * @File: judge_run.java
 * @Date: 2022/11/2 15:05
 */

package judge

import (
	"github.com/SydneyOwl/GoOJ-Sandbox/entity"
	"github.com/SydneyOwl/GoOJ-Sandbox/utils"
	"github.com/panjf2000/ants/v2"
	"github.com/pkg/errors"
	"github.com/tidwall/gjson"
	"io"
	"log"
	"strings"
	"sync"
	"time"
)

type CommandArgs struct {
	Id             int
	JudgeGlobalDTO entity.JudgeGlobalCfg
	Input          string
	channel        chan runResult
}
type runResult struct {
	Id         int
	ExitStatus int
	Time       int64
	Memory     int64
	Stdout     string
	Stderr     string
	Err        error
}

// JudgeAllCase judge all cases; However only default is supported
func JudgeAllCase(problem entity.Problem, judgeLanguage string,
	userFileId string, userFileContent string, judgeCaseMode string) ([]runResult, error) {
	cases := problem.TestCases
	//cases := testCasesInfo.Get("cases").Array()
	testTime := problem.TimeLimit + 200 //HOJ does same here, but I don't know why [默认给题目限制时间+200ms用来测评]
	judgeMode := problem.Type
	if judgeMode.Mode == "" {
		log.Println("Cannot fetch mode.")
		return []runResult{}, NewOJError(OJServerError)
	}
	runConfig := utils.GetRunner(judgeLanguage)
	if !runConfig.IsValid() {
		log.Println("RunConfig not found!")
		return []runResult{}, NewOJError(OJServerError)
	}
	judgeGlobalCfg := entity.JudgeGlobalCfg{
		ProblemId:       problem.ProblemId,
		JudgeMode:       judgeMode,
		UserFileId:      userFileId,
		UserFileContent: userFileContent,
		//RunDir:             runDir,
		SandboxMaxTime: testTime,
		ProblemMaxTime: problem.TimeLimit,
		MaxMemory:      problem.MemoryLimit,
		MaxStack:       problem.StackLimit,
		RunConfig:      runConfig,
	}
	switch judgeCaseMode {
	case utils.DEFAULT.Mode:
		return defaultJudgeAllCase(cases, judgeGlobalCfg)
	default:
		log.Println("Error:Cannot find specified judge")
		return []runResult{}, NewOJError(OJServerError)
	}
}
func runSingleCommand(i interface{}) {
	arg := i.(CommandArgs)
	runCmd := buildRunCommand(arg.JudgeGlobalDTO, arg.Input)
	resp, err := utils.SendJsonPostRequest(SANDBOX_BASE_URL+"/run", runCmd)
	if err != nil {
		arg.channel <- runResult{Err: errors.WithMessage(err, "Cannot sendJson")}
		return
	}
	ans, _ := io.ReadAll(resp.Body)
	ansbytes := gjson.ParseBytes(ans)
	//time.Sleep(time.Minute)
	arg.channel <- runResult{
		Id:         arg.Id,
		ExitStatus: utils.GetStatusCodeByMsg(ansbytes.Array()[0].Get("status").String()),
		Time:       ansbytes.Array()[0].Get("time").Int(),
		Memory:     ansbytes.Array()[0].Get("memory").Int(),
		Stdout:     strings.Replace(ansbytes.Array()[0].Get("files").Get("stdout").String(), "\n", "", -1),
		Stderr:     ansbytes.Array()[0].Get("files").Get("stderr").String(),
		Err:        nil}
}

func defaultJudgeAllCase(testcaseList []gjson.Result, judgeGlobalDTO entity.JudgeGlobalCfg) ([]runResult, error) {
	var wg sync.WaitGroup
	done := make(chan struct{})
	result := make(chan runResult, len(testcaseList))
	final := make([]runResult, 0)
	pool, _ := ants.NewPoolWithFunc(120, func(i interface{}) {
		runSingleCommand(i)
		wg.Done()
	})
	defer pool.Release()
	for i, v := range testcaseList {
		input := v.Get("input").String()
		wg.Add(1)
		_ = pool.Invoke(CommandArgs{
			Id:             i,
			JudgeGlobalDTO: judgeGlobalDTO,
			Input:          input,
			channel:        result,
		})
	}
	go func() {
		wg.Wait()
		close(result)
		done <- struct{}{}
	}()
	select {
	case <-done:
		log.Println("All routines exited successfully.")
		for {
			v, ok := <-result
			if !ok {
				log.Println("Channel closed")
				break
			}
			final = append(final, v)
		}
		for _, v := range final {
			if v.Err != nil {
				return []runResult{}, errors.WithMessagef(v.Err, "Answer invaild since %d has error", v.Id)
			}
		}
		return final, nil
	case <-time.After(10 * time.Second):
		log.Println("Timeout. Kill all routines")
		return []runResult{}, NewOJError(OJServerError) //We consider this as a server error
	}
}

// buildRunCommand build run command for a case
func buildRunCommand(judgeGlobalDTO entity.JudgeGlobalCfg, input string) map[string]interface{} {
	var files = make([]map[string]interface{}, 0)
	files = append(files,
		map[string]interface{}{
			"content": input,
		},
		map[string]interface{}{
			"name": "stdout",
			"max":  10240,
		},
		map[string]interface{}{
			"name": "stderr",
			"max":  10240,
		})
	args, _ := judgeGlobalDTO.RunConfig.GetParsedBuildCommand()
	env := judgeGlobalDTO.RunConfig.Envs
	copyIn := map[string]interface{}{}
	exeFile := map[string]interface{}{}
	if judgeGlobalDTO.UserFileId == "" {
		exeFile["content"] = judgeGlobalDTO.UserFileContent
	} else {
		exeFile["fileId"] = judgeGlobalDTO.UserFileId
	}
	copyIn[judgeGlobalDTO.RunConfig.ExeName] = exeFile
	cmd := map[string]interface{}{
		"args":        args,
		"env":         env,
		"files":       files,
		"cpuLimit":    judgeGlobalDTO.SandboxMaxTime * 1000 * 1000,
		"clockLimit":  judgeGlobalDTO.SandboxMaxTime * 1000 * 1000 * 3,
		"memoryLimit": (judgeGlobalDTO.MaxMemory + 100) * 1024 * 1024,
		"procLimit":   MAX_PROCESS_NUMBER,
		"stackLimit":  judgeGlobalDTO.MaxStack * 1024 * 1024,
		"copyIn":      copyIn,
		"copyOut":     []string{"stdout", "stderr"},
	}
	param := map[string]interface{}{
		"cmd": [1]interface{}{
			cmd,
		},
	}
	return param
}
