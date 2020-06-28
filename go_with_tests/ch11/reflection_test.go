package main

import (
	"reflect"
	"sort"
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
			"Nested fields",
			Person{
				"Chris",
				Profile{33, "Toronto"},
			},
			[]string{"Chris", "Toronto"},
		},
		{
			"Pointers to things",
			&Person{
				"Chris",
				Profile{33, "Toronto"},
			},
			[]string{"Chris", "Toronto"},
		},
		{
			"Slices",
			[]Profile{
				{33, "Toronto"},
				{34, "London"},
			},
			[]string{"Toronto", "London"},
		},
		{
			"Arrays",
			[2]Profile{
				{33, "Toronto"},
				{34, "London"},
			},
			[]string{"Toronto", "London"},
		},
		{
			"Maps",
			map[string]string{
				"TR": "Toronto",
				"LN": "London",
			},
			[]string{"Toronto", "London"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string

			walk(test.Input, func(input string) {
				got = append(got, input)
			})

			sort.Strings(got)
			sort.Strings(test.ExpectedCalls)

			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got %v, want %v", got, test.ExpectedCalls)
			}
		})
	}
}
