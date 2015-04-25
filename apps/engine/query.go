package main

type Query struct {
	Name       string
	Text       string
	Parameters []*QueryParameter
}

type QueryParameter struct {
	Name string
}
