package oanda

import (
	"encoding/json"
)

type Deprecated interface{}

type Undefined interface{}

//
// Account Definitions
//

type AccountIDDefinition = string

type AccountDefinition struct {
	ID                          AccountIDDefinition                   `json:"id,omitempty"`
	Alias                       string                                `json:"alias,omitempty"`
	Currency                    CurrencyDefinition                    `json:"currency,omitempty"`
	Balance                     AccountUnitsDefinition                `json:"balance,omitempty"`
	CreatedByUserID             *int                                  `json:"createdByUserID,omitempty"`
	CreatedTime                 DateTimeDefinition                    `json:"createdTime,omitempty"`
	GuaranteedStopLossOrderMode GuaranteedStopLossOrderModeDefinition `json:"guaranteedStopLossOrderMode,omitempty"`
	PL                          AccountUnitsDefinition                `json:"pl,omitempty"`
	ResettablePL                AccountUnitsDefinition                `json:"resettablePL,omitempty"`
	ResettablePLTime            DateTimeDefinition                    `json:"resettablePLTime,omitempty"`
	Financing                   AccountUnitsDefinition                `json:"financing,omitempty"`
	Commission                  AccountUnitsDefinition                `json:"commission,omitempty"`
	Dividend                    AccountUnitsDefinition                `json:"dividend,omitempty"`
	DividendAdjustment          AccountUnitsDefinition                `json:"dividendAdjustment,omitempty"`
	GuaranteedExecutionFees     AccountUnitsDefinition                `json:"guaranteedExecutionFees,omitempty"`
	MarginRate                  DecimalNumberDefinition               `json:"marginRate,omitempty"`
	MarginCallEnterTime         DateTimeDefinition                    `json:"marginCallEnterTime,omitempty"`
	MarginCallExtensionCount    *int                                  `json:"marginCallExtensionCount,omitempty"`
	LastMarginCallExtensionTime DateTimeDefinition                    `json:"lastMarginCallExtensionTime,omitempty"`

	// undocumented
	LastDividendAdjustmentTimestamps []DividendAdjustmentTimestampsDefinition `json:"lastDividendAdjustmentTimestamps,omitempty"`
	OpenTradeCount                   *int                                     `json:"openTradeCount,omitempty"`
	OpenPositionCount                *int                                     `json:"openPositionCount,omitempty"`
	PendingOrderCount                *int                                     `json:"pendingOrderCount,omitempty"`
	HedgingEnabled                   *bool                                    `json:"hedgingEnabled,omitempty"`
	UnrealizedPL                     AccountUnitsDefinition                   `json:"unrealizedPL,omitempty"`
	NAV                              AccountUnitsDefinition                   `json:"NAV,omitempty"`
	MarginUsed                       AccountUnitsDefinition                   `json:"marginUsed,omitempty"`
	MarginAvailable                  AccountUnitsDefinition                   `json:"marginAvailable,omitempty"`
	PositionValue                    AccountUnitsDefinition                   `json:"positionValue,omitempty"`
	MarginCloseoutUnrealizedPL       AccountUnitsDefinition                   `json:"marginCloseoutUnrealizedPL,omitempty"`
	MarginCloseoutNAV                AccountUnitsDefinition                   `json:"marginCloseoutNAV,omitempty"`
	MarginCloseoutMarginUsed         AccountUnitsDefinition                   `json:"marginCloseoutMarginUsed,omitempty"`
	MarginCloseoutPercent            DecimalNumberDefinition                  `json:"marginCloseoutPercent,omitempty"`
	MarginCloseoutPositionValue      DecimalNumberDefinition                  `json:"marginCloseoutPositionValue,omitempty"`
	WithdrawalLimit                  AccountUnitsDefinition                   `json:"withdrawalLimit,omitempty"`
	MarginCallMarginUsed             AccountUnitsDefinition                   `json:"marginCallMarginUsed,omitempty"`
	MarginCallPercent                DecimalNumberDefinition                  `json:"marginCallPercent,omitempty"`
	LastTransactionID                TransactionIDDefinition                  `json:"lastTransactionID,omitempty"`
	Trades                           []*TradeSummaryDefinition                `json:"trades,omitempty"`
	Positions                        []*PositionDefinition                    `json:"positions,omitempty"`
	Orders                           []*OrderDefinition                       `json:"orders,omitempty"`
}

type AccountChangesStateDefinition struct {
	UnrealizedPL                AccountUnitsDefinition               `json:"unrealizedPL,omitempty"`
	NAV                         AccountUnitsDefinition               `json:"NAV,omitempty"`
	MarginUsed                  AccountUnitsDefinition               `json:"marginUsed,omitempty"`
	MarginAvailable             AccountUnitsDefinition               `json:"marginAvailable,omitempty"`
	PositionValue               AccountUnitsDefinition               `json:"positionValue,omitempty"`
	MarginCloseoutUnrealizedPL  AccountUnitsDefinition               `json:"marginCloseoutUnrealizedPL,omitempty"`
	MarginCloseoutNAV           AccountUnitsDefinition               `json:"marginCloseoutNAV,omitempty"`
	MarginCloseoutMarginUsed    AccountUnitsDefinition               `json:"marginCloseoutMarginUsed,omitempty"`
	MarginCloseoutPercent       DecimalNumberDefinition              `json:"marginCloseoutPercent,omitempty"`
	MarginCloseoutPositionValue DecimalNumberDefinition              `json:"marginCloseoutPositionValue,omitempty"`
	WithdrawalLimit             AccountUnitsDefinition               `json:"withdrawalLimit,omitempty"`
	MarginCallMarginUsed        AccountUnitsDefinition               `json:"marginCallMarginUsed,omitempty"`
	MarginCallPercent           DecimalNumberDefinition              `json:"marginCallPercent,omitempty"`
	Orders                      []*DynamicOrderStateDefinition       `json:"orders,omitempty"`
	Trades                      []*CalculatedTradeStateDefinition    `json:"trades,omitempty"`
	Positions                   []*CalculatedPositionStateDefinition `json:"positions,omitempty"`
}

type AccountPropertiesDefinition struct {
	ID           AccountIDDefinition `json:"id,omitempty"`
	MT4AccountID *int                `json:"mt4AccountID,omitempty"`
	Tags         []string            `json:"tags,omitempty"`
}

type GuaranteedStopLossOrderModeDefinition = string

// FIXED 	Once a guaranteed Stop Loss Order has been created it cannot be replaced or cancelled.
// REPLACEABLE 	An existing guaranteed Stop Loss Order can only be replaced, not cancelled.
// CANCELABLE 	Once a guaranteed Stop Loss Order has been created it can be either replaced or cancelled.
// PRICE_WIDEN_ONLY 	An existing guaranteed Stop Loss Order can only be replaced to widen the gap from the current price, not cancelled.
type GuaranteedStopLossOrderMutabilityDefinition = string

type GuaranteedStopLossOrderParametersDefinition struct {
	// The current guaranteed Stop Loss Order mutability setting of the Account when market is open.
	MutabilityMarketOpen GuaranteedStopLossOrderMutabilityDefinition `json:"mutabilityMarketOpen,omitempty"`
	// The current guaranteed Stop Loss Order mutability setting of the Account
	// when market is halted.
	MutabilityMarketHalted GuaranteedStopLossOrderMutabilityDefinition `json:"mutabilityMarketHalted,omitempty"`
}

type AccountSummaryDefinition struct {
	// The Account’s identifier
	ID AccountIDDefinition `json:"id,omitempty"`

	// Client-assigned alias for the Account. Only provided if the Account has an alias set
	Alias string `json:"alias,omitempty"`

	// The home currency of the Account
	Currency CurrencyDefinition `json:"currency,omitempty"`

	// The current balance of the account.
	Balance AccountUnitsDefinition `json:"balance,omitempty"`

	// ID of the user that created the Account.
	CreatedByUserID *int `json:"createdByUserID,omitempty"`

	// The date/time when the Account was created.
	CreatedTime DateTimeDefinition `json:"createdTime,omitempty"`

	// The current guaranteed Stop Loss Order mode of the Account.
	GuaranteedStopLossOrderMode GuaranteedStopLossOrderModeDefinition `json:"guaranteedStopLossOrderMode,omitempty"`

	// The total profit/loss realized over the lifetime of the Account.
	PL AccountUnitsDefinition `json:"pl,omitempty"`

	// The total realized profit/loss for the account since it was last reset by the client.
	ResettablePL AccountUnitsDefinition `json:"resettablePL,omitempty"`

	// The date/time that the Account’s resettablePL was last reset.
	ResettablePLTime DateTimeDefinition `json:"resettablePLTime,omitempty"`

	// The total amount of financing paid/collected over the lifetime of the account.
	Financing AccountUnitsDefinition `json:"financing,omitempty"`

	// The total amount of commission paid over the lifetime of the Account.
	Commission AccountUnitsDefinition `json:"commission,omitempty"`

	// The total amount of dividend adjustment paid over the lifetime of the Account in the Account’s home currency.
	DividendAdjustment AccountUnitsDefinition `json:"dividendAdjustment,omitempty"`

	// The total amount of fees charged over the lifetime of the Account for the
	// execution of guaranteed Stop Loss Orders.
	GuaranteedExecutionFees AccountUnitsDefinition `json:"guaranteedExecutionFees,omitempty"`

	// The current guaranteed Stop Loss Order settings of the Account. This
	// field will only be present if the guaranteedStopLossOrderMode is not ‘DISABLED’.
	GuaranteedStopLossOrderParameters GuaranteedStopLossOrderParametersDefinition `json:"guaranteedStopLossOrderParameters,omitempty"`

	// Client-provided margin rate override for the Account. The effective
	// margin rate of the Account is the lesser of this value and the OANDA
	// margin rate for the Account’s division. This value is only provided if a
	// margin rate override exists for the Account.
	MarginRate DecimalNumberDefinition `json:"marginRate,omitempty"`

	// The date/time when the Account entered a margin call state. Only provided
	// if the Account is in a margin call.
	MarginCallEnterTime DateTimeDefinition `json:"marginCallEnterTime,omitempty"`

	// The number of times that the Account’s current margin call was extended.
	MarginCallExtensionCount *int `json:"marginCallExtensionCount,omitempty"`

	// The date/time of the Account’s last margin call extension.
	LastMarginCallExtensionTime DateTimeDefinition `json:"lastMarginCallExtensionTime,omitempty"`

	// The number of Trades currently open in the Account.
	OpenTradeCount *int `json:"openTradeCount,omitempty"`

	// The number of Positions currently open in the Account.
	OpenPositionCount *int `json:"openPositionCount,omitempty"`

	// The number of Orders currently pending in the Account.
	PendingOrderCount *int `json:"pendingOrderCount,omitempty"`

	// Flag indicating that the Account has hedging enabled.
	HedgingEnabled *bool `json:"hedgingEnabled,omitempty"`

	// The total unrealized profit/loss for all Trades currently open in the Account.
	UnrealizedPL AccountUnitsDefinition `json:"unrealizedPL,omitempty"`

	// The net asset value of the Account. Equal to Account balance + unrealizedPL.
	NAV AccountUnitsDefinition `json:"NAV,omitempty"`

	// Margin currently used for the Account.
	MarginUsed AccountUnitsDefinition `json:"marginUsed,omitempty"`

	// Margin available for Account currency.
	MarginAvailable AccountUnitsDefinition `json:"marginAvailable,omitempty"`

	// The value of the Account’s open positions represented in the Account’s home currency.
	PositionValue AccountUnitsDefinition `json:"positionValue,omitempty"`

	// The Account’s margin closeout unrealized PL.
	MarginCloseoutUnrealizedPL AccountUnitsDefinition `json:"marginCloseoutUnrealizedPL,omitempty"`

	// The Account’s margin closeout NAV.
	MarginCloseoutNAV AccountUnitsDefinition `json:"marginCloseoutNAV,omitempty"`

	// The Account’s margin closeout margin used.
	MarginCloseoutMarginUsed AccountUnitsDefinition `json:"marginCloseoutMarginUsed,omitempty"`

	// The Account’s margin closeout percentage. When this value is 1.0 or above
	// the Account is in a margin closeout situation.
	MarginCloseoutPercent DecimalNumberDefinition `json:"marginCloseoutPercent,omitempty"`

	// The value of the Account’s open positions as used for margin closeout
	// calculations represented in the Account’s home currency.
	MarginCloseoutPositionValue DecimalNumberDefinition `json:"marginCloseoutPositionValue,omitempty"`

	// The current WithdrawalLimit for the account which will be zero or a
	// positive value indicating how much can be withdrawn from the account.
	WithdrawalLimit AccountUnitsDefinition `json:"withdrawalLimit,omitempty"`

	// The Account’s margin call margin used.
	MarginCallMarginUsed AccountUnitsDefinition `json:"marginCallMarginUsed,omitempty"`

	// The Account’s margin call percentage. When this value is 1.0 or above the
	// Account is in a margin call situation.
	MarginCallPercent DecimalNumberDefinition `json:"marginCallPercent,omitempty"`

	// The ID of the last Transaction created for the Account.
	LastTransactionID TransactionIDDefinition `json:"lastTransactionID,omitempty"`

	// undocumented
	LastDividendAdjustmentTimestamps []DividendAdjustmentTimestampsDefinition `json:"lastDividendAdjustmentTimestamps,omitempty"`
}

