package quoinex

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/sho3imo/quoinex-go-client/v2/models"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"net/http/httputil"
	"net/url"
	"runtime"
	"time"
)

const (
	baseUrl string = "https://api.liquid.com"
	version string = "0.0.1"
)

type Client struct {
	URL        *url.URL
	ApiTokenID string
	ApiSecret  string
	HTTPClient *http.Client
	Logger     *log.Logger
	testServer *httptest.Server
}

var LiquidAlreadyExistError = errors.New(`{"errors":{"client_order_id":["exists"]}}`)

func NewClient(apiTokenID string, apiSecret string, logger *log.Logger) (*Client, error) {
	if len(apiTokenID) == 0 {
		return nil, fmt.Errorf("apiTokenID is not set")
	}

	if len(apiSecret) == 0 {
		return nil, fmt.Errorf("apiSecret is not set")
	}

	url, err := url.ParseRequestURI(baseUrl)
	if err != nil {
		return nil, err
	}

	var discardLogger = log.New(ioutil.Discard, "", log.LstdFlags)
	if logger == nil {
		logger = discardLogger
	}

	client := &http.Client{Timeout: time.Duration(10) * time.Second}
	return &Client{URL: url, ApiTokenID: apiTokenID, ApiSecret: apiSecret, HTTPClient: client, Logger: logger}, nil

}

func (c *Client) GetInterestRates(ctx context.Context, currency string) (*models.InterestRates, error) {
	spath := fmt.Sprintf("/ir_ladders/%s", currency)
	res, err := c.sendRequest(ctx, "GET", spath, nil, nil)
	if err != nil {
		return nil, err
	}

	var interestRates models.InterestRates
	if err := decodeBody(res, &interestRates); err != nil {
		return nil, err
	}

	return &interestRates, nil
}

func (c *Client) GetOrderBook(ctx context.Context, productID int, full bool) (*models.PriceLevels, error) {
	spath := fmt.Sprintf("/products/%d/price_levels", productID)

	var queryParam *map[string]string
	if full {
		queryParam = &map[string]string{
			"full": "1",
		}
	} else {
		queryParam = nil
	}

	res, err := c.sendRequest(ctx, "GET", spath, nil, queryParam)
	if err != nil {
		return nil, err
	}

	var priceLevels models.PriceLevels
	if err := decodeBody(res, &priceLevels); err != nil {
		return nil, err
	}

	return &priceLevels, nil
}

func (c *Client) GetProducts(ctx context.Context) ([]*models.Product, error) {
	spath := fmt.Sprintf("/products")
	res, err := c.sendRequest(ctx, "GET", spath, nil, nil)
	if err != nil {
		return nil, err
	}

	var products []*models.Product
	if err := decodeBody(res, &products); err != nil {
		return nil, err
	}

	return products, nil
}

func (c *Client) GetProduct(ctx context.Context, productID int) (*models.Product, error) {
	spath := fmt.Sprintf("/products/%d", productID)
	res, err := c.sendRequest(ctx, "GET", spath, nil, nil)
	if err != nil {
		return nil, err
	}

	var product models.Product
	if err := decodeBody(res, &product); err != nil {
		return nil, err
	}

	return &product, nil
}

func (c *Client) newRequest(ctx context.Context, method, spath string, body io.Reader, queryParam *map[string]string) (*http.Request, error) {

	// swith client url for unit test
	if c.testServer != nil {
		url, _ := url.ParseRequestURI(c.testServer.URL)
		*c.URL = *url
	}

	u := *c.URL
	// can't use path.Join in case of end with slash ex: http://quoinex/orders/
	// u.Path = path.Join(c.URL.Path, spath)
	u.Path = c.URL.Path + spath

	// build QueryParameter
	if queryParam != nil {
		q := u.Query()
		for k, v := range *queryParam {
			if v != "" {
				q.Set(k, v)
			}
		}
		u.RawQuery = q.Encode()
	}

	userAgent := fmt.Sprintf("GoClient/%s (%s)", version, runtime.Version())
	if u.RawQuery != "" {
		spath = fmt.Sprintf("%s?%s", spath, u.RawQuery)
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"path":     spath,
		"nonce":    time.Now().Unix(),
		"token_id": c.ApiTokenID,
	})

	tokenString, err := token.SignedString([]byte(c.ApiSecret))
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(method, u.String(), body)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Quoine-API-Version", "2")
	req.Header.Set("User-Agent", userAgent)
	req.Header.Set("X-Quoine-Auth", tokenString)

	return req, nil
}

func (c *Client) sendRequest(ctx context.Context, method, spath string, body io.Reader, queryParam *map[string]string) (*http.Response, error) {
	req, err := c.newRequest(ctx, method, spath, body, queryParam)
	c.Logger.Printf("Request:  %s \n", httpRequestLog(req))
	if err != nil {
		c.Logger.Printf("err: %#v \n", err)
		return nil, err
	}

	res, err := c.HTTPClient.Do(req)
	c.Logger.Printf("Response: %s \n", httpResponseLog(res))
	if err != nil {
		c.Logger.Printf("err: %#v \n", err)
		return nil, err
	}

	if res.StatusCode != 200 {
		defer res.Body.Close()
		bytes, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}
		b := string(bytes)
		if b == (`{"errors":{"client_order_id":["exists"]}}`) {
			return nil, LiquidAlreadyExistError
		} else {
			return nil, errors.New(b)
		}
	}
	return res, nil
}

func decodeBody(resp *http.Response, out interface{}) error {
	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)
	return decoder.Decode(out)
}

func httpResponseLog(resp *http.Response) string {
	b, _ := httputil.DumpResponse(resp, true)
	return string(b)
}
func httpRequestLog(req *http.Request) string {
	b, _ := httputil.DumpRequest(req, true)
	return string(b)
}
