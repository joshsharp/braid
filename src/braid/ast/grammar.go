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
						name: "Call",
					},
					&ruleRefExpr{
						pos:  position{line: 30, col: 54, offset: 737},
						name: "ExprLine",
					},
				},
			},
		},
		{
			name: "ExprLine",
			pos:  position{line: 32, col: 1, offset: 747},
			expr: &actionExpr{
				pos: position{line: 32, col: 12, offset: 758},
				run: (*parser).callonExprLine1,
				expr: &seqExpr{
					pos: position{line: 32, col: 12, offset: 758},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 32, col: 12, offset: 758},
							label: "e",
							expr: &ruleRefExpr{
								pos:  position{line: 32, col: 14, offset: 760},
								name: "Expr",
							},
						},
						&andExpr{
							pos: position{line: 32, col: 19, offset: 765},
							expr: &litMatcher{
								pos:        position{line: 32, col: 20, offset: 766},
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
			pos:  position{line: 36, col: 1, offset: 794},
			expr: &actionExpr{
				pos: position{line: 36, col: 11, offset: 804},
				run: (*parser).callonComment1,
				expr: &seqExpr{
					pos: position{line: 36, col: 11, offset: 804},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 36, col: 11, offset: 804},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 36, col: 13, offset: 806},
							val:        "#",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 36, col: 17, offset: 810},
							label: "comment",
							expr: &zeroOrMoreExpr{
								pos: position{line: 36, col: 25, offset: 818},
								expr: &seqExpr{
									pos: position{line: 36, col: 26, offset: 819},
									exprs: []interface{}{
										&notExpr{
											pos: position{line: 36, col: 26, offset: 819},
											expr: &ruleRefExpr{
												pos:  position{line: 36, col: 27, offset: 820},
												name: "EscapedChar",
											},
										},
										&anyMatcher{
											line: 36, col: 39, offset: 832,
										},
									},
								},
							},
						},
						&andExpr{
							pos: position{line: 36, col: 43, offset: 836},
							expr: &litMatcher{
								pos:        position{line: 36, col: 44, offset: 837},
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
			pos:  position{line: 41, col: 1, offset: 987},
			expr: &choiceExpr{
				pos: position{line: 41, col: 12, offset: 998},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 41, col: 12, offset: 998},
						run: (*parser).callonTypeDefn2,
						expr: &seqExpr{
							pos: position{line: 41, col: 12, offset: 998},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 41, col: 12, offset: 998},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 41, col: 14, offset: 1000},
									val:        "type",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 41, col: 21, offset: 1007},
									name: "__",
								},
								&labeledExpr{
									pos:   position{line: 41, col: 24, offset: 1010},
									label: "name",
									expr: &ruleRefExpr{
										pos:  position{line: 41, col: 29, offset: 1015},
										name: "ModuleName",
									},
								},
								&labeledExpr{
									pos:   position{line: 41, col: 40, offset: 1026},
									label: "params",
									expr: &zeroOrMoreExpr{
										pos: position{line: 41, col: 47, offset: 1033},
										expr: &seqExpr{
											pos: position{line: 41, col: 48, offset: 1034},
											exprs: []interface{}{
												&ruleRefExpr{
													pos:  position{line: 41, col: 48, offset: 1034},
													name: "__",
												},
												&ruleRefExpr{
													pos:  position{line: 41, col: 51, offset: 1037},
													name: "TypeParameter",
												},
											},
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 41, col: 67, offset: 1053},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 41, col: 69, offset: 1055},
									val:        "=",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 41, col: 73, offset: 1059},
									label: "types",
									expr: &oneOrMoreExpr{
										pos: position{line: 41, col: 79, offset: 1065},
										expr: &seqExpr{
											pos: position{line: 41, col: 80, offset: 1066},
											exprs: []interface{}{
												&ruleRefExpr{
													pos:  position{line: 41, col: 80, offset: 1066},
													name: "__",
												},
												&ruleRefExpr{
													pos:  position{line: 41, col: 83, offset: 1069},
													name: "AnyType",
												},
											},
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 41, col: 93, offset: 1079},
									name: "__",
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 60, col: 1, offset: 1573},
						run: (*parser).callonTypeDefn22,
						expr: &seqExpr{
							pos: position{line: 60, col: 1, offset: 1573},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 60, col: 1, offset: 1573},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 60, col: 3, offset: 1575},
									val:        "type",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 60, col: 10, offset: 1582},
									name: "__",
								},
								&labeledExpr{
									pos:   position{line: 60, col: 13, offset: 1585},
									label: "name",
									expr: &ruleRefExpr{
										pos:  position{line: 60, col: 18, offset: 1590},
										name: "ModuleName",
									},
								},
								&labeledExpr{
									pos:   position{line: 60, col: 29, offset: 1601},
									label: "params",
									expr: &zeroOrMoreExpr{
										pos: position{line: 60, col: 36, offset: 1608},
										expr: &seqExpr{
											pos: position{line: 60, col: 37, offset: 1609},
											exprs: []interface{}{
												&ruleRefExpr{
													pos:  position{line: 60, col: 37, offset: 1609},
													name: "__",
												},
												&ruleRefExpr{
													pos:  position{line: 60, col: 40, offset: 1612},
													name: "TypeParameter",
												},
											},
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 60, col: 56, offset: 1628},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 60, col: 58, offset: 1630},
									val:        "=",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 60, col: 62, offset: 1634},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 61, col: 5, offset: 1640},
									val:        "{",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 61, col: 9, offset: 1644},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 61, col: 11, offset: 1646},
									label: "first",
									expr: &ruleRefExpr{
										pos:  position{line: 61, col: 17, offset: 1652},
										name: "RecordFieldDefn",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 61, col: 33, offset: 1668},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 61, col: 35, offset: 1670},
									label: "rest",
									expr: &zeroOrMoreExpr{
										pos: position{line: 61, col: 40, offset: 1675},
										expr: &seqExpr{
											pos: position{line: 61, col: 41, offset: 1676},
											exprs: []interface{}{
												&litMatcher{
													pos:        position{line: 61, col: 41, offset: 1676},
													val:        ",",
													ignoreCase: false,
												},
												&ruleRefExpr{
													pos:  position{line: 61, col: 45, offset: 1680},
													name: "_",
												},
												&ruleRefExpr{
													pos:  position{line: 61, col: 47, offset: 1682},
													name: "RecordFieldDefn",
												},
												&ruleRefExpr{
													pos:  position{line: 61, col: 63, offset: 1698},
													name: "_",
												},
											},
										},
									},
								},
								&zeroOrOneExpr{
									pos: position{line: 61, col: 67, offset: 1702},
									expr: &litMatcher{
										pos:        position{line: 61, col: 67, offset: 1702},
										val:        ",",
										ignoreCase: false,
									},
								},
								&ruleRefExpr{
									pos:  position{line: 61, col: 72, offset: 1707},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 61, col: 74, offset: 1709},
									val:        "}",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 61, col: 78, offset: 1713},
									name: "__",
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 79, col: 1, offset: 2196},
						run: (*parser).callonTypeDefn54,
						expr: &seqExpr{
							pos: position{line: 79, col: 1, offset: 2196},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 79, col: 1, offset: 2196},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 79, col: 3, offset: 2198},
									val:        "type",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 79, col: 10, offset: 2205},
									name: "__",
								},
								&labeledExpr{
									pos:   position{line: 79, col: 13, offset: 2208},
									label: "name",
									expr: &ruleRefExpr{
										pos:  position{line: 79, col: 18, offset: 2213},
										name: "ModuleName",
									},
								},
								&labeledExpr{
									pos:   position{line: 79, col: 29, offset: 2224},
									label: "params",
									expr: &zeroOrMoreExpr{
										pos: position{line: 79, col: 36, offset: 2231},
										expr: &seqExpr{
											pos: position{line: 79, col: 37, offset: 2232},
											exprs: []interface{}{
												&ruleRefExpr{
													pos:  position{line: 79, col: 37, offset: 2232},
													name: "__",
												},
												&ruleRefExpr{
													pos:  position{line: 79, col: 40, offset: 2235},
													name: "TypeParameter",
												},
											},
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 79, col: 56, offset: 2251},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 79, col: 58, offset: 2253},
									val:        "=",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 79, col: 62, offset: 2257},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 79, col: 64, offset: 2259},
									label: "rest",
									expr: &oneOrMoreExpr{
										pos: position{line: 79, col: 69, offset: 2264},
										expr: &ruleRefExpr{
											pos:  position{line: 79, col: 70, offset: 2265},
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
			pos:  position{line: 94, col: 1, offset: 2670},
			expr: &actionExpr{
				pos: position{line: 94, col: 19, offset: 2688},
				run: (*parser).callonRecordFieldDefn1,
				expr: &seqExpr{
					pos: position{line: 94, col: 19, offset: 2688},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 94, col: 19, offset: 2688},
							label: "name",
							expr: &ruleRefExpr{
								pos:  position{line: 94, col: 24, offset: 2693},
								name: "VariableName",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 94, col: 37, offset: 2706},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 94, col: 39, offset: 2708},
							val:        ":",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 94, col: 43, offset: 2712},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 94, col: 45, offset: 2714},
							label: "t",
							expr: &ruleRefExpr{
								pos:  position{line: 94, col: 48, offset: 2717},
								name: "AnyType",
							},
						},
					},
				},
			},
		},
		{
			name: "VariantConstructor",
			pos:  position{line: 98, col: 1, offset: 2809},
			expr: &choiceExpr{
				pos: position{line: 98, col: 22, offset: 2830},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 98, col: 22, offset: 2830},
						run: (*parser).callonVariantConstructor2,
						expr: &seqExpr{
							pos: position{line: 98, col: 22, offset: 2830},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 98, col: 22, offset: 2830},
									val:        "|",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 98, col: 26, offset: 2834},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 98, col: 28, offset: 2836},
									label: "name",
									expr: &ruleRefExpr{
										pos:  position{line: 98, col: 33, offset: 2841},
										name: "ModuleName",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 98, col: 44, offset: 2852},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 98, col: 46, offset: 2854},
									val:        "{",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 98, col: 50, offset: 2858},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 98, col: 52, offset: 2860},
									label: "first",
									expr: &ruleRefExpr{
										pos:  position{line: 98, col: 58, offset: 2866},
										name: "RecordFieldDefn",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 98, col: 74, offset: 2882},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 98, col: 76, offset: 2884},
									label: "rest",
									expr: &zeroOrMoreExpr{
										pos: position{line: 98, col: 81, offset: 2889},
										expr: &seqExpr{
											pos: position{line: 98, col: 82, offset: 2890},
											exprs: []interface{}{
												&litMatcher{
													pos:        position{line: 98, col: 82, offset: 2890},
													val:        ",",
													ignoreCase: false,
												},
												&ruleRefExpr{
													pos:  position{line: 98, col: 86, offset: 2894},
													name: "_",
												},
												&ruleRefExpr{
													pos:  position{line: 98, col: 88, offset: 2896},
													name: "RecordFieldDefn",
												},
												&ruleRefExpr{
													pos:  position{line: 98, col: 104, offset: 2912},
													name: "_",
												},
											},
										},
									},
								},
								&zeroOrOneExpr{
									pos: position{line: 98, col: 108, offset: 2916},
									expr: &litMatcher{
										pos:        position{line: 98, col: 108, offset: 2916},
										val:        ",",
										ignoreCase: false,
									},
								},
								&ruleRefExpr{
									pos:  position{line: 98, col: 113, offset: 2921},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 98, col: 115, offset: 2923},
									val:        "}",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 98, col: 119, offset: 2927},
									name: "__",
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 117, col: 1, offset: 3528},
						run: (*parser).callonVariantConstructor26,
						expr: &seqExpr{
							pos: position{line: 117, col: 1, offset: 3528},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 117, col: 1, offset: 3528},
									val:        "|",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 117, col: 5, offset: 3532},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 117, col: 7, offset: 3534},
									label: "name",
									expr: &ruleRefExpr{
										pos:  position{line: 117, col: 12, offset: 3539},
										name: "ModuleName",
									},
								},
								&labeledExpr{
									pos:   position{line: 117, col: 23, offset: 3550},
									label: "rest",
									expr: &zeroOrMoreExpr{
										pos: position{line: 117, col: 28, offset: 3555},
										expr: &seqExpr{
											pos: position{line: 117, col: 29, offset: 3556},
											exprs: []interface{}{
												&ruleRefExpr{
													pos:  position{line: 117, col: 29, offset: 3556},
													name: "__",
												},
												&ruleRefExpr{
													pos:  position{line: 117, col: 32, offset: 3559},
													name: "AnyType",
												},
											},
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 117, col: 42, offset: 3569},
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
			pos:  position{line: 134, col: 1, offset: 4004},
			expr: &choiceExpr{
				pos: position{line: 134, col: 11, offset: 4014},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 134, col: 11, offset: 4014},
						name: "BaseType",
					},
					&ruleRefExpr{
						pos:  position{line: 134, col: 22, offset: 4025},
						name: "TypeParameter",
					},
				},
			},
		},
		{
			name: "Assignment",
			pos:  position{line: 136, col: 1, offset: 4040},
			expr: &choiceExpr{
				pos: position{line: 136, col: 14, offset: 4053},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 136, col: 14, offset: 4053},
						run: (*parser).callonAssignment2,
						expr: &seqExpr{
							pos: position{line: 136, col: 14, offset: 4053},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 136, col: 14, offset: 4053},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 136, col: 16, offset: 4055},
									val:        "let",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 136, col: 22, offset: 4061},
									name: "__",
								},
								&labeledExpr{
									pos:   position{line: 136, col: 25, offset: 4064},
									label: "i",
									expr: &ruleRefExpr{
										pos:  position{line: 136, col: 27, offset: 4066},
										name: "Assignable",
									},
								},
								&labeledExpr{
									pos:   position{line: 136, col: 38, offset: 4077},
									label: "rest",
									expr: &zeroOrMoreExpr{
										pos: position{line: 136, col: 43, offset: 4082},
										expr: &seqExpr{
											pos: position{line: 136, col: 44, offset: 4083},
											exprs: []interface{}{
												&litMatcher{
													pos:        position{line: 136, col: 44, offset: 4083},
													val:        ",",
													ignoreCase: false,
												},
												&ruleRefExpr{
													pos:  position{line: 136, col: 48, offset: 4087},
													name: "_",
												},
												&ruleRefExpr{
													pos:  position{line: 136, col: 50, offset: 4089},
													name: "Assignable",
												},
											},
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 136, col: 63, offset: 4102},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 136, col: 65, offset: 4104},
									val:        "=",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 136, col: 69, offset: 4108},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 136, col: 71, offset: 4110},
									label: "expr",
									expr: &ruleRefExpr{
										pos:  position{line: 136, col: 76, offset: 4115},
										name: "Expr",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 136, col: 81, offset: 4120},
									name: "_",
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 151, col: 1, offset: 4559},
						run: (*parser).callonAssignment21,
						expr: &seqExpr{
							pos: position{line: 151, col: 1, offset: 4559},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 151, col: 1, offset: 4559},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 151, col: 3, offset: 4561},
									val:        "let",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 151, col: 9, offset: 4567},
									name: "__",
								},
								&notExpr{
									pos: position{line: 151, col: 12, offset: 4570},
									expr: &ruleRefExpr{
										pos:  position{line: 151, col: 13, offset: 4571},
										name: "Assignable",
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 155, col: 1, offset: 4679},
						run: (*parser).callonAssignment28,
						expr: &seqExpr{
							pos: position{line: 155, col: 1, offset: 4679},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 155, col: 1, offset: 4679},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 155, col: 3, offset: 4681},
									val:        "let",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 155, col: 9, offset: 4687},
									name: "__",
								},
								&labeledExpr{
									pos:   position{line: 155, col: 12, offset: 4690},
									label: "i",
									expr: &ruleRefExpr{
										pos:  position{line: 155, col: 14, offset: 4692},
										name: "Assignable",
									},
								},
								&labeledExpr{
									pos:   position{line: 155, col: 25, offset: 4703},
									label: "rest",
									expr: &zeroOrMoreExpr{
										pos: position{line: 155, col: 30, offset: 4708},
										expr: &seqExpr{
											pos: position{line: 155, col: 31, offset: 4709},
											exprs: []interface{}{
												&litMatcher{
													pos:        position{line: 155, col: 31, offset: 4709},
													val:        ",",
													ignoreCase: false,
												},
												&ruleRefExpr{
													pos:  position{line: 155, col: 35, offset: 4713},
													name: "_",
												},
												&ruleRefExpr{
													pos:  position{line: 155, col: 37, offset: 4715},
													name: "Assignable",
												},
											},
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 155, col: 50, offset: 4728},
									name: "_",
								},
								&notExpr{
									pos: position{line: 155, col: 52, offset: 4730},
									expr: &litMatcher{
										pos:        position{line: 155, col: 53, offset: 4731},
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
			pos:  position{line: 159, col: 1, offset: 4825},
			expr: &actionExpr{
				pos: position{line: 159, col: 12, offset: 4836},
				run: (*parser).callonFuncDefn1,
				expr: &seqExpr{
					pos: position{line: 159, col: 12, offset: 4836},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 159, col: 12, offset: 4836},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 159, col: 14, offset: 4838},
							val:        "let",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 159, col: 20, offset: 4844},
							name: "__",
						},
						&labeledExpr{
							pos:   position{line: 159, col: 23, offset: 4847},
							label: "i",
							expr: &ruleRefExpr{
								pos:  position{line: 159, col: 25, offset: 4849},
								name: "VariableName",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 159, col: 38, offset: 4862},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 159, col: 40, offset: 4864},
							val:        "=",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 159, col: 44, offset: 4868},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 159, col: 46, offset: 4870},
							val:        "func",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 159, col: 53, offset: 4877},
							name: "__",
						},
						&labeledExpr{
							pos:   position{line: 159, col: 56, offset: 4880},
							label: "ids",
							expr: &zeroOrMoreExpr{
								pos: position{line: 159, col: 60, offset: 4884},
								expr: &seqExpr{
									pos: position{line: 159, col: 61, offset: 4885},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 159, col: 61, offset: 4885},
											name: "VariableName",
										},
										&ruleRefExpr{
											pos:  position{line: 159, col: 74, offset: 4898},
											name: "__",
										},
									},
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 159, col: 79, offset: 4903},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 159, col: 81, offset: 4905},
							val:        "{",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 159, col: 85, offset: 4909},
							name: "__",
						},
						&labeledExpr{
							pos:   position{line: 159, col: 88, offset: 4912},
							label: "statements",
							expr: &oneOrMoreExpr{
								pos: position{line: 159, col: 99, offset: 4923},
								expr: &ruleRefExpr{
									pos:  position{line: 159, col: 100, offset: 4924},
									name: "Statement",
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 159, col: 112, offset: 4936},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 159, col: 114, offset: 4938},
							val:        "}",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 159, col: 118, offset: 4942},
							name: "_",
						},
					},
				},
			},
		},
		{
			name: "Expr",
			pos:  position{line: 182, col: 1, offset: 5596},
			expr: &actionExpr{
				pos: position{line: 182, col: 8, offset: 5603},
				run: (*parser).callonExpr1,
				expr: &labeledExpr{
					pos:   position{line: 182, col: 8, offset: 5603},
					label: "ex",
					expr: &choiceExpr{
						pos: position{line: 182, col: 12, offset: 5607},
						alternatives: []interface{}{
							&ruleRefExpr{
								pos:  position{line: 182, col: 12, offset: 5607},
								name: "IfExpr",
							},
							&ruleRefExpr{
								pos:  position{line: 182, col: 21, offset: 5616},
								name: "Call",
							},
							&ruleRefExpr{
								pos:  position{line: 182, col: 28, offset: 5623},
								name: "CompoundExpr",
							},
						},
					},
				},
			},
		},
		{
			name: "IfExpr",
			pos:  position{line: 188, col: 1, offset: 5740},
			expr: &choiceExpr{
				pos: position{line: 188, col: 10, offset: 5749},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 188, col: 10, offset: 5749},
						run: (*parser).callonIfExpr2,
						expr: &seqExpr{
							pos: position{line: 188, col: 10, offset: 5749},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 188, col: 10, offset: 5749},
									val:        "if",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 188, col: 15, offset: 5754},
									name: "__",
								},
								&labeledExpr{
									pos:   position{line: 188, col: 18, offset: 5757},
									label: "expr",
									expr: &ruleRefExpr{
										pos:  position{line: 188, col: 23, offset: 5762},
										name: "BinOpBool",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 188, col: 33, offset: 5772},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 188, col: 35, offset: 5774},
									val:        "{",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 188, col: 39, offset: 5778},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 188, col: 41, offset: 5780},
									label: "thens",
									expr: &oneOrMoreExpr{
										pos: position{line: 188, col: 47, offset: 5786},
										expr: &ruleRefExpr{
											pos:  position{line: 188, col: 48, offset: 5787},
											name: "Statement",
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 188, col: 60, offset: 5799},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 188, col: 62, offset: 5801},
									val:        "}",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 188, col: 66, offset: 5805},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 188, col: 68, offset: 5807},
									val:        "else",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 188, col: 75, offset: 5814},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 188, col: 77, offset: 5816},
									label: "elseifs",
									expr: &ruleRefExpr{
										pos:  position{line: 188, col: 85, offset: 5824},
										name: "IfExpr",
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 200, col: 1, offset: 6154},
						run: (*parser).callonIfExpr21,
						expr: &seqExpr{
							pos: position{line: 200, col: 1, offset: 6154},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 200, col: 1, offset: 6154},
									val:        "if",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 200, col: 6, offset: 6159},
									name: "__",
								},
								&labeledExpr{
									pos:   position{line: 200, col: 9, offset: 6162},
									label: "expr",
									expr: &ruleRefExpr{
										pos:  position{line: 200, col: 14, offset: 6167},
										name: "BinOpBool",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 200, col: 24, offset: 6177},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 200, col: 26, offset: 6179},
									val:        "{",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 200, col: 30, offset: 6183},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 200, col: 32, offset: 6185},
									label: "thens",
									expr: &oneOrMoreExpr{
										pos: position{line: 200, col: 38, offset: 6191},
										expr: &ruleRefExpr{
											pos:  position{line: 200, col: 39, offset: 6192},
											name: "Statement",
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 200, col: 51, offset: 6204},
									name: "__",
								},
								&litMatcher{
									pos:        position{line: 200, col: 54, offset: 6207},
									val:        "}",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 200, col: 58, offset: 6211},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 200, col: 60, offset: 6213},
									val:        "else",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 200, col: 67, offset: 6220},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 200, col: 69, offset: 6222},
									val:        "{",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 200, col: 73, offset: 6226},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 200, col: 75, offset: 6228},
									label: "elses",
									expr: &oneOrMoreExpr{
										pos: position{line: 200, col: 81, offset: 6234},
										expr: &ruleRefExpr{
											pos:  position{line: 200, col: 82, offset: 6235},
											name: "Statement",
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 200, col: 94, offset: 6247},
									name: "__",
								},
								&litMatcher{
									pos:        position{line: 200, col: 97, offset: 6250},
									val:        "}",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 219, col: 1, offset: 6753},
						run: (*parser).callonIfExpr45,
						expr: &seqExpr{
							pos: position{line: 219, col: 1, offset: 6753},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 219, col: 1, offset: 6753},
									val:        "if",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 219, col: 6, offset: 6758},
									name: "__",
								},
								&labeledExpr{
									pos:   position{line: 219, col: 9, offset: 6761},
									label: "expr",
									expr: &ruleRefExpr{
										pos:  position{line: 219, col: 14, offset: 6766},
										name: "BinOpBool",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 219, col: 24, offset: 6776},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 219, col: 26, offset: 6778},
									val:        "{",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 219, col: 30, offset: 6782},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 219, col: 32, offset: 6784},
									label: "thens",
									expr: &oneOrMoreExpr{
										pos: position{line: 219, col: 38, offset: 6790},
										expr: &ruleRefExpr{
											pos:  position{line: 219, col: 39, offset: 6791},
											name: "Statement",
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 219, col: 51, offset: 6803},
									name: "__",
								},
								&litMatcher{
									pos:        position{line: 219, col: 54, offset: 6806},
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
			pos:  position{line: 231, col: 1, offset: 7104},
			expr: &choiceExpr{
				pos: position{line: 231, col: 8, offset: 7111},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 231, col: 8, offset: 7111},
						run: (*parser).callonCall2,
						expr: &seqExpr{
							pos: position{line: 231, col: 8, offset: 7111},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 231, col: 8, offset: 7111},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 231, col: 10, offset: 7113},
									label: "module",
									expr: &ruleRefExpr{
										pos:  position{line: 231, col: 17, offset: 7120},
										name: "ModuleName",
									},
								},
								&litMatcher{
									pos:        position{line: 231, col: 28, offset: 7131},
									val:        ".",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 231, col: 32, offset: 7135},
									label: "fn",
									expr: &ruleRefExpr{
										pos:  position{line: 231, col: 35, offset: 7138},
										name: "VariableName",
									},
								},
								&labeledExpr{
									pos:   position{line: 231, col: 48, offset: 7151},
									label: "args",
									expr: &ruleRefExpr{
										pos:  position{line: 231, col: 53, offset: 7156},
										name: "Arguments",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 231, col: 63, offset: 7166},
									name: "_",
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 245, col: 1, offset: 7468},
						run: (*parser).callonCall13,
						expr: &seqExpr{
							pos: position{line: 245, col: 1, offset: 7468},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 245, col: 1, offset: 7468},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 245, col: 3, offset: 7470},
									label: "fn",
									expr: &ruleRefExpr{
										pos:  position{line: 245, col: 6, offset: 7473},
										name: "VariableName",
									},
								},
								&labeledExpr{
									pos:   position{line: 245, col: 19, offset: 7486},
									label: "args",
									expr: &ruleRefExpr{
										pos:  position{line: 245, col: 24, offset: 7491},
										name: "Arguments",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 245, col: 34, offset: 7501},
									name: "_",
								},
							},
						},
					},
				},
			},
		},
		{
			name: "Arguments",
			pos:  position{line: 259, col: 1, offset: 7793},
			expr: &choiceExpr{
				pos: position{line: 259, col: 13, offset: 7805},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 259, col: 13, offset: 7805},
						run: (*parser).callonArguments2,
						expr: &seqExpr{
							pos: position{line: 259, col: 13, offset: 7805},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 259, col: 13, offset: 7805},
									val:        "(",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 259, col: 17, offset: 7809},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 259, col: 19, offset: 7811},
									label: "argument",
									expr: &ruleRefExpr{
										pos:  position{line: 259, col: 28, offset: 7820},
										name: "BinOpParens",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 259, col: 40, offset: 7832},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 259, col: 42, offset: 7834},
									label: "rest",
									expr: &zeroOrMoreExpr{
										pos: position{line: 259, col: 47, offset: 7839},
										expr: &seqExpr{
											pos: position{line: 259, col: 48, offset: 7840},
											exprs: []interface{}{
												&litMatcher{
													pos:        position{line: 259, col: 48, offset: 7840},
													val:        ",",
													ignoreCase: false,
												},
												&ruleRefExpr{
													pos:  position{line: 259, col: 52, offset: 7844},
													name: "_",
												},
												&ruleRefExpr{
													pos:  position{line: 259, col: 54, offset: 7846},
													name: "BinOpParens",
												},
											},
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 259, col: 68, offset: 7860},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 259, col: 70, offset: 7862},
									val:        ")",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 276, col: 1, offset: 8284},
						run: (*parser).callonArguments17,
						expr: &seqExpr{
							pos: position{line: 276, col: 1, offset: 8284},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 276, col: 1, offset: 8284},
									val:        "(",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 276, col: 5, offset: 8288},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 276, col: 7, offset: 8290},
									label: "argument",
									expr: &ruleRefExpr{
										pos:  position{line: 276, col: 16, offset: 8299},
										name: "Expr",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 276, col: 21, offset: 8304},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 276, col: 23, offset: 8306},
									val:        ")",
									ignoreCase: false,
								},
							},
						},
					},
					&ruleRefExpr{
						pos:  position{line: 281, col: 1, offset: 8412},
						name: "Unit",
					},
				},
			},
		},
		{
			name: "CompoundExpr",
			pos:  position{line: 283, col: 1, offset: 8418},
			expr: &actionExpr{
				pos: position{line: 283, col: 16, offset: 8433},
				run: (*parser).callonCompoundExpr1,
				expr: &seqExpr{
					pos: position{line: 283, col: 16, offset: 8433},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 283, col: 16, offset: 8433},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 283, col: 18, offset: 8435},
							label: "op",
							expr: &ruleRefExpr{
								pos:  position{line: 283, col: 21, offset: 8438},
								name: "BinOp",
							},
						},
						&labeledExpr{
							pos:   position{line: 283, col: 27, offset: 8444},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 283, col: 32, offset: 8449},
								expr: &seqExpr{
									pos: position{line: 283, col: 33, offset: 8450},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 283, col: 33, offset: 8450},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 283, col: 36, offset: 8453},
											name: "Operator",
										},
										&ruleRefExpr{
											pos:  position{line: 283, col: 45, offset: 8462},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 283, col: 48, offset: 8465},
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
			pos:  position{line: 303, col: 1, offset: 9089},
			expr: &choiceExpr{
				pos: position{line: 303, col: 9, offset: 9097},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 303, col: 9, offset: 9097},
						name: "BinOpBool",
					},
					&ruleRefExpr{
						pos:  position{line: 303, col: 21, offset: 9109},
						name: "BinOpEquality",
					},
					&ruleRefExpr{
						pos:  position{line: 303, col: 37, offset: 9125},
						name: "BinOpLow",
					},
					&ruleRefExpr{
						pos:  position{line: 303, col: 48, offset: 9136},
						name: "BinOpHigh",
					},
					&ruleRefExpr{
						pos:  position{line: 303, col: 60, offset: 9148},
						name: "BinOpParens",
					},
				},
			},
		},
		{
			name: "BinOpBool",
			pos:  position{line: 305, col: 1, offset: 9161},
			expr: &actionExpr{
				pos: position{line: 305, col: 13, offset: 9173},
				run: (*parser).callonBinOpBool1,
				expr: &seqExpr{
					pos: position{line: 305, col: 13, offset: 9173},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 305, col: 13, offset: 9173},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 305, col: 15, offset: 9175},
							label: "first",
							expr: &ruleRefExpr{
								pos:  position{line: 305, col: 21, offset: 9181},
								name: "BinOpEquality",
							},
						},
						&labeledExpr{
							pos:   position{line: 305, col: 35, offset: 9195},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 305, col: 40, offset: 9200},
								expr: &seqExpr{
									pos: position{line: 305, col: 41, offset: 9201},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 305, col: 41, offset: 9201},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 305, col: 44, offset: 9204},
											name: "OperatorBoolean",
										},
										&ruleRefExpr{
											pos:  position{line: 305, col: 60, offset: 9220},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 305, col: 63, offset: 9223},
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
			pos:  position{line: 324, col: 1, offset: 9809},
			expr: &actionExpr{
				pos: position{line: 324, col: 17, offset: 9825},
				run: (*parser).callonBinOpEquality1,
				expr: &seqExpr{
					pos: position{line: 324, col: 17, offset: 9825},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 324, col: 17, offset: 9825},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 324, col: 19, offset: 9827},
							label: "first",
							expr: &ruleRefExpr{
								pos:  position{line: 324, col: 25, offset: 9833},
								name: "BinOpLow",
							},
						},
						&labeledExpr{
							pos:   position{line: 324, col: 34, offset: 9842},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 324, col: 39, offset: 9847},
								expr: &seqExpr{
									pos: position{line: 324, col: 40, offset: 9848},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 324, col: 40, offset: 9848},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 324, col: 43, offset: 9851},
											name: "OperatorEquality",
										},
										&ruleRefExpr{
											pos:  position{line: 324, col: 60, offset: 9868},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 324, col: 63, offset: 9871},
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
			pos:  position{line: 344, col: 1, offset: 10455},
			expr: &actionExpr{
				pos: position{line: 344, col: 12, offset: 10466},
				run: (*parser).callonBinOpLow1,
				expr: &seqExpr{
					pos: position{line: 344, col: 12, offset: 10466},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 344, col: 12, offset: 10466},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 344, col: 14, offset: 10468},
							label: "first",
							expr: &ruleRefExpr{
								pos:  position{line: 344, col: 20, offset: 10474},
								name: "BinOpHigh",
							},
						},
						&labeledExpr{
							pos:   position{line: 344, col: 30, offset: 10484},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 344, col: 35, offset: 10489},
								expr: &seqExpr{
									pos: position{line: 344, col: 36, offset: 10490},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 344, col: 36, offset: 10490},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 344, col: 39, offset: 10493},
											name: "OperatorLow",
										},
										&ruleRefExpr{
											pos:  position{line: 344, col: 51, offset: 10505},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 344, col: 54, offset: 10508},
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
			pos:  position{line: 364, col: 1, offset: 11089},
			expr: &actionExpr{
				pos: position{line: 364, col: 13, offset: 11101},
				run: (*parser).callonBinOpHigh1,
				expr: &seqExpr{
					pos: position{line: 364, col: 13, offset: 11101},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 364, col: 13, offset: 11101},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 364, col: 15, offset: 11103},
							label: "first",
							expr: &ruleRefExpr{
								pos:  position{line: 364, col: 21, offset: 11109},
								name: "BinOpParens",
							},
						},
						&labeledExpr{
							pos:   position{line: 364, col: 33, offset: 11121},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 364, col: 38, offset: 11126},
								expr: &seqExpr{
									pos: position{line: 364, col: 39, offset: 11127},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 364, col: 39, offset: 11127},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 364, col: 42, offset: 11130},
											name: "OperatorHigh",
										},
										&ruleRefExpr{
											pos:  position{line: 364, col: 55, offset: 11143},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 364, col: 58, offset: 11146},
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
			pos:  position{line: 383, col: 1, offset: 11729},
			expr: &choiceExpr{
				pos: position{line: 383, col: 15, offset: 11743},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 383, col: 15, offset: 11743},
						run: (*parser).callonBinOpParens2,
						expr: &seqExpr{
							pos: position{line: 383, col: 15, offset: 11743},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 383, col: 15, offset: 11743},
									val:        "(",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 383, col: 19, offset: 11747},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 383, col: 21, offset: 11749},
									label: "first",
									expr: &ruleRefExpr{
										pos:  position{line: 383, col: 27, offset: 11755},
										name: "BinOp",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 383, col: 33, offset: 11761},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 383, col: 35, offset: 11763},
									val:        ")",
									ignoreCase: false,
								},
							},
						},
					},
					&ruleRefExpr{
						pos:  position{line: 386, col: 5, offset: 11891},
						name: "Value",
					},
				},
			},
		},
		{
			name: "Operator",
			pos:  position{line: 388, col: 1, offset: 11898},
			expr: &choiceExpr{
				pos: position{line: 388, col: 12, offset: 11909},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 388, col: 12, offset: 11909},
						name: "OperatorBoolean",
					},
					&ruleRefExpr{
						pos:  position{line: 388, col: 30, offset: 11927},
						name: "OperatorEquality",
					},
					&ruleRefExpr{
						pos:  position{line: 388, col: 49, offset: 11946},
						name: "OperatorHigh",
					},
					&ruleRefExpr{
						pos:  position{line: 388, col: 64, offset: 11961},
						name: "OperatorLow",
					},
				},
			},
		},
		{
			name: "OperatorBoolean",
			pos:  position{line: 390, col: 1, offset: 11974},
			expr: &actionExpr{
				pos: position{line: 390, col: 19, offset: 11992},
				run: (*parser).callonOperatorBoolean1,
				expr: &choiceExpr{
					pos: position{line: 390, col: 21, offset: 11994},
					alternatives: []interface{}{
						&litMatcher{
							pos:        position{line: 390, col: 21, offset: 11994},
							val:        "not",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 390, col: 29, offset: 12002},
							val:        "or",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 390, col: 36, offset: 12009},
							val:        "and",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "OperatorEquality",
			pos:  position{line: 394, col: 1, offset: 12108},
			expr: &actionExpr{
				pos: position{line: 394, col: 20, offset: 12127},
				run: (*parser).callonOperatorEquality1,
				expr: &choiceExpr{
					pos: position{line: 394, col: 22, offset: 12129},
					alternatives: []interface{}{
						&litMatcher{
							pos:        position{line: 394, col: 22, offset: 12129},
							val:        ">=",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 394, col: 29, offset: 12136},
							val:        "<=",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 394, col: 36, offset: 12143},
							val:        "<",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 394, col: 42, offset: 12149},
							val:        ">",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 394, col: 48, offset: 12155},
							val:        "===",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 394, col: 56, offset: 12163},
							val:        "==",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "OperatorHigh",
			pos:  position{line: 398, col: 1, offset: 12269},
			expr: &choiceExpr{
				pos: position{line: 398, col: 16, offset: 12284},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 398, col: 16, offset: 12284},
						run: (*parser).callonOperatorHigh2,
						expr: &choiceExpr{
							pos: position{line: 398, col: 18, offset: 12286},
							alternatives: []interface{}{
								&litMatcher{
									pos:        position{line: 398, col: 18, offset: 12286},
									val:        "/.",
									ignoreCase: false,
								},
								&litMatcher{
									pos:        position{line: 398, col: 25, offset: 12293},
									val:        "*.",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 401, col: 3, offset: 12399},
						run: (*parser).callonOperatorHigh6,
						expr: &choiceExpr{
							pos: position{line: 401, col: 5, offset: 12401},
							alternatives: []interface{}{
								&litMatcher{
									pos:        position{line: 401, col: 5, offset: 12401},
									val:        "*",
									ignoreCase: false,
								},
								&litMatcher{
									pos:        position{line: 401, col: 11, offset: 12407},
									val:        "/",
									ignoreCase: false,
								},
								&litMatcher{
									pos:        position{line: 401, col: 17, offset: 12413},
									val:        "^",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 404, col: 3, offset: 12516},
						run: (*parser).callonOperatorHigh11,
						expr: &litMatcher{
							pos:        position{line: 404, col: 3, offset: 12516},
							val:        "++",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "OperatorLow",
			pos:  position{line: 408, col: 1, offset: 12620},
			expr: &choiceExpr{
				pos: position{line: 408, col: 15, offset: 12634},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 408, col: 15, offset: 12634},
						run: (*parser).callonOperatorLow2,
						expr: &choiceExpr{
							pos: position{line: 408, col: 17, offset: 12636},
							alternatives: []interface{}{
								&litMatcher{
									pos:        position{line: 408, col: 17, offset: 12636},
									val:        "+.",
									ignoreCase: false,
								},
								&litMatcher{
									pos:        position{line: 408, col: 24, offset: 12643},
									val:        "-.",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 411, col: 3, offset: 12749},
						run: (*parser).callonOperatorLow6,
						expr: &choiceExpr{
							pos: position{line: 411, col: 5, offset: 12751},
							alternatives: []interface{}{
								&litMatcher{
									pos:        position{line: 411, col: 5, offset: 12751},
									val:        "+",
									ignoreCase: false,
								},
								&litMatcher{
									pos:        position{line: 411, col: 11, offset: 12757},
									val:        "-",
									ignoreCase: false,
								},
							},
						},
					},
				},
			},
		},
		{
			name: "Value",
			pos:  position{line: 415, col: 1, offset: 12859},
			expr: &choiceExpr{
				pos: position{line: 415, col: 9, offset: 12867},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 415, col: 9, offset: 12867},
						name: "VariableName",
					},
					&ruleRefExpr{
						pos:  position{line: 415, col: 24, offset: 12882},
						name: "Const",
					},
				},
			},
		},
		{
			name: "Assignable",
			pos:  position{line: 417, col: 1, offset: 12889},
			expr: &choiceExpr{
				pos: position{line: 417, col: 14, offset: 12902},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 417, col: 14, offset: 12902},
						name: "VariableName",
					},
					&ruleRefExpr{
						pos:  position{line: 417, col: 29, offset: 12917},
						name: "Unused",
					},
				},
			},
		},
		{
			name: "Identifier",
			pos:  position{line: 419, col: 1, offset: 12925},
			expr: &choiceExpr{
				pos: position{line: 419, col: 14, offset: 12938},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 419, col: 14, offset: 12938},
						name: "VariableName",
					},
					&ruleRefExpr{
						pos:  position{line: 419, col: 29, offset: 12953},
						name: "ModuleName",
					},
				},
			},
		},
		{
			name: "ArrayLiteral",
			pos:  position{line: 421, col: 1, offset: 12965},
			expr: &actionExpr{
				pos: position{line: 421, col: 16, offset: 12980},
				run: (*parser).callonArrayLiteral1,
				expr: &seqExpr{
					pos: position{line: 421, col: 16, offset: 12980},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 421, col: 16, offset: 12980},
							val:        "[",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 421, col: 20, offset: 12984},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 421, col: 22, offset: 12986},
							label: "first",
							expr: &ruleRefExpr{
								pos:  position{line: 421, col: 28, offset: 12992},
								name: "Expr",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 421, col: 33, offset: 12997},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 421, col: 35, offset: 12999},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 421, col: 40, offset: 13004},
								expr: &seqExpr{
									pos: position{line: 421, col: 41, offset: 13005},
									exprs: []interface{}{
										&litMatcher{
											pos:        position{line: 421, col: 41, offset: 13005},
											val:        ",",
											ignoreCase: false,
										},
										&ruleRefExpr{
											pos:  position{line: 421, col: 45, offset: 13009},
											name: "_",
										},
										&ruleRefExpr{
											pos:  position{line: 421, col: 47, offset: 13011},
											name: "Expr",
										},
										&ruleRefExpr{
											pos:  position{line: 421, col: 52, offset: 13016},
											name: "_",
										},
									},
								},
							},
						},
						&zeroOrOneExpr{
							pos: position{line: 421, col: 56, offset: 13020},
							expr: &litMatcher{
								pos:        position{line: 421, col: 56, offset: 13020},
								val:        ",",
								ignoreCase: false,
							},
						},
						&ruleRefExpr{
							pos:  position{line: 421, col: 61, offset: 13025},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 421, col: 63, offset: 13027},
							val:        "]",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "VariantInstance",
			pos:  position{line: 437, col: 1, offset: 13472},
			expr: &actionExpr{
				pos: position{line: 437, col: 19, offset: 13490},
				run: (*parser).callonVariantInstance1,
				expr: &seqExpr{
					pos: position{line: 437, col: 19, offset: 13490},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 437, col: 19, offset: 13490},
							label: "name",
							expr: &ruleRefExpr{
								pos:  position{line: 437, col: 24, offset: 13495},
								name: "ModuleName",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 437, col: 35, offset: 13506},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 437, col: 37, offset: 13508},
							label: "args",
							expr: &ruleRefExpr{
								pos:  position{line: 437, col: 42, offset: 13513},
								name: "Arguments",
							},
						},
					},
				},
			},
		},
		{
			name: "RecordInstance",
			pos:  position{line: 450, col: 1, offset: 13783},
			expr: &actionExpr{
				pos: position{line: 450, col: 18, offset: 13800},
				run: (*parser).callonRecordInstance1,
				expr: &seqExpr{
					pos: position{line: 450, col: 18, offset: 13800},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 450, col: 18, offset: 13800},
							label: "name",
							expr: &ruleRefExpr{
								pos:  position{line: 450, col: 23, offset: 13805},
								name: "ModuleName",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 450, col: 34, offset: 13816},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 450, col: 36, offset: 13818},
							val:        "{",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 450, col: 40, offset: 13822},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 450, col: 42, offset: 13824},
							label: "firstName",
							expr: &ruleRefExpr{
								pos:  position{line: 450, col: 52, offset: 13834},
								name: "VariableName",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 450, col: 65, offset: 13847},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 450, col: 67, offset: 13849},
							val:        ":",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 450, col: 71, offset: 13853},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 450, col: 73, offset: 13855},
							label: "firstValue",
							expr: &ruleRefExpr{
								pos:  position{line: 450, col: 84, offset: 13866},
								name: "Expr",
							},
						},
						&labeledExpr{
							pos:   position{line: 450, col: 89, offset: 13871},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 450, col: 94, offset: 13876},
								expr: &seqExpr{
									pos: position{line: 450, col: 95, offset: 13877},
									exprs: []interface{}{
										&litMatcher{
											pos:        position{line: 450, col: 95, offset: 13877},
											val:        ",",
											ignoreCase: false,
										},
										&ruleRefExpr{
											pos:  position{line: 450, col: 99, offset: 13881},
											name: "_",
										},
										&ruleRefExpr{
											pos:  position{line: 450, col: 101, offset: 13883},
											name: "VariableName",
										},
										&ruleRefExpr{
											pos:  position{line: 450, col: 114, offset: 13896},
											name: "_",
										},
										&litMatcher{
											pos:        position{line: 450, col: 116, offset: 13898},
											val:        ":",
											ignoreCase: false,
										},
										&ruleRefExpr{
											pos:  position{line: 450, col: 120, offset: 13902},
											name: "_",
										},
										&ruleRefExpr{
											pos:  position{line: 450, col: 122, offset: 13904},
											name: "Expr",
										},
									},
								},
							},
						},
						&litMatcher{
							pos:        position{line: 450, col: 130, offset: 13912},
							val:        "}",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "TypeParameter",
			pos:  position{line: 470, col: 1, offset: 14496},
			expr: &actionExpr{
				pos: position{line: 470, col: 17, offset: 14512},
				run: (*parser).callonTypeParameter1,
				expr: &seqExpr{
					pos: position{line: 470, col: 17, offset: 14512},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 470, col: 17, offset: 14512},
							val:        "'",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 470, col: 22, offset: 14517},
							name: "VariableName",
						},
					},
				},
			},
		},
		{
			name: "VariableName",
			pos:  position{line: 474, col: 1, offset: 14625},
			expr: &actionExpr{
				pos: position{line: 474, col: 16, offset: 14640},
				run: (*parser).callonVariableName1,
				expr: &seqExpr{
					pos: position{line: 474, col: 16, offset: 14640},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 474, col: 16, offset: 14640},
							expr: &ruleRefExpr{
								pos:  position{line: 474, col: 17, offset: 14641},
								name: "Reserved",
							},
						},
						&seqExpr{
							pos: position{line: 474, col: 27, offset: 14651},
							exprs: []interface{}{
								&oneOrMoreExpr{
									pos: position{line: 474, col: 27, offset: 14651},
									expr: &charClassMatcher{
										pos:        position{line: 474, col: 27, offset: 14651},
										val:        "[a-z]",
										ranges:     []rune{'a', 'z'},
										ignoreCase: false,
										inverted:   false,
									},
								},
								&zeroOrMoreExpr{
									pos: position{line: 474, col: 34, offset: 14658},
									expr: &charClassMatcher{
										pos:        position{line: 474, col: 34, offset: 14658},
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
			pos:  position{line: 478, col: 1, offset: 14769},
			expr: &actionExpr{
				pos: position{line: 478, col: 14, offset: 14782},
				run: (*parser).callonModuleName1,
				expr: &seqExpr{
					pos: position{line: 478, col: 15, offset: 14783},
					exprs: []interface{}{
						&oneOrMoreExpr{
							pos: position{line: 478, col: 15, offset: 14783},
							expr: &charClassMatcher{
								pos:        position{line: 478, col: 15, offset: 14783},
								val:        "[A-Z]",
								ranges:     []rune{'A', 'Z'},
								ignoreCase: false,
								inverted:   false,
							},
						},
						&zeroOrMoreExpr{
							pos: position{line: 478, col: 22, offset: 14790},
							expr: &charClassMatcher{
								pos:        position{line: 478, col: 22, offset: 14790},
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
			pos:  position{line: 482, col: 1, offset: 14901},
			expr: &choiceExpr{
				pos: position{line: 482, col: 9, offset: 14909},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 482, col: 9, offset: 14909},
						run: (*parser).callonConst2,
						expr: &seqExpr{
							pos: position{line: 482, col: 9, offset: 14909},
							exprs: []interface{}{
								&zeroOrOneExpr{
									pos: position{line: 482, col: 9, offset: 14909},
									expr: &litMatcher{
										pos:        position{line: 482, col: 9, offset: 14909},
										val:        "-",
										ignoreCase: false,
									},
								},
								&oneOrMoreExpr{
									pos: position{line: 482, col: 14, offset: 14914},
									expr: &charClassMatcher{
										pos:        position{line: 482, col: 14, offset: 14914},
										val:        "[0-9]",
										ranges:     []rune{'0', '9'},
										ignoreCase: false,
										inverted:   false,
									},
								},
								&notExpr{
									pos: position{line: 482, col: 21, offset: 14921},
									expr: &litMatcher{
										pos:        position{line: 482, col: 22, offset: 14922},
										val:        ".",
										ignoreCase: false,
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 489, col: 3, offset: 15098},
						run: (*parser).callonConst10,
						expr: &seqExpr{
							pos: position{line: 489, col: 3, offset: 15098},
							exprs: []interface{}{
								&zeroOrOneExpr{
									pos: position{line: 489, col: 3, offset: 15098},
									expr: &litMatcher{
										pos:        position{line: 489, col: 3, offset: 15098},
										val:        "-",
										ignoreCase: false,
									},
								},
								&oneOrMoreExpr{
									pos: position{line: 489, col: 8, offset: 15103},
									expr: &charClassMatcher{
										pos:        position{line: 489, col: 8, offset: 15103},
										val:        "[0-9]",
										ranges:     []rune{'0', '9'},
										ignoreCase: false,
										inverted:   false,
									},
								},
								&litMatcher{
									pos:        position{line: 489, col: 15, offset: 15110},
									val:        ".",
									ignoreCase: false,
								},
								&oneOrMoreExpr{
									pos: position{line: 489, col: 19, offset: 15114},
									expr: &charClassMatcher{
										pos:        position{line: 489, col: 19, offset: 15114},
										val:        "[0-9]",
										ranges:     []rune{'0', '9'},
										ignoreCase: false,
										inverted:   false,
									},
								},
							},
						},
					},
					&litMatcher{
						pos:        position{line: 496, col: 3, offset: 15304},
						val:        "true",
						ignoreCase: false,
					},
					&actionExpr{
						pos: position{line: 496, col: 12, offset: 15313},
						run: (*parser).callonConst20,
						expr: &litMatcher{
							pos:        position{line: 496, col: 12, offset: 15313},
							val:        "false",
							ignoreCase: false,
						},
					},
					&ruleRefExpr{
						pos:  position{line: 502, col: 3, offset: 15514},
						name: "Unit",
					},
					&actionExpr{
						pos: position{line: 503, col: 3, offset: 15521},
						run: (*parser).callonConst23,
						expr: &seqExpr{
							pos: position{line: 503, col: 3, offset: 15521},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 503, col: 3, offset: 15521},
									val:        "\"",
									ignoreCase: false,
								},
								&zeroOrMoreExpr{
									pos: position{line: 503, col: 7, offset: 15525},
									expr: &seqExpr{
										pos: position{line: 503, col: 8, offset: 15526},
										exprs: []interface{}{
											&notExpr{
												pos: position{line: 503, col: 8, offset: 15526},
												expr: &ruleRefExpr{
													pos:  position{line: 503, col: 9, offset: 15527},
													name: "EscapedChar",
												},
											},
											&anyMatcher{
												line: 503, col: 21, offset: 15539,
											},
										},
									},
								},
								&litMatcher{
									pos:        position{line: 503, col: 25, offset: 15543},
									val:        "\"",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 510, col: 3, offset: 15727},
						run: (*parser).callonConst32,
						expr: &seqExpr{
							pos: position{line: 510, col: 3, offset: 15727},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 510, col: 3, offset: 15727},
									val:        "'",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 510, col: 7, offset: 15731},
									label: "val",
									expr: &seqExpr{
										pos: position{line: 510, col: 12, offset: 15736},
										exprs: []interface{}{
											&notExpr{
												pos: position{line: 510, col: 12, offset: 15736},
												expr: &ruleRefExpr{
													pos:  position{line: 510, col: 13, offset: 15737},
													name: "EscapedChar",
												},
											},
											&anyMatcher{
												line: 510, col: 25, offset: 15749,
											},
										},
									},
								},
								&litMatcher{
									pos:        position{line: 510, col: 28, offset: 15752},
									val:        "'",
									ignoreCase: false,
								},
							},
						},
					},
					&ruleRefExpr{
						pos:  position{line: 512, col: 5, offset: 15844},
						name: "ArrayLiteral",
					},
					&ruleRefExpr{
						pos:  position{line: 512, col: 20, offset: 15859},
						name: "RecordInstance",
					},
					&ruleRefExpr{
						pos:  position{line: 512, col: 37, offset: 15876},
						name: "VariantInstance",
					},
				},
			},
		},
		{
			name: "Unit",
			pos:  position{line: 514, col: 1, offset: 15893},
			expr: &actionExpr{
				pos: position{line: 514, col: 8, offset: 15900},
				run: (*parser).callonUnit1,
				expr: &litMatcher{
					pos:        position{line: 514, col: 8, offset: 15900},
					val:        "()",
					ignoreCase: false,
				},
			},
		},
		{
			name: "Unused",
			pos:  position{line: 518, col: 1, offset: 15962},
			expr: &actionExpr{
				pos: position{line: 518, col: 10, offset: 15971},
				run: (*parser).callonUnused1,
				expr: &litMatcher{
					pos:        position{line: 518, col: 11, offset: 15972},
					val:        "_",
					ignoreCase: false,
				},
			},
		},
		{
			name: "Reserved",
			pos:  position{line: 522, col: 1, offset: 16073},
			expr: &seqExpr{
				pos: position{line: 522, col: 12, offset: 16084},
				exprs: []interface{}{
					&choiceExpr{
						pos: position{line: 522, col: 13, offset: 16085},
						alternatives: []interface{}{
							&litMatcher{
								pos:        position{line: 522, col: 13, offset: 16085},
								val:        "let",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 522, col: 21, offset: 16093},
								val:        "if",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 522, col: 28, offset: 16100},
								val:        "else",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 522, col: 37, offset: 16109},
								val:        "func",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 522, col: 46, offset: 16118},
								val:        "type",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 522, col: 55, offset: 16127},
								val:        "true",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 522, col: 64, offset: 16136},
								val:        "false",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 522, col: 74, offset: 16146},
								val:        "mutable",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 522, col: 86, offset: 16158},
								val:        "type",
								ignoreCase: false,
							},
							&ruleRefExpr{
								pos:  position{line: 522, col: 95, offset: 16167},
								name: "BaseType",
							},
						},
					},
					&notExpr{
						pos: position{line: 522, col: 105, offset: 16177},
						expr: &oneOrMoreExpr{
							pos: position{line: 522, col: 106, offset: 16178},
							expr: &charClassMatcher{
								pos:        position{line: 522, col: 106, offset: 16178},
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
			pos:  position{line: 524, col: 1, offset: 16186},
			expr: &actionExpr{
				pos: position{line: 524, col: 12, offset: 16197},
				run: (*parser).callonBaseType1,
				expr: &choiceExpr{
					pos: position{line: 524, col: 14, offset: 16199},
					alternatives: []interface{}{
						&litMatcher{
							pos:        position{line: 524, col: 14, offset: 16199},
							val:        "int",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 524, col: 22, offset: 16207},
							val:        "bool",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 524, col: 31, offset: 16216},
							val:        "string",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 524, col: 42, offset: 16227},
							val:        "byte",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 524, col: 51, offset: 16236},
							val:        "rune",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 524, col: 60, offset: 16245},
							val:        "float",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 524, col: 70, offset: 16255},
							val:        "list",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "EscapedChar",
			pos:  position{line: 528, col: 1, offset: 16354},
			expr: &charClassMatcher{
				pos:        position{line: 528, col: 15, offset: 16368},
				val:        "[\\x00-\\x1f\"\\\\]",
				chars:      []rune{'"', '\\'},
				ranges:     []rune{'\x00', '\x1f'},
				ignoreCase: false,
				inverted:   false,
			},
		},
		{
			name: "EscapeSequence",
			pos:  position{line: 530, col: 1, offset: 16384},
			expr: &choiceExpr{
				pos: position{line: 530, col: 18, offset: 16401},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 530, col: 18, offset: 16401},
						name: "SingleCharEscape",
					},
					&ruleRefExpr{
						pos:  position{line: 530, col: 37, offset: 16420},
						name: "UnicodeEscape",
					},
				},
			},
		},
		{
			name: "SingleCharEscape",
			pos:  position{line: 532, col: 1, offset: 16435},
			expr: &charClassMatcher{
				pos:        position{line: 532, col: 20, offset: 16454},
				val:        "[\"\\\\/bfnrt]",
				chars:      []rune{'"', '\\', '/', 'b', 'f', 'n', 'r', 't'},
				ignoreCase: false,
				inverted:   false,
			},
		},
		{
			name: "DecimalDigit",
			pos:  position{line: 534, col: 1, offset: 16467},
			expr: &charClassMatcher{
				pos:        position{line: 534, col: 16, offset: 16482},
				val:        "[0-9]",
				ranges:     []rune{'0', '9'},
				ignoreCase: false,
				inverted:   false,
			},
		},
		{
			name: "NonZeroDecimalDigit",
			pos:  position{line: 536, col: 1, offset: 16489},
			expr: &charClassMatcher{
				pos:        position{line: 536, col: 23, offset: 16511},
				val:        "[1-9]",
				ranges:     []rune{'1', '9'},
				ignoreCase: false,
				inverted:   false,
			},
		},
		{
			name: "HexDigit",
			pos:  position{line: 538, col: 1, offset: 16518},
			expr: &charClassMatcher{
				pos:        position{line: 538, col: 12, offset: 16529},
				val:        "[0-9a-f]i",
				ranges:     []rune{'0', '9', 'a', 'f'},
				ignoreCase: true,
				inverted:   false,
			},
		},
		{
			name:        "__",
			displayName: "\"sigwhitespace\"",
			pos:         position{line: 540, col: 1, offset: 16540},
			expr: &choiceExpr{
				pos: position{line: 540, col: 22, offset: 16561},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 540, col: 22, offset: 16561},
						name: "Comment",
					},
					&oneOrMoreExpr{
						pos: position{line: 540, col: 32, offset: 16571},
						expr: &charClassMatcher{
							pos:        position{line: 540, col: 32, offset: 16571},
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
			pos:         position{line: 542, col: 1, offset: 16583},
			expr: &zeroOrMoreExpr{
				pos: position{line: 542, col: 18, offset: 16600},
				expr: &charClassMatcher{
					pos:        position{line: 542, col: 18, offset: 16600},
					val:        "[ \\r\\n\\t]",
					chars:      []rune{' ', '\r', '\n', '\t'},
					ignoreCase: false,
					inverted:   false,
				},
			},
		},
		{
			name: "EOF",
			pos:  position{line: 544, col: 1, offset: 16612},
			expr: &notExpr{
				pos: position{line: 544, col: 7, offset: 16618},
				expr: &anyMatcher{
					line: 544, col: 8, offset: 16619,
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
	return BasicAst{Type: "Comment", StringValue: string(c.text[1:]), ValueType: STRING}, nil
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

	return AliasType{Name: name.(BasicAst).StringValue, Params: parameters, Types: fields}, nil
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

	return RecordType{Name: name.(BasicAst).StringValue, Fields: fields}, nil
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

	return VariantType{Name: name.(BasicAst).StringValue, Params: parameters, Constructors: constructors}, nil
}

func (p *parser) callonTypeDefn54() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onTypeDefn54(stack["name"], stack["params"], stack["rest"])
}

func (c *current) onRecordFieldDefn1(name, t interface{}) (interface{}, error) {
	return RecordField{Name: name.(BasicAst).StringValue, Type: t.(Ast)}, nil
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

	record := RecordType{Name: name.(BasicAst).StringValue, Fields: fields}
	return VariantConstructor{Name: name.(BasicAst).StringValue, Fields: []Ast{record}}, nil
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

	return VariantConstructor{Name: name.(BasicAst).StringValue, Fields: params}, nil
}

func (p *parser) callonVariantConstructor26() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onVariantConstructor26(stack["name"], stack["rest"])
}

func (c *current) onAssignment2(i, rest, expr interface{}) (interface{}, error) {
	//fmt.Println("assignment:", string(c.text))
	vals := []Ast{i.(Ast)}
	if len(rest.([]interface{})) > 0 {
		restSl := toIfaceSlice(rest)
		for _, v := range restSl {
			// we can get each item in the grammar by index
			restExpr := toIfaceSlice(v)
			v := restExpr[2].(Ast)
			vals = append(vals, v)
		}
	}
	return Assignment{Left: vals, Right: expr.(Ast)}, nil
}

func (p *parser) callonAssignment2() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onAssignment2(stack["i"], stack["rest"], stack["expr"])
}