type DividendAdjustmentTimestampsDefinition struct {
	Instrument InstrumentNameDefinition `json:"instrument,omitempty"`
	Timestamp  DateTimeDefinition       `json:"timestamp,omitempty"`
}

type CalculatedAccountStateDefinition struct {
	UnrealizedPL                AccountUnitsDefinition  `json:"unrealizedPL,omitempty"`
	NAV                         AccountUnitsDefinition  `json:"NAV,omitempty"`
	MarginUsed                  AccountUnitsDefinition  `json:"marginUsed,omitempty"`
	MarginAvailable             AccountUnitsDefinition  `json:"marginAvailable,omitempty"`
	PositionValue               AccountUnitsDefinition  `json:"positionValue,omitempty"`
	MarginCloseoutUnrealizedPL  AccountUnitsDefinition  `json:"marginCloseoutUnrealizedPL,omitempty"`
	MarginCloseoutNAV           AccountUnitsDefinition  `json:"marginCloseoutNAV,omitempty"`
	MarginCloseoutMarginUsed    AccountUnitsDefinition  `json:"marginCloseoutMarginUsed,omitempty"`
	MarginCloseoutPercent       DecimalNumberDefinition `json:"marginCloseoutPercent,omitempty"`
	MarginCloseoutPositionValue DecimalNumberDefinition `json:"marginCloseoutPositionValue,omitempty"`
	WithdrawalLimit             AccountUnitsDefinition  `json:"withdrawalLimit,omitempty"`
	MarginCallMarginUsed        AccountUnitsDefinition  `json:"marginCallMarginUsed,omitempty"`
	MarginCallPercent           DecimalNumberDefinition `json:"marginCallPercent,omitempty"`
}

type AccountChangesDefinition struct {
	OrdersCreated   []*OrderDefinition        `json:"ordersCreated,omitempty"`
	OrdersCancelled []*OrderDefinition        `json:"ordersCancelled,omitempty"`
	OrdersFilled    []*OrderDefinition        `json:"ordersFilled,omitempty"`
	OrdersTriggered []*OrderDefinition        `json:"ordersTriggered,omitempty"`
	TradesOpened    []*TradeSummaryDefinition `json:"tradesOpened,omitempty"`
	TradesReduced   []*TradeSummaryDefinition `json:"tradesReduced,omitempty"`
	TradesClosed    []*TradeSummaryDefinition `json:"tradesClosed,omitempty"`
	Positions       []*PositionDefinition     `json:"positions,omitempty"`
	Transactions    []*TransactionDefinition  `json:"transactions,omitempty"`
}

type AccountFinancingModeDefinition = string

type PositionAggregationModeDefinition = string

//
// Instrument Definitions
//

type CandlestickGranularityDefinition = string

const (
	S5  CandlestickGranularityDefinition = "S5"  // 5 second candlesticks, minute alignment
	S10 CandlestickGranularityDefinition = "S10" // 10 second candlesticks, minute alignment
	S15 CandlestickGranularityDefinition = "S15" // 15 second candlesticks, minute alignment
	S30 CandlestickGranularityDefinition = "S30" // 30 second candlesticks, minute alignment
	M1  CandlestickGranularityDefinition = "M1"  // 1 minute candlesticks, minute alignment
	M2  CandlestickGranularityDefinition = "M2"  // 2 minute candlesticks, hour alignment
	M4  CandlestickGranularityDefinition = "M4"  // 4 minute candlesticks, hour alignment
	M5  CandlestickGranularityDefinition = "M5"  // 5 minute candlesticks, hour alignment
	M10 CandlestickGranularityDefinition = "M10" // 10 minute candlesticks, hour alignment
	M15 CandlestickGranularityDefinition = "M15" // 15 minute candlesticks, hour alignment
	M30 CandlestickGranularityDefinition = "M30" // 30 minute candlesticks, hour alignment
	H1  CandlestickGranularityDefinition = "H1"  // 1 hour candlesticks, hour alignment
	H2  CandlestickGranularityDefinition = "H2"  // 2 hour candlesticks, day alignment
	H3  CandlestickGranularityDefinition = "H3"  // 3 hour candlesticks, day alignment
	H4  CandlestickGranularityDefinition = "H4"  // 4 hour candlesticks, day alignment
	H6  CandlestickGranularityDefinition = "H6"  // 6 hour candlesticks, day alignment
	H8  CandlestickGranularityDefinition = "H8"  // 8 hour candlesticks, day alignment
	H12 CandlestickGranularityDefinition = "H12" // 12 hour candlesticks, day alignment
	D   CandlestickGranularityDefinition = "D"   // 1 day candlesticks, day alignment
	W   CandlestickGranularityDefinition = "W"   // 1 week candlesticks, aligned to start of week
	M   CandlestickGranularityDefinition = "M"   // 1 month candlesticks, aligned to first day of the month
)

type WeeklyAlignmentDefinition = string

const (
	Monday    WeeklyAlignmentDefinition = "Monday"
	Tuesday   WeeklyAlignmentDefinition = "Tuesday"
	Wednesday WeeklyAlignmentDefinition = "Wednesday"
	Thursday  WeeklyAlignmentDefinition = "Thursday"
	Friday    WeeklyAlignmentDefinition = "Friday"
	Saturday  WeeklyAlignmentDefinition = "Saturday"
	Sunday    WeeklyAlignmentDefinition = "Sunday"
)

type CandlestickDefinition struct {
	Time     DateTimeDefinition         `json:"time,omitempty"`
	Bid      *CandlestickDataDefinition `json:"bid,omitempty"`
	Ask      *CandlestickDataDefinition `json:"ask,omitempty"`
	Mid      *CandlestickDataDefinition `json:"mid,omitempty"`
	Volume   *int                       `json:"volume,omitempty"`
	Complete *bool                      `json:"complete,omitempty"`
}

type CandlestickDataDefinition struct {
	O PriceValueDefinition `json:"o,omitempty"`
	H PriceValueDefinition `json:"h,omitempty"`
	L PriceValueDefinition `json:"l,omitempty"`
	C PriceValueDefinition `json:"c,omitempty"`
}

type OrderBookDefinition struct {
	Instrument  InstrumentNameDefinition     `json:"instrument,omitempty"`
	Time        DateTimeDefinition           `json:"time,omitempty"`
	Price       PriceValueDefinition         `json:"price,omitempty"`
	BucketWidth PriceValueDefinition         `json:"bucketWidth,omitempty"`
	Buckets     []*OrderBookBucketDefinition `json:"buckets,omitempty"`

	UnixTime Undefined `json:"unixTime,omitempty"`
}

type OrderBookBucketDefinition struct {
	Price             PriceValueDefinition    `json:"price,omitempty"`
	LongCountPercent  DecimalNumberDefinition `json:"longCountPercent,omitempty"`
	ShortCountPercent DecimalNumberDefinition `json:"shortCountPercent,omitempty"`
}

type PositionBookDefinition struct {
	Instrument  InstrumentNameDefinition        `json:"instrument,omitempty"`
	Time        DateTimeDefinition              `json:"time,omitempty"`
	Price       PriceValueDefinition            `json:"price,omitempty"`
	BucketWidth PriceValueDefinition            `json:"bucketWidth,omitempty"`
	Buckets     []*PositionBookBucketDefinition `json:"buckets,omitempty"`

	UnixTime Undefined `json:"unixTime,omitempty"`
}

type PositionBookBucketDefinition struct {
	Price             PriceValueDefinition    `json:"price,omitempty"`
	LongCountPercent  DecimalNumberDefinition `json:"longCountPercent,omitempty"`
	ShortCountPercent DecimalNumberDefinition `json:"shortCountPercent,omitempty"`
}

//
// Order Definitions
//

// Orders

type OrderDefinition struct {
	CancelledTime              DateTimeDefinition                      `json:"cancelledTime,omitempty"`
	CancellingTransactionID    TransactionIDDefinition                 `json:"cancellingTransactionID,omitempty"`
	ClientExtensions           *ClientExtensionsDefinition             `json:"clientExtensions,omitempty"`
	ClientTradeID              string                                  `json:"clientTradeID,omitempty"`
	CreateTime                 DateTimeDefinition                      `json:"createTime,omitempty"`
	DelayedTradeClose          *MarketOrderDelayedTradeCloseDefinition `json:"delayedTradeClose,omitempty"`
	Distance                   DecimalNumberDefinition                 `json:"distance,omitempty"`
	FilledTime                 DateTimeDefinition                      `json:"filledTime,omitempty"`
	FillingTransactionID       TransactionIDDefinition                 `json:"fillingTransactionID,omitempty"`
	GtdTime                    DateTimeDefinition                      `json:"gtdTime,omitempty"`
	Guaranteed                 *bool                                   `json:"guaranteed,omitempty"`
	GuaranteedExecutionPremium DecimalNumberDefinition                 `json:"guaranteedExecutionPremium,omitempty"`
	ID                         string                                  `json:"id,omitempty"`
	InitialMarketPrice         PriceValueDefinition                    `json:"initialMarketPrice,omitempty"`
	Instrument                 InstrumentNameDefinition                `json:"instrument,omitempty"`
	LongPositionCloseout       *MarketOrderPositionCloseoutDefinition  `json:"longPositionCloseout,omitempty"`
	MarginCloseout             *MarketOrderMarginCloseoutDefinition    `json:"marginCloseout,omitempty"`
	PositionFill               OrderPositionFillDefinition             `json:"positionFill,omitempty"`
	Price                      PriceValueDefinition                    `json:"price,omitempty"`
	PriceBound                 PriceValueDefinition                    `json:"priceBound,omitempty"`
	ReplacedByOrderID          string                                  `json:"replacedByOrderID,omitempty"`
	ReplacesOrderID            string                                  `json:"replacesOrderID,omitempty"`
	ShortPositionCloseout      *MarketOrderPositionCloseoutDefinition  `json:"shortPositionCloseout,omitempty"`
	State                      OrderStateDefinition                    `json:"state,omitempty"`
	StopLossOnFill             *StopLossDetailsDefinition              `json:"stopLossOnFill,omitempty"`
	TakeProfitOnFill           *TakeProfitDetailsDefinition            `json:"takeProfitOnFill,omitempty"`
	TimeInForce                TimeInForceDefinition                   `json:"timeInForce,omitempty"`
	TradeClientExtensions      *ClientExtensionsDefinition             `json:"tradeClientExtensions,omitempty"`
	TradeClose                 *MarketOrderTradeCloseDefinition        `json:"tradeClose,omitempty"`
	TradeClosedIDs             []TradeIDDefinition                     `json:"tradeClosedIDs,omitempty"`
	TradeID                    TradeIDDefinition                       `json:"tradeID,omitempty"`
	TradeOpenedID              TradeIDDefinition                       `json:"tradeOpenedID,omitempty"`
	TradeReducedID             TradeIDDefinition                       `json:"tradeReducedID,omitempty"`
	TradeState                 string                                  `json:"tradeState,omitempty"`
	TrailingStopLossOnFill     *TrailingStopLossDetailsDefinition      `json:"trailingStopLossOnFill,omitempty"`
	TrailingStopValue          PriceValueDefinition                    `json:"trailingStopValue,omitempty"`
	TriggerCondition           OrderTriggerConditionDefinition         `json:"triggerCondition,omitempty"`
	TriggerMode                OrderTriggerModeDefinition              `json:"triggerMode,omitempty"`
	Type                       OrderTypeDefinition                     `json:"type,omitempty"`
	Units                      DecimalNumberDefinition                 `json:"units,omitempty"`

	PartialFill Undefined `json:"partialFill"`
}

