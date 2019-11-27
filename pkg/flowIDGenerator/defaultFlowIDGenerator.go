package flowidgenerator

// DefaultFlowIDGenerator is a flowIDGenerator that
// generates a default flowID that makes it convinient for testing
type DefaultFlowIDGenerator struct {}

// New returns a generated flow id for DefaultFlowIDGenerator
func (dfig DefaultFlowIDGenerator) New() string {
	return "ABCD"
}

// NewDefaultFlowIDGenerator should be called when a
// DefaultFlowIDGenerator needs to be created
func NewDefaultFlowIDGenerator() DefaultFlowIDGenerator{
	return DefaultFlowIDGenerator{}
}