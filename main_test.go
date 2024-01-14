package main

import (
	"os"
	"strings"
	"testing"
)

// go test *.go -v

func TestTetris(t *testing.T) {
	successDots := map[string]int{"0": 0, "1": 9, "2": 4, "3": 5, "hard": 1}
	files, err := os.ReadDir("./examples")
	if err != nil {
		t.Log(err)
	}

	for _, file := range files {
		result, err := Start("./examples/" + file.Name())
		if strings.Contains(file.Name(), "bad") {
			if err == nil {
				t.Errorf("file:%s\noutput:\n%s\nwant:\nERROR", file.Name(), result)
			}
		} else {
			for key, value := range successDots {
				if strings.Contains(file.Name(), key) {
					dotsCount := strings.Count(result, ".")
					t.Log(file.Name())
					t.Log("\n" + result)
					if dotsCount != value {
						t.Errorf("file:%s\noutput:\n%d\nwant:\n%d", file.Name(), dotsCount, value)
					}
				}
			}
		}
	}
}
