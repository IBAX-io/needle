// Code generated from NeedleParser.g4 by ANTLR 4.13.1. DO NOT EDIT.

package needle // NeedleParser
import "github.com/antlr4-go/antlr/v4"

// BaseNeedleParserListener is a complete listener for a parse tree produced by NeedleParser.
type BaseNeedleParserListener struct{}

var _ NeedleParserListener = &BaseNeedleParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseNeedleParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseNeedleParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseNeedleParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseNeedleParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterSourceMain is called when production sourceMain is entered.
func (s *BaseNeedleParserListener) EnterSourceMain(ctx *SourceMainContext) {}

// ExitSourceMain is called when production sourceMain is exited.
func (s *BaseNeedleParserListener) ExitSourceMain(ctx *SourceMainContext) {}

// EnterContractDef is called when production contractDef is entered.
func (s *BaseNeedleParserListener) EnterContractDef(ctx *ContractDefContext) {}

// ExitContractDef is called when production contractDef is exited.
func (s *BaseNeedleParserListener) ExitContractDef(ctx *ContractDefContext) {}

// EnterContractPart is called when production contractPart is entered.
func (s *BaseNeedleParserListener) EnterContractPart(ctx *ContractPartContext) {}

// ExitContractPart is called when production contractPart is exited.
func (s *BaseNeedleParserListener) ExitContractPart(ctx *ContractPartContext) {}

// EnterDataDef is called when production dataDef is entered.
func (s *BaseNeedleParserListener) EnterDataDef(ctx *DataDefContext) {}

// ExitDataDef is called when production dataDef is exited.
func (s *BaseNeedleParserListener) ExitDataDef(ctx *DataDefContext) {}

// EnterDataPartList is called when production dataPartList is entered.
func (s *BaseNeedleParserListener) EnterDataPartList(ctx *DataPartListContext) {}

// ExitDataPartList is called when production dataPartList is exited.
func (s *BaseNeedleParserListener) ExitDataPartList(ctx *DataPartListContext) {}

// EnterSettingsDef is called when production settingsDef is entered.
func (s *BaseNeedleParserListener) EnterSettingsDef(ctx *SettingsDefContext) {}

// ExitSettingsDef is called when production settingsDef is exited.
func (s *BaseNeedleParserListener) ExitSettingsDef(ctx *SettingsDefContext) {}

// EnterFuncDef is called when production funcDef is entered.
func (s *BaseNeedleParserListener) EnterFuncDef(ctx *FuncDefContext) {}

// ExitFuncDef is called when production funcDef is exited.
func (s *BaseNeedleParserListener) ExitFuncDef(ctx *FuncDefContext) {}

// EnterDefaultFuncDef is called when production defaultFuncDef is entered.
func (s *BaseNeedleParserListener) EnterDefaultFuncDef(ctx *DefaultFuncDefContext) {}

// ExitDefaultFuncDef is called when production defaultFuncDef is exited.
func (s *BaseNeedleParserListener) ExitDefaultFuncDef(ctx *DefaultFuncDefContext) {}

// EnterFuncDescriptor is called when production funcDescriptor is entered.
func (s *BaseNeedleParserListener) EnterFuncDescriptor(ctx *FuncDescriptorContext) {}

// ExitFuncDescriptor is called when production funcDescriptor is exited.
func (s *BaseNeedleParserListener) ExitFuncDescriptor(ctx *FuncDescriptorContext) {}

// EnterFuncSignature is called when production funcSignature is entered.
func (s *BaseNeedleParserListener) EnterFuncSignature(ctx *FuncSignatureContext) {}

// ExitFuncSignature is called when production funcSignature is exited.
func (s *BaseNeedleParserListener) ExitFuncSignature(ctx *FuncSignatureContext) {}

// EnterFuncTail is called when production funcTail is entered.
func (s *BaseNeedleParserListener) EnterFuncTail(ctx *FuncTailContext) {}

// ExitFuncTail is called when production funcTail is exited.
func (s *BaseNeedleParserListener) ExitFuncTail(ctx *FuncTailContext) {}

// EnterParameterList is called when production parameterList is entered.
func (s *BaseNeedleParserListener) EnterParameterList(ctx *ParameterListContext) {}

// ExitParameterList is called when production parameterList is exited.
func (s *BaseNeedleParserListener) ExitParameterList(ctx *ParameterListContext) {}

// EnterParameter is called when production parameter is entered.
func (s *BaseNeedleParserListener) EnterParameter(ctx *ParameterContext) {}

// ExitParameter is called when production parameter is exited.
func (s *BaseNeedleParserListener) ExitParameter(ctx *ParameterContext) {}

