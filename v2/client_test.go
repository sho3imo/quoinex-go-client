package quoinex

import (
	"context"
	"fmt"
	"github.com/google/go-cmp/cmp"
	"github.com/sho3imo/quoinex-go-client/v2/models"
	"github.com/sho3imo/quoinex-go-client/v2/testutil"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"
	"testing"
	"time"
)

func TestNewClient(t *testing.T) {
	type Param struct {
		apiToken  string
		apiSecret string
		logger    *log.Logger
	}
	type Expect struct {
		client *Client
		err    error
	}
	cases := []struct {
		param  Param
		expect Expect
	}{
		// test case 1
		{
			param:  Param{apiToken: "", apiSecret: "secret", logger: nil},
			expect: Expect{client: nil, err: fmt.Errorf("apiTokenID is not set")},
		},
		// test case 2
		{
			param:  Param{apiToken: "apiToken", apiSecret: "", logger: nil},
			expect: Expect{client: nil, err: fmt.Errorf("apiSecret is not set")},
		},
		// test case 3
		{
			param: Param{apiToken: "apiToken", apiSecret: "apiSecret", logger: nil},
			expect: Expect{client: &Client{
				ApiTokenID: "apiToken",
				ApiSecret:  "apiSecret",
				HTTPClient: &http.Client{Timeout: time.Duration(10) * time.Second},
				Logger:     log.New(ioutil.Discard, "", log.LstdFlags),
			}, err: nil},
		},
	}
	for _, c := range cases {
		client, e := NewClient(c.param.apiToken, c.param.apiSecret, c.param.logger)
		if client == nil && e.Error() != c.expect.err.Error() {
			t.Errorf("Worng err. test set is %+v", c)
		}
		if client == nil {
			t.Logf("client is nil. skip this case.: test set: %+v", c)
			continue
		}
		if client.ApiTokenID != c.expect.client.ApiTokenID {
			t.Errorf("Worng apiToken. test set: %+v", c)
		}
		if client.ApiSecret != c.expect.client.ApiSecret {
			t.Errorf("Worng ApiSecret. test set: %+v", c)
		}
		if reflect.TypeOf(client.HTTPClient) != reflect.TypeOf(c.expect.client.HTTPClient) {
			t.Errorf("Worng HTTPClient. test set: %+v", c)
		}
		if reflect.TypeOf(client.Logger) != reflect.TypeOf(c.expect.client.Logger) {
			t.Errorf("Worng Logger. test set: %+v", c)
		}
	}
}
func TestNewRequest(t *testing.T) {
	type Param struct {
		method     string
		spath      string
		queryParam *map[string]string
	}
	type Expect struct {
		method string
		url    string
	}
	cases := []struct {
		param  Param
		expect Expect
	}{
		// test case 1
		{
			param:  Param{method: "GET", spath: "product/1", queryParam: nil},
			expect: Expect{method: "GET", url: "https://api.quoine.com/product/1"},
		},
		// test case 2
		{
			param:  Param{method: "GET", spath: "product/1", queryParam: &map[string]string{"product_id": "1", "limit": "1", "page": "1"}},
			expect: Expect{method: "GET", url: "https://api.quoine.com/product/1?limit=1&page=1&product_id=1"},
		},
	}

	for _, c := range cases {
		client, _ := NewClient("apiTokenID", "secret", nil)
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		req, _ := client.newRequest(ctx, c.param.method, c.param.spath, nil, c.param.queryParam)
		if req.Method != c.expect.method {
			t.Errorf("Worng method. case: %+v", c)
		}
		if len(req.Header["X-Quoine-Auth"]) < 1 {
			t.Errorf("Worng Header. case: %+v", c)
		}
		if req.URL.String() != c.expect.url {
			t.Errorf("Worng URL case: %+v, actual: %+v", c, req.URL.String())
		}
	}
}

