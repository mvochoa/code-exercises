package main

import (
	"bytes"
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

func TestMain(t *testing.T) {
	ins, _ := ioutil.ReadDir("input")
	for _, f := range ins {
		nTest := strings.Replace(f.Name(), "input", "", -1)
		content, _ := ioutil.ReadFile("input/input" + nTest)

		in, _ := ioutil.TempFile("./", "in")
		in.Write(content)
		in.Seek(0, 0)
		os.Stdin = in

		out, _ := ioutil.TempFile("./", "out")
		os.Stdout = out

		main()

		f1, _ := ioutil.ReadFile(out.Name())
		f2, _ := ioutil.ReadFile("output/output" + nTest)

		if !bytes.Equal(f1, f2) {
			t.Errorf("FALLO PRUEBA: %s, \nSE OBTUVO: \n%s----- \nSE ESPERABA: \n%v",
				strings.Split(nTest, ".")[0],
				string(f1),
				string(f2))
		}

		defer os.Remove(in.Name())
		defer os.Remove(out.Name())
	}
}
