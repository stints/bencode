package bencode

import (
	"os"
	"reflect"
	"strconv"
)

type Encoder struct {
	data []byte
}

func NewEncoder() Encoder {
	enc := Encoder{}
	return enc
}

func (enc Encoder) Write(dst string) {
	f, err := os.Create(dst)
	if err != nil {
		panic(err)
	}
	_, err = f.Write(enc.data)
	if err != nil {
		panic(err)
	}
}

func (enc Encoder) Encode(obj interface{}) []byte {
	v := reflect.ValueOf(obj)
	switch v.Kind() {
	case reflect.Map:
		mapbytes := []byte{'d'}
		for _, key := range v.MapKeys() {
			keyname := enc.Encode(key.String())
			keyval := enc.Encode(v.MapIndex(key).Interface())
			mapbytes = append(mapbytes, append(keyname, keyval...)...)
		}
		mapbytes = append(mapbytes, 'e')
		enc.data = mapbytes
		return mapbytes
	case reflect.Slice, reflect.Array:
		listbytes := []byte{'l'}
		for i := 0; i < v.Len(); i++ {
			listval := enc.Encode(v.Index(i).Interface())
			listbytes = append(listbytes, listval...)
		}
		listbytes = append(listbytes, 'e')
		enc.data = listbytes
		return listbytes
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		intbytes := []byte{'i'}
		intbytes = strconv.AppendInt(intbytes, v.Int(), 10)
		intbytes = append(intbytes, 'e')
		enc.data = intbytes
		return intbytes
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		uintbytes := []byte{'i'}
		uintbytes = strconv.AppendUint(uintbytes, v.Uint(), 10)
		uintbytes = append(uintbytes, 'e')
		enc.data = uintbytes
		return uintbytes
	case reflect.String:
		length := enc.Encode(v.Len())
		bytelength := length[1 : len(length)-1]
		stringbytes := []byte{}
		stringbytes = append(stringbytes, bytelength...)
		stringbytes = append(stringbytes, ':')
		stringbytes = append(stringbytes, []byte(v.String())...)
		enc.data = stringbytes
		return stringbytes
	default:
		panic("Type not supported")
	}
}