func TestGetProducts(t *testing.T) {
	type Param struct {
		jsonResponse string
	}
	type Expect struct {
		path     string
		method   string
		body     string
		products []*models.Product
	}

	cases := []struct {
		param  Param
		expect Expect
	}{
		// test case 1
		{
			param:  Param{jsonResponse: testutil.GetProductsJsonResponse()},
			expect: Expect{path: "/products", method: "GET", body: "", products: testutil.GetExpectedProductsModel()},
		},
		// test case 2
	}
	for _, c := range cases {
		ts := testutil.GenerateTestServer(t, c.expect.path, c.expect.method, c.expect.body, c.param.jsonResponse)
		defer ts.Close()

		client, _ := NewClient("apiTokenID", "secret", nil)
		client.testServer = ts
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		products, err := client.GetProducts(ctx)
		if err != nil {
			t.Errorf("Error. %+v", err)
		}
		if !cmp.Equal(products, c.expect.products) {
			t.Errorf("Worng attribute. %+v", cmp.Diff(products, c.expect.products))

		}

	}
}

func TestGetProduct(t *testing.T) {
	type Param struct {
		productID    int
		jsonResponse string
	}
	type Expect struct {
		path    string
		method  string
		body    string
		product *models.Product
	}

	cases := []struct {
		param  Param
		expect Expect
	}{
		// test case 1
		{
			param:  Param{productID: 1, jsonResponse: testutil.GetProductJsonResponse()},
			expect: Expect{path: "/products/1", method: "GET", body: "", product: testutil.GetExpectedProductmodel()},
		},
		// test case 2
	}
	for _, c := range cases {
		ts := testutil.GenerateTestServer(t, c.expect.path, c.expect.method, c.expect.body, c.param.jsonResponse)
		defer ts.Close()

		client, _ := NewClient("apiTokenID", "secret", nil)
		client.testServer = ts
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		product, err := client.GetProduct(ctx, c.param.productID)
		if err != nil {
			t.Errorf("Error. %+v", err)
		}
		if !cmp.Equal(product, c.expect.product) {
			t.Errorf("Worng attribute. %+v", cmp.Diff(product, c.expect.product))
		}
	}
}

func TestGetOrderBook(t *testing.T) {
	type Param struct {
		productID    int
		full         bool
		jsonResponse string
	}
	type Expect struct {
		path        string
		method      string
		body        string
		priceLevels *models.PriceLevels
	}

	cases := []struct {
		param  Param
		expect Expect
	}{
		// test case 1
		{
			param:  Param{productID: 1, full: true, jsonResponse: testutil.GetOrderBookJsonResponse()},
			expect: Expect{path: "/products/1/price_levels?full=1", method: "GET", body: "", priceLevels: testutil.GetExpectedOrderBookModel()},
		},
		// test case 2
	}
	for _, c := range cases {
		ts := testutil.GenerateTestServer(t, c.expect.path, c.expect.method, c.expect.body, c.param.jsonResponse)
		defer ts.Close()

		client, _ := NewClient("apiTokenID", "secret", nil)
		client.testServer = ts
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		priceLevels, err := client.GetOrderBook(ctx, c.param.productID, true)
		if err != nil {
			t.Errorf("Error. %+v", err)
		}
		if !cmp.Equal(priceLevels, c.expect.priceLevels) {
			t.Errorf("Worng attribute. %+v", cmp.Diff(priceLevels, c.expect.priceLevels))
		}
	}
}

func TestGetInterestRates(t *testing.T) {
	type Param struct {
		currency     string
		jsonResponse string
	}
	type Expect struct {
		path   string
		method string
		body   string
		a      *models.InterestRates
	}
	cases := []struct {
		param  Param
		expect Expect
	}{
		// test case 1
		{
			param:  Param{currency: "USD", jsonResponse: testutil.GetInterestRatesJsonResponse()},
			expect: Expect{path: "/ir_ladders/USD", method: "GET", body: "", a: testutil.GetExpectedInterestRatesModel()},
		},
		// test case 2
	}
	for _, c := range cases {
		ts := testutil.GenerateTestServer(t, c.expect.path, c.expect.method, c.expect.body, c.param.jsonResponse)
		defer ts.Close()

		client, _ := NewClient("apiTokenID", "secret", nil)
		client.testServer = ts
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		r, err := client.GetInterestRates(ctx, c.param.currency)
		if err != nil {
			t.Errorf("Error. %+v", err)
		}
		if !cmp.Equal(r, c.expect.a) {
			t.Errorf("Worng attribute. %+v", cmp.Diff(r, c.expect.a))
		}
	}
}
