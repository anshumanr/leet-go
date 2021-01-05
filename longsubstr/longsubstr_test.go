package longsubstr

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

var result int

func TestLongSubstrSimple(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{"abcabcbb", 3},
		{"abcdefedcba", 6},
		{"bbbbb", 1},
		{"pwwkew", 3},
		{"", 0},
		{"aab", 2},
		{"dvdf", 3},
		{"abcdghefdqwertyuiop[wertzxcvbnm,.?><';", 24},
		{"abcdghefdqwertyuiop[wertzx", 13},
	}

	for _, v := range tests {
		t.Log("input: ", v.input)
		got := lengthOfLongestSubstringV1(v.input)
		assert.Equal(t, v.want, got)
	}
}

func TestLongSubstrSimpleV2(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{"abcabcbb", 3},
		{"bbbbb", 1},
		{"pwwkew", 3},
		{"", 0},
		{"aab", 2},
		{"dvdf", 3},
		{"abcdefedcba", 6},
		{"abcdghefdqwertyuiop[wertzxcvbnm,.?><';", 24},
		{"abcdghefdqwertyuiop[wertzx", 13},
	}

	for _, v := range tests {
		t.Log("input: ", v.input)
		got := lengthOfLongestSubstringV2(v.input)
		assert.Equal(t, v.want, got)
	}
}

func TestLongSubstrLong(t *testing.T) {

	longbuf1, err := ioutil.ReadFile("./longsubstr_test.go")
	if err != nil {
		t.Error("Test failed")
	}
	longstr1 := string(longbuf1)

	longbuf2, err := ioutil.ReadFile("./longsubstr.go")
	if err != nil {
		t.Error("Test failed")
	}
	longstr2 := string(longbuf2)

	sub := "abcdefghijklmnopqrstuvwxyz1234567890 `~!@#$%^&*()-_=+[{]}|;:,<.>/?'\""
	l := len(sub)

	tests := []struct {
		name  string
		input string
		want  int
		f     func(s string) int
	}{
		//len(sub) = len(abcdefghijklmnopqrstuvwxyz1234567890 `~!@#$%^&*()-_=+[{]}|;:,<.>/?'\") = 68; escaped \" is 1 byte
		//length of "abcdefghijklmnopqrstuvwxyz1234567890 `~!@#$%^&*()-_=+[{]}|;:,<.>/?'\ = 69
		{"test1", longstr1 + sub, l + 1, lengthOfLongestSubstringV1},
		{"test2", longstr2 + "a" + sub, l, lengthOfLongestSubstringV1}, //to break & ignore special char like LF, NL
		{"test3", longstr1 + longstr2 + sub, l + 1, lengthOfLongestSubstringV1},
		{"test4", longstr1 + longstr2 + sub + longstr1, l + 1, lengthOfLongestSubstringV1},
		{"test1-v2", longstr1 + sub, l + 1, lengthOfLongestSubstringV2},
		{"test2-v2", longstr2 + "a" + sub, l, lengthOfLongestSubstringV2}, //to break & ignore special char like LF, NL
		{"test3-v2", longstr1 + longstr2 + sub, l + 1, lengthOfLongestSubstringV2},
		{"test4-v2", longstr1 + longstr2 + sub + longstr1, l + 1, lengthOfLongestSubstringV2},
	}

	for _, v := range tests {
		t.Log("test: ", v.name, l)
		got := v.f(v.input)
		assert.Equal(t, v.want, got)
	}
}

func BenchmarkLongSubstr(b *testing.B) {

	longbuf1, err := ioutil.ReadFile("./longsubstr_test.go")
	if err != nil {
		b.Error("Test failed")
	}
	longstr1 := string(longbuf1)

	longbuf2, err := ioutil.ReadFile("./longsubstr.go")
	if err != nil {
		b.Error("Test failed")
	}
	longstr2 := string(longbuf2)

	sub := "abcdefghijklmnopqrstuvwxyz1234567890 `~!@#$%^&*()-_=+[{]}|;:,<.>/?'\""

	input := longstr1 + sub + longstr2 + longstr2 + longstr1 + sub + longstr2 + longstr1 + sub

	funcs := []struct {
		name string
		f    func(s string) int
	}{
		{"myV1", lengthOfLongestSubstringV1},
		{"myV2", lengthOfLongestSubstringV2},
		{"leet-0ms", lengthOfLongestSubstring0ms},
	}

	var res int
	for _, fun := range funcs {

		b.Run(fun.name, func(b *testing.B) {
			b.ReportAllocs()
			b.ResetTimer()
			for k := 0; k < b.N; k++ {
				res = fun.f(input)
			}

		})
	}

	result = res
}
