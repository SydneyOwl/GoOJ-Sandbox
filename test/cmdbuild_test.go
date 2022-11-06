/**
 * @Author: SydneyOwl
 * @Description:
 * @File: cmdbuild_test.go
 * @Date: 2022/11/2 19:15
 */

package test

import (
	"github.com/SydneyOwl/GoOJ-Sandbox/entity"
	"github.com/SydneyOwl/GoOJ-Sandbox/judge"
	"github.com/SydneyOwl/GoOJ-Sandbox/utils"
	"github.com/tidwall/gjson"
	"log"
	"testing"
)

func TestCMD(t *testing.T) {
	//t.Run("cmd", func(t *testing.T) {
	//	jd := entity.JudgeGlobalCfg{
	//		ProblemId:       "123",
	//		JudgeMode:       utils.DEFAULT,
	//		UserFileId:      "1323",
	//		UserFileContent: "fasdfadsfasd",
	//		SandboxMaxTime:  32,
	//		ProblemMaxTime:  322,
	//		MaxMemory:       133,
	//		MaxStack:        34432,
	//		RunConfig:       utils.RCPP,
	//	}
	//	o, _ := json.Marshal(buildRunCommand(jd, "3131232311"))
	//	fmt.Println(string(o))
	//})
	t.Run("judgeC", func(t *testing.T) {
		p := entity.Problem{
			ProblemId:   "1",
			Type:        utils.DEFAULT,
			TimeLimit:   1000,
			MemoryLimit: 100,
			StackLimit:  1000,
			TestCases: gjson.Parse(`
						{
						  "type": "default",
						  "cases": [
							{
							  "input": "1 3",
							  "output": "41",
							  "score": 1
							},
							{
							  "input": "3 337",
							  "output": "340",
							  "score": 1
							},
							{
							  "input": "2 5",
							  "output": "3",
							  "score": 1
							},
						  ]
						}
						`).Get("cases").Array(),
		}
		a := entity.Judge{
			SubmitId:       "kdo",
			ProblemId:      "1",
			SubmitTime:     "",
			Status:         0,
			ErrorMessage:   "",
			TimeConsumed:   0,
			MemoryConsumed: 0,
			Score:          0,
			//PASSED
			Code:     "#include <iostream>\nusing namespace std;\n \nint main()\n{\n    int firstNumber, secondNumber, sumOfTwoNumbers;\n    \n    cout << \"输入两个整数: \";\n    cin >> firstNumber >> secondNumber;\n \n    // 相加\n    sumOfTwoNumbers = firstNumber + secondNumber;\n \n    // 输出\n    cout << sumOfTwoNumbers;     \n \n    return 0;\n}",
			Language: "C++",

			//ERROR
			//Code:     "import java.util.Scanner;\n\npublic class Main {\n    public static void main(String[] args) {\n        Scanner scanner = new Scanner(System.in);\n        while(scanner.hasNext()){\n            int a = scanner.nextInt();\n            int b = scanner.nextInt();\n            int ans = a + b;\n            System.out.println(ans);\n        }\n    }\n}\n",
			//Language: "Java",

			//Code:     "\n#include<stdio.h>\nint main ()\n{\nint x,y,sum=0;\nscanf(\"%d,%d\",&x,&y);\nsum=x+y;\nprintf(\"%d\",sum);\nreturn 0;\n}\n",
			//Language: "C",
		}
		log.Println(judge.Judge(p, a))
	})
}
