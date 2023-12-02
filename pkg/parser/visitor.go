package parser

type ExprVisitor interface {
	VisitListExpr(l ListExpr) (interface{}, error)
	VisitKeywordExpr(k Keyword) (interface{}, error)
	// VisitBinaryExpr(b Binary) (interface{}, error)
	VisitOperatorExpr(o Operator) (interface{}, error)
	VisitAtomExpr(l Atom) (interface{}, error)
	// VisitCondExpr(c Cond) (interface{}, error)
	// VisitUnaryExpr(u Unary) (interface{}, error)
	// VisitVariableExpr(v Variable) (interface{}, error)
	// VisitAssignExpr(a Assign) (interface{}, error)
	// VisitLogicalExpr(l Logical) (interface{}, error)
	// VisitCallExpr(c Call) (interface{}, error)
}

type StmtVisitor interface {
	VisitExprStmt(e ExprStmt) (interface{}, error)
	VisitSetStmt(v SetStmt) (interface{}, error)
	VisitFunctionStmt(f FunctionStmt) (interface{}, error)
}