type TakeProfitOrderDefinition = OrderDefinition
type StopLossOrderDefinition = OrderDefinition
type TrailingStopLossOrderDefinition = OrderDefinition

// Order Requests

// type OrderRequestDefinition
// TODO: Implemented by: MarketOrderRequest, LimitOrderRequest, StopOrderRequest, MarketIfTouchedOrderRequest, TakeProfitOrderRequest, StopLossOrderRequest, TrailingStopLossOrderRequest

type OrderRequestDefinition interface {
	orderRequest()
}

func (MarketOrderRequestDefinition) orderRequest()           {}
func (LimitOrderRequestDefinition) orderRequest()            {}
func (StopOrderRequestDefinition) orderRequest()             {}
func (MarketIfTouchedOrderRequestDefinition) orderRequest()  {}
func (TakeProfitOrderRequestDefinition) orderRequest()       {}
func (StopLossOrderRequestDefinition) orderRequest()         {}
func (TrailingStopLossOrderRequestDefinition) orderRequest() {}

// https://developer.oanda.com/rest-live-v20/order-df/#MarketOrderRequest
type MarketOrderRequestDefinition struct {

	// The type of the Order to Create. Must be set to “MARKET” when creating a
	// Market Order. default=MARKET
	Type OrderTypeDefinition `json:"type,omitempty"`

	// The Market Order’s Instrument.
	Instrument InstrumentNameDefinition `json:"instrument,omitempty"`

	// The quantity requested to be filled by the Market Order. A positive
	// number of units results in a long Order, and a negative number of units
	// results in a short Order.
	Units DecimalNumberDefinition `json:"units,omitempty"`

	// The time-in-force requested for the Market Order. Restricted to FOK or
	// IOC for a MarketOrder. required, default=FOK
	TimeInForce TimeInForceDefinition `json:"timeInForce,omitempty"`

	// The worst price that the client is willing to have the Market Order
	// filled at.
	PriceBound PriceValueDefinition `json:"priceBound,omitempty"`

	// Specification of how Positions in the Account are modified when the Order
	// is filled. required, default=DEFAULT
	PositionFill OrderPositionFillDefinition `json:"positionFill,omitempty"`

	// The client extensions to add to the Order. Do not set, modify, or delete
	// clientExtensions if your account is associated with MT4.
	ClientExtensions *ClientExtensionsDefinition `json:"clientExtensions,omitempty"`

	// TakeProfitDetails specifies the details of a Take Profit Order to be
	// created on behalf of a client. This may happen when an Order is filled
	// that opens a Trade requiring a Take Profit, or when a Trade’s dependent
	// Take Profit Order is modified directly through the Trade.
	TakeProfitOnFill *TakeProfitDetailsDefinition `json:"takeProfitOnFill,omitempty"`

	// StopLossDetails specifies the details of a Stop Loss Order to be created
	// on behalf of a client. This may happen when an Order is filled that opens
	// a Trade requiring a Stop Loss, or when a Trade’s dependent Stop Loss
	// Order is modified directly through the Trade.
	StopLossOnFill *StopLossDetailsDefinition `json:"stopLossOnFill,omitempty"`

	//
	// GuaranteedStopLossDetails specifies the details of a Guaranteed Stop Loss
	// Order to be created on behalf of a client. This may happen when an Order
	// is filled that opens a Trade requiring a Guaranteed Stop Loss, or when a
	// Trade’s dependent Guaranteed Stop Loss Order is modified directly through
	// the Trade.
	//
	// guaranteedStopLossOnFill : (GuaranteedStopLossDetails),

	// TrailingStopLossDetails specifies the details of a Trailing Stop Loss
	// Order to be created on behalf of a client. This may happen when an Order
	// is filled that opens a Trade requiring a Trailing Stop Loss, or when a
	// Trade’s dependent Trailing Stop Loss Order is modified directly through
	// the Trade.
	TrailingStopLossOnFill *TrailingStopLossDetailsDefinition `json:"trailingStopLossOnFill,omitempty"`

	// Client Extensions to add to the Trade created when the Order is filled
	// (if such a Trade is created). Do not set, modify, or delete
	// tradeClientExtensions if your account is associated with MT4.
	TradeClientExtensions *ClientExtensionsDefinition `json:"tradeClientExtensions,omitempty"`
}

type LimitOrderRequestDefinition struct {
	Type                   OrderTypeDefinition                `json:"type,omitempty"`
	Instrument             InstrumentNameDefinition           `json:"instrument,omitempty"`
	Units                  DecimalNumberDefinition            `json:"units,omitempty"`
	Price                  PriceValueDefinition               `json:"price,omitempty"`
	TimeInForce            TimeInForceDefinition              `json:"timeInForce,omitempty"`
	GtdTime                DateTimeDefinition                 `json:"gtdTime,omitempty"`
	PositionFill           OrderPositionFillDefinition        `json:"positionFill,omitempty"`
	TriggerCondition       OrderTriggerConditionDefinition    `json:"triggerCondition,omitempty"`
	ClientExtensions       *ClientExtensionsDefinition        `json:"clientExtensions,omitempty"`
	TakeProfitOnFill       *TakeProfitDetailsDefinition       `json:"takeProfitOnFill,omitempty"`
	StopLossOnFill         *StopLossDetailsDefinition         `json:"stopLossOnFill,omitempty"`
	TrailingStopLossOnFill *TrailingStopLossDetailsDefinition `json:"trailingStopLossOnFill,omitempty"`
	TradeClientExtensions  *ClientExtensionsDefinition        `json:"tradeClientExtensions,omitempty"`
}

type StopOrderRequestDefinition struct {
	Type                   OrderTypeDefinition                `json:"type,omitempty"`
	Instrument             InstrumentNameDefinition           `json:"instrument,omitempty"`
	Units                  DecimalNumberDefinition            `json:"units,omitempty"`
	Price                  PriceValueDefinition               `json:"price,omitempty"`
	PriceBound             PriceValueDefinition               `json:"priceBound,omitempty"`
	TimeInForce            TimeInForceDefinition              `json:"timeInForce,omitempty"`
	GtdTime                DateTimeDefinition                 `json:"gtdTime,omitempty"`
	PositionFill           OrderPositionFillDefinition        `json:"positionFill,omitempty"`
	TriggerCondition       OrderTriggerConditionDefinition    `json:"triggerCondition,omitempty"`
	ClientExtensions       *ClientExtensionsDefinition        `json:"clientExtensions,omitempty"`
	TakeProfitOnFill       *TakeProfitDetailsDefinition       `json:"takeProfitOnFill,omitempty"`
	StopLossOnFill         *StopLossDetailsDefinition         `json:"stopLossOnFill,omitempty"`
	TrailingStopLossOnFill *TrailingStopLossDetailsDefinition `json:"trailingStopLossOnFill,omitempty"`
	TradeClientExtensions  *ClientExtensionsDefinition        `json:"tradeClientExtensions,omitempty"`
}

type MarketIfTouchedOrderRequestDefinition struct {
	Type                   OrderTypeDefinition                `json:"type,omitempty"`
	Instrument             InstrumentNameDefinition           `json:"instrument,omitempty"`
	Units                  DecimalNumberDefinition            `json:"units,omitempty"`
	Price                  PriceValueDefinition               `json:"price,omitempty"`
	PriceBound             PriceValueDefinition               `json:"priceBound,omitempty"`
	TimeInForce            TimeInForceDefinition              `json:"timeInForce,omitempty"`
	GtdTime                DateTimeDefinition                 `json:"gtdTime,omitempty"`
	PositionFill           OrderPositionFillDefinition        `json:"positionFill,omitempty"`
	TriggerCondition       OrderTriggerConditionDefinition    `json:"triggerCondition,omitempty"`
	ClientExtensions       *ClientExtensionsDefinition        `json:"clientExtensions,omitempty"`
	TakeProfitOnFill       *TakeProfitDetailsDefinition       `json:"takeProfitOnFill,omitempty"`
	StopLossOnFill         *StopLossDetailsDefinition         `json:"stopLossOnFill,omitempty"`
	TrailingStopLossOnFill *TrailingStopLossDetailsDefinition `json:"trailingStopLossOnFill,omitempty"`
	TradeClientExtensions  *ClientExtensionsDefinition        `json:"tradeClientExtensions,omitempty"`
}

type TakeProfitOrderRequestDefinition struct {
	Type             OrderTypeDefinition             `json:"type,omitempty"`
	TradeID          TradeIDDefinition               `json:"tradeID,omitempty"`
	ClientTradeID    string                          `json:"clientTradeID,omitempty"`
	Price            PriceValueDefinition            `json:"price,omitempty"`
	TimeInForce      TimeInForceDefinition           `json:"timeInForce,omitempty"`
	GtdTime          DateTimeDefinition              `json:"gtdTime,omitempty"`
	TriggerCondition OrderTriggerConditionDefinition `json:"triggerCondition,omitempty"`
	ClientExtensions *ClientExtensionsDefinition     `json:"clientExtensions,omitempty"`
}

type StopLossOrderRequestDefinition struct {
	Type             OrderTypeDefinition             `json:"type,omitempty"`
	TradeID          TradeIDDefinition               `json:"tradeID,omitempty"`
	ClientTradeID    string                          `json:"clientTradeID,omitempty"`
	Price            PriceValueDefinition            `json:"price,omitempty"`
	Distance         DecimalNumberDefinition         `json:"distance,omitempty"`
	TimeInForce      TimeInForceDefinition           `json:"timeInForce,omitempty"`
	GtdTime          DateTimeDefinition              `json:"gtdTime,omitempty"`
	TriggerCondition OrderTriggerConditionDefinition `json:"triggerCondition,omitempty"`
	Guaranteed       *bool                           `json:"guaranteed,omitempty"`
	ClientExtensions *ClientExtensionsDefinition     `json:"clientExtensions,omitempty"`
}

type TrailingStopLossOrderRequestDefinition struct {
	Type             OrderTypeDefinition             `json:"type,omitempty"`
	TradeID          TradeIDDefinition               `json:"tradeID,omitempty"`
	ClientTradeID    string                          `json:"clientTradeID,omitempty"`
	Distance         DecimalNumberDefinition         `json:"distance,omitempty"`
	TimeInForce      TimeInForceDefinition           `json:"timeInForce,omitempty"`
	GtdTime          DateTimeDefinition              `json:"gtdTime,omitempty"`
	TriggerCondition OrderTriggerConditionDefinition `json:"triggerCondition,omitempty"`
	ClientExtensions *ClientExtensionsDefinition     `json:"clientExtensions,omitempty"`
}

// Order-related Definitions

type OrderTypeDefinition = string

type CancellableOrderTypeDefinition = string

type OrderStateDefinition = string

type OrderStateFilterDefinition = string

type OrderIdentifierDefinition struct {
	OrderID       string `json:"orderID,omitempty"`
	ClientOrderID string `json:"clientOrderID,omitempty"`
}

type OrderSpecifierDefinition = string

type TimeInForceDefinition = string

type OrderPositionFillDefinition = string

type OrderTriggerConditionDefinition = string

