package nbxplorer

import (
	"errors"
)

// Destination struct
type Destination struct {
	Destination   string `json:"destination"`
	Amount        int    `json:"amount"`
	SubstractFees bool   `json:"substractFees"`
	SweepAll      bool   `json:"sweepAll"`
}

// FeePreference struct
type FeePreference struct {
	ExplicitFeeRate int `json:"explicitFeeRate"`
	ExplicitFee     int `json:"explicitFee"`
	BlockTarget     int `json:"blockTarget"`
	FallbackFeeRate int `json:"fallbackFeeRate"`
}

// KeyPath struct
type KeyPath struct {
	AccountKey     string `json:"accountKey"`
	AccountKeyPath string `json:"accountKeyPath"`
}

// PSBT struct
type PSBT struct {
	Seed                            int           `json:"seed"`
	Rbf                             bool          `json:"rbf"`
	Version                         int           `json:"version"`
	TimeLock                        int           `json:"timeLock"`
	ExplicitChangeAddress           string        `json:"explicitChangeAddress"`
	Destinations                    []Destination `json:"destinations"`
	FeePreference                   FeePreference `json:"feePreference"`
	DiscourageFeeSniping            bool          `json:"discourageFeeSniping"`
	ReserveChangeAddress            bool          `json:"reserveChangeAddress"`
	MinConfirmations                int           `json:"minConfirmations"`
	ExcludeOutpoints                []string      `json:"excludeOutpoints"`
	IncludeOnlyOutpoints            []string      `json:"includeOnlyOutpoints"`
	MinValue                        int           `json:"minValue"`
	RebaseKeyPaths                  []KeyPath     `json:"rebaseKeyPaths"`
	DisableFingerprintRandomization bool          `json:"disableFingerprintRandomization"`
	AlwaysIncludeNonWitnessUTXO     bool          `json:"alwaysIncludeNonWitnessUTXO"`
}

// CreatePSBTResponse struct
type CreatePSBTResponse struct {
	Psbt          string `json:"psbt"`
	ChangeAddress string `json:"changeAddress"`
	Suggestions   struct {
		ShouldEnforceLowR bool `json:"shouldEnforceLowR"`
	} `json:"suggestions"`
}

// CreatePSBT
func (c *Client) CreatePSBT(derivationScheme string, psbt PSBT) (CreatePSBTResponse, error) {
	var r ErrorResponse
	var response CreatePSBTResponse
	resp, err := c.R().
		SetBody(psbt).
		SetResult(&response).
		SetError(&r).
		Post("/derivations/" + derivationScheme + "/psbt/create")

	if err != nil {
		return response, err
	}

	if resp.StatusCode() != 200 {
		return response, errors.New(r.Message)
	}

	return response, nil
}
