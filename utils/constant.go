/**
 * @Author: SydneyOwl
 * @Description: Constants...I'd like to say most of the contents in this file comes from HOJ.
 * @File: Constant
 * @Date: 2022/11/1 14:41
 */

package utils

import (
	"reflect"
	"strings"
)

type OJStatus struct {
	Status int
	Name   string
}
type CompileConfig struct {
	Language    string
	SrcName     string
	ExeName     string
	MaxCpuTime  int64
	MaxRealTime int64
	MaxMemory   int64
	Command     string
	Envs        []string
}

type JudgeMode struct {
	Mode string
}
type RunConfig struct {
	Language string
	Command  string
	ExeName  string
	Envs     []string
}

// IsValid checks if struct is not being initialized.
func (compileConfig CompileConfig) IsValid() bool {
	return !reflect.DeepEqual(CompileConfig{}, compileConfig)
}
func (runConfig RunConfig) IsValid() bool {
	return !reflect.DeepEqual(RunConfig{}, runConfig)
}

// GetCompilerByLanguage returns proper compiler
func GetCompilerByLanguage(language string) CompileConfig {
	for _, v := range supported_complier {
		if v.Language == language {
			return v
		}
	}
	return CompileConfig{}
}

// GetJudgeMode returns judge mode
func GetJudgeMode(mode string) JudgeMode {
	for _, v := range supported_mode {
		if v.Mode == mode {
			return v
		}
	}
	return JudgeMode{}
}

// GetRunner returns judge mode
func GetRunner(language string) RunConfig {
	for _, v := range supported_runcfg {
		if v.Language == language {
			return v
		}
	}
	return RunConfig{}
}
func GetStatusMsgByCode(stat int) string {
	for _, v := range supported_ojstatus {
		if v.Status == stat {
			return v.Name
		}
	}
	return "No Status"
}

func GetStatusCodeByMsg(msg string) int {
	for _, v := range supported_ojstatus {
		if v.Name == msg {
			return v.Status
		}
	}
	return STATUS_NULL.Status
}

// GetParsedBuildCommand gives build command which is translated.
func (runConfig RunConfig) GetParsedBuildCommand() ([]string, error) {
	var exp1 = strings.ReplaceAll(runConfig.Command, "{0}", TMPFS_DIR)
	var exp2 = strings.ReplaceAll(exp1, "{1}", runConfig.ExeName)
	//return exp3, nil
	return strings.Split(exp2, " "), nil
}

