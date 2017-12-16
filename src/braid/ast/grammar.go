package ast

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

func toIfaceSlice(v interface{}) []interface{} {
	if v == nil {
		return nil
	}
	return v.([]interface{})
}

var g = &grammar{
	rules: []*rule{
		{
			name: "Module",
			pos:  position{line: 13, col: 1, offset: 142},
			expr: &actionExpr{
				pos: position{line: 13, col: 10, offset: 151},
				run: (*parser).callonModule1,
				expr: &seqExpr{
					pos: position{line: 13, col: 10, offset: 151},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 13, col: 10, offset: 151},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 13, col: 12, offset: 153},
							val:        "module",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 13, col: 21, offset: 162},
							name: "__",
						},
						&labeledExpr{
							pos:   position{line: 13, col: 24, offset: 165},
							label: "name",
							expr: &ruleRefExpr{
								pos:  position{line: 13, col: 29, offset: 170},
								name: "ModuleName",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 13, col: 40, offset: 181},
							name: "__",
						},
						&labeledExpr{
							pos:   position{line: 13, col: 43, offset: 184},
							label: "stat",
							expr: &ruleRefExpr{
								pos:  position{line: 13, col: 48, offset: 189},
								name: "TopLevelStatement",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 13, col: 66, offset: 207},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 13, col: 68, offset: 209},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 13, col: 73, offset: 214},
								expr: &ruleRefExpr{
									pos:  position{line: 13, col: 74, offset: 215},
									name: "TopLevelStatement",
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 13, col: 94, offset: 235},
							name: "_",
						},
						&ruleRefExpr{
							pos:  position{line: 13, col: 96, offset: 237},
							name: "EOF",
						},
					},
				},
			},
		},
		{
			name: "TopLevelStatement",
			pos:  position{line: 28, col: 1, offset: 718},
			expr: &choiceExpr{
				pos: position{line: 28, col: 21, offset: 738},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 28, col: 21, offset: 738},
						name: "FuncDefn",
					},
					&ruleRefExpr{
						pos:  position{line: 28, col: 32, offset: 749},
						name: "TypeDefn",
					},
					&ruleRefExpr{
						pos:  position{line: 28, col: 43, offset: 760},
						name: "ExternDefn",
					},
				},
			},
		},
		{
			name: "Statement",
			pos:  position{line: 30, col: 1, offset: 772},
			expr: &choiceExpr{
				pos: position{line: 30, col: 13, offset: 784},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 30, col: 13, offset: 784},
						name: "FuncDefn",
					},
					&ruleRefExpr{
						pos:  position{line: 30, col: 24, offset: 795},
						name: "Assignment",
					},
					&ruleRefExpr{
						pos:  position{line: 30, col: 37, offset: 808},
						name: "ExprLine",
					},
				},
			},
		},
		{
			name: "ExprLine",
			pos:  position{line: 32, col: 1, offset: 818},
			expr: &actionExpr{
				pos: position{line: 32, col: 12, offset: 829},
				run: (*parser).callonExprLine1,
				expr: &seqExpr{
					pos: position{line: 32, col: 12, offset: 829},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 32, col: 12, offset: 829},
							label: "e",
							expr: &ruleRefExpr{
								pos:  position{line: 32, col: 14, offset: 831},
								name: "Expr",
							},
						},
						&andExpr{
							pos: position{line: 32, col: 19, offset: 836},
							expr: &litMatcher{
								pos:        position{line: 32, col: 20, offset: 837},
								val:        "\n",
								ignoreCase: false,
							},
						},
					},
				},
			},
		},
		{
			name: "ExternDefn",
			pos:  position{line: 46, col: 1, offset: 1198},
			expr: &actionExpr{
				pos: position{line: 46, col: 14, offset: 1211},
				run: (*parser).callonExternDefn1,
				expr: &seqExpr{
					pos: position{line: 46, col: 14, offset: 1211},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 46, col: 14, offset: 1211},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 46, col: 16, offset: 1213},
							val:        "extern",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 46, col: 25, offset: 1222},
							name: "__",
						},
						&labeledExpr{
							pos:   position{line: 46, col: 28, offset: 1225},
							label: "name",
							expr: &ruleRefExpr{
								pos:  position{line: 46, col: 33, offset: 1230},
								name: "Identifier",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 46, col: 44, offset: 1241},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 46, col: 46, offset: 1243},
							val:        "=",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 46, col: 50, offset: 1247},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 46, col: 52, offset: 1249},
							label: "importName",
							expr: &ruleRefExpr{
								pos:  position{line: 46, col: 63, offset: 1260},
								name: "StringLiteral",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 46, col: 77, offset: 1274},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 46, col: 79, offset: 1276},
							label: "args",
							expr: &ruleRefExpr{
								pos:  position{line: 46, col: 84, offset: 1281},
								name: "ArgsDefn",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 46, col: 93, offset: 1290},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 46, col: 95, offset: 1292},
							label: "ret",
							expr: &ruleRefExpr{
								pos:  position{line: 46, col: 99, offset: 1296},
								name: "BaseType",
							},
						},
					},
				},
			},
		},
		{
			name: "TypeDefn",
			pos:  position{line: 52, col: 1, offset: 1508},
			expr: &choiceExpr{
				pos: position{line: 52, col: 12, offset: 1519},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 52, col: 12, offset: 1519},
						run: (*parser).callonTypeDefn2,
						expr: &seqExpr{
							pos: position{line: 52, col: 12, offset: 1519},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 52, col: 12, offset: 1519},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 52, col: 14, offset: 1521},
									val:        "type",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 52, col: 21, offset: 1528},
									name: "__",
								},
								&labeledExpr{
									pos:   position{line: 52, col: 24, offset: 1531},
									label: "name",
									expr: &ruleRefExpr{
										pos:  position{line: 52, col: 29, offset: 1536},
										name: "ModuleName",
									},
								},
								&labeledExpr{
									pos:   position{line: 52, col: 40, offset: 1547},
									label: "params",
									expr: &zeroOrMoreExpr{
										pos: position{line: 52, col: 47, offset: 1554},
										expr: &seqExpr{
											pos: position{line: 52, col: 48, offset: 1555},
											exprs: []interface{}{
												&ruleRefExpr{
													pos:  position{line: 52, col: 48, offset: 1555},
													name: "__",
												},
												&ruleRefExpr{
													pos:  position{line: 52, col: 51, offset: 1558},
													name: "TypeParameter",
												},
											},
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 52, col: 67, offset: 1574},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 52, col: 69, offset: 1576},
									val:        "=",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 52, col: 73, offset: 1580},
									label: "types",
									expr: &oneOrMoreExpr{
										pos: position{line: 52, col: 79, offset: 1586},
										expr: &seqExpr{
											pos: position{line: 52, col: 80, offset: 1587},
											exprs: []interface{}{
												&ruleRefExpr{
													pos:  position{line: 52, col: 80, offset: 1587},
													name: "__",
												},
												&ruleRefExpr{
													pos:  position{line: 52, col: 83, offset: 1590},
													name: "AnyType",
												},
											},
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 52, col: 93, offset: 1600},
									name: "__",
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 71, col: 1, offset: 2096},
						run: (*parser).callonTypeDefn22,
						expr: &seqExpr{
							pos: position{line: 71, col: 1, offset: 2096},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 71, col: 1, offset: 2096},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 71, col: 3, offset: 2098},
									val:        "type",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 71, col: 10, offset: 2105},
									name: "__",
								},
								&labeledExpr{
									pos:   position{line: 71, col: 13, offset: 2108},
									label: "name",
									expr: &ruleRefExpr{
										pos:  position{line: 71, col: 18, offset: 2113},
										name: "ModuleName",
									},
								},
								&labeledExpr{
									pos:   position{line: 71, col: 29, offset: 2124},
									label: "params",
									expr: &zeroOrMoreExpr{
										pos: position{line: 71, col: 36, offset: 2131},
										expr: &seqExpr{
											pos: position{line: 71, col: 37, offset: 2132},
											exprs: []interface{}{
												&ruleRefExpr{
													pos:  position{line: 71, col: 37, offset: 2132},
													name: "__",
												},
												&ruleRefExpr{
													pos:  position{line: 71, col: 40, offset: 2135},
													name: "TypeParameter",
												},
											},
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 71, col: 56, offset: 2151},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 71, col: 58, offset: 2153},
									val:        "=",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 71, col: 62, offset: 2157},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 72, col: 5, offset: 2163},
									val:        "{",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 72, col: 9, offset: 2167},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 72, col: 11, offset: 2169},
									label: "first",
									expr: &ruleRefExpr{
										pos:  position{line: 72, col: 17, offset: 2175},
										name: "RecordFieldDefn",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 72, col: 33, offset: 2191},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 72, col: 35, offset: 2193},
									label: "rest",
									expr: &zeroOrMoreExpr{
										pos: position{line: 72, col: 40, offset: 2198},
										expr: &seqExpr{
											pos: position{line: 72, col: 41, offset: 2199},
											exprs: []interface{}{
												&litMatcher{
													pos:        position{line: 72, col: 41, offset: 2199},
													val:        ",",
													ignoreCase: false,
												},
												&ruleRefExpr{
													pos:  position{line: 72, col: 45, offset: 2203},
													name: "_",
												},
												&ruleRefExpr{
													pos:  position{line: 72, col: 47, offset: 2205},
													name: "RecordFieldDefn",
												},
												&ruleRefExpr{
													pos:  position{line: 72, col: 63, offset: 2221},
													name: "_",
												},
											},
										},
									},
								},
								&zeroOrOneExpr{
									pos: position{line: 72, col: 67, offset: 2225},
									expr: &litMatcher{
										pos:        position{line: 72, col: 67, offset: 2225},
										val:        ",",
										ignoreCase: false,
									},
								},
								&ruleRefExpr{
									pos:  position{line: 72, col: 72, offset: 2230},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 72, col: 74, offset: 2232},
									val:        "}",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 72, col: 78, offset: 2236},
									name: "__",
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 90, col: 1, offset: 2721},
						run: (*parser).callonTypeDefn54,
						expr: &seqExpr{
							pos: position{line: 90, col: 1, offset: 2721},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 90, col: 1, offset: 2721},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 90, col: 3, offset: 2723},
									val:        "type",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 90, col: 10, offset: 2730},
									name: "__",
								},
								&labeledExpr{
									pos:   position{line: 90, col: 13, offset: 2733},
									label: "name",
									expr: &ruleRefExpr{
										pos:  position{line: 90, col: 18, offset: 2738},
										name: "ModuleName",
									},
								},
								&labeledExpr{
									pos:   position{line: 90, col: 29, offset: 2749},
									label: "params",
									expr: &zeroOrMoreExpr{
										pos: position{line: 90, col: 36, offset: 2756},
										expr: &seqExpr{
											pos: position{line: 90, col: 37, offset: 2757},
											exprs: []interface{}{
												&ruleRefExpr{
													pos:  position{line: 90, col: 37, offset: 2757},
													name: "__",
												},
												&ruleRefExpr{
													pos:  position{line: 90, col: 40, offset: 2760},
													name: "TypeParameter",
												},
											},
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 90, col: 56, offset: 2776},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 90, col: 58, offset: 2778},
									val:        "=",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 90, col: 62, offset: 2782},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 90, col: 64, offset: 2784},
									label: "rest",
									expr: &oneOrMoreExpr{
										pos: position{line: 90, col: 69, offset: 2789},
										expr: &ruleRefExpr{
											pos:  position{line: 90, col: 70, offset: 2790},
											name: "VariantConstructor",
										},
									},
								},
							},
						},
					},
				},
			},
		},
		{
			name: "RecordFieldDefn",
			pos:  position{line: 105, col: 1, offset: 3197},
			expr: &actionExpr{
				pos: position{line: 105, col: 19, offset: 3215},
				run: (*parser).callonRecordFieldDefn1,
				expr: &seqExpr{
					pos: position{line: 105, col: 19, offset: 3215},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 105, col: 19, offset: 3215},
							label: "name",
							expr: &ruleRefExpr{
								pos:  position{line: 105, col: 24, offset: 3220},
								name: "VariableName",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 105, col: 37, offset: 3233},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 105, col: 39, offset: 3235},
							val:        ":",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 105, col: 43, offset: 3239},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 105, col: 45, offset: 3241},
							label: "t",
							expr: &ruleRefExpr{
								pos:  position{line: 105, col: 48, offset: 3244},
								name: "AnyType",
							},
						},
					},
				},
			},
		},
		{
			name: "VariantConstructor",
			pos:  position{line: 109, col: 1, offset: 3338},
			expr: &choiceExpr{
				pos: position{line: 109, col: 22, offset: 3359},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 109, col: 22, offset: 3359},
						run: (*parser).callonVariantConstructor2,
						expr: &seqExpr{
							pos: position{line: 109, col: 22, offset: 3359},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 109, col: 22, offset: 3359},
									val:        "|",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 109, col: 26, offset: 3363},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 109, col: 28, offset: 3365},
									label: "name",
									expr: &ruleRefExpr{
										pos:  position{line: 109, col: 33, offset: 3370},
										name: "ModuleName",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 109, col: 44, offset: 3381},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 109, col: 46, offset: 3383},
									val:        "{",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 109, col: 50, offset: 3387},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 109, col: 52, offset: 3389},
									label: "first",
									expr: &ruleRefExpr{
										pos:  position{line: 109, col: 58, offset: 3395},
										name: "RecordFieldDefn",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 109, col: 74, offset: 3411},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 109, col: 76, offset: 3413},
									label: "rest",
									expr: &zeroOrMoreExpr{
										pos: position{line: 109, col: 81, offset: 3418},
										expr: &seqExpr{
											pos: position{line: 109, col: 82, offset: 3419},
											exprs: []interface{}{
												&litMatcher{
													pos:        position{line: 109, col: 82, offset: 3419},
													val:        ",",
													ignoreCase: false,
												},
												&ruleRefExpr{
													pos:  position{line: 109, col: 86, offset: 3423},
													name: "_",
												},
												&ruleRefExpr{
													pos:  position{line: 109, col: 88, offset: 3425},
													name: "RecordFieldDefn",
												},
												&ruleRefExpr{
													pos:  position{line: 109, col: 104, offset: 3441},
													name: "_",
												},
											},
										},
									},
								},
								&zeroOrOneExpr{
									pos: position{line: 109, col: 108, offset: 3445},
									expr: &litMatcher{
										pos:        position{line: 109, col: 108, offset: 3445},
										val:        ",",
										ignoreCase: false,
									},
								},
								&ruleRefExpr{
									pos:  position{line: 109, col: 113, offset: 3450},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 109, col: 115, offset: 3452},
									val:        "}",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 109, col: 119, offset: 3456},
									name: "__",
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 128, col: 1, offset: 4061},
						run: (*parser).callonVariantConstructor26,
						expr: &seqExpr{
							pos: position{line: 128, col: 1, offset: 4061},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 128, col: 1, offset: 4061},
									val:        "|",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 128, col: 5, offset: 4065},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 128, col: 7, offset: 4067},
									label: "name",
									expr: &ruleRefExpr{
										pos:  position{line: 128, col: 12, offset: 4072},
										name: "ModuleName",
									},
								},
								&labeledExpr{
									pos:   position{line: 128, col: 23, offset: 4083},
									label: "rest",
									expr: &zeroOrMoreExpr{
										pos: position{line: 128, col: 28, offset: 4088},
										expr: &seqExpr{
											pos: position{line: 128, col: 29, offset: 4089},
											exprs: []interface{}{
												&ruleRefExpr{
													pos:  position{line: 128, col: 29, offset: 4089},
													name: "__",
												},
												&ruleRefExpr{
													pos:  position{line: 128, col: 32, offset: 4092},
													name: "AnyType",
												},
											},
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 128, col: 42, offset: 4102},
									name: "_",
								},
							},
						},
					},
				},
			},
		},
		{
			name: "AnyType",
			pos:  position{line: 145, col: 1, offset: 4539},
			expr: &choiceExpr{
				pos: position{line: 145, col: 11, offset: 4549},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 145, col: 11, offset: 4549},
						name: "BaseType",
					},
					&ruleRefExpr{
						pos:  position{line: 145, col: 22, offset: 4560},
						name: "TypeParameter",
					},
				},
			},
		},
		{
			name: "Assignment",
			pos:  position{line: 147, col: 1, offset: 4575},
			expr: &choiceExpr{
				pos: position{line: 147, col: 14, offset: 4588},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 147, col: 14, offset: 4588},
						run: (*parser).callonAssignment2,
						expr: &seqExpr{
							pos: position{line: 147, col: 14, offset: 4588},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 147, col: 14, offset: 4588},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 147, col: 16, offset: 4590},
									val:        "let",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 147, col: 22, offset: 4596},
									name: "__",
								},
								&labeledExpr{
									pos:   position{line: 147, col: 25, offset: 4599},
									label: "i",
									expr: &ruleRefExpr{
										pos:  position{line: 147, col: 27, offset: 4601},
										name: "Assignable",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 147, col: 38, offset: 4612},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 147, col: 40, offset: 4614},
									val:        "=",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 147, col: 44, offset: 4618},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 147, col: 46, offset: 4620},
									label: "expr",
									expr: &ruleRefExpr{
										pos:  position{line: 147, col: 51, offset: 4625},
										name: "Expr",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 147, col: 56, offset: 4630},
									name: "_",
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 153, col: 1, offset: 4749},
						run: (*parser).callonAssignment15,
						expr: &seqExpr{
							pos: position{line: 153, col: 1, offset: 4749},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 153, col: 1, offset: 4749},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 153, col: 3, offset: 4751},
									val:        "let",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 153, col: 9, offset: 4757},
									name: "__",
								},
								&notExpr{
									pos: position{line: 153, col: 12, offset: 4760},
									expr: &ruleRefExpr{
										pos:  position{line: 153, col: 13, offset: 4761},
										name: "Assignable",
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 157, col: 1, offset: 4859},
						run: (*parser).callonAssignment22,
						expr: &seqExpr{
							pos: position{line: 157, col: 1, offset: 4859},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 157, col: 1, offset: 4859},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 157, col: 3, offset: 4861},
									val:        "let",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 157, col: 9, offset: 4867},
									name: "__",
								},
								&labeledExpr{
									pos:   position{line: 157, col: 12, offset: 4870},
									label: "i",
									expr: &ruleRefExpr{
										pos:  position{line: 157, col: 14, offset: 4872},
										name: "Assignable",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 157, col: 25, offset: 4883},
									name: "_",
								},
								&notExpr{
									pos: position{line: 157, col: 27, offset: 4885},
									expr: &litMatcher{
										pos:        position{line: 157, col: 28, offset: 4886},
										val:        "=",
										ignoreCase: false,
									},
								},
							},
						},
					},
				},
			},
		},
		{
			name: "FuncDefn",
			pos:  position{line: 161, col: 1, offset: 4980},
			expr: &actionExpr{
				pos: position{line: 161, col: 12, offset: 4991},
				run: (*parser).callonFuncDefn1,
				expr: &seqExpr{
					pos: position{line: 161, col: 12, offset: 4991},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 161, col: 12, offset: 4991},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 161, col: 14, offset: 4993},
							val:        "let",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 161, col: 20, offset: 4999},
							name: "__",
						},
						&labeledExpr{
							pos:   position{line: 161, col: 23, offset: 5002},
							label: "i",
							expr: &ruleRefExpr{
								pos:  position{line: 161, col: 25, offset: 5004},
								name: "VariableName",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 161, col: 38, offset: 5017},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 161, col: 40, offset: 5019},
							val:        "=",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 161, col: 44, offset: 5023},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 161, col: 46, offset: 5025},
							label: "ids",
							expr: &zeroOrOneExpr{
								pos: position{line: 161, col: 50, offset: 5029},
								expr: &seqExpr{
									pos: position{line: 161, col: 51, offset: 5030},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 161, col: 51, offset: 5030},
											name: "ArgsDefn",
										},
										&ruleRefExpr{
											pos:  position{line: 161, col: 60, offset: 5039},
											name: "_",
										},
									},
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 161, col: 64, offset: 5043},
							label: "ret",
							expr: &zeroOrOneExpr{
								pos: position{line: 161, col: 68, offset: 5047},
								expr: &seqExpr{
									pos: position{line: 161, col: 69, offset: 5048},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 161, col: 69, offset: 5048},
											name: "AnyType",
										},
										&ruleRefExpr{
											pos:  position{line: 161, col: 77, offset: 5056},
											name: "_",
										},
									},
								},
							},
						},
						&litMatcher{
							pos:        position{line: 161, col: 81, offset: 5060},
							val:        "{",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 161, col: 85, offset: 5064},
							name: "__",
						},
						&labeledExpr{
							pos:   position{line: 161, col: 88, offset: 5067},
							label: "statements",
							expr: &oneOrMoreExpr{
								pos: position{line: 161, col: 99, offset: 5078},
								expr: &ruleRefExpr{
									pos:  position{line: 161, col: 100, offset: 5079},
									name: "Statement",
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 161, col: 112, offset: 5091},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 161, col: 114, offset: 5093},
							val:        "}",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 161, col: 118, offset: 5097},
							name: "_",
						},
					},
				},
			},
		},
		{
			name: "Expr",
			pos:  position{line: 187, col: 1, offset: 5707},
			expr: &actionExpr{
				pos: position{line: 187, col: 8, offset: 5714},
				run: (*parser).callonExpr1,
				expr: &labeledExpr{
					pos:   position{line: 187, col: 8, offset: 5714},
					label: "ex",
					expr: &choiceExpr{
						pos: position{line: 187, col: 12, offset: 5718},
						alternatives: []interface{}{
							&ruleRefExpr{
								pos:  position{line: 187, col: 12, offset: 5718},
								name: "Call",
							},
							&ruleRefExpr{
								pos:  position{line: 187, col: 19, offset: 5725},
								name: "IfExpr",
							},
							&ruleRefExpr{
								pos:  position{line: 187, col: 28, offset: 5734},
								name: "CompoundExpr",
							},
						},
					},
				},
			},
		},
		{
			name: "IfExpr",
			pos:  position{line: 193, col: 1, offset: 5851},
			expr: &choiceExpr{
				pos: position{line: 193, col: 10, offset: 5860},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 193, col: 10, offset: 5860},
						run: (*parser).callonIfExpr2,
						expr: &seqExpr{
							pos: position{line: 193, col: 10, offset: 5860},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 193, col: 10, offset: 5860},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 193, col: 12, offset: 5862},
									val:        "if",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 193, col: 17, offset: 5867},
									name: "__",
								},
								&labeledExpr{
									pos:   position{line: 193, col: 20, offset: 5870},
									label: "expr",
									expr: &ruleRefExpr{
										pos:  position{line: 193, col: 25, offset: 5875},
										name: "BinOpBool",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 193, col: 35, offset: 5885},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 193, col: 37, offset: 5887},
									val:        "{",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 193, col: 41, offset: 5891},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 193, col: 43, offset: 5893},
									label: "thens",
									expr: &oneOrMoreExpr{
										pos: position{line: 193, col: 49, offset: 5899},
										expr: &ruleRefExpr{
											pos:  position{line: 193, col: 50, offset: 5900},
											name: "Statement",
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 193, col: 62, offset: 5912},
									name: "__",
								},
								&litMatcher{
									pos:        position{line: 193, col: 65, offset: 5915},
									val:        "}",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 193, col: 69, offset: 5919},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 193, col: 71, offset: 5921},
									val:        "else",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 193, col: 78, offset: 5928},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 193, col: 80, offset: 5930},
									val:        "{",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 193, col: 84, offset: 5934},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 193, col: 86, offset: 5936},
									label: "elses",
									expr: &oneOrMoreExpr{
										pos: position{line: 193, col: 92, offset: 5942},
										expr: &ruleRefExpr{
											pos:  position{line: 193, col: 93, offset: 5943},
											name: "Statement",
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 193, col: 105, offset: 5955},
									name: "__",
								},
								&litMatcher{
									pos:        position{line: 193, col: 108, offset: 5958},
									val:        "}",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 212, col: 1, offset: 6461},
						run: (*parser).callonIfExpr27,
						expr: &seqExpr{
							pos: position{line: 212, col: 1, offset: 6461},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 212, col: 1, offset: 6461},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 212, col: 3, offset: 6463},
									val:        "if",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 212, col: 8, offset: 6468},
									name: "__",
								},
								&labeledExpr{
									pos:   position{line: 212, col: 11, offset: 6471},
									label: "expr",
									expr: &ruleRefExpr{
										pos:  position{line: 212, col: 16, offset: 6476},
										name: "BinOpBool",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 212, col: 26, offset: 6486},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 212, col: 28, offset: 6488},
									val:        "{",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 212, col: 32, offset: 6492},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 212, col: 34, offset: 6494},
									label: "thens",
									expr: &oneOrMoreExpr{
										pos: position{line: 212, col: 40, offset: 6500},
										expr: &ruleRefExpr{
											pos:  position{line: 212, col: 41, offset: 6501},
											name: "Statement",
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 212, col: 53, offset: 6513},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 212, col: 55, offset: 6515},
									val:        "}",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 212, col: 59, offset: 6519},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 212, col: 61, offset: 6521},
									val:        "else",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 212, col: 68, offset: 6528},
									label: "elseifs",
									expr: &ruleRefExpr{
										pos:  position{line: 212, col: 76, offset: 6536},
										name: "IfExpr",
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 224, col: 1, offset: 6866},
						run: (*parser).callonIfExpr46,
						expr: &seqExpr{
							pos: position{line: 224, col: 1, offset: 6866},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 224, col: 1, offset: 6866},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 224, col: 3, offset: 6868},
									val:        "if",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 224, col: 8, offset: 6873},
									name: "__",
								},
								&labeledExpr{
									pos:   position{line: 224, col: 11, offset: 6876},
									label: "expr",
									expr: &ruleRefExpr{
										pos:  position{line: 224, col: 16, offset: 6881},
										name: "BinOpBool",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 224, col: 26, offset: 6891},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 224, col: 28, offset: 6893},
									val:        "{",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 224, col: 32, offset: 6897},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 224, col: 34, offset: 6899},
									label: "thens",
									expr: &oneOrMoreExpr{
										pos: position{line: 224, col: 40, offset: 6905},
										expr: &ruleRefExpr{
											pos:  position{line: 224, col: 41, offset: 6906},
											name: "Statement",
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 224, col: 53, offset: 6918},
									name: "__",
								},
								&litMatcher{
									pos:        position{line: 224, col: 56, offset: 6921},
									val:        "}",
									ignoreCase: false,
								},
							},
						},
					},
				},
			},
		},
		{
			name: "Call",
			pos:  position{line: 236, col: 1, offset: 7219},
			expr: &choiceExpr{
				pos: position{line: 236, col: 8, offset: 7226},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 236, col: 8, offset: 7226},
						run: (*parser).callonCall2,
						expr: &seqExpr{
							pos: position{line: 236, col: 8, offset: 7226},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 236, col: 8, offset: 7226},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 236, col: 10, offset: 7228},
									label: "module",
									expr: &ruleRefExpr{
										pos:  position{line: 236, col: 17, offset: 7235},
										name: "ModuleName",
									},
								},
								&litMatcher{
									pos:        position{line: 236, col: 28, offset: 7246},
									val:        ".",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 236, col: 32, offset: 7250},
									label: "fn",
									expr: &ruleRefExpr{
										pos:  position{line: 236, col: 35, offset: 7253},
										name: "VariableName",
									},
								},
								&labeledExpr{
									pos:   position{line: 236, col: 48, offset: 7266},
									label: "args",
									expr: &ruleRefExpr{
										pos:  position{line: 236, col: 53, offset: 7271},
										name: "Arguments",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 236, col: 63, offset: 7281},
									name: "_",
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 250, col: 1, offset: 7597},
						run: (*parser).callonCall13,
						expr: &seqExpr{
							pos: position{line: 250, col: 1, offset: 7597},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 250, col: 1, offset: 7597},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 250, col: 3, offset: 7599},
									label: "fn",
									expr: &ruleRefExpr{
										pos:  position{line: 250, col: 6, offset: 7602},
										name: "VariableName",
									},
								},
								&labeledExpr{
									pos:   position{line: 250, col: 19, offset: 7615},
									label: "args",
									expr: &ruleRefExpr{
										pos:  position{line: 250, col: 24, offset: 7620},
										name: "Arguments",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 250, col: 34, offset: 7630},
									name: "_",
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 264, col: 1, offset: 7939},
						run: (*parser).callonCall21,
						expr: &seqExpr{
							pos: position{line: 264, col: 1, offset: 7939},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 264, col: 1, offset: 7939},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 264, col: 3, offset: 7941},
									label: "fn",
									expr: &ruleRefExpr{
										pos:  position{line: 264, col: 6, offset: 7944},
										name: "VariableName",
									},
								},
								&litMatcher{
									pos:        position{line: 264, col: 19, offset: 7957},
									val:        "()",
									ignoreCase: false,
								},
							},
						},
					},
				},
			},
		},
		{
			name: "ArgsDefn",
			pos:  position{line: 271, col: 1, offset: 8128},
			expr: &actionExpr{
				pos: position{line: 271, col: 12, offset: 8139},
				run: (*parser).callonArgsDefn1,
				expr: &seqExpr{
					pos: position{line: 271, col: 12, offset: 8139},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 271, col: 12, offset: 8139},
							val:        "(",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 271, col: 16, offset: 8143},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 271, col: 18, offset: 8145},
							label: "argument",
							expr: &ruleRefExpr{
								pos:  position{line: 271, col: 27, offset: 8154},
								name: "ArgDefn",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 271, col: 35, offset: 8162},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 271, col: 37, offset: 8164},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 271, col: 42, offset: 8169},
								expr: &seqExpr{
									pos: position{line: 271, col: 43, offset: 8170},
									exprs: []interface{}{
										&litMatcher{
											pos:        position{line: 271, col: 43, offset: 8170},
											val:        ",",
											ignoreCase: false,
										},
										&ruleRefExpr{
											pos:  position{line: 271, col: 47, offset: 8174},
											name: "_",
										},
										&ruleRefExpr{
											pos:  position{line: 271, col: 49, offset: 8176},
											name: "ArgDefn",
										},
									},
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 271, col: 59, offset: 8186},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 271, col: 61, offset: 8188},
							val:        ")",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "ArgDefn",
			pos:  position{line: 289, col: 1, offset: 8610},
			expr: &actionExpr{
				pos: position{line: 289, col: 11, offset: 8620},
				run: (*parser).callonArgDefn1,
				expr: &seqExpr{
					pos: position{line: 289, col: 11, offset: 8620},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 289, col: 11, offset: 8620},
							label: "name",
							expr: &ruleRefExpr{
								pos:  position{line: 289, col: 16, offset: 8625},
								name: "Identifier",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 289, col: 27, offset: 8636},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 289, col: 29, offset: 8638},
							label: "anno",
							expr: &zeroOrOneExpr{
								pos: position{line: 289, col: 34, offset: 8643},
								expr: &seqExpr{
									pos: position{line: 289, col: 35, offset: 8644},
									exprs: []interface{}{
										&litMatcher{
											pos:        position{line: 289, col: 35, offset: 8644},
											val:        ":",
											ignoreCase: false,
										},
										&ruleRefExpr{
											pos:  position{line: 289, col: 39, offset: 8648},
											name: "_",
										},
										&ruleRefExpr{
											pos:  position{line: 289, col: 41, offset: 8650},
											name: "AnyType",
										},
									},
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 289, col: 52, offset: 8661},
							name: "_",
						},
					},
				},
			},
		},
		{
			name: "Arguments",
			pos:  position{line: 309, col: 1, offset: 9149},
			expr: &choiceExpr{
				pos: position{line: 309, col: 13, offset: 9161},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 309, col: 13, offset: 9161},
						run: (*parser).callonArguments2,
						expr: &seqExpr{
							pos: position{line: 309, col: 13, offset: 9161},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 309, col: 13, offset: 9161},
									val:        "(",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 309, col: 17, offset: 9165},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 309, col: 19, offset: 9167},
									label: "argument",
									expr: &ruleRefExpr{
										pos:  position{line: 309, col: 28, offset: 9176},
										name: "BinOpParens",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 309, col: 40, offset: 9188},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 309, col: 42, offset: 9190},
									label: "rest",
									expr: &zeroOrMoreExpr{
										pos: position{line: 309, col: 47, offset: 9195},
										expr: &seqExpr{
											pos: position{line: 309, col: 48, offset: 9196},
											exprs: []interface{}{
												&litMatcher{
													pos:        position{line: 309, col: 48, offset: 9196},
													val:        ",",
													ignoreCase: false,
												},
												&ruleRefExpr{
													pos:  position{line: 309, col: 52, offset: 9200},
													name: "_",
												},
												&ruleRefExpr{
													pos:  position{line: 309, col: 54, offset: 9202},
													name: "BinOpParens",
												},
											},
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 309, col: 68, offset: 9216},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 309, col: 70, offset: 9218},
									val:        ")",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 326, col: 1, offset: 9640},
						run: (*parser).callonArguments17,
						expr: &seqExpr{
							pos: position{line: 326, col: 1, offset: 9640},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 326, col: 1, offset: 9640},
									val:        "(",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 326, col: 5, offset: 9644},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 326, col: 7, offset: 9646},
									label: "argument",
									expr: &ruleRefExpr{
										pos:  position{line: 326, col: 16, offset: 9655},
										name: "Expr",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 326, col: 21, offset: 9660},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 326, col: 23, offset: 9662},
									val:        ")",
									ignoreCase: false,
								},
							},
						},
					},
				},
			},
		},
		{
			name: "CompoundExpr",
			pos:  position{line: 331, col: 1, offset: 9767},
			expr: &actionExpr{
				pos: position{line: 331, col: 16, offset: 9782},
				run: (*parser).callonCompoundExpr1,
				expr: &seqExpr{
					pos: position{line: 331, col: 16, offset: 9782},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 331, col: 16, offset: 9782},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 331, col: 18, offset: 9784},
							label: "op",
							expr: &ruleRefExpr{
								pos:  position{line: 331, col: 21, offset: 9787},
								name: "BinOp",
							},
						},
						&labeledExpr{
							pos:   position{line: 331, col: 27, offset: 9793},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 331, col: 32, offset: 9798},
								expr: &seqExpr{
									pos: position{line: 331, col: 33, offset: 9799},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 331, col: 33, offset: 9799},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 331, col: 36, offset: 9802},
											name: "Operator",
										},
										&ruleRefExpr{
											pos:  position{line: 331, col: 45, offset: 9811},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 331, col: 48, offset: 9814},
											name: "BinOp",
										},
									},
								},
							},
						},
					},
				},
			},
		},
		{
			name: "BinOp",
			pos:  position{line: 351, col: 1, offset: 10420},
			expr: &choiceExpr{
				pos: position{line: 351, col: 9, offset: 10428},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 351, col: 9, offset: 10428},
						name: "BinOpBool",
					},
					&ruleRefExpr{
						pos:  position{line: 351, col: 21, offset: 10440},
						name: "BinOpEquality",
					},
					&ruleRefExpr{
						pos:  position{line: 351, col: 37, offset: 10456},
						name: "BinOpLow",
					},
					&ruleRefExpr{
						pos:  position{line: 351, col: 48, offset: 10467},
						name: "BinOpHigh",
					},
					&ruleRefExpr{
						pos:  position{line: 351, col: 60, offset: 10479},
						name: "BinOpParens",
					},
				},
			},
		},
		{
			name: "BinOpBool",
			pos:  position{line: 353, col: 1, offset: 10492},
			expr: &actionExpr{
				pos: position{line: 353, col: 13, offset: 10504},
				run: (*parser).callonBinOpBool1,
				expr: &seqExpr{
					pos: position{line: 353, col: 13, offset: 10504},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 353, col: 13, offset: 10504},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 353, col: 15, offset: 10506},
							label: "first",
							expr: &ruleRefExpr{
								pos:  position{line: 353, col: 21, offset: 10512},
								name: "BinOpEquality",
							},
						},
						&labeledExpr{
							pos:   position{line: 353, col: 35, offset: 10526},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 353, col: 40, offset: 10531},
								expr: &seqExpr{
									pos: position{line: 353, col: 41, offset: 10532},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 353, col: 41, offset: 10532},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 353, col: 44, offset: 10535},
											name: "OperatorBoolean",
										},
										&ruleRefExpr{
											pos:  position{line: 353, col: 60, offset: 10551},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 353, col: 63, offset: 10554},
											name: "BinOpEquality",
										},
									},
								},
							},
						},
					},
				},
			},
		},
		{
			name: "BinOpEquality",
			pos:  position{line: 386, col: 1, offset: 11447},
			expr: &actionExpr{
				pos: position{line: 386, col: 17, offset: 11463},
				run: (*parser).callonBinOpEquality1,
				expr: &seqExpr{
					pos: position{line: 386, col: 17, offset: 11463},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 386, col: 17, offset: 11463},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 386, col: 19, offset: 11465},
							label: "first",
							expr: &ruleRefExpr{
								pos:  position{line: 386, col: 25, offset: 11471},
								name: "BinOpLow",
							},
						},
						&labeledExpr{
							pos:   position{line: 386, col: 34, offset: 11480},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 386, col: 39, offset: 11485},
								expr: &seqExpr{
									pos: position{line: 386, col: 40, offset: 11486},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 386, col: 40, offset: 11486},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 386, col: 43, offset: 11489},
											name: "OperatorEquality",
										},
										&ruleRefExpr{
											pos:  position{line: 386, col: 60, offset: 11506},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 386, col: 63, offset: 11509},
											name: "BinOpLow",
										},
									},
								},
							},
						},
					},
				},
			},
		},
		{
			name: "BinOpLow",
			pos:  position{line: 418, col: 1, offset: 12396},
			expr: &actionExpr{
				pos: position{line: 418, col: 12, offset: 12407},
				run: (*parser).callonBinOpLow1,
				expr: &seqExpr{
					pos: position{line: 418, col: 12, offset: 12407},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 418, col: 12, offset: 12407},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 418, col: 14, offset: 12409},
							label: "first",
							expr: &ruleRefExpr{
								pos:  position{line: 418, col: 20, offset: 12415},
								name: "BinOpHigh",
							},
						},
						&labeledExpr{
							pos:   position{line: 418, col: 30, offset: 12425},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 418, col: 35, offset: 12430},
								expr: &seqExpr{
									pos: position{line: 418, col: 36, offset: 12431},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 418, col: 36, offset: 12431},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 418, col: 39, offset: 12434},
											name: "OperatorLow",
										},
										&ruleRefExpr{
											pos:  position{line: 418, col: 51, offset: 12446},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 418, col: 54, offset: 12449},
											name: "BinOpHigh",
										},
									},
								},
							},
						},
					},
				},
			},
		},
		{
			name: "BinOpHigh",
			pos:  position{line: 450, col: 1, offset: 13337},
			expr: &actionExpr{
				pos: position{line: 450, col: 13, offset: 13349},
				run: (*parser).callonBinOpHigh1,
				expr: &seqExpr{
					pos: position{line: 450, col: 13, offset: 13349},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 450, col: 13, offset: 13349},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 450, col: 15, offset: 13351},
							label: "first",
							expr: &ruleRefExpr{
								pos:  position{line: 450, col: 21, offset: 13357},
								name: "BinOpParens",
							},
						},
						&labeledExpr{
							pos:   position{line: 450, col: 33, offset: 13369},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 450, col: 38, offset: 13374},
								expr: &seqExpr{
									pos: position{line: 450, col: 39, offset: 13375},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 450, col: 39, offset: 13375},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 450, col: 42, offset: 13378},
											name: "OperatorHigh",
										},
										&ruleRefExpr{
											pos:  position{line: 450, col: 55, offset: 13391},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 450, col: 58, offset: 13394},
											name: "BinOpParens",
										},
									},
								},
							},
						},
					},
				},
			},
		},
		{
			name: "BinOpParens",
			pos:  position{line: 481, col: 1, offset: 14283},
			expr: &choiceExpr{
				pos: position{line: 481, col: 15, offset: 14297},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 481, col: 15, offset: 14297},
						run: (*parser).callonBinOpParens2,
						expr: &seqExpr{
							pos: position{line: 481, col: 15, offset: 14297},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 481, col: 15, offset: 14297},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 481, col: 17, offset: 14299},
									val:        "(",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 481, col: 21, offset: 14303},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 481, col: 23, offset: 14305},
									label: "first",
									expr: &ruleRefExpr{
										pos:  position{line: 481, col: 29, offset: 14311},
										name: "BinOp",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 481, col: 35, offset: 14317},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 481, col: 37, offset: 14319},
									val:        ")",
									ignoreCase: false,
								},
							},
						},
					},
					&ruleRefExpr{
						pos:  position{line: 484, col: 5, offset: 14442},
						name: "Value",
					},
				},
			},
		},
		{
			name: "Operator",
			pos:  position{line: 486, col: 1, offset: 14449},
			expr: &choiceExpr{
				pos: position{line: 486, col: 12, offset: 14460},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 486, col: 12, offset: 14460},
						name: "OperatorBoolean",
					},
					&ruleRefExpr{
						pos:  position{line: 486, col: 30, offset: 14478},
						name: "OperatorEquality",
					},
					&ruleRefExpr{
						pos:  position{line: 486, col: 49, offset: 14497},
						name: "OperatorHigh",
					},
					&ruleRefExpr{
						pos:  position{line: 486, col: 64, offset: 14512},
						name: "OperatorLow",
					},
				},
			},
		},
		{
			name: "OperatorBoolean",
			pos:  position{line: 488, col: 1, offset: 14525},
			expr: &actionExpr{
				pos: position{line: 488, col: 19, offset: 14543},
				run: (*parser).callonOperatorBoolean1,
				expr: &choiceExpr{
					pos: position{line: 488, col: 21, offset: 14545},
					alternatives: []interface{}{
						&litMatcher{
							pos:        position{line: 488, col: 21, offset: 14545},
							val:        "or",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 488, col: 28, offset: 14552},
							val:        "and",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "OperatorEquality",
			pos:  position{line: 492, col: 1, offset: 14634},
			expr: &actionExpr{
				pos: position{line: 492, col: 20, offset: 14653},
				run: (*parser).callonOperatorEquality1,
				expr: &choiceExpr{
					pos: position{line: 492, col: 22, offset: 14655},
					alternatives: []interface{}{
						&litMatcher{
							pos:        position{line: 492, col: 22, offset: 14655},
							val:        ">=",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 492, col: 29, offset: 14662},
							val:        "<=",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 492, col: 36, offset: 14669},
							val:        "<",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 492, col: 42, offset: 14675},
							val:        ">",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 492, col: 48, offset: 14681},
							val:        "===",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 492, col: 56, offset: 14689},
							val:        "==",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "OperatorHigh",
			pos:  position{line: 496, col: 1, offset: 14768},
			expr: &choiceExpr{
				pos: position{line: 496, col: 16, offset: 14783},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 496, col: 16, offset: 14783},
						run: (*parser).callonOperatorHigh2,
						expr: &choiceExpr{
							pos: position{line: 496, col: 18, offset: 14785},
							alternatives: []interface{}{
								&litMatcher{
									pos:        position{line: 496, col: 18, offset: 14785},
									val:        "/",
									ignoreCase: false,
								},
								&litMatcher{
									pos:        position{line: 496, col: 24, offset: 14791},
									val:        "*",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 499, col: 3, offset: 14874},
						run: (*parser).callonOperatorHigh6,
						expr: &litMatcher{
							pos:        position{line: 499, col: 5, offset: 14876},
							val:        "^",
							ignoreCase: false,
						},
					},
					&actionExpr{
						pos: position{line: 502, col: 3, offset: 14956},
						run: (*parser).callonOperatorHigh8,
						expr: &litMatcher{
							pos:        position{line: 502, col: 3, offset: 14956},
							val:        "++",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "OperatorLow",
			pos:  position{line: 506, col: 1, offset: 15037},
			expr: &actionExpr{
				pos: position{line: 506, col: 15, offset: 15051},
				run: (*parser).callonOperatorLow1,
				expr: &choiceExpr{
					pos: position{line: 506, col: 17, offset: 15053},
					alternatives: []interface{}{
						&litMatcher{
							pos:        position{line: 506, col: 17, offset: 15053},
							val:        "+",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 506, col: 23, offset: 15059},
							val:        "-",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "Value",
			pos:  position{line: 510, col: 1, offset: 15141},
			expr: &choiceExpr{
				pos: position{line: 510, col: 9, offset: 15149},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 510, col: 9, offset: 15149},
						name: "Call",
					},
					&ruleRefExpr{
						pos:  position{line: 510, col: 16, offset: 15156},
						name: "VariableName",
					},
					&ruleRefExpr{
						pos:  position{line: 510, col: 31, offset: 15171},
						name: "Const",
					},
				},
			},
		},
		{
			name: "Assignable",
			pos:  position{line: 512, col: 1, offset: 15178},
			expr: &choiceExpr{
				pos: position{line: 512, col: 14, offset: 15191},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 512, col: 14, offset: 15191},
						name: "VariableName",
					},
					&ruleRefExpr{
						pos:  position{line: 512, col: 29, offset: 15206},
						name: "Unused",
					},
				},
			},
		},
		{
			name: "Identifier",
			pos:  position{line: 514, col: 1, offset: 15214},
			expr: &choiceExpr{
				pos: position{line: 514, col: 14, offset: 15227},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 514, col: 14, offset: 15227},
						name: "VariableName",
					},
					&ruleRefExpr{
						pos:  position{line: 514, col: 29, offset: 15242},
						name: "ModuleName",
					},
				},
			},
		},
		{
			name: "ArrayLiteral",
			pos:  position{line: 516, col: 1, offset: 15254},
			expr: &actionExpr{
				pos: position{line: 516, col: 16, offset: 15269},
				run: (*parser).callonArrayLiteral1,
				expr: &seqExpr{
					pos: position{line: 516, col: 16, offset: 15269},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 516, col: 16, offset: 15269},
							val:        "[",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 516, col: 20, offset: 15273},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 516, col: 22, offset: 15275},
							label: "first",
							expr: &ruleRefExpr{
								pos:  position{line: 516, col: 28, offset: 15281},
								name: "Expr",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 516, col: 33, offset: 15286},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 516, col: 35, offset: 15288},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 516, col: 40, offset: 15293},
								expr: &seqExpr{
									pos: position{line: 516, col: 41, offset: 15294},
									exprs: []interface{}{
										&litMatcher{
											pos:        position{line: 516, col: 41, offset: 15294},
											val:        ",",
											ignoreCase: false,
										},
										&ruleRefExpr{
											pos:  position{line: 516, col: 45, offset: 15298},
											name: "_",
										},
										&ruleRefExpr{
											pos:  position{line: 516, col: 47, offset: 15300},
											name: "Expr",
										},
										&ruleRefExpr{
											pos:  position{line: 516, col: 52, offset: 15305},
											name: "_",
										},
									},
								},
							},
						},
						&zeroOrOneExpr{
							pos: position{line: 516, col: 56, offset: 15309},
							expr: &litMatcher{
								pos:        position{line: 516, col: 56, offset: 15309},
								val:        ",",
								ignoreCase: false,
							},
						},
						&ruleRefExpr{
							pos:  position{line: 516, col: 61, offset: 15314},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 516, col: 63, offset: 15316},
							val:        "]",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "VariantInstance",
			pos:  position{line: 532, col: 1, offset: 15761},
			expr: &actionExpr{
				pos: position{line: 532, col: 19, offset: 15779},
				run: (*parser).callonVariantInstance1,
				expr: &seqExpr{
					pos: position{line: 532, col: 19, offset: 15779},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 532, col: 19, offset: 15779},
							label: "name",
							expr: &ruleRefExpr{
								pos:  position{line: 532, col: 24, offset: 15784},
								name: "ModuleName",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 532, col: 35, offset: 15795},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 532, col: 37, offset: 15797},
							label: "args",
							expr: &ruleRefExpr{
								pos:  position{line: 532, col: 42, offset: 15802},
								name: "Arguments",
							},
						},
					},
				},
			},
		},
		{
			name: "RecordInstance",
			pos:  position{line: 545, col: 1, offset: 16072},
			expr: &actionExpr{
				pos: position{line: 545, col: 18, offset: 16089},
				run: (*parser).callonRecordInstance1,
				expr: &seqExpr{
					pos: position{line: 545, col: 18, offset: 16089},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 545, col: 18, offset: 16089},
							label: "name",
							expr: &ruleRefExpr{
								pos:  position{line: 545, col: 23, offset: 16094},
								name: "ModuleName",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 545, col: 34, offset: 16105},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 545, col: 36, offset: 16107},
							val:        "{",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 545, col: 40, offset: 16111},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 545, col: 42, offset: 16113},
							label: "firstName",
							expr: &ruleRefExpr{
								pos:  position{line: 545, col: 52, offset: 16123},
								name: "VariableName",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 545, col: 65, offset: 16136},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 545, col: 67, offset: 16138},
							val:        ":",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 545, col: 71, offset: 16142},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 545, col: 73, offset: 16144},
							label: "firstValue",
							expr: &ruleRefExpr{
								pos:  position{line: 545, col: 84, offset: 16155},
								name: "Expr",
							},
						},
						&labeledExpr{
							pos:   position{line: 545, col: 89, offset: 16160},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 545, col: 94, offset: 16165},
								expr: &seqExpr{
									pos: position{line: 545, col: 95, offset: 16166},
									exprs: []interface{}{
										&litMatcher{
											pos:        position{line: 545, col: 95, offset: 16166},
											val:        ",",
											ignoreCase: false,
										},
										&ruleRefExpr{
											pos:  position{line: 545, col: 99, offset: 16170},
											name: "_",
										},
										&ruleRefExpr{
											pos:  position{line: 545, col: 101, offset: 16172},
											name: "VariableName",
										},
										&ruleRefExpr{
											pos:  position{line: 545, col: 114, offset: 16185},
											name: "_",
										},
										&litMatcher{
											pos:        position{line: 545, col: 116, offset: 16187},
											val:        ":",
											ignoreCase: false,
										},
										&ruleRefExpr{
											pos:  position{line: 545, col: 120, offset: 16191},
											name: "_",
										},
										&ruleRefExpr{
											pos:  position{line: 545, col: 122, offset: 16193},
											name: "Expr",
										},
									},
								},
							},
						},
						&litMatcher{
							pos:        position{line: 545, col: 130, offset: 16201},
							val:        "}",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "TypeParameter",
			pos:  position{line: 565, col: 1, offset: 16785},
			expr: &actionExpr{
				pos: position{line: 565, col: 17, offset: 16801},
				run: (*parser).callonTypeParameter1,
				expr: &seqExpr{
					pos: position{line: 565, col: 17, offset: 16801},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 565, col: 17, offset: 16801},
							val:        "'",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 565, col: 22, offset: 16806},
							name: "VariableName",
						},
					},
				},
			},
		},
		{
			name: "VariableName",
			pos:  position{line: 569, col: 1, offset: 16879},
			expr: &actionExpr{
				pos: position{line: 569, col: 16, offset: 16894},
				run: (*parser).callonVariableName1,
				expr: &seqExpr{
					pos: position{line: 569, col: 16, offset: 16894},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 569, col: 16, offset: 16894},
							expr: &ruleRefExpr{
								pos:  position{line: 569, col: 17, offset: 16895},
								name: "Reserved",
							},
						},
						&seqExpr{
							pos: position{line: 569, col: 27, offset: 16905},
							exprs: []interface{}{
								&oneOrMoreExpr{
									pos: position{line: 569, col: 27, offset: 16905},
									expr: &charClassMatcher{
										pos:        position{line: 569, col: 27, offset: 16905},
										val:        "[a-z]",
										ranges:     []rune{'a', 'z'},
										ignoreCase: false,
										inverted:   false,
									},
								},
								&zeroOrMoreExpr{
									pos: position{line: 569, col: 34, offset: 16912},
									expr: &charClassMatcher{
										pos:        position{line: 569, col: 34, offset: 16912},
										val:        "[a-zA-Z0-9_]",
										chars:      []rune{'_'},
										ranges:     []rune{'a', 'z', 'A', 'Z', '0', '9'},
										ignoreCase: false,
										inverted:   false,
									},
								},
							},
						},
					},
				},
			},
		},
		{
			name: "ModuleName",
			pos:  position{line: 573, col: 1, offset: 16987},
			expr: &actionExpr{
				pos: position{line: 573, col: 14, offset: 17000},
				run: (*parser).callonModuleName1,
				expr: &seqExpr{
					pos: position{line: 573, col: 15, offset: 17001},
					exprs: []interface{}{
						&oneOrMoreExpr{
							pos: position{line: 573, col: 15, offset: 17001},
							expr: &charClassMatcher{
								pos:        position{line: 573, col: 15, offset: 17001},
								val:        "[A-Z]",
								ranges:     []rune{'A', 'Z'},
								ignoreCase: false,
								inverted:   false,
							},
						},
						&zeroOrMoreExpr{
							pos: position{line: 573, col: 22, offset: 17008},
							expr: &charClassMatcher{
								pos:        position{line: 573, col: 22, offset: 17008},
								val:        "[a-zA-Z0-9_]",
								chars:      []rune{'_'},
								ranges:     []rune{'a', 'z', 'A', 'Z', '0', '9'},
								ignoreCase: false,
								inverted:   false,
							},
						},
					},
				},
			},
		},
		{
			name: "Const",
			pos:  position{line: 577, col: 1, offset: 17083},
			expr: &choiceExpr{
				pos: position{line: 577, col: 9, offset: 17091},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 577, col: 9, offset: 17091},
						run: (*parser).callonConst2,
						expr: &seqExpr{
							pos: position{line: 577, col: 9, offset: 17091},
							exprs: []interface{}{
								&zeroOrOneExpr{
									pos: position{line: 577, col: 9, offset: 17091},
									expr: &litMatcher{
										pos:        position{line: 577, col: 9, offset: 17091},
										val:        "-",
										ignoreCase: false,
									},
								},
								&oneOrMoreExpr{
									pos: position{line: 577, col: 14, offset: 17096},
									expr: &charClassMatcher{
										pos:        position{line: 577, col: 14, offset: 17096},
										val:        "[0-9]",
										ranges:     []rune{'0', '9'},
										ignoreCase: false,
										inverted:   false,
									},
								},
								&notExpr{
									pos: position{line: 577, col: 21, offset: 17103},
									expr: &litMatcher{
										pos:        position{line: 577, col: 22, offset: 17104},
										val:        ".",
										ignoreCase: false,
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 584, col: 3, offset: 17279},
						run: (*parser).callonConst10,
						expr: &seqExpr{
							pos: position{line: 584, col: 3, offset: 17279},
							exprs: []interface{}{
								&zeroOrOneExpr{
									pos: position{line: 584, col: 3, offset: 17279},
									expr: &litMatcher{
										pos:        position{line: 584, col: 3, offset: 17279},
										val:        "-",
										ignoreCase: false,
									},
								},
								&oneOrMoreExpr{
									pos: position{line: 584, col: 8, offset: 17284},
									expr: &charClassMatcher{
										pos:        position{line: 584, col: 8, offset: 17284},
										val:        "[0-9]",
										ranges:     []rune{'0', '9'},
										ignoreCase: false,
										inverted:   false,
									},
								},
								&litMatcher{
									pos:        position{line: 584, col: 15, offset: 17291},
									val:        ".",
									ignoreCase: false,
								},
								&oneOrMoreExpr{
									pos: position{line: 584, col: 19, offset: 17295},
									expr: &charClassMatcher{
										pos:        position{line: 584, col: 19, offset: 17295},
										val:        "[0-9]",
										ranges:     []rune{'0', '9'},
										ignoreCase: false,
										inverted:   false,
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 591, col: 3, offset: 17484},
						run: (*parser).callonConst19,
						expr: &litMatcher{
							pos:        position{line: 591, col: 3, offset: 17484},
							val:        "true",
							ignoreCase: false,
						},
					},
					&actionExpr{
						pos: position{line: 595, col: 3, offset: 17569},
						run: (*parser).callonConst21,
						expr: &litMatcher{
							pos:        position{line: 595, col: 3, offset: 17569},
							val:        "false",
							ignoreCase: false,
						},
					},
					&ruleRefExpr{
						pos:  position{line: 598, col: 3, offset: 17655},
						name: "Unit",
					},
					&ruleRefExpr{
						pos:  position{line: 599, col: 3, offset: 17662},
						name: "StringLiteral",
					},
					&actionExpr{
						pos: position{line: 600, col: 3, offset: 17678},
						run: (*parser).callonConst25,
						expr: &seqExpr{
							pos: position{line: 600, col: 3, offset: 17678},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 600, col: 3, offset: 17678},
									val:        "'",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 600, col: 7, offset: 17682},
									label: "val",
									expr: &seqExpr{
										pos: position{line: 600, col: 12, offset: 17687},
										exprs: []interface{}{
											&notExpr{
												pos: position{line: 600, col: 12, offset: 17687},
												expr: &ruleRefExpr{
													pos:  position{line: 600, col: 13, offset: 17688},
													name: "EscapedChar",
												},
											},
											&anyMatcher{
												line: 600, col: 25, offset: 17700,
											},
										},
									},
								},
								&litMatcher{
									pos:        position{line: 600, col: 28, offset: 17703},
									val:        "'",
									ignoreCase: false,
								},
							},
						},
					},
					&ruleRefExpr{
						pos:  position{line: 602, col: 5, offset: 17795},
						name: "ArrayLiteral",
					},
					&ruleRefExpr{
						pos:  position{line: 602, col: 20, offset: 17810},
						name: "RecordInstance",
					},
					&ruleRefExpr{
						pos:  position{line: 602, col: 37, offset: 17827},
						name: "VariantInstance",
					},
				},
			},
		},
		{
			name: "Unit",
			pos:  position{line: 604, col: 1, offset: 17844},
			expr: &actionExpr{
				pos: position{line: 604, col: 8, offset: 17851},
				run: (*parser).callonUnit1,
				expr: &litMatcher{
					pos:        position{line: 604, col: 8, offset: 17851},
					val:        "()",
					ignoreCase: false,
				},
			},
		},
		{
			name: "StringLiteral",
			pos:  position{line: 608, col: 1, offset: 17914},
			expr: &actionExpr{
				pos: position{line: 608, col: 17, offset: 17930},
				run: (*parser).callonStringLiteral1,
				expr: &seqExpr{
					pos: position{line: 608, col: 17, offset: 17930},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 608, col: 17, offset: 17930},
							val:        "\"",
							ignoreCase: false,
						},
						&zeroOrMoreExpr{
							pos: position{line: 608, col: 21, offset: 17934},
							expr: &seqExpr{
								pos: position{line: 608, col: 22, offset: 17935},
								exprs: []interface{}{
									&notExpr{
										pos: position{line: 608, col: 22, offset: 17935},
										expr: &ruleRefExpr{
											pos:  position{line: 608, col: 23, offset: 17936},
											name: "EscapedChar",
										},
									},
									&anyMatcher{
										line: 608, col: 35, offset: 17948,
									},
								},
							},
						},
						&litMatcher{
							pos:        position{line: 608, col: 39, offset: 17952},
							val:        "\"",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "Unused",
			pos:  position{line: 616, col: 1, offset: 18135},
			expr: &actionExpr{
				pos: position{line: 616, col: 10, offset: 18144},
				run: (*parser).callonUnused1,
				expr: &litMatcher{
					pos:        position{line: 616, col: 11, offset: 18145},
					val:        "_",
					ignoreCase: false,
				},
			},
		},
		{
			name: "Reserved",
			pos:  position{line: 620, col: 1, offset: 18200},
			expr: &seqExpr{
				pos: position{line: 620, col: 12, offset: 18211},
				exprs: []interface{}{
					&choiceExpr{
						pos: position{line: 620, col: 13, offset: 18212},
						alternatives: []interface{}{
							&litMatcher{
								pos:        position{line: 620, col: 13, offset: 18212},
								val:        "let",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 620, col: 21, offset: 18220},
								val:        "if",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 620, col: 28, offset: 18227},
								val:        "else",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 620, col: 37, offset: 18236},
								val:        "func",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 620, col: 46, offset: 18245},
								val:        "type",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 620, col: 55, offset: 18254},
								val:        "true",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 620, col: 64, offset: 18263},
								val:        "false",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 620, col: 74, offset: 18273},
								val:        "mutable",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 620, col: 86, offset: 18285},
								val:        "type",
								ignoreCase: false,
							},
							&ruleRefExpr{
								pos:  position{line: 620, col: 95, offset: 18294},
								name: "BaseType",
							},
						},
					},
					&notExpr{
						pos: position{line: 620, col: 105, offset: 18304},
						expr: &oneOrMoreExpr{
							pos: position{line: 620, col: 106, offset: 18305},
							expr: &charClassMatcher{
								pos:        position{line: 620, col: 106, offset: 18305},
								val:        "[a-z]",
								ranges:     []rune{'a', 'z'},
								ignoreCase: false,
								inverted:   false,
							},
						},
					},
				},
			},
		},
		{
			name: "BaseType",
			pos:  position{line: 622, col: 1, offset: 18313},
			expr: &choiceExpr{
				pos: position{line: 622, col: 12, offset: 18324},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 622, col: 12, offset: 18324},
						run: (*parser).callonBaseType2,
						expr: &choiceExpr{
							pos: position{line: 622, col: 14, offset: 18326},
							alternatives: []interface{}{
								&litMatcher{
									pos:        position{line: 622, col: 14, offset: 18326},
									val:        "int",
									ignoreCase: false,
								},
								&litMatcher{
									pos:        position{line: 622, col: 22, offset: 18334},
									val:        "bool",
									ignoreCase: false,
								},
								&litMatcher{
									pos:        position{line: 622, col: 31, offset: 18343},
									val:        "string",
									ignoreCase: false,
								},
								&litMatcher{
									pos:        position{line: 622, col: 42, offset: 18354},
									val:        "byte",
									ignoreCase: false,
								},
								&litMatcher{
									pos:        position{line: 622, col: 51, offset: 18363},
									val:        "rune",
									ignoreCase: false,
								},
								&litMatcher{
									pos:        position{line: 622, col: 60, offset: 18372},
									val:        "float",
									ignoreCase: false,
								},
								&litMatcher{
									pos:        position{line: 622, col: 70, offset: 18382},
									val:        "()",
									ignoreCase: false,
								},
							},
						},
					},
					&ruleRefExpr{
						pos:  position{line: 625, col: 3, offset: 18479},
						name: "Unit",
					},
				},
			},
		},
		{
			name: "EscapedChar",
			pos:  position{line: 627, col: 1, offset: 18485},
			expr: &charClassMatcher{
				pos:        position{line: 627, col: 15, offset: 18499},
				val:        "[\\x00-\\x1f\"\\\\]",
				chars:      []rune{'"', '\\'},
				ranges:     []rune{'\x00', '\x1f'},
				ignoreCase: false,
				inverted:   false,
			},
		},
		{
			name: "EscapeSequence",
			pos:  position{line: 629, col: 1, offset: 18515},
			expr: &choiceExpr{
				pos: position{line: 629, col: 18, offset: 18532},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 629, col: 18, offset: 18532},
						name: "SingleCharEscape",
					},
					&ruleRefExpr{
						pos:  position{line: 629, col: 37, offset: 18551},
						name: "UnicodeEscape",
					},
				},
			},
		},
		{
			name: "SingleCharEscape",
			pos:  position{line: 631, col: 1, offset: 18566},
			expr: &charClassMatcher{
				pos:        position{line: 631, col: 20, offset: 18585},
				val:        "[\"\\\\/bfnrt]",
				chars:      []rune{'"', '\\', '/', 'b', 'f', 'n', 'r', 't'},
				ignoreCase: false,
				inverted:   false,
			},
		},
		{
			name: "DecimalDigit",
			pos:  position{line: 633, col: 1, offset: 18598},
			expr: &charClassMatcher{
				pos:        position{line: 633, col: 16, offset: 18613},
				val:        "[0-9]",
				ranges:     []rune{'0', '9'},
				ignoreCase: false,
				inverted:   false,
			},
		},
		{
			name: "NonZeroDecimalDigit",
			pos:  position{line: 635, col: 1, offset: 18620},
			expr: &charClassMatcher{
				pos:        position{line: 635, col: 23, offset: 18642},
				val:        "[1-9]",
				ranges:     []rune{'1', '9'},
				ignoreCase: false,
				inverted:   false,
			},
		},
		{
			name: "HexDigit",
			pos:  position{line: 637, col: 1, offset: 18649},
			expr: &charClassMatcher{
				pos:        position{line: 637, col: 12, offset: 18660},
				val:        "[0-9a-f]i",
				ranges:     []rune{'0', '9', 'a', 'f'},
				ignoreCase: true,
				inverted:   false,
			},
		},
		{
			name:        "__",
			displayName: "\"sigwhitespace\"",
			pos:         position{line: 639, col: 1, offset: 18671},
			expr: &choiceExpr{
				pos: position{line: 639, col: 22, offset: 18692},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 639, col: 22, offset: 18692},
						name: "Comments",
					},
					&oneOrMoreExpr{
						pos: position{line: 639, col: 33, offset: 18703},
						expr: &charClassMatcher{
							pos:        position{line: 639, col: 33, offset: 18703},
							val:        "[ \\r\\n\\t]",
							chars:      []rune{' ', '\r', '\n', '\t'},
							ignoreCase: false,
							inverted:   false,
						},
					},
				},
			},
		},
		{
			name:        "_",
			displayName: "\"whitespace\"",
			pos:         position{line: 641, col: 1, offset: 18715},
			expr: &choiceExpr{
				pos: position{line: 641, col: 18, offset: 18732},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 641, col: 18, offset: 18732},
						name: "Comments",
					},
					&zeroOrMoreExpr{
						pos: position{line: 641, col: 29, offset: 18743},
						expr: &charClassMatcher{
							pos:        position{line: 641, col: 29, offset: 18743},
							val:        "[ \\r\\n\\t]",
							chars:      []rune{' ', '\r', '\n', '\t'},
							ignoreCase: false,
							inverted:   false,
						},
					},
				},
			},
		},
		{
			name: "Comments",
			pos:  position{line: 643, col: 1, offset: 18755},
			expr: &oneOrMoreExpr{
				pos: position{line: 643, col: 12, offset: 18766},
				expr: &ruleRefExpr{
					pos:  position{line: 643, col: 13, offset: 18767},
					name: "Comment",
				},
			},
		},
		{
			name: "Comment",
			pos:  position{line: 645, col: 1, offset: 18778},
			expr: &seqExpr{
				pos: position{line: 645, col: 11, offset: 18788},
				exprs: []interface{}{
					&zeroOrMoreExpr{
						pos: position{line: 645, col: 11, offset: 18788},
						expr: &charClassMatcher{
							pos:        position{line: 645, col: 11, offset: 18788},
							val:        "[ \\r\\n\\t]",
							chars:      []rune{' ', '\r', '\n', '\t'},
							ignoreCase: false,
							inverted:   false,
						},
					},
					&litMatcher{
						pos:        position{line: 645, col: 22, offset: 18799},
						val:        "#",
						ignoreCase: false,
					},
					&zeroOrMoreExpr{
						pos: position{line: 645, col: 26, offset: 18803},
						expr: &seqExpr{
							pos: position{line: 645, col: 27, offset: 18804},
							exprs: []interface{}{
								&notExpr{
									pos: position{line: 645, col: 27, offset: 18804},
									expr: &charClassMatcher{
										pos:        position{line: 645, col: 28, offset: 18805},
										val:        "[\\n]",
										chars:      []rune{'\n'},
										ignoreCase: false,
										inverted:   false,
									},
								},
								&anyMatcher{
									line: 645, col: 33, offset: 18810,
								},
							},
						},
					},
					&andExpr{
						pos: position{line: 645, col: 37, offset: 18814},
						expr: &litMatcher{
							pos:        position{line: 645, col: 38, offset: 18815},
							val:        "\n",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "EOF",
			pos:  position{line: 647, col: 1, offset: 18821},
			expr: &notExpr{
				pos: position{line: 647, col: 7, offset: 18827},
				expr: &anyMatcher{
					line: 647, col: 8, offset: 18828,
				},
			},
		},
	},
}

func (c *current) onModule1(name, stat, rest interface{}) (interface{}, error) {
	//fmt.Println("beginning module")
	vals := rest.([]interface{})
	if len(vals) > 0 {
		//fmt.Println("multiple statements")
		subvalues := []Ast{stat.(Ast)}
		for _, el := range vals {
			subvalues = append(subvalues, el.(Ast))
		}
		return Module{Name: name.(Identifier).StringValue, Subvalues: subvalues}, nil
	} else {
		return Module{Name: name.(Identifier).StringValue, Subvalues: []Ast{stat.(Ast)}}, nil
	}
}

func (p *parser) callonModule1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onModule1(stack["name"], stack["stat"], stack["rest"])
}

func (c *current) onExprLine1(e interface{}) (interface{}, error) {
	fmt.Println("line:", e)
	// wrap calls as statements in an expr
	switch e.(type) {
	case Call:
		ex := Expr{Subvalues: []Ast{e.(Call)}, AsStatement: true}
		return ex, nil
	case Expr:
		ex := Expr{Subvalues: e.(Expr).Subvalues, AsStatement: true}
		return ex, nil
	}
	return e, nil
}

func (p *parser) callonExprLine1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onExprLine1(stack["e"])
}

