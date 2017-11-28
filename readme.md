# linksys
A Go library for interacting with Linksys Smart WiFi-enabled routers using the JNAP protocol described below.

## HTTP API overview
All API requests are POST requests made to `http://192.168.1.1/JNAP/`. The API endpoint can be changed if `192.168.1.1` does not point to the router. The special domain [LinksysSmartWIFI.com](http://linksyssmartwifi.com) points to the router ONLY if the DNS server is still the default. If using any other DNS server, such as Google's `8.8.8.8`, it will point to `54.183.80.213`.

All requests contain a JSON object as the POST data. Even if the action doesn't require any parameters, the POST data must be `{}`. All actions below take no parameters unless otherwise documented. The `Content-Type` header does not matter and is assumed to be `application/json`.

All responses are JSON objects with two keys: `output` and `response`.
TODO: document keys

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

#### `core/GetDeviceInfo`
This action returns information about the router including the model, firmware, and a list of supported actions.

#### `devicelist/GetDevices`
This action returns every device that has connected to the router (whether it is currently connected or not) with information such as it's local IP address (if currently connected), hostname, MAC addresses, device model (if detected), and operating system (if detected). Surprisingly, this action does not require authorization.

#### `networkconnections/GetNetworkConnections`
This action returns information associated with every connected device including its MAC address, Mbps, band (2.4GHz, 5GHz, ...), and signal decibels. Surprisingly, this action does not require authorization.

#### `networkconnections/GetNetworkConnections2`
Same as `networkconnections/GetNetworkConnections`, but with the radio ID.

#### `router/GetWANStatus`
This action returns information about the WAN status of the router.

#### `router/GetWANStatus2`
This action returns a bit more information than `core/GetWANStatus`.

#### `router/GetWANStatus3`
This action returns a bit more information than `core/GetWANStatus2`.

### Undocumented actions
* `ui/GetRemoteSetting`
