package tracing

type ITracing interface {
	Log(inTrace ...*TraceNode)
}