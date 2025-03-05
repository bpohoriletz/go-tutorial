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

func TestWalkTable(t *testing.T) {
	tests := []struct {
		Name          string
		Input         any
		ExpectedCalls []string
	}{
		{
			Name: "struct with one field",
			Input: struct {
				Name string
			}{"Chris"},
			ExpectedCalls: []string{"Chris"},
		},
		{
			Name: "struct with two fields",
			Input: struct {
				Name  string
				Title string
			}{"Bob", "Sponge"},
			ExpectedCalls: []string{"Bob", "Sponge"},
		},
		{
			Name: "struct with string and int",
			Input: struct {
				Name string
				Age  int
			}{"Chris", 33},
			ExpectedCalls: []string{"Chris"},
		},
		{"Nested struct", Person{"Chris", Profile{33, "London"}}, []string{"Chris", "London"}},
		{"Pointer struct", &Person{"Chris", Profile{33, "London"}}, []string{"Chris", "London"}},
		{"Slices", []Profile{{33, "London"}, {34, "Lviv"}}, []string{"London", "Lviv"}},
		{"Arrays", [2]Profile{{33, "London"}, {34, "Lviv"}}, []string{"London", "Lviv"}},
		{"Maps", map[string]string{"Cow": "Moo", "Dog": "Bark"}, []string{"Moo", "Bark"}},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			walk(test.Input, func(input string) {
				got = append(got, input)
			})
			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("Expected calls %v, got calls %v", test.ExpectedCalls, got)
			}
		})
	}
}

func TestWalk(t *testing.T) {
	t.Run("Must work", func(t *testing.T) {
		expected := "Chris"
		var got []string

		x := struct {
			Name string
		}{expected}

		walk(x, func(input string) {
			got = append(got, input)
		})

		if len(got) != 1 {
			t.Errorf("wrong number of function calls, got %d, want %d", len(got), 1)
		}

		if !reflect.DeepEqual(got[0], expected) {
			t.Errorf("Got ya! Expected %v, got %v", expected, got[0])
		}
	})
}
