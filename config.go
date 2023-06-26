package ezenvconfig

import (
	"fmt"
	"os"
)

// Entry describes an env var entry
type Entry struct {
	// Name is the name of the entry
	Name string
	// Aliases describes the aliases for the entry
	Aliases []string
	// OnNotFound is executed when the entry is not found.
	// It's executed when the package couldn't find the entry.
	// It's executed with the default flag too.
	// It doesn't stops the error returning in  ExtractFromEnv when it's not optional.
	OnNotFound func()
	// Default describes the default value.
	Default string
	// Optional describes if the entry is optional.
	// If true, the ExtractFromEnv doesn't return an error when the entry is not found.
	Optional bool
}

var (
	NoValueForEntry = func(entryName string, aliases []string) error {
		return fmt.Errorf(
			`could not find a value for entry "%s", tried aliases: %v`,
			entryName,
			aliases,
		)
	}
)

func ExtractFromEnv(entry Entry) (value string, err error) {
	for _, alias := range entry.Aliases {
		value, ok := os.LookupEnv(alias)
		if !ok {
			continue
		}

		return value, nil
	}

	if entry.OnNotFound != nil {
		entry.OnNotFound()
	}

	if entry.Default != "" {
		return entry.Default, nil
	}

	if entry.Optional {
		return
	}

	err = NoValueForEntry(entry.Name, entry.Aliases)
	return
}
