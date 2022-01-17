package oanda

import (
	"context"

	"github.com/pkg/errors"
)

/* Receivers */

type ReceiverPositions struct {
	AccountID  string
	Connection *Connection
}

func (r *ReceiverAccountID) Positions() *ReceiverPositions {
	return &ReceiverPositions{
		AccountID:  r.AccountID,
		Connection: r.Connection,
	}
}

type ReceiverOpenPositions struct {
	AccountID  string
	Connection *Connection
}

func (r *ReceiverAccountID) OpenPositions() *ReceiverOpenPositions {
	return &ReceiverOpenPositions{
		AccountID:  r.AccountID,
		Connection: r.Connection,
	}
}

type ReceiverPositionsInstrument struct {
	AccountID  string
	Connection *Connection
	Instrument string
}

func (r *ReceiverPositions) Instrument(instrument string) *ReceiverPositionsInstrument {
	return &ReceiverPositionsInstrument{
		AccountID:  r.AccountID,
		Connection: r.Connection,
		Instrument: instrument,
	}
}

type ReceiverPositionsInstrumentClose struct {
	AccountID  string
	Connection *Connection
	Instrument string
}

func (r *ReceiverPositionsInstrument) Close() *ReceiverPositionsInstrumentClose {
	return &ReceiverPositionsInstrumentClose{
		AccountID:  r.AccountID,
		Connection: r.Connection,
		Instrument: r.Instrument,
	}
}

/* Params */

type PutPositionsInstrumentCloseBodyParams struct {
	// Indication of how much of the long Position to closeout. Either the
	// string “ALL”, the string “NONE”, or a DecimalNumber representing how many
	// units of the long position to close using a PositionCloseout MarketOrder.
	// The units specified must always be positive.
	// default=ALL
	LongUnits string `json:"longUnits,omitempty"`

	// The client extensions to add to the MarketOrder used to close the long
	// position.
	LongClientExtensions *ClientExtensionsDefinition `json:"longClientExtensions,omitempty"`

	// Indication of how much of the short Position to closeout. Either the
	// string “ALL”, the string “NONE”, or a DecimalNumber representing how many
	// units of the short position to close using a PositionCloseout
	// MarketOrder. The units specified must always be positive.
	// default=ALL
	ShortUnits string `json:"shortUnits,omitempty"`

	// The client extensions to add to the MarketOrder used to close the short
	// position.
	ShortClientExtensions *ClientExtensionsDefinition `json:"shortClientExtensions,omitempty"`
}

type PutPositionsInstrumentCloseParams struct {
	Body *PutPositionsInstrumentCloseBodyParams
}

/* Schemas */

type GetPositionsSchema struct {
	Positions         []*PositionDefinition   `json:"positions,omitempty"`
	LastTransactionID TransactionIDDefinition `json:"lastTransactionID,omitempty"`
}

type GetOpenPositionsSchema struct {
	Positions         []*PositionDefinition   `json:"positions,omitempty"`
	LastTransactionID TransactionIDDefinition `json:"lastTransactionID,omitempty"`
}

type GetPositionsInstrumentSchema struct {
	Position          *PositionDefinition     `json:"position,omitempty"`
	LastTransactionID TransactionIDDefinition `json:"lastTransactionID,omitempty"`
}

type PutPositionsInstrumentCloseSchema struct {
	LongOrderCreateTransaction  *TransactionDefinition    `json:"longOrderCreateTransaction,omitempty"`
	LongOrderFillTransaction    *TransactionDefinition    `json:"longOrderFillTransaction,omitempty"`
	LongOrderCancelTransaction  *TransactionDefinition    `json:"longOrderCancelTransaction,omitempty"`
	ShortOrderCreateTransaction *TransactionDefinition    `json:"shortOrderCreateTransaction,omitempty"`
	ShortOrderFillTransaction   *TransactionDefinition    `json:"shortOrderFillTransaction,omitempty"`
	ShortOrderCancelTransaction *TransactionDefinition    `json:"shortOrderCancelTransaction,omitempty"`
	RelatedTransactionIDs       []TransactionIDDefinition `json:"relatedTransactionIDs,omitempty"`
	LastTransactionID           TransactionIDDefinition   `json:"lastTransactionID,omitempty"`
}

func (p *PutPositionsInstrumentCloseSchema) ToTradeResult() *TradeResult {
	res := &TradeResult{}

	if p.LongOrderFillTransaction != nil {
		res.OrderFillTransaction = p.LongOrderFillTransaction
	}

	if p.ShortOrderFillTransaction != nil {
		res.OrderFillTransaction = p.ShortOrderFillTransaction
	}

	if p.LongOrderCancelTransaction != nil {
		res.OrderCancelTransaction = p.LongOrderCancelTransaction
	}

	if p.ShortOrderCancelTransaction != nil {
		res.OrderCancelTransaction = p.ShortOrderCancelTransaction
	}

	return res
}

/* Errors */

type PutPositionsInstrumentCloseBadRequestError struct {
	LongOrderRejectTransaction  *TransactionDefinition    `json:"longOrderRejectTransaction,omitempty"`
	ShortOrderRejectTransaction *TransactionDefinition    `json:"shortOrderRejectTransaction,omitempty"`
	RelatedTransactionIDs       []TransactionIDDefinition `json:"relatedTransactionIDs,omitempty"`
	LastTransactionID           TransactionIDDefinition   `json:"lastTransactionID,omitempty"`
	ErrorCode                   string                    `json:"errorCode,omitempty"`
	ErrorMessage                string                    `json:"errorMessage,omitempty"`
}