// EnterReturnParameters is called when production returnParameters is entered.
func (s *BaseNeedleParserListener) EnterReturnParameters(ctx *ReturnParametersContext) {}

// ExitReturnParameters is called when production returnParameters is exited.
func (s *BaseNeedleParserListener) ExitReturnParameters(ctx *ReturnParametersContext) {}

// EnterBlock is called when production block is entered.
func (s *BaseNeedleParserListener) EnterBlock(ctx *BlockContext) {}

// ExitBlock is called when production block is exited.
func (s *BaseNeedleParserListener) ExitBlock(ctx *BlockContext) {}

// EnterStatementList is called when production statementList is entered.
func (s *BaseNeedleParserListener) EnterStatementList(ctx *StatementListContext) {}

// ExitStatementList is called when production statementList is exited.
func (s *BaseNeedleParserListener) ExitStatementList(ctx *StatementListContext) {}

// EnterStatement is called when production statement is entered.
func (s *BaseNeedleParserListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BaseNeedleParserListener) ExitStatement(ctx *StatementContext) {}

// EnterSimpleStmt is called when production simpleStmt is entered.
func (s *BaseNeedleParserListener) EnterSimpleStmt(ctx *SimpleStmtContext) {}

// ExitSimpleStmt is called when production simpleStmt is exited.
func (s *BaseNeedleParserListener) ExitSimpleStmt(ctx *SimpleStmtContext) {}

// EnterIncDecStmt is called when production incDecStmt is entered.
func (s *BaseNeedleParserListener) EnterIncDecStmt(ctx *IncDecStmtContext) {}

// ExitIncDecStmt is called when production incDecStmt is exited.
func (s *BaseNeedleParserListener) ExitIncDecStmt(ctx *IncDecStmtContext) {}

// EnterAssignment is called when production assignment is entered.
func (s *BaseNeedleParserListener) EnterAssignment(ctx *AssignmentContext) {}

// ExitAssignment is called when production assignment is exited.
func (s *BaseNeedleParserListener) ExitAssignment(ctx *AssignmentContext) {}

// EnterVarDef is called when production varDef is entered.
func (s *BaseNeedleParserListener) EnterVarDef(ctx *VarDefContext) {}

// ExitVarDef is called when production varDef is exited.
func (s *BaseNeedleParserListener) ExitVarDef(ctx *VarDefContext) {}

// EnterIfStmt is called when production ifStmt is entered.
func (s *BaseNeedleParserListener) EnterIfStmt(ctx *IfStmtContext) {}

// ExitIfStmt is called when production ifStmt is exited.
func (s *BaseNeedleParserListener) ExitIfStmt(ctx *IfStmtContext) {}

// EnterIfBody is called when production ifBody is entered.
func (s *BaseNeedleParserListener) EnterIfBody(ctx *IfBodyContext) {}

// ExitIfBody is called when production ifBody is exited.
func (s *BaseNeedleParserListener) ExitIfBody(ctx *IfBodyContext) {}

// EnterElseBody is called when production elseBody is entered.
func (s *BaseNeedleParserListener) EnterElseBody(ctx *ElseBodyContext) {}

// ExitElseBody is called when production elseBody is exited.
func (s *BaseNeedleParserListener) ExitElseBody(ctx *ElseBodyContext) {}

// EnterReturnStmt is called when production returnStmt is entered.
func (s *BaseNeedleParserListener) EnterReturnStmt(ctx *ReturnStmtContext) {}

// ExitReturnStmt is called when production returnStmt is exited.
func (s *BaseNeedleParserListener) ExitReturnStmt(ctx *ReturnStmtContext) {}

// EnterContinueStmt is called when production continueStmt is entered.
func (s *BaseNeedleParserListener) EnterContinueStmt(ctx *ContinueStmtContext) {}

// ExitContinueStmt is called when production continueStmt is exited.
func (s *BaseNeedleParserListener) ExitContinueStmt(ctx *ContinueStmtContext) {}

// EnterBreakStmt is called when production breakStmt is entered.
func (s *BaseNeedleParserListener) EnterBreakStmt(ctx *BreakStmtContext) {}

// ExitBreakStmt is called when production breakStmt is exited.
func (s *BaseNeedleParserListener) ExitBreakStmt(ctx *BreakStmtContext) {}

// EnterWhileStmt is called when production whileStmt is entered.
func (s *BaseNeedleParserListener) EnterWhileStmt(ctx *WhileStmtContext) {}

// ExitWhileStmt is called when production whileStmt is exited.
func (s *BaseNeedleParserListener) ExitWhileStmt(ctx *WhileStmtContext) {}

