package main

import (
	"bufio"
	"fmt"
	"slices"
	"strings"
)

type request struct {
	method  string
	path    string
	version string
	headers map[string][]string
	body    []byte
}

func parseRequest(reader *bufio.Reader) (*request, error) {
	var req request

	line, err := reader.ReadString('\n')
	if err != nil {
		return nil, fmt.Errorf("reading request line: %w", err)
	}

	if !strings.HasSuffix(line, "\r\n") {
		return nil, fmt.Errorf("missing CRLF: %q", line)
	}

	line = strings.TrimSuffix(line, "\r\n")
	splitLine := strings.Split(line, " ")

	if len(splitLine) != 3 {
		return nil, fmt.Errorf("incorrect number of parts: %q", line)
	}

	if slices.Contains(splitLine, "") {
		return nil, fmt.Errorf("incorrect spacing between request values: %q", line)
	}

	if splitLine[2] != "HTTP/1.1" {
		return nil, fmt.Errorf("http version %q not accepted", splitLine[2])
	}

	req.method = splitLine[0]
	req.path = splitLine[1]
	req.version = splitLine[2]

	return &req, nil
}
