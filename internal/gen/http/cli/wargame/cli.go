// Code generated by goa v3.5.2, DO NOT EDIT.
//
// wargame HTTP client CLI support package
//
// Command:
// $ goa gen github.com/sakerhetsm/ssm-wargame/internal/design -o internal/

package cli

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	adminc "github.com/sakerhetsm/ssm-wargame/internal/gen/http/admin/client"
	authc "github.com/sakerhetsm/ssm-wargame/internal/gen/http/auth/client"
	challengec "github.com/sakerhetsm/ssm-wargame/internal/gen/http/challenge/client"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// UsageCommands returns the set of commands and sub-commands using the format
//
//    command (subcommand1|subcommand2|...)
//
func UsageCommands() string {
	return `admin (list-challenges|create-challenge|list-monthly-challenges|delete-monthly-challenge|create-monthly-challenge|list-users)
auth (generate-discord-auth-url|exchange-discord)
challenge (list-challenges|list-monthly-challenges|submit-flag)
`
}

// UsageExamples produces an example of a valid invocation of the CLI tool.
func UsageExamples() string {
	return os.Args[0] + ` admin list-challenges --token "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6InN1cCAoIDoiLCJpYXQiOjE1MTYyMzkwMjJ9.niAX9xS6jNYQSX6hleuwGmzkUCuR9OXPRb5BksyMlkg"` + "\n" +
		os.Args[0] + ` auth generate-discord-auth-url` + "\n" +
		os.Args[0] + ` challenge list-challenges --token "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6InN1cCAoIDoiLCJpYXQiOjE1MTYyMzkwMjJ9.niAX9xS6jNYQSX6hleuwGmzkUCuR9OXPRb5BksyMlkg"` + "\n" +
		""
}

