# linksys
A Go library for interacting with Linksys Smart WiFi-enabled routers using the JNAP protocol described below.

## HTTP API overview
All API requests are POST requests made to `http://192.168.1.1/JNAP/`. The API endpoint can be changed if `192.168.1.1` does not point to the router. The special domain [LinksysSmartWIFI.com](http://linksyssmartwifi.com) points to the router ONLY if the DNS server is still the default. If using any other DNS server, such as Google's `8.8.8.8`, it will point to `54.183.80.213`.

All requests contain a JSON object as the POST data. Even if the action doesn't require any parameters, the POST data must be `{}`. All actions below take no parameters unless otherwise documented. The `Content-Type` header does not matter and is assumed to be `application/json`.

All responses are JSON objects with two keys: `output` and `response`. The `response` key indicates the status of the action. A value of `OK` indicates the operation was successful, any other value indicates an error. When an error occurs, another field, `error`, may be returned.

#### `X-JNAP-Action`
Every request has a HTTP header, `X-JNAP-Action`, that describes the action, request, and response parameters. Every action starts with `http://linksys.com/jnap/` (alternatively, `http://cisco.com/jnap/` could be used, but it seems to be deprecated). The actions described below have the prefix omitted, but are required in the header. For example, if the heading below is `core/GetDeviceInfo`, the HTTP header `X-JNAP-Action` in the request will be `http://linksys.com/jnap/core/GetDeviceInfo`.

#### `X-JNAP-Authorization`
Most actions require authorization provided by the `X-JNAP-Authorization` HTTP header. The header is a [basic access authentication](https://en.wikipedia.org/wiki/Basic_access_authentication) with the username `admin`. An example for the value of `X-JNAP-Authorization` in pseudocode would be:
```
"Basic "+base64("admin:"+password)
```

For example, if the password was `password`, the header would be `Basic YWRtaW46cGFzc3dvcmQ=`

### Unauthorized requests
#### `core/GetAdminPasswordHint`
This action returns the password hint.

#### `core/GetAdminPasswordRestrictions`
This action returns the restrictions for the admin password.

#### `core/GetDeviceInfo`
This action returns information about the router including the model, firmware, and a list of supported actions.

#### `core/IsAdminPasswordDefault`
This action returns whether the admin password is the default password.

#### `devicelist/GetDevices`
This action returns every device that has connected to the router (whether it is currently connected or not) with information such as it's local IP address (if currently connected), hostname, MAC addresses, device model (if detected), and operating system (if detected). Surprisingly, this action does not require authorization.
* `sinceRevision` an integer that, when provided, makes the response only contain devices that have connected since the specified revision.

#### `devicelist/GetLocalDevice`
This action returns the router-assigned GUID of the device performing the request.

#### `firewall/GetALGSettings`
This action returns the firewall ALG settings.

#### `firmwareupdate/GetFirmwareUpdateStatus`
This action returns the timestamp of the last time the router checked for an update.
<!-- TODO: check response when update is available -->

#### `locale/GetTimeSettings`
This action returns the time zone of the router.

#### `networkconnections/GetNetworkConnections`
This action returns information associated with every connected device including its MAC address, Mbps, band (2.4GHz, 5GHz, ...), and signal decibels. Surprisingly, this action does not require authorization.

#### `networkconnections/GetNetworkConnections2`
Same as `networkconnections/GetNetworkConnections`, but with the radio ID.

#### `parentalcontrol/GetParentalControlSettings`
This action returns the parentral controls settings, such as if it's enabled and rules.

#### `qos/GetQoSSettings`
This actions returns QoS settings.

#### `qos/GetWLANQoSSettings`
This actions returns WLAN QoS settings.

#### `router/GetDHCPClientLeases`
This action returns DHCP leases for every connected device.

#### `router/GetEthernetPortConnections`
This action returns ethernet port connections.

#### `router/GetExpressForwardingSettings`
This action returns whether express forwarding is enabled.

#### `router/GetIPv6Settings`
This action returns IPv6 settings.

#### `router/GetIPv6Settings2`
Same as `router/GetIPv6Settings`, but with more information.

#### `router/GetLANSettings`
This actions returns LAN settings.

#### `router/GetMACAddressCloneSettings`
This action returns the MAC address clone settings.

#### `router/GetRoutingSettings`
This action returns routing settings.

#### `router/GetWANStatus`
This action returns information about the WAN status of the router.

#### `router/GetWANStatus2`
This action returns a bit more information than `core/GetWANStatus`.

#### `router/GetWANStatus3`
This action returns a bit more information than `core/GetWANStatus2`.

#### `routerleds/GetRouterLEDSettings`
This action returns a list of activated LEDs on the router.

#### `routerlog/GetDHCPLogEntries`
This action returns DHCP log entries.
* `firstEntryIndex` an integer
* `entryCount` an integer

#### `routerlog/GetIncomingLogEntries`
This action returns incoming log entries.
* `firstEntryIndex` an integer
* `entryCount` an integer

#### `routerlog/GetLogSettings`
This action returns log settings.

#### `routerlog/GetOutgoingLogEntries`
This action returns outgoing log entries.
* `firstEntryIndex` an integer
* `entryCount` an integer

#### `routerlog/GetSecurityLogEntries`
This action returns security log entries.
* `firstEntryIndex` an integer
* `entryCount` an integer

#### `routermanagement/GetManagementSettings`
This action returns management settings of the router.

#### `routermanagement/GetManagementSettings2`
Same as `routermanagement/GetManagementSettings`, but with more information.

#### `routermanagement/GetRemoteManagementStatus`
This action returns the remote management status of the router.

#### `routerupnp/GetUPnPSettings`
This action returns UPnP settings.

#### `storage/GetPartitions`
This action returns a list of external storage devices.

#### `wirelessscheduler/GetWirelessSchedulerSettings`
This actions returns the hours that wireless access is permitted by parental controls.

### Authorized requests
#### `core/CheckAdminPassword`
This action valides the password provided with the `X-JNAP-Authorization`.

#### `core/Reboot`
This action instructs the router to reboot.

#### `core/SetAdminPassword2`
This action sets the _router_ admin password, NOT the network password.
* `adminPassword` a string set to the password
* `passwordHint` a hint viewable by anyone connected to the network

#### `devicelist/SetDeviceProperties`
This action sets properties for a specified device.
* `deviceID` the device GUID assigned by the router
* `propertiesToModify` an array of objects of properties to modify
	* `name` the property to modify. known values are: `urn:cisco-com:ui:qos` (QoS)
	* `value` the value to set the property
		* QoS
			* `userPrioritized` prioritizes the device over other devices

#### `firmwareupdate/SetFirmwareUpdateSettings`
This action sets the interval at which to check for updates and whether to automatically check for updates.
* `autoUpdateWindow` an object containing information on when to check for updates
	* `durationMinutes` interval at which to check for updates in minutes
	* `startMinute`
* `updatePolicy` a string set to `Manual` or `AutomaticallyCheckAndInstall`

#### `firmwareupdate/UpdateFirmwareNow`
This action instructs the router to perform a firmware upgrade.
<!-- TODO: check response when update is available -->

#### `guestnetwork/GetGuestRadioSettings`
This action returns information about the guest network, such as if it's enabled, its SSID, and password.

#### `guestnetwork/SetGuestRadioSettings`
This action sets the settings for the guest network.
* `isGuestNetworkEnabled` a boolean that is set to true when the guest network will be enabled.
* `maxSimultaneousGuests` an integer that is set to the maximum number of devices that are allowed to be connected to the guest network.
* `radios` an array containing information about the network
	* `broadcastGuestSSID` a string containing the SSID for the guest network
	* `guestPassword` a string containing the guest password
	* `guestSSID` a string containing the guest SSID
	* `isEnabled` a boolean that is set to true when the guest network is enabled. only takes effect when `isGuestNetworkEnabled` is set to true
	* `radioID` a string containing the radio ID. can be obtained with `guestnetwork/GetGuestRadioSettings`

#### `locale/SetTimeSettings`
This action sets the time zone for the router.
* `autoAdjustForDST` a boolean set to true when configured to automatically adject for daylight savings time
* `timeZoneID` the id for the time zone

#### `parentalcontrol/SetParentalControlSettings`
This action sets the parental controls settings.
* `isParentalControlEnabled` a boolean that is set to true when parental controls are enabled
* `rules` an array of objects of rules
	* `blockedURLs` an array of blocked URLs
	* `description` a string with no meaning
	* `isEnabled` a boolean
	* `macAddresses` an array of affected MAC addresses
	* `wanSchedule` an object of days and restricted hours
		* `sunday` a 48-length string of restricted hours. every character in the string is either a 0 or a 1. 0 means restricted and 1 means allowed. the numbers are in 30-minute intervals. for example, a value of "0101..." would restrict access from 12am-12:30am and 1am-1:30am
		* `monday` same as sunday
		* `tuesday` same as sunday
		* `wednesday` same as sunday
		* `thursday` same as sunday
		* `friday` same as sunday
		* `saturday` same as sunday

#### `routerleds/SetRouterLEDSettings`
This actions enables or disables LEDs on the router.
* `isSwitchportLEDEnabled` a boolean

#### `routerlog/SetLogSettings`
This action enables or disables logging.
* `isLoggingEnabled` a boolean

#### `vlantagging/GetProfiles`
This action returns VLAN tagging profiles.

#### `vlantagging/GetVLANTaggingSettings`
This action returns VLAN tagging settings.

#### `wirelessap/GetRadioInfo`
This action returns information about the main network, such as its SSID and password.

#### `wirelessap/GetRadioInfo2`
Same as `wirelessap/GetRadioInfo`, but with more information.

#### `wirelessap/GetRadioInfo3`
Same as `wirelessap/GetRadioInfo2`, but with more information.

#### `wirelessap/SetRadioSettings3`
This action sets the SSID and password for the main main networks.
* `bandSteeringMode` a string set to the band steering mode
* `isBandSteeringEnabled` a boolean set to true when band steering is enabled
* `radios` an array of objects describing settings for different frequencies
	* `radioID` the ID of the frequency
	* `settings` an object containing the settings of the frequency
		* `broadcastSSID` a boolean set to false when the network is hidden
		* `channel`
		* `channelWidth`
		* `isEnabled`
		* `mode`
		* `security` a string set to the security protocol
		* `ssid`
		* `wpaPersonalSettings` an object containing the password
			* `passphrase` a string set to the password of the network

### Undocumented actions
* `ui/GetRemoteSetting`
