/**
 * @Author: SydneyOwl
 * @Description:
 * @File: judge
 * @Date: 2022/11/2 13:50
 */

package entity

type Judge struct {
	SubmitId       string
	ProblemId      string
	SubmitTime     string
	Status         int
	ErrorMessage   string
	TimeConsumed   int //ms
	MemoryConsumed int //k
	Score          int
	Code           string
	Language       string
}
