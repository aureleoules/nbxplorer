package nbxplorer

import (
	"errors"
	"strconv"
)

// EventType type
type EventType string

const (
	NewTransaction EventType = "newtransaction"
	NewBlock       EventType = "newblock"
)

// Event struct
type Event struct {
	EventID int         `json:"eventId"`
	Type    EventType   `json:"type"`
	Data    interface{} `json:"data"`
}

func (c *Client) GetEventStream(lastEventID int, longPolling bool, limit *int) ([]Event, error) {
	var r ErrorResponse
	var events []Event
	req := c.httpClient.R().
		SetResult(&events).
		SetQueryParam("lastEventId", strconv.Itoa(lastEventID)).
		SetQueryParam("longPolling", strconv.FormatBool(longPolling)).
		SetError(&r)

	if limit != nil {
		req.SetQueryParam("limit", strconv.Itoa(*limit))
	}

	resp, err := req.Get("/events")
	if err != nil {
		return nil, err
	}

	if resp.StatusCode() != 200 {
		return nil, errors.New(r.Message)
	}

	return events, nil
}

// GetRecentEventStream
func (c *Client) GetRecentEventStream(limit int) ([]Event, error) {
	var r ErrorResponse
	var events []Event
	resp, err := c.httpClient.R().
		SetResult(&events).
		SetQueryParam("limit", strconv.Itoa(limit)).
		SetError(&r).Get("/events/latest")

	if err != nil {
		return nil, err
	}

	if resp.StatusCode() != 200 {
		return nil, errors.New(r.Message)
	}

	return events, nil
}