func (c *current) onExternDefn1(name, importName, args, ret interface{}) (interface{}, error) {

	return Extern{Name: name.(Identifier).StringValue, Import: importName.(BasicAst).StringValue,
		Arguments: args.(Container).Subvalues, ReturnAnnotation: ret.(BasicAst).StringValue}, nil
}

func (p *parser) callonExternDefn1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onExternDefn1(stack["name"], stack["importName"], stack["args"], stack["ret"])
}

func (c *current) onTypeDefn2(name, params, types interface{}) (interface{}, error) {
	// alias type
	parameters := []Ast{}
	fields := []Ast{}

	vals := types.([]interface{})
	if len(vals) > 0 {
		restSl := toIfaceSlice(types)
		for _, v := range restSl {
			// we can get each item in the grammar by index
			restExpr := toIfaceSlice(v)
			v := restExpr[1].(Ast)
			fields = append(fields, v)
		}
	}

	return AliasType{Name: name.(Identifier).StringValue, Params: parameters, Types: fields}, nil
}

func (p *parser) callonTypeDefn2() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onTypeDefn2(stack["name"], stack["params"], stack["types"])
}

func (c *current) onTypeDefn22(name, params, first, rest interface{}) (interface{}, error) {
	// record type
	fields := []RecordField{first.(RecordField)}

	vals := rest.([]interface{})
	if len(vals) > 0 {
		restSl := toIfaceSlice(rest)
		for _, v := range restSl {
			// we can get each item in the grammar by index
			restExpr := toIfaceSlice(v)
			v := restExpr[2].(RecordField)
			fields = append(fields, v)
		}
	}

	return RecordType{Name: name.(Identifier).StringValue, Fields: fields}, nil
}

