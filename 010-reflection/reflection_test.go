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
		{
			Name: "Arrays",
			Input: [2]Profile{
				{33, "Dortmund"},
				{34, "Ellwangen"},
			},
			ExpectedCalls: []string{"Dortmund", "Ellwangen"},
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

	t.Run("with maps", func(t *testing.T) {
		aMap := map[string]string{
			"Foo": "Bar",
			"Baz": "Boz",
		}

		var got []string
		walk(aMap, func(input string) {
			got = append(got, input)
		})

		assertContains(t, got, "Bar")
		assertContains(t, got, "Boz")
	})

	t.Run("with channels", func(t *testing.T) {
		aChannel := make(chan Profile)

		go func() {
			aChannel <- Profile{33, "Marburg"}
			aChannel <- Profile{34, "Ellwangen"}
			close(aChannel)
		}()

		var got []string
		want := []string{"Marburg", "Ellwangen"}

		walk(aChannel, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("with function", func(t *testing.T) {
		aFunction := func() (Profile, Profile) {
			return Profile{33, "Marburg"}, Profile{34, "Ellwangen"}
		}

		var got []string
		want := []string{"Marburg", "Ellwangen"}

		walk(aFunction, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}

func assertContains(t testing.TB, haystack []string, needle string) {
	t.Helper()
	contains := false
	for _, x := range haystack {
		if x == needle {
			contains = true
			break
		}
	}
	if !contains {
		t.Errorf("expected %+v to contain %q but it didn't", haystack, needle)
	}
}
