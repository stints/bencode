# bencode
Bencoding library for go.

## encoding
```
import "github.com/stints/bencode"

mapObj = make(map[string]interface{})
mapObj["key1"] = "value1"
mapObj["key2"] = []int{1,2,3}
mapObj["key3"] = 234523

enc = bencode.NewEncoder()
encData = enc.Encode(mapObj) // encodes to []byte -> "d4:key16:value14:key2li1ei2ei3ee4:key3i234523ee"

\\ write current encode to local file
enc.Write("file.txt")
```

## decoding
```
import github.com/stints/bencode"

// Creates a new decoder, reads data from file into []byte
dec = bencode.NewDecoder("file.txt")
decObj = dec.Decode().(map[interface{}]interface{}) // assertion required as any type can be a key

// decObj["key1"] -> "value1"
```
