package oanda

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/pkg/errors"
)

type header struct {
	key   string
	value string
}

type query struct {
	key   string
	value string
}

type requestParams struct {
	method   string
	endPoint string
	headers  []header
	queries  []query
	body     interface{}
}

type baseURLs struct {
	rest   *url.URL
	stream *url.URL
}

func oandaBaseURL(env OandaEnvironment) *baseURLs {
	var urls *baseURLs

	switch env {
	case OandaPractice:
		urls = &baseURLs{
			rest:   parseURL("https://api-fxpractice.oanda.com"),
			stream: parseURL("https://stream-fxpractice.oanda.com"),
		}
	case OandaLive:
		urls = &baseURLs{
			rest:   parseURL("https://api-fxtrade.oanda.com"),
			stream: parseURL("https://stream-fxtrade.oanda.com"),
		}

	case oandaDummy:
		urls = &baseURLs{
			rest:   parseURL("https://192.0.2.1"),
			stream: parseURL("https://192.0.2.2"),
		}
	}

	return urls
}

func parseURL(urlString string) *url.URL {
	urlStructure, _ := url.Parse(urlString)

	return urlStructure
}

func Int(v int) *int          { return &v }
func String(v string) *string { return &v }
func Bool(v bool) *bool       { return &v }

type schemas interface {
	setHeaders(*http.Response) error
}

func copyHeader(resp *http.Response, header string) ([]string, error) {
	src, ok := resp.Header[header]
	if !ok {
		return nil, errors.Errorf("Header \"%s\" does not exist", header)
	}
	dst := make([]string, len(src))
	copy(dst, src)
	return dst, nil
}

func parseResponse(resp *http.Response, data interface{}, strict bool) (interface{}, error) {
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, errors.Errorf("Read response body failed: %v", err)
	}

	var errMessage string
	switch resp.StatusCode {
	case 200, 201:
		errMessage = ""
		if data == nil {
			return nil, errors.New("Variable that receives the response is nil")
		}
	case 400:
		errMessage = "400 bad request"
		if data == nil {
			data = new(BadRequestError)
		}
	case 401:
		errMessage = "401 unauthorized"
		if data == nil {
			data = new(UnauthorizedError)
		}
	case 403:
		errMessage = "403 forbidden"
		if data == nil {
			data = new(ForbiddenError)
		}
	case 404:
		errMessage = "404 not found"
		if data == nil {
			data = new(NotFoundError)
		}
	// TODO: 405
	// TODO: 416
	default:
		return nil, errors.Errorf("Unexpected status code(%d)", resp.StatusCode)
	}

	if err := json.Unmarshal(body, data); err != nil {
		spew.Dump(body)
		return nil, errors.Errorf("Unmarshal response body failed: %v", err)
	}

	if strict {
		if err := compareJson(data, body); err != nil {
			return nil, errors.Errorf("Response body JSON is different from unmarshalled it: %v", err)
		}
	}

	if resp.StatusCode/100 != 2 {
		return nil, errors.Errorf("%s: %v", errMessage, data)
	}

	{
		sm, ok := data.(schemas)
		if !ok {
			// TODO: ヘッダ未実装時の回避、全て実装したら消す
			return data, nil
		}
		if err := sm.setHeaders(resp); err != nil {
			return nil, errors.Errorf("Set headers failed: %v", err)
		}
	}

	return data, nil
}

func compareJson(jsonObj interface{}, jsonStr []byte) error {
	bytes, err := json.Marshal(jsonObj)
	if err != nil {
		return errors.Errorf("Marshal JSON object failed: %v", err)
	}

	actual := new(interface{})
	if err = json.Unmarshal(bytes, actual); err != nil {
		return errors.Errorf("Reunmarshal JSON string failed: %v", err)
	}

	expect := new(interface{})
	if err = json.Unmarshal(jsonStr, expect); err != nil {
		return errors.Errorf("Unmarshal JSON string failed: %v", err)
	}

	if err := deepEqual(expect, actual, []string{reflect.TypeOf(jsonObj).String()}); err != nil {
		return errors.Errorf("Unmarshalled JSON is lacking:\nExpect: %s\nActual: %s\n: %v", string(jsonStr), string(bytes), err)
	}

	return nil
}

