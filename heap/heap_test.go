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

		if got := New(comparator); !(reflect.ValueOf(comparator).Pointer() == reflect.ValueOf(got.comparator).Pointer()) {
			t.Errorf("The comparators are different")
		}
	})
}
