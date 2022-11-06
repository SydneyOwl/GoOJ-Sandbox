# README

## 部署

这是连接沙盒程序的服务端部分。

### 沙箱

docker:

```shell
docker build -t xxx/xxx -f xxx .
docker run --name exec -d --privileged -e ES_ENABLE_GRPC=1 -e ES_ENABLE_METRICS=1 -e ES_ENABLE_DEBUG=1 -e ES_GRPC_ADDR=:6051 -e ES_HTTP_ADDR=:6050 -p 6051:6051 -p 6050:6050 judger_exec
```

### 调用

#### Terminal

```bash
go build command_line.go
./command_line ./example/problem ./example/answer
```

#### Internal

程序内部调用：
`go get github.com/SydneyOwl/GoOJ-Sandbox`后

```go
package main

import (
	"fmt"
	judge "github.com/SydneyOwl/GoOJ-Sandbox/judge"
	entity "github.com/SydneyOwl/GoOJ-Sandbox/entity"
)

func main() {
	p := entity.Problem{
		ProblemId:   "1",
		Type:        utils.DEFAULT,
		TimeLimit:   100,
		MemoryLimit: 100,
		StackLimit:  1000,
		TestCases: gjson.Parse(`
						{
						  "type": "default",
						  "cases": [
							{
							  "input": "1 3",
							  "output": "4",
							  "score": 1
							},
							{
							  "input": "3 337",
							  "output": "340",
							  "score": 2
							},
							{
							  "input": "2 5",
							  "output": "7",
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
		//PASSED(O0)
		//Code:     "#include <iostream>\nusing namespace std;\n \nint main()\n{\n    int firstNumber, secondNumber, sumOfTwoNumbers;\n    \n     cin >> firstNumber >> secondNumber;\n \n    // 相加\n    sumOfTwoNumbers = firstNumber + secondNumber;\n \n    // 输出\n    cout << sumOfTwoNumbers;     \n \n    return 0;\n}",
		//Language: "C++",

		//PASSED
		//Code:     "import java.util.Scanner;\n\npublic class Main {\n    public static void main(String[] args) {\n        Scanner scanner = new Scanner(System.in);\n        while(scanner.hasNext()){\n            int a = scanner.nextInt();\n            int b = scanner.nextInt();\n            int ans = a + b;\n            System.out.print(ans);\n        }\n    }\n}\n",
		//Language: "Java",

		//PASSED(O0)
		//Code:     "\n#include<stdio.h>\nint main (void)\n{\nint num1 = 0;\n    int num2 = 0;\n scanf(\"%d%d\", &num1, &num2);printf(\"%d\",num1 + num2);\nreturn 0;\n}\n",
		//Language: "C",

		//PASS
		//Code:     "print sum([int(x) for x in input().split(\" \")])",
		//Language: "Python2",

		//PASS
		//Code:     "print(sum([int(x) for x in input().split(\" \")]),end=\"\")",
		//Language: "Python3",

		//PASSED
		//Code:     "package main\nimport \"fmt\"\nfunc main(){\nvar a,b int\nfmt.Scan(&a,&b)\nfmt.Print(a+b)}",
		//Language: "Golang",

		//PASSED
		//Code:     "using System;\nusing System.Linq;\n\nclass Program {\n    public static void Main(string[] args) {\n        Console.WriteLine(Console.ReadLine().Split().Select(int.Parse).Sum());\n    }\n}",
		//Language: "C#",

		//PASSED
		//Code:     "<?=array_sum(fscanf(STDIN, \"%d %d\"));",
		//Language: "PHP",

		//PASSED
		Code:     "var readline = require('readline');\nconst rl = readline.createInterface({\n        input: process.stdin,\n        output: process.stdout\n});\nrl.on('line', function(line){\n   var tokens = line.split(' ');\n    console.log(parseInt(tokens[0]) + parseInt(tokens[1]));\n});",
		Language: "JavaScript Node",
	}
	fmt.Println(judge.Judge(p, a))
}
```
