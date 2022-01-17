package oanda

import (
	"net/url"
	"os"
	"reflect"
	"testing"
	"time"

	dot "github.com/MakeNowJust/heredoc/dot"
	"github.com/joho/godotenv"
	"github.com/pkg/errors"
)

func init() {
	if _, err := os.Stat("./.env"); err != nil {
		file, fErr := os.OpenFile("./.env", os.O_WRONLY|os.O_CREATE, 0644)
		if fErr != nil {
			panic(errors.Errorf("Create .env failed: %v", err))
		}
		_, fErr = file.WriteString(dot.D(`
			TOKEN=
			ACCOUNT_ID=
		`))
		if fErr != nil {
			panic(errors.Errorf("Write .env failed: %v", err))
		}

		panic(errors.Errorf(".env file not exist: %v", err))
	}

	err := godotenv.Load()
	if err != nil {
		panic(errors.Errorf("Error loading .env file: %v", err))
	}

	if os.Getenv("TOKEN") == "" {
		panic(errors.Errorf("Env 'TOKEN' is empty"))
	}

	if os.Getenv("ACCOUNT_ID") == "" {
		panic(errors.Errorf("Env 'ACCOUNT_ID' is empty"))
	}
}

func Test_oandaBaseURL(t *testing.T) {
	var expect *baseURLs

	expect = new(baseURLs)
	expect.rest, _ = url.Parse("https://api-fxpractice.oanda.com")
	expect.stream, _ = url.Parse("https://stream-fxpractice.oanda.com")
	if actual := oandaBaseURL(OandaPractice); !reflect.DeepEqual(actual, expect) {
		if !reflect.DeepEqual(actual.rest, expect.rest) {
			t.Errorf("\ngot:  %#v\nwant: %#v", actual.rest, expect.rest)
		}
		if !reflect.DeepEqual(actual.stream, expect.stream) {
			t.Errorf("\ngot:  %#v\nwant: %#v", actual.stream, expect.stream)
		}
	}

	expect = new(baseURLs)
	expect.rest, _ = url.Parse("https://api-fxtrade.oanda.com")
	expect.stream, _ = url.Parse("https://stream-fxtrade.oanda.com")
	if actual := oandaBaseURL(OandaLive); !reflect.DeepEqual(actual, expect) {
		if !reflect.DeepEqual(actual.rest, expect.rest) {
			t.Errorf("\ngot:  %#v\nwant: %#v", actual.rest, expect.rest)
		}
		if !reflect.DeepEqual(actual.stream, expect.stream) {
			t.Errorf("\ngot:  %#v\nwant: %#v", actual.stream, expect.stream)
		}
	}
}

func newConnection(t *testing.T, env OandaEnvironment) *Connection {
	if env == OandaLive {
		t.Fatal("Live environment for testing is prohibited.")
	}

	connection := &Connection{
		Token:       Getenv("TOKEN"),
		Environemnt: env,
		Timeout:     time.Second * 30,
		Strict:      true,
	}

	return connection
}

func Getenv(k string) string {
	return os.Getenv(k)
}
