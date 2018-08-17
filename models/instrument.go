// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Instrument Tradeable Contracts, Indices, and History
// swagger:model Instrument
type Instrument struct {

	// ask price
	AskPrice float64 `json:"askPrice,omitempty"`

	// bankrupt limit down price
	BankruptLimitDownPrice float64 `json:"bankruptLimitDownPrice,omitempty"`

	// bankrupt limit up price
	BankruptLimitUpPrice float64 `json:"bankruptLimitUpPrice,omitempty"`

	// bid price
	BidPrice float64 `json:"bidPrice,omitempty"`

	// buy leg
	BuyLeg string `json:"buyLeg,omitempty"`

	// calc interval
	// Format: date-time
	CalcInterval strfmt.DateTime `json:"calcInterval,omitempty"`

	// capped
	Capped bool `json:"capped,omitempty"`

	// closing timestamp
	// Format: date-time
	ClosingTimestamp strfmt.DateTime `json:"closingTimestamp,omitempty"`

	// deleverage
	Deleverage bool `json:"deleverage,omitempty"`

	// expiry
	// Format: date-time
	Expiry strfmt.DateTime `json:"expiry,omitempty"`

	// fair basis
	FairBasis float64 `json:"fairBasis,omitempty"`

	// fair basis rate
	FairBasisRate float64 `json:"fairBasisRate,omitempty"`

	// fair method
	FairMethod string `json:"fairMethod,omitempty"`

	// fair price
	FairPrice float64 `json:"fairPrice,omitempty"`

	// foreign notional24h
	ForeignNotional24h float64 `json:"foreignNotional24h,omitempty"`

	// front
	// Format: date-time
	Front strfmt.DateTime `json:"front,omitempty"`

	// funding base symbol
	FundingBaseSymbol string `json:"fundingBaseSymbol,omitempty"`

	// funding interval
	// Format: date-time
	FundingInterval strfmt.DateTime `json:"fundingInterval,omitempty"`

	// funding premium symbol
	FundingPremiumSymbol string `json:"fundingPremiumSymbol,omitempty"`

	// funding quote symbol
	FundingQuoteSymbol string `json:"fundingQuoteSymbol,omitempty"`

	// funding rate
	FundingRate float64 `json:"fundingRate,omitempty"`

	// funding timestamp
	// Format: date-time
	FundingTimestamp strfmt.DateTime `json:"fundingTimestamp,omitempty"`

	// has liquidity
	HasLiquidity bool `json:"hasLiquidity,omitempty"`

	// high price
	HighPrice float64 `json:"highPrice,omitempty"`

	// home notional24h
	HomeNotional24h float64 `json:"homeNotional24h,omitempty"`

	// impact ask price
	ImpactAskPrice float64 `json:"impactAskPrice,omitempty"`

	// impact bid price
	ImpactBidPrice float64 `json:"impactBidPrice,omitempty"`

	// impact mid price
	ImpactMidPrice float64 `json:"impactMidPrice,omitempty"`

	// indicative funding rate
	IndicativeFundingRate float64 `json:"indicativeFundingRate,omitempty"`

	// indicative settle price
	IndicativeSettlePrice float64 `json:"indicativeSettlePrice,omitempty"`

	// indicative tax rate
	IndicativeTaxRate float64 `json:"indicativeTaxRate,omitempty"`

	// init margin
	InitMargin float64 `json:"initMargin,omitempty"`

	// insurance fee
	InsuranceFee float64 `json:"insuranceFee,omitempty"`

	// inverse leg
	InverseLeg string `json:"inverseLeg,omitempty"`

	// is inverse
	IsInverse bool `json:"isInverse,omitempty"`

	// is quanto
	IsQuanto bool `json:"isQuanto,omitempty"`

	// last change pcnt
	LastChangePcnt float64 `json:"lastChangePcnt,omitempty"`

	// last price
	LastPrice float64 `json:"lastPrice,omitempty"`

	// last price protected
	LastPriceProtected float64 `json:"lastPriceProtected,omitempty"`

	// last tick direction
	LastTickDirection string `json:"lastTickDirection,omitempty"`

	// limit
	Limit float64 `json:"limit,omitempty"`

	// limit down price
	LimitDownPrice float64 `json:"limitDownPrice,omitempty"`

	// limit up price
	LimitUpPrice float64 `json:"limitUpPrice,omitempty"`

	// listing
	// Format: date-time
	Listing strfmt.DateTime `json:"listing,omitempty"`

	// lot size
	LotSize int64 `json:"lotSize,omitempty"`

	// low price
	LowPrice float64 `json:"lowPrice,omitempty"`

	// maint margin
	MaintMargin float64 `json:"maintMargin,omitempty"`

	// maker fee
	MakerFee float64 `json:"makerFee,omitempty"`

	// mark method
	MarkMethod string `json:"markMethod,omitempty"`

	// mark price
	MarkPrice float64 `json:"markPrice,omitempty"`

	// max order qty
	MaxOrderQty int64 `json:"maxOrderQty,omitempty"`

	// max price
	MaxPrice float64 `json:"maxPrice,omitempty"`

	// mid price
	MidPrice float64 `json:"midPrice,omitempty"`

	// multiplier
	Multiplier int64 `json:"multiplier,omitempty"`

	// open interest
	OpenInterest int64 `json:"openInterest,omitempty"`

	// open value
	OpenValue int64 `json:"openValue,omitempty"`

	// opening timestamp
	// Format: date-time
	OpeningTimestamp strfmt.DateTime `json:"openingTimestamp,omitempty"`

	// option multiplier
	OptionMultiplier float64 `json:"optionMultiplier,omitempty"`

	// option strike pcnt
	OptionStrikePcnt float64 `json:"optionStrikePcnt,omitempty"`

	// option strike price
	OptionStrikePrice float64 `json:"optionStrikePrice,omitempty"`

	// option strike round
	OptionStrikeRound float64 `json:"optionStrikeRound,omitempty"`

	// option underlying price
	OptionUnderlyingPrice float64 `json:"optionUnderlyingPrice,omitempty"`

	// position currency
	PositionCurrency string `json:"positionCurrency,omitempty"`

	// prev close price
	PrevClosePrice float64 `json:"prevClosePrice,omitempty"`

	// prev price24h
	PrevPrice24h float64 `json:"prevPrice24h,omitempty"`

	// prev total turnover
	PrevTotalTurnover int64 `json:"prevTotalTurnover,omitempty"`

	// prev total volume
	PrevTotalVolume int64 `json:"prevTotalVolume,omitempty"`

	// publish interval
	// Format: date-time
	PublishInterval strfmt.DateTime `json:"publishInterval,omitempty"`

	// publish time
	// Format: date-time
	PublishTime strfmt.DateTime `json:"publishTime,omitempty"`

	// quote currency
	QuoteCurrency string `json:"quoteCurrency,omitempty"`

	// quote to settle multiplier
	QuoteToSettleMultiplier int64 `json:"quoteToSettleMultiplier,omitempty"`

	// rebalance interval
	// Format: date-time
	RebalanceInterval strfmt.DateTime `json:"rebalanceInterval,omitempty"`

	// rebalance timestamp
	// Format: date-time
	RebalanceTimestamp strfmt.DateTime `json:"rebalanceTimestamp,omitempty"`

	// reference
	Reference string `json:"reference,omitempty"`

	// reference symbol
	ReferenceSymbol string `json:"referenceSymbol,omitempty"`

	// relist interval
	// Format: date-time
	RelistInterval strfmt.DateTime `json:"relistInterval,omitempty"`

	// risk limit
	RiskLimit int64 `json:"riskLimit,omitempty"`

	// risk step
	RiskStep int64 `json:"riskStep,omitempty"`

	// root symbol
	RootSymbol string `json:"rootSymbol,omitempty"`

	// sell leg
	SellLeg string `json:"sellLeg,omitempty"`

	// session interval
	// Format: date-time
	SessionInterval strfmt.DateTime `json:"sessionInterval,omitempty"`

	// settl currency
	SettlCurrency string `json:"settlCurrency,omitempty"`

	// settle
	// Format: date-time
	Settle strfmt.DateTime `json:"settle,omitempty"`

	// settled price
	SettledPrice float64 `json:"settledPrice,omitempty"`

	// settlement fee
	SettlementFee float64 `json:"settlementFee,omitempty"`

	// state
	State string `json:"state,omitempty"`

	// symbol
	// Required: true
	Symbol *string `json:"symbol"`

	// taker fee
	TakerFee float64 `json:"takerFee,omitempty"`

	// taxed
	Taxed bool `json:"taxed,omitempty"`

	// tick size
	TickSize float64 `json:"tickSize,omitempty"`

	// timestamp
	// Format: date-time
	Timestamp strfmt.DateTime `json:"timestamp,omitempty"`

	// total turnover
	TotalTurnover int64 `json:"totalTurnover,omitempty"`

	// total volume
	TotalVolume int64 `json:"totalVolume,omitempty"`

	// turnover
	Turnover int64 `json:"turnover,omitempty"`

	// turnover24h
	Turnover24h int64 `json:"turnover24h,omitempty"`

	// typ
	Typ string `json:"typ,omitempty"`

	// underlying
	Underlying string `json:"underlying,omitempty"`

	// underlying symbol
	UnderlyingSymbol string `json:"underlyingSymbol,omitempty"`

	// underlying to position multiplier
	UnderlyingToPositionMultiplier int64 `json:"underlyingToPositionMultiplier,omitempty"`

	// underlying to settle multiplier
	UnderlyingToSettleMultiplier int64 `json:"underlyingToSettleMultiplier,omitempty"`

	// volume
	Volume int64 `json:"volume,omitempty"`

	// volume24h
	Volume24h int64 `json:"volume24h,omitempty"`

	// vwap
	Vwap float64 `json:"vwap,omitempty"`
}

