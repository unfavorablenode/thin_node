package utils

import "testing"

func TestStringSlicesEqual(t *testing.T)    {
    testcases := []struct   {
	in1, in2 []string
	want bool
    }{
	{
	    []string{ "testing1", "testing2" },
	    []string{ "testing1", "testing2" },
	    true,
	},
	{
	    []string{ "testing1", "testing2" },
	    []string{ "testing1" },
	    false,
	},
	{
	    []string{},
	    []string{},
	    true,
	},
    }

    for _, tc := range testcases    {
	if result := StringSlicesEqual(tc.in1, tc.in2); result != tc.want   {
	    t.Errorf("StringSlicesEqual: %t, want: %t", result, tc.want)
	}
    }
}
