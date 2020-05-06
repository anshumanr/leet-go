package atoi

import "testing"

var result int

var tests = []struct {
	str    string
	intVal int
}{
	{"-3212", -3212},
	{" +234", 234},
	{"  -23 words", -23},
	{"23423", 23423},
	{"   -42  ", -42},
	{"-4293 with words", -4293},
	{"   4193 with words", 4193},
	{"   +4193 with words", 4193},
	{"-91283472332", -2147483648},
	{"91283472332", 2147483647},
	{"2147483648", 2147483647},
	{"9223372036854775808", 2147483647},
	{"-9223372036854775808", -2147483648},
	{"2147483647", 2147483647},
	{"-2147483649", -2147483648},
	{"words with 123", 0},
	{" -+34 ", 0},
	{" +-34", 0},
	{" - 34", 0},
	{" + 43 wer", 0},
	{"+", 0},
	{"-+", 0},
	{"+-", 0},
	{"-", 0},
	{"", 0},
}

func TestAtoi1(t *testing.T) {
	for i, v := range tests {
		got := myAtoi1(v.str)
		if got != v.intVal {
			t.Errorf("Atoi1: Test %d, got: %d, want: %d", i+1, got, v.intVal)
		}
	}
}

func TestAtoi2(t *testing.T) {
	for i, v := range tests {
		got := myAtoi2(v.str)
		if got != v.intVal {
			t.Errorf("Atoi2: Test %d, got: %d, want: %d", i+1, got, v.intVal)
		}
	}
}

func TestAtoi3(t *testing.T) {
	for i, v := range tests {
		got := myAtoi3(v.str)
		if got != v.intVal {
			t.Errorf("Atoi2: Test %d, got: %d, want: %d", i+1, got, v.intVal)
		}
	}
}

func TestAtoi0ms(t *testing.T) {
	for i, v := range tests {
		got := atoiLeet0ms(v.str)
		if got != v.intVal {
			t.Errorf("Atoi0ms: Test %d, got: %d, want: %d", i+1, got, v.intVal)
		}
	}
}

func BenchmarkAtoi(b *testing.B) {
	funcs := []struct {
		name string
		f    func(string) int
	}{
		{"myAtoi-1", myAtoi1},
		{"myAtoi-2", myAtoi2},
		{"myAtoi-3", myAtoi3},
		{"atoiLeetCode-0ms", atoiLeet0ms},
	}

	var res int
	for _, fun := range funcs {

		b.Run(fun.name, func(b *testing.B) {
			b.ReportAllocs()
			b.ResetTimer()
			for k := 0; k < b.N; k++ {
				res = fun.f(tests[k%14].str)
			}

		})

		//fmt.Println(result, arr[result[0]], arr[result[1]])
	}

	result = res
}
