package env

import (
    "testing"
)

func stringSlicesEqual(a, b []string) bool  {
    if len(a) != len(b)	{
	return false
    }
    for i, v := range a	{
	if v != b[i]	{
	    return false
	}
    }

    return true
}

func TestRetrieveValidLinesFromContent(t *testing.T)	{
    testcases := []struct   {
	in string 
	want []string
    }{
	{"", []string{}},
	{"password=", []string{}},
	{"password=pass123", []string{"password=pass123"}},
	{"password=pass123#database password", []string{"password=pass123"}},
	{"password=pass123\nusername=root", []string{"password=pass123", "username=root"}},
	{"#password=pass123\nusername=root", []string{"username=root"}},
    }

    for _, tc := range testcases    {
	result, err := retrieveValidLinesFromContent(tc.in)
	if err != nil	{
	    t.Errorf("Retrieve valid lines error encountered: %q", err)
	}
	if !stringSlicesEqual(result, tc.want) {
	    t.Errorf("Retrieve valid lines: %q, want %q", result, tc.want)
	}
    }
}