func (c *current) onAssignment21() (interface{}, error) {
	return nil, errors.New("Variable name or '_' (unused result character) required here")
}

func (p *parser) callonAssignment21() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onAssignment21()
}

func (c *current) onAssignment28(i, rest interface{}) (interface{}, error) {
	return nil, errors.New("When assigning a value to a variable, you must use '='")
}

func (p *parser) callonAssignment28() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onAssignment28(stack["i"], stack["rest"])
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
	return Func{Name: i.(BasicAst).StringValue, Arguments: args, Subvalues: subvalues}, nil
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

func (c *current) onIfExpr21(expr, thens, elses interface{}) (interface{}, error) {
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

func (p *parser) callonIfExpr21() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onIfExpr21(stack["expr"], stack["thens"], stack["elses"])
}

func (c *current) onIfExpr45(expr, thens interface{}) (interface{}, error) {
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

func (p *parser) callonIfExpr45() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onIfExpr45(stack["expr"], stack["thens"])
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

func (c *current) onCall13(fn, args interface{}) (interface{}, error) {
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

func (p *parser) callonCall13() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onCall13(stack["fn"], stack["args"])
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

		return Container{Type: "CompoundExpr", Subvalues: subvalues}, nil
	} else {
		return Container{Type: "CompoundExpr", Subvalues: []Ast{op.(Ast)}}, nil
	}
}