// Validate validates this instrument
func (m *Instrument) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCalcInterval(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClosingTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExpiry(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFront(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFundingInterval(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFundingTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateListing(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOpeningTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePublishInterval(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePublishTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRebalanceInterval(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRebalanceTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRelistInterval(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSessionInterval(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSettle(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSymbol(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Instrument) validateCalcInterval(formats strfmt.Registry) error {

	if swag.IsZero(m.CalcInterval) { // not required
		return nil
	}

	if err := validate.FormatOf("calcInterval", "body", "date-time", m.CalcInterval.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Instrument) validateClosingTimestamp(formats strfmt.Registry) error {

	if swag.IsZero(m.ClosingTimestamp) { // not required
		return nil
	}

	if err := validate.FormatOf("closingTimestamp", "body", "date-time", m.ClosingTimestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Instrument) validateExpiry(formats strfmt.Registry) error {

	if swag.IsZero(m.Expiry) { // not required
		return nil
	}

	if err := validate.FormatOf("expiry", "body", "date-time", m.Expiry.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Instrument) validateFront(formats strfmt.Registry) error {

	if swag.IsZero(m.Front) { // not required
		return nil
	}

	if err := validate.FormatOf("front", "body", "date-time", m.Front.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Instrument) validateFundingInterval(formats strfmt.Registry) error {

	if swag.IsZero(m.FundingInterval) { // not required
		return nil
	}

	if err := validate.FormatOf("fundingInterval", "body", "date-time", m.FundingInterval.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Instrument) validateFundingTimestamp(formats strfmt.Registry) error {

	if swag.IsZero(m.FundingTimestamp) { // not required
		return nil
	}

	if err := validate.FormatOf("fundingTimestamp", "body", "date-time", m.FundingTimestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Instrument) validateListing(formats strfmt.Registry) error {

	if swag.IsZero(m.Listing) { // not required
		return nil
	}

	if err := validate.FormatOf("listing", "body", "date-time", m.Listing.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Instrument) validateOpeningTimestamp(formats strfmt.Registry) error {

	if swag.IsZero(m.OpeningTimestamp) { // not required
		return nil
	}

	if err := validate.FormatOf("openingTimestamp", "body", "date-time", m.OpeningTimestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Instrument) validatePublishInterval(formats strfmt.Registry) error {

	if swag.IsZero(m.PublishInterval) { // not required
		return nil
	}

	if err := validate.FormatOf("publishInterval", "body", "date-time", m.PublishInterval.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Instrument) validatePublishTime(formats strfmt.Registry) error {

	if swag.IsZero(m.PublishTime) { // not required
		return nil
	}

	if err := validate.FormatOf("publishTime", "body", "date-time", m.PublishTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Instrument) validateRebalanceInterval(formats strfmt.Registry) error {

	if swag.IsZero(m.RebalanceInterval) { // not required
		return nil
	}

	if err := validate.FormatOf("rebalanceInterval", "body", "date-time", m.RebalanceInterval.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Instrument) validateRebalanceTimestamp(formats strfmt.Registry) error {

	if swag.IsZero(m.RebalanceTimestamp) { // not required
		return nil
	}

	if err := validate.FormatOf("rebalanceTimestamp", "body", "date-time", m.RebalanceTimestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Instrument) validateRelistInterval(formats strfmt.Registry) error {

	if swag.IsZero(m.RelistInterval) { // not required
		return nil
	}

	if err := validate.FormatOf("relistInterval", "body", "date-time", m.RelistInterval.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Instrument) validateSessionInterval(formats strfmt.Registry) error {

	if swag.IsZero(m.SessionInterval) { // not required
		return nil
	}

	if err := validate.FormatOf("sessionInterval", "body", "date-time", m.SessionInterval.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Instrument) validateSettle(formats strfmt.Registry) error {

	if swag.IsZero(m.Settle) { // not required
		return nil
	}

	if err := validate.FormatOf("settle", "body", "date-time", m.Settle.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Instrument) validateSymbol(formats strfmt.Registry) error {

	if err := validate.Required("symbol", "body", m.Symbol); err != nil {
		return err
	}

	return nil
}

func (m *Instrument) validateTimestamp(formats strfmt.Registry) error {

	if swag.IsZero(m.Timestamp) { // not required
		return nil
	}

	if err := validate.FormatOf("timestamp", "body", "date-time", m.Timestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Instrument) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Instrument) UnmarshalBinary(b []byte) error {
	var res Instrument
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
