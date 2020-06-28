/*
Copyright 2020 RS4
@Author: Weny Xu
@Date: 2020/06/28 20:53
*/

package traceAdapter

import (
	"context"
	http2 "net/http"
)

var (
	headerFlags = []string{"x-request-id", "x-b3-traceid", "x-b3-spanid", "x-b3-parentspanid", "x-b3-sampled", "x-b3-flags", "x-ot-span-context"}
)

func DefaultServerBefore(ctx context.Context, r *http2.Request) context.Context {
	return context.WithValue(ctx, "x-header", r.Header)
}

func DefaultServerAfter(ctx context.Context, w http2.ResponseWriter) context.Context {
	r, ok := ctx.Value("x-header").(http2.Header)
	if ok {
		for _, s := range headerFlags {
			w.Header().Set(s, r.Get(s))
		}
	}
	return ctx
}
