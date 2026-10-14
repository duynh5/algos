package commonprefixscan

import "testing"

func TestLongestCommonPrefix(t *testing.T) {
	tests := []struct {
		name string
		strs []string
		want string
	}{
		{name: "shared prefix", strs: []string{"flower", "flow", "flight"}, want: "fl"},
		{name: "no shared prefix", strs: []string{"dog", "racecar", "car"}, want: ""},
		{name: "shortest string in middle", strs: []string{"prefix", "pre", "prevent"}, want: "pre"},
		{name: "identical strings", strs: []string{"same", "same", "same"}, want: "same"},
		{name: "single string", strs: []string{"alone"}, want: "alone"},
		{name: "single empty string", strs: []string{""}, want: ""},
		{name: "empty collection", strs: nil, want: ""},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := LongestCommonPrefix(test.strs); got != test.want {
				t.Fatalf("LongestCommonPrefix(%q) = %q, want %q", test.strs, got, test.want)
			}
		})
	}
}
