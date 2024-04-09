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

	// EnterDataPartList is called when entering the dataPartList production.
	EnterDataPartList(c *DataPartListContext)

	// EnterSettingsDef is called when entering the settingsDef production.
	EnterSettingsDef(c *SettingsDefContext)

	// EnterSettingsValue is called when entering the settingsValue production.
	EnterSettingsValue(c *SettingsValueContext)

	// EnterFuncDef is called when entering the funcDef production.
	EnterFuncDef(c *FuncDefContext)

	// EnterInnerFuncDef is called when entering the innerFuncDef production.
	EnterInnerFuncDef(c *InnerFuncDefContext)

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

	// EnterStatList is called when entering the statList production.
	EnterStatList(c *StatListContext)

	// EnterStat is called when entering the stat production.
	EnterStat(c *StatContext)

	// EnterVarDef is called when entering the varDef production.
	EnterVarDef(c *VarDefContext)

	// EnterIfStat is called when entering the ifStat production.
	EnterIfStat(c *IfStatContext)

	// EnterReturnStat is called when entering the returnStat production.
	EnterReturnStat(c *ReturnStatContext)

	// EnterContinueStat is called when entering the continueStat production.
	EnterContinueStat(c *ContinueStatContext)

	// EnterBreakStat is called when entering the breakStat production.
	EnterBreakStat(c *BreakStatContext)

	// EnterWhileStat is called when entering the whileStat production.
	EnterWhileStat(c *WhileStatContext)

	// EnterErrorStat is called when entering the errorStat production.
	EnterErrorStat(c *ErrorStatContext)

	// EnterTypeName is called when entering the typeName production.
	EnterTypeName(c *TypeNameContext)

	// EnterSliceStat is called when entering the sliceStat production.
	EnterSliceStat(c *SliceStatContext)

	// EnterIndexNumber is called when entering the indexNumber production.
	EnterIndexNumber(c *IndexNumberContext)

	// EnterArrayStat is called when entering the arrayStat production.
	EnterArrayStat(c *ArrayStatContext)

	// EnterArrayList is called when entering the arrayList production.
	EnterArrayList(c *ArrayListContext)

	// EnterArrayValue is called when entering the arrayValue production.
	EnterArrayValue(c *ArrayValueContext)

	// EnterIndexStat is called when entering the indexStat production.
	EnterIndexStat(c *IndexStatContext)

	// EnterObjectStat is called when entering the objectStat production.
	EnterObjectStat(c *ObjectStatContext)

	// EnterPairList is called when entering the pairList production.
	EnterPairList(c *PairListContext)

	// EnterPair is called when entering the pair production.
	EnterPair(c *PairContext)

	// EnterPairValue is called when entering the pairValue production.
	EnterPairValue(c *PairValueContext)

	// EnterArguments is called when entering the arguments production.
	EnterArguments(c *ArgumentsContext)

	// EnterArgumentsList is called when entering the argumentsList production.
	EnterArgumentsList(c *ArgumentsListContext)

	// EnterSimpleStat is called when entering the simpleStat production.
	EnterSimpleStat(c *SimpleStatContext)

	// EnterIncDecStat is called when entering the incDecStat production.
	EnterIncDecStat(c *IncDecStatContext)

	// EnterExprStat is called when entering the exprStat production.
	EnterExprStat(c *ExprStatContext)

	// EnterAssignMapArrStat is called when entering the assignMapArrStat production.
	EnterAssignMapArrStat(c *AssignMapArrStatContext)

	// EnterInitMapArrStat is called when entering the initMapArrStat production.
	EnterInitMapArrStat(c *InitMapArrStatContext)

	// EnterAssignment is called when entering the assignment production.
	EnterAssignment(c *AssignmentContext)

	// EnterPrimaryExpr is called when entering the primaryExpr production.
	EnterPrimaryExpr(c *PrimaryExprContext)

	// EnterOperand is called when entering the operand production.
	EnterOperand(c *OperandContext)

	// EnterExprList is called when entering the exprList production.
	EnterExprList(c *ExprListContext)

	// EnterExpr is called when entering the expr production.
	EnterExpr(c *ExprContext)

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

	// EnterIdentifierFull is called when entering the identifierFull production.
	EnterIdentifierFull(c *IdentifierFullContext)

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

	// ExitDataPartList is called when exiting the dataPartList production.
	ExitDataPartList(c *DataPartListContext)

	// ExitSettingsDef is called when exiting the settingsDef production.
	ExitSettingsDef(c *SettingsDefContext)

	// ExitSettingsValue is called when exiting the settingsValue production.
	ExitSettingsValue(c *SettingsValueContext)

	// ExitFuncDef is called when exiting the funcDef production.
	ExitFuncDef(c *FuncDefContext)

	// ExitInnerFuncDef is called when exiting the innerFuncDef production.
	ExitInnerFuncDef(c *InnerFuncDefContext)

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

	// ExitStatList is called when exiting the statList production.
	ExitStatList(c *StatListContext)

	// ExitStat is called when exiting the stat production.
	ExitStat(c *StatContext)

	// ExitVarDef is called when exiting the varDef production.
	ExitVarDef(c *VarDefContext)

	// ExitIfStat is called when exiting the ifStat production.
	ExitIfStat(c *IfStatContext)

	// ExitReturnStat is called when exiting the returnStat production.
	ExitReturnStat(c *ReturnStatContext)

	// ExitContinueStat is called when exiting the continueStat production.
	ExitContinueStat(c *ContinueStatContext)

	// ExitBreakStat is called when exiting the breakStat production.
	ExitBreakStat(c *BreakStatContext)

	// ExitWhileStat is called when exiting the whileStat production.
	ExitWhileStat(c *WhileStatContext)

	// ExitErrorStat is called when exiting the errorStat production.
	ExitErrorStat(c *ErrorStatContext)

	// ExitTypeName is called when exiting the typeName production.
	ExitTypeName(c *TypeNameContext)

	// ExitSliceStat is called when exiting the sliceStat production.
	ExitSliceStat(c *SliceStatContext)

	// ExitIndexNumber is called when exiting the indexNumber production.
	ExitIndexNumber(c *IndexNumberContext)

	// ExitArrayStat is called when exiting the arrayStat production.
	ExitArrayStat(c *ArrayStatContext)

	// ExitArrayList is called when exiting the arrayList production.
	ExitArrayList(c *ArrayListContext)

	// ExitArrayValue is called when exiting the arrayValue production.
	ExitArrayValue(c *ArrayValueContext)

	// ExitIndexStat is called when exiting the indexStat production.
	ExitIndexStat(c *IndexStatContext)

	// ExitObjectStat is called when exiting the objectStat production.
	ExitObjectStat(c *ObjectStatContext)

	// ExitPairList is called when exiting the pairList production.
	ExitPairList(c *PairListContext)

	// ExitPair is called when exiting the pair production.
	ExitPair(c *PairContext)

	// ExitPairValue is called when exiting the pairValue production.
	ExitPairValue(c *PairValueContext)

	// ExitArguments is called when exiting the arguments production.
	ExitArguments(c *ArgumentsContext)

	// ExitArgumentsList is called when exiting the argumentsList production.
	ExitArgumentsList(c *ArgumentsListContext)

	// ExitSimpleStat is called when exiting the simpleStat production.
	ExitSimpleStat(c *SimpleStatContext)

	// ExitIncDecStat is called when exiting the incDecStat production.
	ExitIncDecStat(c *IncDecStatContext)

	// ExitExprStat is called when exiting the exprStat production.
	ExitExprStat(c *ExprStatContext)

	// ExitAssignMapArrStat is called when exiting the assignMapArrStat production.
	ExitAssignMapArrStat(c *AssignMapArrStatContext)

	// ExitInitMapArrStat is called when exiting the initMapArrStat production.
	ExitInitMapArrStat(c *InitMapArrStatContext)

	// ExitAssignment is called when exiting the assignment production.
	ExitAssignment(c *AssignmentContext)

	// ExitPrimaryExpr is called when exiting the primaryExpr production.
	ExitPrimaryExpr(c *PrimaryExprContext)

	// ExitOperand is called when exiting the operand production.
	ExitOperand(c *OperandContext)

	// ExitExprList is called when exiting the exprList production.
	ExitExprList(c *ExprListContext)

	// ExitExpr is called when exiting the expr production.
	ExitExpr(c *ExprContext)

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

	// ExitIdentifierFull is called when exiting the identifierFull production.
	ExitIdentifierFull(c *IdentifierFullContext)

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
