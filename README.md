# wsdyn_router
A small lightweigth dynamic proxy/ webservice registry relying on etcd


Idea is to have a central registry where available webservices can be registered dynamically. The Host/URL for the client should remain static, but requests are routed transparently to the back end. The configuration can be changed on the fly as the router stores it's configuration in etcd.

This is a POC only at the moment - ugly code, just to get it up and running.

Additional topics which might be considered in future development:
- add stability
- check memory usage
- error handling, when backend is not available
- make more testing for RESTful services
- start testing for SOAP/http services

# License

I put the software temporarily under the Go-compatible BSD license, if this prevents someone from using the software, do let mee know and I'll consider changing it.

At any rate, user feedback is very important for me, so I'll be delighted to know if you're using this package.
