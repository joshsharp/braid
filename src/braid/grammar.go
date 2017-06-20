package main

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
			pos:  position{line: 13, col: 1, offset: 143},
			expr: &actionExpr{
				pos: position{line: 13, col: 10, offset: 152},
				run: (*parser).callonModule1,
				expr: &seqExpr{
					pos: position{line: 13, col: 10, offset: 152},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 13, col: 10, offset: 152},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 13, col: 12, offset: 154},
							label: "stat",
							expr: &ruleRefExpr{
								pos:  position{line: 13, col: 17, offset: 159},
								name: "TopLevelStatement",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 13, col: 35, offset: 177},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 13, col: 37, offset: 179},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 13, col: 42, offset: 184},
								expr: &ruleRefExpr{
									pos:  position{line: 13, col: 43, offset: 185},
									name: "TopLevelStatement",
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 13, col: 63, offset: 205},
							name: "_",
						},
						&ruleRefExpr{
							pos:  position{line: 13, col: 65, offset: 207},
							name: "EOF",
						},
					},
				},
			},
		},
		{
			name: "TopLevelStatement",
			pos:  position{line: 28, col: 1, offset: 695},
			expr: &choiceExpr{
				pos: position{line: 28, col: 21, offset: 715},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 28, col: 21, offset: 715},
						name: "Comment",
					},
					&ruleRefExpr{
						pos:  position{line: 28, col: 31, offset: 725},
						name: "FuncDefn",
					},
					&ruleRefExpr{
						pos:  position{line: 28, col: 42, offset: 736},
						name: "TypeDefn",
					},
				},
			},
		},
		{
			name: "Statement",
			pos:  position{line: 30, col: 1, offset: 746},
			expr: &choiceExpr{
				pos: position{line: 30, col: 13, offset: 758},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 30, col: 13, offset: 758},
						name: "Comment",
					},
					&ruleRefExpr{
						pos:  position{line: 30, col: 23, offset: 768},
						name: "FuncDefn",
					},
					&ruleRefExpr{
						pos:  position{line: 30, col: 34, offset: 779},
						name: "Assignment",
					},
					&ruleRefExpr{
						pos:  position{line: 30, col: 47, offset: 792},
						name: "ExprLine",
					},
				},
			},
		},
		{
			name: "ExprLine",
			pos:  position{line: 32, col: 1, offset: 802},
			expr: &actionExpr{
				pos: position{line: 32, col: 12, offset: 813},
				run: (*parser).callonExprLine1,
				expr: &seqExpr{
					pos: position{line: 32, col: 12, offset: 813},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 32, col: 12, offset: 813},
							label: "e",
							expr: &ruleRefExpr{
								pos:  position{line: 32, col: 14, offset: 815},
								name: "Expr",
							},
						},
						&andExpr{
							pos: position{line: 32, col: 19, offset: 820},
							expr: &litMatcher{
								pos:        position{line: 32, col: 20, offset: 821},
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
			pos:  position{line: 36, col: 1, offset: 849},
			expr: &actionExpr{
				pos: position{line: 36, col: 11, offset: 859},
				run: (*parser).callonComment1,
				expr: &seqExpr{
					pos: position{line: 36, col: 11, offset: 859},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 36, col: 11, offset: 859},
							val:        "#",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 36, col: 15, offset: 863},
							label: "comment",
							expr: &zeroOrMoreExpr{
								pos: position{line: 36, col: 23, offset: 871},
								expr: &seqExpr{
									pos: position{line: 36, col: 24, offset: 872},
									exprs: []interface{}{
										&notExpr{
											pos: position{line: 36, col: 24, offset: 872},
											expr: &ruleRefExpr{
												pos:  position{line: 36, col: 25, offset: 873},
												name: "EscapedChar",
											},
										},
										&anyMatcher{
											line: 36, col: 37, offset: 885,
										},
									},
								},
							},
						},
						&andExpr{
							pos: position{line: 36, col: 41, offset: 889},
							expr: &litMatcher{
								pos:        position{line: 36, col: 42, offset: 890},
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
			pos:  position{line: 41, col: 1, offset: 1040},
			expr: &choiceExpr{
				pos: position{line: 41, col: 12, offset: 1051},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 41, col: 12, offset: 1051},
						run: (*parser).callonTypeDefn2,
						expr: &seqExpr{
							pos: position{line: 41, col: 12, offset: 1051},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 41, col: 12, offset: 1051},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 41, col: 14, offset: 1053},
									val:        "type",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 41, col: 21, offset: 1060},
									name: "__",
								},
								&labeledExpr{
									pos:   position{line: 41, col: 24, offset: 1063},
									label: "name",
									expr: &ruleRefExpr{
										pos:  position{line: 41, col: 29, offset: 1068},
										name: "ModuleName",
									},
								},
								&labeledExpr{
									pos:   position{line: 41, col: 40, offset: 1079},
									label: "params",
									expr: &zeroOrMoreExpr{
										pos: position{line: 41, col: 47, offset: 1086},
										expr: &seqExpr{
											pos: position{line: 41, col: 48, offset: 1087},
											exprs: []interface{}{
												&ruleRefExpr{
													pos:  position{line: 41, col: 48, offset: 1087},
													name: "__",
												},
												&ruleRefExpr{
													pos:  position{line: 41, col: 51, offset: 1090},
													name: "TypeParameter",
												},
											},
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 41, col: 67, offset: 1106},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 41, col: 69, offset: 1108},
									val:        "=",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 41, col: 73, offset: 1112},
									label: "types",
									expr: &oneOrMoreExpr{
										pos: position{line: 41, col: 79, offset: 1118},
										expr: &seqExpr{
											pos: position{line: 41, col: 80, offset: 1119},
											exprs: []interface{}{
												&ruleRefExpr{
													pos:  position{line: 41, col: 80, offset: 1119},
													name: "__",
												},
												&ruleRefExpr{
													pos:  position{line: 41, col: 83, offset: 1122},
													name: "AnyType",
												},
											},
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 41, col: 93, offset: 1132},
									name: "__",
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 46, col: 1, offset: 1222},
						run: (*parser).callonTypeDefn22,
						expr: &seqExpr{
							pos: position{line: 46, col: 1, offset: 1222},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 46, col: 1, offset: 1222},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 46, col: 3, offset: 1224},
									val:        "type",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 46, col: 10, offset: 1231},
									name: "__",
								},
								&labeledExpr{
									pos:   position{line: 46, col: 13, offset: 1234},
									label: "name",
									expr: &ruleRefExpr{
										pos:  position{line: 46, col: 18, offset: 1239},
										name: "ModuleName",
									},
								},
								&labeledExpr{
									pos:   position{line: 46, col: 29, offset: 1250},
									label: "params",
									expr: &zeroOrMoreExpr{
										pos: position{line: 46, col: 36, offset: 1257},
										expr: &seqExpr{
											pos: position{line: 46, col: 37, offset: 1258},
											exprs: []interface{}{
												&ruleRefExpr{
													pos:  position{line: 46, col: 37, offset: 1258},
													name: "__",
												},
												&ruleRefExpr{
													pos:  position{line: 46, col: 40, offset: 1261},
													name: "TypeParameter",
												},
											},
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 46, col: 56, offset: 1277},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 46, col: 58, offset: 1279},
									val:        "=",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 46, col: 62, offset: 1283},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 47, col: 5, offset: 1289},
									val:        "{",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 47, col: 9, offset: 1293},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 47, col: 11, offset: 1295},
									label: "first",
									expr: &ruleRefExpr{
										pos:  position{line: 47, col: 17, offset: 1301},
										name: "RecordFieldDefn",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 47, col: 33, offset: 1317},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 47, col: 35, offset: 1319},
									label: "rest",
									expr: &zeroOrMoreExpr{
										pos: position{line: 47, col: 40, offset: 1324},
										expr: &seqExpr{
											pos: position{line: 47, col: 41, offset: 1325},
											exprs: []interface{}{
												&litMatcher{
													pos:        position{line: 47, col: 41, offset: 1325},
													val:        ",",
													ignoreCase: false,
												},
												&ruleRefExpr{
													pos:  position{line: 47, col: 45, offset: 1329},
													name: "_",
												},
												&ruleRefExpr{
													pos:  position{line: 47, col: 47, offset: 1331},
													name: "RecordFieldDefn",
												},
												&ruleRefExpr{
													pos:  position{line: 47, col: 63, offset: 1347},
													name: "_",
												},
											},
										},
									},
								},
								&zeroOrOneExpr{
									pos: position{line: 47, col: 67, offset: 1351},
									expr: &litMatcher{
										pos:        position{line: 47, col: 67, offset: 1351},
										val:        ",",
										ignoreCase: false,
									},
								},
								&ruleRefExpr{
									pos:  position{line: 47, col: 72, offset: 1356},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 47, col: 74, offset: 1358},
									val:        "}",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 47, col: 78, offset: 1362},
									name: "__",
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 65, col: 1, offset: 1845},
						run: (*parser).callonTypeDefn54,
						expr: &seqExpr{
							pos: position{line: 65, col: 1, offset: 1845},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 65, col: 1, offset: 1845},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 65, col: 3, offset: 1847},
									val:        "type",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 65, col: 10, offset: 1854},
									name: "__",
								},
								&labeledExpr{
									pos:   position{line: 65, col: 13, offset: 1857},
									label: "name",
									expr: &ruleRefExpr{
										pos:  position{line: 65, col: 18, offset: 1862},
										name: "ModuleName",
									},
								},
								&labeledExpr{
									pos:   position{line: 65, col: 29, offset: 1873},
									label: "params",
									expr: &zeroOrMoreExpr{
										pos: position{line: 65, col: 36, offset: 1880},
										expr: &seqExpr{
											pos: position{line: 65, col: 37, offset: 1881},
											exprs: []interface{}{
												&ruleRefExpr{
													pos:  position{line: 65, col: 37, offset: 1881},
													name: "__",
												},
												&ruleRefExpr{
													pos:  position{line: 65, col: 40, offset: 1884},
													name: "TypeParameter",
												},
											},
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 65, col: 56, offset: 1900},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 65, col: 58, offset: 1902},
									val:        "=",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 65, col: 62, offset: 1906},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 65, col: 64, offset: 1908},
									label: "rest",
									expr: &oneOrMoreExpr{
										pos: position{line: 65, col: 69, offset: 1913},
										expr: &ruleRefExpr{
											pos:  position{line: 65, col: 70, offset: 1914},
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
			name: "VariantConstructor",
			pos:  position{line: 80, col: 1, offset: 2319},
			expr: &actionExpr{
				pos: position{line: 80, col: 22, offset: 2340},
				run: (*parser).callonVariantConstructor1,
				expr: &seqExpr{
					pos: position{line: 80, col: 22, offset: 2340},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 80, col: 22, offset: 2340},
							val:        "|",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 80, col: 26, offset: 2344},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 80, col: 28, offset: 2346},
							label: "name",
							expr: &ruleRefExpr{
								pos:  position{line: 80, col: 33, offset: 2351},
								name: "ModuleName",
							},
						},
						&labeledExpr{
							pos:   position{line: 80, col: 44, offset: 2362},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 80, col: 49, offset: 2367},
								expr: &seqExpr{
									pos: position{line: 80, col: 50, offset: 2368},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 80, col: 50, offset: 2368},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 80, col: 53, offset: 2371},
											name: "AnyType",
										},
									},
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 80, col: 63, offset: 2381},
							name: "_",
						},
					},
				},
			},
		},
		{
			name: "RecordFieldDefn",
			pos:  position{line: 97, col: 1, offset: 2816},
			expr: &actionExpr{
				pos: position{line: 97, col: 19, offset: 2834},
				run: (*parser).callonRecordFieldDefn1,
				expr: &seqExpr{
					pos: position{line: 97, col: 19, offset: 2834},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 97, col: 19, offset: 2834},
							label: "name",
							expr: &ruleRefExpr{
								pos:  position{line: 97, col: 24, offset: 2839},
								name: "VariableName",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 97, col: 37, offset: 2852},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 97, col: 39, offset: 2854},
							val:        ":",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 97, col: 43, offset: 2858},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 97, col: 45, offset: 2860},
							label: "t",
							expr: &ruleRefExpr{
								pos:  position{line: 97, col: 48, offset: 2863},
								name: "AnyType",
							},
						},
					},
				},
			},
		},
		{
			name: "AnyType",
			pos:  position{line: 101, col: 1, offset: 2955},
			expr: &choiceExpr{
				pos: position{line: 101, col: 11, offset: 2965},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 101, col: 11, offset: 2965},
						name: "BaseType",
					},
					&ruleRefExpr{
						pos:  position{line: 101, col: 22, offset: 2976},
						name: "TypeParameter",
					},
				},
			},
		},
		{
			name: "Assignment",
			pos:  position{line: 103, col: 1, offset: 2991},
			expr: &choiceExpr{
				pos: position{line: 103, col: 14, offset: 3004},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 103, col: 14, offset: 3004},
						run: (*parser).callonAssignment2,
						expr: &seqExpr{
							pos: position{line: 103, col: 14, offset: 3004},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 103, col: 14, offset: 3004},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 103, col: 16, offset: 3006},
									val:        "let",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 103, col: 22, offset: 3012},
									name: "__",
								},
								&labeledExpr{
									pos:   position{line: 103, col: 25, offset: 3015},
									label: "i",
									expr: &ruleRefExpr{
										pos:  position{line: 103, col: 27, offset: 3017},
										name: "Assignable",
									},
								},
								&labeledExpr{
									pos:   position{line: 103, col: 38, offset: 3028},
									label: "rest",
									expr: &zeroOrMoreExpr{
										pos: position{line: 103, col: 43, offset: 3033},
										expr: &seqExpr{
											pos: position{line: 103, col: 44, offset: 3034},
											exprs: []interface{}{
												&litMatcher{
													pos:        position{line: 103, col: 44, offset: 3034},
													val:        ",",
													ignoreCase: false,
												},
												&ruleRefExpr{
													pos:  position{line: 103, col: 48, offset: 3038},
													name: "_",
												},
												&ruleRefExpr{
													pos:  position{line: 103, col: 50, offset: 3040},
													name: "Assignable",
												},
											},
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 103, col: 63, offset: 3053},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 103, col: 65, offset: 3055},
									val:        "=",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 103, col: 69, offset: 3059},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 103, col: 71, offset: 3061},
									label: "expr",
									expr: &ruleRefExpr{
										pos:  position{line: 103, col: 76, offset: 3066},
										name: "Expr",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 103, col: 81, offset: 3071},
									name: "_",
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 118, col: 1, offset: 3510},
						run: (*parser).callonAssignment21,
						expr: &seqExpr{
							pos: position{line: 118, col: 1, offset: 3510},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 118, col: 1, offset: 3510},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 118, col: 3, offset: 3512},
									val:        "let",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 118, col: 9, offset: 3518},
									name: "__",
								},
								&notExpr{
									pos: position{line: 118, col: 12, offset: 3521},
									expr: &ruleRefExpr{
										pos:  position{line: 118, col: 13, offset: 3522},
										name: "Assignable",
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 122, col: 1, offset: 3630},
						run: (*parser).callonAssignment28,
						expr: &seqExpr{
							pos: position{line: 122, col: 1, offset: 3630},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 122, col: 1, offset: 3630},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 122, col: 3, offset: 3632},
									val:        "let",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 122, col: 9, offset: 3638},
									name: "__",
								},
								&labeledExpr{
									pos:   position{line: 122, col: 12, offset: 3641},
									label: "i",
									expr: &ruleRefExpr{
										pos:  position{line: 122, col: 14, offset: 3643},
										name: "Assignable",
									},
								},
								&labeledExpr{
									pos:   position{line: 122, col: 25, offset: 3654},
									label: "rest",
									expr: &zeroOrMoreExpr{
										pos: position{line: 122, col: 30, offset: 3659},
										expr: &seqExpr{
											pos: position{line: 122, col: 31, offset: 3660},
											exprs: []interface{}{
												&litMatcher{
													pos:        position{line: 122, col: 31, offset: 3660},
													val:        ",",
													ignoreCase: false,
												},
												&ruleRefExpr{
													pos:  position{line: 122, col: 35, offset: 3664},
													name: "_",
												},
												&ruleRefExpr{
													pos:  position{line: 122, col: 37, offset: 3666},
													name: "Assignable",
												},
											},
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 122, col: 50, offset: 3679},
									name: "_",
								},
								&notExpr{
									pos: position{line: 122, col: 52, offset: 3681},
									expr: &litMatcher{
										pos:        position{line: 122, col: 53, offset: 3682},
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
			pos:  position{line: 126, col: 1, offset: 3776},
			expr: &actionExpr{
				pos: position{line: 126, col: 12, offset: 3787},
				run: (*parser).callonFuncDefn1,
				expr: &seqExpr{
					pos: position{line: 126, col: 12, offset: 3787},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 126, col: 12, offset: 3787},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 126, col: 14, offset: 3789},
							val:        "let",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 126, col: 20, offset: 3795},
							name: "__",
						},
						&labeledExpr{
							pos:   position{line: 126, col: 23, offset: 3798},
							label: "i",
							expr: &ruleRefExpr{
								pos:  position{line: 126, col: 25, offset: 3800},
								name: "VariableName",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 126, col: 38, offset: 3813},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 126, col: 40, offset: 3815},
							val:        "=",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 126, col: 44, offset: 3819},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 126, col: 46, offset: 3821},
							val:        "func",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 126, col: 53, offset: 3828},
							name: "__",
						},
						&labeledExpr{
							pos:   position{line: 126, col: 56, offset: 3831},
							label: "ids",
							expr: &zeroOrMoreExpr{
								pos: position{line: 126, col: 60, offset: 3835},
								expr: &seqExpr{
									pos: position{line: 126, col: 61, offset: 3836},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 126, col: 61, offset: 3836},
											name: "VariableName",
										},
										&ruleRefExpr{
											pos:  position{line: 126, col: 74, offset: 3849},
											name: "__",
										},
									},
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 126, col: 79, offset: 3854},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 126, col: 81, offset: 3856},
							val:        "{",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 126, col: 85, offset: 3860},
							name: "__",
						},
						&labeledExpr{
							pos:   position{line: 126, col: 88, offset: 3863},
							label: "statements",
							expr: &oneOrMoreExpr{
								pos: position{line: 126, col: 99, offset: 3874},
								expr: &ruleRefExpr{
									pos:  position{line: 126, col: 100, offset: 3875},
									name: "Statement",
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 126, col: 112, offset: 3887},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 126, col: 114, offset: 3889},
							val:        "}",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 126, col: 118, offset: 3893},
							name: "_",
						},
					},
				},
			},
		},
		{
			name: "Expr",
			pos:  position{line: 149, col: 1, offset: 4568},
			expr: &actionExpr{
				pos: position{line: 149, col: 8, offset: 4575},
				run: (*parser).callonExpr1,
				expr: &labeledExpr{
					pos:   position{line: 149, col: 8, offset: 4575},
					label: "ex",
					expr: &choiceExpr{
						pos: position{line: 149, col: 12, offset: 4579},
						alternatives: []interface{}{
							&ruleRefExpr{
								pos:  position{line: 149, col: 12, offset: 4579},
								name: "IfExpr",
							},
							&ruleRefExpr{
								pos:  position{line: 149, col: 21, offset: 4588},
								name: "Call",
							},
							&ruleRefExpr{
								pos:  position{line: 149, col: 28, offset: 4595},
								name: "CompoundExpr",
							},
						},
					},
				},
			},
		},
		{
			name: "IfExpr",
			pos:  position{line: 154, col: 1, offset: 4690},
			expr: &choiceExpr{
				pos: position{line: 154, col: 10, offset: 4699},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 154, col: 10, offset: 4699},
						run: (*parser).callonIfExpr2,
						expr: &seqExpr{
							pos: position{line: 154, col: 10, offset: 4699},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 154, col: 10, offset: 4699},
									val:        "if",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 154, col: 15, offset: 4704},
									name: "__",
								},
								&labeledExpr{
									pos:   position{line: 154, col: 18, offset: 4707},
									label: "expr",
									expr: &ruleRefExpr{
										pos:  position{line: 154, col: 23, offset: 4712},
										name: "BinOpBool",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 154, col: 33, offset: 4722},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 154, col: 35, offset: 4724},
									val:        "{",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 154, col: 39, offset: 4728},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 154, col: 41, offset: 4730},
									label: "thens",
									expr: &oneOrMoreExpr{
										pos: position{line: 154, col: 47, offset: 4736},
										expr: &ruleRefExpr{
											pos:  position{line: 154, col: 48, offset: 4737},
											name: "Statement",
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 154, col: 60, offset: 4749},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 154, col: 62, offset: 4751},
									val:        "}",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 154, col: 66, offset: 4755},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 154, col: 68, offset: 4757},
									val:        "else",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 154, col: 75, offset: 4764},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 154, col: 77, offset: 4766},
									label: "elseifs",
									expr: &ruleRefExpr{
										pos:  position{line: 154, col: 85, offset: 4774},
										name: "IfExpr",
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 166, col: 1, offset: 5104},
						run: (*parser).callonIfExpr21,
						expr: &seqExpr{
							pos: position{line: 166, col: 1, offset: 5104},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 166, col: 1, offset: 5104},
									val:        "if",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 166, col: 6, offset: 5109},
									name: "__",
								},
								&labeledExpr{
									pos:   position{line: 166, col: 9, offset: 5112},
									label: "expr",
									expr: &ruleRefExpr{
										pos:  position{line: 166, col: 14, offset: 5117},
										name: "BinOpBool",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 166, col: 24, offset: 5127},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 166, col: 26, offset: 5129},
									val:        "{",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 166, col: 30, offset: 5133},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 166, col: 32, offset: 5135},
									label: "thens",
									expr: &oneOrMoreExpr{
										pos: position{line: 166, col: 38, offset: 5141},
										expr: &ruleRefExpr{
											pos:  position{line: 166, col: 39, offset: 5142},
											name: "Statement",
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 166, col: 51, offset: 5154},
									name: "__",
								},
								&litMatcher{
									pos:        position{line: 166, col: 54, offset: 5157},
									val:        "}",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 166, col: 58, offset: 5161},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 166, col: 60, offset: 5163},
									val:        "else",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 166, col: 67, offset: 5170},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 166, col: 69, offset: 5172},
									val:        "{",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 166, col: 73, offset: 5176},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 166, col: 75, offset: 5178},
									label: "elses",
									expr: &oneOrMoreExpr{
										pos: position{line: 166, col: 81, offset: 5184},
										expr: &ruleRefExpr{
											pos:  position{line: 166, col: 82, offset: 5185},
											name: "Statement",
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 166, col: 94, offset: 5197},
									name: "__",
								},
								&litMatcher{
									pos:        position{line: 166, col: 97, offset: 5200},
									val:        "}",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 185, col: 1, offset: 5703},
						run: (*parser).callonIfExpr45,
						expr: &seqExpr{
							pos: position{line: 185, col: 1, offset: 5703},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 185, col: 1, offset: 5703},
									val:        "if",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 185, col: 6, offset: 5708},
									name: "__",
								},
								&labeledExpr{
									pos:   position{line: 185, col: 9, offset: 5711},
									label: "expr",
									expr: &ruleRefExpr{
										pos:  position{line: 185, col: 14, offset: 5716},
										name: "BinOpBool",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 185, col: 24, offset: 5726},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 185, col: 26, offset: 5728},
									val:        "{",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 185, col: 30, offset: 5732},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 185, col: 32, offset: 5734},
									label: "thens",
									expr: &oneOrMoreExpr{
										pos: position{line: 185, col: 38, offset: 5740},
										expr: &ruleRefExpr{
											pos:  position{line: 185, col: 39, offset: 5741},
											name: "Statement",
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 185, col: 51, offset: 5753},
									name: "__",
								},
								&litMatcher{
									pos:        position{line: 185, col: 54, offset: 5756},
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
			pos:  position{line: 198, col: 1, offset: 6055},
			expr: &choiceExpr{
				pos: position{line: 198, col: 8, offset: 6062},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 198, col: 8, offset: 6062},
						run: (*parser).callonCall2,
						expr: &seqExpr{
							pos: position{line: 198, col: 8, offset: 6062},
							exprs: []interface{}{
								&labeledExpr{
									pos:   position{line: 198, col: 8, offset: 6062},
									label: "module",
									expr: &ruleRefExpr{
										pos:  position{line: 198, col: 15, offset: 6069},
										name: "ModuleName",
									},
								},
								&litMatcher{
									pos:        position{line: 198, col: 26, offset: 6080},
									val:        ".",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 198, col: 30, offset: 6084},
									label: "fn",
									expr: &ruleRefExpr{
										pos:  position{line: 198, col: 33, offset: 6087},
										name: "VariableName",
									},
								},
								&labeledExpr{
									pos:   position{line: 198, col: 46, offset: 6100},
									label: "arguments",
									expr: &zeroOrMoreExpr{
										pos: position{line: 198, col: 56, offset: 6110},
										expr: &seqExpr{
											pos: position{line: 198, col: 57, offset: 6111},
											exprs: []interface{}{
												&ruleRefExpr{
													pos:  position{line: 198, col: 57, offset: 6111},
													name: "__",
												},
												&ruleRefExpr{
													pos:  position{line: 198, col: 60, offset: 6114},
													name: "BinOpParens",
												},
											},
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 198, col: 74, offset: 6128},
									name: "_",
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 216, col: 1, offset: 6626},
						run: (*parser).callonCall15,
						expr: &seqExpr{
							pos: position{line: 216, col: 1, offset: 6626},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 216, col: 1, offset: 6626},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 216, col: 3, offset: 6628},
									label: "fn",
									expr: &ruleRefExpr{
										pos:  position{line: 216, col: 6, offset: 6631},
										name: "VariableName",
									},
								},
								&labeledExpr{
									pos:   position{line: 216, col: 19, offset: 6644},
									label: "arguments",
									expr: &oneOrMoreExpr{
										pos: position{line: 216, col: 29, offset: 6654},
										expr: &seqExpr{
											pos: position{line: 216, col: 30, offset: 6655},
											exprs: []interface{}{
												&ruleRefExpr{
													pos:  position{line: 216, col: 30, offset: 6655},
													name: "__",
												},
												&ruleRefExpr{
													pos:  position{line: 216, col: 33, offset: 6658},
													name: "BinOpParens",
												},
											},
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 216, col: 47, offset: 6672},
									name: "_",
								},
							},
						},
					},
				},
			},
		},
		{
			name: "CompoundExpr",
			pos:  position{line: 234, col: 1, offset: 7182},
			expr: &actionExpr{
				pos: position{line: 234, col: 16, offset: 7197},
				run: (*parser).callonCompoundExpr1,
				expr: &seqExpr{
					pos: position{line: 234, col: 16, offset: 7197},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 234, col: 16, offset: 7197},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 234, col: 18, offset: 7199},
							label: "op",
							expr: &ruleRefExpr{
								pos:  position{line: 234, col: 21, offset: 7202},
								name: "BinOp",
							},
						},
						&labeledExpr{
							pos:   position{line: 234, col: 27, offset: 7208},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 234, col: 32, offset: 7213},
								expr: &seqExpr{
									pos: position{line: 234, col: 33, offset: 7214},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 234, col: 33, offset: 7214},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 234, col: 36, offset: 7217},
											name: "Operator",
										},
										&ruleRefExpr{
											pos:  position{line: 234, col: 45, offset: 7226},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 234, col: 48, offset: 7229},
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
			pos:  position{line: 254, col: 1, offset: 7893},
			expr: &choiceExpr{
				pos: position{line: 254, col: 9, offset: 7901},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 254, col: 9, offset: 7901},
						name: "BinOpBool",
					},
					&ruleRefExpr{
						pos:  position{line: 254, col: 21, offset: 7913},
						name: "BinOpEquality",
					},
					&ruleRefExpr{
						pos:  position{line: 254, col: 37, offset: 7929},
						name: "BinOpLow",
					},
					&ruleRefExpr{
						pos:  position{line: 254, col: 48, offset: 7940},
						name: "BinOpHigh",
					},
					&ruleRefExpr{
						pos:  position{line: 254, col: 60, offset: 7952},
						name: "BinOpParens",
					},
				},
			},
		},
		{
			name: "BinOpBool",
			pos:  position{line: 256, col: 1, offset: 7965},
			expr: &actionExpr{
				pos: position{line: 256, col: 13, offset: 7977},
				run: (*parser).callonBinOpBool1,
				expr: &seqExpr{
					pos: position{line: 256, col: 13, offset: 7977},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 256, col: 13, offset: 7977},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 256, col: 15, offset: 7979},
							label: "first",
							expr: &ruleRefExpr{
								pos:  position{line: 256, col: 21, offset: 7985},
								name: "BinOpEquality",
							},
						},
						&labeledExpr{
							pos:   position{line: 256, col: 35, offset: 7999},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 256, col: 40, offset: 8004},
								expr: &seqExpr{
									pos: position{line: 256, col: 41, offset: 8005},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 256, col: 41, offset: 8005},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 256, col: 44, offset: 8008},
											name: "OperatorBoolean",
										},
										&ruleRefExpr{
											pos:  position{line: 256, col: 60, offset: 8024},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 256, col: 63, offset: 8027},
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
			pos:  position{line: 275, col: 1, offset: 8633},
			expr: &actionExpr{
				pos: position{line: 275, col: 17, offset: 8649},
				run: (*parser).callonBinOpEquality1,
				expr: &seqExpr{
					pos: position{line: 275, col: 17, offset: 8649},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 275, col: 17, offset: 8649},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 275, col: 19, offset: 8651},
							label: "first",
							expr: &ruleRefExpr{
								pos:  position{line: 275, col: 25, offset: 8657},
								name: "BinOpLow",
							},
						},
						&labeledExpr{
							pos:   position{line: 275, col: 34, offset: 8666},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 275, col: 39, offset: 8671},
								expr: &seqExpr{
									pos: position{line: 275, col: 40, offset: 8672},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 275, col: 40, offset: 8672},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 275, col: 43, offset: 8675},
											name: "OperatorEquality",
										},
										&ruleRefExpr{
											pos:  position{line: 275, col: 60, offset: 8692},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 275, col: 63, offset: 8695},
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
			pos:  position{line: 295, col: 1, offset: 9299},
			expr: &actionExpr{
				pos: position{line: 295, col: 12, offset: 9310},
				run: (*parser).callonBinOpLow1,
				expr: &seqExpr{
					pos: position{line: 295, col: 12, offset: 9310},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 295, col: 12, offset: 9310},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 295, col: 14, offset: 9312},
							label: "first",
							expr: &ruleRefExpr{
								pos:  position{line: 295, col: 20, offset: 9318},
								name: "BinOpHigh",
							},
						},
						&labeledExpr{
							pos:   position{line: 295, col: 30, offset: 9328},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 295, col: 35, offset: 9333},
								expr: &seqExpr{
									pos: position{line: 295, col: 36, offset: 9334},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 295, col: 36, offset: 9334},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 295, col: 39, offset: 9337},
											name: "OperatorLow",
										},
										&ruleRefExpr{
											pos:  position{line: 295, col: 51, offset: 9349},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 295, col: 54, offset: 9352},
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
			pos:  position{line: 315, col: 1, offset: 9953},
			expr: &actionExpr{
				pos: position{line: 315, col: 13, offset: 9965},
				run: (*parser).callonBinOpHigh1,
				expr: &seqExpr{
					pos: position{line: 315, col: 13, offset: 9965},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 315, col: 13, offset: 9965},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 315, col: 15, offset: 9967},
							label: "first",
							expr: &ruleRefExpr{
								pos:  position{line: 315, col: 21, offset: 9973},
								name: "BinOpParens",
							},
						},
						&labeledExpr{
							pos:   position{line: 315, col: 33, offset: 9985},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 315, col: 38, offset: 9990},
								expr: &seqExpr{
									pos: position{line: 315, col: 39, offset: 9991},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 315, col: 39, offset: 9991},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 315, col: 42, offset: 9994},
											name: "OperatorHigh",
										},
										&ruleRefExpr{
											pos:  position{line: 315, col: 55, offset: 10007},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 315, col: 58, offset: 10010},
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
			pos:  position{line: 334, col: 1, offset: 10614},
			expr: &choiceExpr{
				pos: position{line: 334, col: 15, offset: 10628},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 334, col: 15, offset: 10628},
						run: (*parser).callonBinOpParens2,
						expr: &seqExpr{
							pos: position{line: 334, col: 15, offset: 10628},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 334, col: 15, offset: 10628},
									val:        "(",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 334, col: 19, offset: 10632},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 334, col: 21, offset: 10634},
									label: "first",
									expr: &ruleRefExpr{
										pos:  position{line: 334, col: 27, offset: 10640},
										name: "BinOp",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 334, col: 33, offset: 10646},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 334, col: 35, offset: 10648},
									val:        ")",
									ignoreCase: false,
								},
							},
						},
					},
					&ruleRefExpr{
						pos:  position{line: 337, col: 5, offset: 10797},
						name: "Value",
					},
				},
			},
		},
		{
			name: "Operator",
			pos:  position{line: 339, col: 1, offset: 10804},
			expr: &choiceExpr{
				pos: position{line: 339, col: 12, offset: 10815},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 339, col: 12, offset: 10815},
						name: "OperatorBoolean",
					},
					&ruleRefExpr{
						pos:  position{line: 339, col: 30, offset: 10833},
						name: "OperatorEquality",
					},
					&ruleRefExpr{
						pos:  position{line: 339, col: 49, offset: 10852},
						name: "OperatorHigh",
					},
					&ruleRefExpr{
						pos:  position{line: 339, col: 64, offset: 10867},
						name: "OperatorLow",
					},
				},
			},
		},
		{
			name: "OperatorBoolean",
			pos:  position{line: 341, col: 1, offset: 10880},
			expr: &actionExpr{
				pos: position{line: 341, col: 19, offset: 10898},
				run: (*parser).callonOperatorBoolean1,
				expr: &choiceExpr{
					pos: position{line: 341, col: 21, offset: 10900},
					alternatives: []interface{}{
						&litMatcher{
							pos:        position{line: 341, col: 21, offset: 10900},
							val:        "not",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 341, col: 29, offset: 10908},
							val:        "or",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 341, col: 36, offset: 10915},
							val:        "and",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "OperatorEquality",
			pos:  position{line: 345, col: 1, offset: 11014},
			expr: &actionExpr{
				pos: position{line: 345, col: 20, offset: 11033},
				run: (*parser).callonOperatorEquality1,
				expr: &choiceExpr{
					pos: position{line: 345, col: 22, offset: 11035},
					alternatives: []interface{}{
						&litMatcher{
							pos:        position{line: 345, col: 22, offset: 11035},
							val:        ">=",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 345, col: 29, offset: 11042},
							val:        "<=",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 345, col: 36, offset: 11049},
							val:        "<",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 345, col: 42, offset: 11055},
							val:        ">",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 345, col: 48, offset: 11061},
							val:        "===",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 345, col: 56, offset: 11069},
							val:        "==",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "OperatorHigh",
			pos:  position{line: 349, col: 1, offset: 11175},
			expr: &choiceExpr{
				pos: position{line: 349, col: 16, offset: 11190},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 349, col: 16, offset: 11190},
						run: (*parser).callonOperatorHigh2,
						expr: &choiceExpr{
							pos: position{line: 349, col: 18, offset: 11192},
							alternatives: []interface{}{
								&litMatcher{
									pos:        position{line: 349, col: 18, offset: 11192},
									val:        "/.",
									ignoreCase: false,
								},
								&litMatcher{
									pos:        position{line: 349, col: 25, offset: 11199},
									val:        "*.",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 352, col: 3, offset: 11305},
						run: (*parser).callonOperatorHigh6,
						expr: &choiceExpr{
							pos: position{line: 352, col: 5, offset: 11307},
							alternatives: []interface{}{
								&litMatcher{
									pos:        position{line: 352, col: 5, offset: 11307},
									val:        "*",
									ignoreCase: false,
								},
								&litMatcher{
									pos:        position{line: 352, col: 11, offset: 11313},
									val:        "/",
									ignoreCase: false,
								},
								&litMatcher{
									pos:        position{line: 352, col: 17, offset: 11319},
									val:        "^",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 355, col: 3, offset: 11422},
						run: (*parser).callonOperatorHigh11,
						expr: &litMatcher{
							pos:        position{line: 355, col: 3, offset: 11422},
							val:        "++",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "OperatorLow",
			pos:  position{line: 359, col: 1, offset: 11526},
			expr: &choiceExpr{
				pos: position{line: 359, col: 15, offset: 11540},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 359, col: 15, offset: 11540},
						run: (*parser).callonOperatorLow2,
						expr: &choiceExpr{
							pos: position{line: 359, col: 17, offset: 11542},
							alternatives: []interface{}{
								&litMatcher{
									pos:        position{line: 359, col: 17, offset: 11542},
									val:        "+.",
									ignoreCase: false,
								},
								&litMatcher{
									pos:        position{line: 359, col: 24, offset: 11549},
									val:        "-.",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 362, col: 3, offset: 11655},
						run: (*parser).callonOperatorLow6,
						expr: &choiceExpr{
							pos: position{line: 362, col: 5, offset: 11657},
							alternatives: []interface{}{
								&litMatcher{
									pos:        position{line: 362, col: 5, offset: 11657},
									val:        "+",
									ignoreCase: false,
								},
								&litMatcher{
									pos:        position{line: 362, col: 11, offset: 11663},
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
			pos:  position{line: 366, col: 1, offset: 11765},
			expr: &choiceExpr{
				pos: position{line: 366, col: 9, offset: 11773},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 366, col: 9, offset: 11773},
						name: "VariableName",
					},
					&ruleRefExpr{
						pos:  position{line: 366, col: 24, offset: 11788},
						name: "Const",
					},
				},
			},
		},
		{
			name: "Assignable",
			pos:  position{line: 368, col: 1, offset: 11795},
			expr: &choiceExpr{
				pos: position{line: 368, col: 14, offset: 11808},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 368, col: 14, offset: 11808},
						name: "VariableName",
					},
					&ruleRefExpr{
						pos:  position{line: 368, col: 29, offset: 11823},
						name: "Unused",
					},
				},
			},
		},
		{
			name: "Identifier",
			pos:  position{line: 370, col: 1, offset: 11831},
			expr: &choiceExpr{
				pos: position{line: 370, col: 14, offset: 11844},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 370, col: 14, offset: 11844},
						name: "VariableName",
					},
					&ruleRefExpr{
						pos:  position{line: 370, col: 29, offset: 11859},
						name: "ModuleName",
					},
				},
			},
		},
		{
			name: "ArrayLiteral",
			pos:  position{line: 372, col: 1, offset: 11871},
			expr: &actionExpr{
				pos: position{line: 372, col: 16, offset: 11886},
				run: (*parser).callonArrayLiteral1,
				expr: &seqExpr{
					pos: position{line: 372, col: 16, offset: 11886},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 372, col: 16, offset: 11886},
							val:        "[",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 372, col: 20, offset: 11890},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 372, col: 22, offset: 11892},
							label: "first",
							expr: &ruleRefExpr{
								pos:  position{line: 372, col: 28, offset: 11898},
								name: "Expr",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 372, col: 33, offset: 11903},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 372, col: 35, offset: 11905},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 372, col: 40, offset: 11910},
								expr: &seqExpr{
									pos: position{line: 372, col: 41, offset: 11911},
									exprs: []interface{}{
										&litMatcher{
											pos:        position{line: 372, col: 41, offset: 11911},
											val:        ",",
											ignoreCase: false,
										},
										&ruleRefExpr{
											pos:  position{line: 372, col: 45, offset: 11915},
											name: "_",
										},
										&ruleRefExpr{
											pos:  position{line: 372, col: 47, offset: 11917},
											name: "Expr",
										},
										&ruleRefExpr{
											pos:  position{line: 372, col: 52, offset: 11922},
											name: "_",
										},
									},
								},
							},
						},
						&zeroOrOneExpr{
							pos: position{line: 372, col: 56, offset: 11926},
							expr: &litMatcher{
								pos:        position{line: 372, col: 56, offset: 11926},
								val:        ",",
								ignoreCase: false,
							},
						},
						&ruleRefExpr{
							pos:  position{line: 372, col: 61, offset: 11931},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 372, col: 63, offset: 11933},
							val:        "]",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "TypeParameter",
			pos:  position{line: 388, col: 1, offset: 12414},
			expr: &actionExpr{
				pos: position{line: 388, col: 17, offset: 12430},
				run: (*parser).callonTypeParameter1,
				expr: &seqExpr{
					pos: position{line: 388, col: 17, offset: 12430},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 388, col: 17, offset: 12430},
							val:        "'",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 388, col: 22, offset: 12435},
							name: "VariableName",
						},
					},
				},
			},
		},
		{
			name: "VariableName",
			pos:  position{line: 392, col: 1, offset: 12543},
			expr: &actionExpr{
				pos: position{line: 392, col: 16, offset: 12558},
				run: (*parser).callonVariableName1,
				expr: &seqExpr{
					pos: position{line: 392, col: 16, offset: 12558},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 392, col: 16, offset: 12558},
							expr: &ruleRefExpr{
								pos:  position{line: 392, col: 17, offset: 12559},
								name: "Reserved",
							},
						},
						&seqExpr{
							pos: position{line: 392, col: 27, offset: 12569},
							exprs: []interface{}{
								&oneOrMoreExpr{
									pos: position{line: 392, col: 27, offset: 12569},
									expr: &charClassMatcher{
										pos:        position{line: 392, col: 27, offset: 12569},
										val:        "[a-z]",
										ranges:     []rune{'a', 'z'},
										ignoreCase: false,
										inverted:   false,
									},
								},
								&zeroOrMoreExpr{
									pos: position{line: 392, col: 34, offset: 12576},
									expr: &charClassMatcher{
										pos:        position{line: 392, col: 34, offset: 12576},
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
			pos:  position{line: 396, col: 1, offset: 12687},
			expr: &actionExpr{
				pos: position{line: 396, col: 14, offset: 12700},
				run: (*parser).callonModuleName1,
				expr: &seqExpr{
					pos: position{line: 396, col: 15, offset: 12701},
					exprs: []interface{}{
						&oneOrMoreExpr{
							pos: position{line: 396, col: 15, offset: 12701},
							expr: &charClassMatcher{
								pos:        position{line: 396, col: 15, offset: 12701},
								val:        "[A-Z]",
								ranges:     []rune{'A', 'Z'},
								ignoreCase: false,
								inverted:   false,
							},
						},
						&zeroOrMoreExpr{
							pos: position{line: 396, col: 22, offset: 12708},
							expr: &charClassMatcher{
								pos:        position{line: 396, col: 22, offset: 12708},
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
			pos:  position{line: 400, col: 1, offset: 12819},
			expr: &choiceExpr{
				pos: position{line: 400, col: 9, offset: 12827},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 400, col: 9, offset: 12827},
						run: (*parser).callonConst2,
						expr: &seqExpr{
							pos: position{line: 400, col: 9, offset: 12827},
							exprs: []interface{}{
								&zeroOrOneExpr{
									pos: position{line: 400, col: 9, offset: 12827},
									expr: &litMatcher{
										pos:        position{line: 400, col: 9, offset: 12827},
										val:        "-",
										ignoreCase: false,
									},
								},
								&oneOrMoreExpr{
									pos: position{line: 400, col: 14, offset: 12832},
									expr: &charClassMatcher{
										pos:        position{line: 400, col: 14, offset: 12832},
										val:        "[0-9]",
										ranges:     []rune{'0', '9'},
										ignoreCase: false,
										inverted:   false,
									},
								},
								&notExpr{
									pos: position{line: 400, col: 21, offset: 12839},
									expr: &litMatcher{
										pos:        position{line: 400, col: 22, offset: 12840},
										val:        ".",
										ignoreCase: false,
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 407, col: 3, offset: 13016},
						run: (*parser).callonConst10,
						expr: &seqExpr{
							pos: position{line: 407, col: 3, offset: 13016},
							exprs: []interface{}{
								&zeroOrOneExpr{
									pos: position{line: 407, col: 3, offset: 13016},
									expr: &litMatcher{
										pos:        position{line: 407, col: 3, offset: 13016},
										val:        "-",
										ignoreCase: false,
									},
								},
								&oneOrMoreExpr{
									pos: position{line: 407, col: 8, offset: 13021},
									expr: &charClassMatcher{
										pos:        position{line: 407, col: 8, offset: 13021},
										val:        "[0-9]",
										ranges:     []rune{'0', '9'},
										ignoreCase: false,
										inverted:   false,
									},
								},
								&litMatcher{
									pos:        position{line: 407, col: 15, offset: 13028},
									val:        ".",
									ignoreCase: false,
								},
								&oneOrMoreExpr{
									pos: position{line: 407, col: 19, offset: 13032},
									expr: &charClassMatcher{
										pos:        position{line: 407, col: 19, offset: 13032},
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
						pos:        position{line: 414, col: 3, offset: 13222},
						val:        "true",
						ignoreCase: false,
					},
					&actionExpr{
						pos: position{line: 414, col: 12, offset: 13231},
						run: (*parser).callonConst20,
						expr: &litMatcher{
							pos:        position{line: 414, col: 12, offset: 13231},
							val:        "false",
							ignoreCase: false,
						},
					},
					&actionExpr{
						pos: position{line: 420, col: 3, offset: 13432},
						run: (*parser).callonConst22,
						expr: &litMatcher{
							pos:        position{line: 420, col: 3, offset: 13432},
							val:        "()",
							ignoreCase: false,
						},
					},
					&actionExpr{
						pos: position{line: 423, col: 3, offset: 13495},
						run: (*parser).callonConst24,
						expr: &seqExpr{
							pos: position{line: 423, col: 3, offset: 13495},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 423, col: 3, offset: 13495},
									val:        "\"",
									ignoreCase: false,
								},
								&zeroOrMoreExpr{
									pos: position{line: 423, col: 7, offset: 13499},
									expr: &seqExpr{
										pos: position{line: 423, col: 8, offset: 13500},
										exprs: []interface{}{
											&notExpr{
												pos: position{line: 423, col: 8, offset: 13500},
												expr: &ruleRefExpr{
													pos:  position{line: 423, col: 9, offset: 13501},
													name: "EscapedChar",
												},
											},
											&anyMatcher{
												line: 423, col: 21, offset: 13513,
											},
										},
									},
								},
								&litMatcher{
									pos:        position{line: 423, col: 25, offset: 13517},
									val:        "\"",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 430, col: 3, offset: 13701},
						run: (*parser).callonConst33,
						expr: &seqExpr{
							pos: position{line: 430, col: 3, offset: 13701},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 430, col: 3, offset: 13701},
									val:        "'",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 430, col: 7, offset: 13705},
									label: "val",
									expr: &seqExpr{
										pos: position{line: 430, col: 12, offset: 13710},
										exprs: []interface{}{
											&notExpr{
												pos: position{line: 430, col: 12, offset: 13710},
												expr: &ruleRefExpr{
													pos:  position{line: 430, col: 13, offset: 13711},
													name: "EscapedChar",
												},
											},
											&anyMatcher{
												line: 430, col: 25, offset: 13723,
											},
										},
									},
								},
								&litMatcher{
									pos:        position{line: 430, col: 28, offset: 13726},
									val:        "'",
									ignoreCase: false,
								},
							},
						},
					},
					&ruleRefExpr{
						pos:  position{line: 432, col: 5, offset: 13818},
						name: "ArrayLiteral",
					},
				},
			},
		},
		{
			name: "Unused",
			pos:  position{line: 434, col: 1, offset: 13832},
			expr: &actionExpr{
				pos: position{line: 434, col: 10, offset: 13841},
				run: (*parser).callonUnused1,
				expr: &litMatcher{
					pos:        position{line: 434, col: 11, offset: 13842},
					val:        "_",
					ignoreCase: false,
				},
			},
		},
		{
			name: "Reserved",
			pos:  position{line: 438, col: 1, offset: 13943},
			expr: &seqExpr{
				pos: position{line: 438, col: 12, offset: 13954},
				exprs: []interface{}{
					&choiceExpr{
						pos: position{line: 438, col: 13, offset: 13955},
						alternatives: []interface{}{
							&litMatcher{
								pos:        position{line: 438, col: 13, offset: 13955},
								val:        "let",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 438, col: 21, offset: 13963},
								val:        "if",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 438, col: 28, offset: 13970},
								val:        "else",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 438, col: 37, offset: 13979},
								val:        "func",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 438, col: 46, offset: 13988},
								val:        "type",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 438, col: 55, offset: 13997},
								val:        "true",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 438, col: 64, offset: 14006},
								val:        "false",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 438, col: 74, offset: 14016},
								val:        "mutable",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 438, col: 86, offset: 14028},
								val:        "type",
								ignoreCase: false,
							},
							&ruleRefExpr{
								pos:  position{line: 438, col: 95, offset: 14037},
								name: "BaseType",
							},
						},
					},
					&notExpr{
						pos: position{line: 438, col: 105, offset: 14047},
						expr: &oneOrMoreExpr{
							pos: position{line: 438, col: 106, offset: 14048},
							expr: &charClassMatcher{
								pos:        position{line: 438, col: 106, offset: 14048},
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
			pos:  position{line: 440, col: 1, offset: 14056},
			expr: &actionExpr{
				pos: position{line: 440, col: 12, offset: 14067},
				run: (*parser).callonBaseType1,
				expr: &choiceExpr{
					pos: position{line: 440, col: 14, offset: 14069},
					alternatives: []interface{}{
						&litMatcher{
							pos:        position{line: 440, col: 14, offset: 14069},
							val:        "int",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 440, col: 22, offset: 14077},
							val:        "bool",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 440, col: 31, offset: 14086},
							val:        "string",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 440, col: 42, offset: 14097},
							val:        "byte",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 440, col: 51, offset: 14106},
							val:        "rune",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 440, col: 60, offset: 14115},
							val:        "float",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 440, col: 70, offset: 14125},
							val:        "list",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "EscapedChar",
			pos:  position{line: 444, col: 1, offset: 14224},
			expr: &charClassMatcher{
				pos:        position{line: 444, col: 15, offset: 14238},
				val:        "[\\x00-\\x1f\"\\\\]",
				chars:      []rune{'"', '\\'},
				ranges:     []rune{'\x00', '\x1f'},
				ignoreCase: false,
				inverted:   false,
			},
		},
		{
			name: "EscapeSequence",
			pos:  position{line: 446, col: 1, offset: 14254},
			expr: &choiceExpr{
				pos: position{line: 446, col: 18, offset: 14271},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 446, col: 18, offset: 14271},
						name: "SingleCharEscape",
					},
					&ruleRefExpr{
						pos:  position{line: 446, col: 37, offset: 14290},
						name: "UnicodeEscape",
					},
				},
			},
		},
		{
			name: "SingleCharEscape",
			pos:  position{line: 448, col: 1, offset: 14305},
			expr: &charClassMatcher{
				pos:        position{line: 448, col: 20, offset: 14324},
				val:        "[\"\\\\/bfnrt]",
				chars:      []rune{'"', '\\', '/', 'b', 'f', 'n', 'r', 't'},
				ignoreCase: false,
				inverted:   false,
			},
		},
		{
			name: "DecimalDigit",
			pos:  position{line: 450, col: 1, offset: 14337},
			expr: &charClassMatcher{
				pos:        position{line: 450, col: 16, offset: 14352},
				val:        "[0-9]",
				ranges:     []rune{'0', '9'},
				ignoreCase: false,
				inverted:   false,
			},
		},
		{
			name: "NonZeroDecimalDigit",
			pos:  position{line: 452, col: 1, offset: 14359},
			expr: &charClassMatcher{
				pos:        position{line: 452, col: 23, offset: 14381},
				val:        "[1-9]",
				ranges:     []rune{'1', '9'},
				ignoreCase: false,
				inverted:   false,
			},
		},
		{
			name: "HexDigit",
			pos:  position{line: 454, col: 1, offset: 14388},
			expr: &charClassMatcher{
				pos:        position{line: 454, col: 12, offset: 14399},
				val:        "[0-9a-f]i",
				ranges:     []rune{'0', '9', 'a', 'f'},
				ignoreCase: true,
				inverted:   false,
			},
		},
		{
			name:        "__",
			displayName: "\"sigwhitespace\"",
			pos:         position{line: 456, col: 1, offset: 14410},
			expr: &oneOrMoreExpr{
				pos: position{line: 456, col: 22, offset: 14431},
				expr: &charClassMatcher{
					pos:        position{line: 456, col: 22, offset: 14431},
					val:        "[ \\r\\n\\t]",
					chars:      []rune{' ', '\r', '\n', '\t'},
					ignoreCase: false,
					inverted:   false,
				},
			},
		},
		{
			name:        "_",
			displayName: "\"whitespace\"",
			pos:         position{line: 458, col: 1, offset: 14443},
			expr: &zeroOrMoreExpr{
				pos: position{line: 458, col: 18, offset: 14460},
				expr: &charClassMatcher{
					pos:        position{line: 458, col: 18, offset: 14460},
					val:        "[ \\r\\n\\t]",
					chars:      []rune{' ', '\r', '\n', '\t'},
					ignoreCase: false,
					inverted:   false,
				},
			},
		},
		{
			name: "EOF",
			pos:  position{line: 460, col: 1, offset: 14472},
			expr: &notExpr{
				pos: position{line: 460, col: 7, offset: 14478},
				expr: &anyMatcher{
					line: 460, col: 8, offset: 14479,
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
		return BasicAst{Type: "Module", Subvalues: subvalues, ValueType: CONTAINER}, nil
	} else {
		return BasicAst{Type: "Module", Subvalues: []Ast{stat.(Ast)}, ValueType: CONTAINER}, nil
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
	return AliasType{Name: name.(BasicAst).StringValue}, nil
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

func (c *current) onVariantConstructor1(name, rest interface{}) (interface{}, error) {
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

func (p *parser) callonVariantConstructor1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onVariantConstructor1(stack["name"], stack["rest"])
}

func (c *current) onRecordFieldDefn1(name, t interface{}) (interface{}, error) {
	return RecordField{Name: name.(BasicAst).StringValue, Type: t.(Ast)}, nil
}

func (p *parser) callonRecordFieldDefn1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onRecordFieldDefn1(stack["name"], stack["t"])
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
	return Func{Name: i.(BasicAst).StringValue, Arguments: args, Subvalues: subvalues, ValueType: CONTAINER}, nil
}

func (p *parser) callonFuncDefn1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onFuncDefn1(stack["i"], stack["ids"], stack["statements"])
}

func (c *current) onExpr1(ex interface{}) (interface{}, error) {
	//fmt.Printf("top-level expr: %s\n", string(c.text))
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

func (c *current) onCall2(module, fn, arguments interface{}) (interface{}, error) {
	//fmt.Println("call", string(c.text))

	args := []Ast{}
	vals := arguments.([]interface{})
	if len(vals) > 0 {
		restSl := toIfaceSlice(arguments)
		for _, v := range restSl {
			// we can get each item in the grammar by index
			restExpr := toIfaceSlice(v)
			v := restExpr[1].(Ast)
			args = append(args, v)
		}
	}

	return Call{Module: module.(Ast), Function: fn.(Ast), Arguments: args, ValueType: CONTAINER}, nil
}

func (p *parser) callonCall2() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onCall2(stack["module"], stack["fn"], stack["arguments"])
}

func (c *current) onCall15(fn, arguments interface{}) (interface{}, error) {
	//fmt.Println("call", string(c.text))
	//var mod BasicAst
	args := []Ast{}
	vals := arguments.([]interface{})
	if len(vals) > 0 {
		restSl := toIfaceSlice(arguments)
		for _, v := range restSl {
			// we can get each item in the grammar by index
			restExpr := toIfaceSlice(v)
			v := restExpr[1].(Ast)
			args = append(args, v)
		}
	}

	return Call{Module: nil, Function: fn.(Ast), Arguments: args, ValueType: CONTAINER}, nil
}

func (p *parser) callonCall15() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onCall15(stack["fn"], stack["arguments"])
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

		return BasicAst{Type: "CompoundExpr", Subvalues: subvalues, ValueType: CONTAINER}, nil
	} else {
		return BasicAst{Type: "CompoundExpr", Subvalues: []Ast{op.(Ast)}, ValueType: CONTAINER}, nil
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
		return BasicAst{Type: "BinOpBool", Subvalues: subvalues, ValueType: CONTAINER}, nil
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
		return BasicAst{Type: "BinOpEquality", Subvalues: subvalues, ValueType: CONTAINER}, nil
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
		return BasicAst{Type: "BinOpLow", Subvalues: subvalues, ValueType: CONTAINER}, nil
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
		return BasicAst{Type: "BinOpHigh", Subvalues: subvalues, ValueType: CONTAINER}, nil
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
	return BasicAst{Type: "BinOpParens", Subvalues: []Ast{first.(Ast)}, ValueType: CONTAINER}, nil
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
	return BasicAst{Type: "Array", Subvalues: subvalues, ValueType: CONTAINER}, nil
}

func (p *parser) callonArrayLiteral1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onArrayLiteral1(stack["first"], stack["rest"])
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

func (c *current) onConst22() (interface{}, error) {
	return BasicAst{Type: "Nil", ValueType: NIL}, nil
}

func (p *parser) callonConst22() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onConst22()
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
