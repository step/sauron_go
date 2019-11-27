package flowidgenerator

// FlowIDGenerator is a interface for generating flowID
type FlowIDGenerator interface{
	New() string
}
