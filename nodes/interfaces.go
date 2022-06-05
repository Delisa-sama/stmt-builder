package nodes

// Node abstract node that can accept translate visitor
type Node interface {
	Accept(visitor TranslateVisitor) string
}

// NodeWithChilds abstract node with childs
type NodeWithChilds interface {
	Node
	Childs() []Node
}

// NodeWithName abstract node with name
type NodeWithName interface {
	Node
	Name() NameNode
}

// NodeWithValue abstract node with value
type NodeWithValue interface {
	Node
	Value() interface{}
}

// TranslateVisitor abstract visitor that translates nodes to some string
type TranslateVisitor interface {
	TranslateAndNode(AndNode) string
	TranslateArrayNode(ArrayNode) string
	TranslateEqNode(EqNode) string
	TranslateGeNode(GeNode) string
	TranslateGtNode(GtNode) string
	TranslateInNode(InNode) string
	TranslateNotInNode(NotInNode) string
	TranslateLeNode(LeNode) string
	TranslateLtNode(LtNode) string
	TranslateNameNode(NameNode) string
	TranslateNeNode(NeNode) string
	TranslateNotNode(NotNode) string
	TranslateNullNode(NullNode) string
	TranslateOrNode(OrNode) string
	TranslateStringNode(StringNode) string
	TranslateValueNode(ValueNode) string
	TranslateTimeNode(TimeNode) string
	TranslateIntNode(IntNode) string
	TranslateFloatNode(FloatNode) string
	TranslateUintNode(UintNode) string
}
