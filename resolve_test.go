package dots

import (
	"log"
	"testing"
)

func TestResolve(t *testing.T) {
	result, err := ResolveFiles([]string{"fixtures/..."}, []string{"fixtures/foo"})

	files := map[string]bool{
		"fixtures/bar/bar1.go": true,
		"fixtures/bar/bar2.go": true,
		"fixtures/baz/baz1.go": true,
		"fixtures/baz/baz2.go": true,
		"fixtures/baz/baz3.go": true,
	}

	if err != nil {
		t.Error("Got an error: " + err.Error())
	}

	if len(result) != len(files) {
		t.Error("Matched different number of files")
	}

	if err != nil {
		log.Fatal(err)
	}

	for _, r := range result {
		if _, ok := files[r]; !ok {
			t.Error("Not supposed to match" + r)
		}
	}
}