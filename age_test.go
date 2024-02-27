package age

import (
	"strconv"
	"testing"
	"time"
)

func TestAge(t *testing.T) {
	cases := []struct {
		birth  string
		target string
		age    int
	}{
		{
			birth:  "2023-02-27T20:28:01+01:00",
			target: "2022-02-27T20:28:01+01:00",
			age:    -1,
		},
		{
			birth:  "2023-02-27T20:28:01+01:00",
			target: "2024-02-27T19:28:02+00:00",
			age:    1,
		},
		{
			birth:  "2023-02-27T20:28:01+01:00",
			target: "2024-02-27T19:28:00+00:00",
			age:    0,
		},
	}

	for i, tt := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			birthT, err := time.Parse(time.RFC3339, tt.birth)
			if err != nil {
				t.Fatal(err)
			}
			targetT, err := time.Parse(time.RFC3339, tt.target)
			if err != nil {
				t.Fatal(err)
			}
			actual := Age(birthT, targetT)
			if actual != tt.age {
				t.Fatalf("expect=%v, got=%v", tt.age, actual)
			}
		})
	}
}

func TestParse(t *testing.T) {
	cases := []struct {
		input  string
		expect T
	}{
		{
			input: "2024-02-27T20:28:01+01:00",
			expect: T{
				Year:  2024,
				Month: 2,
				Day:   27,
				Hour:  19,
				Min:   28,
				Sec:   01,
			},
		},
	}

	for _, tt := range cases {
		t.Run(tt.input, func(t *testing.T) {
			it, err := time.Parse(time.RFC3339, tt.input)
			if err != nil {
				t.Fatal(err)
			}
			actual := parse(it)
			if actual != tt.expect {
				t.Fatalf("expect=%v, got=%v", tt.expect, actual)
			}
		})
	}
}