func (p *parser) callonCompoundExpr1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onCompoundExpr1(stack["op"], stack["rest"])
}

func (c *current) onBinOpBool1(first, rest interface{}) (interface{}, error) {
	//fmt.Println("binopbool", first, rest)
	vals := rest.([]interface{})
	if len(vals) > 0 {
		subvalues := []Ast{first.(Ast)}
		restSl := toIfaceSlice(rest)
		for _, v := range restSl {
			// we can get each item in the grammar by index
			restExpr := toIfaceSlice(v)
			v := restExpr[3].(Ast)
			op := restExpr[1].(Ast)
			subvalues = append(subvalues, op, v)
		}
		return Container{Type: "BinOpBool", Subvalues: subvalues}, nil
	} else {
		return first.(Ast), nil
	}
}

func (p *parser) callonBinOpBool1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onBinOpBool1(stack["first"], stack["rest"])
}

func (c *current) onBinOpEquality1(first, rest interface{}) (interface{}, error) {
	//fmt.Println("binopeq", first, rest)
	vals := rest.([]interface{})
	if len(vals) > 0 {
		subvalues := []Ast{first.(Ast)}
		restSl := toIfaceSlice(rest)
		for _, v := range restSl {
			// we can get each item in the grammar by index
			restExpr := toIfaceSlice(v)
			v := restExpr[3].(Ast)
			op := restExpr[1].(Ast)
			subvalues = append(subvalues, op, v)
		}
		return Container{Type: "BinOpEquality", Subvalues: subvalues}, nil
	} else {
		return first.(Ast), nil
	}

}

