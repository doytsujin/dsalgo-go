package structure

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMonoid(t *testing.T) {
	m := Monoid{
		Op: func(x, y interface{}) interface{} {
			return x.(int) + y.(int)
		},
		E: func() interface{} { return 0 },
	}
	if s := m.Op(1, 1); !reflect.DeepEqual(s, 2) {
		t.Fatal(fmt.Sprintf("answer should be 2 but the result was %v", s))
	}
}
