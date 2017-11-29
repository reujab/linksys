package linksys

// GetAdminPasswordHint returns the admin password hint.
func (client Client) GetAdminPasswordHint() (string, error) {
	var res struct {
		Hint string `json:"passwordHint"`
	}
	err := client.MakeRequest("core/GetAdminPasswordHint", nil, &res)
	return res.Hint, err
}

// RouterInfo represents information about the router.
type RouterInfo struct {
	Description     string   `json:"description"`
	FirmwareDate    string   `json:"firmwareDate"`
	FirmwareVersion string   `json:"firmwareVersion"`
	HardwareVersion string   `json:"hardwareVersion"`
	Manufacturer    string   `json:"manufacturer"`
	ModelNumber     string   `json:"modelNumber"`
	SerialNumber    string   `json:"serialNumber"`
	Services        []string `json:"services"`
}

// GetRouterInfo returns information about the router.
func (client Client) GetRouterInfo() (RouterInfo, error) {
	var info RouterInfo
	err := client.MakeRequest("core/GetDeviceInfo", nil, &info)
	return info, err
}

// Reboot reboots the router.
func (client Client) Reboot() error {
	return client.MakeRequest("core/Reboot", nil, nil)
}

// SetAdminPassword sets the router's admin password, NOT the network password. Requires authentication. Automatically reauthenticates using new password.
func (client Client) SetAdminPassword(password, hint string) error {
	err := client.MakeRequest("core/SetAdminPassword2", map[string]string{
		"adminPassword": password,
		"passwordHint":  hint,
	}, nil)
	if err != nil {
		return err
	}

	return client.Authorize(password)
}
