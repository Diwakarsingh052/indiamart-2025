package sum

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// create file with _test.go to register them as test file

// go test -v ./... // run all the tests for the project

func TestSumInt(t *testing.T) {
	input := []int{1, 2, 3, 4, 5}
	want := 15
	got := SumInt(input)

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}

	want = 0
	got = SumInt(nil)
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}

}

func TestSumIntV2(t *testing.T) {

	tt := []struct {
		name  string
		input []int
		want  int
	}{
		{
			name:  "1 to 5 numbers",
			input: []int{1, 2, 3, 4, 5},
			want:  15,
		},
		{
			name:  "empty slice",
			input: []int{},
			want:  0,
		},
		{
			name:  "nil slice",
			input: nil,
			want:  0,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			got := SumInt(tc.input)
			require.Equal(t, tc.want, got)

		})

	}
}
