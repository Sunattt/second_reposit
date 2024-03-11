package Days11

import (
	"fmt"
	"reflect"
	"testing"
)

func TestTransStr(t *testing.T) {
	var tests = []struct {
		have string
		want []uint8
	}{
		{"abcdef", []uint8{97, 98, 99, 100, 101, 102}},
		{"rlmjnios", []uint8{114, 108, 109, 106, 110, 105, 111, 115}},
		{"sunat", []uint8{115, 117, 110, 97, 116}},
		{"zimeston", []uint8{122, 105, 109, 101, 115, 116, 111, 110}},
	}
	for _, tt := range tests {
		testName := fmt.Sprintf("%v", tt.have)
		t.Run(testName, func(t *testing.T) {
			answer := TransStr(tt.have)
			if !reflect.DeepEqual(answer, tt.want) {
				t.Errorf("got %d, want %d", answer, tt.want)
			}
		})
	}
}

func TestTransByte(t *testing.T) {
	var tests = []struct {
		have []uint8
		want []string
	}{
		{[]uint8{109, 121, 32, 110, 97, 109, 101, 32, 105, 115, 32, 115, 117, 110, 97, 116}, []string{"m", "y", " ", "n", "a", "m", "e", " ", "i", "s", " ", "s", "u", "n", "a", "t"}},
		{[]uint8{119, 111, 114 ,108, 100 }, []string{"w","o","r","l","d"}},
		{[]uint8{104, 117, 109, 111, 108 ,97 ,98}, []string{"h","u","m","o","l","a","b"}},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v", tt.have)
		t.Run(testname, func (t *testing.T) {
			answer := TransByte(tt.have)
			if !reflect.DeepEqual(answer, tt.want){
				t.Errorf("got %v, want %v", answer, tt.want)
			}
		})
	}

}
