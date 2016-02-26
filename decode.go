package bencode

import (
	"bytes"
	"io/ioutil"
	"strconv"
)

type Decoder struct {
	length    int
	filebytes []byte
	reader    *bytes.Reader
}

func NewDecoder(dst string) Decoder {
	filebytes, err := ioutil.ReadFile(dst)
	if err != nil {
		panic(err)
	}

	dec := Decoder{}
	dec.reader = bytes.NewReader(filebytes)
	dec.length = len(filebytes)
	dec.filebytes = filebytes

	return dec
}

func (dec Decoder) Decode() interface{} {

	b, _ := dec.reader.ReadByte()

	switch b {
	case 'd':
		next, err := dec.reader.ReadByte()
		if err != nil {
			panic(err)
		}
		dict := make(map[interface{}]interface{})
		for next != 'e' {
			dec.reader.UnreadByte()
			key := dec.Decode()
			val := dec.Decode()
			dict[key] = val
			next, err = dec.reader.ReadByte()
			if err != nil {
				panic(err)
			}
		}
		return dict
	case 'l':
		next, err := dec.reader.ReadByte()
		if err != nil {
			panic(err)
		}
		list := []interface{}{}
		for next != 'e' {
			dec.reader.UnreadByte()
			list = append(list, dec.Decode())
			next, err = dec.reader.ReadByte()
			if err != nil {
				panic(err)
			}
		}
		return list
	case 'i':
		numBytes := dec.readBytesTill('e')
		data, err := strconv.ParseInt(string(numBytes), 10, 64)
		if err != nil {
			data, err := strconv.ParseUint(string(numBytes), 10, 64)
			if err != nil {
				panic(err)
			}
			return data
		}
		return data
	default:
		dec.reader.UnreadByte()
		lengthBytes := dec.readBytesTill(':')
		length, err := strconv.ParseInt(string(lengthBytes), 10, 64)
		if err != nil {
			panic(err)
		}

		stringbytes := make([]byte, length)
		dec.reader.Read(stringbytes)
		data := string(stringbytes)
		return data
	}
}

func (dec Decoder) readBytesTill(delim byte) []byte {
	i := bytes.IndexByte(dec.filebytes[dec.length-dec.reader.Len():], delim)
	byteslice := make([]byte, i)
	dec.reader.Read(byteslice)
	dec.reader.Seek(1, 1)
	return byteslice
}
