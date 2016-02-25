package bencode

import (
	"bytes"
	"io"
	"io/ioutil"
	"strconv"
)

type Decoder struct {
	Data interface{}
}

func NewDecoder(dst string) Decoder {
	filebytes, err := ioutil.ReadFile(dst)
	if err != nil {
		panic(err)
	}

	dec := Decoder{}
	data := dec.Decode(filebytes, 0)
	dec.Data = data
	return dec
}

func (dec Decoder) Decode(filebytes []byte, pos int64) interface{}, int {
	reader := bytes.NewReader(filebytes)

	pos, err := reader.Seek(pos, 0)
	if err != nil {
		panic(err)
	}
	b, _ := reader.ReadByte()
	switch b {
		case 'd':
		case 'i':
			i := pos + int64(bytes.IndexByte(filebytes[pos:], 'e'))
			data, err := strconv.ParseInt(string(filebytes[pos + 1:i]), 10, 64)
			if err != nil {
				data, err := strconv.ParseUint(string(filebytes[pos + 1:i]), 10, 64)
				if err != nil {
					panic(err)
				}
				return data
			}
			return data, XXXXX
		case 'l':
			data := make([]interface{})
			next, _ := reader.ReadByte()
			for next != 'e' {
				next, _ = reader.ReadByte()
			}
		default:
			i := pos + int64(bytes.IndexByte(filebytes[pos:], ':'))
			length, err := strconv.ParseInt(string(filebytes[pos:i]), 10, 64)
			if err != nil {
				panic(err)
			}
			pos, err = reader.Seek(i + 1, 0)
			if err != nil {
				panic(err)
			}
			stringbytes := make([]byte, length)
			reader.Read(stringbytes)
			data := string(stringbytes)
			return data, pos
	}

	return "null", pos
}
