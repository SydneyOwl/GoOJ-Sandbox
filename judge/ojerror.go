/**
 * @Author: SydneyOwl
 * @Description:
 * @File: Exceptions
 * @Date: 2022/11/1 20:18
 */

package judge

import (
	"fmt"
	"github.com/SydneyOwl/GoOJ-Sandbox/utils"
	"github.com/pkg/errors"
	"reflect"
)

type OJError struct {
	Code   ErrCode
	Msg    string
	Stdout string
	Stderr string
}

// Error returns full err msg.
func (err *OJError) Error() string {
	return utils.MergeNonEmptyStr(err.Msg, err.Stdout, err.Stderr)
}
func (err *OJError) GetType() string {
	return err.Code.String()
}
func (err *OJError) Is(err1 error) bool {
	statusCode := reflect.ValueOf(err1).Elem().FieldByName("Code").Int()
	name := reflect.TypeOf(err1).Elem().Name()
	//a := reflect.ValueOf(errUp).FieldByName("Code").String()
	//log.Println(a)
	//if errUp == nil {
	//	return false
	//}
	//log.Println(reflect.TypeOf(err1).Elem().Name())
	if name != "OJError" {
		return false
	}
	if statusCode == int64(err.Code) {
		return true
	}
	return false
}

type ErrCode int

//go:generate stringer -type ErrCode -linecomment
const (
	OJServerError ErrCode = 10001

	CompileError             ErrCode = 20001
	SubmitError              ErrCode = 20002
	UnsupportedLanguageError ErrCode = 20003
)

// Deprecated: NewDetailedOJError creates an FULL_OJError with wrap option.
func NewDetailedOJError(code ErrCode, msg string, stdout string, stderr string) error {
	return errors.Wrap(&OJError{
		Code:   code,
		Msg:    msg,
		Stdout: stdout,
		Stderr: stderr,
	}, "")
}

// NewOJError creates an OJError with wrap option.
func NewOJError(code ErrCode) error {
	return errors.Wrap(&OJError{
		Code: code,
		Msg:  code.String(),
	}, "")
}

func NewFullOJError(code ErrCode, msg string, stdout string, stderr string) error {
	return errors.Wrap(&OJError{
		Code:   code,
		Msg:    msg,
		Stdout: stdout,
		Stderr: stderr,
	}, "")
}
func NewRawOJError(code ErrCode) error {
	return &OJError{
		Code:   code,
		Msg:    "",
		Stdout: "",
		Stderr: "",
	}
}

// StackTraceOJError gives all message and stack info if there's.
func StackTraceOJError(e error) string {
	return fmt.Sprintf("OJError:%+v", e)
}