func (p *parser) callonTypeDefn22() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onTypeDefn22(stack["name"], stack["params"], stack["first"], stack["rest"])
}

func (c *current) onTypeDefn54(name, params, rest interface{}) (interface{}, error) {
	// variant type
	parameters := []Ast{}
	constructors := []VariantConstructor{}

	vals := rest.([]interface{})
	if len(vals) > 0 {
		for _, v := range vals {
			constructors = append(constructors, v.(VariantConstructor))
		}
	}

	return VariantType{Name: name.(Identifier).StringValue, Params: parameters, Constructors: constructors}, nil
}

func (p *parser) callonTypeDefn54() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onTypeDefn54(stack["name"], stack["params"], stack["rest"])
}

func (c *current) onRecordFieldDefn1(name, t interface{}) (interface{}, error) {
	return RecordField{Name: name.(Identifier).StringValue, Type: t.(Ast)}, nil
}

func (p *parser) callonRecordFieldDefn1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onRecordFieldDefn1(stack["name"], stack["t"])
}

func (c *current) onVariantConstructor2(name, first, rest interface{}) (interface{}, error) {
	// variant constructor with a record type
	fields := []RecordField{first.(RecordField)}

	vals := rest.([]interface{})
	if len(vals) > 0 {
		restSl := toIfaceSlice(rest)
		for _, v := range restSl {
			// we can get each item in the grammar by index
			restExpr := toIfaceSlice(v)
			v := restExpr[2].(RecordField)
			fields = append(fields, v)
		}
	}

	record := RecordType{Name: name.(Identifier).StringValue, Fields: fields}
	return VariantConstructor{Name: name.(Identifier).StringValue, Fields: []Ast{record}}, nil
}

