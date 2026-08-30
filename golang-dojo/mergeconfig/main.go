package main

import (
	"fmt"
	"maps"
	"slices"
)

type Config struct {
	Name   string
	Tags   []string
	Limits map[string]int
}

// MergeConfig returns a new Config combining base and override.
//
// Override rule — a single rule applicable to all fields:
// an override field left at its zero value is not applied.
//
//   - Name    : replaced if override.Name != ""
//   - Tags    : replaced wholesale if override.Tags != nil
//     (a non-nil empty list, []string{}, clears the tags)
//   - Limits  : merged key-by-key if override.Limits != nil
//     (keys absent from override are preserved)
//
// Guarantees:
//   - neither base nor override is modified;
//   - the result shares no mutable structure with them.
//
// KNOWN LIMITATION: `Name` cannot be reset to the empty string, because `string`
// has no value distinct from its valid values ​​to signify "absent".
// If this need arises, this signature will no longer suffice.
func MergeConfig(base, override Config) Config {
	result := Config{
		Name:   base.Name,
		Tags:   slices.Clone(base.Tags),
		Limits: maps.Clone(base.Limits),
	}

	if override.Name != "" {
		result.Name = override.Name
	}
	if override.Tags != nil {
		result.Tags = slices.Clone(override.Tags)
	}
	if override.Limits != nil {
		if result.Limits == nil {
			result.Limits = make(map[string]int, len(override.Limits))
		}
		maps.Copy(result.Limits, override.Limits)
	}
	return result
}

func main() {
	bases := []Config{
		{
			Name:   "PAYMENT1",
			Tags:   []string{"prod"},
			Limits: map[string]int{"max": 10},
		},
		{
			Tags:   []string{"prod"},
			Limits: map[string]int{"max": 10},
		},
		{
			Name:   "PAYMENT1",
			Limits: map[string]int{"max": 10},
		},
		{
			Name: "PAYMENT1",
			Tags: []string{"prod"},
		},
		{
			Name: "PAYMENT1",
			Tags: []string{"prod"},
		},
	}
	overrides := []Config{
		{
			Name:   "",
			Tags:   []string{},
			Limits: map[string]int{"max": 20, "timeout": 10},
		},
		{
			Name:   "PAYMENT2",
			Tags:   []string{},
			Limits: map[string]int{"max": 20, "timeout": 10},
		},
		{
			Name:   "PAYMENT2",
			Tags:   []string{},
			Limits: map[string]int{"max": 20, "timeout": 10},
		},
		{
			Name:   "PAYMENT2",
			Tags:   []string{},
			Limits: map[string]int{"max": 20, "timeout": 10},
		},
		{
			Name: "PAYMENT2",
			Tags: []string{},
		},
	}

	for index, override := range overrides {
		mergedConfig := MergeConfig(bases[index], override)
		fmt.Println("Merged Config : #", index, " - ", mergedConfig)
	}
}
