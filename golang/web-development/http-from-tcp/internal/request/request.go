package request

import (
	"bytes"
	"cmd/tcplistener/internal/headers"
	"errors"
	"fmt"
	"io"
	"strings"
)

type RequestState int
type Request struct {
	RequestLine Line
	Headers     headers.Headers
	State       RequestState
}

type Line struct {
	HttpVersion string
	Target      string
	Method      string
}

const (
	RequestStateInitialized RequestState = iota
	RequestStateDone
	RequestStateParsingHeaders
)
const crlf = "\r\n"
const bufferSize = 8

func FromReader(reader io.Reader) (*Request, error) {
	//initialize buffer to store data which will be processed
	buf := make([]byte, bufferSize)
	readToIndex := 0
	req := &Request{
		State: RequestStateInitialized,
	}
	for req.State != RequestStateDone {
		if readToIndex >= len(buf) {
			newBuf := make([]byte, len(buf)*2)
			copy(newBuf, buf)
			buf = newBuf
		}
		//Read writes the new bytes into the buffer
		numBytesRead, err := reader.Read(buf[readToIndex:])
		if err != nil {
			if errors.Is(err, io.EOF) {
				req.State = 1
				break
			}
			return nil, err
		}
		readToIndex += numBytesRead
		numBytesParsed, err := req.parse(buf[:readToIndex])
		if err != nil {
			return nil, err
		}

		//once it's parsed, it can be discarded from the buffer to save memory.
		copy(buf, buf[numBytesParsed:])
		readToIndex -= numBytesParsed
	}

	return req, nil
}

func (r *Request) parse(data []byte) (int, error) {
	if r.State == RequestStateDone {
		return 0, fmt.Errorf("error: trying to read data in a done state")
	}

	if r.State == RequestStateParsingHeaders {
	}

	requestLine, bytesRead, err := parseRequestLine(data)
	//fmt.Println(string(data))
	if err != nil {
		//Something went wrong
		return 0, err
	}
	if bytesRead == 0 {
		//Just need more data
		return 0, nil
	}
	r.RequestLine = *requestLine
	r.State = RequestStateParsingHeaders
	return bytesRead, nil
}

// At the end of the day, I must pass the complete line to this method.
func parseRequestLine(data []byte) (*Line, int, error) {
	idx := bytes.Index(data, []byte(crlf))
	if idx == -1 {
		return nil, 0, nil
	}

	requestLineStr := string(data[:idx])
	requestLine, err := requestLineFromString(requestLineStr)
	if err != nil {
		return nil, 0, err
	}

	return requestLine, idx + len(crlf), nil
}

func requestLineFromString(str string) (*Line, error) {
	requestLineParts := strings.Split(str, " ")
	if len(requestLineParts) != 3 {
		return nil, fmt.Errorf("invalid number of parts in request line")
	}

	method := requestLineParts[0]
	if strings.ToUpper(method) != method {
		return nil, fmt.Errorf("method must only contains capital alphabetic characters")
	}

	requestTarget := requestLineParts[1]

	versionParts := strings.Split(requestLineParts[2], "/")
	if len(versionParts) != 2 {
		return nil, fmt.Errorf("malformed start-line: %s", str)
	}

	httpPart := versionParts[0]
	if httpPart != "HTTP" {
		return nil, fmt.Errorf("unrecognized HTTP-version: %s", httpPart)
	}

	httpVersion := versionParts[1]
	if httpVersion != "1.1" {
		return nil, fmt.Errorf("unrecognized HTTP-version: %s", httpVersion)
	}

	return &Line{
		HttpVersion: httpVersion,
		Target:      requestTarget,
		Method:      method,
	}, nil
}