func (p *parser) callonVariantConstructor2() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onVariantConstructor2(stack["name"], stack["first"], stack["rest"])
}

func (c *current) onVariantConstructor26(name, rest interface{}) (interface{}, error) {
	params := []Ast{}

	vals := rest.([]interface{})
	if len(vals) > 0 {
		restSl := toIfaceSlice(rest)
		for _, v := range restSl {
			// we can get each item in the grammar by index
			restExpr := toIfaceSlice(v)
			v := restExpr[1].(Ast)
			params = append(params, v)
		}
	}

	return VariantConstructor{Name: name.(Identifier).StringValue, Fields: params}, nil
}

func (p *parser) callonVariantConstructor26() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onVariantConstructor26(stack["name"], stack["rest"])
}

func (c *current) onAssignment2(i, expr interface{}) (interface{}, error) {
	//fmt.Println("assignment:", string(c.text))

	return Assignment{Left: i.(Ast), Right: expr.(Ast)}, nil
}

func (p *parser) callonAssignment2() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onAssignment2(stack["i"], stack["expr"])
}

func (c *current) onAssignment15() (interface{}, error) {
	return nil, errors.New("Variable name or '_' (unused result) required here")
}

func (p *parser) callonAssignment15() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onAssignment15()
}

func (c *current) onAssignment22(i interface{}) (interface{}, error) {
	return nil, errors.New("When assigning a value to a variable, you must use '='")
}