func deepEqual(expect, actual interface{}, breadcrumbs []string) error {
	if reflect.TypeOf(expect).String() == "*interface {}" {
		expect = reflect.Indirect(reflect.ValueOf(expect)).Interface()
	}

	if actual == nil {
		return errors.Errorf("Actual value is not equal to expect\nBreadcrumbs: %s\nExpect: %sActual: %s", strings.Join(breadcrumbs, " > "), spew.Sdump(expect), spew.Sdump(actual))
	}
	if reflect.TypeOf(actual).String() == "*interface {}" {
		actual = reflect.Indirect(reflect.ValueOf(actual)).Interface()
	}

	switch expectValue := expect.(type) {
	case map[string]interface{}:
		actualValue, ok := actual.(map[string]interface{})
		if !ok {
			return errors.Errorf("Actual value is not equal to expect\nBreadcrumbs: %s\nExpect: %+v\nActual: %+v", strings.Join(breadcrumbs, " > "), expect, actual)
		}

		for k := range expectValue {
			if v, ok := expectValue[k].([]interface{}); ok && len(v) == 0 {
				if _, ok := actualValue[k]; !ok {
					return nil
				}
			}

			if err := deepEqual(expectValue[k], actualValue[k], append(breadcrumbs, k)); err != nil {
				return err
			}
		}
	case []interface{}:
		actualValue, ok := actual.([]interface{})
		if !ok {
			return errors.Errorf("Actual value is not equal to expect\nBreadcrumbs: %s\nExpect: %+v\nActual: %+v", strings.Join(breadcrumbs, " > "), expect, actual)
		}

		for n := range expectValue {
			if err := deepEqual(expectValue[n], actualValue[n], append(breadcrumbs, "["+strconv.Itoa(n)+"]")); err != nil {
				return err
			}
		}
	case string:
		switch actualValue := actual.(type) {
		case string:
			if expectValue != actualValue {
				return errors.Errorf("Actual value is not equal to expect\nBreadcrumbs: %s\nExpect: %sActual: %s", strings.Join(breadcrumbs, " > "), spew.Sdump(expect), spew.Sdump(actual))
			}
		case float64:
			ev, _ := strconv.ParseFloat(expectValue, 64)
			if ev != actualValue {
				return errors.Errorf("Actual value is not equal to expect\nBreadcrumbs: %s\nExpect: %sActual: %s", strings.Join(breadcrumbs, " > "), spew.Sdump(expect), spew.Sdump(actual))
			}
		default:
			return errors.Errorf("Actual value is not equal to expect\nBreadcrumbs: %s\nExpect: %sActual: %s", strings.Join(breadcrumbs, " > "), spew.Sdump(expect), spew.Sdump(actual))
		}
	case float64:
		switch actualValue := actual.(type) {
		case float64:
			if expectValue != actualValue {
				return errors.Errorf("Actual value is not equal to expect\nBreadcrumbs: %s\nExpect: %sActual: %s", strings.Join(breadcrumbs, " > "), spew.Sdump(expect), spew.Sdump(actual))
			}
		default:
			return errors.Errorf("Actual value is not equal to expect\nBreadcrumbs: %s\nExpect: %sActual: %s", strings.Join(breadcrumbs, " > "), spew.Sdump(expect), spew.Sdump(actual))
		}
	case bool:
		switch actualValue := actual.(type) {
		case bool:
			if expectValue != actualValue {
				return errors.Errorf("Actual value is not equal to expect\nBreadcrumbs: %s\nExpect: %sActual: %s", strings.Join(breadcrumbs, " > "), spew.Sdump(expect), spew.Sdump(actual))
			}
		default:
			return errors.Errorf("Actual value is not equal to expect\nBreadcrumbs: %s\nExpect: %sActual: %s", strings.Join(breadcrumbs, " > "), spew.Sdump(expect), spew.Sdump(actual))
		}
	default:
		return errors.Errorf("Unexpected type was given\nBreadcrumbs: %s\nType: %s\nExpect: %sActual: %s", strings.Join(breadcrumbs, " > "), reflect.TypeOf(expect).String(), spew.Sdump(expect), spew.Sdump(actual))
	}

	return nil
}

type CandleDataRangeDefinition struct {
	From  time.Time
	To    time.Time
	Count int
}

func NewCandleDataRange() *CandleDataRangeDefinition {
	return &CandleDataRangeDefinition{}
}

func (p *CandleDataRangeDefinition) WithFrom(dt time.Time) *CandleDataRangeDefinition {
	p.From = dt
	if !p.To.IsZero() {
		p.Count = 0
	}
	return p
}

func (p *CandleDataRangeDefinition) WithTo(dt time.Time) *CandleDataRangeDefinition {
	p.To = dt
	if !p.From.IsZero() {
		p.Count = 0
	}
	return p
}

func (p *CandleDataRangeDefinition) WithCount(count int) *CandleDataRangeDefinition {
	p.Count = count
	return p
}

func OandaTimeFrame2Duration(timeFrame CandlestickGranularityDefinition) time.Duration {
	switch timeFrame {
	case S5:
		return time.Second * 5 // 5 second candlesticks, minute alignment
	case S10:
		return time.Second * 10 // 10 second candlesticks, minute alignment
	case S15:
		return time.Second * 15 // 15 second candlesticks, minute alignment
	case S30:
		return time.Second * 30 // 30 second candlesticks, minute alignment
	case M1:
		return time.Minute * 1 // 1 minute candlesticks, minute alignment
	case M2:
		return time.Minute * 2 // 2 minute candlesticks, hour alignment
	case M4:
		return time.Minute * 4 // 4 minute candlesticks, hour alignment
	case M5:
		return time.Minute * 5 // 5 minute candlesticks, hour alignment
	case M10:
		return time.Minute * 10 // 10 minute candlesticks, hour alignment
	case M15:
		return time.Minute * 15 // 15 minute candlesticks, hour alignment
	case M30:
		return time.Minute * 30 // 30 minute candlesticks, hour alignment
	case H1:
		return time.Hour * 1 // 1 hour candlesticks, hour alignment
	case H2:
		return time.Hour * 2 // 2 hour candlesticks, day alignment
	case H3:
		return time.Hour * 3 // 3 hour candlesticks, day alignment
	case H4:
		return time.Hour * 4 // 4 hour candlesticks, day alignment
	case H6:
		return time.Hour * 6 // 6 hour candlesticks, day alignment
	case H8:
		return time.Hour * 8 // 8 hour candlesticks, day alignment
	case H12:
		return time.Hour * 12 // 12 hour candlesticks, day alignment
	case D:
		return time.Hour * 24 // 1 day candlesticks, day alignment
	case W:
		return time.Hour * 24 * 7 // 1 week candlesticks, aligned to start of week
	default:
		panic(fmt.Sprintf("timeframe %s not handled", timeFrame))
	}
}
