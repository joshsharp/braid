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
						&labeledExpr{
							pos:   position{line: 13, col: 12, offset: 153},
							label: "stat",
							expr: &ruleRefExpr{
								pos:  position{line: 13, col: 17, offset: 158},
								name: "TopLevelStatement",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 13, col: 35, offset: 176},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 13, col: 37, offset: 178},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 13, col: 42, offset: 183},
								expr: &ruleRefExpr{
									pos:  position{line: 13, col: 43, offset: 184},
									name: "TopLevelStatement",
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 13, col: 63, offset: 204},
							name: "_",
						},
						&ruleRefExpr{
							pos:  position{line: 13, col: 65, offset: 206},
							name: "EOF",
						},
					},
				},
			},
		},
		{
			name: "TopLevelStatement",
			pos:  position{line: 28, col: 1, offset: 633},
			expr: &choiceExpr{
				pos: position{line: 28, col: 21, offset: 653},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 28, col: 21, offset: 653},
						name: "Comment",
					},
					&ruleRefExpr{
						pos:  position{line: 28, col: 31, offset: 663},
						name: "FuncDefn",
					},
					&ruleRefExpr{
						pos:  position{line: 28, col: 42, offset: 674},
						name: "TypeDefn",
					},
				},
			},
		},
		{
			name: "Statement",
			pos:  position{line: 30, col: 1, offset: 684},
			expr: &choiceExpr{
				pos: position{line: 30, col: 13, offset: 696},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 30, col: 13, offset: 696},
						name: "Comment",
					},
					&ruleRefExpr{
						pos:  position{line: 30, col: 23, offset: 706},
						name: "FuncDefn",
					},
					&ruleRefExpr{
						pos:  position{line: 30, col: 34, offset: 717},
						name: "Assignment",
					},
					&ruleRefExpr{
						pos:  position{line: 30, col: 47, offset: 730},
						name: "ExprLine",
					},
				},
			},
		},
		{
			name: "ExprLine",
			pos:  position{line: 32, col: 1, offset: 740},
			expr: &actionExpr{
				pos: position{line: 32, col: 12, offset: 751},
				run: (*parser).callonExprLine1,
				expr: &seqExpr{
					pos: position{line: 32, col: 12, offset: 751},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 32, col: 12, offset: 751},
							label: "e",
							expr: &ruleRefExpr{
								pos:  position{line: 32, col: 14, offset: 753},
								name: "Expr",
							},
						},
						&andExpr{
							pos: position{line: 32, col: 19, offset: 758},
							expr: &litMatcher{
								pos:        position{line: 32, col: 20, offset: 759},
								val:        "\n",
								ignoreCase: false,
							},
						},
					},
				},
			},
		},
		{
			name: "Comment",
			pos:  position{line: 36, col: 1, offset: 787},
			expr: &actionExpr{
				pos: position{line: 36, col: 11, offset: 797},
				run: (*parser).callonComment1,
				expr: &seqExpr{
					pos: position{line: 36, col: 11, offset: 797},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 36, col: 11, offset: 797},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 36, col: 13, offset: 799},
							val:        "#",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 36, col: 17, offset: 803},
							label: "comment",
							expr: &zeroOrMoreExpr{
								pos: position{line: 36, col: 25, offset: 811},
								expr: &seqExpr{
									pos: position{line: 36, col: 26, offset: 812},
									exprs: []interface{}{
										&notExpr{
											pos: position{line: 36, col: 26, offset: 812},
											expr: &ruleRefExpr{
												pos:  position{line: 36, col: 27, offset: 813},
												name: "EscapedChar",
											},
										},
										&anyMatcher{
											line: 36, col: 39, offset: 825,
										},
									},
								},
							},
						},
						&andExpr{
							pos: position{line: 36, col: 43, offset: 829},
							expr: &litMatcher{
								pos:        position{line: 36, col: 44, offset: 830},
								val:        "\n",
								ignoreCase: false,
							},
						},
					},
				},
			},
		},
		{
			name: "TypeDefn",
			pos:  position{line: 41, col: 1, offset: 943},
			expr: &choiceExpr{
				pos: position{line: 41, col: 12, offset: 954},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 41, col: 12, offset: 954},
						run: (*parser).callonTypeDefn2,
						expr: &seqExpr{
							pos: position{line: 41, col: 12, offset: 954},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 41, col: 12, offset: 954},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 41, col: 14, offset: 956},
									val:        "type",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 41, col: 21, offset: 963},
									name: "__",
								},
								&labeledExpr{
									pos:   position{line: 41, col: 24, offset: 966},
									label: "name",
									expr: &ruleRefExpr{
										pos:  position{line: 41, col: 29, offset: 971},
										name: "ModuleName",
									},
								},
								&labeledExpr{
									pos:   position{line: 41, col: 40, offset: 982},
									label: "params",
									expr: &zeroOrMoreExpr{
										pos: position{line: 41, col: 47, offset: 989},
										expr: &seqExpr{
											pos: position{line: 41, col: 48, offset: 990},
											exprs: []interface{}{
												&ruleRefExpr{
													pos:  position{line: 41, col: 48, offset: 990},
													name: "__",
												},
												&ruleRefExpr{
													pos:  position{line: 41, col: 51, offset: 993},
													name: "TypeParameter",
												},
											},
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 41, col: 67, offset: 1009},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 41, col: 69, offset: 1011},
									val:        "=",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 41, col: 73, offset: 1015},
									label: "types",
									expr: &oneOrMoreExpr{
										pos: position{line: 41, col: 79, offset: 1021},
										expr: &seqExpr{
											pos: position{line: 41, col: 80, offset: 1022},
											exprs: []interface{}{
												&ruleRefExpr{
													pos:  position{line: 41, col: 80, offset: 1022},
													name: "__",
												},
												&ruleRefExpr{
													pos:  position{line: 41, col: 83, offset: 1025},
													name: "AnyType",
												},
											},
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 41, col: 93, offset: 1035},
									name: "__",
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 60, col: 1, offset: 1531},
						run: (*parser).callonTypeDefn22,
						expr: &seqExpr{
							pos: position{line: 60, col: 1, offset: 1531},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 60, col: 1, offset: 1531},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 60, col: 3, offset: 1533},
									val:        "type",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 60, col: 10, offset: 1540},
									name: "__",
								},
								&labeledExpr{
									pos:   position{line: 60, col: 13, offset: 1543},
									label: "name",
									expr: &ruleRefExpr{
										pos:  position{line: 60, col: 18, offset: 1548},
										name: "ModuleName",
									},
								},
								&labeledExpr{
									pos:   position{line: 60, col: 29, offset: 1559},
									label: "params",
									expr: &zeroOrMoreExpr{
										pos: position{line: 60, col: 36, offset: 1566},
										expr: &seqExpr{
											pos: position{line: 60, col: 37, offset: 1567},
											exprs: []interface{}{
												&ruleRefExpr{
													pos:  position{line: 60, col: 37, offset: 1567},
													name: "__",
												},
												&ruleRefExpr{
													pos:  position{line: 60, col: 40, offset: 1570},
													name: "TypeParameter",
												},
											},
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 60, col: 56, offset: 1586},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 60, col: 58, offset: 1588},
									val:        "=",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 60, col: 62, offset: 1592},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 61, col: 5, offset: 1598},
									val:        "{",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 61, col: 9, offset: 1602},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 61, col: 11, offset: 1604},
									label: "first",
									expr: &ruleRefExpr{
										pos:  position{line: 61, col: 17, offset: 1610},
										name: "RecordFieldDefn",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 61, col: 33, offset: 1626},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 61, col: 35, offset: 1628},
									label: "rest",
									expr: &zeroOrMoreExpr{
										pos: position{line: 61, col: 40, offset: 1633},
										expr: &seqExpr{
											pos: position{line: 61, col: 41, offset: 1634},
											exprs: []interface{}{
												&litMatcher{
													pos:        position{line: 61, col: 41, offset: 1634},
													val:        ",",
													ignoreCase: false,
												},
												&ruleRefExpr{
													pos:  position{line: 61, col: 45, offset: 1638},
													name: "_",
												},
												&ruleRefExpr{
													pos:  position{line: 61, col: 47, offset: 1640},
													name: "RecordFieldDefn",
												},
												&ruleRefExpr{
													pos:  position{line: 61, col: 63, offset: 1656},
													name: "_",
												},
											},
										},
									},
								},
								&zeroOrOneExpr{
									pos: position{line: 61, col: 67, offset: 1660},
									expr: &litMatcher{
										pos:        position{line: 61, col: 67, offset: 1660},
										val:        ",",
										ignoreCase: false,
									},
								},
								&ruleRefExpr{
									pos:  position{line: 61, col: 72, offset: 1665},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 61, col: 74, offset: 1667},
									val:        "}",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 61, col: 78, offset: 1671},
									name: "__",
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 79, col: 1, offset: 2156},
						run: (*parser).callonTypeDefn54,
						expr: &seqExpr{
							pos: position{line: 79, col: 1, offset: 2156},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 79, col: 1, offset: 2156},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 79, col: 3, offset: 2158},
									val:        "type",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 79, col: 10, offset: 2165},
									name: "__",
								},
								&labeledExpr{
									pos:   position{line: 79, col: 13, offset: 2168},
									label: "name",
									expr: &ruleRefExpr{
										pos:  position{line: 79, col: 18, offset: 2173},
										name: "ModuleName",
									},
								},
								&labeledExpr{
									pos:   position{line: 79, col: 29, offset: 2184},
									label: "params",
									expr: &zeroOrMoreExpr{
										pos: position{line: 79, col: 36, offset: 2191},
										expr: &seqExpr{
											pos: position{line: 79, col: 37, offset: 2192},
											exprs: []interface{}{
												&ruleRefExpr{
													pos:  position{line: 79, col: 37, offset: 2192},
													name: "__",
												},
												&ruleRefExpr{
													pos:  position{line: 79, col: 40, offset: 2195},
													name: "TypeParameter",
												},
											},
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 79, col: 56, offset: 2211},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 79, col: 58, offset: 2213},
									val:        "=",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 79, col: 62, offset: 2217},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 79, col: 64, offset: 2219},
									label: "rest",
									expr: &oneOrMoreExpr{
										pos: position{line: 79, col: 69, offset: 2224},
										expr: &ruleRefExpr{
											pos:  position{line: 79, col: 70, offset: 2225},
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
			pos:  position{line: 94, col: 1, offset: 2632},
			expr: &actionExpr{
				pos: position{line: 94, col: 19, offset: 2650},
				run: (*parser).callonRecordFieldDefn1,
				expr: &seqExpr{
					pos: position{line: 94, col: 19, offset: 2650},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 94, col: 19, offset: 2650},
							label: "name",
							expr: &ruleRefExpr{
								pos:  position{line: 94, col: 24, offset: 2655},
								name: "VariableName",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 94, col: 37, offset: 2668},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 94, col: 39, offset: 2670},
							val:        ":",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 94, col: 43, offset: 2674},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 94, col: 45, offset: 2676},
							label: "t",
							expr: &ruleRefExpr{
								pos:  position{line: 94, col: 48, offset: 2679},
								name: "AnyType",
							},
						},
					},
				},
			},
		},
		{
			name: "VariantConstructor",
			pos:  position{line: 98, col: 1, offset: 2773},
			expr: &choiceExpr{
				pos: position{line: 98, col: 22, offset: 2794},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 98, col: 22, offset: 2794},
						run: (*parser).callonVariantConstructor2,
						expr: &seqExpr{
							pos: position{line: 98, col: 22, offset: 2794},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 98, col: 22, offset: 2794},
									val:        "|",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 98, col: 26, offset: 2798},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 98, col: 28, offset: 2800},
									label: "name",
									expr: &ruleRefExpr{
										pos:  position{line: 98, col: 33, offset: 2805},
										name: "ModuleName",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 98, col: 44, offset: 2816},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 98, col: 46, offset: 2818},
									val:        "{",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 98, col: 50, offset: 2822},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 98, col: 52, offset: 2824},
									label: "first",
									expr: &ruleRefExpr{
										pos:  position{line: 98, col: 58, offset: 2830},
										name: "RecordFieldDefn",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 98, col: 74, offset: 2846},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 98, col: 76, offset: 2848},
									label: "rest",
									expr: &zeroOrMoreExpr{
										pos: position{line: 98, col: 81, offset: 2853},
										expr: &seqExpr{
											pos: position{line: 98, col: 82, offset: 2854},
											exprs: []interface{}{
												&litMatcher{
													pos:        position{line: 98, col: 82, offset: 2854},
													val:        ",",
													ignoreCase: false,
												},
												&ruleRefExpr{
													pos:  position{line: 98, col: 86, offset: 2858},
													name: "_",
												},
												&ruleRefExpr{
													pos:  position{line: 98, col: 88, offset: 2860},
													name: "RecordFieldDefn",
												},
												&ruleRefExpr{
													pos:  position{line: 98, col: 104, offset: 2876},
													name: "_",
												},
											},
										},
									},
								},
								&zeroOrOneExpr{
									pos: position{line: 98, col: 108, offset: 2880},
									expr: &litMatcher{
										pos:        position{line: 98, col: 108, offset: 2880},
										val:        ",",
										ignoreCase: false,
									},
								},
								&ruleRefExpr{
									pos:  position{line: 98, col: 113, offset: 2885},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 98, col: 115, offset: 2887},
									val:        "}",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 98, col: 119, offset: 2891},
									name: "__",
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 117, col: 1, offset: 3496},
						run: (*parser).callonVariantConstructor26,
						expr: &seqExpr{
							pos: position{line: 117, col: 1, offset: 3496},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 117, col: 1, offset: 3496},
									val:        "|",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 117, col: 5, offset: 3500},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 117, col: 7, offset: 3502},
									label: "name",
									expr: &ruleRefExpr{
										pos:  position{line: 117, col: 12, offset: 3507},
										name: "ModuleName",
									},
								},
								&labeledExpr{
									pos:   position{line: 117, col: 23, offset: 3518},
									label: "rest",
									expr: &zeroOrMoreExpr{
										pos: position{line: 117, col: 28, offset: 3523},
										expr: &seqExpr{
											pos: position{line: 117, col: 29, offset: 3524},
											exprs: []interface{}{
												&ruleRefExpr{
													pos:  position{line: 117, col: 29, offset: 3524},
													name: "__",
												},
												&ruleRefExpr{
													pos:  position{line: 117, col: 32, offset: 3527},
													name: "AnyType",
												},
											},
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 117, col: 42, offset: 3537},
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
			pos:  position{line: 134, col: 1, offset: 3974},
			expr: &choiceExpr{
				pos: position{line: 134, col: 11, offset: 3984},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 134, col: 11, offset: 3984},
						name: "BaseType",
					},
					&ruleRefExpr{
						pos:  position{line: 134, col: 22, offset: 3995},
						name: "TypeParameter",
					},
				},
			},
		},
		{
			name: "Assignment",
			pos:  position{line: 136, col: 1, offset: 4010},
			expr: &choiceExpr{
				pos: position{line: 136, col: 14, offset: 4023},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 136, col: 14, offset: 4023},
						run: (*parser).callonAssignment2,
						expr: &seqExpr{
							pos: position{line: 136, col: 14, offset: 4023},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 136, col: 14, offset: 4023},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 136, col: 16, offset: 4025},
									val:        "let",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 136, col: 22, offset: 4031},
									name: "__",
								},
								&labeledExpr{
									pos:   position{line: 136, col: 25, offset: 4034},
									label: "i",
									expr: &ruleRefExpr{
										pos:  position{line: 136, col: 27, offset: 4036},
										name: "Assignable",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 136, col: 38, offset: 4047},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 136, col: 40, offset: 4049},
									val:        "=",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 136, col: 44, offset: 4053},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 136, col: 46, offset: 4055},
									label: "expr",
									expr: &ruleRefExpr{
										pos:  position{line: 136, col: 51, offset: 4060},
										name: "Expr",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 136, col: 56, offset: 4065},
									name: "_",
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 142, col: 1, offset: 4184},
						run: (*parser).callonAssignment15,
						expr: &seqExpr{
							pos: position{line: 142, col: 1, offset: 4184},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 142, col: 1, offset: 4184},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 142, col: 3, offset: 4186},
									val:        "let",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 142, col: 9, offset: 4192},
									name: "__",
								},
								&notExpr{
									pos: position{line: 142, col: 12, offset: 4195},
									expr: &ruleRefExpr{
										pos:  position{line: 142, col: 13, offset: 4196},
										name: "Assignable",
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 146, col: 1, offset: 4294},
						run: (*parser).callonAssignment22,
						expr: &seqExpr{
							pos: position{line: 146, col: 1, offset: 4294},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 146, col: 1, offset: 4294},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 146, col: 3, offset: 4296},
									val:        "let",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 146, col: 9, offset: 4302},
									name: "__",
								},
								&labeledExpr{
									pos:   position{line: 146, col: 12, offset: 4305},
									label: "i",
									expr: &ruleRefExpr{
										pos:  position{line: 146, col: 14, offset: 4307},
										name: "Assignable",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 146, col: 25, offset: 4318},
									name: "_",
								},
								&notExpr{
									pos: position{line: 146, col: 27, offset: 4320},
									expr: &litMatcher{
										pos:        position{line: 146, col: 28, offset: 4321},
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
			pos:  position{line: 150, col: 1, offset: 4415},
			expr: &actionExpr{
				pos: position{line: 150, col: 12, offset: 4426},
				run: (*parser).callonFuncDefn1,
				expr: &seqExpr{
					pos: position{line: 150, col: 12, offset: 4426},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 150, col: 12, offset: 4426},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 150, col: 14, offset: 4428},
							val:        "let",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 150, col: 20, offset: 4434},
							name: "__",
						},
						&labeledExpr{
							pos:   position{line: 150, col: 23, offset: 4437},
							label: "i",
							expr: &ruleRefExpr{
								pos:  position{line: 150, col: 25, offset: 4439},
								name: "VariableName",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 150, col: 38, offset: 4452},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 150, col: 40, offset: 4454},
							val:        "=",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 150, col: 44, offset: 4458},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 150, col: 46, offset: 4460},
							val:        "func",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 150, col: 53, offset: 4467},
							name: "__",
						},
						&labeledExpr{
							pos:   position{line: 150, col: 56, offset: 4470},
							label: "ids",
							expr: &zeroOrMoreExpr{
								pos: position{line: 150, col: 60, offset: 4474},
								expr: &seqExpr{
									pos: position{line: 150, col: 61, offset: 4475},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 150, col: 61, offset: 4475},
											name: "VariableName",
										},
										&ruleRefExpr{
											pos:  position{line: 150, col: 74, offset: 4488},
											name: "__",
										},
									},
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 150, col: 79, offset: 4493},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 150, col: 81, offset: 4495},
							val:        "{",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 150, col: 85, offset: 4499},
							name: "__",
						},
						&labeledExpr{
							pos:   position{line: 150, col: 88, offset: 4502},
							label: "statements",
							expr: &oneOrMoreExpr{
								pos: position{line: 150, col: 99, offset: 4513},
								expr: &ruleRefExpr{
									pos:  position{line: 150, col: 100, offset: 4514},
									name: "Statement",
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 150, col: 112, offset: 4526},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 150, col: 114, offset: 4528},
							val:        "}",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 150, col: 118, offset: 4532},
							name: "_",
						},
					},
				},
			},
		},
		{
			name: "Expr",
			pos:  position{line: 173, col: 1, offset: 5188},
			expr: &actionExpr{
				pos: position{line: 173, col: 8, offset: 5195},
				run: (*parser).callonExpr1,
				expr: &labeledExpr{
					pos:   position{line: 173, col: 8, offset: 5195},
					label: "ex",
					expr: &choiceExpr{
						pos: position{line: 173, col: 12, offset: 5199},
						alternatives: []interface{}{
							&ruleRefExpr{
								pos:  position{line: 173, col: 12, offset: 5199},
								name: "IfExpr",
							},
							&ruleRefExpr{
								pos:  position{line: 173, col: 21, offset: 5208},
								name: "BinOp",
							},
							&ruleRefExpr{
								pos:  position{line: 173, col: 29, offset: 5216},
								name: "Call",
							},
						},
					},
				},
			},
		},
		{
			name: "IfExpr",
			pos:  position{line: 179, col: 1, offset: 5325},
			expr: &choiceExpr{
				pos: position{line: 179, col: 10, offset: 5334},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 179, col: 10, offset: 5334},
						run: (*parser).callonIfExpr2,
						expr: &seqExpr{
							pos: position{line: 179, col: 10, offset: 5334},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 179, col: 10, offset: 5334},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 179, col: 12, offset: 5336},
									val:        "if",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 179, col: 17, offset: 5341},
									name: "__",
								},
								&labeledExpr{
									pos:   position{line: 179, col: 20, offset: 5344},
									label: "expr",
									expr: &ruleRefExpr{
										pos:  position{line: 179, col: 25, offset: 5349},
										name: "BinOpBool",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 179, col: 35, offset: 5359},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 179, col: 37, offset: 5361},
									val:        "{",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 179, col: 41, offset: 5365},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 179, col: 43, offset: 5367},
									label: "thens",
									expr: &oneOrMoreExpr{
										pos: position{line: 179, col: 49, offset: 5373},
										expr: &ruleRefExpr{
											pos:  position{line: 179, col: 50, offset: 5374},
											name: "Statement",
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 179, col: 62, offset: 5386},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 179, col: 64, offset: 5388},
									val:        "}",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 179, col: 68, offset: 5392},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 179, col: 70, offset: 5394},
									val:        "else",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 179, col: 77, offset: 5401},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 179, col: 79, offset: 5403},
									label: "elseifs",
									expr: &ruleRefExpr{
										pos:  position{line: 179, col: 87, offset: 5411},
										name: "IfExpr",
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 191, col: 1, offset: 5741},
						run: (*parser).callonIfExpr22,
						expr: &seqExpr{
							pos: position{line: 191, col: 1, offset: 5741},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 191, col: 1, offset: 5741},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 191, col: 3, offset: 5743},
									val:        "if",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 191, col: 8, offset: 5748},
									name: "__",
								},
								&labeledExpr{
									pos:   position{line: 191, col: 11, offset: 5751},
									label: "expr",
									expr: &ruleRefExpr{
										pos:  position{line: 191, col: 16, offset: 5756},
										name: "BinOpBool",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 191, col: 26, offset: 5766},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 191, col: 28, offset: 5768},
									val:        "{",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 191, col: 32, offset: 5772},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 191, col: 34, offset: 5774},
									label: "thens",
									expr: &oneOrMoreExpr{
										pos: position{line: 191, col: 40, offset: 5780},
										expr: &ruleRefExpr{
											pos:  position{line: 191, col: 41, offset: 5781},
											name: "Statement",
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 191, col: 53, offset: 5793},
									name: "__",
								},
								&litMatcher{
									pos:        position{line: 191, col: 56, offset: 5796},
									val:        "}",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 191, col: 60, offset: 5800},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 191, col: 62, offset: 5802},
									val:        "else",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 191, col: 69, offset: 5809},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 191, col: 71, offset: 5811},
									val:        "{",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 191, col: 75, offset: 5815},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 191, col: 77, offset: 5817},
									label: "elses",
									expr: &oneOrMoreExpr{
										pos: position{line: 191, col: 83, offset: 5823},
										expr: &ruleRefExpr{
											pos:  position{line: 191, col: 84, offset: 5824},
											name: "Statement",
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 191, col: 96, offset: 5836},
									name: "__",
								},
								&litMatcher{
									pos:        position{line: 191, col: 99, offset: 5839},
									val:        "}",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 210, col: 1, offset: 6342},
						run: (*parser).callonIfExpr47,
						expr: &seqExpr{
							pos: position{line: 210, col: 1, offset: 6342},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 210, col: 1, offset: 6342},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 210, col: 3, offset: 6344},
									val:        "if",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 210, col: 8, offset: 6349},
									name: "__",
								},
								&labeledExpr{
									pos:   position{line: 210, col: 11, offset: 6352},
									label: "expr",
									expr: &ruleRefExpr{
										pos:  position{line: 210, col: 16, offset: 6357},
										name: "BinOpBool",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 210, col: 26, offset: 6367},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 210, col: 28, offset: 6369},
									val:        "{",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 210, col: 32, offset: 6373},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 210, col: 34, offset: 6375},
									label: "thens",
									expr: &oneOrMoreExpr{
										pos: position{line: 210, col: 40, offset: 6381},
										expr: &ruleRefExpr{
											pos:  position{line: 210, col: 41, offset: 6382},
											name: "Statement",
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 210, col: 53, offset: 6394},
									name: "__",
								},
								&litMatcher{
									pos:        position{line: 210, col: 56, offset: 6397},
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
			pos:  position{line: 222, col: 1, offset: 6695},
			expr: &choiceExpr{
				pos: position{line: 222, col: 8, offset: 6702},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 222, col: 8, offset: 6702},
						run: (*parser).callonCall2,
						expr: &seqExpr{
							pos: position{line: 222, col: 8, offset: 6702},
							exprs: []interface{}{
								&labeledExpr{
									pos:   position{line: 222, col: 8, offset: 6702},
									label: "module",
									expr: &ruleRefExpr{
										pos:  position{line: 222, col: 15, offset: 6709},
										name: "ModuleName",
									},
								},
								&litMatcher{
									pos:        position{line: 222, col: 26, offset: 6720},
									val:        ".",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 222, col: 30, offset: 6724},
									label: "fn",
									expr: &ruleRefExpr{
										pos:  position{line: 222, col: 33, offset: 6727},
										name: "VariableName",
									},
								},
								&labeledExpr{
									pos:   position{line: 222, col: 46, offset: 6740},
									label: "args",
									expr: &ruleRefExpr{
										pos:  position{line: 222, col: 51, offset: 6745},
										name: "Arguments",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 222, col: 61, offset: 6755},
									name: "_",
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 236, col: 1, offset: 7057},
						run: (*parser).callonCall12,
						expr: &seqExpr{
							pos: position{line: 236, col: 1, offset: 7057},
							exprs: []interface{}{
								&labeledExpr{
									pos:   position{line: 236, col: 1, offset: 7057},
									label: "fn",
									expr: &ruleRefExpr{
										pos:  position{line: 236, col: 4, offset: 7060},
										name: "VariableName",
									},
								},
								&labeledExpr{
									pos:   position{line: 236, col: 17, offset: 7073},
									label: "args",
									expr: &ruleRefExpr{
										pos:  position{line: 236, col: 22, offset: 7078},
										name: "Arguments",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 236, col: 32, offset: 7088},
									name: "_",
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 250, col: 1, offset: 7381},
						run: (*parser).callonCall19,
						expr: &seqExpr{
							pos: position{line: 250, col: 1, offset: 7381},
							exprs: []interface{}{
								&labeledExpr{
									pos:   position{line: 250, col: 1, offset: 7381},
									label: "fn",
									expr: &ruleRefExpr{
										pos:  position{line: 250, col: 4, offset: 7384},
										name: "VariableName",
									},
								},
								&litMatcher{
									pos:        position{line: 250, col: 17, offset: 7397},
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
			name: "Arguments",
			pos:  position{line: 258, col: 1, offset: 7553},
			expr: &choiceExpr{
				pos: position{line: 258, col: 13, offset: 7565},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 258, col: 13, offset: 7565},
						run: (*parser).callonArguments2,
						expr: &seqExpr{
							pos: position{line: 258, col: 13, offset: 7565},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 258, col: 13, offset: 7565},
									val:        "(",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 258, col: 17, offset: 7569},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 258, col: 19, offset: 7571},
									label: "argument",
									expr: &ruleRefExpr{
										pos:  position{line: 258, col: 28, offset: 7580},
										name: "BinOpParens",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 258, col: 40, offset: 7592},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 258, col: 42, offset: 7594},
									label: "rest",
									expr: &zeroOrMoreExpr{
										pos: position{line: 258, col: 47, offset: 7599},
										expr: &seqExpr{
											pos: position{line: 258, col: 48, offset: 7600},
											exprs: []interface{}{
												&litMatcher{
													pos:        position{line: 258, col: 48, offset: 7600},
													val:        ",",
													ignoreCase: false,
												},
												&ruleRefExpr{
													pos:  position{line: 258, col: 52, offset: 7604},
													name: "_",
												},
												&ruleRefExpr{
													pos:  position{line: 258, col: 54, offset: 7606},
													name: "BinOpParens",
												},
											},
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 258, col: 68, offset: 7620},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 258, col: 70, offset: 7622},
									val:        ")",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 275, col: 1, offset: 8044},
						run: (*parser).callonArguments17,
						expr: &seqExpr{
							pos: position{line: 275, col: 1, offset: 8044},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 275, col: 1, offset: 8044},
									val:        "(",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 275, col: 5, offset: 8048},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 275, col: 7, offset: 8050},
									label: "argument",
									expr: &ruleRefExpr{
										pos:  position{line: 275, col: 16, offset: 8059},
										name: "Expr",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 275, col: 21, offset: 8064},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 275, col: 23, offset: 8066},
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
			pos:  position{line: 280, col: 1, offset: 8171},
			expr: &actionExpr{
				pos: position{line: 280, col: 16, offset: 8186},
				run: (*parser).callonCompoundExpr1,
				expr: &seqExpr{
					pos: position{line: 280, col: 16, offset: 8186},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 280, col: 16, offset: 8186},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 280, col: 18, offset: 8188},
							label: "op",
							expr: &ruleRefExpr{
								pos:  position{line: 280, col: 21, offset: 8191},
								name: "BinOp",
							},
						},
						&labeledExpr{
							pos:   position{line: 280, col: 27, offset: 8197},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 280, col: 32, offset: 8202},
								expr: &seqExpr{
									pos: position{line: 280, col: 33, offset: 8203},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 280, col: 33, offset: 8203},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 280, col: 36, offset: 8206},
											name: "Operator",
										},
										&ruleRefExpr{
											pos:  position{line: 280, col: 45, offset: 8215},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 280, col: 48, offset: 8218},
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
			pos:  position{line: 300, col: 1, offset: 8824},
			expr: &choiceExpr{
				pos: position{line: 300, col: 9, offset: 8832},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 300, col: 9, offset: 8832},
						name: "BinOpBool",
					},
					&ruleRefExpr{
						pos:  position{line: 300, col: 21, offset: 8844},
						name: "BinOpEquality",
					},
					&ruleRefExpr{
						pos:  position{line: 300, col: 37, offset: 8860},
						name: "BinOpLow",
					},
					&ruleRefExpr{
						pos:  position{line: 300, col: 48, offset: 8871},
						name: "BinOpHigh",
					},
					&ruleRefExpr{
						pos:  position{line: 300, col: 60, offset: 8883},
						name: "BinOpParens",
					},
				},
			},
		},
		{
			name: "BinOpBool",
			pos:  position{line: 302, col: 1, offset: 8896},
			expr: &actionExpr{
				pos: position{line: 302, col: 13, offset: 8908},
				run: (*parser).callonBinOpBool1,
				expr: &seqExpr{
					pos: position{line: 302, col: 13, offset: 8908},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 302, col: 13, offset: 8908},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 302, col: 15, offset: 8910},
							label: "first",
							expr: &ruleRefExpr{
								pos:  position{line: 302, col: 21, offset: 8916},
								name: "BinOpEquality",
							},
						},
						&labeledExpr{
							pos:   position{line: 302, col: 35, offset: 8930},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 302, col: 40, offset: 8935},
								expr: &seqExpr{
									pos: position{line: 302, col: 41, offset: 8936},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 302, col: 41, offset: 8936},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 302, col: 44, offset: 8939},
											name: "OperatorBoolean",
										},
										&ruleRefExpr{
											pos:  position{line: 302, col: 60, offset: 8955},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 302, col: 63, offset: 8958},
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
			pos:  position{line: 335, col: 1, offset: 9851},
			expr: &actionExpr{
				pos: position{line: 335, col: 17, offset: 9867},
				run: (*parser).callonBinOpEquality1,
				expr: &seqExpr{
					pos: position{line: 335, col: 17, offset: 9867},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 335, col: 17, offset: 9867},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 335, col: 19, offset: 9869},
							label: "first",
							expr: &ruleRefExpr{
								pos:  position{line: 335, col: 25, offset: 9875},
								name: "BinOpLow",
							},
						},
						&labeledExpr{
							pos:   position{line: 335, col: 34, offset: 9884},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 335, col: 39, offset: 9889},
								expr: &seqExpr{
									pos: position{line: 335, col: 40, offset: 9890},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 335, col: 40, offset: 9890},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 335, col: 43, offset: 9893},
											name: "OperatorEquality",
										},
										&ruleRefExpr{
											pos:  position{line: 335, col: 60, offset: 9910},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 335, col: 63, offset: 9913},
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
			pos:  position{line: 367, col: 1, offset: 10800},
			expr: &actionExpr{
				pos: position{line: 367, col: 12, offset: 10811},
				run: (*parser).callonBinOpLow1,
				expr: &seqExpr{
					pos: position{line: 367, col: 12, offset: 10811},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 367, col: 12, offset: 10811},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 367, col: 14, offset: 10813},
							label: "first",
							expr: &ruleRefExpr{
								pos:  position{line: 367, col: 20, offset: 10819},
								name: "BinOpHigh",
							},
						},
						&labeledExpr{
							pos:   position{line: 367, col: 30, offset: 10829},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 367, col: 35, offset: 10834},
								expr: &seqExpr{
									pos: position{line: 367, col: 36, offset: 10835},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 367, col: 36, offset: 10835},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 367, col: 39, offset: 10838},
											name: "OperatorLow",
										},
										&ruleRefExpr{
											pos:  position{line: 367, col: 51, offset: 10850},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 367, col: 54, offset: 10853},
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
			pos:  position{line: 399, col: 1, offset: 11741},
			expr: &actionExpr{
				pos: position{line: 399, col: 13, offset: 11753},
				run: (*parser).callonBinOpHigh1,
				expr: &seqExpr{
					pos: position{line: 399, col: 13, offset: 11753},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 399, col: 13, offset: 11753},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 399, col: 15, offset: 11755},
							label: "first",
							expr: &ruleRefExpr{
								pos:  position{line: 399, col: 21, offset: 11761},
								name: "BinOpParens",
							},
						},
						&labeledExpr{
							pos:   position{line: 399, col: 33, offset: 11773},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 399, col: 38, offset: 11778},
								expr: &seqExpr{
									pos: position{line: 399, col: 39, offset: 11779},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 399, col: 39, offset: 11779},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 399, col: 42, offset: 11782},
											name: "OperatorHigh",
										},
										&ruleRefExpr{
											pos:  position{line: 399, col: 55, offset: 11795},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 399, col: 58, offset: 11798},
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
			pos:  position{line: 430, col: 1, offset: 12687},
			expr: &choiceExpr{
				pos: position{line: 430, col: 15, offset: 12701},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 430, col: 15, offset: 12701},
						run: (*parser).callonBinOpParens2,
						expr: &seqExpr{
							pos: position{line: 430, col: 15, offset: 12701},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 430, col: 15, offset: 12701},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 430, col: 17, offset: 12703},
									val:        "(",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 430, col: 21, offset: 12707},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 430, col: 23, offset: 12709},
									label: "first",
									expr: &ruleRefExpr{
										pos:  position{line: 430, col: 29, offset: 12715},
										name: "BinOp",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 430, col: 35, offset: 12721},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 430, col: 37, offset: 12723},
									val:        ")",
									ignoreCase: false,
								},
							},
						},
					},
					&ruleRefExpr{
						pos:  position{line: 433, col: 5, offset: 12846},
						name: "Value",
					},
				},
			},
		},
		{
			name: "Operator",
			pos:  position{line: 435, col: 1, offset: 12853},
			expr: &choiceExpr{
				pos: position{line: 435, col: 12, offset: 12864},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 435, col: 12, offset: 12864},
						name: "OperatorBoolean",
					},
					&ruleRefExpr{
						pos:  position{line: 435, col: 30, offset: 12882},
						name: "OperatorEquality",
					},
					&ruleRefExpr{
						pos:  position{line: 435, col: 49, offset: 12901},
						name: "OperatorHigh",
					},
					&ruleRefExpr{
						pos:  position{line: 435, col: 64, offset: 12916},
						name: "OperatorLow",
					},
				},
			},
		},
		{
			name: "OperatorBoolean",
			pos:  position{line: 437, col: 1, offset: 12929},
			expr: &actionExpr{
				pos: position{line: 437, col: 19, offset: 12947},
				run: (*parser).callonOperatorBoolean1,
				expr: &choiceExpr{
					pos: position{line: 437, col: 21, offset: 12949},
					alternatives: []interface{}{
						&litMatcher{
							pos:        position{line: 437, col: 21, offset: 12949},
							val:        "or",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 437, col: 28, offset: 12956},
							val:        "and",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "OperatorEquality",
			pos:  position{line: 441, col: 1, offset: 13038},
			expr: &actionExpr{
				pos: position{line: 441, col: 20, offset: 13057},
				run: (*parser).callonOperatorEquality1,
				expr: &choiceExpr{
					pos: position{line: 441, col: 22, offset: 13059},
					alternatives: []interface{}{
						&litMatcher{
							pos:        position{line: 441, col: 22, offset: 13059},
							val:        ">=",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 441, col: 29, offset: 13066},
							val:        "<=",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 441, col: 36, offset: 13073},
							val:        "<",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 441, col: 42, offset: 13079},
							val:        ">",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 441, col: 48, offset: 13085},
							val:        "===",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 441, col: 56, offset: 13093},
							val:        "==",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "OperatorHigh",
			pos:  position{line: 445, col: 1, offset: 13172},
			expr: &choiceExpr{
				pos: position{line: 445, col: 16, offset: 13187},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 445, col: 16, offset: 13187},
						run: (*parser).callonOperatorHigh2,
						expr: &choiceExpr{
							pos: position{line: 445, col: 18, offset: 13189},
							alternatives: []interface{}{
								&litMatcher{
									pos:        position{line: 445, col: 18, offset: 13189},
									val:        "/",
									ignoreCase: false,
								},
								&litMatcher{
									pos:        position{line: 445, col: 24, offset: 13195},
									val:        "*",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 448, col: 3, offset: 13278},
						run: (*parser).callonOperatorHigh6,
						expr: &litMatcher{
							pos:        position{line: 448, col: 5, offset: 13280},
							val:        "^",
							ignoreCase: false,
						},
					},
					&actionExpr{
						pos: position{line: 451, col: 3, offset: 13360},
						run: (*parser).callonOperatorHigh8,
						expr: &litMatcher{
							pos:        position{line: 451, col: 3, offset: 13360},
							val:        "++",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "OperatorLow",
			pos:  position{line: 455, col: 1, offset: 13441},
			expr: &actionExpr{
				pos: position{line: 455, col: 15, offset: 13455},
				run: (*parser).callonOperatorLow1,
				expr: &choiceExpr{
					pos: position{line: 455, col: 17, offset: 13457},
					alternatives: []interface{}{
						&litMatcher{
							pos:        position{line: 455, col: 17, offset: 13457},
							val:        "+",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 455, col: 23, offset: 13463},
							val:        "-",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "Value",
			pos:  position{line: 459, col: 1, offset: 13545},
			expr: &choiceExpr{
				pos: position{line: 459, col: 9, offset: 13553},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 459, col: 9, offset: 13553},
						name: "Call",
					},
					&ruleRefExpr{
						pos:  position{line: 459, col: 16, offset: 13560},
						name: "VariableName",
					},
					&ruleRefExpr{
						pos:  position{line: 459, col: 31, offset: 13575},
						name: "Const",
					},
				},
			},
		},
		{
			name: "Assignable",
			pos:  position{line: 461, col: 1, offset: 13582},
			expr: &choiceExpr{
				pos: position{line: 461, col: 14, offset: 13595},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 461, col: 14, offset: 13595},
						name: "VariableName",
					},
					&ruleRefExpr{
						pos:  position{line: 461, col: 29, offset: 13610},
						name: "Unused",
					},
				},
			},
		},
		{
			name: "Identifier",
			pos:  position{line: 463, col: 1, offset: 13618},
			expr: &choiceExpr{
				pos: position{line: 463, col: 14, offset: 13631},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 463, col: 14, offset: 13631},
						name: "VariableName",
					},
					&ruleRefExpr{
						pos:  position{line: 463, col: 29, offset: 13646},
						name: "ModuleName",
					},
				},
			},
		},
		{
			name: "ArrayLiteral",
			pos:  position{line: 465, col: 1, offset: 13658},
			expr: &actionExpr{
				pos: position{line: 465, col: 16, offset: 13673},
				run: (*parser).callonArrayLiteral1,
				expr: &seqExpr{
					pos: position{line: 465, col: 16, offset: 13673},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 465, col: 16, offset: 13673},
							val:        "[",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 465, col: 20, offset: 13677},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 465, col: 22, offset: 13679},
							label: "first",
							expr: &ruleRefExpr{
								pos:  position{line: 465, col: 28, offset: 13685},
								name: "Expr",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 465, col: 33, offset: 13690},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 465, col: 35, offset: 13692},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 465, col: 40, offset: 13697},
								expr: &seqExpr{
									pos: position{line: 465, col: 41, offset: 13698},
									exprs: []interface{}{
										&litMatcher{
											pos:        position{line: 465, col: 41, offset: 13698},
											val:        ",",
											ignoreCase: false,
										},
										&ruleRefExpr{
											pos:  position{line: 465, col: 45, offset: 13702},
											name: "_",
										},
										&ruleRefExpr{
											pos:  position{line: 465, col: 47, offset: 13704},
											name: "Expr",
										},
										&ruleRefExpr{
											pos:  position{line: 465, col: 52, offset: 13709},
											name: "_",
										},
									},
								},
							},
						},
						&zeroOrOneExpr{
							pos: position{line: 465, col: 56, offset: 13713},
							expr: &litMatcher{
								pos:        position{line: 465, col: 56, offset: 13713},
								val:        ",",
								ignoreCase: false,
							},
						},
						&ruleRefExpr{
							pos:  position{line: 465, col: 61, offset: 13718},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 465, col: 63, offset: 13720},
							val:        "]",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "VariantInstance",
			pos:  position{line: 481, col: 1, offset: 14165},
			expr: &actionExpr{
				pos: position{line: 481, col: 19, offset: 14183},
				run: (*parser).callonVariantInstance1,
				expr: &seqExpr{
					pos: position{line: 481, col: 19, offset: 14183},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 481, col: 19, offset: 14183},
							label: "name",
							expr: &ruleRefExpr{
								pos:  position{line: 481, col: 24, offset: 14188},
								name: "ModuleName",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 481, col: 35, offset: 14199},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 481, col: 37, offset: 14201},
							label: "args",
							expr: &ruleRefExpr{
								pos:  position{line: 481, col: 42, offset: 14206},
								name: "Arguments",
							},
						},
					},
				},
			},
		},
		{
			name: "RecordInstance",
			pos:  position{line: 494, col: 1, offset: 14476},
			expr: &actionExpr{
				pos: position{line: 494, col: 18, offset: 14493},
				run: (*parser).callonRecordInstance1,
				expr: &seqExpr{
					pos: position{line: 494, col: 18, offset: 14493},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 494, col: 18, offset: 14493},
							label: "name",
							expr: &ruleRefExpr{
								pos:  position{line: 494, col: 23, offset: 14498},
								name: "ModuleName",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 494, col: 34, offset: 14509},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 494, col: 36, offset: 14511},
							val:        "{",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 494, col: 40, offset: 14515},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 494, col: 42, offset: 14517},
							label: "firstName",
							expr: &ruleRefExpr{
								pos:  position{line: 494, col: 52, offset: 14527},
								name: "VariableName",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 494, col: 65, offset: 14540},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 494, col: 67, offset: 14542},
							val:        ":",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 494, col: 71, offset: 14546},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 494, col: 73, offset: 14548},
							label: "firstValue",
							expr: &ruleRefExpr{
								pos:  position{line: 494, col: 84, offset: 14559},
								name: "Expr",
							},
						},
						&labeledExpr{
							pos:   position{line: 494, col: 89, offset: 14564},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 494, col: 94, offset: 14569},
								expr: &seqExpr{
									pos: position{line: 494, col: 95, offset: 14570},
									exprs: []interface{}{
										&litMatcher{
											pos:        position{line: 494, col: 95, offset: 14570},
											val:        ",",
											ignoreCase: false,
										},
										&ruleRefExpr{
											pos:  position{line: 494, col: 99, offset: 14574},
											name: "_",
										},
										&ruleRefExpr{
											pos:  position{line: 494, col: 101, offset: 14576},
											name: "VariableName",
										},
										&ruleRefExpr{
											pos:  position{line: 494, col: 114, offset: 14589},
											name: "_",
										},
										&litMatcher{
											pos:        position{line: 494, col: 116, offset: 14591},
											val:        ":",
											ignoreCase: false,
										},
										&ruleRefExpr{
											pos:  position{line: 494, col: 120, offset: 14595},
											name: "_",
										},
										&ruleRefExpr{
											pos:  position{line: 494, col: 122, offset: 14597},
											name: "Expr",
										},
									},
								},
							},
						},
						&litMatcher{
							pos:        position{line: 494, col: 130, offset: 14605},
							val:        "}",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "TypeParameter",
			pos:  position{line: 514, col: 1, offset: 15189},
			expr: &actionExpr{
				pos: position{line: 514, col: 17, offset: 15205},
				run: (*parser).callonTypeParameter1,
				expr: &seqExpr{
					pos: position{line: 514, col: 17, offset: 15205},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 514, col: 17, offset: 15205},
							val:        "'",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 514, col: 22, offset: 15210},
							name: "VariableName",
						},
					},
				},
			},
		},
		{
			name: "VariableName",
			pos:  position{line: 518, col: 1, offset: 15283},
			expr: &actionExpr{
				pos: position{line: 518, col: 16, offset: 15298},
				run: (*parser).callonVariableName1,
				expr: &seqExpr{
					pos: position{line: 518, col: 16, offset: 15298},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 518, col: 16, offset: 15298},
							expr: &ruleRefExpr{
								pos:  position{line: 518, col: 17, offset: 15299},
								name: "Reserved",
							},
						},
						&seqExpr{
							pos: position{line: 518, col: 27, offset: 15309},
							exprs: []interface{}{
								&oneOrMoreExpr{
									pos: position{line: 518, col: 27, offset: 15309},
									expr: &charClassMatcher{
										pos:        position{line: 518, col: 27, offset: 15309},
										val:        "[a-z]",
										ranges:     []rune{'a', 'z'},
										ignoreCase: false,
										inverted:   false,
									},
								},
								&zeroOrMoreExpr{
									pos: position{line: 518, col: 34, offset: 15316},
									expr: &charClassMatcher{
										pos:        position{line: 518, col: 34, offset: 15316},
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
			pos:  position{line: 522, col: 1, offset: 15391},
			expr: &actionExpr{
				pos: position{line: 522, col: 14, offset: 15404},
				run: (*parser).callonModuleName1,
				expr: &seqExpr{
					pos: position{line: 522, col: 15, offset: 15405},
					exprs: []interface{}{
						&oneOrMoreExpr{
							pos: position{line: 522, col: 15, offset: 15405},
							expr: &charClassMatcher{
								pos:        position{line: 522, col: 15, offset: 15405},
								val:        "[A-Z]",
								ranges:     []rune{'A', 'Z'},
								ignoreCase: false,
								inverted:   false,
							},
						},
						&zeroOrMoreExpr{
							pos: position{line: 522, col: 22, offset: 15412},
							expr: &charClassMatcher{
								pos:        position{line: 522, col: 22, offset: 15412},
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
			pos:  position{line: 526, col: 1, offset: 15487},
			expr: &choiceExpr{
				pos: position{line: 526, col: 9, offset: 15495},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 526, col: 9, offset: 15495},
						run: (*parser).callonConst2,
						expr: &seqExpr{
							pos: position{line: 526, col: 9, offset: 15495},
							exprs: []interface{}{
								&zeroOrOneExpr{
									pos: position{line: 526, col: 9, offset: 15495},
									expr: &litMatcher{
										pos:        position{line: 526, col: 9, offset: 15495},
										val:        "-",
										ignoreCase: false,
									},
								},
								&oneOrMoreExpr{
									pos: position{line: 526, col: 14, offset: 15500},
									expr: &charClassMatcher{
										pos:        position{line: 526, col: 14, offset: 15500},
										val:        "[0-9]",
										ranges:     []rune{'0', '9'},
										ignoreCase: false,
										inverted:   false,
									},
								},
								&notExpr{
									pos: position{line: 526, col: 21, offset: 15507},
									expr: &litMatcher{
										pos:        position{line: 526, col: 22, offset: 15508},
										val:        ".",
										ignoreCase: false,
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 533, col: 3, offset: 15683},
						run: (*parser).callonConst10,
						expr: &seqExpr{
							pos: position{line: 533, col: 3, offset: 15683},
							exprs: []interface{}{
								&zeroOrOneExpr{
									pos: position{line: 533, col: 3, offset: 15683},
									expr: &litMatcher{
										pos:        position{line: 533, col: 3, offset: 15683},
										val:        "-",
										ignoreCase: false,
									},
								},
								&oneOrMoreExpr{
									pos: position{line: 533, col: 8, offset: 15688},
									expr: &charClassMatcher{
										pos:        position{line: 533, col: 8, offset: 15688},
										val:        "[0-9]",
										ranges:     []rune{'0', '9'},
										ignoreCase: false,
										inverted:   false,
									},
								},
								&litMatcher{
									pos:        position{line: 533, col: 15, offset: 15695},
									val:        ".",
									ignoreCase: false,
								},
								&oneOrMoreExpr{
									pos: position{line: 533, col: 19, offset: 15699},
									expr: &charClassMatcher{
										pos:        position{line: 533, col: 19, offset: 15699},
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
						pos: position{line: 540, col: 3, offset: 15888},
						run: (*parser).callonConst19,
						expr: &litMatcher{
							pos:        position{line: 540, col: 3, offset: 15888},
							val:        "true",
							ignoreCase: false,
						},
					},
					&actionExpr{
						pos: position{line: 544, col: 3, offset: 15973},
						run: (*parser).callonConst21,
						expr: &litMatcher{
							pos:        position{line: 544, col: 3, offset: 15973},
							val:        "false",
							ignoreCase: false,
						},
					},
					&ruleRefExpr{
						pos:  position{line: 547, col: 3, offset: 16059},
						name: "Unit",
					},
					&actionExpr{
						pos: position{line: 548, col: 3, offset: 16066},
						run: (*parser).callonConst24,
						expr: &seqExpr{
							pos: position{line: 548, col: 3, offset: 16066},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 548, col: 3, offset: 16066},
									val:        "\"",
									ignoreCase: false,
								},
								&zeroOrMoreExpr{
									pos: position{line: 548, col: 7, offset: 16070},
									expr: &seqExpr{
										pos: position{line: 548, col: 8, offset: 16071},
										exprs: []interface{}{
											&notExpr{
												pos: position{line: 548, col: 8, offset: 16071},
												expr: &ruleRefExpr{
													pos:  position{line: 548, col: 9, offset: 16072},
													name: "EscapedChar",
												},
											},
											&anyMatcher{
												line: 548, col: 21, offset: 16084,
											},
										},
									},
								},
								&litMatcher{
									pos:        position{line: 548, col: 25, offset: 16088},
									val:        "\"",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 555, col: 3, offset: 16272},
						run: (*parser).callonConst33,
						expr: &seqExpr{
							pos: position{line: 555, col: 3, offset: 16272},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 555, col: 3, offset: 16272},
									val:        "'",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 555, col: 7, offset: 16276},
									label: "val",
									expr: &seqExpr{
										pos: position{line: 555, col: 12, offset: 16281},
										exprs: []interface{}{
											&notExpr{
												pos: position{line: 555, col: 12, offset: 16281},
												expr: &ruleRefExpr{
													pos:  position{line: 555, col: 13, offset: 16282},
													name: "EscapedChar",
												},
											},
											&anyMatcher{
												line: 555, col: 25, offset: 16294,
											},
										},
									},
								},
								&litMatcher{
									pos:        position{line: 555, col: 28, offset: 16297},
									val:        "'",
									ignoreCase: false,
								},
							},
						},
					},
					&ruleRefExpr{
						pos:  position{line: 557, col: 5, offset: 16389},
						name: "ArrayLiteral",
					},
					&ruleRefExpr{
						pos:  position{line: 557, col: 20, offset: 16404},
						name: "RecordInstance",
					},
					&ruleRefExpr{
						pos:  position{line: 557, col: 37, offset: 16421},
						name: "VariantInstance",
					},
				},
			},
		},
		{
			name: "Unit",
			pos:  position{line: 559, col: 1, offset: 16438},
			expr: &actionExpr{
				pos: position{line: 559, col: 8, offset: 16445},
				run: (*parser).callonUnit1,
				expr: &litMatcher{
					pos:        position{line: 559, col: 8, offset: 16445},
					val:        "()",
					ignoreCase: false,
				},
			},
		},
		{
			name: "Unused",
			pos:  position{line: 563, col: 1, offset: 16508},
			expr: &actionExpr{
				pos: position{line: 563, col: 10, offset: 16517},
				run: (*parser).callonUnused1,
				expr: &litMatcher{
					pos:        position{line: 563, col: 11, offset: 16518},
					val:        "_",
					ignoreCase: false,
				},
			},
		},
		{
			name: "Reserved",
			pos:  position{line: 567, col: 1, offset: 16573},
			expr: &seqExpr{
				pos: position{line: 567, col: 12, offset: 16584},
				exprs: []interface{}{
					&choiceExpr{
						pos: position{line: 567, col: 13, offset: 16585},
						alternatives: []interface{}{
							&litMatcher{
								pos:        position{line: 567, col: 13, offset: 16585},
								val:        "let",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 567, col: 21, offset: 16593},
								val:        "if",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 567, col: 28, offset: 16600},
								val:        "else",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 567, col: 37, offset: 16609},
								val:        "func",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 567, col: 46, offset: 16618},
								val:        "type",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 567, col: 55, offset: 16627},
								val:        "true",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 567, col: 64, offset: 16636},
								val:        "false",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 567, col: 74, offset: 16646},
								val:        "mutable",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 567, col: 86, offset: 16658},
								val:        "type",
								ignoreCase: false,
							},
							&ruleRefExpr{
								pos:  position{line: 567, col: 95, offset: 16667},
								name: "BaseType",
							},
						},
					},
					&notExpr{
						pos: position{line: 567, col: 105, offset: 16677},
						expr: &oneOrMoreExpr{
							pos: position{line: 567, col: 106, offset: 16678},
							expr: &charClassMatcher{
								pos:        position{line: 567, col: 106, offset: 16678},
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
			pos:  position{line: 569, col: 1, offset: 16686},
			expr: &actionExpr{
				pos: position{line: 569, col: 12, offset: 16697},
				run: (*parser).callonBaseType1,
				expr: &choiceExpr{
					pos: position{line: 569, col: 14, offset: 16699},
					alternatives: []interface{}{
						&litMatcher{
							pos:        position{line: 569, col: 14, offset: 16699},
							val:        "int",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 569, col: 22, offset: 16707},
							val:        "bool",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 569, col: 31, offset: 16716},
							val:        "string",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 569, col: 42, offset: 16727},
							val:        "byte",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 569, col: 51, offset: 16736},
							val:        "rune",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 569, col: 60, offset: 16745},
							val:        "float",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 569, col: 70, offset: 16755},
							val:        "list",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "EscapedChar",
			pos:  position{line: 573, col: 1, offset: 16854},
			expr: &charClassMatcher{
				pos:        position{line: 573, col: 15, offset: 16868},
				val:        "[\\x00-\\x1f\"\\\\]",
				chars:      []rune{'"', '\\'},
				ranges:     []rune{'\x00', '\x1f'},
				ignoreCase: false,
				inverted:   false,
			},
		},
		{
			name: "EscapeSequence",
			pos:  position{line: 575, col: 1, offset: 16884},
			expr: &choiceExpr{
				pos: position{line: 575, col: 18, offset: 16901},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 575, col: 18, offset: 16901},
						name: "SingleCharEscape",
					},
					&ruleRefExpr{
						pos:  position{line: 575, col: 37, offset: 16920},
						name: "UnicodeEscape",
					},
				},
			},
		},
		{
			name: "SingleCharEscape",
			pos:  position{line: 577, col: 1, offset: 16935},
			expr: &charClassMatcher{
				pos:        position{line: 577, col: 20, offset: 16954},
				val:        "[\"\\\\/bfnrt]",
				chars:      []rune{'"', '\\', '/', 'b', 'f', 'n', 'r', 't'},
				ignoreCase: false,
				inverted:   false,
			},
		},
		{
			name: "DecimalDigit",
			pos:  position{line: 579, col: 1, offset: 16967},
			expr: &charClassMatcher{
				pos:        position{line: 579, col: 16, offset: 16982},
				val:        "[0-9]",
				ranges:     []rune{'0', '9'},
				ignoreCase: false,
				inverted:   false,
			},
		},
		{
			name: "NonZeroDecimalDigit",
			pos:  position{line: 581, col: 1, offset: 16989},
			expr: &charClassMatcher{
				pos:        position{line: 581, col: 23, offset: 17011},
				val:        "[1-9]",
				ranges:     []rune{'1', '9'},
				ignoreCase: false,
				inverted:   false,
			},
		},
		{
			name: "HexDigit",
			pos:  position{line: 583, col: 1, offset: 17018},
			expr: &charClassMatcher{
				pos:        position{line: 583, col: 12, offset: 17029},
				val:        "[0-9a-f]i",
				ranges:     []rune{'0', '9', 'a', 'f'},
				ignoreCase: true,
				inverted:   false,
			},
		},
		{
			name:        "__",
			displayName: "\"sigwhitespace\"",
			pos:         position{line: 585, col: 1, offset: 17040},
			expr: &choiceExpr{
				pos: position{line: 585, col: 22, offset: 17061},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 585, col: 22, offset: 17061},
						name: "Comment",
					},
					&oneOrMoreExpr{
						pos: position{line: 585, col: 32, offset: 17071},
						expr: &charClassMatcher{
							pos:        position{line: 585, col: 32, offset: 17071},
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
			pos:         position{line: 587, col: 1, offset: 17083},
			expr: &zeroOrMoreExpr{
				pos: position{line: 587, col: 18, offset: 17100},
				expr: &charClassMatcher{
					pos:        position{line: 587, col: 18, offset: 17100},
					val:        "[ \\r\\n\\t]",
					chars:      []rune{' ', '\r', '\n', '\t'},
					ignoreCase: false,
					inverted:   false,
				},
			},
		},
		{
			name: "EOF",
			pos:  position{line: 589, col: 1, offset: 17112},
			expr: &notExpr{
				pos: position{line: 589, col: 7, offset: 17118},
				expr: &anyMatcher{
					line: 589, col: 8, offset: 17119,
				},
			},
		},
	},
}

