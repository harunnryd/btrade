package olymptrade

import "context"

// WebsocketServiceClient ...
type WebsocketServiceClient interface {
	// SetOptions ...
	SetOptions(options Options)

	// StartCandleStream ...
	StartCandleStream(ctx context.Context)

	// Analysis ...
	Analysis(ctx context.Context, df DataFrame)

	// Buy ...
	Buy(ctx context.Context, dir string)
}
