/**
 * @Author: SydneyOwl
 * @Description: For oj purpose only.
 * @File: Compiler
 * @Date: 2022/11/1 17:05
 */

package judge

import (
	"github.com/SydneyOwl/GoOJ-Sandbox/utils"
	"log"
)

// Compile compiles specified code and gives an fileId.
func Compile(compileConfig utils.CompileConfig, code string, language string) (string, error) {
	cmd, err := compileConfig.GetParsedBuildCommand()
	//log.Println(compileConfig)
	if err != nil {
		log.Println("Failed to build Command.Abort.")
		return "", NewOJError(OJServerError)
	}
	if !compileConfig.IsValid() {
		log.Println("Unsupported Language" + language)
		return "", NewOJError(UnsupportedLanguageError)
	}
	result := SendCompileReq(compileConfig.MaxCpuTime, compileConfig.MaxRealTime, compileConfig.MaxMemory, 256*1024*1024, compileConfig.SrcName, compileConfig.ExeName, cmd, compileConfig.Envs, code)
	//if parsed fail
	if result.Raw == "" {
		log.Println("Failed to parse json")
		return "", NewOJError(OJServerError)
	}
	if result.Array()[0].Get("exitStatus").Int() != int64(utils.STATUS_ACCEPTED.Status) {
		log.Println("Not accepted")
		//TODO:add detail here with NDOE
		return "", NewFullOJError(CompileError, CompileError.String(), result.Array()[0].Get("files").Get("stderr").String(), result.Array()[0].Get("files").Get("stdout").String())
	}
	//Java has bug here
	//fileId := result.Array()[0].Get("fileIds").Get(compileConfig.ExeName).String()
	filemap := result.Array()[0].Get("fileIds").Value().(map[string]interface{})
	fileId := ""
	for _, v := range filemap {
		fileId = v.(string)
	}
	log.Println(fileId, compileConfig.ExeName)
	//fileId := result.Array()[0].Get("fileIds").Get().String()
	if fileId == "" {
		log.Println("Executable file not found.")
		//TODO:add detail here
		return "", NewOJError(SubmitError)
	}
	return fileId, nil
}
