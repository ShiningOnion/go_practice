package errorSample

import "testing"

func TestIsMyError(t *testing.T) {
	err := ErrUserNameExist{UserName: "Leo"}

	ok := IsErrUserNameExist(err)

	if !ok {
		t.Fatal("testing error")
	}
}
