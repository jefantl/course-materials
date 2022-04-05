// Optional Todo

package hscan

import (
	"fmt"
	"testing"
)

func TestGuessSingle(t *testing.T) {
	got, err := GuessSingle("77f62e3524cd583d698d51fa24fdff4f", "../main/Top304Thousand-probable-v2.txt")
	if err != nil {
		fmt.Println(err)
	}

	want := "foo"
	if got != want {
		t.Errorf("got %s, wanted %s", got, want)
	}

}

func TestGenHashMaps(t *testing.T) {
	GenHashMaps("../main/Top304Thousand-probable-v2.txt")
}

func TestGetMD5(t *testing.T) {
	GenHashMaps("../main/Top304Thousand-probable-v2.txt")

	s, err := GetMD5("90f2c9c53f66540e67349e0ab83d8cd0")
	if err != nil {
		t.Errorf(err.Error())
	}

	if s != "p@ssword" {
		t.Errorf("got %s, wanted %s", s, "p@ssword")
	}
}
