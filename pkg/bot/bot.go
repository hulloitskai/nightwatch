package bot

import (
	"fmt"
	"strings"
)

// A Bot is capable of receiving events from Reddit (as a logged-in account),
// using package github.com/stevenxie/graw.
type Bot struct {
	*creds          // credentials for authentication
	version, uagent string
}

// New returns a Bot with the default configuration, and which uses the provided
// version string.
//
// The Bot will use credentials read from the environment; if these variables
// are not set, New will return an error (with code InvalidConfig).
func New(version string) (*Bot, error) {
	if version != "" {
		version = strings.TrimPrefix(version, "v") // trim 'v' prefix
	} else {
		version = "unset"
	}

	// Read creds from environment.
	c, err := readCreds()
	if err != nil {
		return nil, err
	}

	// Only accept valid auth configurations.
	if !c.IsValid() {
		c = nil
		err = &Error{
			Code: InvalidConfig,
			msg:  "bot: environment does not contain valid auth configuration",
		}
	}
	return &Bot{
		creds:   c,
		version: version,
	}, err
}

// UserAgent returns the Bot's user agent string.
func (b *Bot) UserAgent() string {
	if b.uagent == "" { // generate user agent string
		b.uagent = fmt.Sprintf("%s:%s:%s (by %s)", platform, appid, b.version,
			author)
	}
	return b.uagent
}

// Version describes the Bot's version number.
func (b *Bot) Version() string { return b.version }
