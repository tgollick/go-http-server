package main

type request struct {
	method  string
	path    string
	version string
	headers map[string][]string
	body    []byte
}
