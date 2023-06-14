package ezenvconfig

import (
	"fmt"
	"os"
)

type Entry struct {
	Name       string
	Aliases    []string
	OnNotFound func()
	Default    string
	Optional   bool
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

	if entry.Default != "" {
		return entry.Default, nil
	}

	err = NoValueForEntry(entry.Name, entry.Aliases)
	return
}
