package recovery

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// wrappedStream ensures that SetTrailer calls are delegated to the underlying stream.
type wrappedStream struct {
	grpc.ServerStream
}

func (w *wrappedStream) SetTrailer(md metadata.MD) {
	w.ServerStream.SetTrailer(md)
}

// The recovery interceptor logic needs to ensure that the context used in the 
// recovery handler is aware of the stream so that grpc.SetTrailer(ctx, ...) works.
// We ensure the stream is wrapped to delegate SetTrailer correctly.