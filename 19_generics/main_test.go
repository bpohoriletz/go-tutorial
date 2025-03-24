package generics

import "testing"

func TestAssertFunctions(t *testing.T) {
	t.Run("Integers", func(t *testing.T) {
		AssertEqual(t, 1, 1)
		AssertNotEqual(t, 2, 1)
	})

	t.Run("Strings", func(t *testing.T) {
		AssertEqual(t, "1", "1")
		AssertNotEqual(t, "2", "1")
	})

	//AssertEqual(t, 1, "1")
}

func AssertEqual[T comparable](t *testing.T, a, b T) {
	t.Helper()
	if a != b {
		t.Errorf("Not equal %+v, and %+v", a, b)
	}
}

func AssertNotEqual[T comparable](t *testing.T, a, b T) {
	t.Helper()
	if a == b {
		t.Errorf("Equal %+v, and %+v", a, b)
	}
}
