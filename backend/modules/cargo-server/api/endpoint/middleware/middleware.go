package middleware

import (
	// this service api
	"github.com/chaos-io/cargo/backend/modules/cargo-server/api/endpoint"
	pb "github.com/chaos-io/cargo/genproto/cargo/v1"
)

// WrapEndpoints accepts the service's entire collection of endpoints, so that a
// set of middlewares can be wrapped around every middleware (e.g., access
// logging and instrumentation), and others wrapped selectively around some
// endpoints and not others (e.g., endpoints requiring authenticated access).
// Note that the final middleware wrapped will be the outermost middleware
// (i.e. applied first)
func WrapEndpoints(in endpoint.Endpoints, options map[string]interface{}) endpoint.Endpoints {

	// Pass a middleware you want applied to every endpoint.
	// optionally pass in endpoints by name that you want to be excluded
	// e.g.
	// in.WrapAllExcept(authMiddleware, "Status", "Ping")

	// Pass in a svc.LabeledMiddleware you want applied to every endpoint.
	// These middlewares get passed the endpoints name as their first argument when applied.
	// This can be used to write generic metric gathering middlewares that can
	// report the endpoint name for free.
	// in.WrapAllLabeledExcept(errorCounter(statsdCounter), "Status", "Ping")

	// How to apply a middleware to a single endpoint.
	// in.ExampleEndpoint = authMiddleware(in.ExampleEndpoint)

	wrapTracer(in, options)
	wrapInstrumenting(in, options)
	wrapValidator(in, options)

	return in
}

func WrapService(in pb.CargoServiceServer, options map[string]interface{}) pb.CargoServiceServer {
	return in
}
