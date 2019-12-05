package utils

import (
	"github.com/kataras/golog"
	"github.com/spf13/cast"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func HttpGet(url string, params map[string]interface{}) string {
	c := &http.Client{
		Timeout: 5 * time.Second,//5s超时
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		golog.Error("[http]method:HttpGet, NewRequest err: ", err)
		return ""
	}
	q := req.URL.Query()
	for k, v := range params {
		q.Add(k, cast.ToString(v))
	}
	req.URL.RawQuery = q.Encode()
	golog.Info("[http]method:HttpGet, http url: ", req.URL.String())
	resp, err := c.Do(req)
	if err != nil {
		golog.Error("[http]method:HttpGet, err: ", err)
		return ""
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		golog.Error("[http]method:HttpGet, body err: ", err)
		return ""
	}
	return string(body)
}

func HttpPost(url1 string, params map[string]interface{}) string {
	contentType := "application/x-www-form-urlencoded"
	q := make(url.Values)

	for k, v := range params {
		q[k] = []string{cast.ToString(v)}
	}

	c := &http.Client{
		Timeout: 10 * time.Second, //10s超时
	}
	golog.Info("[http]method:HttpPost, http url: ", url1, strings.NewReader(q.Encode()))

	req, err := http.NewRequest("POST", url1, strings.NewReader(q.Encode()))
	if err != nil {
		golog.Error("[http]method:HttpPost, err: ", err)
		return ""
	}
	req.Header.Set("Content-Type", contentType)
	resp, err := c.Do(req)

	if err != nil {
		golog.Error("[http]method:HttpPost, c.Do err: ", err)
		return ""
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		golog.Error("[http]method:HttpPost, body err: ", err)
		return ""
	}
	return string(body)
}