// Code generated from NeedleParser.g4 by ANTLR 4.13.1. DO NOT EDIT.

package needle // NeedleParser
import "github.com/antlr4-go/antlr/v4"

// NeedleParserListener is a complete listener for a parse tree produced by NeedleParser.
type NeedleParserListener interface {
	antlr.ParseTreeListener

	// EnterSourceMain is called when entering the sourceMain production.
	EnterSourceMain(c *SourceMainContext)

	// EnterContractDef is called when entering the contractDef production.
	EnterContractDef(c *ContractDefContext)

	// EnterContractPart is called when entering the contractPart production.
	EnterContractPart(c *ContractPartContext)

	// EnterDataDef is called when entering the dataDef production.
	EnterDataDef(c *DataDefContext)

	// EnterDataPart is called when entering the dataPart production.
	EnterDataPart(c *DataPartContext)

	// EnterSettingsDef is called when entering the settingsDef production.
	EnterSettingsDef(c *SettingsDefContext)

	// EnterFuncDef is called when entering the funcDef production.
	EnterFuncDef(c *FuncDefContext)

	// EnterDefaultFuncDef is called when entering the defaultFuncDef production.
	EnterDefaultFuncDef(c *DefaultFuncDefContext)

	// EnterFuncDescriptor is called when entering the funcDescriptor production.
	EnterFuncDescriptor(c *FuncDescriptorContext)

	// EnterFuncSignature is called when entering the funcSignature production.
	EnterFuncSignature(c *FuncSignatureContext)

	// EnterFuncTail is called when entering the funcTail production.
	EnterFuncTail(c *FuncTailContext)

	// EnterParameterList is called when entering the parameterList production.
	EnterParameterList(c *ParameterListContext)

	// EnterParameter is called when entering the parameter production.
	EnterParameter(c *ParameterContext)

	// EnterReturnParameters is called when entering the returnParameters production.
	EnterReturnParameters(c *ReturnParametersContext)

	// EnterBlock is called when entering the block production.
	EnterBlock(c *BlockContext)

	// EnterStatementList is called when entering the statementList production.
	EnterStatementList(c *StatementListContext)

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

	// EnterSimpleStmt is called when entering the simpleStmt production.
	EnterSimpleStmt(c *SimpleStmtContext)

	// EnterIncDecStmt is called when entering the incDecStmt production.
	EnterIncDecStmt(c *IncDecStmtContext)

	// EnterAssignment is called when entering the assignment production.
	EnterAssignment(c *AssignmentContext)

	// EnterVarDef is called when entering the varDef production.
	EnterVarDef(c *VarDefContext)

	// EnterIfStmt is called when entering the ifStmt production.
	EnterIfStmt(c *IfStmtContext)

	// EnterIfBody is called when entering the ifBody production.
	EnterIfBody(c *IfBodyContext)

	// EnterElseBody is called when entering the elseBody production.
	EnterElseBody(c *ElseBodyContext)

	// EnterReturnStmt is called when entering the returnStmt production.
	EnterReturnStmt(c *ReturnStmtContext)

	// EnterContinueStmt is called when entering the continueStmt production.
	EnterContinueStmt(c *ContinueStmtContext)

	// EnterBreakStmt is called when entering the breakStmt production.
	EnterBreakStmt(c *BreakStmtContext)

	// EnterWhileStmt is called when entering the whileStmt production.
	EnterWhileStmt(c *WhileStmtContext)

	// EnterErrorStmt is called when entering the errorStmt production.
	EnterErrorStmt(c *ErrorStmtContext)

	// EnterArrayExpr is called when entering the arrayExpr production.
	EnterArrayExpr(c *ArrayExprContext)

	// EnterArrayList is called when entering the arrayList production.
	EnterArrayList(c *ArrayListContext)

	// EnterMapExpr is called when entering the mapExpr production.
	EnterMapExpr(c *MapExprContext)

	// EnterPairList is called when entering the pairList production.
	EnterPairList(c *PairListContext)

	// EnterPair is called when entering the pair production.
	EnterPair(c *PairContext)

	// EnterArguments is called when entering the arguments production.
	EnterArguments(c *ArgumentsContext)

	// EnterExprList is called when entering the exprList production.
	EnterExprList(c *ExprListContext)

	// EnterExpr is called when entering the expr production.
	EnterExpr(c *ExprContext)

	// EnterPrimaryExpr is called when entering the primaryExpr production.
	EnterPrimaryExpr(c *PrimaryExprContext)

	// EnterIndexExpr is called when entering the indexExpr production.
	EnterIndexExpr(c *IndexExprContext)

	// EnterSliceExpr is called when entering the sliceExpr production.
	EnterSliceExpr(c *SliceExprContext)

	// EnterContractCall is called when entering the contractCall production.
	EnterContractCall(c *ContractCallContext)

	// EnterOperand is called when entering the operand production.
	EnterOperand(c *OperandContext)

	// EnterLiteral is called when entering the literal production.
	EnterLiteral(c *LiteralContext)

	// EnterTypeName is called when entering the typeName production.
	EnterTypeName(c *TypeNameContext)

	// EnterIncDec_op is called when entering the incDec_op production.
	EnterIncDec_op(c *IncDec_opContext)

	// EnterMul_op is called when entering the mul_op production.
	EnterMul_op(c *Mul_opContext)

	// EnterUnary_op is called when entering the unary_op production.
	EnterUnary_op(c *Unary_opContext)

	// EnterAdd_op is called when entering the add_op production.
	EnterAdd_op(c *Add_opContext)

	// EnterLogical_op is called when entering the logical_op production.
	EnterLogical_op(c *Logical_opContext)

	// EnterRel_op is called when entering the rel_op production.
	EnterRel_op(c *Rel_opContext)

	// EnterAssign_op is called when entering the assign_op production.
	EnterAssign_op(c *Assign_opContext)

	// EnterIdentifierVar is called when entering the identifierVar production.
	EnterIdentifierVar(c *IdentifierVarContext)

	// EnterIdentifierList is called when entering the identifierList production.
	EnterIdentifierList(c *IdentifierListContext)

	// EnterStringLiteral is called when entering the stringLiteral production.
	EnterStringLiteral(c *StringLiteralContext)

	// EnterNumberLiteral is called when entering the numberLiteral production.
	EnterNumberLiteral(c *NumberLiteralContext)

	// EnterBooleanLiteral is called when entering the booleanLiteral production.
	EnterBooleanLiteral(c *BooleanLiteralContext)

	// EnterEos is called when entering the eos production.
	EnterEos(c *EosContext)

	// ExitSourceMain is called when exiting the sourceMain production.
	ExitSourceMain(c *SourceMainContext)

	// ExitContractDef is called when exiting the contractDef production.
	ExitContractDef(c *ContractDefContext)

	// ExitContractPart is called when exiting the contractPart production.
	ExitContractPart(c *ContractPartContext)

	// ExitDataDef is called when exiting the dataDef production.
	ExitDataDef(c *DataDefContext)

	// ExitDataPart is called when exiting the dataPart production.
	ExitDataPart(c *DataPartContext)

	// ExitSettingsDef is called when exiting the settingsDef production.
	ExitSettingsDef(c *SettingsDefContext)

	// ExitFuncDef is called when exiting the funcDef production.
	ExitFuncDef(c *FuncDefContext)

	// ExitDefaultFuncDef is called when exiting the defaultFuncDef production.
	ExitDefaultFuncDef(c *DefaultFuncDefContext)

	// ExitFuncDescriptor is called when exiting the funcDescriptor production.
	ExitFuncDescriptor(c *FuncDescriptorContext)

	// ExitFuncSignature is called when exiting the funcSignature production.
	ExitFuncSignature(c *FuncSignatureContext)

	// ExitFuncTail is called when exiting the funcTail production.
	ExitFuncTail(c *FuncTailContext)

	// ExitParameterList is called when exiting the parameterList production.
	ExitParameterList(c *ParameterListContext)

	// ExitParameter is called when exiting the parameter production.
	ExitParameter(c *ParameterContext)

	// ExitReturnParameters is called when exiting the returnParameters production.
	ExitReturnParameters(c *ReturnParametersContext)

	// ExitBlock is called when exiting the block production.
	ExitBlock(c *BlockContext)

	// ExitStatementList is called when exiting the statementList production.
	ExitStatementList(c *StatementListContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitSimpleStmt is called when exiting the simpleStmt production.
	ExitSimpleStmt(c *SimpleStmtContext)

	// ExitIncDecStmt is called when exiting the incDecStmt production.
	ExitIncDecStmt(c *IncDecStmtContext)

	// ExitAssignment is called when exiting the assignment production.
	ExitAssignment(c *AssignmentContext)

	// ExitVarDef is called when exiting the varDef production.
	ExitVarDef(c *VarDefContext)

	// ExitIfStmt is called when exiting the ifStmt production.
	ExitIfStmt(c *IfStmtContext)

	// ExitIfBody is called when exiting the ifBody production.
	ExitIfBody(c *IfBodyContext)

	// ExitElseBody is called when exiting the elseBody production.
	ExitElseBody(c *ElseBodyContext)

	// ExitReturnStmt is called when exiting the returnStmt production.
	ExitReturnStmt(c *ReturnStmtContext)

	// ExitContinueStmt is called when exiting the continueStmt production.
	ExitContinueStmt(c *ContinueStmtContext)

	// ExitBreakStmt is called when exiting the breakStmt production.
	ExitBreakStmt(c *BreakStmtContext)

	// ExitWhileStmt is called when exiting the whileStmt production.
	ExitWhileStmt(c *WhileStmtContext)

	// ExitErrorStmt is called when exiting the errorStmt production.
	ExitErrorStmt(c *ErrorStmtContext)

	// ExitArrayExpr is called when exiting the arrayExpr production.
	ExitArrayExpr(c *ArrayExprContext)

	// ExitArrayList is called when exiting the arrayList production.
	ExitArrayList(c *ArrayListContext)

	// ExitMapExpr is called when exiting the mapExpr production.
	ExitMapExpr(c *MapExprContext)

	// ExitPairList is called when exiting the pairList production.
	ExitPairList(c *PairListContext)

	// ExitPair is called when exiting the pair production.
	ExitPair(c *PairContext)

	// ExitArguments is called when exiting the arguments production.
	ExitArguments(c *ArgumentsContext)

	// ExitExprList is called when exiting the exprList production.
	ExitExprList(c *ExprListContext)

	// ExitExpr is called when exiting the expr production.
	ExitExpr(c *ExprContext)

	// ExitPrimaryExpr is called when exiting the primaryExpr production.
	ExitPrimaryExpr(c *PrimaryExprContext)

	// ExitIndexExpr is called when exiting the indexExpr production.
	ExitIndexExpr(c *IndexExprContext)

	// ExitSliceExpr is called when exiting the sliceExpr production.
	ExitSliceExpr(c *SliceExprContext)

	// ExitContractCall is called when exiting the contractCall production.
	ExitContractCall(c *ContractCallContext)

	// ExitOperand is called when exiting the operand production.
	ExitOperand(c *OperandContext)

	// ExitLiteral is called when exiting the literal production.
	ExitLiteral(c *LiteralContext)

	// ExitTypeName is called when exiting the typeName production.
	ExitTypeName(c *TypeNameContext)

	// ExitIncDec_op is called when exiting the incDec_op production.
	ExitIncDec_op(c *IncDec_opContext)

	// ExitMul_op is called when exiting the mul_op production.
	ExitMul_op(c *Mul_opContext)

	// ExitUnary_op is called when exiting the unary_op production.
	ExitUnary_op(c *Unary_opContext)

	// ExitAdd_op is called when exiting the add_op production.
	ExitAdd_op(c *Add_opContext)

	// ExitLogical_op is called when exiting the logical_op production.
	ExitLogical_op(c *Logical_opContext)

	// ExitRel_op is called when exiting the rel_op production.
	ExitRel_op(c *Rel_opContext)

	// ExitAssign_op is called when exiting the assign_op production.
	ExitAssign_op(c *Assign_opContext)

	// ExitIdentifierVar is called when exiting the identifierVar production.
	ExitIdentifierVar(c *IdentifierVarContext)

	// ExitIdentifierList is called when exiting the identifierList production.
	ExitIdentifierList(c *IdentifierListContext)

	// ExitStringLiteral is called when exiting the stringLiteral production.
	ExitStringLiteral(c *StringLiteralContext)

	// ExitNumberLiteral is called when exiting the numberLiteral production.
	ExitNumberLiteral(c *NumberLiteralContext)

	// ExitBooleanLiteral is called when exiting the booleanLiteral production.
	ExitBooleanLiteral(c *BooleanLiteralContext)

	// ExitEos is called when exiting the eos production.
	ExitEos(c *EosContext)
}
