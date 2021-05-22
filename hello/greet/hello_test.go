package greet

import (
	"fmt"
	"testing"
)

func TestSay(t *testing.T) {
	subTests := []struct {
		items  []string
		result string
	}{
		{
			items:  []string{"Uzama"},
			result: "Hello, Uzama!",
		},
		{
			items:  []string{"Zaid"},
			result: "Hello, Zaid!",
		},
		{
			items:  []string{"Uzama", "Zaid"},
			result: "Hello, Uzama, Zaid!",
		},
	}

	for i, st := range subTests {
		t.Run(fmt.Sprintf("test-case-%d", i), func(t *testing.T) {
			if s := Say(st.items); s != st.result {
				t.Errorf("want: %s, but got %s", st.result, s)
			}
		})

	}
}
