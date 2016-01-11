package main

type Dict struct {
	name string
	url  string
	path string // html path for goquery
}

type Config struct {
	dicts []Dict
}
