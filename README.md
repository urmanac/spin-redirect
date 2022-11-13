# spin-redirect

This is a simple HTTP redirect component written in Go.

This is not a Spin application in itself but a component that can be used in applications to redirect a route.

Example usage:

```
# Redirect / to /index.html
[[component]]
id = "redirect-to-index"
source = "path/to/redirect.wasm"
environment = { DESTINATION = "/index.html" }
[component.trigger]
route = "/"
executor = { type = "wagi" }
```