// ParseEndpoint returns the endpoint and payload as specified on the command
// line.
func ParseEndpoint(
	scheme, host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restore bool,
) (goa.Endpoint, interface{}, error) {
	var (
		adminFlags = flag.NewFlagSet("admin", flag.ContinueOnError)

		adminListChallengesFlags     = flag.NewFlagSet("list-challenges", flag.ExitOnError)
		adminListChallengesTokenFlag = adminListChallengesFlags.String("token", "REQUIRED", "")

		adminCreateChallengeFlags     = flag.NewFlagSet("create-challenge", flag.ExitOnError)
		adminCreateChallengeBodyFlag  = adminCreateChallengeFlags.String("body", "REQUIRED", "")
		adminCreateChallengeTokenFlag = adminCreateChallengeFlags.String("token", "REQUIRED", "")

		adminListMonthlyChallengesFlags     = flag.NewFlagSet("list-monthly-challenges", flag.ExitOnError)
		adminListMonthlyChallengesTokenFlag = adminListMonthlyChallengesFlags.String("token", "REQUIRED", "")

		adminDeleteMonthlyChallengeFlags           = flag.NewFlagSet("delete-monthly-challenge", flag.ExitOnError)
		adminDeleteMonthlyChallengeChallengeIDFlag = adminDeleteMonthlyChallengeFlags.String("challenge-id", "REQUIRED", "ID of a challenge")
		adminDeleteMonthlyChallengeTokenFlag       = adminDeleteMonthlyChallengeFlags.String("token", "REQUIRED", "")

		adminCreateMonthlyChallengeFlags     = flag.NewFlagSet("create-monthly-challenge", flag.ExitOnError)
		adminCreateMonthlyChallengeBodyFlag  = adminCreateMonthlyChallengeFlags.String("body", "REQUIRED", "")
		adminCreateMonthlyChallengeTokenFlag = adminCreateMonthlyChallengeFlags.String("token", "REQUIRED", "")

		adminListUsersFlags     = flag.NewFlagSet("list-users", flag.ExitOnError)
		adminListUsersTokenFlag = adminListUsersFlags.String("token", "REQUIRED", "")

		authFlags = flag.NewFlagSet("auth", flag.ContinueOnError)

		authGenerateDiscordAuthURLFlags = flag.NewFlagSet("generate-discord-auth-url", flag.ExitOnError)

		authExchangeDiscordFlags    = flag.NewFlagSet("exchange-discord", flag.ExitOnError)
		authExchangeDiscordBodyFlag = authExchangeDiscordFlags.String("body", "REQUIRED", "")

		challengeFlags = flag.NewFlagSet("challenge", flag.ContinueOnError)

		challengeListChallengesFlags     = flag.NewFlagSet("list-challenges", flag.ExitOnError)
		challengeListChallengesTokenFlag = challengeListChallengesFlags.String("token", "", "")

		challengeListMonthlyChallengesFlags     = flag.NewFlagSet("list-monthly-challenges", flag.ExitOnError)
		challengeListMonthlyChallengesTokenFlag = challengeListMonthlyChallengesFlags.String("token", "", "")

		challengeSubmitFlagFlags           = flag.NewFlagSet("submit-flag", flag.ExitOnError)
		challengeSubmitFlagBodyFlag        = challengeSubmitFlagFlags.String("body", "REQUIRED", "")
		challengeSubmitFlagChallengeIDFlag = challengeSubmitFlagFlags.String("challenge-id", "REQUIRED", "ID of a challenge")
		challengeSubmitFlagTokenFlag       = challengeSubmitFlagFlags.String("token", "REQUIRED", "")
	)
	adminFlags.Usage = adminUsage
	adminListChallengesFlags.Usage = adminListChallengesUsage
	adminCreateChallengeFlags.Usage = adminCreateChallengeUsage
	adminListMonthlyChallengesFlags.Usage = adminListMonthlyChallengesUsage
	adminDeleteMonthlyChallengeFlags.Usage = adminDeleteMonthlyChallengeUsage
	adminCreateMonthlyChallengeFlags.Usage = adminCreateMonthlyChallengeUsage
	adminListUsersFlags.Usage = adminListUsersUsage

	authFlags.Usage = authUsage
	authGenerateDiscordAuthURLFlags.Usage = authGenerateDiscordAuthURLUsage
	authExchangeDiscordFlags.Usage = authExchangeDiscordUsage

	challengeFlags.Usage = challengeUsage
	challengeListChallengesFlags.Usage = challengeListChallengesUsage
	challengeListMonthlyChallengesFlags.Usage = challengeListMonthlyChallengesUsage
	challengeSubmitFlagFlags.Usage = challengeSubmitFlagUsage

	if err := flag.CommandLine.Parse(os.Args[1:]); err != nil {
		return nil, nil, err
	}

	if flag.NArg() < 2 { // two non flag args are required: SERVICE and ENDPOINT (aka COMMAND)
		return nil, nil, fmt.Errorf("not enough arguments")
	}

	var (
		svcn string
		svcf *flag.FlagSet
	)
	{
		svcn = flag.Arg(0)
		switch svcn {
		case "admin":
			svcf = adminFlags
		case "auth":
			svcf = authFlags
		case "challenge":
			svcf = challengeFlags
		default:
			return nil, nil, fmt.Errorf("unknown service %q", svcn)
		}
	}
	if err := svcf.Parse(flag.Args()[1:]); err != nil {
		return nil, nil, err
	}

	var (
		epn string
		epf *flag.FlagSet
	)
	{
		epn = svcf.Arg(0)
		switch svcn {
		case "admin":
			switch epn {
			case "list-challenges":
				epf = adminListChallengesFlags

			case "create-challenge":
				epf = adminCreateChallengeFlags

			case "list-monthly-challenges":
				epf = adminListMonthlyChallengesFlags

			case "delete-monthly-challenge":
				epf = adminDeleteMonthlyChallengeFlags

			case "create-monthly-challenge":
				epf = adminCreateMonthlyChallengeFlags

			case "list-users":
				epf = adminListUsersFlags

			}

		case "auth":
			switch epn {
			case "generate-discord-auth-url":
				epf = authGenerateDiscordAuthURLFlags

			case "exchange-discord":
				epf = authExchangeDiscordFlags

			}

		case "challenge":
			switch epn {
			case "list-challenges":
				epf = challengeListChallengesFlags

			case "list-monthly-challenges":
				epf = challengeListMonthlyChallengesFlags

			case "submit-flag":
				epf = challengeSubmitFlagFlags

			}

		}
	}
	if epf == nil {
		return nil, nil, fmt.Errorf("unknown %q endpoint %q", svcn, epn)
	}

	// Parse endpoint flags if any
	if svcf.NArg() > 1 {
		if err := epf.Parse(svcf.Args()[1:]); err != nil {
			return nil, nil, err
		}
	}

	var (
		data     interface{}
		endpoint goa.Endpoint
		err      error
	)
	{
		switch svcn {
		case "admin":
			c := adminc.NewClient(scheme, host, doer, enc, dec, restore)
			switch epn {
			case "list-challenges":
				endpoint = c.ListChallenges()
				data, err = adminc.BuildListChallengesPayload(*adminListChallengesTokenFlag)
			case "create-challenge":
				endpoint = c.CreateChallenge()
				data, err = adminc.BuildCreateChallengePayload(*adminCreateChallengeBodyFlag, *adminCreateChallengeTokenFlag)
			case "list-monthly-challenges":
				endpoint = c.ListMonthlyChallenges()
				data, err = adminc.BuildListMonthlyChallengesPayload(*adminListMonthlyChallengesTokenFlag)
			case "delete-monthly-challenge":
				endpoint = c.DeleteMonthlyChallenge()
				data, err = adminc.BuildDeleteMonthlyChallengePayload(*adminDeleteMonthlyChallengeChallengeIDFlag, *adminDeleteMonthlyChallengeTokenFlag)
			case "create-monthly-challenge":
				endpoint = c.CreateMonthlyChallenge()
				data, err = adminc.BuildCreateMonthlyChallengePayload(*adminCreateMonthlyChallengeBodyFlag, *adminCreateMonthlyChallengeTokenFlag)
			case "list-users":
				endpoint = c.ListUsers()
				data, err = adminc.BuildListUsersPayload(*adminListUsersTokenFlag)
			}
		case "auth":
			c := authc.NewClient(scheme, host, doer, enc, dec, restore)
			switch epn {
			case "generate-discord-auth-url":
				endpoint = c.GenerateDiscordAuthURL()
				data = nil
			case "exchange-discord":
				endpoint = c.ExchangeDiscord()
				data, err = authc.BuildExchangeDiscordPayload(*authExchangeDiscordBodyFlag)
			}
		case "challenge":
			c := challengec.NewClient(scheme, host, doer, enc, dec, restore)
			switch epn {
			case "list-challenges":
				endpoint = c.ListChallenges()
				data, err = challengec.BuildListChallengesPayload(*challengeListChallengesTokenFlag)
			case "list-monthly-challenges":
				endpoint = c.ListMonthlyChallenges()
				data, err = challengec.BuildListMonthlyChallengesPayload(*challengeListMonthlyChallengesTokenFlag)
			case "submit-flag":
				endpoint = c.SubmitFlag()
				data, err = challengec.BuildSubmitFlagPayload(*challengeSubmitFlagBodyFlag, *challengeSubmitFlagChallengeIDFlag, *challengeSubmitFlagTokenFlag)
			}
		}
	}
	if err != nil {
		return nil, nil, err
	}

	return endpoint, data, nil
}

