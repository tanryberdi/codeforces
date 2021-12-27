package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strconv"
)

func nGroups(r io.Reader, w io.Writer) (int, error) {
	if _, ok := r.(*bufio.Reader); !ok {
		r = bufio.NewReader(r)
	}
	br := r.(*bufio.Reader)

	buf, err := br.ReadBytes('\n')
	if err != nil {
		return 0, err
	}
	n, err := strconv.Atoi(string(bytes.TrimSpace(buf)))
	if err != nil {
		return 0, err
	}

	ng := 0
	var prev byte
	for i := 0; i < n; i++ {
		m, err := br.ReadSlice('\n')
		if err != nil {
			if err != io.EOF || i < n-1 {
				return 0, err
			}
		}
		if m[0] != prev {
			ng++
		}
		prev = m[0]
	}
	return ng, nil
}

func main() {
	ng, err := nGroups(os.Stdin, os.Stdout)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(ng)
}
