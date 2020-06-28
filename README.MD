# QuickStart

```go
// Transports expose the service to the network. In this first example we utilize JSON over HTTP.

func main() {
	svc := stringService{}

	uppercaseHandler := httptransport.NewServer(
		makeUppercaseEndpoint(svc),
		decodeUppercaseRequest,
		encodeResponse,
                //Add ServerBefore & ServerAfter func 
                httptransport.ServerBefore(gokit_istio_trace_adapter.DefaultServerBefore),
		httptransport.ServerAfter(gokit_istio_trace_adapter.DefaultServerAfter),
	)

	http.Handle("/uppercase", uppercaseHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
```