func (p *parser) callonAssignment22() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onAssignment22(stack["i"])
}

func (c *current) onFuncDefn1(i, ids, ret, statements interface{}) (interface{}, error) {
	//fmt.Println(string(c.text))
	subvalues := []Ast{}
	vals := statements.([]interface{})
	args := []Ast{}

	if len(vals) > 0 {
		for _, el := range vals {
			subvalues = append(subvalues, el.(Ast))
		}
	}

	if ids != nil {
		vals = ids.([]interface{})
		args = vals[0].(Container).Subvalues
	}

	retType := ""
	if ret != nil {
		vals = ret.([]interface{})
		retType = vals[0].(BasicAst).StringValue
	}

	return Func{Name: i.(Identifier).StringValue, Arguments: args, Subvalues: subvalues, ReturnAnnotation: retType}, nil
}

func (p *parser) callonFuncDefn1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onFuncDefn1(stack["i"], stack["ids"], stack["ret"], stack["statements"])
}

func (c *current) onExpr1(ex interface{}) (interface{}, error) {
	//fmt.Printf("top-level expr: %s\n", string(c.text))
	//fmt.Println(ex)
	return ex, nil
}

func (p *parser) callonExpr1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onExpr1(stack["ex"])
}

func (c *current) onIfExpr2(expr, thens, elses interface{}) (interface{}, error) {
	//fmt.Printf("if: %s\n", string(c.text))
	subvalues := []Ast{}
	vals := thens.([]interface{})
	if len(vals) > 0 {
		for _, el := range vals {
			subvalues = append(subvalues, el.(Ast))
		}
	}
	elsevalues := []Ast{}
	vals = elses.([]interface{})
	if len(vals) > 0 {
		for _, el := range vals {
			elsevalues = append(elsevalues, el.(Ast))
		}
	}
	return If{Condition: expr.(Ast), Then: subvalues, Else: elsevalues}, nil
}