// values "TOP_OF_BOOK"
type OrderTriggerModeDefinition = string

type DynamicOrderStateDefinition struct {
	ID                     string               `json:"id,omitempty"`
	TrailingStopValue      PriceValueDefinition `json:"trailingStopValue,omitempty"`
	TriggerDistance        PriceValueDefinition `json:"triggerDistance,omitempty"`
	IsTriggerDistanceExact *bool                `json:"isTriggerDistanceExact,omitempty"`
}

type UnitsAvailableDetailsDefinition struct {
	Long  DecimalNumberDefinition `json:"long,omitempty"`
	Short DecimalNumberDefinition `json:"short,omitempty"`
}

type UnitsAvailableDefinition struct {
	Default     *UnitsAvailableDetailsDefinition `json:"default,omitempty"`
	ReduceFirst *UnitsAvailableDetailsDefinition `json:"reduceFirst,omitempty"`
	ReduceOnly  *UnitsAvailableDetailsDefinition `json:"reduceOnly,omitempty"`
	OpenOnly    *UnitsAvailableDetailsDefinition `json:"openOnly,omitempty"`
}

type GuaranteedStopLossOrderEntryDataDefinition struct {
	MinimumDistance  DecimalNumberDefinition                            `json:"minimumDistance,omitempty"`
	Premium          DecimalNumberDefinition                            `json:"premium,omitempty"`
	LevelRestriction *GuaranteedStopLossOrderLevelRestrictionDefinition `json:"levelRestriction,omitempty"`
}

//
// Trade Definitions
//

type TradeIDDefinition = string

type TradeStateDefinition = string

type TradeStateFilterDefinition = string

type TradeSpecifierDefinition = string

type TradeDefinition struct {
	ID                    TradeIDDefinition                `json:"id,omitempty"`
	Instrument            InstrumentNameDefinition         `json:"instrument,omitempty"`
	Price                 PriceValueDefinition             `json:"price,omitempty"`
	OpenTime              DateTimeDefinition               `json:"openTime,omitempty"`
	State                 TradeStateDefinition             `json:"state,omitempty"`
	InitialUnits          DecimalNumberDefinition          `json:"initialUnits,omitempty"`
	InitialMarginRequired AccountUnitsDefinition           `json:"initialMarginRequired,omitempty"`
	CurrentUnits          DecimalNumberDefinition          `json:"currentUnits,omitempty"`
	RealizedPL            AccountUnitsDefinition           `json:"realizedPL,omitempty"`
	UnrealizedPL          AccountUnitsDefinition           `json:"unrealizedPL,omitempty"`
	MarginUsed            AccountUnitsDefinition           `json:"marginUsed,omitempty"`
	AverageClosePrice     PriceValueDefinition             `json:"averageClosePrice,omitempty"`
	ClosingTransactionIDs []TransactionIDDefinition        `json:"closingTransactionIDs,omitempty"`
	Financing             AccountUnitsDefinition           `json:"financing,omitempty"`
	Dividend              AccountUnitsDefinition           `json:"dividend,omitempty"`
	CloseTime             DateTimeDefinition               `json:"closeTime,omitempty"`
	ClientExtensions      *ClientExtensionsDefinition      `json:"clientExtensions,omitempty"`
	TakeProfitOrder       *TakeProfitOrderDefinition       `json:"takeProfitOrder,omitempty"`
	StopLossOrder         *StopLossOrderDefinition         `json:"stopLossOrder,omitempty"`
	TrailingStopLossOrder *TrailingStopLossOrderDefinition `json:"trailingStopLossOrder,omitempty"`

	DividendAdjustment Undefined `json:"dividendAdjustment,omitempty"`
}

type TradeSummaryDefinition struct {
	ID                      TradeIDDefinition           `json:"id,omitempty"`
	Instrument              InstrumentNameDefinition    `json:"instrument,omitempty"`
	Price                   PriceValueDefinition        `json:"price,omitempty"`
	OpenTime                DateTimeDefinition          `json:"openTime,omitempty"`
	State                   TradeStateDefinition        `json:"state,omitempty"`
	InitialUnits            DecimalNumberDefinition     `json:"initialUnits,omitempty"`
	InitialMarginRequired   AccountUnitsDefinition      `json:"initialMarginRequired,omitempty"`
	CurrentUnits            DecimalNumberDefinition     `json:"currentUnits,omitempty"`
	RealizedPL              AccountUnitsDefinition      `json:"realizedPL,omitempty"`
	UnrealizedPL            AccountUnitsDefinition      `json:"unrealizedPL,omitempty"`
	MarginUsed              AccountUnitsDefinition      `json:"marginUsed,omitempty"`
	AverageClosePrice       PriceValueDefinition        `json:"averageClosePrice,omitempty"`
	ClosingTransactionIDs   []TransactionIDDefinition   `json:"closingTransactionIDs,omitempty"`
	Financing               AccountUnitsDefinition      `json:"financing,omitempty"`
	Dividend                AccountUnitsDefinition      `json:"dividend,omitempty"`
	CloseTime               DateTimeDefinition          `json:"closeTime,omitempty"`
	ClientExtensions        *ClientExtensionsDefinition `json:"clientExtensions,omitempty"`
	TakeProfitOrderID       string                      `json:"takeProfitOrderID,omitempty"`
	StopLossOrderID         string                      `json:"stopLossOrderID,omitempty"`
	TrailingStopLossOrderID string                      `json:"trailingStopLossOrderID,omitempty"`

	DividendAdjustment Undefined `json:"dividendAdjustment,omitempty"`
}

type CalculatedTradeStateDefinition struct {
	ID           TradeIDDefinition      `json:"id,omitempty"`
	UnrealizedPL AccountUnitsDefinition `json:"unrealizedPL,omitempty"`
	MarginUsed   AccountUnitsDefinition `json:"marginUsed,omitempty"`
}

type TradePLDefinition = string

//
// Position Definitions
//

type PositionDefinition struct {
	Instrument              InstrumentNameDefinition `json:"instrument,omitempty"`
	PL                      AccountUnitsDefinition   `json:"pl,omitempty"`
	UnrealizedPL            AccountUnitsDefinition   `json:"unrealizedPL,omitempty"`
	MarginUsed              AccountUnitsDefinition   `json:"marginUsed,omitempty"`
	ResettablePL            AccountUnitsDefinition   `json:"resettablePL,omitempty"`
	Financing               AccountUnitsDefinition   `json:"financing,omitempty"`
	Commission              AccountUnitsDefinition   `json:"commission,omitempty"`
	Dividend                AccountUnitsDefinition   `json:"dividend,omitempty"`
	GuaranteedExecutionFees AccountUnitsDefinition   `json:"guaranteedExecutionFees,omitempty"`
	Long                    *PositionSideDefinition  `json:"long,omitempty"`
	Short                   *PositionSideDefinition  `json:"short,omitempty"`

	DividendAdjustment Undefined `json:"dividendAdjustment,omitempty"`
}

type PositionSideDefinition struct {
	Units                   DecimalNumberDefinition `json:"units,omitempty"`
	AveragePrice            PriceValueDefinition    `json:"averagePrice,omitempty"`
	TradeIDs                []TradeIDDefinition     `json:"tradeIDs,omitempty"`
	PL                      AccountUnitsDefinition  `json:"pl,omitempty"`
	UnrealizedPL            AccountUnitsDefinition  `json:"unrealizedPL,omitempty"`
	ResettablePL            AccountUnitsDefinition  `json:"resettablePL,omitempty"`
	Financing               AccountUnitsDefinition  `json:"financing,omitempty"`
	Dividend                AccountUnitsDefinition  `json:"dividend,omitempty"`
	GuaranteedExecutionFees AccountUnitsDefinition  `json:"guaranteedExecutionFees,omitempty"`

	DividendAdjustment Undefined `json:"dividendAdjustment,omitempty"`
}

type CalculatedPositionStateDefinition struct {
	Instrument        InstrumentNameDefinition `json:"instrument,omitempty"`
	NetUnrealizedPL   AccountUnitsDefinition   `json:"netUnrealizedPL,omitempty"`
	LongUnrealizedPL  AccountUnitsDefinition   `json:"longUnrealizedPL,omitempty"`
	ShortUnrealizedPL AccountUnitsDefinition   `json:"shortUnrealizedPL,omitempty"`
	MarginUsed        AccountUnitsDefinition   `json:"marginUsed,omitempty"`
}

//
// Transaction Definitions
//

// Transactions

// https://developer.oanda.com/rest-live-v20/transaction-df/

