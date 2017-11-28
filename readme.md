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
Most actions require authentication provided by the `X-JNAP-Authorization` HTTP header. The header is a [basic access authentication](https://en.wikipedia.org/wiki/Basic_access_authentication) with the username `admin`. An example for the value of `X-JNAP-Authorization` in pseudocode would be:
```
"Basic "+base64("admin:"+password)
```

For example, if the password was `password`, the header would be `Basic YWRtaW46cGFzc3dvcmQ=`