// GetParsedBuildCommand Reversed func GetParsedBuildCommand
func (compileConfig CompileConfig) GetParsedBuildCommand() ([]string, error) {
	var exp1 = strings.ReplaceAll(compileConfig.Command, "{0}", TMPFS_DIR)
	var exp2 = strings.ReplaceAll(exp1, "{1}", compileConfig.SrcName)
	var exp3 = strings.ReplaceAll(exp2, "{2}", compileConfig.ExeName)
	//log.Println(exp3)
	//Here we translate the command line. Same as HOJ but translated into golang...but seems useless

	//if exp3 != "" {
	//	state := 0
	//	sepCount := strings.Count(exp3, "\"' ")
	//	temp := strings.Split(exp3, "\"' ")
	//	tok := make([]string, 0)
	//	// Re-add sep
	//	for k, v := range temp {
	//		tok = append(tok, v)
	//		if (k == 0 || k%2 == 0) && sepCount > 0 {
	//			sepCount--
	//			tok = append(tok, "\"' ")
	//		}
	//	}
	//
	//	result := make([]string, 0)
	//	current := ""
	//	lastTokenHasBeenQuoted := false
	//	pos := 0
	//	for {
	//		for pos < len(tok)-1 {
	//			pos++
	//			nextTok := tok[pos]
	//			switch state {
	//			case 1:
	//				if "'" == nextTok {
	//					lastTokenHasBeenQuoted = true
	//					state = 0
	//				} else {
	//					current += nextTok
	//				}
	//				continue
	//			case 2:
	//				if "\"" == nextTok {
	//					lastTokenHasBeenQuoted = true
	//					state = 0
	//				} else {
	//					current += nextTok
	//				}
	//				continue
	//			}
	//
	//			if "'" == nextTok {
	//				state = 1
	//			} else if "\"" == nextTok {
	//				state = 2
	//			} else if " " == nextTok {
	//				if lastTokenHasBeenQuoted || len(current) > 0 {
	//					result = append(result, current)
	//					current = ""
	//				}
	//			} else {
	//				current += nextTok
	//			}
	//
	//			lastTokenHasBeenQuoted = false
	//		}
	//
	//		if lastTokenHasBeenQuoted || len(current) > 0 {
	//			result = append(result, current)
	//		}
	//
	//		if state != 1 && state != 2 {
	//			return result, nil
	//		}
	//		return make([]string, 0), errors.New("unbalanced quote")
	//	}
	//} else {
	//	return make([]string, 0), errors.New("emptyCommand")
	//}
	//....
	if compileConfig.Language == "Java" { //bin/bash -c "javac -encoding=utf-8 {1} && jar -cvf {2} *.class"
		return []string{"/bin/bash", "-c", "javac -encoding utf-8 Main.java && jar -cvf Main.jar *.class"}, nil
	}
	print(exp3)
	return strings.Split(exp3, " "), nil
}

const (
	TMPFS_DIR         = "/w"
	TEST_CASE_DIR     = "_temp_data/problems"
	RUN_WORKSPACE_DIR = "_temp/workspace"
)

