package flowIDGenerator

type DefaultFlowIDGenerator struct {}

func (dfig DefaultFlowIDGenerator) New() string {
	return "ABCD"
}

func NewDefaultFlowIDGenerator() DefaultFlowIDGenerator{
	return DefaultFlowIDGenerator{}
}