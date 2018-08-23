package util

import "testing"

func TestIPToCountry(t *testing.T) {
	var tests = []struct {
		ip string
		ec string
	}{
		{
			ip: "81.2.69.142",
			ec: "GB",
		},
		{
			ip: "14.177.12.126",
			ec: "VN",
		},
	}
	il, err := NewIPLocator()
	if err != nil {
		t.Fatal(err)
	}
	for _, test := range tests {
		if c, tErr := il.IPToCountry(test.ip); tErr != nil {
			t.Error(tErr)
		} else {
			if c != test.ec {
				t.Errorf("expected country %q for IP %q, got %q", test.ec, test.ip, c)
			}
		}
	}
}
