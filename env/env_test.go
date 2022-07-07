package env

import (
	"testing"

	"github.com/unfavorablenode/thin_node/utils"
)

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
	if !utils.StringSlicesEqual(result, tc.want) {
	    t.Errorf("Retrieve valid lines: %q, want %q", result, tc.want)
	}
    }
}

func TestReturnDefaultRouteIfNonePassed(t *testing.T)	{
    testcases := []struct   {
	in, want []string
    }{
	{ in: []string{}, want: []string{".env"}, },
	{ in: []string{".env", ".env.template"}, want: []string{".env", ".env.template"}, },
    }

    for _, tc := range testcases    {
	result := returnDefaultRouteIfNonePassed(tc.in)

	if !utils.StringSlicesEqual(result, tc.want){
	    t.Errorf("ReturnDefaultRouteIfNonePassed: %q, want %q", result, tc.want)
	}
    }
}