var (
	STATUS_NOT_SUBMITTED         = OJStatus{-10, "Not Submitted"}
	STATUS_CANCELLED             = OJStatus{-4, "Cancelled"}
	STATUS_PRESENTATION_ERROR    = OJStatus{-3, "Presentation Error"}
	STATUS_COMPILE_ERROR         = OJStatus{-2, "Compile Error"}
	STATUS_WRONG_ANSWER          = OJStatus{-1, "Wrong Answer"}
	STATUS_ACCEPTED              = OJStatus{0, "Accepted"}
	STATUS_TIME_LIMIT_EXCEEDED   = OJStatus{1, "Time Limit Exceeded"}
	STATUS_MEMORY_LIMIT_EXCEEDED = OJStatus{2, "Memory Limit Exceeded"}
	STATUS_RUNTIME_ERROR         = OJStatus{3, "Runtime Error"}
	STATUS_SYSTEM_ERROR          = OJStatus{4, "System Error"}
	STATUS_PENDING               = OJStatus{5, "Pending"}
	STATUS_COMPILING             = OJStatus{6, "Compiling"}
	STATUS_JUDGING               = OJStatus{7, "Judging"}
	STATUS_PARTIAL_ACCEPTED      = OJStatus{8, "Partial Accepted"}
	STATUS_SUBMITTING            = OJStatus{9, "Submitting"}
	STATUS_SUBMITTED_FAILED      = OJStatus{10, "Submitted Failed"}
	STATUS_NULL                  = OJStatus{15, "No Status"}

	defaultEnv = []string{
		"PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin",
		"LANG=en_US.UTF-8",
		"LC_ALL=en_US.UTF-8",
		"LANGUAGE=en_US:en",
		"HOME=/w",
	}

	python3Env = []string{
		"LANG=en_US.UTF-8",
		"LANGUAGE=en_US:en",
		"LC_ALL=en_US.UTF-8",
		"PYTHONIOENCODING=utf-8",
	}

	golangEnv = []string{
		"GODEBUG=madvdontneed=1",
		"GOCACHE=off",
		"PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin",
		"LANG=en_US.UTF-8",
		"LANGUAGE=en_US:en",
		"LC_ALL=en_US.UTF-8",
	}
	//compile
	CC         = CompileConfig{"C", "main.c", "main", 3000, 10000, 256 * 1024 * 1024, "/usr/bin/gcc -DONLINE_JUDGE -w -fmax-errors=1 -std=c11 {1} -lm -o {2}", defaultEnv}
	CCWithO2   = CompileConfig{"C With O2", "main.c", "main", 3000, 10000, 256 * 1024 * 1024, "/usr/bin/gcc -DONLINE_JUDGE -O2 -w -fmax-errors=1 -std=c11 {1} -lm -o {2}", defaultEnv}
	CCPP       = CompileConfig{"C++", "main.cpp", "main", 10000, 20000, 512 * 1024 * 1024, "/usr/bin/g++ -DONLINE_JUDGE -w -fmax-errors=1 -std=c++14 {1} -lm -o {2}", defaultEnv}
	CCPPWithO2 = CompileConfig{"C++ With O2", "main.cpp", "main", 10000, 20000, 512 * 1024 * 1024, "/usr/bin/g++ -DONLINE_JUDGE -O2 -w -fmax-errors=1 -std=c++14 {1} -lm -o {2}", defaultEnv}
	CJAVA      = CompileConfig{"Java", "Main.java", "Main.jar", 10000, 20000, 512 * 1024 * 1024, "/bin/bash -c \"javac -encoding=utf-8 {1} && jar -cvf {2} *.class\"", defaultEnv}
	CPYTHON2   = CompileConfig{"Python2", "main.py", "main.pyc", 3000, 10000, 128 * 1024 * 1024, "/usr/bin/python -m py_compile ./{1}", defaultEnv}
	CPYTHON3   = CompileConfig{"Python3", "main.py", "__pycache__/main.cpython-37.pyc", 3000, 10000, 128 * 1024 * 1024, "/usr/bin/python3.7 -m py_compile ./{1}", defaultEnv}
	CGOLANG    = CompileConfig{"Golang", "main.go", "main", 3000, 5000, 512 * 1024 * 1024, "/usr/bin/go build -o {2} {1}", defaultEnv}
	CCS        = CompileConfig{"C#", "Main.cs", "main", 5000, 10000, 512 * 1024 * 1024, "/usr/bin/mcs -optimize+ -out:{0}/{2} {0}/{1}", defaultEnv}
	CPyPy2     = CompileConfig{"PyPy2", "main.py", "__pycache__/main.pypy-73.pyc", 3000, 10000, 256 * 1024 * 1024, "/usr/bin/pypy -m py_compile {0}/{1}", defaultEnv}
	CPyPy3     = CompileConfig{"PyPy3", "main.py", "__pycache__/main.pypy38.pyc", 3000, 10000, 256 * 1024 * 1024, "/usr/bin/pypy3 -m py_compile {0}/{1}", defaultEnv}

	/*
		CSPJ_C           = CompileConfig{"SPJ-C", "spj.c", "spj", 3000, 5000, 512 * 1024 * 1024, "/usr/bin/gcc -DONLINE_JUDGE -O2 -w -fmax-errors=3 -std=c99 {1} -lm -o {2}", defaultEnv}
		CSPJ_CPP         = CompileConfig{"SPJ-C++", "spj.cpp", "spj", 10000, 20000, 512 * 1024 * 1024, "/usr/bin/g++ -DONLINE_JUDGE -O2 -w -fmax-errors=3 -std=c++14 {1} -lm -o {2}", defaultEnv}
		CINTERACTIVE_C   = CompileConfig{"INTERACTIVE-C", "interactive.c", "interactive", 3000, 5000, 512 * 1024 * 1024, "/usr/bin/gcc -DONLINE_JUDGE -O2 -w -fmax-errors=3 -std=c99 {1} -lm -o {2}", defaultEnv}
		CINTERACTIVE_CPP = CompileConfig{"INTERACTIVE-C++", "interactive.cpp", "interactive", 10000, 20000, 512 * 1024 * 1024, "/usr/bin/g++ -DONLINE_JUDGE -O2 -w -fmax-errors=3 -std=c++14 {1} -lm -o {2}", defaultEnv}
	*/

	RC       = RunConfig{"C", "{0}/{1}", "main", defaultEnv}
	RCPP     = RunConfig{"C++", "{0}/{1}", "main", defaultEnv}
	RJAVA    = RunConfig{"Java", "/usr/bin/java -Dfile.encoding=UTF-8 -cp {0}/{1} Main", "Main.jar", defaultEnv}
	RPYTHON2 = RunConfig{"Python2", "/usr/bin/python {1}", "main", defaultEnv}
	RPYTHON3 = RunConfig{"Python3", "/usr/bin/python3.7 {1}", "main", python3Env}
	RGOLANG  = RunConfig{"Golang", "{0}/{1}", "main", golangEnv}
	RCS      = RunConfig{"C#", "/usr/bin/mono {0}/{1}", "main", defaultEnv}
	RPHP     = RunConfig{"PHP", "/usr/bin/php {1}", "main.php", defaultEnv}
	RJS_NODE = RunConfig{"JavaScript Node", "/usr/bin/node {1}", "main.js", defaultEnv}
	//RJS_V8   = RunConfig{"JavaScript V8", "/usr/bin/jsv8/d8 {1}", "main.js", defaultEnv}

	/*
		RCWithO2   = RunConfig{"C With O2", "{0}/{1}", "main", defaultEnv}
		RCPPWithO2 = RunConfig{"C++ With O2", "{0}/{1}", "main", defaultEnv}
		RPyPy2     = RunConfig{"PyPy2", "/usr/bin/pypy {1}", "main.pyc", defaultEnv}
		RPyPy3     = RunConfig{"PyPy3", "/usr/bin/pypy3 {1}", "main.pyc", python3Env}
		RSPJ_C = RunConfig{"SPJ-C", "{0}/{1} {2} {3} {4}", "spj", defaultEnv}
		RSPJ_CPP = RunConfig{"SPJ-C++", "{0}/{1} {2} {3} {4}", "spj", defaultEnv}
		RINTERACTIVE_C = RunConfig{"INTERACTIVE-C", "{0}/{1} {2} {3} {4}", "interactive", defaultEnv}
		RINTERACTIVE_CPP    = RunConfig{"INTERACTIVE-C++", "{0}/{1} {2} {3} {4}", "interactive", defaultEnv}
	*/

	DEFAULT            = JudgeMode{"default"}
	supported_complier = []CompileConfig{CC, CCWithO2, CCPP, CCPPWithO2, CJAVA, CPYTHON2, CPYTHON3, CGOLANG, CCS, CPyPy2, CPyPy3} //, CSPJ_C, CSPJ_CPP, CINTERACTIVE_C, CINTERACTIVE_CPP}
	supported_mode     = []JudgeMode{DEFAULT}
	supported_runcfg   = []RunConfig{RC, RCPP, RJAVA, RPYTHON2, RPYTHON3, RGOLANG, RCS, RPHP, RJS_NODE}
	supported_ojstatus = []OJStatus{STATUS_NOT_SUBMITTED, STATUS_CANCELLED, STATUS_PRESENTATION_ERROR, STATUS_COMPILE_ERROR, STATUS_WRONG_ANSWER,
		STATUS_ACCEPTED, STATUS_TIME_LIMIT_EXCEEDED, STATUS_MEMORY_LIMIT_EXCEEDED, STATUS_RUNTIME_ERROR, STATUS_SYSTEM_ERROR, STATUS_PENDING, STATUS_COMPILING,
		STATUS_JUDGING, STATUS_PARTIAL_ACCEPTED, STATUS_SUBMITTING, STATUS_SUBMITTED_FAILED, STATUS_NULL}
)