// adminUsage displays the usage of the admin command and its subcommands.
func adminUsage() {
	fmt.Fprintf(os.Stderr, `Service is the admin service interface.
Usage:
    %[1]s [globalflags] admin COMMAND [flags]

COMMAND:
    list-challenges: ListChallenges implements ListChallenges.
    create-challenge: CreateChallenge implements CreateChallenge.
    list-monthly-challenges: ListMonthlyChallenges implements ListMonthlyChallenges.
    delete-monthly-challenge: DeleteMonthlyChallenge implements DeleteMonthlyChallenge.
    create-monthly-challenge: CreateMonthlyChallenge implements CreateMonthlyChallenge.
    list-users: ListUsers implements ListUsers.

Additional help:
    %[1]s admin COMMAND --help
`, os.Args[0])
}
func adminListChallengesUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] admin list-challenges -token STRING

ListChallenges implements ListChallenges.
    -token STRING: 

Example:
    %[1]s admin list-challenges --token "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6InN1cCAoIDoiLCJpYXQiOjE1MTYyMzkwMjJ9.niAX9xS6jNYQSX6hleuwGmzkUCuR9OXPRb5BksyMlkg"
`, os.Args[0])
}

func adminCreateChallengeUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] admin create-challenge -body JSON -token STRING

CreateChallenge implements CreateChallenge.
    -body JSON: 
    -token STRING: 

Example:
    %[1]s admin create-challenge --body '{
      "description": "A heap overflow challenge",
      "score": 50,
      "slug": "pwnme",
      "title": "pwnme"
   }' --token "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6InN1cCAoIDoiLCJpYXQiOjE1MTYyMzkwMjJ9.niAX9xS6jNYQSX6hleuwGmzkUCuR9OXPRb5BksyMlkg"
`, os.Args[0])
}

func adminListMonthlyChallengesUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] admin list-monthly-challenges -token STRING

ListMonthlyChallenges implements ListMonthlyChallenges.
    -token STRING: 

Example:
    %[1]s admin list-monthly-challenges --token "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6InN1cCAoIDoiLCJpYXQiOjE1MTYyMzkwMjJ9.niAX9xS6jNYQSX6hleuwGmzkUCuR9OXPRb5BksyMlkg"
