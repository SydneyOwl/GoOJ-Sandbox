/**
 * @Author: SydneyOwl
 * @Description:
 * @File: judge_mode
 * @Date: 2022/11/2 13:41
 */

package judge

import (
	"github.com/SydneyOwl/GoOJ-Sandbox/entity"
	"github.com/SydneyOwl/GoOJ-Sandbox/utils"
	"github.com/pkg/errors"
	"github.com/tidwall/gjson"
	"log"
	"os"
	"path"
	"sort"
)

func loadProblemInfo(dir string) (gjson.Result, error) {
	file, err := os.ReadFile(path.Join(dir, "info"))
	if err != nil {
		return gjson.Result{}, NewOJError(OJServerError)
	}
	return gjson.ParseBytes(file), nil
}

// Judge judges problem; However we only support normal mode:(
func Judge(problem entity.Problem, judge entity.Judge) map[string]interface{} {
	compileConfig := utils.GetCompilerByLanguage(judge.Language)
	userFileId := ""
	var err error
	var mode = problem.Type
	//var testDir string
	//var allCase []gjson.Result
	defer func() {
		if userFileId != "" {
			DelFile(userFileId)
		}
	}()
	// Exclude javascript and php
	if compileConfig.IsValid() {
		userFileId, err = Compile(compileConfig, judge.Code, judge.Language)
	}
	//if err == nil {
	//	testDir = path.Join(utils.TEST_CASE_DIR, "problem_"+problem.ProblemId)
	//	testCasesInfo, err = loadProblemInfo(testDir)
	//	mode = utils.GetJudgeMode(testCasesInfo.Get("type").String())
	//}
	var allCase []runResult
	if err == nil {
		allCase, err = JudgeAllCase(problem, judge.Language, userFileId, judge.Code, mode.Mode)
	}
	//if err == nil {
	//	//return allCase
	//	return getJudgeInfo(allCase, problem, judge, mode)
	//}
	//log.Println(errors.Cause(err))
	// TODO:Here we update the db.....(status
	if err != nil {
		result := make(map[string]interface{})
		if errors.Is(err, NewRawOJError(OJServerError)) {
			result["code"] = utils.STATUS_SYSTEM_ERROR.Status
			result["errMsg"] = "Something went wrong with Gooj"
			log.Printf("Judge system error!Submit id:%s,Problem id:%s,Error:%s\n", judge.SubmitId, judge.ProblemId, StackTraceOJError(err))
		} else if errors.Is(err, NewRawOJError(SubmitError)) {
			result["code"] = utils.STATUS_SUBMITTED_FAILED.Status
			result["errMsg"] = err.Error()
			log.Printf("Submit error!Submit id:%s,Problem id:%s,Error:%s\n", judge.SubmitId, judge.ProblemId, StackTraceOJError(err))
		} else if errors.Is(err, NewRawOJError(CompileError)) {
			result["code"] = utils.STATUS_COMPILE_ERROR.Status
			result["errMsg"] = err.Error()
			log.Printf("Compile Error!Submit id:%s,Problem id:%s,Error:%s\n", judge.SubmitId, judge.ProblemId, StackTraceOJError(err))
		} else {
			result["code"] = utils.STATUS_SYSTEM_ERROR.Status
			result["errMsg"] = "Something went wrong with oj."
			log.Printf("Error!Submit id:%s,Problem id:%s,Error:%s\n", judge.SubmitId, judge.ProblemId, StackTraceOJError(err))
		}
		//errors.Cause()
		//errors.As()
		return map[string]interface{}{"answer": result, "status": utils.STATUS_CANCELLED.Status}
	} else {
		return getJudgeInfo(allCase, problem.TestCases)
	}
}

func getJudgeInfo(allCase []runResult, cases []gjson.Result) map[string]interface{} {
	//var allCase []runResult = make([]runResult, 0)
	//copy(allCase, al)
	//log.Println(allCase)
	allRes := make([]map[string]interface{}, 1, 1)
	sort.Slice(allCase, func(i, j int) bool {
		return allCase[i].Id < allCase[j].Id
	})
	var total int
	isErr, isPss := false, false
	//Here we judge total point
	for i, v := range allCase {
		var score int
		//log.Println([]byte(utils.CompressStr(v.Stdout)), []byte(cases[i].Get("output").String()))
		if v.ExitStatus == utils.STATUS_ACCEPTED.Status {
			if v.Stdout == cases[i].Get("output").String() {
				score = int(cases[i].Get("score").Int())
				isPss = true
			} else {
				isErr = true
				v.ExitStatus = utils.STATUS_WRONG_ANSWER.Status
			}
		} else {
			isErr = true
		}
		total += score
		//log.Println(v.ExitStatus)
		allRes = append(allRes, map[string]interface{}{
			"caseId":      v.Id,
			"code":        v.ExitStatus,
			"msg":         utils.GetStatusMsgByCode(int(v.ExitStatus)),
			"time":        v.Time,
			"memory":      v.Memory,
			"singleScore": score,
		})
	}
	status := utils.STATUS_NULL
	if isErr && isPss {
		status = utils.STATUS_PARTIAL_ACCEPTED
	} else if isErr && !isPss {
		status = utils.STATUS_WRONG_ANSWER
	} else {
		status = utils.STATUS_ACCEPTED
	}
	//this gives directly to frontend?Nope;However this is only a test.
	return map[string]interface{}{
		"answer":     allRes,
		"totalScore": total,
		"status":     status.Status,
		"msg":        status.Name,
	}
}
