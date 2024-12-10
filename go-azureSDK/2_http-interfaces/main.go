package main

import (
	"fmt"
	"io"
	"log"
)

type MySlowReader struct {
	//attributes
	content string
	pos     int
}

// WRiting our own Streaming function

func (m *MySlowReader) Read(p []byte) (n int, err error) {
	if m.pos+1 <= len(m.content) {
		n := copy(p, m.content[m.pos:m.pos+1])
		m.pos++
		return n, nil
	}
	return 0, io.EOF
}

func main() {

	// instantiating and providing values to the struct
	MySlowReaderInstance := &MySlowReader{
		content: "Hello World!",
	}

	body, err := io.ReadAll(MySlowReaderInstance)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Body/output: %s\n", body)
}
