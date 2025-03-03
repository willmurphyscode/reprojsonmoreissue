package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"maps"
	"os"

	"github.com/CalebQ42/squashfs"
)

func main() {
	// make a squashfs file from the test data
	fh, err := os.Open("testdata/test.squashfs")
	if err != nil {
		panic(err)
	}
	sqfs, err := squashfs.NewReader(fh)
	if err != nil {
		panic(err)
	}
	r, err := sqfs.Open("package.json")
	if err != nil {
		panic(err)
	}
	data := make(map[string]interface{})
	var decodedItems []map[string]interface{}
	decoder := json.NewDecoder(r)
	giveUpAfter := 100
	iterations := 0
	for decoder.More() {
		iterations++
		if iterations > giveUpAfter {
			break
		}
		err = decoder.Decode(&data)
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			panic(err)
		}
		decodedItems = append(decodedItems, maps.Clone(data))
	}
	fmt.Printf("Decoded %d items in %d iterations\n", len(decodedItems), iterations)
	if len(decodedItems) > 0 {
		fmt.Printf("First item: %v\n", decodedItems[0])
	}
}
