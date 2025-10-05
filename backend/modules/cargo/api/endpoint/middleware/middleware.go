package middleware

import (
	"github.com/chaos-io/cargo/backend/modules/cargo/api/endpoint"
	"github.com/chaos-io/gokit/middleware"
	"github.com/chaos-io/gokit/tracing"
	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	"go.opentelemetry.io/otel/trace"

	// this service api
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

	var tracer trace.Tracer
	if value, ok := options["tracer"]; ok && value != nil {
		tracer = value.(trace.Tracer)
	}
	var count *kitprometheus.Counter
	if value, ok := options["count"]; ok && value != nil {
		count = value.(*kitprometheus.Counter)
	}
	var latency *kitprometheus.Histogram
	if value, ok := options["latency"]; ok && value != nil {
		latency = value.(*kitprometheus.Histogram)
	}
	// var validator *middleware.Validator
	// if value, ok := options["validator"]; ok && value != nil {
	// 	validator = value.(*middleware.Validator)
	// }

	{ // CreateResource
		if tracer != nil {
			in.CreateResourceEndpoint = tracing.TraceServer(tracer, "CreateResource")(in.CreateResourceEndpoint)
		}
		if count != nil && latency != nil {
			in.CreateResourceEndpoint = middleware.Instrumenting(latency.With("method", "CreateResource"), count.With("method", "CreateResource"))(in.CreateResourceEndpoint)
		}
		// if validator != nil {
		// 	in.CreateResourceEndpoint = validator.Validate()(in.CreateResourceEndpoint)
		// }
	}
	{ // GetResource
		if tracer != nil {
			in.GetResourceEndpoint = tracing.TraceServer(tracer, "GetResource")(in.GetResourceEndpoint)
		}
		if count != nil && latency != nil {
			in.GetResourceEndpoint = middleware.Instrumenting(latency.With("method", "GetResource"), count.With("method", "GetResource"))(in.GetResourceEndpoint)
		}
		// if validator != nil {
		// 	in.GetResourceEndpoint = validator.Validate()(in.GetResourceEndpoint)
		// }
	}
	{ // ListResources
		if tracer != nil {
			in.ListResourcesEndpoint = tracing.TraceServer(tracer, "ListResources")(in.ListResourcesEndpoint)
		}
		if count != nil && latency != nil {
			in.ListResourcesEndpoint = middleware.Instrumenting(latency.With("method", "ListResources"), count.With("method", "ListResources"))(in.ListResourcesEndpoint)
		}
		// if validator != nil {
		// 	in.ListResourcesEndpoint = validator.Validate()(in.ListResourcesEndpoint)
		// }
	}
	{ // UpdateResource
		if tracer != nil {
			in.UpdateResourceEndpoint = tracing.TraceServer(tracer, "UpdateResource")(in.UpdateResourceEndpoint)
		}
		if count != nil && latency != nil {
			in.UpdateResourceEndpoint = middleware.Instrumenting(latency.With("method", "UpdateResource"), count.With("method", "UpdateResource"))(in.UpdateResourceEndpoint)
		}
		// if validator != nil {
		// 	in.UpdateResourceEndpoint = validator.Validate()(in.UpdateResourceEndpoint)
		// }
	}
	{ // DeleteResource
		if tracer != nil {
			in.DeleteResourceEndpoint = tracing.TraceServer(tracer, "DeleteResource")(in.DeleteResourceEndpoint)
		}
		if count != nil && latency != nil {
			in.DeleteResourceEndpoint = middleware.Instrumenting(latency.With("method", "DeleteResource"), count.With("method", "DeleteResource"))(in.DeleteResourceEndpoint)
		}
		// if validator != nil {
		// 	in.DeleteResourceEndpoint = validator.Validate()(in.DeleteResourceEndpoint)
		// }
	}

	return in
}

func WrapService(in pb.CargoServer, options map[string]interface{}) pb.CargoServer {
	return in
}