func (c *current) onModule1(stat, rest interface{}) (interface{}, error) {
	//fmt.Println("beginning module")
	vals := rest.([]interface{})
	if len(vals) > 0 {
		//fmt.Println("multiple statements")
		subvalues := []Ast{stat.(Ast)}
		for _, el := range vals {
			subvalues = append(subvalues, el.(Ast))
		}
		return Module{Name: "", Subvalues: subvalues}, nil
	} else {
		return Module{Name: "", Subvalues: []Ast{stat.(Ast)}}, nil
	}
}

func (p *parser) callonModule1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onModule1(stack["stat"], stack["rest"])
}

func (c *current) onExprLine1(e interface{}) (interface{}, error) {
	return e, nil
}

func (p *parser) callonExprLine1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onExprLine1(stack["e"])
}

func (c *current) onComment1(comment interface{}) (interface{}, error) {
	//fmt.Println("comment:", string(c.text))
	return Comment{StringValue: string(c.text[1:])}, nil
}

func (p *parser) callonComment1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onComment1(stack["comment"])
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

func (c *current) onFuncDefn1(i, ids, statements interface{}) (interface{}, error) {
	//fmt.Println(string(c.text))
	subvalues := []Ast{}
	args := []Ast{}
	vals := statements.([]interface{})
	if len(vals) > 0 {
		for _, el := range vals {
			subvalues = append(subvalues, el.(Ast))
		}
	}
	vals = ids.([]interface{})
	if len(vals) > 0 {
		restSl := toIfaceSlice(ids)
		for _, v := range restSl {
			// we can get each item in the grammar by index
			restExpr := toIfaceSlice(v)
			v := restExpr[0].(Ast)
			args = append(args, v)
		}
	}
	return Func{Name: i.(Identifier).StringValue, Arguments: args, Subvalues: subvalues}, nil
}

