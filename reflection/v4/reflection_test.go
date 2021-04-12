package reflection

import (
	"reflect"
	"testing"
)

func TestWalk(t *testing.T) {
	cases := []struct {
		Name          string
		input         interface{}
		ExpectedCalls []string
	}{
		{
			Name: "Struct with one string field",
			input: struct {
				Name string
				Age  int
			}{"gavin", 23},
			ExpectedCalls: []string{"gavin"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			Walk(test.input, func(input string) {
				got = append(got, input)
			})
			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got %v, want %v", got, test.ExpectedCalls)
			}
		})
	}
}