func (p *parser) callonBinOpEquality1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onBinOpEquality1(stack["first"], stack["rest"])
}

func (c *current) onBinOpLow1(first, rest interface{}) (interface{}, error) {
	//fmt.Println("binoplow", first, rest)
	vals := rest.([]interface{})
	if len(vals) > 0 {
		subvalues := []Ast{first.(Ast)}
		restSl := toIfaceSlice(rest)
		for _, v := range restSl {
			// we can get each item in the grammar by index
			restExpr := toIfaceSlice(v)
			v := restExpr[3].(Ast)
			op := restExpr[1].(Ast)
			subvalues = append(subvalues, op, v)
		}
		return Container{Type: "BinOpLow", Subvalues: subvalues}, nil
	} else {
		return first.(Ast), nil
	}

}

func (p *parser) callonBinOpLow1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onBinOpLow1(stack["first"], stack["rest"])
}

func (c *current) onBinOpHigh1(first, rest interface{}) (interface{}, error) {
	//fmt.Println("binophigh", first, rest)
	vals := rest.([]interface{})
	if len(vals) > 0 {
		subvalues := []Ast{first.(Ast)}
		restSl := toIfaceSlice(rest)
		for _, v := range restSl {
			// we can get each item in the grammar by index
			restExpr := toIfaceSlice(v)
			v := restExpr[3].(Ast)
			op := restExpr[1].(Ast)
			subvalues = append(subvalues, op, v)
		}
		return Container{Type: "BinOpHigh", Subvalues: subvalues}, nil
	} else {
		return first.(Ast), nil
	}
}

