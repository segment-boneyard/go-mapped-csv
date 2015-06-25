package mapped_csv

import "github.com/bmizerany/assert"
import "testing"
import "bytes"

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func TestWrite(t *testing.T) {
	b := bytes.NewBuffer(nil)

	w := New(b, []string{"first_name", "last_name", "species"})

	err := w.Write(map[string]string{
		"first_name": "Tobi",
		"last_name":  "Ferret",
		"species":    "ferret",
		"more":       "stuff",
	})

	check(err)

	err = w.Write(map[string]string{
		"first_name": "Manny",
		"last_name":  "Cat",
		"species":    "just a cat",
		"more":       "stuff",
	})

	check(err)

	w.Flush()

	exp := `Tobi,Ferret,ferret
Manny,Cat,just a cat
`

	assert.Equal(t, exp, b.String())
}

func TestWriteHeader(t *testing.T) {
	b := bytes.NewBuffer(nil)

	w := New(b, []string{"first_name", "last_name", "species"})

	err := w.WriteHeader()
	check(err)

	err = w.Write(map[string]string{
		"first_name": "Tobi",
		"last_name":  "Ferret",
		"species":    "ferret",
		"more":       "stuff",
	})

	check(err)

	err = w.Write(map[string]string{
		"first_name": "Manny",
		"last_name":  "Cat",
		"species":    "just a cat",
		"more":       "stuff",
	})

	check(err)

	w.Flush()

	exp := `first_name,last_name,species
Tobi,Ferret,ferret
Manny,Cat,just a cat
`

	assert.Equal(t, exp, b.String())
}

func Benchmark(b *testing.B) {
	buf := bytes.NewBuffer(nil)

	w := New(buf, []string{"first_name", "last_name", "species"})

	err := w.WriteHeader()
	check(err)

	for i := 0; i < b.N; i++ {
		err = w.Write(map[string]string{
			"first_name": "Tobi",
			"last_name":  "Ferret",
			"species":    "ferret",
			"more":       "stuff",
		})

		check(err)
	}
}
