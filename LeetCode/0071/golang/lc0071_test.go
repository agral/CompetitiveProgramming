package lc0071

import "testing"

func Test_simplifyPath(t *testing.T) {
	testcases := []struct {
		path     string
		expected string
	}{
		{"/home/", "/home"},
		{"/home//foo/", "/home/foo"},
		{"/home/user/Documents/../Pictures", "/home/user/Pictures"},
		{"/../", "/"},
		{"/.../a/../b/c/../d/./", "/.../b/d"}, // note: "..." is a valid directory name.
	}

	for i, tc := range testcases {
		actual := simplifyPath(tc.path)
		if actual != tc.expected {
			t.Errorf("Testcase %02d (%v) failed: want %s, got %s", i+1, tc.path, tc.expected, actual)
		}
	}
}
