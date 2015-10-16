package main

import(
        "os"
	"io"
)

type Msg string

func (m Msg) WriteTo(w io.Writer) (int, error){
	return w.Write([]byte(m))
}

func ( m *Msg) Reset(){
	*m = ""
}

func main() {
	m := Msg("Hello World!\n")
	m.WriteTo(os.Stdout)
	m.Reset()
	m.WriteTo(os.Stdout)
}