func (p *parser) callonIfExpr2() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onIfExpr2(stack["expr"], stack["thens"], stack["elses"])
}

func (c *current) onIfExpr27(expr, thens, elseifs interface{}) (interface{}, error) {
	//fmt.Printf("if: %s\n", string(c.text))
	subvalues := []Ast{}
	vals := thens.([]interface{})
	if len(vals) > 0 {
		for _, el := range vals {
			subvalues = append(subvalues, el.(Ast))
		}
	}
	return If{Condition: expr.(Ast), Then: subvalues, Else: []Ast{elseifs.(Ast)}}, nil
}

func (p *parser) callonIfExpr27() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onIfExpr27(stack["expr"], stack["thens"], stack["elseifs"])
}

func (c *current) onIfExpr46(expr, thens interface{}) (interface{}, error) {
	//fmt.Printf("if: %s\n", string(c.text))
	subvalues := []Ast{}
	vals := thens.([]interface{})
	if len(vals) > 0 {
		for _, el := range vals {
			subvalues = append(subvalues, el.(Ast))
		}
	}
	return If{Condition: expr.(Ast), Then: subvalues}, nil
}

func (p *parser) callonIfExpr46() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onIfExpr46(stack["expr"], stack["thens"])
}

func (c *current) onCall2(module, fn, args interface{}) (interface{}, error) {
	//fmt.Println("call", string(c.text))
	arguments := []Ast{}

	switch args.(type) {
	case Container:
		arguments = args.(Container).Subvalues
	default:
		arguments = []Ast{}
	}

	return Call{Module: module.(Identifier), Function: fn.(Identifier), Arguments: arguments}, nil
}

func (p *parser) callonCall2() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onCall2(stack["module"], stack["fn"], stack["args"])
}

func (c *current) onCall13(fn, args interface{}) (interface{}, error) {
	//fmt.Println("call", string(c.text))
	arguments := []Ast{}

	switch args.(type) {
	case Container:
		arguments = args.(Container).Subvalues
	default:
		arguments = []Ast{}
	}

	return Call{Module: Identifier{}, Function: fn.(Identifier), Arguments: arguments}, nil
}

func (p *parser) callonCall13() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onCall13(stack["fn"], stack["args"])
}

func (c *current) onCall21(fn interface{}) (interface{}, error) {
	//fmt.Println("call", string(c.text))
	arguments := []Ast{}

	return Call{Module: Identifier{}, Function: fn.(Identifier), Arguments: arguments}, nil
}

func (p *parser) callonCall21() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onCall21(stack["fn"])
}

func (c *current) onArgsDefn1(argument, rest interface{}) (interface{}, error) {

	args := []Ast{argument.(Ast)}

	vals := rest.([]interface{})
	if len(vals) > 0 {
		restSl := toIfaceSlice(rest)
		for _, v := range restSl {
			// we can get each item in the grammar by index
			restExpr := toIfaceSlice(v)
			v := restExpr[2].(Ast)
			args = append(args, v)
		}
	}

	return Container{Type: "Arguments", Subvalues: args}, nil
}

func (p *parser) callonArgsDefn1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onArgsDefn1(stack["argument"], stack["rest"])
}

func (c *current) onArgDefn1(name, anno interface{}) (interface{}, error) {
	//fmt.Println("parsing arg:", string(c.text))
	arg := name.(Identifier)

	if anno != nil {
		vals := anno.([]interface{})
		//fmt.Println(vals)
		//restSl := toIfaceSlice(vals[0])

		switch vals[2].(type) {
		case BasicAst:
			arg.Annotation = vals[2].(BasicAst).StringValue
		case Identifier:
			arg.Annotation = vals[2].(Identifier).StringValue
		}
	}
	//fmt.Println("parsed:", arg)
	return arg, nil
}

func (p *parser) callonArgDefn1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onArgDefn1(stack["name"], stack["anno"])
}

func (c *current) onArguments2(argument, rest interface{}) (interface{}, error) {
	args := []Ast{argument.(Ast)}

	vals := rest.([]interface{})
	if len(vals) > 0 {
		restSl := toIfaceSlice(rest)
		for _, v := range restSl {
			// we can get each item in the grammar by index
			restExpr := toIfaceSlice(v)
			v := restExpr[2].(Ast)
			args = append(args, v)
		}
	}

	return Container{Type: "Arguments", Subvalues: args}, nil
}

func (p *parser) callonArguments2() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onArguments2(stack["argument"], stack["rest"])
}

func (c *current) onArguments17(argument interface{}) (interface{}, error) {
	args := []Ast{argument.(Ast)}
	return Container{Type: "Arguments", Subvalues: args}, nil
}

func (p *parser) callonArguments17() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onArguments17(stack["argument"])
}

func (c *current) onCompoundExpr1(op, rest interface{}) (interface{}, error) {
	//fmt.Println("compound", op, rest)
	vals := rest.([]interface{})
	if len(vals) > 0 {
		subvalues := []Ast{op.(Ast)}
		restSl := toIfaceSlice(rest)
		for _, v := range restSl {
			// we can get each item in the grammar by index
			restExpr := toIfaceSlice(v)
			v := restExpr[3].(Ast)
			op := restExpr[1].(Ast)
			subvalues = append(subvalues, op, v)
		}

		return Expr{Type: "Compound", Subvalues: subvalues}, nil
	} else {
		return Expr{Type: "Compound", Subvalues: []Ast{op.(Ast)}}, nil
	}
}

func (p *parser) callonCompoundExpr1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onCompoundExpr1(stack["op"], stack["rest"])
}

func (c *current) onBinOpBool1(first, rest interface{}) (interface{}, error) {

	subvalues := []Ast{first.(Ast)}

	//fmt.Println("binopbool", first, rest)
	vals := rest.([]interface{})
	if len(vals) > 0 {

		restSl := toIfaceSlice(rest)
		for _, v := range restSl {
			// we can get each item in the grammar by index
			restExpr := toIfaceSlice(v)
			e := restExpr[3].(Ast)
			op := restExpr[1].(Ast)
			subvalues = append(subvalues, op, e)
		}
		//return Expr{Type: "BinOpBool", Subvalues: subvalues }, nil
	}

	for len(subvalues) > 1 {
		length := len(subvalues)
		right := subvalues[length-1].(Ast)
		op := subvalues[length-2].(Operator)
		left := subvalues[length-3].(Ast)
		binop := BinOp{Operator: op, Left: left, Right: right}
		subvalues = append(subvalues[:length-3], binop)

	}

	return subvalues[0].(Ast), nil

}

func (p *parser) callonBinOpBool1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onBinOpBool1(stack["first"], stack["rest"])
}

func (c *current) onBinOpEquality1(first, rest interface{}) (interface{}, error) {
	subvalues := []Ast{first.(Ast)}

	//fmt.Println("binopbool", first, rest)
	vals := rest.([]interface{})
	if len(vals) > 0 {

		restSl := toIfaceSlice(rest)
		for _, v := range restSl {
			// we can get each item in the grammar by index
			restExpr := toIfaceSlice(v)
			e := restExpr[3].(Ast)
			op := restExpr[1].(Ast)
			subvalues = append(subvalues, op, e)
		}
		//return Expr{Type: "BinOpBool", Subvalues: subvalues }, nil
	}

	for len(subvalues) > 1 {
		length := len(subvalues)
		right := subvalues[length-1].(Ast)
		op := subvalues[length-2].(Operator)
		left := subvalues[length-3].(Ast)
		binop := BinOp{Operator: op, Left: left, Right: right}
		subvalues = append(subvalues[:length-3], binop)

	}

	return subvalues[0].(Ast), nil

}

func (p *parser) callonBinOpEquality1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onBinOpEquality1(stack["first"], stack["rest"])
}

func (c *current) onBinOpLow1(first, rest interface{}) (interface{}, error) {
	subvalues := []Ast{first.(Ast)}

	//fmt.Println("binopbool", first, rest)
	vals := rest.([]interface{})
	if len(vals) > 0 {

		restSl := toIfaceSlice(rest)
		for _, v := range restSl {
			// we can get each item in the grammar by index
			restExpr := toIfaceSlice(v)
			e := restExpr[3].(Ast)
			op := restExpr[1].(Ast)
			subvalues = append(subvalues, op, e)
		}
		//return Expr{Type: "BinOpBool", Subvalues: subvalues }, nil
	}

	for len(subvalues) > 1 {
		length := len(subvalues)
		right := subvalues[length-1].(Ast)
		op := subvalues[length-2].(Operator)
		left := subvalues[length-3].(Ast)
		binop := BinOp{Operator: op, Left: left, Right: right}
		subvalues = append(subvalues[:length-3], binop)

	}

	return subvalues[0].(Ast), nil

}

func (p *parser) callonBinOpLow1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onBinOpLow1(stack["first"], stack["rest"])
}

