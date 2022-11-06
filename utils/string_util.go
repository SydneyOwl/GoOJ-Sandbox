/**
 * @Author: SydneyOwl
 * @Description:
 * @File: string_util.go
 * @Date: 2022/11/2 16:39
 */

package utils

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func MergeNonEmptyStr(concat ...string) string {
	var str strings.Builder
	for _, v := range concat {
		if v != "" {
			if len(v) > 1024 {
				_, _ = fmt.Fprint(&str, v[0:1024])
				_, _ = fmt.Fprint(&str, "\n")
			} else {
				_, _ = fmt.Fprint(&str, v)
				_, _ = fmt.Fprint(&str, "\n")
			}
		}
	}
	return str.String()
}
func CompressStr(str string) string {
	if str == "" {
		return ""
	}
	reg := regexp.MustCompile("\\s+")
	return reg.ReplaceAllString(str, "")
}
func DelUnprintableWords(str string) string {
	target := []rune(str)
	var ans strings.Builder
	for _, v := range target {
		if strconv.IsPrint(v) {
			_, _ = fmt.Fprint(&ans, string(v))
		}
	}
	return ans.String()
}
