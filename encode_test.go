package bencode

import (
	"testing"
)

func TestMapEncode(t *testing.T) {
	encoder := NewEncoder()

	mapObj := make(map[string]string)
	mapObj["key1"] = "value1"

	mapObjEnc := encoder.Encode(mapObj)

	if string(mapObjEnc) !=  "d4:key16:value1e" {
		t.Error("Map encoded value incorrect")
	}
}

func TestListEncode(t *testing.T) {
	encoder := NewEncoder()

	listObj := []string{"first","second","third"}

	listObjEnc := encoder.Encode(listObj)

	if string(listObjEnc) != "l5:first6:second5:thirde" {
		t.Error("List encoded value incorrect")
	}
}

func TestInnerMapEncode(t *testing.T) {
	encoder := NewEncoder()

	outerMap := make(map[string]interface{})
	innerMap := make(map[string]int)

	innerMap["one"] = 1

	outerMap["inner"] = innerMap

	twoMapObjEnc := encoder.Encode(outerMap)

	if string(twoMapObjEnc) != "d5:innerd3:onei1eee" {
		t.Error("Inner map encoded value incorrect")
	}
}
