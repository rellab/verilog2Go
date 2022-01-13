// Code generated from Expression.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // Expression

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 42, 100,
	4, 2, 9, 2, 4, 3, 9, 3, 3, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 5, 3, 32, 10, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 7, 3, 95, 10, 3,
	12, 3, 14, 3, 98, 11, 3, 3, 3, 2, 3, 4, 4, 2, 4, 2, 11, 3, 2, 3, 8, 3,
	2, 10, 12, 3, 2, 4, 5, 3, 2, 13, 16, 3, 2, 17, 20, 3, 2, 21, 24, 4, 2,
	8, 8, 25, 25, 3, 2, 26, 27, 4, 2, 7, 7, 28, 28, 2, 114, 2, 6, 3, 2, 2,
	2, 4, 31, 3, 2, 2, 2, 6, 7, 5, 4, 3, 2, 7, 8, 7, 2, 2, 3, 8, 3, 3, 2, 2,
	2, 9, 10, 8, 3, 1, 2, 10, 11, 9, 2, 2, 2, 11, 12, 5, 4, 3, 20, 12, 13,
	8, 3, 1, 2, 13, 32, 3, 2, 2, 2, 14, 15, 7, 35, 2, 2, 15, 32, 8, 3, 1, 2,
	16, 17, 7, 36, 2, 2, 17, 32, 8, 3, 1, 2, 18, 19, 7, 31, 2, 2, 19, 20, 5,
	4, 3, 2, 20, 21, 7, 32, 2, 2, 21, 22, 8, 3, 1, 2, 22, 32, 3, 2, 2, 2, 23,
	24, 7, 36, 2, 2, 24, 25, 7, 33, 2, 2, 25, 26, 5, 4, 3, 2, 26, 27, 7, 34,
	2, 2, 27, 28, 8, 3, 1, 2, 28, 32, 3, 2, 2, 2, 29, 30, 7, 41, 2, 2, 30,
	32, 8, 3, 1, 2, 31, 9, 3, 2, 2, 2, 31, 14, 3, 2, 2, 2, 31, 16, 3, 2, 2,
	2, 31, 18, 3, 2, 2, 2, 31, 23, 3, 2, 2, 2, 31, 29, 3, 2, 2, 2, 32, 96,
	3, 2, 2, 2, 33, 34, 12, 19, 2, 2, 34, 35, 7, 9, 2, 2, 35, 36, 5, 4, 3,
	20, 36, 37, 8, 3, 1, 2, 37, 95, 3, 2, 2, 2, 38, 39, 12, 18, 2, 2, 39, 40,
	9, 3, 2, 2, 40, 41, 5, 4, 3, 19, 41, 42, 8, 3, 1, 2, 42, 95, 3, 2, 2, 2,
	43, 44, 12, 17, 2, 2, 44, 45, 9, 4, 2, 2, 45, 46, 5, 4, 3, 18, 46, 47,
	8, 3, 1, 2, 47, 95, 3, 2, 2, 2, 48, 49, 12, 16, 2, 2, 49, 50, 9, 5, 2,
	2, 50, 51, 5, 4, 3, 17, 51, 52, 8, 3, 1, 2, 52, 95, 3, 2, 2, 2, 53, 54,
	12, 15, 2, 2, 54, 55, 9, 6, 2, 2, 55, 56, 5, 4, 3, 16, 56, 57, 8, 3, 1,
	2, 57, 95, 3, 2, 2, 2, 58, 59, 12, 14, 2, 2, 59, 60, 9, 7, 2, 2, 60, 61,
	5, 4, 3, 15, 61, 62, 8, 3, 1, 2, 62, 95, 3, 2, 2, 2, 63, 64, 12, 13, 2,
	2, 64, 65, 9, 8, 2, 2, 65, 66, 5, 4, 3, 14, 66, 67, 8, 3, 1, 2, 67, 95,
	3, 2, 2, 2, 68, 69, 12, 12, 2, 2, 69, 70, 9, 9, 2, 2, 70, 71, 5, 4, 3,
	13, 71, 72, 8, 3, 1, 2, 72, 95, 3, 2, 2, 2, 73, 74, 12, 11, 2, 2, 74, 75,
	9, 10, 2, 2, 75, 76, 5, 4, 3, 12, 76, 77, 8, 3, 1, 2, 77, 95, 3, 2, 2,
	2, 78, 79, 12, 10, 2, 2, 79, 80, 7, 29, 2, 2, 80, 81, 5, 4, 3, 11, 81,
	82, 8, 3, 1, 2, 82, 95, 3, 2, 2, 2, 83, 84, 12, 9, 2, 2, 84, 85, 7, 30,
	2, 2, 85, 86, 5, 4, 3, 10, 86, 87, 8, 3, 1, 2, 87, 95, 3, 2, 2, 2, 88,
	89, 12, 5, 2, 2, 89, 90, 7, 31, 2, 2, 90, 91, 5, 4, 3, 2, 91, 92, 7, 32,
	2, 2, 92, 93, 8, 3, 1, 2, 93, 95, 3, 2, 2, 2, 94, 33, 3, 2, 2, 2, 94, 38,
	3, 2, 2, 2, 94, 43, 3, 2, 2, 2, 94, 48, 3, 2, 2, 2, 94, 53, 3, 2, 2, 2,
	94, 58, 3, 2, 2, 2, 94, 63, 3, 2, 2, 2, 94, 68, 3, 2, 2, 2, 94, 73, 3,
	2, 2, 2, 94, 78, 3, 2, 2, 2, 94, 83, 3, 2, 2, 2, 94, 88, 3, 2, 2, 2, 95,
	98, 3, 2, 2, 2, 96, 94, 3, 2, 2, 2, 96, 97, 3, 2, 2, 2, 97, 5, 3, 2, 2,
	2, 98, 96, 3, 2, 2, 2, 5, 31, 94, 96,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'!'", "'+'", "'-'", "'~'", "'|'", "'&'", "'**'", "'*'", "'/'", "'%'",
	"'<<'", "'>>'", "'<<<'", "'>>>'", "'<'", "'<='", "'>'", "'>='", "'=='",
	"'!='", "'==='", "'!=='", "'~&'", "'^'", "'~^'", "'~|'", "'&&'", "'||'",
	"'('", "')'", "'['", "']'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "Number", "ID",
	"Decimal_number", "Binary_number", "Octal_number", "Hex_number", "Decimal",
	"White_space",
}

