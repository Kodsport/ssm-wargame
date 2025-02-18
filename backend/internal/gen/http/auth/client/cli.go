// Code generated by goa v3.5.2, DO NOT EDIT.
//
// auth HTTP client CLI support package
//
// Command:
// $ goa gen github.com/sakerhetsm/ssm-wargame/internal/design -o internal/

package client

import (
	"encoding/json"
	"fmt"

	auth "github.com/sakerhetsm/ssm-wargame/internal/gen/auth"
)

// BuildExchangeDiscordPayload builds the payload for the auth ExchangeDiscord
// endpoint from CLI flags.
func BuildExchangeDiscordPayload(authExchangeDiscordBody string) (*auth.ExchangeDiscordPayload, error) {
	var err error
	var body ExchangeDiscordRequestBody
	{
		err = json.Unmarshal([]byte(authExchangeDiscordBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"code\": \"123abc\",\n      \"state\": \"15773059ghq9183habn\"\n   }'")
		}
	}
	v := &auth.ExchangeDiscordPayload{
		Code:  body.Code,
		State: body.State,
	}

	return v, nil
}
