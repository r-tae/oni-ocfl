package main

import (
	"bufio"
	"bytes"
	"github.com/richardlehane/siegfried"
	//"io/ioutil"
	"log"
	"os"
	"syscall/js"
)

var sf *siegfried.Siegfried

func identify(this js.Value, inputs []js.Value) interface{} {
	//data := js.Global().Get("Uint8Array").New(inputs[0])
	data := inputs[0]
	dst := make([]byte, data.Get("length").Int())
	dst2 := make([]byte, data.Get("length").Int())

	// This returns the number of bytes copied, its definitely working
	js.CopyBytesToGo(dst, data)
	js.CopyBytesToGo(dst2, data)

	bytes.NewReader(dst)
	r := bytes.NewReader([]byte{10, 20, 30})
	// out has a length, so the reader is working
	//out, _ := ioutil.ReadAll(r)

	//d_out := js.Global().Get("Uint8Array").New(10)
	// `return data.Get("length").Int()` gives a length, also working
	// `js.CopyBytesToJS(d_out, out)` gives a length, also working
	//js.CopyBytesToJS(d_out, out)

	filename := inputs[1]
	//siegfried identify(io.Reader, filename, mimeType),
	// can leave filename and mimetype blank if you don't have them
	ids, _ := sf.Identify(r, filename.String(), "")
	return 0
	return ids
}

func main() {
	quit := make(chan int)
	js.Global().Set("identify", js.FuncOf(identify))

	file, err := os.Open("./node_modules/oni-ocfl/default.sig")
	if err != nil {
		log.Fatal(err)
	}

	reader := bufio.NewReader(file)

	sf, err = siegfried.LoadReader(reader)
	if err != nil {
		log.Fatal(err)
	}

	<-quit
}
