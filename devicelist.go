package linksys

// Devices represents a list of devices that have connected to the router.
type Devices struct {
	Devices  []Device `json:"devices"`
	Revision int      `json:"revision"`
}

// Device represents a device that has connected to the router.
type Device struct {
	Connections []struct {
		IP  string `json:"ipAddress"`
		MAC string `json:"macAddress"`
	} `json:"connections"`
	GUID                 string   `json:"deviceID"`
	Hostname             string   `json:"friendlyName"`
	Authority            bool     `json:"isAuthority"`
	MACAddresses         []string `json:"knownMACAddresses"`
	LastChangeRevision   int      `json:"lastChangeRevision"`
	MaxAllowedProperties int      `json:"maxAllowedProperties"`
	Model                struct {
		Type         string `json:"deviceType"`
		Manufacturer string `json:"manufacturer"`
		Name         string `json:"modelNumber"`
	} `json:"model"`
	Properties []struct {
		Name  string `json:"name"`
		Value string `json:"value"`
	} `json:"properties"`
	Unit struct {
		OS           string `json:"operatingSystem"`
		SerialNumber string `json:"serialNumber"`
	}
}

// GetDevices returns every device that has connected to the router (whether it is currently connected or not). `revision` can be set to 0 to get all devices.
func (client Client) GetDevices(revision int) (Devices, error) {
	var devices Devices
	err := client.MakeRequest("devicelist/GetDevices", map[string]int{
		"sinceRevision": revision,
	}, &devices)
	return devices, err
}

// GetCurrentDeviceGUID returns the current machine's GUID provided by the router.
func (client Client) GetCurrentDeviceGUID() (string, error) {
	var guid struct {
		GUID string `json:"deviceID"`
	}
	err := client.MakeRequest("devicelist/GetLocalDevice", nil, &guid)
	return guid.GUID, err
}