type TransactionDefinition struct {
	// #
	// # The Account’s balance after the Order was filled.
	// #
	AccountBalance       AccountUnitsDefinition         `json:"accountBalance,omitempty"`
	AccountFinancingMode AccountFinancingModeDefinition `json:"accountFinancingMode,omitempty"`
	// #
	// # The ID of the Account the Transaction was created for.
	// #
	AccountID     AccountIDDefinition    `json:"accountID,omitempty"`
	AccountNumber *int                   `json:"accountNumber,omitempty"`
	AccountUserID *int                   `json:"accountUserID,omitempty"`
	Alias         string                 `json:"alias,omitempty"`
	Amount        AccountUnitsDefinition `json:"amount,omitempty"`

	// The financing paid or collected when the Order was filled, in the Instrument’s base currency.
	BaseFinancing DecimalNumberDefinition `json:"baseFinancing,omitempty"`
	// #
	// # The ID of the “batch” that the Transaction belongs to. Transactions in
	// # the same batch are applied to the Account simultaneously.
	// #
	BatchID                 TransactionIDDefinition     `json:"batchID,omitempty"`
	CancellingTransactionID TransactionIDDefinition     `json:"cancellingTransactionID,omitempty"`
	ClientExtensions        *ClientExtensionsDefinition `json:"clientExtensions,omitempty"`
	ClientExtensionsModify  *ClientExtensionsDefinition `json:"clientExtensionsModify,omitempty"`
	// #
	// # The client Order ID of the Order filled (only provided if the client has
	// # assigned one).
	// #
	ClientOrderID string `json:"clientOrderID,omitempty"`
	ClientTradeID string `json:"clientTradeID,omitempty"`
	Comment       string `json:"comment,omitempty"`
	// #
	// # The commission charged in the Account’s home currency as a result of
	// # filling the Order. The commission is always represented as a positive
	// # quantity of the Account’s home currency, however it reduces the balance
	// # in the Account.
	// #
	Commission        AccountUnitsDefinition                  `json:"commission,omitempty"`
	DelayedTradeClose *MarketOrderDelayedTradeCloseDefinition `json:"delayedTradeClose,omitempty"`
	Distance          DecimalNumberDefinition                 `json:"distance,omitempty"`
	DivisionID        *int                                    `json:"divisionID,omitempty"`
	ExtensionNumber   *int                                    `json:"extensionNumber,omitempty"`
	// The financing paid or collected when the Order was filled.
	Financing AccountUnitsDefinition `json:"financing,omitempty"`
	// #
	// # The price in effect for the account at the time of the Order fill.
	// #
	FullPrice *ClientPriceDefinition `json:"fullPrice,omitempty"`
	// #
	// # The price that all of the units of the OrderFill should have been filled
	// # at, in the absence of guaranteed price execution. This factors in the
	// # Account’s current ClientPrice, used liquidity and the units of the
	// # OrderFill only. If no Trades were closed with their price clamped for
	// # guaranteed stop loss enforcement, then this value will match the price
	// # fields of each Trade opened, closed, and reduced, and they will all be
	// # the exact same.
	// #
	FullVWAP                      Undefined               `json:"fullVWAP,omitempty"`
	FundingReason                 FundingReasonDefinition `json:"fundingReason,omitempty"`
	GainQuoteHomeConversionFactor Deprecated              `json:"gainQuoteHomeConversionFactor,omitempty"`
	GtdTime                       DateTimeDefinition      `json:"gtdTime,omitempty"`
	Guaranteed                    Deprecated              `json:"guaranteed,omitempty"`

	// The total guaranteed execution fees charged for all Trades opened, closed or reduced with guaranteed Stop Loss Orders.
	GuaranteedExecutionFee AccountUnitsDefinition `json:"guaranteedExecutionFee,omitempty"`

	GuaranteedExecutionPremium Deprecated `json:"guaranteedExecutionPremium,omitempty"`

	// #
	// # The half spread cost for the OrderFill, which is the sum of the
	// # halfSpreadCost values in the tradeOpened, tradesClosed and tradeReduced
	// # fields. This can be a positive or negative value and is represented in
	// # the home currency of the Account.
	// #
	HalfSpreadCost AccountUnitsDefinition `json:"halfSpreadCost,omitempty"`
	HomeCurrency   CurrencyDefinition     `json:"homeCurrency,omitempty"`

	// The HomeConversionFactors in effect at the time of the OrderFill.
	HomeConversionFactors HomeConversionFactorsDefinition `json:"homeConversionFactors,omitempty"`
	// #
	// # The Transaction’s Identifier.
	// #
	ID TransactionIDDefinition `json:"id,omitempty"`
	// #
	// # The name of the filled Order’s instrument.
	// #
	Instrument                    InstrumentNameDefinition               `json:"instrument,omitempty"`
	IntendedReplacesOrderID       string                                 `json:"intendedReplacesOrderID,omitempty"`
	LongPositionCloseout          *MarketOrderPositionCloseoutDefinition `json:"longPositionCloseout,omitempty"`
	LossQuoteHomeConversionFactor Deprecated                             `json:"lossQuoteHomeConversionFactor,omitempty"`
	MarginCloseout                *MarketOrderMarginCloseoutDefinition   `json:"marginCloseout,omitempty"`
	MarginRate                    DecimalNumberDefinition                `json:"marginRate,omitempty"`
	OrderFillTransactionID        TransactionIDDefinition                `json:"orderFillTransactionID,omitempty"`
	// #
	// # The ID of the Order filled.
	// #
	OrderID string `json:"orderID,omitempty"`

	// The profit or loss incurred when the Order was filled.
	PL AccountUnitsDefinition `json:"pl,omitempty"`

	PositionFill       OrderPositionFillDefinition    `json:"positionFill,omitempty"`
	PositionFinancings []*PositionFinancingDefinition `json:"positionFinancings,omitempty"`
	Price              Deprecated                     `json:"price,omitempty"`
	PriceBound         PriceValueDefinition           `json:"priceBound,omitempty"`

	// The financing paid or collected when the Order was filled, in the Instrument’s quote currency.
	QuoteFinancing DecimalNumberDefinition `json:"quoteFinancing,omitempty"`

	// The profit or loss incurred when the Order was filled, in the Instrument’s quote currency.
	QuotePL DecimalNumberDefinition `json:"quotePL,omitempty"`

	// The total guaranteed execution fees charged for all Trades opened, closed
	// or reduced with guaranteed Stop Loss Orders, expressed in the
	// Instrument’s quote currency.
	QuoteGuaranteedExecutionFee DecimalNumberDefinition `json:"quoteGuaranteedExecutionFee,omitempty"`

	// #
	// # The reason that an Order was filled
	// #
	Reason            Reason                            `json:"reason,omitempty"`
	RejectReason      TransactionRejectReasonDefinition `json:"rejectReason,omitempty"`
	ReplacedByOrderID string                            `json:"replacedByOrderID,omitempty"`
	ReplacesOrderID   string                            `json:"replacesOrderID,omitempty"`
	// #
	// # The Request ID of the request which generated the transaction.
	// #
	RequestID             RequestIDDefinition                    `json:"requestID,omitempty"`
	ShortPositionCloseout *MarketOrderPositionCloseoutDefinition `json:"shortPositionCloseout,omitempty"`
	SiteID                *int                                   `json:"siteID,omitempty"`
	StopLossOnFill        *StopLossDetailsDefinition             `json:"stopLossOnFill,omitempty"`
	TakeProfitOnFill      *TakeProfitDetailsDefinition           `json:"takeProfitOnFill,omitempty"`
	// #
	// # The date/time when the Transaction was created.
	// #
	Time                        DateTimeDefinition               `json:"time,omitempty"`
	TimeInForce                 TimeInForceDefinition            `json:"timeInForce,omitempty"`
	TradeClientExtensions       *ClientExtensionsDefinition      `json:"tradeClientExtensions,omitempty"`
	TradeClientExtensionsModify *ClientExtensionsDefinition      `json:"tradeClientExtensionsModify,omitempty"`
	TradeClose                  *MarketOrderTradeCloseDefinition `json:"tradeClose,omitempty"`
	TradeID                     TradeIDDefinition                `json:"tradeID,omitempty"`
	TradeIDs                    TradeIDDefinition                `json:"tradeIDs,omitempty"`
	// #
	// # The Trade that was opened when the Order was filled (only provided if
	// # filling the Order resulted in a new Trade).
	// #
	TradeOpened *TradeOpenDefinition `json:"tradeOpened,omitempty"`
	// #
	// # The Trade that was reduced when the Order was filled (only provided if
	// # filling the Order resulted in reducing an open Trade).
	// #
	TradeReduced *TradeReduceDefinition `json:"tradeReduced,omitempty"`
	TradeState   string                 `json:"tradeState,omitempty"`
	// #
	// # The Trades that were closed when the Order was filled (only provided if
	// # filling the Order resulted in a closing open Trades).
	// #
	TradesClosed           []*TradeReduceDefinition           `json:"tradesClosed,omitempty"`
	TrailingStopLossOnFill *TrailingStopLossDetailsDefinition `json:"trailingStopLossOnFill,omitempty"`
	TriggerCondition       OrderTriggerConditionDefinition    `json:"triggerCondition,omitempty"`
	TriggerMode            OrderTriggerModeDefinition         `json:"triggerMode,omitempty"`
	// #
	// # The Type of the Transaction. Always set to “ORDER_FILL” for an
	// # OrderFillTransaction.
	// #
	Type TransactionTypeDefinition `json:"type,omitempty"`
	// #
	// # The number of units filled by the OrderFill.
	// #
	Units DecimalNumberDefinition `json:"units,omitempty"`
	// #
	// # The ID of the user that initiated the creation of the Transaction.
	// #
	UserID *int `json:"userID,omitempty"`

	RequestedUnits Undefined `json:"requestedUnits,omitempty"`

	PartialFill             Undefined `json:"partialFill,omitempty"`
	TradeCloseTransactionID Undefined `json:"tradeCloseTransactionID,omitempty"`
	ClosedTradeID           Undefined `json:"closedTradeID,omitempty"`
}

// Transaction-related Definitions

type TransactionIDDefinition = string

type TransactionTypeDefinition = string

type FundingReasonDefinition = string

type ClientTagDefinition = string

type ClientCommentDefinition = string

type ClientExtensionsDefinition struct {
	ID      string                  `json:"id,omitempty"`
	Tag     ClientTagDefinition     `json:"tag,omitempty"`
	Comment ClientCommentDefinition `json:"comment,omitempty"`
}

type TakeProfitDetailsDefinition struct {
	Price            PriceValueDefinition        `json:"price,omitempty"`
	TimeInForce      TimeInForceDefinition       `json:"timeInForce,omitempty"`
	GtdTime          DateTimeDefinition          `json:"gtdTime,omitempty"`
	ClientExtensions *ClientExtensionsDefinition `json:"clientExtensions,omitempty"`
}

type StopLossDetailsDefinition struct {
	Price            PriceValueDefinition        `json:"price,omitempty"`
	Distance         DecimalNumberDefinition     `json:"distance,omitempty"`
	TimeInForce      TimeInForceDefinition       `json:"timeInForce,omitempty"`
	GtdTime          DateTimeDefinition          `json:"gtdTime,omitempty"`
	ClientExtensions *ClientExtensionsDefinition `json:"clientExtensions,omitempty"`
	Guaranteed       *bool                       `json:"guaranteed,omitempty"`
}

type TrailingStopLossDetailsDefinition struct {
	Distance         DecimalNumberDefinition     `json:"distance,omitempty"`
	TimeInForce      TimeInForceDefinition       `json:"timeInForce,omitempty"`
	GtdTime          DateTimeDefinition          `json:"gtdTime,omitempty"`
	ClientExtensions *ClientExtensionsDefinition `json:"clientExtensions,omitempty"`
}

type TradeOpenDefinition struct {
	// The ID of the Trade that was opened
	TradeID TradeIDDefinition `json:"tradeID,omitempty"`
	// The number of units opened by the Trade
	Units DecimalNumberDefinition `json:"units,omitempty"`
	// The average price that the units were opened at.
	Price PriceValueDefinition `json:"price,omitempty"`
	// This is the fee charged for opening the trade if it has a guaranteed Stop Loss Order attached to it.
	GuaranteedExecutionFee AccountUnitsDefinition `json:"guaranteedExecutionFee,omitempty"`
	//This is the fee charged for opening the trade if it has a guaranteed Stop Loss Order attached to it, expressed in the Instrument’s quote currency.
	QuoteGuaranteedExecutionFee DecimalNumberDefinition `json:"quoteGuaranteedExecutionFee,omitempty"`
	// The client extensions for the newly opened Trade
	ClientExtensions *ClientExtensionsDefinition `json:"clientExtensions,omitempty"`
	// The half spread cost for the trade open. This can be a positive or negative value and is represented in the home currency of the Account.
	HalfSpreadCost AccountUnitsDefinition `json:"halfSpreadCost,omitempty"`
	// The margin required at the time the Trade was created. Note, this is the ‘pure’ margin required, it is not the ‘effective’ margin used that factors in the trade risk if a GSLO is attached to the trade.
	InitialMarginRequired AccountUnitsDefinition `json:"initialMarginRequired,omitempty"`
}

type HomeConversionFactorsDefinition struct {

	// The ConversionFactor in effect for the Account for converting any gains
	// realized in Instrument quote units into units of the Account’s home
	// currency.

	GainQuoteHome ConversionFactorDefinition `json:"gainQuoteHome,omitempty"`

	// The ConversionFactor in effect for the Account for converting any losses
	// realized in Instrument quote units into units of the Account’s home
	// currency.

	LossQuoteHome ConversionFactorDefinition `json:"lossQuoteHome,omitempty"`

	// The ConversionFactor in effect for the Account for converting any gains
	// realized in Instrument base units into units of the Account’s home
	// currency.

	GainBaseHome ConversionFactorDefinition `json:"gainBaseHome,omitempty"`

	// The ConversionFactor in effect for the Account for converting any losses
	// realized in Instrument base units into units of the Account’s home
	// currency.

	LossBaseHome ConversionFactorDefinition `json:"lossBaseHome,omitempty"`
}

type TradeReduceDefinition struct {

	// The ID of the Trade that was reduced or closed
	TradeID TradeIDDefinition `json:"tradeID,omitempty"`

	// The number of units that the Trade was reduced by
	Units DecimalNumberDefinition `json:"units,omitempty"`

	// The average price that the units were closed at. This price may be
	// clamped for guaranteed Stop Loss Orders.
	Price PriceValueDefinition `json:"price,omitempty"`

	// The PL realized when reducing the Trade
	RealizedPL AccountUnitsDefinition `json:"realizedPL,omitempty"`

	// The financing paid/collected when reducing the Trade
	Financing AccountUnitsDefinition `json:"financing,omitempty"`

	// The base financing paid/collected when reducing the Trade
	BaseFinancing DecimalNumberDefinition `json:"baseFinancing,omitempty"`

	// The quote financing paid/collected when reducing the Trade
	QuoteFinancing DecimalNumberDefinition `json:"quoteFinancing,omitempty"`

	//  The financing rate in effect for the instrument used to calculate the
	//  amount of financing paid/collected when reducing the Trade. This field
	//  will only be set if the AccountFinancingMode at the time of the order
	//  fill is SECOND_BY_SECOND_INSTRUMENT. The value is in decimal rather than
	//  percentage points, e.g. 5% is represented as 0.05.
	FinancingRate DecimalNumberDefinition `json:"financingRate,omitempty"`

	// This is the fee that is charged for closing the Trade if it has a
	// guaranteed Stop Loss Order attached to it.
	GuaranteedExecutionFee AccountUnitsDefinition `json:"guaranteedExecutionFee,omitempty"`

	// This is the fee that is charged for closing the Trade if it has a
	// guaranteed Stop Loss Order attached to it, expressed in the Instrument’s
	// quote currency.
	QuoteGuaranteedExecutionFee DecimalNumberDefinition `json:"quoteGuaranteedExecutionFee,omitempty"`

	// The half spread cost for the trade reduce/close. This can be a positive or negative value and is represented in the home currency of the Account.
	HalfSpreadCost AccountUnitsDefinition `json:"halfSpreadCost,omitempty"`

	ClientTradeID Undefined `json:"clientTradeID,omitempty"`
}