`, os.Args[0])
}

func adminDeleteMonthlyChallengeUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] admin delete-monthly-challenge -challenge-id STRING -token STRING

DeleteMonthlyChallenge implements DeleteMonthlyChallenge.
    -challenge-id STRING: ID of a challenge
    -token STRING: 

Example:
    %[1]s admin delete-monthly-challenge --challenge-id "195229b0-b15f-4ee5-9a99-94bfff492967" --token "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6InN1cCAoIDoiLCJpYXQiOjE1MTYyMzkwMjJ9.niAX9xS6jNYQSX6hleuwGmzkUCuR9OXPRb5BksyMlkg"
`, os.Args[0])
}

func adminCreateMonthlyChallengeUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] admin create-monthly-challenge -body JSON -token STRING

CreateMonthlyChallenge implements CreateMonthlyChallenge.
    -body JSON: 
    -token STRING: 

Example:
    %[1]s admin create-monthly-challenge --body '{
      "challengeID": "195229b0-b15f-4ee5-9a99-94bfff492967",
      "display_month": "Januari/Februari",
      "end_date": "2006-02-01",
      "start_date": "2006-02-01"
   }' --token "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6InN1cCAoIDoiLCJpYXQiOjE1MTYyMzkwMjJ9.niAX9xS6jNYQSX6hleuwGmzkUCuR9OXPRb5BksyMlkg"
`, os.Args[0])
}

func adminListUsersUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] admin list-users -token STRING

ListUsers implements ListUsers.
    -token STRING: 

Example:
    %[1]s admin list-users --token "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6InN1cCAoIDoiLCJpYXQiOjE1MTYyMzkwMjJ9.niAX9xS6jNYQSX6hleuwGmzkUCuR9OXPRb5BksyMlkg"
`, os.Args[0])
}

// authUsage displays the usage of the auth command and its subcommands.
func authUsage() {
	fmt.Fprintf(os.Stderr, `Service is the auth service interface.
Usage:
    %[1]s [globalflags] auth COMMAND [flags]

COMMAND:
    generate-discord-auth-url: GenerateDiscordAuthURL implements GenerateDiscordAuthURL.
    exchange-discord: ExchangeDiscord implements ExchangeDiscord.

Additional help:
    %[1]s auth COMMAND --help
`, os.Args[0])
}
func authGenerateDiscordAuthURLUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] auth generate-discord-auth-url

GenerateDiscordAuthURL implements GenerateDiscordAuthURL.

Example:
    %[1]s auth generate-discord-auth-url
`, os.Args[0])
}

func authExchangeDiscordUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] auth exchange-discord -body JSON

ExchangeDiscord implements ExchangeDiscord.
    -body JSON: 

Example:
    %[1]s auth exchange-discord --body '{
      "code": "123abc",
      "state": "15773059ghq9183habn"
   }'
`, os.Args[0])
}

// challengeUsage displays the usage of the challenge command and its
// subcommands.
func challengeUsage() {
	fmt.Fprintf(os.Stderr, `Service is the challenge service interface.
Usage:
    %[1]s [globalflags] challenge COMMAND [flags]

COMMAND:
    list-challenges: ListChallenges implements ListChallenges.
    list-monthly-challenges: ListMonthlyChallenges implements ListMonthlyChallenges.
    submit-flag: SubmitFlag implements SubmitFlag.

Additional help:
    %[1]s challenge COMMAND --help
`, os.Args[0])
}
func challengeListChallengesUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] challenge list-challenges -token STRING

ListChallenges implements ListChallenges.
    -token STRING: 

Example:
    %[1]s challenge list-challenges --token "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6InN1cCAoIDoiLCJpYXQiOjE1MTYyMzkwMjJ9.niAX9xS6jNYQSX6hleuwGmzkUCuR9OXPRb5BksyMlkg"
`, os.Args[0])
}

func challengeListMonthlyChallengesUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] challenge list-monthly-challenges -token STRING

ListMonthlyChallenges implements ListMonthlyChallenges.
    -token STRING: 

Example:
    %[1]s challenge list-monthly-challenges --token "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6InN1cCAoIDoiLCJpYXQiOjE1MTYyMzkwMjJ9.niAX9xS6jNYQSX6hleuwGmzkUCuR9OXPRb5BksyMlkg"
`, os.Args[0])
}

func challengeSubmitFlagUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] challenge submit-flag -body JSON -challenge-id STRING -token STRING

SubmitFlag implements SubmitFlag.
    -body JSON: 
    -challenge-id STRING: ID of a challenge
    -token STRING: 

Example:
    %[1]s challenge submit-flag --body '{
      "flag": "SSM{flag}"
   }' --challenge-id "195229b0-b15f-4ee5-9a99-94bfff492967" --token "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6InN1cCAoIDoiLCJpYXQiOjE1MTYyMzkwMjJ9.niAX9xS6jNYQSX6hleuwGmzkUCuR9OXPRb5BksyMlkg"
`, os.Args[0])
}