func (p *parser) callonFuncDefn1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onFuncDefn1(stack["i"], stack["ids"], stack["statements"])
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

func (c *current) onIfExpr2(expr, thens, elseifs interface{}) (interface{}, error) {
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

func (p *parser) callonIfExpr2() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onIfExpr2(stack["expr"], stack["thens"], stack["elseifs"])
}

func (c *current) onIfExpr22(expr, thens, elses interface{}) (interface{}, error) {
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

func (p *parser) callonIfExpr22() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onIfExpr22(stack["expr"], stack["thens"], stack["elses"])
}

func (c *current) onIfExpr47(expr, thens interface{}) (interface{}, error) {
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

func (p *parser) callonIfExpr47() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onIfExpr47(stack["expr"], stack["thens"])
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

	return Call{Module: module.(Ast), Function: fn.(Ast), Arguments: arguments}, nil
}

func (p *parser) callonCall2() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onCall2(stack["module"], stack["fn"], stack["args"])
}

func (c *current) onCall12(fn, args interface{}) (interface{}, error) {
	//fmt.Println("call", string(c.text))
	arguments := []Ast{}

	switch args.(type) {
	case Container:
		arguments = args.(Container).Subvalues
	default:
		arguments = []Ast{}
	}

	return Call{Module: nil, Function: fn.(Ast), Arguments: arguments}, nil
}

