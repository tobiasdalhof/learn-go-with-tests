package reflection

import (
	"reflect"
	"testing"
)

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age  int
	City string
}

func TestWalk(t *testing.T) {

	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			Name: "Struct with one string field",
			Input: struct {
				Name string
			}{"Jerome"},
			ExpectedCalls: []string{"Jerome"},
		},
		{
			Name: "Struct with two string fields",
			Input: struct {
				Name string
				City string
			}{"Jerome", "Marburg"},
			ExpectedCalls: []string{"Jerome", "Marburg"},
		},
		{
			Name: "Struct with non string fields",
			Input: struct {
				Name string
				Age  int
			}{"Jerome", 42},
			ExpectedCalls: []string{"Jerome"},
		},
		{
			Name: "Nested fields",
			Input: Person{
				Name: "Jerome",
				Profile: Profile{
					Age:  42,
					City: "Marburg",
				},
			},
			ExpectedCalls: []string{"Jerome", "Marburg"},
		},
		{
			Name: "Pointers to thing",
			Input: &Person{
				Name: "Jerome",
				Profile: Profile{
					Age:  42,
					City: "Marburg",
				},
			},
			ExpectedCalls: []string{"Jerome", "Marburg"},
		},
		{
			Name: "Slices",
			Input: []Profile{
				{33, "Dortmund"},
				{34, "Ellwangen"},
				{35, "Marburg"},
			},
			ExpectedCalls: []string{"Dortmund", "Ellwangen", "Marburg"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			walk(test.Input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got %v, want %v", got, test.ExpectedCalls)
			}
		})
	}
}
