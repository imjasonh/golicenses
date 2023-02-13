package golicenses

import (
	"bytes"
	"compress/gzip"
	_ "embed"
	"encoding/csv"
	"errors"
	"io"
	"runtime"
	"sync"
	"time"

	"github.com/dghubble/trie"
)

//go:embed licenses.csv.gz
var b []byte

var (
	t    *trie.RuneTrie
	once sync.Once

	// LoadTime is the time it took to load the dataset.
	// It is populated after the first call to Get.
	LoadTime time.Duration

	// NumRecords is the total number of records in the dataset.
	// It is populated after the first call to Get.
	NumRecords int

	Alloc uint64

	ErrNotFound = errors.New("not found")
)

// Get returns the reported license for the package.
//
// The first time Get is called, the dataset is loaded and parsed and stored in
// memory, populating LoadTime and NumRecords. Subsequent calls to Get read
// from memory.
func Get(p string) (string, error) {
	var lerr error

	once.Do(func() {
		var mem runtime.MemStats
		runtime.ReadMemStats(&mem)
		prev := mem.Alloc

		start := time.Now()
		t = trie.NewRuneTrie()
		gr, err := gzip.NewReader(bytes.NewReader(b))
		if err != nil {
			lerr = err
			return
		}
		r := csv.NewReader(gr)
		r.FieldsPerRecord = 2
		c := 0
		for {
			rec, err := r.Read()
			if err == io.EOF {
				break
			} else if err != nil {
				lerr = err
				return
			}
			t.Put(rec[0], rec[1])
			c++
		}

		LoadTime = time.Since(start)
		NumRecords = c

		runtime.ReadMemStats(&mem)
		Alloc = mem.Alloc - prev
	})
	if lerr != nil {
		return "", lerr
	}

	return t.Get(p).(string), nil
}
