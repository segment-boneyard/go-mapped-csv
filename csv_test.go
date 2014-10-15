package csv

import "github.com/bmizerany/assert"
import "testing"
import "bytes"

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func Test(t *testing.T) {
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

	assert.Equal(t, exp, string(b.Bytes()))
}
