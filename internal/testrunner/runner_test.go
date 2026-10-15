package testrunner

import (
	"bytes"
	"slices"
	"strings"
	"testing"
)

func TestDecodeJSONArgs(t *testing.T) {
	var nums []int
	var target int

	err := DecodeJSONArgs(`[[2,7,11,15],9]`, &nums, &target)
	if err != nil {
		t.Fatalf("DecodeJSONArgs() error = %v, want nil", err)
	}
	if !slices.Equal(nums, []int{2, 7, 11, 15}) {
		t.Fatalf("DecodeJSONArgs() nums = %v, want [2 7 11 15]", nums)
	}
	if target != 9 {
		t.Fatalf("DecodeJSONArgs() target = %d, want 9", target)
	}
}

func TestDecodeJSONArgsRejectsWrongArgumentCount(t *testing.T) {
	var nums []int

	err := DecodeJSONArgs(`[[2,7,11,15],9]`, &nums)
	if err == nil || !strings.Contains(err.Error(), "got 2 arguments, want 1") {
		t.Fatalf("DecodeJSONArgs() error = %v, want argument-count error", err)
	}
}

func TestDecodeJSONArgsReturnsArgumentContext(t *testing.T) {
	var nums []int
	var target int

	err := DecodeJSONArgs(`[[2,7,11,15],"nine"]`, &nums, &target)
	if err == nil || !strings.Contains(err.Error(), "decode JSON argument 2") {
		t.Fatalf("DecodeJSONArgs() error = %v, want contextual argument error", err)
	}
}

func TestRunReportsPassesAndFailures(t *testing.T) {
	cases := []Case{{Name: "Example", Input: "hello", Want: "HELLO"}}
	solutions := []Solution{
		{Name: "correct", Run: strings.ToUpper},
		{Name: "incorrect", Run: func(input string) string { return input }},
	}
	var output bytes.Buffer

	err := Run(&output, cases, solutions)
	if err == nil || err.Error() != "1 sample checks failed" {
		t.Fatalf("Run() error = %v, want one failed check", err)
	}
	if !strings.Contains(output.String(), "PASS  correct / Example") {
		t.Fatalf("Run() output missing pass result:\n%s", output.String())
	}
	if !strings.Contains(output.String(), "FAIL  incorrect / Example") {
		t.Fatalf("Run() output missing failure result:\n%s", output.String())
	}
}

func TestRunSelectsConfiguredCase(t *testing.T) {
	t.Setenv(caseNumberEnvironmentVariable, "2")
	cases := []Case{
		{Name: "First", Input: "first", Want: "FIRST"},
		{Name: "Second", Input: "second", Want: "SECOND"},
	}
	var output bytes.Buffer

	if err := Run(&output, cases, []Solution{{Name: "uppercase", Run: strings.ToUpper}}); err != nil {
		t.Fatalf("Run() error = %v, want nil", err)
	}
	if strings.Contains(output.String(), "First") {
		t.Fatalf("Run() output includes unselected case:\n%s", output.String())
	}
	if !strings.Contains(output.String(), "PASS  uppercase / Second") {
		t.Fatalf("Run() output missing selected case:\n%s", output.String())
	}
}

func TestRunRejectsOutOfRangeConfiguredCase(t *testing.T) {
	t.Setenv(caseNumberEnvironmentVariable, "2")
	cases := []Case{{Name: "Only case"}}

	err := Run(&bytes.Buffer{}, cases, []Solution{{Name: "identity", Run: func(input string) string { return input }}})
	if err == nil || err.Error() != "case 2 is out of range; 1 cases configured" {
		t.Fatalf("Run() error = %v, want out-of-range error", err)
	}
}

func TestRunRejectsInvalidConfiguredCase(t *testing.T) {
	t.Setenv(caseNumberEnvironmentVariable, "invalid")

	err := Run(&bytes.Buffer{}, []Case{{Name: "Only case"}}, []Solution{{Name: "identity", Run: func(input string) string { return input }}})
	if err == nil || !strings.Contains(err.Error(), "must be a positive integer") {
		t.Fatalf("Run() error = %v, want positive-integer error", err)
	}
}

func TestRunRejectsInvalidRegistration(t *testing.T) {
	cases := []Case{{Name: "Example"}}
	solutions := []Solution{
		{Name: "duplicate", Run: strings.ToUpper},
		{Name: "duplicate", Run: strings.ToLower},
	}

	err := Run(&bytes.Buffer{}, cases, solutions)
	if err == nil || !strings.Contains(err.Error(), "registered more than once") {
		t.Fatalf("Run() error = %v, want duplicate-name error", err)
	}
}

func TestRunAllowsNoConfiguredCases(t *testing.T) {
	var output bytes.Buffer

	if err := Run(&output, nil, nil); err != nil {
		t.Fatalf("Run() error = %v, want nil", err)
	}
	if !strings.Contains(output.String(), "No sample cases configured") {
		t.Fatalf("Run() output missing empty-workbench guidance: %s", output.String())
	}
}