func (p *parser) callonCall12() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onCall12(stack["fn"], stack["args"])
}

func (c *current) onCall19(fn interface{}) (interface{}, error) {
	//fmt.Println("call", string(c.text))
	arguments := []Ast{}

	return Call{Module: nil, Function: fn.(Ast), Arguments: arguments}, nil
}

func (p *parser) callonCall19() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onCall19(stack["fn"])
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

func (c *current) onConst24() (interface{}, error) {
	val, err := strconv.Unquote(string(c.text))
	if err == nil {
		return BasicAst{Type: "String", StringValue: val, ValueType: STRING}, nil
	}
	return nil, err
}

func (p *parser) callonConst24() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onConst24()
}

func (c *current) onConst33(val interface{}) (interface{}, error) {
	return BasicAst{Type: "Char", CharValue: rune(c.text[1]), ValueType: CHAR}, nil
}

func (p *parser) callonConst33() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onConst33(stack["val"])
}

func (c *current) onUnit1() (interface{}, error) {
	return BasicAst{Type: "Unit", ValueType: NIL}, nil
}

func (p *parser) callonUnit1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onUnit1()
}

func (c *current) onUnused1() (interface{}, error) {
	return Identifier{StringValue: "_"}, nil
}

func (p *parser) callonUnused1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onUnused1()
}

func (c *current) onBaseType1() (interface{}, error) {
	return BasicAst{Type: "Type", StringValue: string(c.text), ValueType: STRING}, nil
}

func (p *parser) callonBaseType1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onBaseType1()
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
