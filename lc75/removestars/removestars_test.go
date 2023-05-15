package removestars

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var result string

var tests = []struct {
	input    string
	expected string
}{
	{"leet**cod*e", "lecoe"},
	{"erase*****", ""},
	{
		"thquickb**rownfox*jumped*overthel***azydog*thequickb**rownfox*jumped*overthel***azydogthequickb**rownfox*jumped*overthel***azydothequickb**rownfox*jumped*overthel***azydog*thequickb**rownfox*jumped*overthel***azydogthquickb**rownfox*jumped*overthel***azydog*thequickb**rownfox*jumped*overthel***azydogthequickb**rownfox*jumped*overthel***azydothequickb**rownfox*jumped*overthel***azydog*thequickb**rownfox*jumped*overthel***azydog*asthquickb**rownfox*jumped*overthel***azydog*thequickb**rownfox*jumped*overthel***azydogthequickb**rownfox*jumped*overthel***azydothequickb**rownfox*jumped*overthel***azydog*thequickb**rownfox*jumped*overthel***azydog*bytthquickb**rownfox*jumped*overthel***azydog*thequickb**rownfox*jumped*overthel***azydogthequickb**rownfox*jumped*overthel***azydothequickb**rownfox*jumped*overthel***azydog*thequickb**rownfox*jumped*overthel***azydog*zoothquickb**rownfox*jumped*overthel***azydog*thequickb**rownfox*jumped*overthel***azydogthequickb**rownfox*jumped*overthel***azydothequickb**rownfox*jumped*overthel***azydog*thequickb**rownfox*jumped*overthel***azydog*thankkthquickb**rownfox*jumped*overthel***azydog*thequickb**rownfox*jumped*overthel***azydogthequickb**rownfox*jumped*overthel***azydothequickb**rownfox*jumped*overthel***azydog*thequickb**rownfox*jumped*overthel***azydog*asitthquickb**rownfox*jumped*overthel***azydog*thequickb**rownfox*jumped*overthel***azydogthequickb**rownfox*jumped*overthel***azydothequickb**rownfox*jumped*overthel***azydog*thequickb**rownfox*jumped*overthel***azydog*",
		"thquicrownfojumpeovertazydothequicrownfojumpeovertazydogthequicrownfojumpeovertazydothequicrownfojumpeovertazydothequicrownfojumpeovertazydogthquicrownfojumpeovertazydothequicrownfojumpeovertazydogthequicrownfojumpeovertazydothequicrownfojumpeovertazydothequicrownfojumpeovertazydoasthquicrownfojumpeovertazydothequicrownfojumpeovertazydogthequicrownfojumpeovertazydothequicrownfojumpeovertazydothequicrownfojumpeovertazydobytthquicrownfojumpeovertazydothequicrownfojumpeovertazydogthequicrownfojumpeovertazydothequicrownfojumpeovertazydothequicrownfojumpeovertazydozoothquicrownfojumpeovertazydothequicrownfojumpeovertazydogthequicrownfojumpeovertazydothequicrownfojumpeovertazydothequicrownfojumpeovertazydothankkthquicrownfojumpeovertazydothequicrownfojumpeovertazydogthequicrownfojumpeovertazydothequicrownfojumpeovertazydothequicrownfojumpeovertazydoasitthquicrownfojumpeovertazydothequicrownfojumpeovertazydogthequicrownfojumpeovertazydothequicrownfojumpeovertazydothequicrownfojumpeovertazydo",
	},
}

func TestReverseVowels(t *testing.T) {

	for i, v := range tests {
		t.Log("test: ", i+1)
		got := removeStarsV2(v.input)
		assert.Equal(t, v.expected, got)
	}
}

func BenchmarkReverseVowels(b *testing.B) {
	funcs := []struct {
		name string
		f    func(string) string
	}{
		{"mine", removeStars},
		{"using-stack", removeStarsV2},
		{"leet-fastest", removeStars_leet32ms},
	}

	var res string
	for _, fun := range funcs {
		b.Run(fun.name, func(b *testing.B) {
			b.ReportAllocs()
			b.ResetTimer()
			for k := 0; k < b.N; k++ {
				res = fun.f(tests[k%len(tests)].input)
			}

		})
	}

	result = res
}
