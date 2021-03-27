package tracing

type TraceNode struct {
	Key string
	Value string
	Next []*TraceNode
}

type tracingImpl struct {}

func (*tracingImpl) Log () {
	//
}

func New () ITracing {
	return &tracingImpl{}
}