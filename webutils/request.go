package webutils

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"sync"
	"time"
)

const (
	METHODNOTSUPPORT = "METHODNOTSUPPORT"
	URLINVALID       = "URLINVALID"
	REQUESTFAIL      = "REQUESTFAIL"
)

type Query struct {
	Method  string
	Url     string
	Headers map[string]string
	Data    map[string]string
	Tags    interface{}
}

type Feedback struct {
	Statuscode int
	Message    string
	Error      error
	Tags       interface{}
}

func doRequest(httpclient *http.Client, query Query) (feedback Feedback) {
	if query.Method == "GET" {
		var err error
		var req *http.Request
		var resp *http.Response
		req, err = http.NewRequest(query.Method, query.Url, nil)
		if err != nil {
			return Feedback{Error: err}
		}
		for k, v := range query.Headers {
			req.Header.Add(k, v)
		}
		resp, err = httpclient.Do(req)
		if err != nil {
			return Feedback{Error: errors.New(REQUESTFAIL)}
		}
		defer resp.Body.Close()
		statuscode := resp.StatusCode
		resultbyte, _ := ioutil.ReadAll(resp.Body)
		result := string(resultbyte)
		return Feedback{Statuscode: statuscode, Message: result, Tags: query.Tags}
	} else if query.Method == "POST" {
		var err error
		var req *http.Request
		var resp *http.Response
		postvalue := url.Values{}
		for k, v := range query.Data {
			postvalue.Set(k, v)
		}
		postdatastr := postvalue.Encode()
		postdatabytes := []byte(postdatastr)
		postbytesreader := bytes.NewReader(postdatabytes)
		if err != nil {
			return Feedback{Error: errors.New(URLINVALID)}
		}
		req, err = http.NewRequest(query.Method, query.Url, postbytesreader)
		if err != nil {
			return Feedback{Error: errors.New(URLINVALID)}
		}
		for k, v := range query.Headers {
			req.Header.Add(k, v)
		}
		resp, err = httpclient.Do(req)
		if err != nil {
			return Feedback{Error: errors.New(REQUESTFAIL)}
		}
		defer resp.Body.Close()
		statuscode := resp.StatusCode
		resultbyte, _ := ioutil.ReadAll(resp.Body)
		result := string(resultbyte)
		return Feedback{Statuscode: statuscode, Message: result, Tags: query.Tags}
	} else {
		return Feedback{Error: errors.New(METHODNOTSUPPORT)}
	}
}

func CaptureSingle(query Query, timeout time.Duration) (feedback Feedback) {
	httpclient := &http.Client{Timeout: timeout}
	feedback = doRequest(httpclient, query)
	return feedback
}

func CaptureMulti(query []Query, timeout time.Duration) (feedbacks []Feedback) {
	httpclient := &http.Client{Timeout: timeout}
	feedbackchan := make(chan Feedback)
	wg := sync.WaitGroup{}
	wg.Add(len(query))
	for _, q := range query {
		go func(httpclient *http.Client, query Query) {
			defer wg.Done()
			fmt.Println("send request")
			feedback := doRequest(httpclient, query)
			// fmt.Println(feedback)
			feedbackchan <- feedback
		}(httpclient, q)
	}
	go func() {
		for r := range feedbackchan {
			feedbacks = append(feedbacks, r)
		}
	}()
	wg.Wait()
	return feedbacks
}
