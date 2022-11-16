package goparsify

import (
	"errors"
	"testing"
)

func TestUnparsedInputError_Is(t *testing.T) {
	t.Run("some other kind of error is not an UnparsedInputError", func(t *testing.T) {
		err := errors.New("something went wrong")
		if errors.Is(err, UnparsedInputError{}) {
			t.Fatal("error got classified as UnparsedInputError")
		}
	})
	t.Run("UnparsedInputError is an UnparsedInputError", func(t *testing.T) {
		err := UnparsedInputError{"more stuff"}
		if !errors.Is(err, UnparsedInputError{}) {
			t.Fatal("error did not get classified as UnparsedInputError")
		}
	})
}
