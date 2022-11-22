# SRVR

SRVR is pretty close to the most simple http server you could write in go. The
point is to show the power of Go's interfaces in terms of decoupling. By
explicitly specifying interfaces in the controller and business modules
they become agnostic to the implementation of the logger which allows us to
inject the logger at runtime.

# Getting Started

Run the server with:

```bash
go run .
```

Then hit the hello route with to get a response for the user.

```bash
curl "http://localhost:7777/hello?user_id=1"
```