type MarketOrderTradeCloseDefinition struct {
	TradeID       TradeIDDefinition `json:"tradeID,omitempty"`
	ClientTradeID string            `json:"clientTradeID,omitempty"`
	Units         string            `json:"units,omitempty"`
}

type MarketOrderMarginCloseoutDefinition struct {
	Reason MarketOrderMarginCloseoutReasonDefinition `json:"reason,omitempty"`
}

type MarketOrderMarginCloseoutReasonDefinition = string

type MarketOrderDelayedTradeCloseDefinition struct {
	TradeID             TradeIDDefinition       `json:"tradeID,omitempty"`
	ClientTradeID       TradeIDDefinition       `json:"clientTradeID,omitempty"`
	SourceTransactionID TransactionIDDefinition `json:"sourceTransactionID,omitempty"`
}

type MarketOrderPositionCloseoutDefinition struct {
	Instrument InstrumentNameDefinition `json:"instrument,omitempty"`
	Units      string                   `json:"units,omitempty"`
}

type LiquidityRegenerationScheduleDefinition struct {
	Steps []*LiquidityRegenerationScheduleStepDefinition `json:"steps,omitempty"`
}

type LiquidityRegenerationScheduleStepDefinition struct {
	Timestamp        DateTimeDefinition      `json:"timestamp,omitempty"`
	BidLiquidityUsed DecimalNumberDefinition `json:"bidLiquidityUsed,omitempty"`
	AskLiquidityUsed DecimalNumberDefinition `json:"askLiquidityUsed,omitempty"`
}

type OpenTradeFinancingDefinition struct {
	TradeID   TradeIDDefinition      `json:"tradeID,omitempty"`
	Financing AccountUnitsDefinition `json:"financing,omitempty"`
}

type PositionFinancingDefinition struct {
	Instrument          InstrumentNameDefinition        `json:"instrument,omitempty"`
	Financing           AccountUnitsDefinition          `json:"financing,omitempty"`
	OpenTradeFinancings []*OpenTradeFinancingDefinition `json:"openTradeFinancings,omitempty"`
}

type RequestIDDefinition = string

type TransactionRejectReasonDefinition = string

type Reason string

func (p Reason) Description() string {
	switch p {
	case "LIMIT_ORDER":
		return "The Order filled was a Limit Order"
	case "STOP_ORDER":
		return "The Order filled was a Stop Order"
	case "MARKET_IF_TOUCHED_ORDER":
		return "The Order filled was a Market-if-touched Order"
	case "TAKE_PROFIT_ORDER":
		return "The Order filled was a Take Profit Order"
	case "STOP_LOSS_ORDER":
		return "The Order filled was a Stop Loss Order"
	case "GUARANTEED_STOP_LOSS_ORDER":
		return "The Order filled was a Guaranteed Stop Loss Order"
	case "TRAILING_STOP_LOSS_ORDER":
		return "The Order filled was a Trailing Stop Loss Order"
	case "MARKET_ORDER":
		return "The Order filled was a Market Order"
	case "MARKET_ORDER_TRADE_CLOSE":
		return "The Order filled was a Market Order used to explicitly close a Trade"
	case "MARKET_ORDER_POSITION_CLOSEOUT":
		return "The Order filled was a Market Order used to explicitly close a Position"
	case "MARKET_ORDER_MARGIN_CLOSEOUT":
		return "The Order filled was a Market Order used for a Margin Closeout"
	case "MARKET_ORDER_DELAYED_TRADE_CLOSE":
		return "The Order filled was a Market Order used for a delayed Trade close"
	case "FIXED_PRICE_ORDER":
		return "The Order filled was a Fixed Price Order"
	case "FIXED_PRICE_ORDER_PLATFORM_ACCOUNT_MIGRATION":
		return "The Order filled was a Fixed Price Order created as part of a platform account migration"
	case "FIXED_PRICE_ORDER_DIVISION_ACCOUNT_MIGRATION":
		return "The Order filled was a Fixed Price Order created to close a Trade as part of division account migration"
	case "FIXED_PRICE_ORDER_ADMINISTRATIVE_ACTION":
		return "The Order filled was a Fixed Price Order created to close a Trade administratively"
	case "INTERNAL_SERVER_ERROR":
		return "The Order was cancelled because at the time of filling, an unexpected internal server error occurred."
	case "ACCOUNT_LOCKED":
		return "The Order was cancelled because at the time of filling the account was locked."
	case "ACCOUNT_NEW_POSITIONS_LOCKED":
		return "The order was to be filled, however the account is configured to not allow new positions to be created."
	case "ACCOUNT_ORDER_CREATION_LOCKED":
		return "Filling the Order wasn’t possible because it required the creation of a dependent Order and the Account is locked for Order creation."
	case "ACCOUNT_ORDER_FILL_LOCKED":
		return "Filling the Order was not possible because the Account is locked for filling Orders."
	case "CLIENT_REQUEST":
		return "The Order was cancelled explicitly at the request of the client."
	case "MIGRATION":
		return "The Order cancelled because it is being migrated to another account."
	case "MARKET_HALTED":
		return "Filling the Order wasn’t possible because the Order’s instrument was halted."
	case "LINKED_TRADE_CLOSED":
		return "The Order is linked to an open Trade that was closed."
	case "TIME_IN_FORCE_EXPIRED":
		return "The time in force specified for this order has passed."
	case "INSUFFICIENT_MARGIN":
		return "Filling the Order wasn’t possible because the Account had insufficient margin."
	case "FIFO_VIOLATION":
		return "Filling the Order would have resulted in a a FIFO violation."
	case "BOUNDS_VIOLATION":
		return "Filling the Order would have violated the Order’s price bound."
	case "CLIENT_REQUEST_REPLACED":
		return "The Order was cancelled for replacement at the request of the client."
	case "DIVIDEND_ADJUSTMENT_REPLACED":
		return "The Order was cancelled for replacement with an adjusted fillPrice to accommodate for the price movement caused by a dividendAdjustment."
	case "INSUFFICIENT_LIQUIDITY":
		return "Filling the Order wasn’t possible because enough liquidity available."
	case "TAKE_PROFIT_ON_FILL_GTD_TIMESTAMP_IN_PAST":
		return "Filling the Order would have resulted in the creation of a Take Profit Order with a GTD time in the past."
	case "TAKE_PROFIT_ON_FILL_LOSS":
		return "Filling the Order would result in the creation of a Take Profit Order that would have been filled immediately, closing the new Trade at a loss."
	case "LOSING_TAKE_PROFIT":
		return "Filling the Order would result in the creation of a Take Profit Loss Order that would close the new Trade at a loss when filled."
	case "STOP_LOSS_ON_FILL_GTD_TIMESTAMP_IN_PAST":
		return "Filling the Order would have resulted in the creation of a Stop Loss Order with a GTD time in the past."
	case "STOP_LOSS_ON_FILL_LOSS":
		return "Filling the Order would result in the creation of a Stop Loss Order that would have been filled immediately, closing the new Trade at a loss."
	case "STOP_LOSS_ON_FILL_PRICE_DISTANCE_MAXIMUM_EXCEEDED":
		return "Filling the Order would result in the creation of a Stop Loss Order whose price would be zero or negative due to the specified distance."
	case "STOP_LOSS_ON_FILL_REQUIRED":
		return "Filling the Order would not result in the creation of Stop Loss Order, however the Account’s configuration requires that all Trades have a Stop Loss Order attached to them."
	case "STOP_LOSS_ON_FILL_GUARANTEED_REQUIRED":
		return "Filling the Order would not result in the creation of a guaranteed Stop Loss Order, however the Account’s configuration requires that all Trades have a guaranteed Stop Loss Order attached to them."
	case "STOP_LOSS_ON_FILL_GUARANTEED_NOT_ALLOWED":
		return "Filling the Order would result in the creation of a guaranteed Stop Loss Order, however the Account’s configuration does not allow guaranteed Stop Loss Orders."
	case "STOP_LOSS_ON_FILL_GUARANTEED_MINIMUM_DISTANCE_NOT_MET":
		return "Filling the Order would result in the creation of a guaranteed Stop Loss Order with a distance smaller than the configured minimum distance."
	case "STOP_LOSS_ON_FILL_GUARANTEED_LEVEL_RESTRICTION_EXCEEDED":
		return "Filling the Order would result in the creation of a guaranteed Stop Loss Order with trigger price and number of units that that violates the account’s guaranteed Stop Loss Order level restriction."
	case "STOP_LOSS_ON_FILL_GUARANTEED_HEDGING_NOT_ALLOWED":
		return "Filling the Order would result in the creation of a guaranteed Stop Loss Order for a hedged Trade, however the Account’s configuration does not allow guaranteed Stop Loss Orders for hedged Trades/Positions."
	case "STOP_LOSS_ON_FILL_TIME_IN_FORCE_INVALID":
		return "Filling the Order would result in the creation of a Stop Loss Order whose TimeInForce value is invalid. A likely cause would be if the Account requires guaranteed stop loss orders and the TimeInForce value were not GTC."
	case "STOP_LOSS_ON_FILL_TRIGGER_CONDITION_INVALID":
		return "Filling the Order would result in the creation of a Stop Loss Order whose TriggerCondition value is invalid. A likely cause would be if the stop loss order is guaranteed and the TimeInForce is not TRIGGER_DEFAULT or TRIGGER_BID for a long trade, or not TRIGGER_DEFAULT or TRIGGER_ASK for a short trade."
	case "GUARANTEED_STOP_LOSS_ON_FILL_GTD_TIMESTAMP_IN_PAST":
		return "Filling the Order would have resulted in the creation of a Guaranteed Stop Loss Order with a GTD time in the past."
	case "GUARANTEED_STOP_LOSS_ON_FILL_LOSS":
		return "Filling the Order would result in the creation of a Guaranteed Stop Loss Order that would have been filled immediately, closing the new Trade at a loss."
	case "GUARANTEED_STOP_LOSS_ON_FILL_PRICE_DISTANCE_MAXIMUM_EXCEEDED":
		return "Filling the Order would result in the creation of a Guaranteed Stop Loss Order whose price would be zero or negative due to the specified distance."
	case "GUARANTEED_STOP_LOSS_ON_FILL_REQUIRED":
		return "Filling the Order would not result in the creation of a Guaranteed Stop Loss Order, however the Account’s configuration requires that all Trades have a Guaranteed Stop Loss Order attached to them."
	case "GUARANTEED_STOP_LOSS_ON_FILL_NOT_ALLOWED":
		return "Filling the Order would result in the creation of a Guaranteed Stop Loss Order, however the Account’s configuration does not allow Guaranteed Stop Loss Orders."
	case "GUARANTEED_STOP_LOSS_ON_FILL_MINIMUM_DISTANCE_NOT_MET":
		return "Filling the Order would result in the creation of a Guaranteed Stop Loss Order with a distance smaller than the configured minimum distance."
	case "GUARANTEED_STOP_LOSS_ON_FILL_LEVEL_RESTRICTION_VOLUME_EXCEEDED":
		return "Filling the Order would result in the creation of a Guaranteed Stop Loss Order with trigger number of units that violates the account’s Guaranteed Stop Loss Order level restriction volume."
	case "GUARANTEED_STOP_LOSS_ON_FILL_LEVEL_RESTRICTION_PRICE_RANGE_EXCEEDED":
		return "Filling the Order would result in the creation of a Guaranteed Stop Loss Order with trigger price that violates the account’s Guaranteed Stop Loss Order level restriction price range."
	case "GUARANTEED_STOP_LOSS_ON_FILL_HEDGING_NOT_ALLOWED":
		return "Filling the Order would result in the creation of a Guaranteed Stop Loss Order for a hedged Trade, however the Account’s configuration does not allow Guaranteed Stop Loss Orders for hedged Trades/Positions."
	case "GUARANTEED_STOP_LOSS_ON_FILL_TIME_IN_FORCE_INVALID":
		return "Filling the Order would result in the creation of a Guaranteed Stop Loss Order whose TimeInForce value is invalid. A likely cause would be if the Account requires guaranteed stop loss orders and the TimeInForce value were not GTC."
	case "GUARANTEED_STOP_LOSS_ON_FILL_TRIGGER_CONDITION_INVALID":
		return "Filling the Order would result in the creation of a Guaranteed Stop Loss Order whose TriggerCondition value is invalid. A likely cause would be the TimeInForce is not TRIGGER_DEFAULT or TRIGGER_BID for a long trade, or not TRIGGER_DEFAULT or TRIGGER_ASK for a short trade."
	case "TAKE_PROFIT_ON_FILL_PRICE_DISTANCE_MAXIMUM_EXCEEDED":
		return "Filling the Order would result in the creation of a Take Profit Order whose price would be zero or negative due to the specified distance."
	case "TRAILING_STOP_LOSS_ON_FILL_GTD_TIMESTAMP_IN_PAST":
		return "Filling the Order would have resulted in the creation of a Trailing Stop Loss Order with a GTD time in the past."
	case "CLIENT_TRADE_ID_ALREADY_EXISTS":
		return "Filling the Order would result in the creation of a new Open Trade with a client Trade ID already in use."
	case "POSITION_CLOSEOUT_FAILED":
		return "Closing out a position wasn’t fully possible."
	case "OPEN_TRADES_ALLOWED_EXCEEDED":
		return "Filling the Order would cause the maximum open trades allowed for the Account to be exceeded."
	case "PENDING_ORDERS_ALLOWED_EXCEEDED":
		return "Filling the Order would have resulted in exceeding the number of pending Orders allowed for the Account."
	case "TAKE_PROFIT_ON_FILL_CLIENT_ORDER_ID_ALREADY_EXISTS":
		return "Filling the Order would have resulted in the creation of a Take Profit Order with a client Order ID that is already in use."
	case "STOP_LOSS_ON_FILL_CLIENT_ORDER_ID_ALREADY_EXISTS":
		return "Filling the Order would have resulted in the creation of a Stop Loss Order with a client Order ID that is already in use."
	case "GUARANTEED_STOP_LOSS_ON_FILL_CLIENT_ORDER_ID_ALREADY_EXISTS":
		return "Filling the Order would have resulted in the creation of a Guaranteed Stop Loss Order with a client Order ID that is already in use."
	case "TRAILING_STOP_LOSS_ON_FILL_CLIENT_ORDER_ID_ALREADY_EXISTS":
		return "Filling the Order would have resulted in the creation of a Trailing Stop Loss Order with a client Order ID that is already in use."
	case "POSITION_SIZE_EXCEEDED":
		return "Filling the Order would have resulted in the Account’s maximum position size limit being exceeded for the Order’s instrument."
	case "HEDGING_GSLO_VIOLATION":
		return "Filling the Order would result in the creation of a Trade, however there already exists an opposing (hedged) Trade that has a guaranteed Stop Loss Order attached to it. Guaranteed Stop Loss Orders cannot be combined with hedged positions."
	case "ACCOUNT_POSITION_VALUE_LIMIT_EXCEEDED":
		return "Filling the order would cause the maximum position value allowed for the account to be exceeded. The Order has been cancelled as a result."
	case "INSTRUMENT_BID_REDUCE_ONLY":
		return "Filling the order would require the creation of a short trade, however the instrument is configured such that orders being filled using bid prices can only reduce existing positions. New short positions cannot be created, but existing long positions may be reduced or closed."
	case "INSTRUMENT_ASK_REDUCE_ONLY":
		return "Filling the order would require the creation of a long trade, however the instrument is configured such that orders being filled using ask prices can only reduce existing positions. New long positions cannot be created, but existing short positions may be reduced or closed."
	case "INSTRUMENT_BID_HALTED":
		return "Filling the order would require using the bid, however the instrument is configured such that the bids are halted, and so no short orders may be filled."
	case "INSTRUMENT_ASK_HALTED":
		return "Filling the order would require using the ask, however the instrument is configured such that the asks are halted, and so no long orders may be filled."
	case "STOP_LOSS_ON_FILL_GUARANTEED_BID_HALTED":
		return "Filling the Order would result in the creation of a Guaranteed Stop Loss Order (GSLO). Since the trade is long the GSLO would be short, however the bid side is currently halted. GSLOs cannot be created in this situation."
	case "STOP_LOSS_ON_FILL_GUARANTEED_ASK_HALTED":
		return "Filling the Order would result in the creation of a Guaranteed Stop Loss Order (GSLO). Since the trade is short the GSLO would be long, however the ask side is currently halted. GSLOs cannot be created in this situation."
	case "GUARANTEED_STOP_LOSS_ON_FILL_BID_HALTED":
		return "Filling the Order would result in the creation of a Guaranteed Stop Loss Order (GSLO). Since the trade is long the GSLO would be short, however the bid side is currently halted. GSLOs cannot be created in this situation."
	case "GUARANTEED_STOP_LOSS_ON_FILL_ASK_HALTED":
		return "Filling the Order would result in the creation of a Guaranteed Stop Loss Order (GSLO). Since the trade is short the GSLO would be long, however the ask side is currently halted. GSLOs cannot be created in this situation."
	case "FIFO_VIOLATION_SAFEGUARD_VIOLATION":
		return "Filling the Order would have resulted in a new Trade that violates the FIFO violation safeguard constraints."
	case "FIFO_VIOLATION_SAFEGUARD_PARTIAL_CLOSE_VIOLATION":
		return "Filling the Order would have reduced an existing Trade such that the reduced Trade violates the FIFO violation safeguard constraints."
	case "ORDERS_ON_FILL_RMO_MUTUAL_EXCLUSIVITY_MUTUALLY_EXCLUSIVE_VIOLATION":
		return "The Orders on fill would be in violation of the risk management Order mutual exclusivity configuration specifying that only one risk management Order can be attached to a Trade."
	}

	return string(p)
}