func (c *current) onBinOpHigh1(first, rest interface{}) (interface{}, error) {
	subvalues := []Ast{first.(Ast)}

	//fmt.Println("binopbool", first, rest)
	vals := rest.([]interface{})
	if len(vals) > 0 {

		restSl := toIfaceSlice(rest)
		for _, v := range restSl {
			// we can get each item in the grammar by index
			restExpr := toIfaceSlice(v)
			e := restExpr[3].(Ast)
			op := restExpr[1].(Ast)
			subvalues = append(subvalues, op, e)
		}
		//return Expr{Type: "BinOpBool", Subvalues: subvalues }, nil
	}

	for len(subvalues) > 1 {
		length := len(subvalues)
		right := subvalues[length-1].(Ast)
		op := subvalues[length-2].(Operator)
		left := subvalues[length-3].(Ast)
		binop := BinOp{Operator: op, Left: left, Right: right}
		subvalues = append(subvalues[:length-3], binop)

	}

	return subvalues[0].(Ast), nil
}

func (p *parser) callonBinOpHigh1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onBinOpHigh1(stack["first"], stack["rest"])
}

func (c *current) onBinOpParens2(first interface{}) (interface{}, error) {
	//fmt.Println("binopparens", first)
	return Expr{Type: "BinOpParens", Subvalues: []Ast{first.(Ast)}}, nil
}

func (p *parser) callonBinOpParens2() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onBinOpParens2(stack["first"])
}

func (c *current) onOperatorBoolean1() (interface{}, error) {
	return Operator{StringValue: string(c.text), ValueType: BOOL}, nil
}

func (p *parser) callonOperatorBoolean1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onOperatorBoolean1()
}

func (c *current) onOperatorEquality1() (interface{}, error) {
	return Operator{StringValue: string(c.text), ValueType: BOOL}, nil
}

func (p *parser) callonOperatorEquality1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onOperatorEquality1()
}

func (c *current) onOperatorHigh2() (interface{}, error) {
	return Operator{StringValue: string(c.text), ValueType: NUMBER}, nil
}

func (p *parser) callonOperatorHigh2() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onOperatorHigh2()
}

func (c *current) onOperatorHigh6() (interface{}, error) {
	return Operator{StringValue: string(c.text), ValueType: INT}, nil
}

func (p *parser) callonOperatorHigh6() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onOperatorHigh6()
}

func (c *current) onOperatorHigh8() (interface{}, error) {
	return Operator{StringValue: string(c.text), ValueType: STRING}, nil
}

func (p *parser) callonOperatorHigh8() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onOperatorHigh8()
}

func (c *current) onOperatorLow1() (interface{}, error) {
	return Operator{StringValue: string(c.text), ValueType: NUMBER}, nil
}

func (p *parser) callonOperatorLow1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onOperatorLow1()
}

func (c *current) onArrayLiteral1(first, rest interface{}) (interface{}, error) {
	// rest:(_ ',' _ Expr)*
	vals := rest.([]interface{})
	subvalues := []Ast{first.(Ast)}
	if len(vals) > 0 {
		restSl := toIfaceSlice(rest)
		for _, v := range restSl {
			// we can get each item in the grammar by index
			restExpr := toIfaceSlice(v)
			v := restExpr[2].(Ast)
			subvalues = append(subvalues, v)
		}
	}
	return ArrayType{Subvalues: subvalues}, nil
}

func (p *parser) callonArrayLiteral1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onArrayLiteral1(stack["first"], stack["rest"])
}

func (c *current) onVariantInstance1(name, args interface{}) (interface{}, error) {
	arguments := []Ast{}

	switch args.(type) {
	case Container:
		arguments = args.(Container).Subvalues
	default:
		arguments = []Ast{}
	}

	return VariantInstance{Name: name.(BasicAst).StringValue, Arguments: arguments}, nil
}

func (p *parser) callonVariantInstance1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onVariantInstance1(stack["name"], stack["args"])
}

func (c *current) onRecordInstance1(name, firstName, firstValue, rest interface{}) (interface{}, error) {
	instance := RecordInstance{Name: name.(BasicAst).StringValue}
	instance.Values = make(map[string]Ast)

	vals := rest.([]interface{})
	instance.Values[firstName.(BasicAst).StringValue] = firstValue.(Ast)

	if len(vals) > 0 {
		restSl := toIfaceSlice(rest)
		for _, v := range restSl {
			// we can get each item in the grammar by index
			restExpr := toIfaceSlice(v)
			k := restExpr[2].(BasicAst).StringValue
			v := restExpr[6].(Ast)
			instance.Values[k] = v
		}
	}
	return instance, nil
}

func (p *parser) callonRecordInstance1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onRecordInstance1(stack["name"], stack["firstName"], stack["firstValue"], stack["rest"])
}

func (c *current) onTypeParameter1() (interface{}, error) {
	return Identifier{StringValue: string(c.text)}, nil
}

func (p *parser) callonTypeParameter1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onTypeParameter1()
}

func (c *current) onVariableName1() (interface{}, error) {
	return Identifier{StringValue: string(c.text)}, nil
}

func (p *parser) callonVariableName1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onVariableName1()
}

func (c *current) onModuleName1() (interface{}, error) {
	return Identifier{StringValue: string(c.text)}, nil
}

func (p *parser) callonModuleName1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onModuleName1()
}

func (c *current) onConst2() (interface{}, error) {
	val, err := strconv.Atoi(string(c.text))
	if err != nil {
		return nil, err
	}
	return BasicAst{Type: "Integer", IntValue: val, ValueType: INT}, nil
}

func (p *parser) callonConst2() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onConst2()
}

func (c *current) onConst10() (interface{}, error) {
	val, err := strconv.ParseFloat(string(c.text), 64)
	if err != nil {
		return nil, err
	}
	return BasicAst{Type: "Float", FloatValue: val, ValueType: FLOAT}, nil
}

func (p *parser) callonConst10() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onConst10()
}

func (c *current) onConst19() (interface{}, error) {
	return BasicAst{Type: "Bool", BoolValue: true, ValueType: BOOL}, nil

}

func (p *parser) callonConst19() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onConst19()
}

func (c *current) onConst21() (interface{}, error) {
	return BasicAst{Type: "Bool", BoolValue: false, ValueType: BOOL}, nil
}

func (p *parser) callonConst21() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onConst21()
}

func (c *current) onConst25(val interface{}) (interface{}, error) {
	return BasicAst{Type: "Char", CharValue: rune(c.text[1]), ValueType: CHAR}, nil
}

func (p *parser) callonConst25() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onConst25(stack["val"])
}

func (c *current) onUnit1() (interface{}, error) {
	return BasicAst{Type: "Unit", ValueType: NIL}, nil
}

func (p *parser) callonUnit1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onUnit1()
}

func (c *current) onStringLiteral1() (interface{}, error) {
	val, err := strconv.Unquote(string(c.text))
	if err == nil {
		return BasicAst{Type: "String", StringValue: val, ValueType: STRING}, nil
	}
	return nil, err
}

func (p *parser) callonStringLiteral1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onStringLiteral1()
}

func (c *current) onUnused1() (interface{}, error) {
	return Identifier{StringValue: "_"}, nil
}

func (p *parser) callonUnused1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onUnused1()
}

func (c *current) onBaseType2() (interface{}, error) {
	return BasicAst{Type: "Type", StringValue: string(c.text), ValueType: STRING}, nil
}

func (p *parser) callonBaseType2() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onBaseType2()
}

var (
	// errNoRule is returned when the grammar to parse has no rule.
	errNoRule = errors.New("grammar has no rule")

	// errInvalidEncoding is returned when the source is not properly
	// utf8-encoded.
	errInvalidEncoding = errors.New("invalid encoding")
)

// Option is a function that can set an option on the parser. It returns
// the previous setting as an Option.
type Option func(*parser) Option

// Debug creates an Option to set the debug flag to b. When set to true,
// debugging information is printed to stdout while parsing.
//
// The default is false.
func Debug(b bool) Option {
	return func(p *parser) Option {
		old := p.debug
		p.debug = b
		return Debug(old)
	}
}

// Memoize creates an Option to set the memoize flag to b. When set to true,
// the parser will cache all results so each expression is evaluated only
// once. This guarantees linear parsing time even for pathological cases,
// at the expense of more memory and slower times for typical cases.
//
// The default is false.
func Memoize(b bool) Option {
	return func(p *parser) Option {
		old := p.memoize
		p.memoize = b
		return Memoize(old)
	}
}

// Recover creates an Option to set the recover flag to b. When set to
// true, this causes the parser to recover from panics and convert it
// to an error. Setting it to false can be useful while debugging to
// access the full stack trace.
//
// The default is true.
func Recover(b bool) Option {
	return func(p *parser) Option {
		old := p.recover
		p.recover = b
		return Recover(old)
	}
}

// ParseFile parses the file identified by filename.
func ParseFile(filename string, opts ...Option) (i interface{}, err error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer func() {
		err = f.Close()
	}()
	return ParseReader(filename, f, opts...)
}

// ParseReader parses the data from r using filename as information in the
// error messages.
func ParseReader(filename string, r io.Reader, opts ...Option) (interface{}, error) {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	return Parse(filename, b, opts...)
}

// Parse parses the data from b using filename as information in the
// error messages.
func Parse(filename string, b []byte, opts ...Option) (interface{}, error) {
	return newParser(filename, b, opts...).parse(g)
}

// position records a position in the text.
type position struct {
	line, col, offset int
}

func (p position) String() string {
	return fmt.Sprintf("%d:%d [%d]", p.line, p.col, p.offset)
}

// savepoint stores all state required to go back to this point in the
// parser.
type savepoint struct {
	position
	rn rune
	w  int
}

type current struct {
	pos  position // start position of the match
	text []byte   // raw text of the match
}

// the AST types...

type grammar struct {
	pos   position
	rules []*rule
}

type rule struct {
	pos         position
	name        string
	displayName string
	expr        interface{}
}

type choiceExpr struct {
	pos          position
	alternatives []interface{}
}

type actionExpr struct {
	pos  position
	expr interface{}
	run  func(*parser) (interface{}, error)
}

type seqExpr struct {
	pos   position
	exprs []interface{}
}

type labeledExpr struct {
	pos   position
	label string
	expr  interface{}
}

type expr struct {
	pos  position
	expr interface{}
}

type andExpr expr
type notExpr expr
type zeroOrOneExpr expr
type zeroOrMoreExpr expr
type oneOrMoreExpr expr

type ruleRefExpr struct {
	pos  position
	name string
}

type andCodeExpr struct {
	pos position
	run func(*parser) (bool, error)
}

type notCodeExpr struct {
	pos position
	run func(*parser) (bool, error)
}

type litMatcher struct {
	pos        position
	val        string
	ignoreCase bool
}

type charClassMatcher struct {
	pos        position
	val        string
	chars      []rune
	ranges     []rune
	classes    []*unicode.RangeTable
	ignoreCase bool
	inverted   bool
}

type anyMatcher position

// errList cumulates the errors found by the parser.
type errList []error

func (e *errList) add(err error) {
	*e = append(*e, err)
}

func (e errList) err() error {
	if len(e) == 0 {
		return nil
	}
	e.dedupe()
	return e
}

func (e *errList) dedupe() {
	var cleaned []error
	set := make(map[string]bool)
	for _, err := range *e {
		if msg := err.Error(); !set[msg] {
			set[msg] = true
			cleaned = append(cleaned, err)
		}
	}
	*e = cleaned
}

func (e errList) Error() string {
	switch len(e) {
	case 0:
		return ""
	case 1:
		return e[0].Error()
	default:
		var buf bytes.Buffer

		for i, err := range e {
			if i > 0 {
				buf.WriteRune('\n')
			}
			buf.WriteString(err.Error())
		}
		return buf.String()
	}
}

// parserError wraps an error with a prefix indicating the rule in which
// the error occurred. The original error is stored in the Inner field.
type parserError struct {
	Inner    error
	pos      position
	prefix   string
	expected []string
}

// Error returns the error message.
func (p *parserError) Error() string {
	return p.prefix + ": " + p.Inner.Error()
}

// newParser creates a parser with the specified input source and options.
func newParser(filename string, b []byte, opts ...Option) *parser {
	p := &parser{
		filename:        filename,
		errs:            new(errList),
		data:            b,
		pt:              savepoint{position: position{line: 1}},
		recover:         true,
		maxFailPos:      position{col: 1, line: 1},
		maxFailExpected: make(map[string]struct{}),
	}
	p.setOptions(opts)
	return p
}

// setOptions applies the options to the parser.
func (p *parser) setOptions(opts []Option) {
	for _, opt := range opts {
		opt(p)
	}
}

type resultTuple struct {
	v   interface{}
	b   bool
	end savepoint
}

type parser struct {
	filename string
	pt       savepoint
	cur      current

	data []byte
	errs *errList

	depth   int
	recover bool
	debug   bool

	memoize bool
	// memoization table for the packrat algorithm:
	// map[offset in source] map[expression or rule] {value, match}
	memo map[int]map[interface{}]resultTuple

	// rules table, maps the rule identifier to the rule node
	rules map[string]*rule
	// variables stack, map of label to value
	vstack []map[string]interface{}
	// rule stack, allows identification of the current rule in errors
	rstack []*rule

	// stats
	exprCnt int

	// parse fail
	maxFailPos            position
	maxFailExpected       map[string]struct{}
	maxFailInvertExpected bool
}

// push a variable set on the vstack.
func (p *parser) pushV() {
	if cap(p.vstack) == len(p.vstack) {
		// create new empty slot in the stack
		p.vstack = append(p.vstack, nil)
	} else {
		// slice to 1 more
		p.vstack = p.vstack[:len(p.vstack)+1]
	}

	// get the last args set
	m := p.vstack[len(p.vstack)-1]
	if m != nil && len(m) == 0 {
		// empty map, all good
		return
	}

	m = make(map[string]interface{})
	p.vstack[len(p.vstack)-1] = m
}

// pop a variable set from the vstack.
func (p *parser) popV() {
	// if the map is not empty, clear it
	m := p.vstack[len(p.vstack)-1]
	if len(m) > 0 {
		// GC that map
		p.vstack[len(p.vstack)-1] = nil
	}
	p.vstack = p.vstack[:len(p.vstack)-1]
}

func (p *parser) print(prefix, s string) string {
	if !p.debug {
		return s
	}

	fmt.Printf("%s %d:%d:%d: %s [%#U]\n",
		prefix, p.pt.line, p.pt.col, p.pt.offset, s, p.pt.rn)
	return s
}

