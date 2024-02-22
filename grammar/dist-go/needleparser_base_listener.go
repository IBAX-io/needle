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

// EnterSettingsValue is called when production settingsValue is entered.
func (s *BaseNeedleParserListener) EnterSettingsValue(ctx *SettingsValueContext) {}

// ExitSettingsValue is called when production settingsValue is exited.
func (s *BaseNeedleParserListener) ExitSettingsValue(ctx *SettingsValueContext) {}

// EnterFuncDef is called when production funcDef is entered.
func (s *BaseNeedleParserListener) EnterFuncDef(ctx *FuncDefContext) {}

// ExitFuncDef is called when production funcDef is exited.
func (s *BaseNeedleParserListener) ExitFuncDef(ctx *FuncDefContext) {}

// EnterInnerFuncDef is called when production innerFuncDef is entered.
func (s *BaseNeedleParserListener) EnterInnerFuncDef(ctx *InnerFuncDefContext) {}

// ExitInnerFuncDef is called when production innerFuncDef is exited.
func (s *BaseNeedleParserListener) ExitInnerFuncDef(ctx *InnerFuncDefContext) {}

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

// EnterStatList is called when production statList is entered.
func (s *BaseNeedleParserListener) EnterStatList(ctx *StatListContext) {}

// ExitStatList is called when production statList is exited.
func (s *BaseNeedleParserListener) ExitStatList(ctx *StatListContext) {}

// EnterStat is called when production stat is entered.
func (s *BaseNeedleParserListener) EnterStat(ctx *StatContext) {}

// ExitStat is called when production stat is exited.
func (s *BaseNeedleParserListener) ExitStat(ctx *StatContext) {}

// EnterVarDef is called when production varDef is entered.
func (s *BaseNeedleParserListener) EnterVarDef(ctx *VarDefContext) {}

// ExitVarDef is called when production varDef is exited.
func (s *BaseNeedleParserListener) ExitVarDef(ctx *VarDefContext) {}

// EnterIfStat is called when production ifStat is entered.
func (s *BaseNeedleParserListener) EnterIfStat(ctx *IfStatContext) {}

// ExitIfStat is called when production ifStat is exited.
func (s *BaseNeedleParserListener) ExitIfStat(ctx *IfStatContext) {}

// EnterReturnStat is called when production returnStat is entered.
func (s *BaseNeedleParserListener) EnterReturnStat(ctx *ReturnStatContext) {}

// ExitReturnStat is called when production returnStat is exited.
func (s *BaseNeedleParserListener) ExitReturnStat(ctx *ReturnStatContext) {}

// EnterContinueStat is called when production continueStat is entered.
func (s *BaseNeedleParserListener) EnterContinueStat(ctx *ContinueStatContext) {}

// ExitContinueStat is called when production continueStat is exited.
func (s *BaseNeedleParserListener) ExitContinueStat(ctx *ContinueStatContext) {}

// EnterBreakStat is called when production breakStat is entered.
func (s *BaseNeedleParserListener) EnterBreakStat(ctx *BreakStatContext) {}

// ExitBreakStat is called when production breakStat is exited.
func (s *BaseNeedleParserListener) ExitBreakStat(ctx *BreakStatContext) {}

// EnterWhileStat is called when production whileStat is entered.
func (s *BaseNeedleParserListener) EnterWhileStat(ctx *WhileStatContext) {}

// ExitWhileStat is called when production whileStat is exited.
func (s *BaseNeedleParserListener) ExitWhileStat(ctx *WhileStatContext) {}

// EnterErrorStat is called when production errorStat is entered.
func (s *BaseNeedleParserListener) EnterErrorStat(ctx *ErrorStatContext) {}

// ExitErrorStat is called when production errorStat is exited.
func (s *BaseNeedleParserListener) ExitErrorStat(ctx *ErrorStatContext) {}

// EnterTypeName is called when production typeName is entered.
func (s *BaseNeedleParserListener) EnterTypeName(ctx *TypeNameContext) {}

// ExitTypeName is called when production typeName is exited.
func (s *BaseNeedleParserListener) ExitTypeName(ctx *TypeNameContext) {}

// EnterSliceStat is called when production sliceStat is entered.
func (s *BaseNeedleParserListener) EnterSliceStat(ctx *SliceStatContext) {}

// ExitSliceStat is called when production sliceStat is exited.
func (s *BaseNeedleParserListener) ExitSliceStat(ctx *SliceStatContext) {}

// EnterIndexNumber is called when production indexNumber is entered.
func (s *BaseNeedleParserListener) EnterIndexNumber(ctx *IndexNumberContext) {}

// ExitIndexNumber is called when production indexNumber is exited.
func (s *BaseNeedleParserListener) ExitIndexNumber(ctx *IndexNumberContext) {}

// EnterArrayStat is called when production arrayStat is entered.
func (s *BaseNeedleParserListener) EnterArrayStat(ctx *ArrayStatContext) {}

