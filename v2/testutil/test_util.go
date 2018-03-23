package testutil

import (
	"github.com/jjjjpppp/quoinex-go-client/v2/models"
)

func GetOrderJsonResponse() string {
	return `{
  	"id": 2157479,
  	"order_type": "limit",
  	"quantity": "0.01",
  	"disc_quantity": "0.0",
  	"iceberg_total_quantity": "0.0",
  	"side": "sell",
  	"filled_quantity": "0.01",
  	"price": "500.0",
  	"created_at": 1462123639,
  	"updated_at": 1462123639,
  	"status": "filled",
  	"leverage_level": 2,
  	"source_exchange": "QUOINE",
  	"product_id": 1,
  	"product_code": "CASH",
  	"funding_currency": "USD",
  	"currency_pair_code": "BTCUSD",
  	"order_fee": "0.0",
  	"executions": [
  	  {
  	    "id": 4566133,
  	    "quantity": "0.01",
  	    "price": "500.0",
  	    "taker_side": "buy",
  	    "my_side": "sell",
  	    "created_at": 1465396785
  	  }
  	]
	}`
}

func GetProductsJsonResponse() string {
	return `
	[
    {
        "id": "5",
        "product_type": "CurrencyPair",
        "code": "CASH",
        "name": "CASH Trading",
        "market_ask": "48203.05",
        "market_bid": "48188.15",
        "indicator": -1,
        "currency": "JPY",
        "currency_pair_code": "BTCJPY",
        "symbol": "¥",
        "fiat_minimum_withdraw": "1500.0",
        "pusher_channel": "product_cash_btcjpy_5",
        "taker_fee": "0.0",
        "maker_fee": "0.0",
        "low_market_bid": "47630.99",
        "high_market_ask": "48396.71",
        "volume_24h": "2915.627366519999999998",
        "last_price_24h": "48217.2",
        "last_traded_price": "48203.05",
        "last_traded_quantity": "1.0",
        "quoted_currency": "JPY",
        "base_currency": "BTC",
        "exchange_rate": "0.009398151671149725"
    }
  ]`
}

func GetProductJsonResponse() string {
	return `{
        "id": "5",
        "product_type": "CurrencyPair",
        "code": "CASH",
        "name": "CASH Trading",
        "market_ask": "48203.05",
        "market_bid": "48188.15",
        "indicator": -1,
        "currency": "JPY",
        "currency_pair_code": "BTCJPY",
        "symbol": "¥",
        "fiat_minimum_withdraw": "1500.0",
        "pusher_channel": "product_cash_btcjpy_5",
        "taker_fee": "0.0",
        "maker_fee": "0.0",
        "low_market_bid": "47630.99",
        "high_market_ask": "48396.71",
        "volume_24h": "2915.627366519999999998",
        "last_price_24h": "48217.2",
        "last_traded_price": "48203.05",
        "last_traded_quantity": "1.0",
        "quoted_currency": "JPY",
        "base_currency": "BTC",
        "exchange_rate": "0.009398151671149725"
    }`
}

func GetExpectedOrderModel() *models.Order {
	return &models.Order{
		ID:                   2157479,
		OrderType:            "limit",
		Quantity:             "0.01",
		DiscQuantity:         "0.0",
		IcebergTotalQuantity: "0.0",
		Side:                 "sell",
		FilledQuantity:       "0.01",
		Price:                "500.0",
		CreatedAt:            1462123639,
		UpdatedAt:            1462123639,
		Status:               "filled",
		LeverageLevel:        2,
		SourceExchange:       "QUOINE",
		ProductID:            1,
		ProductCode:          "CASH",
		FundingCurrency:      "USD",
		CurrencyPairCode:     "BTCUSD",
		OrderFee:             "0.0",
		Executions: models.OrderExecutions{
			{
				ID:        4566133,
				Quantity:  "0.01",
				Price:     "500.0",
				TakerSide: "buy",
				MySide:    "sell",
				CreatedAt: 1465396785,
			},
		},
	}
}

func GetExpectedProductsModel() []*models.Product {
	model := &models.Product{
		ID:                  "5",
		ProductType:         "CurrencyPair",
		Code:                "CASH",
		Name:                "CASH Trading",
		MarketAsk:           "48203.05",
		MarketBid:           "48188.15",
		Indicator:           -1,
		Currency:            "JPY",
		CurrencyPairCode:    "BTCJPY",
		Symbol:              "¥",
		FiatMinimumWithdraw: "1500.0",
		PusherChannel:       "product_cash_btcjpy_5",
		TakerFee:            "0.0",
		MakerFee:            "0.0",
		LowMarketBid:        "47630.99",
		HighMarketAsk:       "48396.71",
		Volume24H:           "2915.627366519999999998",
		LastPrice24H:        "48217.2",
		LastTradedPrice:     "48203.05",
		LastTradedQuantity:  "1.0",
		QuotedCurrency:      "JPY",
		BaseCurrency:        "BTC",
		ExchangeRate:        "0.009398151671149725",
	}
	return []*models.Product{model}
}