var ruleNames = []string{
	"start", "expression",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type ExpressionParser struct {
	*antlr.BaseParser
}

func NewExpressionParser(input antlr.TokenStream) *ExpressionParser {
	this := new(ExpressionParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "Expression.g4"

	return this
}

// ExpressionParser tokens.
const (
	ExpressionParserEOF            = antlr.TokenEOF
	ExpressionParserT__0           = 1
	ExpressionParserT__1           = 2
	ExpressionParserT__2           = 3
	ExpressionParserT__3           = 4
	ExpressionParserT__4           = 5
	ExpressionParserT__5           = 6
	ExpressionParserT__6           = 7
	ExpressionParserT__7           = 8
	ExpressionParserT__8           = 9
	ExpressionParserT__9           = 10
	ExpressionParserT__10          = 11
	ExpressionParserT__11          = 12
	ExpressionParserT__12          = 13
	ExpressionParserT__13          = 14
	ExpressionParserT__14          = 15
	ExpressionParserT__15          = 16
	ExpressionParserT__16          = 17
	ExpressionParserT__17          = 18
	ExpressionParserT__18          = 19
	ExpressionParserT__19          = 20
	ExpressionParserT__20          = 21
	ExpressionParserT__21          = 22
	ExpressionParserT__22          = 23
	ExpressionParserT__23          = 24
	ExpressionParserT__24          = 25
	ExpressionParserT__25          = 26
	ExpressionParserT__26          = 27
	ExpressionParserT__27          = 28
	ExpressionParserT__28          = 29
	ExpressionParserT__29          = 30
	ExpressionParserT__30          = 31
	ExpressionParserT__31          = 32
	ExpressionParserNumber         = 33
	ExpressionParserID             = 34
	ExpressionParserDecimal_number = 35
	ExpressionParserBinary_number  = 36
	ExpressionParserOctal_number   = 37
	ExpressionParserHex_number     = 38
	ExpressionParserDecimal        = 39
	ExpressionParserWhite_space    = 40
)

// ExpressionParser rules.
const (
	ExpressionParserRULE_start      = 0
	ExpressionParserRULE_expression = 1
)

// IStartContext is an interface to support dynamic dispatch.
type IStartContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStartContext differentiates from other interfaces.
	IsStartContext()
}

type StartContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStartContext() *StartContext {
	var p = new(StartContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ExpressionParserRULE_start
	return p
}

func (*StartContext) IsStartContext() {}

func NewStartContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StartContext {
	var p = new(StartContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ExpressionParserRULE_start

	return p
}

func (s *StartContext) GetParser() antlr.Parser { return s.parser }

func (s *StartContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *StartContext) EOF() antlr.TerminalNode {
	return s.GetToken(ExpressionParserEOF, 0)
}

func (s *StartContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StartContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StartContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExpressionListener); ok {
		listenerT.EnterStart(s)
	}
}

func (s *StartContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExpressionListener); ok {
		listenerT.ExitStart(s)
	}
}

func (p *ExpressionParser) Start() (localctx IStartContext) {
	localctx = NewStartContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, ExpressionParserRULE_start)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(4)
		p.expression(0)
	}
	{
		p.SetState(5)
		p.Match(ExpressionParserEOF)
	}

	return localctx
}

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOp returns the op token.
	GetOp() antlr.Token

	// GetId returns the id token.
	GetId() antlr.Token

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// SetId sets the id token.
	SetId(antlr.Token)

	// GetNumber returns the number attribute.
	GetNumber() int

	// SetNumber sets the number attribute.
	SetNumber(int)

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	number int
	op     antlr.Token
	id     antlr.Token
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ExpressionParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ExpressionParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) GetOp() antlr.Token { return s.op }