// ExitArrayStat is called when production arrayStat is exited.
func (s *BaseNeedleParserListener) ExitArrayStat(ctx *ArrayStatContext) {}

// EnterArrayList is called when production arrayList is entered.
func (s *BaseNeedleParserListener) EnterArrayList(ctx *ArrayListContext) {}

// ExitArrayList is called when production arrayList is exited.
func (s *BaseNeedleParserListener) ExitArrayList(ctx *ArrayListContext) {}

// EnterArrayValue is called when production arrayValue is entered.
func (s *BaseNeedleParserListener) EnterArrayValue(ctx *ArrayValueContext) {}

// ExitArrayValue is called when production arrayValue is exited.
func (s *BaseNeedleParserListener) ExitArrayValue(ctx *ArrayValueContext) {}

// EnterIndexStat is called when production indexStat is entered.
func (s *BaseNeedleParserListener) EnterIndexStat(ctx *IndexStatContext) {}

// ExitIndexStat is called when production indexStat is exited.
func (s *BaseNeedleParserListener) ExitIndexStat(ctx *IndexStatContext) {}

// EnterObjectStat is called when production objectStat is entered.
func (s *BaseNeedleParserListener) EnterObjectStat(ctx *ObjectStatContext) {}

// ExitObjectStat is called when production objectStat is exited.
func (s *BaseNeedleParserListener) ExitObjectStat(ctx *ObjectStatContext) {}

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

// EnterArgumentsList is called when production argumentsList is entered.
func (s *BaseNeedleParserListener) EnterArgumentsList(ctx *ArgumentsListContext) {}

// ExitArgumentsList is called when production argumentsList is exited.
func (s *BaseNeedleParserListener) ExitArgumentsList(ctx *ArgumentsListContext) {}

// EnterSimpleStat is called when production simpleStat is entered.
func (s *BaseNeedleParserListener) EnterSimpleStat(ctx *SimpleStatContext) {}

// ExitSimpleStat is called when production simpleStat is exited.
func (s *BaseNeedleParserListener) ExitSimpleStat(ctx *SimpleStatContext) {}

// EnterIncDecStat is called when production incDecStat is entered.
func (s *BaseNeedleParserListener) EnterIncDecStat(ctx *IncDecStatContext) {}

// ExitIncDecStat is called when production incDecStat is exited.
func (s *BaseNeedleParserListener) ExitIncDecStat(ctx *IncDecStatContext) {}

// EnterExprStat is called when production exprStat is entered.
func (s *BaseNeedleParserListener) EnterExprStat(ctx *ExprStatContext) {}

// ExitExprStat is called when production exprStat is exited.
func (s *BaseNeedleParserListener) ExitExprStat(ctx *ExprStatContext) {}

// EnterAssignMapArrStat is called when production assignMapArrStat is entered.
func (s *BaseNeedleParserListener) EnterAssignMapArrStat(ctx *AssignMapArrStatContext) {}

// ExitAssignMapArrStat is called when production assignMapArrStat is exited.
func (s *BaseNeedleParserListener) ExitAssignMapArrStat(ctx *AssignMapArrStatContext) {}

// EnterInitMapArrStat is called when production initMapArrStat is entered.
func (s *BaseNeedleParserListener) EnterInitMapArrStat(ctx *InitMapArrStatContext) {}

// ExitInitMapArrStat is called when production initMapArrStat is exited.
func (s *BaseNeedleParserListener) ExitInitMapArrStat(ctx *InitMapArrStatContext) {}

// EnterAssignment is called when production assignment is entered.
func (s *BaseNeedleParserListener) EnterAssignment(ctx *AssignmentContext) {}

// ExitAssignment is called when production assignment is exited.
func (s *BaseNeedleParserListener) ExitAssignment(ctx *AssignmentContext) {}

// EnterPrimaryExpr is called when production primaryExpr is entered.
func (s *BaseNeedleParserListener) EnterPrimaryExpr(ctx *PrimaryExprContext) {}

// ExitPrimaryExpr is called when production primaryExpr is exited.
func (s *BaseNeedleParserListener) ExitPrimaryExpr(ctx *PrimaryExprContext) {}

// EnterOperand is called when production operand is entered.
func (s *BaseNeedleParserListener) EnterOperand(ctx *OperandContext) {}

// ExitOperand is called when production operand is exited.
func (s *BaseNeedleParserListener) ExitOperand(ctx *OperandContext) {}

// EnterExprList is called when production exprList is entered.
func (s *BaseNeedleParserListener) EnterExprList(ctx *ExprListContext) {}

// ExitExprList is called when production exprList is exited.
func (s *BaseNeedleParserListener) ExitExprList(ctx *ExprListContext) {}

// EnterExpr is called when production expr is entered.
func (s *BaseNeedleParserListener) EnterExpr(ctx *ExprContext) {}

// ExitExpr is called when production expr is exited.
func (s *BaseNeedleParserListener) ExitExpr(ctx *ExprContext) {}

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