// EnterErrorStmt is called when production errorStmt is entered.
func (s *BaseNeedleParserListener) EnterErrorStmt(ctx *ErrorStmtContext) {}

// ExitErrorStmt is called when production errorStmt is exited.
func (s *BaseNeedleParserListener) ExitErrorStmt(ctx *ErrorStmtContext) {}

// EnterArrayExpr is called when production arrayExpr is entered.
func (s *BaseNeedleParserListener) EnterArrayExpr(ctx *ArrayExprContext) {}

// ExitArrayExpr is called when production arrayExpr is exited.
func (s *BaseNeedleParserListener) ExitArrayExpr(ctx *ArrayExprContext) {}

// EnterArrayList is called when production arrayList is entered.
func (s *BaseNeedleParserListener) EnterArrayList(ctx *ArrayListContext) {}

// ExitArrayList is called when production arrayList is exited.
func (s *BaseNeedleParserListener) ExitArrayList(ctx *ArrayListContext) {}

// EnterMapExpr is called when production mapExpr is entered.
func (s *BaseNeedleParserListener) EnterMapExpr(ctx *MapExprContext) {}

// ExitMapExpr is called when production mapExpr is exited.
func (s *BaseNeedleParserListener) ExitMapExpr(ctx *MapExprContext) {}

// EnterPairList is called when production pairList is entered.
func (s *BaseNeedleParserListener) EnterPairList(ctx *PairListContext) {}

// ExitPairList is called when production pairList is exited.
func (s *BaseNeedleParserListener) ExitPairList(ctx *PairListContext) {}

// EnterPair is called when production pair is entered.
func (s *BaseNeedleParserListener) EnterPair(ctx *PairContext) {}

// ExitPair is called when production pair is exited.
func (s *BaseNeedleParserListener) ExitPair(ctx *PairContext) {}

// EnterPairValue is called when production pairValue is entered.
func (s *BaseNeedleParserListener) EnterPairValue(ctx *PairValueContext) {}

// ExitPairValue is called when production pairValue is exited.
func (s *BaseNeedleParserListener) ExitPairValue(ctx *PairValueContext) {}

// EnterArguments is called when production arguments is entered.
func (s *BaseNeedleParserListener) EnterArguments(ctx *ArgumentsContext) {}

// ExitArguments is called when production arguments is exited.
func (s *BaseNeedleParserListener) ExitArguments(ctx *ArgumentsContext) {}

// EnterExprList is called when production exprList is entered.
func (s *BaseNeedleParserListener) EnterExprList(ctx *ExprListContext) {}

// ExitExprList is called when production exprList is exited.
func (s *BaseNeedleParserListener) ExitExprList(ctx *ExprListContext) {}

// EnterExpr is called when production expr is entered.
func (s *BaseNeedleParserListener) EnterExpr(ctx *ExprContext) {}

// ExitExpr is called when production expr is exited.
func (s *BaseNeedleParserListener) ExitExpr(ctx *ExprContext) {}

// EnterPrimaryExpr is called when production primaryExpr is entered.
func (s *BaseNeedleParserListener) EnterPrimaryExpr(ctx *PrimaryExprContext) {}

// ExitPrimaryExpr is called when production primaryExpr is exited.
func (s *BaseNeedleParserListener) ExitPrimaryExpr(ctx *PrimaryExprContext) {}

// EnterIndexExpr is called when production indexExpr is entered.
func (s *BaseNeedleParserListener) EnterIndexExpr(ctx *IndexExprContext) {}

// ExitIndexExpr is called when production indexExpr is exited.
func (s *BaseNeedleParserListener) ExitIndexExpr(ctx *IndexExprContext) {}

// EnterSliceExpr is called when production sliceExpr is entered.
func (s *BaseNeedleParserListener) EnterSliceExpr(ctx *SliceExprContext) {}

// ExitSliceExpr is called when production sliceExpr is exited.
func (s *BaseNeedleParserListener) ExitSliceExpr(ctx *SliceExprContext) {}

// EnterOperand is called when production operand is entered.
func (s *BaseNeedleParserListener) EnterOperand(ctx *OperandContext) {}

// ExitOperand is called when production operand is exited.
func (s *BaseNeedleParserListener) ExitOperand(ctx *OperandContext) {}

// EnterLiteral is called when production literal is entered.
func (s *BaseNeedleParserListener) EnterLiteral(ctx *LiteralContext) {}

// ExitLiteral is called when production literal is exited.
func (s *BaseNeedleParserListener) ExitLiteral(ctx *LiteralContext) {}

// EnterTypeName is called when production typeName is entered.
func (s *BaseNeedleParserListener) EnterTypeName(ctx *TypeNameContext) {}

