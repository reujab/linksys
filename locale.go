package linksys

import "time"

// TimeSettings represents the data returned by Client#GetTime()
type TimeSettings struct {
	AutoAdjustForDST bool
	CurrentTime      time.Time
	TimeZones        []struct {
		Description string
		ObserveDST  bool
		TimeZoneID  string
		// the UTC offset in minutes
		UTCOffset int `json:"utcOffsetMinutes"`
	} `json:"supportedTimeZones"`
	TimeZone string `json:"timeZoneID"`
}

// GetTime returns the current time and time zone of the router.
func (client Client) GetTime() (TimeSettings, error) {
	var timeSettings TimeSettings
	err := client.MakeRequest("locale/GetTimeSettings", nil, &timeSettings)
	return timeSettings, err
}

// SetTime sets the current time zone of the router.
func (client Client) SetTime(timeZone string, autoAdjustForDST bool) error {
	return client.MakeRequest("locale/SetTimeSettings", struct {
		AutoAdjustForDST bool   `json:"autoAdjustForDST"`
		TimeZone         string `json:"timeZoneID"`
	}{autoAdjustForDST, timeZone}, nil)
}
