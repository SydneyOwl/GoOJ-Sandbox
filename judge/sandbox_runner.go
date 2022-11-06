/**
 * @Author: SydneyOwl
 * @Description:
 * @File: sandbox_runner.go
 * @Date: 2022/11/1 14:42
 */

package judge

import (
	"github.com/SydneyOwl/GoOJ-Sandbox/utils"
	"github.com/tidwall/gjson"
	"io"
	"log"
)

const (
	STACK_LIMIT_MB     = 128
	STDIO_SIZE_MB      = 32
	SANDBOX_BASE_URL   = "http://localhost:6050"
	MAX_PROCESS_NUMBER = 128
)

var COMPILE_FILES = []map[string]interface{}{
	{
		"content": "",
	},
	{
		"name": "stdout",
		"max":  1024 * 1024 * STDIO_SIZE_MB,
	},
	{
		"name": "stderr",
		"max":  1024 * 1024 * STDIO_SIZE_MB,
	},
}

func init() {
}

// DelFile Deletes file in container.
func DelFile(fileId string) {
	res, err := utils.SendSimpleDeleteRequest(SANDBOX_BASE_URL + "/file/" + fileId)
	if err != nil || res.StatusCode != 200 {
		log.Println("未找到文件缓存，id:", fileId)
	}
}

// runCode sends compile request.
func runCode(url string, param map[string]interface{}) (gjson.Result, error) {
	res, err := utils.SendJsonPostRequest(url, param)
	if err != nil || res.StatusCode != 200 {
		log.Println("RunCode请求出错")
		return gjson.Result{}, NewOJError(OJServerError)
	}
	ans, _ := io.ReadAll(res.Body)
	gsonResult := gjson.Parse(string(ans))
	return gsonResult, nil
}

// SendCompileReq build and give compile request.
func SendCompileReq(maxCpuTime int64, maxRealTime int64, maxMemory int64, maxStack int64, srcName string, exeName string, args []string, envs []string, code string) gjson.Result {
	// 构建命令
	cmd := map[string]interface{}{
		"args":        args,
		"env":         envs,
		"files":       COMPILE_FILES,
		"cpuLimit":    maxCpuTime * 1000 * 1000,
		"clockLimit":  maxRealTime * 1000 * 1000,
		"memoryLimit": maxMemory,
		"procLimit":   MAX_PROCESS_NUMBER,
		"stackLimit":  maxStack,
	}
	fileContent := map[string]string{
		"content": code,
	}
	copyIn := map[string]interface{}{
		srcName: fileContent,
	}
	cmd["copyIn"] = copyIn
	cmd["copyOut"] = [2]string{"stdout", "stderr"}
	//生成用户程序的缓存文件，即生成用户程序id
	cmd["copyOutCached"] = [1]string{exeName}
	param := map[string]interface{}{
		"cmd": [1]interface{}{
			cmd,
		},
	}
	//a, _ := json.Marshal(param)
	//log.Println(string(a))
	result, err := runCode(SANDBOX_BASE_URL+"/run", param)
	if err != nil {
		log.Println("无法编译！")
		return gjson.Result{}
	}
	return result
}