func (p *parser) callonBinOpHigh1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onBinOpHigh1(stack["first"], stack["rest"])
}

func (c *current) onBinOpParens2(first interface{}) (interface{}, error) {
	//fmt.Println("binopparens", first)
	return Container{Type: "BinOpParens", Subvalues: []Ast{first.(Ast)}}, nil
}

func (p *parser) callonBinOpParens2() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onBinOpParens2(stack["first"])
}

func (c *current) onOperatorBoolean1() (interface{}, error) {
	return BasicAst{Type: "BoolOp", StringValue: string(c.text), ValueType: STRING}, nil
}

func (p *parser) callonOperatorBoolean1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onOperatorBoolean1()
}

func (c *current) onOperatorEquality1() (interface{}, error) {
	return BasicAst{Type: "EqualityOperator", StringValue: string(c.text), ValueType: STRING}, nil
}

func (p *parser) callonOperatorEquality1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onOperatorEquality1()
}

func (c *current) onOperatorHigh2() (interface{}, error) {
	return BasicAst{Type: "FloatOperator", StringValue: string(c.text), ValueType: STRING}, nil
}

func (p *parser) callonOperatorHigh2() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onOperatorHigh2()
}

func (c *current) onOperatorHigh6() (interface{}, error) {
	return BasicAst{Type: "IntOperator", StringValue: string(c.text), ValueType: STRING}, nil
}

