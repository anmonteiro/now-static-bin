package main

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/cenkalti/backoff"
)

type Request struct {
	Host     string            `json:"host"`
	Path     string            `json:"path"`
	Method   string            `json:"method"`
	Headers  map[string]string `json:"headers"`
	Encoding string            `json:"encoding,omitempty"`
	Body     string            `json:"body"`
}

type ResponseErrorWrapper struct {
	Error ResponseError `json:"error"`
}

type ResponseError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

var cmdName = ""
var cmd *exec.Cmd
var cleanup func()

func init() {
	var err error
	ex, _ := os.Executable()
	scriptPath := os.Getenv("NOW_STATIC_BIN_LOCATION")
	cmdName = path.Join(filepath.Dir(ex), scriptPath)
	cmd = exec.Command(cmdName)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err = cmd.Start(); err != nil {
		fmt.Fprintln(os.Stderr, "There was an error starting your program: ", err)
		os.Exit(1)
	}
	cleanup = func() {
		cmd.Process.Kill()
	}
}

func createErrorResponse(message string, code string, statusCode int) (events.APIGatewayProxyResponse, error) {
	obj := ResponseErrorWrapper{
		Error: ResponseError{
			Code:    code,
			Message: message,
		},
	}

	body, _ := json.Marshal(obj)

	return events.APIGatewayProxyResponse{
		StatusCode: statusCode,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Body: string(body),
	}, nil
}

func handler(ctx context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var req Request
	json.Unmarshal([]byte(event.Body), &req)

	// Should always be present, the builders sets it.
	port := os.Getenv("NOW_STATIC_BIN_PORT")
	url := "http://127.0.0.1:" + port + req.Path

	var body string
	if req.Encoding == "base64" {
		decoded, _ := base64.StdEncoding.DecodeString(req.Body)
		body = string(decoded)
	} else {
		body = string(req.Body)
	}

	bodyReader := strings.NewReader(body)
	localHttpReq, err := http.NewRequest(req.Method, url, bodyReader)
	if err != nil {
		fmt.Println(err)
		return createErrorResponse("Bad gateway internal req failed", "bad_gateway", 502)
	}

	for k, v := range req.Headers {
		localHttpReq.Header.Add(k, v)
		if strings.ToLower(k) == "host" {
			localHttpReq.Host = v
		}
	}
	for k, v := range req.Headers {
		localHttpReq.Header.Add(k, v)
		switch strings.ToLower(k) {
		case "host":
			// we need to set `Host` in the request
			// because Go likes to ignore the `Host` header
			// see https://github.com/golang/go/issues/7682
			localHttpReq.Host = v
		case "content-length":
			contentLength, _ := strconv.ParseInt(v, 10, 64)
			localHttpReq.ContentLength = contentLength
		case "x-forwarded-for":
		case "x-real-ip":
			localHttpReq.RemoteAddr = v
		}
	}
	// Taken from https://github.com/zeit/now-builders/pull/67
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			// we don't want to follow any redirects. that's something the actual client
			// (i.e., the end user) should do. this internal request is just an implementation detail
			// that acts like a proxy, so it shouldn't do things that are supposed to be done by the client
			return http.ErrUseLastResponse
		},
	}

	var internalRes *http.Response
	// Should always be present, the builders sets it.
	timeoutMillis, _ := strconv.Atoi(os.Getenv("NOW_STATIC_BIN_TIMEOUT"))
	backoffStrategy := backoff.WithMaxRetries(backoff.NewConstantBackOff(time.Duration(timeoutMillis)*time.Millisecond), 5)
	retryFunc := func() error {
		internalRes, err = client.Do(localHttpReq)
		return err
	}
	err = backoff.Retry(retryFunc, backoffStrategy)
	if err != nil {
		fmt.Println(err)
		return createErrorResponse("Bad gateway couldn't reach internal binary", "bad_gateway", 502)
	}

	defer internalRes.Body.Close()

	resHeaders := make(map[string]string, len(internalRes.Header))
	for k, v := range internalRes.Header {
		// FIXME: support multiple values via concatenating with ','
		// see RFC 7230, section 3.2.2
		resHeaders[k] = v[0]
	}

	bodyBytes, err := ioutil.ReadAll(internalRes.Body)
	if err != nil {
		return createErrorResponse("Bad gateway ReadAll bytes from response failed", "bad_gateway", 502)
	}

	return events.APIGatewayProxyResponse{
		StatusCode: internalRes.StatusCode,
		Headers:    resHeaders,
		Body:       string(bodyBytes)}, nil
}

func main() {
	lambda.Start(handler)
}