func (p Reason) IsInsufficientLiquidity() bool {
	return p == "INSUFFICIENT_LIQUIDITY"
}

func (p Reason) IsInsufficientMargin() bool {
	return p == "INSUFFICIENT_MARGIN"
}

type TransactionFilterDefinition = string

const (
	OrderTransaction                             TransactionFilterDefinition = "ORDER"                                 //	Order-related Transactions. These are the Transactions that create, cancel, fill or trigger Orders
	FundingTransaction                           TransactionFilterDefinition = "FUNDING"                               //	Funding-related Transactions
	AdminTransaction                             TransactionFilterDefinition = "ADMIN"                                 //	Administrative Transactions
	CreateTransaction                            TransactionFilterDefinition = "CREATE"                                //	Account Create Transaction
	CloseTransaction                             TransactionFilterDefinition = "CLOSE"                                 //	Account Close Transaction
	ReopenTransaction                            TransactionFilterDefinition = "REOPEN"                                //	Account Reopen Transaction
	ClientConfigureTransaction                   TransactionFilterDefinition = "CLIENT_CONFIGURE"                      //	Client Configuration Transaction
	ClientConfigureRejectTransaction             TransactionFilterDefinition = "CLIENT_CONFIGURE_REJECT"               //	Client Configuration Reject Transaction
	TransferFundsTransaction                     TransactionFilterDefinition = "TRANSFER_FUNDS"                        //	Transfer Funds Transaction
	TransferFundsRejectTransaction               TransactionFilterDefinition = "TRANSFER_FUNDS_REJECT"                 //	Transfer Funds Reject Transaction
	MarketOrderTransaction                       TransactionFilterDefinition = "MARKET_ORDER"                          //	Market Order Transaction
	MarketOrderRejectTransaction                 TransactionFilterDefinition = "MARKET_ORDER_REJECT"                   //	Market Order Reject Transaction
	LimitOrderTransaction                        TransactionFilterDefinition = "LIMIT_ORDER"                           //	Limit Order Transaction
	LimitOrderRejectTransaction                  TransactionFilterDefinition = "LIMIT_ORDER_REJECT"                    //	Limit Order Reject Transaction
	StopOrderTransaction                         TransactionFilterDefinition = "STOP_ORDER"                            //	Stop Order Transaction
	StopOrderRejectTransaction                   TransactionFilterDefinition = "STOP_ORDER_REJECT"                     //	Stop Order Reject Transaction
	MarketIfTouchedOrderTransaction              TransactionFilterDefinition = "MARKET_IF_TOUCHED_ORDER"               //	Market if Touched Order Transaction
	MarketIfTouchedOrderRejectTransaction        TransactionFilterDefinition = "MARKET_IF_TOUCHED_ORDER_REJECT"        //	Market if Touched Order Reject Transaction
	TakeProfitOrderTransaction                   TransactionFilterDefinition = "TAKE_PROFIT_ORDER"                     //	Take Profit Order Transaction
	TakeProfitOrderRejectTransaction             TransactionFilterDefinition = "TAKE_PROFIT_ORDER_REJECT"              //	Take Profit Order Reject Transaction
	StopLossOrderTransaction                     TransactionFilterDefinition = "STOP_LOSS_ORDER"                       //	Stop Loss Order Transaction
	StopLossOrderRejectTransaction               TransactionFilterDefinition = "STOP_LOSS_ORDER_REJECT"                //	Stop Loss Order Reject Transaction
	TrailingStopLossOrderTransaction             TransactionFilterDefinition = "TRAILING_STOP_LOSS_ORDER"              //	Trailing Stop Loss Order Transaction
	TrailingStopLossOrderRejectTransaction       TransactionFilterDefinition = "TRAILING_STOP_LOSS_ORDER_REJECT"       //	Trailing Stop Loss Order Reject Transaction
	OneCancelsAllOrderTransaction                TransactionFilterDefinition = "ONE_CANCELS_ALL_ORDER"                 //	One Cancels All Order Transaction
	OneCancelsAllOrderRejectTransaction          TransactionFilterDefinition = "ONE_CANCELS_ALL_ORDER_REJECT"          //	One Cancels All Order Reject Transaction
	OneCancelsAllOrderTriggeredTransaction       TransactionFilterDefinition = "ONE_CANCELS_ALL_ORDER_TRIGGERED"       //	One Cancels All Order Trigger Transaction
	OrderFillTransaction                         TransactionFilterDefinition = "ORDER_FILL"                            //	Order Fill Transaction
	OrderCancelTransaction                       TransactionFilterDefinition = "ORDER_CANCEL"                          //	Order Cancel Transaction
	OrderCancelRejectTransaction                 TransactionFilterDefinition = "ORDER_CANCEL_REJECT"                   //	Order Cancel Reject Transaction
	OrderClientExtensionsModifyTransaction       TransactionFilterDefinition = "ORDER_CLIENT_EXTENSIONS_MODIFY"        //	Order Client Extensions Modify Transaction
	OrderClientExtensionsModifyRejectTransaction TransactionFilterDefinition = "ORDER_CLIENT_EXTENSIONS_MODIFY_REJECT" //	Order Client Extensions Modify Reject Transaction
	TradeClientExtensionsModifyTransaction       TransactionFilterDefinition = "TRADE_CLIENT_EXTENSIONS_MODIFY"        //	Trade Client Extensions Modify Transaction
	TradeClientExtensionsModifyRejectTransaction TransactionFilterDefinition = "TRADE_CLIENT_EXTENSIONS_MODIFY_REJECT" //	Trade Client Extensions Modify Reject Transaction
	MarginCallEnterTransaction                   TransactionFilterDefinition = "MARGIN_CALL_ENTER"                     //	Margin Call Enter Transaction
	MarginCallExtendTransaction                  TransactionFilterDefinition = "MARGIN_CALL_EXTEND"                    //	Margin Call Extend Transaction
	MarginCallExitTransaction                    TransactionFilterDefinition = "MARGIN_CALL_EXIT"                      //	Margin Call Exit Transaction
	DelayedTradeClosureTransaction               TransactionFilterDefinition = "DELAYED_TRADE_CLOSURE"                 //	Delayed Trade Closure Transaction
	DailyFinancingTransaction                    TransactionFilterDefinition = "DAILY_FINANCING"                       //	Daily Financing Transaction
	ResetResettablePLTransaction                 TransactionFilterDefinition = "RESET_RESETTABLE_PL"                   //	Reset Resettable PL Transaction
)

