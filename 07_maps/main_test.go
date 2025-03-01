package maps

import "testing"

func TestDelete(t *testing.T) {
	t.Run("existing", func(t *testing.T) {
		dic := Dictionary{"thing": "was before"}

		dic.Delete("thing")
		_, err := dic.Search("thing")

		assertError(t, err)
	})

	t.Run("new", func(t *testing.T) {
		dic := Dictionary{}

		err := dic.Delete("thing")

		assertError(t, err)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("existing", func(t *testing.T) {
		dic := Dictionary{"thing": "was before"}

		dic.Update("thing", "is now")
		got, err := dic.Search("thing")

		assertNoError(t, err)
		assertStrings(t, got, "is now")
	})

	t.Run("new", func(t *testing.T) {
		dic := Dictionary{}

		err := dic.Update("thing", "new stuff")

		assertError(t, err)
	})
}

func TestSearch(t *testing.T) {
	t.Run("known", func(t *testing.T) {
		got, _ := Dictionary{"test": "This is just a test"}.Search("test")
		want := "This is just a test"

		assertStrings(t, got, want)
	})

	t.Run("unknown", func(t *testing.T) {
		_, err := Dictionary{"test": "This is just a test"}.Search("something")

		if nil == err {
			t.Fatal("expected an exception.")
		}

		assertStrings(t, err.Error(), ErrNotFound.Error())
	})
}

func TestAdd(t *testing.T) {
	t.Run("new", func(t *testing.T) {
		dic := Dictionary{}
		want := "some text"

		dic.Add("some", want)
		got, err := dic.Search("some")

		assertNoError(t, err)
		assertStrings(t, got, want)

	})
	t.Run("existing", func(t *testing.T) {
		dic := Dictionary{}
		want := "some text"

		dic.Add("some", want)
		dic.Add("some", "not want")
		got, err := dic.Search("some")

		assertNoError(t, err)
		assertStrings(t, got, want)
	})
}

func assertError(t *testing.T, err error) {
	t.Helper()
	if nil == err {
		t.Fatal("expected an exception")
	}
}

func assertNoError(t *testing.T, err error) {
	t.Helper()
	if nil != err {
		t.Fatal("expected no exception")
	}
}

func assertStrings(t *testing.T, got string, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q, want %q, given %q", got, want, "test")
	}
}
