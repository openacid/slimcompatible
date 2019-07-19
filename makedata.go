package main

import (
	"encoding/binary"
	"flag"
	"fmt"
	"os"

	"github.com/golang/protobuf/proto"
	"github.com/openacid/slim/trie"
	"github.com/openacid/slimcompatible/testkeys"
)

type CmdFlag struct {
	Ver string
}

func main() {
	f := &CmdFlag{}
	flag.StringVar(&f.Ver, "ver", "1.0.0", "version in fn of slimtrie")
	flag.Parse()

	for typ, ks := range testkeys.Keys {
		prf := "slimtrie-data-" + typ + "-"
		makeMarshaledData(prf+"%s", f.Ver, ks)
	}
}

type I32 struct{}

func (c I32) Encode(d interface{}) []byte {
	b := make([]byte, 4)
	v := uint32(d.(int32))
	binary.LittleEndian.PutUint32(b, v)
	return b
}

func (c I32) Decode(b []byte) (int, interface{}) {

	size := int(4)
	s := b[:size]

	d := int32(binary.LittleEndian.Uint32(s))
	return size, d
}

func (c I32) GetSize(d interface{}) int {
	return 4
}

func (c I32) GetEncodedSize(b []byte) int {
	return 4
}

func makeMarshaledData(fn string, defaultVer string, keys []string) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("failed:", fn, defaultVer, r)
		}
	}()
	makeData(fn, defaultVer, keys)
}

func makeData(fn string, defaultVer string, keys []string) {

	type gv interface {
		GetVersion() string
	}

	n := len(keys)
	values := make([]int32, n)
	for i := 0; i < n; i++ {
		values[i] = int32(i)
	}

	st, err := trie.NewSlimTrie(I32{}, keys, values)
	if err != nil {
		panic(err)
	}

	b, err := proto.Marshal(st)
	if err != nil {
		panic(err)
	}

	var ver string
	ii := interface{}(st)
	vst, ok := ii.(gv)
	if ok {
		ver = vst.GetVersion()
	} else {
		ver = defaultVer
	}

	fn = fmt.Sprintf(fn, ver)
	f := newFile(fn)
	defer f.Close()

	_, err = f.Write(b)
	if err != nil {
		panic(err)
	}
}

func newFile(fn string) *os.File {
	f, err := os.OpenFile(fn, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	}

	err = f.Truncate(0)
	if err != nil {
		panic(err)
	}
	return f
}
