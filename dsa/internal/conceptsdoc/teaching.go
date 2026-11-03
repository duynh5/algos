package main

import (
	"bufio"
	"bytes"
	"fmt"
)

var modernTeachingFields = []string{
	"Problem and contract",
	"Mental model",
	"Worked walkthrough",
	"Implementation",
	"Why it works",
	"Costs and tradeoffs",
	"Edge cases and mistakes",
	"When to choose it",
}

func validateTeachingEntries(
	spec conceptGuideSpec,
	prose []byte,
) ([]string, int) {
	issues := make([]string, 0)
	fields := spec.requiredFields
	expected := 0
	entryCount := 0
	entryLevel := 0
	parentLevel := 0
	lineNumber := 0
	// An unfinished lesson must fail at its own boundary, even if a
	// later lesson happens to contain the remaining field names.
	finish := func() {
		if expected > 0 {
			issues = append(issues, fmt.Sprintf(
				"%s:%d: teaching entry stops before %q",
				spec.path, lineNumber, fields[expected],
			))
		}
		expected = 0
		entryLevel = 0
	}
	scanner := bufio.NewScanner(bytes.NewReader(prose))
	for scanner.Scan() {
		lineNumber++
		match := headingPattern.FindStringSubmatch(scanner.Text())
		if match == nil {
			continue
		}
		level := len(match[1])
		name := match[2]
		if entryLevel > 0 && level < entryLevel {
			finish()
		}
		if name == "The idea in simple terms" {
			issues = append(issues, fmt.Sprintf(
				"%s:%d: deprecated teaching field %q; merge it "+
					"into %q",
				spec.path, lineNumber, name, "How it works",
			))
			continue
		}
		// The first field chooses a layout for this lesson only.
		if name == spec.requiredFields[0] ||
			name == modernTeachingFields[0] {
			finish()
			fields = spec.requiredFields
			if name == modernTeachingFields[0] {
				fields = modernTeachingFields
			}
			entryLevel = level
			if level != parentLevel+1 {
				issues = append(issues, fmt.Sprintf(
					"%s:%d: teaching fields must be one level "+
						"below their concept",
					spec.path, lineNumber,
				))
			}
		}
		index := -1
		for fieldIndex, field := range fields {
			if name == field {
				index = fieldIndex
				break
			}
		}
		if index < 0 {
			if entryLevel == 0 {
				parentLevel = level
			}
			continue
		}
		if index != expected {
			issues = append(issues, fmt.Sprintf(
				"%s:%d: teaching field %q appears where %q is "+
					"required",
				spec.path, lineNumber, name, fields[expected],
			))
		}
		if level != entryLevel {
			issues = append(issues, fmt.Sprintf(
				"%s:%d: teaching field %q uses level %d; "+
					"entry started at level %d",
				spec.path, lineNumber, name, level, entryLevel,
			))
		}
		expected = index + 1
		if expected == len(fields) {
			entryCount++
			expected = 0
			entryLevel = 0
		}
	}
	finish()
	if entryCount != spec.expectedEntries {
		issues = append(issues, fmt.Sprintf(
			"%s: teaching entries: got %d, want %d",
			spec.path, entryCount, spec.expectedEntries,
		))
	}
	return issues, entryCount
}
