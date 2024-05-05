package snowflake

import (
	"fmt"
	"testing"
)

// depend on etc.cfg
func TestGen(t *testing.T) {
	id, err := New()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(id)

	ids, err := NewBuffer(2000)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(ids)
}
