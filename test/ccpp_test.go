/**
 * @Author: SydneyOwl
 * @Description: Unit test
 * @File: test_ccpp
 * @Date: 2022/11/2 10:42
 */

package test

import (
	"fmt"
	"github.com/SydneyOwl/GoOJ-Sandbox/judge"
	"github.com/SydneyOwl/GoOJ-Sandbox/utils"
	"testing"
)

// TestCompile tests if compile module works.
func TestCompile(t *testing.T) {
	t.Run("judge.compile", func(t *testing.T) {
		if a, err := judge.Compile(utils.CCPP, "#include <iostream>\nusing namespace std;\nint main() {\nint a, b;\ncin >> a >> b;\ncout << a + b << endl;}", "c++"); err != nil {
			t.Error("judge.compile failed!")
		} else {
			fmt.Println(a)
		}
	})
	t.Run("http.sendCR", func(t *testing.T) {
		a := judge.SendCompileReq(10000, 30000, 104811111, 134217728, "a.cc", "a", []string{"/usr/bin/g++", "a.cc", "-o", "a"}, []string{"PATH=/usr/bin:/bin"}, "#include <iostream>\nusing namespace std;\nint main() {\nint a, b;\ncin >> a >> b;\ncout << a + b << endl;\n}")
		if a.Raw == "" {
			t.Error("http.sendCR failed")
		}

	})
}