type TransactionHeartbeatDefinition struct {
	Type              string                  `json:"type,omitempty"`
	LastTransactionID TransactionIDDefinition `json:"lastTransactionID,omitempty"`
	Time              DateTimeDefinition      `json:"time,omitempty"`
}

//
// Pricing Definitions
//

type PriceDefinition struct {

	// The string “PRICE”. Used to identify the a Price object when found in a stream.
	Type string `json:"type,omitempty"`

	// The Price’s Instrument.
	Instrument InstrumentNameDefinition `json:"instrument,omitempty"`

	// The date/time when the Price was created
	Time DateTimeDefinition `json:"time,omitempty"`

	// The status of the Price.
	// Deprecated: Will be removed in a future API update.
	Status PriceStatusDefinition `json:"status,omitempty"`

	// Flag indicating if the Price is tradeable or not
	Tradeable *bool `json:"tradeable,omitempty"`

	// The list of prices and liquidity available on the Instrument’s bid side.
	// It is possible for this list to be empty if there is no bid liquidity
	// currently available for the Instrument in the Account.
	Bids []*PriceBucketDefinition `json:"bids,omitempty"`

	// The list of prices and liquidity available on the Instrument’s ask side.
	// It is possible for this list to be empty if there is no ask liquidity
	// currently available for the Instrument in the Account.
	Asks []*PriceBucketDefinition `json:"asks,omitempty"`

	// The closeout bid Price. This Price is used when a bid is required to
	// closeout a Position (margin closeout or manual) yet there is no bid
	// liquidity. The closeout bid is never used to open a new position.
	CloseoutBid PriceValueDefinition `json:"closeoutBid,omitempty"`

	// The closeout ask Price. This Price is used when a ask is required to
	// closeout a Position (margin closeout or manual) yet there is no ask
	// liquidity. The closeout ask is never used to open a new position.
	CloseoutAsk PriceValueDefinition `json:"closeoutAsk,omitempty"`

	// The factors used to convert quantities of this price’s Instrument’s quote
	// currency into a quantity of the Account’s home currency. When the
	// includeHomeConversions is present in the pricing request (regardless of
	// its value), this field will not be present.
	// Deprecated: Will be removed in a future API update.
	QuoteHomeConversionFactors *QuoteHomeConversionFactorsDefinition `json:"quoteHomeConversionFactors,omitempty"`

	// Representation of how many units of an Instrument are available to be
	// traded by an Order depending on its positionFill option.
	// Deprecated: Will be removed in a future API update.
	UnitsAvailable *UnitsAvailableDefinition `json:"unitsAvailable,omitempty"`
}

type PriceValueDefinition = string

type PriceBucketDefinition struct {
	Price     PriceValueDefinition `json:"price,omitempty"`
	Liquidity json.Number          `json:"liquidity,omitempty"`
}

type PriceStatusDefinition = string

type QuoteHomeConversionFactorsDefinition struct {
	PositiveUnits DecimalNumberDefinition `json:"positiveUnits,omitempty"`
	NegativeUnits DecimalNumberDefinition `json:"negativeUnits,omitempty"`
}

type HomeConversionsDefinition struct {
	Currency      CurrencyDefinition      `json:"currency,omitempty"`
	AccountGain   DecimalNumberDefinition `json:"accountGain,omitempty"`
	AccountLoss   DecimalNumberDefinition `json:"accountLoss,omitempty"`
	PositionValue DecimalNumberDefinition `json:"positionValue,omitempty"`
}

type ClientPriceDefinition struct {
	Bids        []*PriceBucketDefinition `json:"bids,omitempty"`
	Asks        []*PriceBucketDefinition `json:"asks,omitempty"`
	CloseoutBid PriceValueDefinition     `json:"closeoutBid,omitempty"`
	CloseoutAsk PriceValueDefinition     `json:"closeoutAsk,omitempty"`
	Timestamp   DateTimeDefinition       `json:"timestamp,omitempty"`
}

type PricingHeartbeatDefinition struct {
	Type string             `json:"type,omitempty"`
	Time DateTimeDefinition `json:"time,omitempty"`
}

type ConversionFactorDefinition struct {
	Factor DecimalNumberDefinition `json:"factor,omitempty"`
}

//
// Primitives Definitions
//

type DecimalNumberDefinition = string

type AccountUnitsDefinition = string

type CurrencyDefinition = string

type TagDefinition struct {
	Type string `json:"type,omitempty"`
	Name string `json:"name,omitempty"`
}

type InstrumentNameDefinition = string

type InstrumentTypeDefinition = string

type InstrumentDefinition struct {
	// The name of the Instrument
	Name InstrumentNameDefinition `json:"name,omitempty"`

	// The type of the Instrument
	Type InstrumentTypeDefinition `json:"type,omitempty"`

	// The display name of the Instrument
	DisplayName string `json:"displayName,omitempty"`

	// The amount that is charged to the account if a guaranteed Stop Loss Order
	// is triggered and filled. The value is in price units and is charged for
	// each unit of the Trade. This field will only be present if the Account’s
	// guaranteedStopLossOrderMode for this Instrument is not ‘DISABLED’.
	GuaranteedStopLossOrderExecutionPremium DecimalNumberDefinition `json:"guaranteedStopLossOrderExecutionPremium,omitempty"`

	// The guaranteed Stop Loss Order level restriction for this instrument.
	// This field will only be present if the Account’s
	// guaranteedStopLossOrderMode for this Instrument is not ‘DISABLED’.
	GuaranteedStopLossOrderLevelRestriction GuaranteedStopLossOrderLevelRestrictionDefinition `json:"guaranteedStopLossOrderLevelRestriction,omitempty"`

	// The current Guaranteed Stop Loss Order mode of the Account for this
	// Instrument.
	GuaranteedStopLossOrderMode GuaranteedStopLossOrderModeForInstrumentDefinition `json:"guaranteedStopLossOrderMode,omitempty"`

	// The location of the “pip” for this instrument. The decimal position of
	// the pip in this Instrument’s price can be found at 10 ^ pipLocation (e.g.
	// -4 pipLocation results in a decimal pip position of 10 ^ -4 = 0.0001).
	PipLocation *int `json:"pipLocation,omitempty"`

	// The number of decimal places that should be used to display prices for
	// this instrument. (e.g. a displayPrecision of 5 would result in a price of
	// “1” being displayed as “1.00000”)
	DisplayPrecision *int `json:"displayPrecision,omitempty"`

	// The amount of decimal places that may be provided when specifying the
	// number of units traded for this instrument.
	TradeUnitsPrecision *int `json:"tradeUnitsPrecision,omitempty"`

	// The minimum distance allowed between the Trade’s fill price and the
	// configured price for guaranteed Stop Loss Orders created for this
	// instrument. Specified in price units.
	MinimumGuaranteedStopLossDistance DecimalNumberDefinition `json:"minimumGuaranteedStopLossDistance,omitempty"`

	// The smallest number of units allowed to be traded for this instrument.
	MinimumTradeSize DecimalNumberDefinition `json:"minimumTradeSize,omitempty"`

	// The maximum trailing stop distance allowed for a trailing stop loss
	// created for this instrument. Specified in price units.
	MaximumTrailingStopDistance DecimalNumberDefinition `json:"maximumTrailingStopDistance,omitempty"`

	// The minimum trailing stop distance allowed for a trailing stop loss
	// created for this instrument. Specified in price units.
	MinimumTrailingStopDistance DecimalNumberDefinition `json:"minimumTrailingStopDistance,omitempty"`

	// The maximum position size allowed for this instrument. Specified in
	// units.
	MaximumPositionSize DecimalNumberDefinition `json:"maximumPositionSize,omitempty"`

	// The maximum units allowed for an Order placed for this instrument.
	// Specified in units.
	MaximumOrderUnits DecimalNumberDefinition `json:"maximumOrderUnits,omitempty"`

	// The margin rate for this instrument.
	MarginRate DecimalNumberDefinition `json:"marginRate,omitempty"`

	// The commission structure for this instrument.
	Commission *InstrumentCommissionDefinition `json:"commission,omitempty"`

	// The tags associated with this instrument.
	Tags []*TagDefinition `json:"tags,omitempty"`

	// Financing data for this instrument.
	Financing InstrumentFinancingDefinition `json:"financing,omitempty"`
}

type DateTimeDefinition = string

type AcceptDatetimeFormatDefinition = string

type InstrumentCommissionDefinition struct {
	Commission        DecimalNumberDefinition `json:"commission,omitempty"`
	UnitsTraded       DecimalNumberDefinition `json:"unitsTraded,omitempty"`
	MinimumCommission DecimalNumberDefinition `json:"minimumCommission,omitempty"`
}

type FinancingDayOfWeekDefinition struct {
	// The day of the week to charge the financing.
	DayOfWeek DayOfWeekDefinition `json:"dayOfWeek,omitempty"`

	// The number of days worth of financing to be charged on dayOfWeek.
	DaysCharged *int `json:"daysCharged,omitempty"`
}

type InstrumentFinancingDefinition struct {

	// The financing rate to be used for a long position for the instrument. The
	// value is in decimal rather than percentage points, i.e. 5% is represented
	// as 0.05.
	LongRate DecimalNumberDefinition `json:"longRate,omitempty"`

	// The financing rate to be used for a short position for the instrument.
	// The value is in decimal rather than percentage points, i.e. 5% is
	// represented as 0.05.
	ShortRate DecimalNumberDefinition `json:"shortRate,omitempty"`

	// The days of the week to debit or credit financing charges; the exact time
	// of day at which to charge the financing is set in the
	// DivisionTradingGroup for the client’s account.

	FinancingDaysOfWeek []FinancingDayOfWeekDefinition `json:"financingDaysOfWeek,omitempty"`
}

type GuaranteedStopLossOrderLevelRestrictionDefinition struct {
	// Applies to Trades with a guaranteed Stop Loss Order attached for the
	// specified Instrument. This is the total allowed Trade volume that can
	// exist within the priceRange based on the trigger prices of the guaranteed
	// Stop Loss Orders.
	Volume DecimalNumberDefinition `json:"volume,omitempty"`

	// The price range the volume applies to. This value is in price units.
	PriceRange DecimalNumberDefinition `json:"priceRange,omitempty"`
}

// DISABLED 	The Account is not permitted to create Guaranteed Stop Loss Orders for this Instrument.
// ALLOWED 	The Account is able, but not required to have Guaranteed Stop Loss Orders for open Trades for this Instrument.
// REQUIRED 	The Account is required to have Guaranteed Stop Loss Orders for all open Trades for this Instrument.
type GuaranteedStopLossOrderModeForInstrumentDefinition = string

// SUNDAY 	Sunday
// MONDAY 	Monday
// TUESDAY 	Tuesday
// WEDNESDAY 	Wednesday
// THURSDAY 	Thursday
// FRIDAY 	Friday
// SATURDAY 	Saturday
type DayOfWeekDefinition = string

type DirectionDefinition = string
