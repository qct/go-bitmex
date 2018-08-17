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

// UserPreferences user preferences
// swagger:model UserPreferences
type UserPreferences struct {

	// alert on liquidations
	AlertOnLiquidations bool `json:"alertOnLiquidations,omitempty"`

	// animations enabled
	AnimationsEnabled bool `json:"animationsEnabled,omitempty"`

	// announcements last seen
	// Format: date-time
	AnnouncementsLastSeen strfmt.DateTime `json:"announcementsLastSeen,omitempty"`

	// chat channel ID
	ChatChannelID float64 `json:"chatChannelID,omitempty"`

	// color theme
	ColorTheme string `json:"colorTheme,omitempty"`

	// currency
	Currency string `json:"currency,omitempty"`

	// debug
	Debug bool `json:"debug,omitempty"`

	// disable emails
	DisableEmails []string `json:"disableEmails"`

	// hide confirm dialogs
	HideConfirmDialogs []string `json:"hideConfirmDialogs"`

	// hide connection modal
	HideConnectionModal bool `json:"hideConnectionModal,omitempty"`

	// hide from leaderboard
	HideFromLeaderboard *bool `json:"hideFromLeaderboard,omitempty"`

	// hide name from leaderboard
	HideNameFromLeaderboard *bool `json:"hideNameFromLeaderboard,omitempty"`

	// hide notifications
	HideNotifications []string `json:"hideNotifications"`

	// locale
	Locale *string `json:"locale,omitempty"`

	// msgs seen
	MsgsSeen []string `json:"msgsSeen"`

	// order book binning
	OrderBookBinning interface{} `json:"orderBookBinning,omitempty"`

	// order book type
	OrderBookType string `json:"orderBookType,omitempty"`

	// order clear immediate
	OrderClearImmediate *bool `json:"orderClearImmediate,omitempty"`

	// order controls plus minus
	OrderControlsPlusMinus bool `json:"orderControlsPlusMinus,omitempty"`

	// show locale numbers
	ShowLocaleNumbers *bool `json:"showLocaleNumbers,omitempty"`

	// sounds
	Sounds []string `json:"sounds"`

	// strict IP check
	StrictIPCheck *bool `json:"strictIPCheck,omitempty"`

	// strict timeout
	StrictTimeout *bool `json:"strictTimeout,omitempty"`

	// ticker group
	TickerGroup string `json:"tickerGroup,omitempty"`

	// ticker pinned
	TickerPinned bool `json:"tickerPinned,omitempty"`

	// trade layout
	TradeLayout string `json:"tradeLayout,omitempty"`
}

// Validate validates this user preferences
func (m *UserPreferences) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAnnouncementsLastSeen(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UserPreferences) validateAnnouncementsLastSeen(formats strfmt.Registry) error {

	if swag.IsZero(m.AnnouncementsLastSeen) { // not required
		return nil
	}

	if err := validate.FormatOf("announcementsLastSeen", "body", "date-time", m.AnnouncementsLastSeen.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *UserPreferences) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UserPreferences) UnmarshalBinary(b []byte) error {
	var res UserPreferences
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
