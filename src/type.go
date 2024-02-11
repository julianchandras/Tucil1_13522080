package main

type matrix struct {
	Width  int
	Height int
	Buffer [][]string
}

type sequence struct {
	Tokens []string
	Reward int
}

type listOfSequence struct {
	Buffer []sequence
	Neff   int
}

type point struct {
	Column int
	Row    int
}

type path struct {
	Coordinates []point
	Neff        int
}

type listOfPath struct {
	Buffer []path
	Neff   int
}