func (p *parser) in(s string) string {
	p.depth++
	return p.print(strings.Repeat(" ", p.depth)+">", s)
}

func (p *parser) out(s string) string {
	p.depth--
	return p.print(strings.Repeat(" ", p.depth)+"<", s)
}

func (p *parser) addErr(err error) {
	p.addErrAt(err, p.pt.position, []string{})
}

func (p *parser) addErrAt(err error, pos position, expected []string) {
	var buf bytes.Buffer
	if p.filename != "" {
		buf.WriteString(p.filename)
	}
	if buf.Len() > 0 {
		buf.WriteString(":")
	}
	buf.WriteString(fmt.Sprintf("%d:%d (%d)", pos.line, pos.col, pos.offset))
	if len(p.rstack) > 0 {
		if buf.Len() > 0 {
			buf.WriteString(": ")
		}
		rule := p.rstack[len(p.rstack)-1]
		if rule.displayName != "" {
			buf.WriteString("rule " + rule.displayName)
		} else {
			buf.WriteString("rule " + rule.name)
		}
	}
	pe := &parserError{Inner: err, pos: pos, prefix: buf.String(), expected: expected}
	p.errs.add(pe)
}

func (p *parser) failAt(fail bool, pos position, want string) {
	// process fail if parsing fails and not inverted or parsing succeeds and invert is set
	if fail == p.maxFailInvertExpected {
		if pos.offset < p.maxFailPos.offset {
			return
		}

		if pos.offset > p.maxFailPos.offset {
			p.maxFailPos = pos
			p.maxFailExpected = make(map[string]struct{})
		}

		if p.maxFailInvertExpected {
			want = "!" + want
		}
		p.maxFailExpected[want] = struct{}{}
	}
}

// read advances the parser to the next rune.
func (p *parser) read() {
	p.pt.offset += p.pt.w
	rn, n := utf8.DecodeRune(p.data[p.pt.offset:])
	p.pt.rn = rn
	p.pt.w = n
	p.pt.col++
	if rn == '\n' {
		p.pt.line++
		p.pt.col = 0
	}

	if rn == utf8.RuneError {
		if n == 1 {
			p.addErr(errInvalidEncoding)
		}
	}
}

// restore parser position to the savepoint pt.
func (p *parser) restore(pt savepoint) {
	if p.debug {
		defer p.out(p.in("restore"))
	}
	if pt.offset == p.pt.offset {
		return
	}
	p.pt = pt
}

// get the slice of bytes from the savepoint start to the current position.
func (p *parser) sliceFrom(start savepoint) []byte {
	return p.data[start.position.offset:p.pt.position.offset]
}

func (p *parser) getMemoized(node interface{}) (resultTuple, bool) {
	if len(p.memo) == 0 {
		return resultTuple{}, false
	}
	m := p.memo[p.pt.offset]
	if len(m) == 0 {
		return resultTuple{}, false
	}
	res, ok := m[node]
	return res, ok
}

func (p *parser) setMemoized(pt savepoint, node interface{}, tuple resultTuple) {
	if p.memo == nil {
		p.memo = make(map[int]map[interface{}]resultTuple)
	}
	m := p.memo[pt.offset]
	if m == nil {
		m = make(map[interface{}]resultTuple)
		p.memo[pt.offset] = m
	}
	m[node] = tuple
}

func (p *parser) buildRulesTable(g *grammar) {
	p.rules = make(map[string]*rule, len(g.rules))
	for _, r := range g.rules {
		p.rules[r.name] = r
	}
}

func (p *parser) parse(g *grammar) (val interface{}, err error) {
	if len(g.rules) == 0 {
		p.addErr(errNoRule)
		return nil, p.errs.err()
	}

	// TODO : not super critical but this could be generated
	p.buildRulesTable(g)

	if p.recover {
		// panic can be used in action code to stop parsing immediately
		// and return the panic as an error.
		defer func() {
			if e := recover(); e != nil {
				if p.debug {
					defer p.out(p.in("panic handler"))
				}
				val = nil
				switch e := e.(type) {
				case error:
					p.addErr(e)
				default:
					p.addErr(fmt.Errorf("%v", e))
				}
				err = p.errs.err()
			}
		}()
	}

	// start rule is rule [0]
	p.read() // advance to first rune
	val, ok := p.parseRule(g.rules[0])
	if !ok {
		if len(*p.errs) == 0 {
			// If parsing fails, but no errors have been recorded, the expected values
			// for the farthest parser position are returned as error.
			expected := make([]string, 0, len(p.maxFailExpected))
			eof := false
			if _, ok := p.maxFailExpected["!."]; ok {
				delete(p.maxFailExpected, "!.")
				eof = true
			}
			for k := range p.maxFailExpected {
				expected = append(expected, k)
			}
			sort.Strings(expected)
			if eof {
				expected = append(expected, "EOF")
			}
			p.addErrAt(errors.New("no match found, expected: "+listJoin(expected, ", ", "or")), p.maxFailPos, expected)
		}
		return nil, p.errs.err()
	}
	return val, p.errs.err()
}

func listJoin(list []string, sep string, lastSep string) string {
	switch len(list) {
	case 0:
		return ""
	case 1:
		return list[0]
	default:
		return fmt.Sprintf("%s %s %s", strings.Join(list[:len(list)-1], sep), lastSep, list[len(list)-1])
	}
}

func (p *parser) parseRule(rule *rule) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseRule " + rule.name))
	}

	if p.memoize {
		res, ok := p.getMemoized(rule)
		if ok {
			p.restore(res.end)
			return res.v, res.b
		}
	}

	start := p.pt
	p.rstack = append(p.rstack, rule)
	p.pushV()
	val, ok := p.parseExpr(rule.expr)
	p.popV()
	p.rstack = p.rstack[:len(p.rstack)-1]
	if ok && p.debug {
		p.print(strings.Repeat(" ", p.depth)+"MATCH", string(p.sliceFrom(start)))
	}

	if p.memoize {
		p.setMemoized(start, rule, resultTuple{val, ok, p.pt})
	}
	return val, ok
}

func (p *parser) parseExpr(expr interface{}) (interface{}, bool) {
	var pt savepoint

	if p.memoize {
		res, ok := p.getMemoized(expr)
		if ok {
			p.restore(res.end)
			return res.v, res.b
		}
		pt = p.pt
	}

	p.exprCnt++
	var val interface{}
	var ok bool
	switch expr := expr.(type) {
	case *actionExpr:
		val, ok = p.parseActionExpr(expr)
	case *andCodeExpr:
		val, ok = p.parseAndCodeExpr(expr)
	case *andExpr:
		val, ok = p.parseAndExpr(expr)
	case *anyMatcher:
		val, ok = p.parseAnyMatcher(expr)
	case *charClassMatcher:
		val, ok = p.parseCharClassMatcher(expr)
	case *choiceExpr:
		val, ok = p.parseChoiceExpr(expr)
	case *labeledExpr:
		val, ok = p.parseLabeledExpr(expr)
	case *litMatcher:
		val, ok = p.parseLitMatcher(expr)
	case *notCodeExpr:
		val, ok = p.parseNotCodeExpr(expr)
	case *notExpr:
		val, ok = p.parseNotExpr(expr)
	case *oneOrMoreExpr:
		val, ok = p.parseOneOrMoreExpr(expr)
	case *ruleRefExpr:
		val, ok = p.parseRuleRefExpr(expr)
	case *seqExpr:
		val, ok = p.parseSeqExpr(expr)
	case *zeroOrMoreExpr:
		val, ok = p.parseZeroOrMoreExpr(expr)
	case *zeroOrOneExpr:
		val, ok = p.parseZeroOrOneExpr(expr)
	default:
		panic(fmt.Sprintf("unknown expression type %T", expr))
	}
	if p.memoize {
		p.setMemoized(pt, expr, resultTuple{val, ok, p.pt})
	}
	return val, ok
}

func (p *parser) parseActionExpr(act *actionExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseActionExpr"))
	}

	start := p.pt
	val, ok := p.parseExpr(act.expr)
	if ok {
		p.cur.pos = start.position
		p.cur.text = p.sliceFrom(start)
		actVal, err := act.run(p)
		if err != nil {
			p.addErrAt(err, start.position, []string{})
		}
		val = actVal
	}
	if ok && p.debug {
		p.print(strings.Repeat(" ", p.depth)+"MATCH", string(p.sliceFrom(start)))
	}
	return val, ok
}

func (p *parser) parseAndCodeExpr(and *andCodeExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseAndCodeExpr"))
	}

	ok, err := and.run(p)
	if err != nil {
		p.addErr(err)
	}
	return nil, ok
}

func (p *parser) parseAndExpr(and *andExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseAndExpr"))
	}

	pt := p.pt
	p.pushV()
	_, ok := p.parseExpr(and.expr)
	p.popV()
	p.restore(pt)
	return nil, ok
}

func (p *parser) parseAnyMatcher(any *anyMatcher) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseAnyMatcher"))
	}

	if p.pt.rn != utf8.RuneError {
		start := p.pt
		p.read()
		p.failAt(true, start.position, ".")
		return p.sliceFrom(start), true
	}
	p.failAt(false, p.pt.position, ".")
	return nil, false
}

func (p *parser) parseCharClassMatcher(chr *charClassMatcher) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseCharClassMatcher"))
	}

	cur := p.pt.rn
	start := p.pt
	// can't match EOF
	if cur == utf8.RuneError {
		p.failAt(false, start.position, chr.val)
		return nil, false
	}
	if chr.ignoreCase {
		cur = unicode.ToLower(cur)
	}

	// try to match in the list of available chars
	for _, rn := range chr.chars {
		if rn == cur {
			if chr.inverted {
				p.failAt(false, start.position, chr.val)
				return nil, false
			}
			p.read()
			p.failAt(true, start.position, chr.val)
			return p.sliceFrom(start), true
		}
	}

	// try to match in the list of ranges
	for i := 0; i < len(chr.ranges); i += 2 {
		if cur >= chr.ranges[i] && cur <= chr.ranges[i+1] {
			if chr.inverted {
				p.failAt(false, start.position, chr.val)
				return nil, false
			}
			p.read()
			p.failAt(true, start.position, chr.val)
			return p.sliceFrom(start), true
		}
	}

	// try to match in the list of Unicode classes
	for _, cl := range chr.classes {
		if unicode.Is(cl, cur) {
			if chr.inverted {
				p.failAt(false, start.position, chr.val)
				return nil, false
			}
			p.read()
			p.failAt(true, start.position, chr.val)
			return p.sliceFrom(start), true
		}
	}

	if chr.inverted {
		p.read()
		p.failAt(true, start.position, chr.val)
		return p.sliceFrom(start), true
	}
	p.failAt(false, start.position, chr.val)
	return nil, false
}

func (p *parser) parseChoiceExpr(ch *choiceExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseChoiceExpr"))
	}

	for _, alt := range ch.alternatives {
		p.pushV()
		val, ok := p.parseExpr(alt)
		p.popV()
		if ok {
			return val, ok
		}
	}
	return nil, false
}

func (p *parser) parseLabeledExpr(lab *labeledExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseLabeledExpr"))
	}

	p.pushV()
	val, ok := p.parseExpr(lab.expr)
	p.popV()
	if ok && lab.label != "" {
		m := p.vstack[len(p.vstack)-1]
		m[lab.label] = val
	}
	return val, ok
}

func (p *parser) parseLitMatcher(lit *litMatcher) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseLitMatcher"))
	}

	ignoreCase := ""
	if lit.ignoreCase {
		ignoreCase = "i"
	}
	val := fmt.Sprintf("%q%s", lit.val, ignoreCase)
	start := p.pt
	for _, want := range lit.val {
		cur := p.pt.rn
		if lit.ignoreCase {
			cur = unicode.ToLower(cur)
		}
		if cur != want {
			p.failAt(false, start.position, val)
			p.restore(start)
			return nil, false
		}
		p.read()
	}
	p.failAt(true, start.position, val)
	return p.sliceFrom(start), true
}

func (p *parser) parseNotCodeExpr(not *notCodeExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseNotCodeExpr"))
	}

	ok, err := not.run(p)
	if err != nil {
		p.addErr(err)
	}
	return nil, !ok
}

func (p *parser) parseNotExpr(not *notExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseNotExpr"))
	}

	pt := p.pt
	p.pushV()
	p.maxFailInvertExpected = !p.maxFailInvertExpected
	_, ok := p.parseExpr(not.expr)
	p.maxFailInvertExpected = !p.maxFailInvertExpected
	p.popV()
	p.restore(pt)
	return nil, !ok
}

func (p *parser) parseOneOrMoreExpr(expr *oneOrMoreExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseOneOrMoreExpr"))
	}

	var vals []interface{}

	for {
		p.pushV()
		val, ok := p.parseExpr(expr.expr)
		p.popV()
		if !ok {
			if len(vals) == 0 {
				// did not match once, no match
				return nil, false
			}
			return vals, true
		}
		vals = append(vals, val)
	}
}

func (p *parser) parseRuleRefExpr(ref *ruleRefExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseRuleRefExpr " + ref.name))
	}

	if ref.name == "" {
		panic(fmt.Sprintf("%s: invalid rule: missing name", ref.pos))
	}

	rule := p.rules[ref.name]
	if rule == nil {
		p.addErr(fmt.Errorf("undefined rule: %s", ref.name))
		return nil, false
	}
	return p.parseRule(rule)
}

func (p *parser) parseSeqExpr(seq *seqExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseSeqExpr"))
	}

	var vals []interface{}

	pt := p.pt
	for _, expr := range seq.exprs {
		val, ok := p.parseExpr(expr)
		if !ok {
			p.restore(pt)
			return nil, false
		}
		vals = append(vals, val)
	}
	return vals, true
}

func (p *parser) parseZeroOrMoreExpr(expr *zeroOrMoreExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseZeroOrMoreExpr"))
	}

	var vals []interface{}

	for {
		p.pushV()
		val, ok := p.parseExpr(expr.expr)
		p.popV()
		if !ok {
			return vals, true
		}
		vals = append(vals, val)
	}
}

func (p *parser) parseZeroOrOneExpr(expr *zeroOrOneExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseZeroOrOneExpr"))
	}

	p.pushV()
	val, _ := p.parseExpr(expr.expr)
	p.popV()
	// whether it matched or not, consider it a match
	return val, true
}