// ExitTypeName is called when production typeName is exited.
func (s *BaseNeedleParserListener) ExitTypeName(ctx *TypeNameContext) {}

// EnterIncDec_op is called when production incDec_op is entered.
func (s *BaseNeedleParserListener) EnterIncDec_op(ctx *IncDec_opContext) {}

// ExitIncDec_op is called when production incDec_op is exited.
func (s *BaseNeedleParserListener) ExitIncDec_op(ctx *IncDec_opContext) {}

// EnterMul_op is called when production mul_op is entered.
func (s *BaseNeedleParserListener) EnterMul_op(ctx *Mul_opContext) {}

// ExitMul_op is called when production mul_op is exited.
func (s *BaseNeedleParserListener) ExitMul_op(ctx *Mul_opContext) {}

// EnterUnary_op is called when production unary_op is entered.
func (s *BaseNeedleParserListener) EnterUnary_op(ctx *Unary_opContext) {}

// ExitUnary_op is called when production unary_op is exited.
func (s *BaseNeedleParserListener) ExitUnary_op(ctx *Unary_opContext) {}

// EnterAdd_op is called when production add_op is entered.
func (s *BaseNeedleParserListener) EnterAdd_op(ctx *Add_opContext) {}

// ExitAdd_op is called when production add_op is exited.
func (s *BaseNeedleParserListener) ExitAdd_op(ctx *Add_opContext) {}

// EnterLogical_op is called when production logical_op is entered.
func (s *BaseNeedleParserListener) EnterLogical_op(ctx *Logical_opContext) {}

// ExitLogical_op is called when production logical_op is exited.
func (s *BaseNeedleParserListener) ExitLogical_op(ctx *Logical_opContext) {}

// EnterRel_op is called when production rel_op is entered.
func (s *BaseNeedleParserListener) EnterRel_op(ctx *Rel_opContext) {}

// ExitRel_op is called when production rel_op is exited.
func (s *BaseNeedleParserListener) ExitRel_op(ctx *Rel_opContext) {}

// EnterAssign_op is called when production assign_op is entered.
func (s *BaseNeedleParserListener) EnterAssign_op(ctx *Assign_opContext) {}

// ExitAssign_op is called when production assign_op is exited.
func (s *BaseNeedleParserListener) ExitAssign_op(ctx *Assign_opContext) {}

// EnterIdentifierFull is called when production identifierFull is entered.
func (s *BaseNeedleParserListener) EnterIdentifierFull(ctx *IdentifierFullContext) {}

// ExitIdentifierFull is called when production identifierFull is exited.
func (s *BaseNeedleParserListener) ExitIdentifierFull(ctx *IdentifierFullContext) {}

// EnterIdentifierVar is called when production identifierVar is entered.
func (s *BaseNeedleParserListener) EnterIdentifierVar(ctx *IdentifierVarContext) {}

// ExitIdentifierVar is called when production identifierVar is exited.
func (s *BaseNeedleParserListener) ExitIdentifierVar(ctx *IdentifierVarContext) {}

// EnterIdentifierList is called when production identifierList is entered.
func (s *BaseNeedleParserListener) EnterIdentifierList(ctx *IdentifierListContext) {}

// ExitIdentifierList is called when production identifierList is exited.
func (s *BaseNeedleParserListener) ExitIdentifierList(ctx *IdentifierListContext) {}

// EnterStringLiteral is called when production stringLiteral is entered.
func (s *BaseNeedleParserListener) EnterStringLiteral(ctx *StringLiteralContext) {}

// ExitStringLiteral is called when production stringLiteral is exited.
func (s *BaseNeedleParserListener) ExitStringLiteral(ctx *StringLiteralContext) {}

// EnterNumberLiteral is called when production numberLiteral is entered.
func (s *BaseNeedleParserListener) EnterNumberLiteral(ctx *NumberLiteralContext) {}

// ExitNumberLiteral is called when production numberLiteral is exited.
func (s *BaseNeedleParserListener) ExitNumberLiteral(ctx *NumberLiteralContext) {}

// EnterBooleanLiteral is called when production booleanLiteral is entered.
func (s *BaseNeedleParserListener) EnterBooleanLiteral(ctx *BooleanLiteralContext) {}

// ExitBooleanLiteral is called when production booleanLiteral is exited.
func (s *BaseNeedleParserListener) ExitBooleanLiteral(ctx *BooleanLiteralContext) {}

// EnterEos is called when production eos is entered.
func (s *BaseNeedleParserListener) EnterEos(ctx *EosContext) {}

// ExitEos is called when production eos is exited.
func (s *BaseNeedleParserListener) ExitEos(ctx *EosContext) {}