func (s *ExpressionContext) GetId() antlr.Token { return s.id }

func (s *ExpressionContext) SetOp(v antlr.Token) { s.op = v }

func (s *ExpressionContext) SetId(v antlr.Token) { s.id = v }

func (s *ExpressionContext) GetNumber() int { return s.number }

func (s *ExpressionContext) SetNumber(v int) { s.number = v }

func (s *ExpressionContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *ExpressionContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ExpressionContext) Number() antlr.TerminalNode {
	return s.GetToken(ExpressionParserNumber, 0)
}

func (s *ExpressionContext) ID() antlr.TerminalNode {
	return s.GetToken(ExpressionParserID, 0)
}

func (s *ExpressionContext) Decimal() antlr.TerminalNode {
	return s.GetToken(ExpressionParserDecimal, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExpressionListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExpressionListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *ExpressionParser) Expression() (localctx IExpressionContext) {
	return p.expression(0)
}

func (p *ExpressionParser) expression(_p int) (localctx IExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 2
	p.EnterRecursionRule(localctx, 2, ExpressionParserRULE_expression, _p)
	var _la int

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(29)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(8)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*ExpressionContext).op = _lt

			_la = p.GetTokenStream().LA(1)

			if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ExpressionParserT__0)|(1<<ExpressionParserT__1)|(1<<ExpressionParserT__2)|(1<<ExpressionParserT__3)|(1<<ExpressionParserT__4)|(1<<ExpressionParserT__5))) != 0) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*ExpressionContext).op = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(9)
			p.expression(18)
		}
		localctx.(*ExpressionContext).SetNumber(1)

	case 2:
		{
			p.SetState(12)

			var _m = p.Match(ExpressionParserNumber)

			localctx.(*ExpressionContext).id = _m
		}
		localctx.(*ExpressionContext).SetNumber(17)

	case 3:
		{
			p.SetState(14)

			var _m = p.Match(ExpressionParserID)

			localctx.(*ExpressionContext).id = _m
		}
		localctx.(*ExpressionContext).SetNumber(13)

	case 4:
		{
			p.SetState(16)
			p.Match(ExpressionParserT__28)
		}
		{
			p.SetState(17)
			p.expression(0)
		}
		{
			p.SetState(18)
			p.Match(ExpressionParserT__29)
		}
		localctx.(*ExpressionContext).SetNumber(14)

	case 5:
		{
			p.SetState(21)

			var _m = p.Match(ExpressionParserID)

			localctx.(*ExpressionContext).id = _m
		}
		{
			p.SetState(22)
			p.Match(ExpressionParserT__30)
		}
		{
			p.SetState(23)
			p.expression(0)
		}
		{
			p.SetState(24)
			p.Match(ExpressionParserT__31)
		}
		localctx.(*ExpressionContext).SetNumber(16)

	case 6:
		{
			p.SetState(27)
			p.Match(ExpressionParserDecimal)
		}
		localctx.(*ExpressionContext).SetNumber(18)

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(94)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(92)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, ExpressionParserRULE_expression)
				p.SetState(31)

				if !(p.Precpred(p.GetParserRuleContext(), 17)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 17)", ""))
				}
				{
					p.SetState(32)

					var _m = p.Match(ExpressionParserT__6)

					localctx.(*ExpressionContext).op = _m
				}
				{
					p.SetState(33)
					p.expression(18)
				}
				localctx.(*ExpressionContext).SetNumber(2)

			case 2:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, ExpressionParserRULE_expression)
				p.SetState(36)

				if !(p.Precpred(p.GetParserRuleContext(), 16)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 16)", ""))
				}
				{
					p.SetState(37)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExpressionContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ExpressionParserT__7)|(1<<ExpressionParserT__8)|(1<<ExpressionParserT__9))) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExpressionContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(38)
					p.expression(17)
				}
				localctx.(*ExpressionContext).SetNumber(3)

			case 3:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, ExpressionParserRULE_expression)
				p.SetState(41)

				if !(p.Precpred(p.GetParserRuleContext(), 15)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 15)", ""))
				}
				{
					p.SetState(42)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExpressionContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == ExpressionParserT__1 || _la == ExpressionParserT__2) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExpressionContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(43)
					p.expression(16)
				}
				localctx.(*ExpressionContext).SetNumber(4)

			case 4:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, ExpressionParserRULE_expression)
				p.SetState(46)

				if !(p.Precpred(p.GetParserRuleContext(), 14)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 14)", ""))
				}
				{
					p.SetState(47)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExpressionContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ExpressionParserT__10)|(1<<ExpressionParserT__11)|(1<<ExpressionParserT__12)|(1<<ExpressionParserT__13))) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExpressionContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(48)
					p.expression(15)
				}
				localctx.(*ExpressionContext).SetNumber(5)

			case 5:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, ExpressionParserRULE_expression)
				p.SetState(51)

				if !(p.Precpred(p.GetParserRuleContext(), 13)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 13)", ""))
				}
				{
					p.SetState(52)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExpressionContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ExpressionParserT__14)|(1<<ExpressionParserT__15)|(1<<ExpressionParserT__16)|(1<<ExpressionParserT__17))) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExpressionContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(53)
					p.expression(14)
				}
				localctx.(*ExpressionContext).SetNumber(6)

			case 6:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, ExpressionParserRULE_expression)
				p.SetState(56)

				if !(p.Precpred(p.GetParserRuleContext(), 12)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 12)", ""))
				}
				{
					p.SetState(57)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExpressionContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ExpressionParserT__18)|(1<<ExpressionParserT__19)|(1<<ExpressionParserT__20)|(1<<ExpressionParserT__21))) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExpressionContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(58)
					p.expression(13)
				}
				localctx.(*ExpressionContext).SetNumber(7)

			case 7:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, ExpressionParserRULE_expression)
				p.SetState(61)

				if !(p.Precpred(p.GetParserRuleContext(), 11)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 11)", ""))
				}
				{
					p.SetState(62)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExpressionContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == ExpressionParserT__5 || _la == ExpressionParserT__22) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExpressionContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(63)
					p.expression(12)
				}
				localctx.(*ExpressionContext).SetNumber(8)

			case 8:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, ExpressionParserRULE_expression)
				p.SetState(66)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
				}
				{
					p.SetState(67)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExpressionContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == ExpressionParserT__23 || _la == ExpressionParserT__24) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExpressionContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(68)
					p.expression(11)
				}
				localctx.(*ExpressionContext).SetNumber(9)

			case 9:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, ExpressionParserRULE_expression)
				p.SetState(71)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
				}
				{
					p.SetState(72)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExpressionContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == ExpressionParserT__4 || _la == ExpressionParserT__25) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExpressionContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(73)
					p.expression(10)
				}
				localctx.(*ExpressionContext).SetNumber(10)

			case 10:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, ExpressionParserRULE_expression)
				p.SetState(76)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
				}
				{
					p.SetState(77)

					var _m = p.Match(ExpressionParserT__26)

					localctx.(*ExpressionContext).op = _m
				}
				{
					p.SetState(78)
					p.expression(9)
				}
				localctx.(*ExpressionContext).SetNumber(11)

			case 11:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, ExpressionParserRULE_expression)
				p.SetState(81)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
				}
				{
					p.SetState(82)

					var _m = p.Match(ExpressionParserT__27)

					localctx.(*ExpressionContext).op = _m
				}
				{
					p.SetState(83)
					p.expression(8)
				}
				localctx.(*ExpressionContext).SetNumber(12)

			case 12:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, ExpressionParserRULE_expression)
				p.SetState(86)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(87)
					p.Match(ExpressionParserT__28)
				}
				{
					p.SetState(88)
					p.expression(0)
				}
				{
					p.SetState(89)
					p.Match(ExpressionParserT__29)
				}
				localctx.(*ExpressionContext).SetNumber(15)

			}

		}
		p.SetState(96)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())
	}

	return localctx
}

func (p *ExpressionParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 1:
		var t *ExpressionContext = nil
		if localctx != nil {
			t = localctx.(*ExpressionContext)
		}
		return p.Expression_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *ExpressionParser) Expression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 17)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 16)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 15)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 14)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 13)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 12)

	case 6:
		return p.Precpred(p.GetParserRuleContext(), 11)

	case 7:
		return p.Precpred(p.GetParserRuleContext(), 10)

	case 8:
		return p.Precpred(p.GetParserRuleContext(), 9)

	case 9:
		return p.Precpred(p.GetParserRuleContext(), 8)

	case 10:
		return p.Precpred(p.GetParserRuleContext(), 7)

	case 11:
		return p.Precpred(p.GetParserRuleContext(), 3)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
