package tools

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"strings"
	"time"
)

/**
@description
@date: 03/13 13:17
@author Gk
**/

var cli *http.Client

const (
	httpTimeout     = 60 * time.Second
	dialTimeout     = 60 * time.Second
	aliveTimeout    = 60 * time.Second
	shakeTimeout    = 60 * time.Second
	reqPrintLength  = 2000
	respPrintLength = 2000
)

type Logger interface {
	Printf(format string, a ...interface{})
}

type PrintReqMiddleware struct {
	Logger Logger
	Proxy  *http.Transport
}

func (lrt *PrintReqMiddleware) RoundTrip(req *http.Request) (*http.Response, error) {
	resp, e := lrt.Proxy.RoundTrip(req)
	//PrintCurl(req, resp, e, lrt.Logger)
	return resp, e
}

func PrintCurl(req *http.Request, resp *http.Response, respErr error, logger Logger) {

	var builder strings.Builder
	builder.WriteString(fmt.Sprintf("curl -v -X %s '%s' ", req.Method, req.URL))
	i := 0
	for k, v := range req.Header {
		value := ""
		i++
		for _, vone := range v {
			value += vone
		}
		builder.WriteString(fmt.Sprintf("\\\n-H '%s: %s' ", k, value))
	}
	if req.Body != nil {
		builder.WriteString(fmt.Sprintf("\\\n"))
		builder.WriteString("-d $'")
		body, err := req.GetBody()
		if err != nil {
			logger.Printf("get body error %v", err)
		}
		buf := make([]byte, reqPrintLength)
		n, err := body.Read(buf)
		if err != nil {
			if err != io.EOF {
				fmt.Println("read error:", err)
			}
		}
		builder.WriteString(string(buf[:n]))
		if err != io.EOF && n >= reqPrintLength {
			builder.WriteString(fmt.Sprintf("......large bigger than %d }", reqPrintLength))
		}
		builder.WriteString("'")
	}
	//logger.Printf("[CURL] request:\n%s \n ", builder.String())

	var respBuilder strings.Builder
	respBuilder.WriteString(fmt.Sprintf("url: %s\n", req.URL))
	if respErr != nil {
		respBuilder.WriteString(fmt.Sprintf("error: %v\n", respErr))
	}
	respBuilder.WriteString(fmt.Sprintf("resp: "))
	if resp == nil {
		respBuilder.WriteString("nil")
	} else {
		respBuilder.WriteString(fmt.Sprintf("\n    status: %d", resp.StatusCode))
		respBuilder.WriteString("\n    header:")
		for k, v := range resp.Header {
			value := ""
			i++
			for _, vone := range v {
				value += vone
			}
			respBuilder.WriteString(fmt.Sprintf("\n        %s: %s", k, value))
		}
		if resp.ContentLength > respPrintLength {
			respBuilder.WriteString(fmt.Sprintf("\n    body: response too large bigger than %d", respPrintLength))
		} else {
			var b bytes.Buffer
			copyN, err := io.CopyN(bufio.NewWriter(&b), resp.Body, respPrintLength)
			if err != nil {
				if err != io.EOF {
					logger.Printf("copy resp body error %v", err)
				}
			}
			if err != io.EOF && copyN >= respPrintLength {
				respBuilder.WriteString(fmt.Sprintf("\n    body: %s", b.String()+"......"))
			} else {
				respBuilder.WriteString(fmt.Sprintf("\n    body: %s", b.String()))
			}
			resp.Body = ioutil.NopCloser(io.MultiReader(bufio.NewReader(&b), resp.Body))
		}
		respBuilder.WriteString("\n")
	}
	logger.Printf("[CURL] response:\n%s \n ", respBuilder.String())
}

func init() {
	//var err error

	cli = &http.Client{
		Timeout: 900 * time.Second,
		Transport: &PrintReqMiddleware{
			Logger: LogrusLogger,
			Proxy: &http.Transport{
				DialContext: (&net.Dialer{
					Timeout:   dialTimeout,
					KeepAlive: aliveTimeout,
				}).DialContext,
				TLSHandshakeTimeout: shakeTimeout,
			},
		},
	}
}

func HTTPGet(url string, headers map[string]string) (resp *http.Response, body []byte, err error) {
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, nil, fmt.Errorf("HTTP: create request failed: %w", err)
	}

	for k, v := range headers {
		request.Header.Set(k, v)
	}

	resp, err = cli.Do(request)
	if err != nil {
		return resp, nil, fmt.Errorf("HTTP: do request failed: %w", err)
	}

	httpResp, err := dealHttpResp(resp)
	return resp, httpResp, err
}

func HTTPPost(url string, headers map[string]string, data []byte) (resp *http.Response, body []byte, err error) {
	request, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(data))
	if err != nil {
		return nil, nil, fmt.Errorf("HTTP: create request failed: %w", err)
	}

	for k, v := range headers {
		request.Header.Set(k, v)
	}

	resp, err = cli.Do(request)
	if err != nil {
		return resp, nil, fmt.Errorf("HTTP: do request failed: %w", err)
	}

	httpResp, err := dealHttpResp(resp)
	return resp, httpResp, err
}

func HTTPPatch(url string, headers map[string]string, data []byte) (resp *http.Response, body []byte, err error) {
	request, err := http.NewRequest(http.MethodPatch, url, bytes.NewBuffer(data))
	if err != nil {
		return nil, nil, fmt.Errorf("HTTP: create request failed: %w", err)
	}

	for k, v := range headers {
		request.Header.Set(k, v)
	}

	resp, err = cli.Do(request)
	if err != nil {
		return resp, nil, fmt.Errorf("HTTP: do request failed: %w", err)
	}

	httpResp, err := dealHttpResp(resp)
	return resp, httpResp, err
}

func dealHttpResp(resp *http.Response) (body []byte, err error) {
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("HTTP: read body failed: %w", err)
	}

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("HTTP: %v, msg=%v", resp.StatusCode, string(body))
	}

	return body, nil
}
