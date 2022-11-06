/**
 * @Author: SydneyOwl
 * @Description: Send HTTP requests
 * @File: http_client.go
 * @Date: 2022/11/1 15:26
 */

package utils

import (
	"bytes"
	"encoding/json"
	"github.com/pkg/errors"
	"net/http"
)

// SendSimpleDeleteRequest sends reqDELETE
func SendSimpleDeleteRequest(url string) (resp http.Response, err error) {
	req, _ := http.NewRequest("DELETE", url, nil)
	res, err := http.DefaultClient.Do(req)
	defer res.Body.Close()
	if err != nil {
		return http.Response{}, errors.Wrap(err, "Send request failed!")
	}
	return *res, nil
}

// SendJsonPostRequest sends reqJSON
func SendJsonPostRequest(url string, param map[string]interface{}) (resp http.Response, err error) {
	data, err := json.Marshal(param)
	if err != nil {
		return http.Response{}, err
	}
	reader := bytes.NewReader(data)
	request, _ := http.NewRequest("POST", url, reader)
	request.Header.Set("Content-Type", "application/json;charset=UTF-8")
	res, err := http.DefaultClient.Do(request)
	if err != nil {
		return http.Response{}, errors.Wrap(err, "Send JP request failed!")
	}
	return *res, nil
}