func (r *PutPositionsInstrumentCloseBadRequestError) Error() string {
	// TODO: エラーを整える
	return r.ErrorMessage
}

type PutPositionsInstrumentCloseNotFoundError struct {
	LongOrderRejectTransaction  *TransactionDefinition    `json:"longOrderRejectTransaction,omitempty"`
	ShortOrderRejectTransaction *TransactionDefinition    `json:"shortOrderRejectTransaction,omitempty"`
	RelatedTransactionIDs       []TransactionIDDefinition `json:"relatedTransactionIDs,omitempty"`
	LastTransactionID           TransactionIDDefinition   `json:"lastTransactionID,omitempty"`
	ErrorCode                   string                    `json:"errorCode,omitempty"`
	ErrorMessage                string                    `json:"errorMessage,omitempty"`
}

func (r *PutPositionsInstrumentCloseNotFoundError) Error() string {
	// TODO: エラーを整える
	return r.ErrorMessage
}

/* API */

// GET /v3/accounts/{accountID}/positions
func (r *ReceiverPositions) Get(ctx context.Context) (*GetPositionsSchema, error) {
	childCtx, cancel := context.WithCancel(ctx)
	defer cancel()

	resp, err := r.Connection.request(
		childCtx,
		&requestParams{
			method:   "GET",
			endPoint: "/v3/accounts/" + r.AccountID + "/positions",
			headers: []header{
				{key: "Accept-Datetime-Format", value: "RFC3339"},
			},
		},
	)
	if err != nil {
		return nil, errors.Errorf("Get positions canceled: %v", err)
	}
	defer resp.Body.Close()

	var data interface{}
	switch resp.StatusCode {
	case 200:
		data = new(GetPositionsSchema)
	}

	data, err = parseResponse(resp, data, r.Connection.Strict)
	if err != nil {
		return nil, errors.Errorf("Get positions failed: %v", err)
	}
	return data.(*GetPositionsSchema), nil
}

// GET /v3/accounts/{accountID}/openPositions
func (r *ReceiverOpenPositions) Get(ctx context.Context) (*GetOpenPositionsSchema, error) {
	childCtx, cancel := context.WithCancel(ctx)
	defer cancel()

	resp, err := r.Connection.request(
		childCtx,
		&requestParams{
			method:   "GET",
			endPoint: "/v3/accounts/" + r.AccountID + "/openPositions",
			headers: []header{
				{key: "Accept-Datetime-Format", value: "RFC3339"},
			},
		},
	)
	if err != nil {
		return nil, errors.Errorf("Get open positions canceled: %v", err)
	}
	defer resp.Body.Close()

	var data interface{}
	switch resp.StatusCode {
	case 200:
		data = new(GetOpenPositionsSchema)
	}

	data, err = parseResponse(resp, data, r.Connection.Strict)
	if err != nil {
		return nil, errors.Errorf("Get open positions failed: %v", err)
	}
	return data.(*GetOpenPositionsSchema), nil
}

// GET /v3/accounts/{accountID}/positions/{instrument}
func (r *ReceiverPositionsInstrument) Get(ctx context.Context) (*GetPositionsInstrumentSchema, error) {
	childCtx, cancel := context.WithCancel(ctx)
	defer cancel()

	resp, err := r.Connection.request(
		childCtx,
		&requestParams{
			method:   "GET",
			endPoint: "/v3/accounts/" + r.AccountID + "/positions/" + r.Instrument,
			headers: []header{
				{key: "Accept-Datetime-Format", value: "RFC3339"},
			},
		},
	)
	if err != nil {
		return nil, errors.Errorf("Get positions instrument canceled: %v", err)
	}
	defer resp.Body.Close()

	var data interface{}
	switch resp.StatusCode {
	case 200:
		data = new(GetPositionsInstrumentSchema)
	}

	data, err = parseResponse(resp, data, r.Connection.Strict)
	if err != nil {
		return nil, errors.Errorf("Get positions instrument failed: %v", err)
	}
	return data.(*GetPositionsInstrumentSchema), nil
}

// PUT /v3/accounts/{accountID}/positions/{instrument}/close
func (r *ReceiverPositionsInstrumentClose) Put(ctx context.Context, params *PutPositionsInstrumentCloseParams) (*PutPositionsInstrumentCloseSchema, error) {
	childCtx, cancel := context.WithCancel(ctx)
	defer cancel()

	resp, err := r.Connection.request(
		childCtx,
		&requestParams{
			method:   "PUT",
			endPoint: "/v3/accounts/" + r.AccountID + "/positions/" + r.Instrument + "/close",
			headers: []header{
				{key: "Accept-Datetime-Format", value: "RFC3339"},
			},
			body: params.Body,
		},
	)
	if err != nil {
		return nil, errors.Errorf("Put positions instrument close canceled: %v", err)
	}
	defer resp.Body.Close()

	var data interface{}
	switch resp.StatusCode {
	case 200:
		data = new(PutPositionsInstrumentCloseSchema)
	case 400:
		data = new(PutPositionsInstrumentCloseBadRequestError)
	case 404:
		data = new(PutPositionsInstrumentCloseNotFoundError)
	}

	data, err = parseResponse(resp, data, r.Connection.Strict)
	if err != nil {
		return nil, errors.Errorf("Put positions instrument close failed: %v", err)
	}
	return data.(*PutPositionsInstrumentCloseSchema), nil
}
