package parser

type ExprVisitor interface {
	VisitListExpr(l ListExpr) (interface{}, error)
	VisitKeywordExpr(k Keyword) (interface{}, error)
	VisitOperatorExpr(o Operator) (interface{}, error)
	VisitAtomExpr(l Atom) (interface{}, error)
	VisitSymbolExpr(s Symbol) (interface{}, error)
	VisitFuncDefinitionExpr(f FuncDefinition) (interface{}, error)
	VisitCallExpr(c Call) (interface{}, error)
}