func (p *parser) callonOperatorHigh6() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onOperatorHigh6()
}

func (c *current) onOperatorHigh11() (interface{}, error) {
	return BasicAst{Type: "StringOperator", StringValue: string(c.text), ValueType: STRING}, nil
}

func (p *parser) callonOperatorHigh11() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onOperatorHigh11()
}

func (c *current) onOperatorLow2() (interface{}, error) {
	return BasicAst{Type: "FloatOperator", StringValue: string(c.text), ValueType: STRING}, nil
}

func (p *parser) callonOperatorLow2() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onOperatorLow2()
}

func (c *current) onOperatorLow6() (interface{}, error) {
	return BasicAst{Type: "IntOperator", StringValue: string(c.text), ValueType: STRING}, nil
}

func (p *parser) callonOperatorLow6() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onOperatorLow6()
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
	return BasicAst{Type: "TypeParam", StringValue: string(c.text), ValueType: STRING}, nil
}

func (p *parser) callonTypeParameter1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onTypeParameter1()
}

func (c *current) onVariableName1() (interface{}, error) {
	return BasicAst{Type: "Identifier", StringValue: string(c.text), ValueType: STRING}, nil
}

func (p *parser) callonVariableName1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onVariableName1()
}

func (c *current) onModuleName1() (interface{}, error) {
	return BasicAst{Type: "Identifier", StringValue: string(c.text), ValueType: STRING}, nil
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

func (c *current) onConst20() (interface{}, error) {
	if string(c.text) == "true" {
		return BasicAst{Type: "Bool", BoolValue: true, ValueType: BOOL}, nil
	}
	return BasicAst{Type: "Bool", BoolValue: false, ValueType: BOOL}, nil
}

func (p *parser) callonConst20() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onConst20()
}

func (c *current) onConst23() (interface{}, error) {
	val, err := strconv.Unquote(string(c.text))
	if err == nil {
		return BasicAst{Type: "String", StringValue: val, ValueType: STRING}, nil
	}
	return nil, err
}

func (p *parser) callonConst23() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onConst23()
}

func (c *current) onConst32(val interface{}) (interface{}, error) {
	return BasicAst{Type: "Char", CharValue: rune(c.text[1]), ValueType: CHAR}, nil
}

func (p *parser) callonConst32() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onConst32(stack["val"])
}

func (c *current) onUnit1() (interface{}, error) {
	return BasicAst{Type: "Nil", ValueType: NIL}, nil
}

func (p *parser) callonUnit1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onUnit1()
}

func (c *current) onUnused1() (interface{}, error) {
	return BasicAst{Type: "Identifier", StringValue: string(c.text), ValueType: STRING}, nil
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
