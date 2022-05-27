package nodes

type TranslateVisitor interface {
	TranslateAndNode(node AndNode) string
	TranslateArrayNode(node ArrayNode) string
	TranslateEqNode(node EqNode) string
	TranslateGeNode(node GeNode) string
	TranslateGtNode(node GtNode) string
	TranslateInNode(node InNode) string
	TranslateLeNode(node LeNode) string
	TranslateLtNode(node LtNode) string
	TranslateNameNode(node NameNode) string
	TranslateNeNode(node NeNode) string
	TranslateNotNode(node NotNode) string
	TranslateNullNode(node NullNode) string
	TranslateOrNode(node OrNode) string
	TranslateStringNode(node StringNode) string
	TranslateValueNode(node ValueNode) string
	TranslateTimeNode(n TimeNode) string
}
