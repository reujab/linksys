package linksys

import "time"

// GetLastUpgradeCheck returns the date of the last firmware upgrade check.
func (client Client) GetLastUpgradeCheck() (time.Time, error) {
	var check struct {
		LastCheck       string `json:"lastSuccessfulCheckTime"`
		AvailableUpdate *struct {
			Version     string `json:"firmwareVersion"`
			Date        string `json:"firmwareDate"`
			Description string `json:"description"`
		} `json:"availableUpdate"`
	}
	err := client.MakeRequest("firmwareupdate/GetFirmwareUpdateStatus", nil, &check)
	if err != nil {
		return time.Unix(0, 0), err
	}

	return time.Parse(time.RFC3339, check.LastCheck)
}

// UpgradeFirmware checks for an update and updates if available. Requires authorization.
func (client Client) UpgradeFirmware() error {
	return client.MakeRequest("firmwareupdate/UpdateFirmwareNow")
}
