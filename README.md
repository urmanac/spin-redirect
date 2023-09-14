# spin-redirect

This is a simple HTTP redirect component written in Go.

This is not a Spin application in itself but a component that can be used in applications to redirect a route.

## Configuration

The `spin-redirect` component can be configured to address your needs in different scenarios. `spin-redirect` tries to load configuration data from multiple places in the specified order:

1. Spin component configuration
2. Environment variables

The following table outlines available configuration values:

| Key           | Description                                           | Default Value    |
|---------------|-------------------------------------------------------|------------------|
| `destination` | Where should the component redirect to                | `/`              |
| `statuscode`  | What HTTP status code should be used when redirecting | `302`            |

The `spin-redirect` component tries to look up the config value in the Spin component configuration using the keys shown in the table above (lower case). If desired key is not present, it transforms the key to upper case (e.g., `DESTINATION`) and checks environment variables.

### Valid redirection status codes

The `spin-redirect` component supports the following HTTP status codes to perform a redirect:

- `301` Moved Permanently
- `302` Found (Moved Temporarily)
- `303` See Other: Only supported for `PUT` and `POST` requests
- `307` Temporary Redirect
- `308` Permanent Redirect

## Example usage

The following snippet shows how to add and configure `spin-redirect` in your `spin.toml` using environment variables

```toml
spin_manifest_version = "1"
description = ""
name = "test"
trigger = { type = "http", base = "/" }
version = "0.1.0"

# Redirect / to /index.html using HTTP status code 301
[[component]]
id = "redirect-sample"
source = "path/to/redirect.wasm"
environment = { DESTINATION = "/index.html", STATUSCODE = "301" }

[component.trigger]
route = "/"
```

Alternatively, you can use component configuration to configure `spin-redirect` as shown below:

```toml
spin_manifest_version = "1"
description = ""
name = "test"
trigger = { type = "http", base = "/" }
version = "0.1.0"

# Redirect / to /index.html using HTTP status code 301
[[component]]
id = "redirect-sample"
source = "path/to/redirect.wasm"

[component.config]
destination="/index.html"
statuscode="301"

[component.trigger]
route = "/"

```
