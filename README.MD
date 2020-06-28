# QuickStart

```go
// Transports expose the service to the network. In this first example we utilize JSON over HTTP.
import (
        ...
	traceAdapter "github.com/WenyXu/gokit-istio-trace-adapter"
)

...

func main() {
	svc := stringService{}

	uppercaseHandler := httptransport.NewServer(
		makeUppercaseEndpoint(svc),
		decodeUppercaseRequest,
		encodeResponse,
                //Add ServerBefore & ServerAfter func 
                httptransport.ServerBefore(traceAdapter.DefaultServerBefore),
		httptransport.ServerAfter(traceAdapter.DefaultServerAfter),
	)

	http.Handle("/uppercase", uppercaseHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
```