func GetExpectedProductmodel() *models.Product {
	return &models.Product{
		ID:                  "5",
		ProductType:         "CurrencyPair",
		Code:                "CASH",
		Name:                "CASH Trading",
		MarketAsk:           "48203.05",
		MarketBid:           "48188.15",
		Indicator:           -1,
		Currency:            "JPY",
		CurrencyPairCode:    "BTCJPY",
		Symbol:              "¥",
		FiatMinimumWithdraw: "1500.0",
		PusherChannel:       "product_cash_btcjpy_5",
		TakerFee:            "0.0",
		MakerFee:            "0.0",
		LowMarketBid:        "47630.99",
		HighMarketAsk:       "48396.71",
		Volume24H:           "2915.627366519999999998",
		LastPrice24H:        "48217.2",
		LastTradedPrice:     "48203.05",
		LastTradedQuantity:  "1.0",
		QuotedCurrency:      "JPY",
		BaseCurrency:        "BTC",
		ExchangeRate:        "0.009398151671149725",
	}
}

func GetOrderBookJsonResponse() string {
	return `{
    "buy_price_levels": [
      ["416.23000", "1.75000"], ["0","0"]
    ],
    "sell_price_levels": [
      ["416.47000", "0.28675"], ["1","1"]
    ]
  }`
}

func GetExpectedOrderBookModel() *models.PriceLevels {
	buyPriceLevels := [][]string{{"416.23000", "1.75000"}, {"0", "0"}}
	sellPriceLevels := [][]string{{"416.47000", "0.28675"}, {"1", "1"}}

	return &models.PriceLevels{BuyPriceLevels: buyPriceLevels, SellPriceLevels: sellPriceLevels}
}

func GetExecutionsJsonResponse() string {
	return `{
    "models": [
      {
        "id": 1011880,
        "quantity": "6.118954",
        "price": "409.78",
        "taker_side": "sell",
        "created_at": 1457370745
      },
      {
        "id": 1011791,
        "quantity": "1.15",
        "price": "409.12",
        "taker_side": "sell",
        "created_at": 1457365585
      }
    ],
    "current_page": 2,
    "total_pages": 1686
  }`
}

func GetExpectedExecutionsModel() *models.Executions {
	model1 := &models.ExecutionsModels{ID: 1011880, Quantity: "6.118954", Price: "409.78", TakerSide: "sell", CreatedAt: 1457370745}
	model2 := &models.ExecutionsModels{ID: 1011791, Quantity: "1.15", Price: "409.12", TakerSide: "sell", CreatedAt: 1457365585}
	return &models.Executions{Models: []*models.ExecutionsModels{model1, model2}, CurrentPage: 2, TotalPages: 1686}
}

func GetExecutionsByTimestampJsonResponse() string {
	return `[
    {
      "id": 960598,
      "quantity": "5.6",
      "price": "431.89",
      "taker_side": "buy",
      "created_at": 1456705487
    },
    {
      "id": 960603,
      "quantity": "0.06",
      "price": "431.74",
      "taker_side": "buy",
      "created_at": 1456705564
    }
  ]`
}

func GetExpectedExecutionsByTimestampModel() []*models.ExecutionsModels {
	model1 := &models.ExecutionsModels{ID: 960598, Quantity: "5.6", Price: "431.89", TakerSide: "buy", CreatedAt: 1456705487}
	model2 := &models.ExecutionsModels{ID: 960603, Quantity: "0.06", Price: "431.74", TakerSide: "buy", CreatedAt: 1456705564}

	return []*models.ExecutionsModels{model1, model2}
}

func GetInterestRatesJsonResponse() string {
	return `{
    "bids": [
      [
        "0.00020",
        "23617.81698"
      ],
      [
        "0.00040",
        "50050.42000"
      ],
      [
        "0.00050",
        "100000.00000"
      ]
    ],
    "asks": []
  }`
}

func GetExpectedInterestRatesModel() *models.InterestRates {
	bids := [][]string{{"0.00020", "23617.81698"}, {"0.00040", "50050.42000"}, {"0.00050", "100000.00000"}}
	return &models.InterestRates{Bids: bids, Asks: []interface{}{}}
}

func GetCreateAnOrderJsonResponse() string {
	return `{
    "id": 2157474,
    "order_type": "limit",
    "quantity": "0.01",
    "disc_quantity": "0.0",
    "iceberg_total_quantity": "0.0",
    "side": "sell",
    "filled_quantity": "0.0",
    "price": "500.0",
    "created_at": 1462123639,
    "updated_at": 1462123639,
    "status": "live",
    "leverage_level": 1,
    "source_exchange": "QUOINE",
    "product_id": 1,
    "product_code": "CASH",
    "funding_currency": "USD",
    "currency_pair_code": "BTCUSD",
    "order_fee": "0.0"
  }`
}

func GetExpectedCreateAnOrderModel() *models.Order {
	return &models.Order{
		ID:                   2157474,
		OrderType:            "limit",
		Quantity:             "0.01",
		DiscQuantity:         "0.0",
		IcebergTotalQuantity: "0.0",
		Side:                 "sell",
		FilledQuantity:       "0.0",
		Price:                "500.0",
		CreatedAt:            1462123639,
		UpdatedAt:            1462123639,
		Status:               "live",
		LeverageLevel:        1,
		SourceExchange:       "QUOINE",
		ProductID:            1,
		ProductCode:          "CASH",
		FundingCurrency:      "USD",
		CurrencyPairCode:     "BTCUSD",
		OrderFee:             "0.0",
	}
}
