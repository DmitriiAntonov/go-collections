package heap

import (
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {

	t.Run("Should create new heap", func(t *testing.T) {
		comparator := func(i int, j int) bool {
			return i < j
		}

		if got := New(comparator); !(reflect.ValueOf(comparator).Pointer() == reflect.ValueOf(got.less).Pointer()) {
			t.Errorf("The comparators are different")
		}
	})
}

func TestPushAndPop(t *testing.T) {
	type args struct {
		numbers []int
	}

	type testCase struct {
		name string
		args args
		less func(i, j int) bool
		want []int
	}

	testCases := []testCase{
		{
			name: "Should create a min heap and push values",
			args: args{numbers: []int{10, -5, -3, 4, 7, 2, -1, 1}},
			less: func(i, j int) bool {
				return i < j
			},
			want: []int{-5, -3, -1, 1, 2, 4, 7, 10},
		},
		{
			name: "Should create a max heap and push values",
			args: args{numbers: []int{10, -5, -3, 4, 7, 2, -1, 1}},
			less: func(i, j int) bool {
				return i > j
			},
			want: []int{10, 7, 4, 2, 1, -1, -3, -5},
		},
	}

	for _, tt := range testCases {

		t.Run(tt.name, func(t *testing.T) {
			heap := New[int](tt.less)

			for _, number := range tt.args.numbers {
				heap.Push(number)
			}

			for index := 0; index < len(tt.args.numbers); index++ {
				got, want := heap.Pop(), tt.want[index]
				if !reflect.DeepEqual(got, want) {
					t.Errorf("Want: %d, but got %d", want, got)
				}
			}
		})
	}
}
