package request

import "time"

type Request struct {
	Handler Handler `json:"handler"`
	Intent  Intent  `json:"intent"`
	Scene   Scene   `json:"scene"`
	Session Session `json:"session"`
	User    User    `json:"user"`
	Home    Home    `json:"home"`
	Device  Device  `json:"device"`
}
type Handler struct {
	Name string `json:"name"`
}
type Resolved struct {
	Month int `json:"month"`
	Year  int `json:"year"`
	Day   int `json:"day"`
}
type IntentParam struct {
	Original string   `json:"original"`
	Resolved Resolved `json:"resolved"`
}
type IntentParams struct {
	DateFrom   IntentParam `json:"DateFrom"`
	DateTo     IntentParam `json:"DateTo"`
	DateAmount IntentParam `json:"DateAmount"`
}
type Intent struct {
	Name   string       `json:"name"`
	Params IntentParams `json:"params"`
	Query  string       `json:"query"`
}
type Slots struct {
}
type Scene struct {
	Name              string `json:"name"`
	SlotFillingStatus string `json:"slotFillingStatus"`
	Slots             Slots  `json:"slots"`
}
type Params struct {
}
type Session struct {
	ID            string        `json:"id"`
	Params        Params        `json:"params"`
	TypeOverrides []interface{} `json:"typeOverrides"`
	LanguageCode  string        `json:"languageCode"`
}
type User struct {
	Locale               string        `json:"locale"`
	Params               Params        `json:"params"`
	AccountLinkingStatus string        `json:"accountLinkingStatus"`
	VerificationStatus   string        `json:"verificationStatus"`
	PackageEntitlements  []interface{} `json:"packageEntitlements"`
	LastSeenTime         time.Time     `json:"lastSeenTime"`
}
type Home struct {
	Params Params `json:"params"`
}
type Device struct {
	Capabilities []string `json:"capabilities"`
}
