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
										name: "VariableName",
									},
								},
								&labeledExpr{
									pos:   position{line: 41, col: 42, offset: 1081},
									label: "params",
									expr: &zeroOrMoreExpr{
										pos: position{line: 41, col: 49, offset: 1088},
										expr: &seqExpr{
											pos: position{line: 41, col: 50, offset: 1089},
											exprs: []interface{}{
												&ruleRefExpr{
													pos:  position{line: 41, col: 50, offset: 1089},
													name: "__",
												},
												&litMatcher{
													pos:        position{line: 41, col: 53, offset: 1092},
													val:        "'",
													ignoreCase: false,
												},
												&ruleRefExpr{
													pos:  position{line: 41, col: 58, offset: 1097},
													name: "VariableName",
												},
											},
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 41, col: 73, offset: 1112},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 41, col: 75, offset: 1114},
									val:        "=",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 41, col: 79, offset: 1118},
									label: "types",
									expr: &oneOrMoreExpr{
										pos: position{line: 41, col: 85, offset: 1124},
										expr: &choiceExpr{
											pos: position{line: 41, col: 86, offset: 1125},
											alternatives: []interface{}{
												&seqExpr{
													pos: position{line: 41, col: 86, offset: 1125},
													exprs: []interface{}{
														&ruleRefExpr{
															pos:  position{line: 41, col: 86, offset: 1125},
															name: "__",
														},
														&ruleRefExpr{
															pos:  position{line: 41, col: 89, offset: 1128},
															name: "BaseType",
														},
													},
												},
												&ruleRefExpr{
													pos:  position{line: 41, col: 100, offset: 1139},
													name: "TypeParameter",
												},
											},
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 41, col: 116, offset: 1155},
									name: "__",
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 46, col: 1, offset: 1220},
						run: (*parser).callonTypeDefn25,
						expr: &seqExpr{
							pos: position{line: 46, col: 1, offset: 1220},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 46, col: 1, offset: 1220},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 46, col: 3, offset: 1222},
									val:        "type",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 46, col: 10, offset: 1229},
									name: "__",
								},
								&labeledExpr{
									pos:   position{line: 46, col: 13, offset: 1232},
									label: "name",
									expr: &ruleRefExpr{
										pos:  position{line: 46, col: 18, offset: 1237},
										name: "VariableName",
									},
								},
								&labeledExpr{
									pos:   position{line: 46, col: 31, offset: 1250},
									label: "params",
									expr: &zeroOrMoreExpr{
										pos: position{line: 46, col: 38, offset: 1257},
										expr: &seqExpr{
											pos: position{line: 46, col: 39, offset: 1258},
											exprs: []interface{}{
												&ruleRefExpr{
													pos:  position{line: 46, col: 39, offset: 1258},
													name: "__",
												},
												&ruleRefExpr{
													pos:  position{line: 46, col: 42, offset: 1261},
													name: "TypeParameter",
												},
											},
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 46, col: 58, offset: 1277},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 46, col: 60, offset: 1279},
									val:        "=",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 46, col: 64, offset: 1283},
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
						pos: position{line: 52, col: 1, offset: 1429},
						run: (*parser).callonTypeDefn57,
						expr: &seqExpr{
							pos: position{line: 52, col: 1, offset: 1429},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 52, col: 1, offset: 1429},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 52, col: 3, offset: 1431},
									val:        "type",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 52, col: 10, offset: 1438},
									name: "__",
								},
								&labeledExpr{
									pos:   position{line: 52, col: 13, offset: 1441},
									label: "name",
									expr: &ruleRefExpr{
										pos:  position{line: 52, col: 18, offset: 1446},
										name: "VariableName",
									},
								},
								&labeledExpr{
									pos:   position{line: 52, col: 31, offset: 1459},
									label: "params",
									expr: &zeroOrMoreExpr{
										pos: position{line: 52, col: 38, offset: 1466},
										expr: &seqExpr{
											pos: position{line: 52, col: 39, offset: 1467},
											exprs: []interface{}{
												&ruleRefExpr{
													pos:  position{line: 52, col: 39, offset: 1467},
													name: "__",
												},
												&ruleRefExpr{
													pos:  position{line: 52, col: 42, offset: 1470},
													name: "TypeParameter",
												},
											},
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 52, col: 58, offset: 1486},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 52, col: 60, offset: 1488},
									val:        "=",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 52, col: 64, offset: 1492},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 52, col: 66, offset: 1494},
									label: "rest",
									expr: &oneOrMoreExpr{
										pos: position{line: 52, col: 71, offset: 1499},
										expr: &ruleRefExpr{
											pos:  position{line: 52, col: 72, offset: 1500},
											name: "VariantConstructor",
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 52, col: 93, offset: 1521},
									name: "__",
								},
							},
						},
					},
				},
			},
		},
		{
			name: "VariantConstructor",
			pos:  position{line: 57, col: 1, offset: 1589},
			expr: &actionExpr{
				pos: position{line: 57, col: 22, offset: 1610},
				run: (*parser).callonVariantConstructor1,
				expr: &seqExpr{
					pos: position{line: 57, col: 22, offset: 1610},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 57, col: 22, offset: 1610},
							val:        "|",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 57, col: 26, offset: 1614},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 57, col: 28, offset: 1616},
							label: "name",
							expr: &ruleRefExpr{
								pos:  position{line: 57, col: 33, offset: 1621},
								name: "ModuleName",
							},
						},
						&labeledExpr{
							pos:   position{line: 57, col: 44, offset: 1632},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 57, col: 49, offset: 1637},
								expr: &choiceExpr{
									pos: position{line: 57, col: 50, offset: 1638},
									alternatives: []interface{}{
										&seqExpr{
											pos: position{line: 57, col: 50, offset: 1638},
											exprs: []interface{}{
												&ruleRefExpr{
													pos:  position{line: 57, col: 50, offset: 1638},
													name: "__",
												},
												&ruleRefExpr{
													pos:  position{line: 57, col: 53, offset: 1641},
													name: "BaseType",
												},
											},
										},
										&ruleRefExpr{
											pos:  position{line: 57, col: 64, offset: 1652},
											name: "TypeParameter",
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
			pos:  position{line: 61, col: 1, offset: 1710},
			expr: &actionExpr{
				pos: position{line: 61, col: 19, offset: 1728},
				run: (*parser).callonRecordFieldDefn1,
				expr: &seqExpr{
					pos: position{line: 61, col: 19, offset: 1728},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 61, col: 19, offset: 1728},
							label: "name",
							expr: &ruleRefExpr{
								pos:  position{line: 61, col: 24, offset: 1733},
								name: "VariableName",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 61, col: 37, offset: 1746},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 61, col: 39, offset: 1748},
							val:        ":",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 61, col: 43, offset: 1752},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 61, col: 45, offset: 1754},
							label: "t",
							expr: &choiceExpr{
								pos: position{line: 61, col: 48, offset: 1757},
								alternatives: []interface{}{
									&ruleRefExpr{
										pos:  position{line: 61, col: 48, offset: 1757},
										name: "BaseType",
									},
									&ruleRefExpr{
										pos:  position{line: 61, col: 59, offset: 1768},
										name: "TypeParameter",
									},
								},
							},
						},
					},
				},
			},
		},
		{
			name: "Assignment",
			pos:  position{line: 65, col: 1, offset: 1818},
			expr: &choiceExpr{
				pos: position{line: 65, col: 14, offset: 1831},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 65, col: 14, offset: 1831},
						run: (*parser).callonAssignment2,
						expr: &seqExpr{
							pos: position{line: 65, col: 14, offset: 1831},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 65, col: 14, offset: 1831},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 65, col: 16, offset: 1833},
									val:        "let",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 65, col: 22, offset: 1839},
									name: "__",
								},
								&labeledExpr{
									pos:   position{line: 65, col: 25, offset: 1842},
									label: "i",
									expr: &ruleRefExpr{
										pos:  position{line: 65, col: 27, offset: 1844},
										name: "Assignable",
									},
								},
								&labeledExpr{
									pos:   position{line: 65, col: 38, offset: 1855},
									label: "rest",
									expr: &zeroOrMoreExpr{
										pos: position{line: 65, col: 43, offset: 1860},
										expr: &seqExpr{
											pos: position{line: 65, col: 44, offset: 1861},
											exprs: []interface{}{
												&litMatcher{
													pos:        position{line: 65, col: 44, offset: 1861},
													val:        ",",
													ignoreCase: false,
												},
												&ruleRefExpr{
													pos:  position{line: 65, col: 48, offset: 1865},
													name: "_",
												},
												&ruleRefExpr{
													pos:  position{line: 65, col: 50, offset: 1867},
													name: "Assignable",
												},
											},
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 65, col: 63, offset: 1880},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 65, col: 65, offset: 1882},
									val:        "=",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 65, col: 69, offset: 1886},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 65, col: 71, offset: 1888},
									label: "expr",
									expr: &ruleRefExpr{
										pos:  position{line: 65, col: 76, offset: 1893},
										name: "Expr",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 65, col: 81, offset: 1898},
									name: "_",
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 80, col: 1, offset: 2337},
						run: (*parser).callonAssignment21,
						expr: &seqExpr{
							pos: position{line: 80, col: 1, offset: 2337},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 80, col: 1, offset: 2337},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 80, col: 3, offset: 2339},
									val:        "let",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 80, col: 9, offset: 2345},
									name: "__",
								},
								&notExpr{
									pos: position{line: 80, col: 12, offset: 2348},
									expr: &ruleRefExpr{
										pos:  position{line: 80, col: 13, offset: 2349},
										name: "Assignable",
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 84, col: 1, offset: 2457},
						run: (*parser).callonAssignment28,
						expr: &seqExpr{
							pos: position{line: 84, col: 1, offset: 2457},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 84, col: 1, offset: 2457},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 84, col: 3, offset: 2459},
									val:        "let",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 84, col: 9, offset: 2465},
									name: "__",
								},
								&labeledExpr{
									pos:   position{line: 84, col: 12, offset: 2468},
									label: "i",
									expr: &ruleRefExpr{
										pos:  position{line: 84, col: 14, offset: 2470},
										name: "Assignable",
									},
								},
								&labeledExpr{
									pos:   position{line: 84, col: 25, offset: 2481},
									label: "rest",
									expr: &zeroOrMoreExpr{
										pos: position{line: 84, col: 30, offset: 2486},
										expr: &seqExpr{
											pos: position{line: 84, col: 31, offset: 2487},
											exprs: []interface{}{
												&litMatcher{
													pos:        position{line: 84, col: 31, offset: 2487},
													val:        ",",
													ignoreCase: false,
												},
												&ruleRefExpr{
													pos:  position{line: 84, col: 35, offset: 2491},
													name: "_",
												},
												&ruleRefExpr{
													pos:  position{line: 84, col: 37, offset: 2493},
													name: "Assignable",
												},
											},
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 84, col: 50, offset: 2506},
									name: "_",
								},
								&notExpr{
									pos: position{line: 84, col: 52, offset: 2508},
									expr: &litMatcher{
										pos:        position{line: 84, col: 53, offset: 2509},
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
			pos:  position{line: 88, col: 1, offset: 2603},
			expr: &actionExpr{
				pos: position{line: 88, col: 12, offset: 2614},
				run: (*parser).callonFuncDefn1,
				expr: &seqExpr{
					pos: position{line: 88, col: 12, offset: 2614},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 88, col: 12, offset: 2614},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 88, col: 14, offset: 2616},
							val:        "let",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 88, col: 20, offset: 2622},
							name: "__",
						},
						&labeledExpr{
							pos:   position{line: 88, col: 23, offset: 2625},
							label: "i",
							expr: &ruleRefExpr{
								pos:  position{line: 88, col: 25, offset: 2627},
								name: "VariableName",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 88, col: 38, offset: 2640},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 88, col: 40, offset: 2642},
							val:        "=",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 88, col: 44, offset: 2646},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 88, col: 46, offset: 2648},
							val:        "func",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 88, col: 53, offset: 2655},
							name: "__",
						},
						&labeledExpr{
							pos:   position{line: 88, col: 56, offset: 2658},
							label: "ids",
							expr: &zeroOrMoreExpr{
								pos: position{line: 88, col: 60, offset: 2662},
								expr: &seqExpr{
									pos: position{line: 88, col: 61, offset: 2663},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 88, col: 61, offset: 2663},
											name: "VariableName",
										},
										&ruleRefExpr{
											pos:  position{line: 88, col: 74, offset: 2676},
											name: "__",
										},
									},
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 88, col: 79, offset: 2681},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 88, col: 81, offset: 2683},
							val:        "{",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 88, col: 85, offset: 2687},
							name: "__",
						},
						&labeledExpr{
							pos:   position{line: 88, col: 88, offset: 2690},
							label: "statements",
							expr: &oneOrMoreExpr{
								pos: position{line: 88, col: 99, offset: 2701},
								expr: &ruleRefExpr{
									pos:  position{line: 88, col: 100, offset: 2702},
									name: "Statement",
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 88, col: 112, offset: 2714},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 88, col: 114, offset: 2716},
							val:        "}",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 88, col: 118, offset: 2720},
							name: "_",
						},
					},
				},
			},
		},
		{
			name: "Expr",
			pos:  position{line: 111, col: 1, offset: 3395},
			expr: &actionExpr{
				pos: position{line: 111, col: 8, offset: 3402},
				run: (*parser).callonExpr1,
				expr: &labeledExpr{
					pos:   position{line: 111, col: 8, offset: 3402},
					label: "ex",
					expr: &choiceExpr{
						pos: position{line: 111, col: 12, offset: 3406},
						alternatives: []interface{}{
							&ruleRefExpr{
								pos:  position{line: 111, col: 12, offset: 3406},
								name: "IfExpr",
							},
							&ruleRefExpr{
								pos:  position{line: 111, col: 21, offset: 3415},
								name: "Call",
							},
							&ruleRefExpr{
								pos:  position{line: 111, col: 28, offset: 3422},
								name: "CompoundExpr",
							},
						},
					},
				},
			},
		},
		{
			name: "IfExpr",
			pos:  position{line: 116, col: 1, offset: 3517},
			expr: &choiceExpr{
				pos: position{line: 116, col: 10, offset: 3526},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 116, col: 10, offset: 3526},
						run: (*parser).callonIfExpr2,
						expr: &seqExpr{
							pos: position{line: 116, col: 10, offset: 3526},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 116, col: 10, offset: 3526},
									val:        "if",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 116, col: 15, offset: 3531},
									name: "__",
								},
								&labeledExpr{
									pos:   position{line: 116, col: 18, offset: 3534},
									label: "expr",
									expr: &ruleRefExpr{
										pos:  position{line: 116, col: 23, offset: 3539},
										name: "BinOpBool",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 116, col: 33, offset: 3549},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 116, col: 35, offset: 3551},
									val:        "{",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 116, col: 39, offset: 3555},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 116, col: 41, offset: 3557},
									label: "thens",
									expr: &oneOrMoreExpr{
										pos: position{line: 116, col: 47, offset: 3563},
										expr: &ruleRefExpr{
											pos:  position{line: 116, col: 48, offset: 3564},
											name: "Statement",
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 116, col: 60, offset: 3576},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 116, col: 62, offset: 3578},
									val:        "}",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 116, col: 66, offset: 3582},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 116, col: 68, offset: 3584},
									val:        "else",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 116, col: 75, offset: 3591},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 116, col: 77, offset: 3593},
									label: "elseifs",
									expr: &ruleRefExpr{
										pos:  position{line: 116, col: 85, offset: 3601},
										name: "IfExpr",
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 128, col: 1, offset: 3931},
						run: (*parser).callonIfExpr21,
						expr: &seqExpr{
							pos: position{line: 128, col: 1, offset: 3931},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 128, col: 1, offset: 3931},
									val:        "if",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 128, col: 6, offset: 3936},
									name: "__",
								},
								&labeledExpr{
									pos:   position{line: 128, col: 9, offset: 3939},
									label: "expr",
									expr: &ruleRefExpr{
										pos:  position{line: 128, col: 14, offset: 3944},
										name: "BinOpBool",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 128, col: 24, offset: 3954},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 128, col: 26, offset: 3956},
									val:        "{",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 128, col: 30, offset: 3960},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 128, col: 32, offset: 3962},
									label: "thens",
									expr: &oneOrMoreExpr{
										pos: position{line: 128, col: 38, offset: 3968},
										expr: &ruleRefExpr{
											pos:  position{line: 128, col: 39, offset: 3969},
											name: "Statement",
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 128, col: 51, offset: 3981},
									name: "__",
								},
								&litMatcher{
									pos:        position{line: 128, col: 54, offset: 3984},
									val:        "}",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 128, col: 58, offset: 3988},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 128, col: 60, offset: 3990},
									val:        "else",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 128, col: 67, offset: 3997},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 128, col: 69, offset: 3999},
									val:        "{",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 128, col: 73, offset: 4003},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 128, col: 75, offset: 4005},
									label: "elses",
									expr: &oneOrMoreExpr{
										pos: position{line: 128, col: 81, offset: 4011},
										expr: &ruleRefExpr{
											pos:  position{line: 128, col: 82, offset: 4012},
											name: "Statement",
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 128, col: 94, offset: 4024},
									name: "__",
								},
								&litMatcher{
									pos:        position{line: 128, col: 97, offset: 4027},
									val:        "}",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 147, col: 1, offset: 4530},
						run: (*parser).callonIfExpr45,
						expr: &seqExpr{
							pos: position{line: 147, col: 1, offset: 4530},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 147, col: 1, offset: 4530},
									val:        "if",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 147, col: 6, offset: 4535},
									name: "__",
								},
								&labeledExpr{
									pos:   position{line: 147, col: 9, offset: 4538},
									label: "expr",
									expr: &ruleRefExpr{
										pos:  position{line: 147, col: 14, offset: 4543},
										name: "BinOpBool",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 147, col: 24, offset: 4553},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 147, col: 26, offset: 4555},
									val:        "{",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 147, col: 30, offset: 4559},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 147, col: 32, offset: 4561},
									label: "thens",
									expr: &oneOrMoreExpr{
										pos: position{line: 147, col: 38, offset: 4567},
										expr: &ruleRefExpr{
											pos:  position{line: 147, col: 39, offset: 4568},
											name: "Statement",
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 147, col: 51, offset: 4580},
									name: "__",
								},
								&litMatcher{
									pos:        position{line: 147, col: 54, offset: 4583},
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
			pos:  position{line: 160, col: 1, offset: 4882},
			expr: &choiceExpr{
				pos: position{line: 160, col: 8, offset: 4889},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 160, col: 8, offset: 4889},
						run: (*parser).callonCall2,
						expr: &seqExpr{
							pos: position{line: 160, col: 8, offset: 4889},
							exprs: []interface{}{
								&labeledExpr{
									pos:   position{line: 160, col: 8, offset: 4889},
									label: "module",
									expr: &ruleRefExpr{
										pos:  position{line: 160, col: 15, offset: 4896},
										name: "ModuleName",
									},
								},
								&litMatcher{
									pos:        position{line: 160, col: 26, offset: 4907},
									val:        ".",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 160, col: 30, offset: 4911},
									label: "fn",
									expr: &ruleRefExpr{
										pos:  position{line: 160, col: 33, offset: 4914},
										name: "VariableName",
									},
								},
								&labeledExpr{
									pos:   position{line: 160, col: 46, offset: 4927},
									label: "arguments",
									expr: &zeroOrMoreExpr{
										pos: position{line: 160, col: 56, offset: 4937},
										expr: &seqExpr{
											pos: position{line: 160, col: 57, offset: 4938},
											exprs: []interface{}{
												&ruleRefExpr{
													pos:  position{line: 160, col: 57, offset: 4938},
													name: "__",
												},
												&ruleRefExpr{
													pos:  position{line: 160, col: 60, offset: 4941},
													name: "BinOpParens",
												},
											},
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 160, col: 74, offset: 4955},
									name: "_",
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 178, col: 1, offset: 5453},
						run: (*parser).callonCall15,
						expr: &seqExpr{
							pos: position{line: 178, col: 1, offset: 5453},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 178, col: 1, offset: 5453},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 178, col: 3, offset: 5455},
									label: "fn",
									expr: &ruleRefExpr{
										pos:  position{line: 178, col: 6, offset: 5458},
										name: "VariableName",
									},
								},
								&labeledExpr{
									pos:   position{line: 178, col: 19, offset: 5471},
									label: "arguments",
									expr: &oneOrMoreExpr{
										pos: position{line: 178, col: 29, offset: 5481},
										expr: &seqExpr{
											pos: position{line: 178, col: 30, offset: 5482},
											exprs: []interface{}{
												&ruleRefExpr{
													pos:  position{line: 178, col: 30, offset: 5482},
													name: "__",
												},
												&ruleRefExpr{
													pos:  position{line: 178, col: 33, offset: 5485},
													name: "BinOpParens",
												},
											},
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 178, col: 47, offset: 5499},
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
			pos:  position{line: 196, col: 1, offset: 6009},
			expr: &actionExpr{
				pos: position{line: 196, col: 16, offset: 6024},
				run: (*parser).callonCompoundExpr1,
				expr: &seqExpr{
					pos: position{line: 196, col: 16, offset: 6024},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 196, col: 16, offset: 6024},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 196, col: 18, offset: 6026},
							label: "op",
							expr: &ruleRefExpr{
								pos:  position{line: 196, col: 21, offset: 6029},
								name: "BinOp",
							},
						},
						&labeledExpr{
							pos:   position{line: 196, col: 27, offset: 6035},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 196, col: 32, offset: 6040},
								expr: &seqExpr{
									pos: position{line: 196, col: 33, offset: 6041},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 196, col: 33, offset: 6041},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 196, col: 36, offset: 6044},
											name: "Operator",
										},
										&ruleRefExpr{
											pos:  position{line: 196, col: 45, offset: 6053},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 196, col: 48, offset: 6056},
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
			pos:  position{line: 216, col: 1, offset: 6720},
			expr: &choiceExpr{
				pos: position{line: 216, col: 9, offset: 6728},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 216, col: 9, offset: 6728},
						name: "BinOpBool",
					},
					&ruleRefExpr{
						pos:  position{line: 216, col: 21, offset: 6740},
						name: "BinOpEquality",
					},
					&ruleRefExpr{
						pos:  position{line: 216, col: 37, offset: 6756},
						name: "BinOpLow",
					},
					&ruleRefExpr{
						pos:  position{line: 216, col: 48, offset: 6767},
						name: "BinOpHigh",
					},
					&ruleRefExpr{
						pos:  position{line: 216, col: 60, offset: 6779},
						name: "BinOpParens",
					},
				},
			},
		},
		{
			name: "BinOpBool",
			pos:  position{line: 218, col: 1, offset: 6792},
			expr: &actionExpr{
				pos: position{line: 218, col: 13, offset: 6804},
				run: (*parser).callonBinOpBool1,
				expr: &seqExpr{
					pos: position{line: 218, col: 13, offset: 6804},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 218, col: 13, offset: 6804},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 218, col: 15, offset: 6806},
							label: "first",
							expr: &ruleRefExpr{
								pos:  position{line: 218, col: 21, offset: 6812},
								name: "BinOpEquality",
							},
						},
						&labeledExpr{
							pos:   position{line: 218, col: 35, offset: 6826},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 218, col: 40, offset: 6831},
								expr: &seqExpr{
									pos: position{line: 218, col: 41, offset: 6832},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 218, col: 41, offset: 6832},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 218, col: 44, offset: 6835},
											name: "OperatorBoolean",
										},
										&ruleRefExpr{
											pos:  position{line: 218, col: 60, offset: 6851},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 218, col: 63, offset: 6854},
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
			pos:  position{line: 237, col: 1, offset: 7460},
			expr: &actionExpr{
				pos: position{line: 237, col: 17, offset: 7476},
				run: (*parser).callonBinOpEquality1,
				expr: &seqExpr{
					pos: position{line: 237, col: 17, offset: 7476},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 237, col: 17, offset: 7476},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 237, col: 19, offset: 7478},
							label: "first",
							expr: &ruleRefExpr{
								pos:  position{line: 237, col: 25, offset: 7484},
								name: "BinOpLow",
							},
						},
						&labeledExpr{
							pos:   position{line: 237, col: 34, offset: 7493},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 237, col: 39, offset: 7498},
								expr: &seqExpr{
									pos: position{line: 237, col: 40, offset: 7499},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 237, col: 40, offset: 7499},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 237, col: 43, offset: 7502},
											name: "OperatorEquality",
										},
										&ruleRefExpr{
											pos:  position{line: 237, col: 60, offset: 7519},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 237, col: 63, offset: 7522},
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
			pos:  position{line: 257, col: 1, offset: 8126},
			expr: &actionExpr{
				pos: position{line: 257, col: 12, offset: 8137},
				run: (*parser).callonBinOpLow1,
				expr: &seqExpr{
					pos: position{line: 257, col: 12, offset: 8137},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 257, col: 12, offset: 8137},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 257, col: 14, offset: 8139},
							label: "first",
							expr: &ruleRefExpr{
								pos:  position{line: 257, col: 20, offset: 8145},
								name: "BinOpHigh",
							},
						},
						&labeledExpr{
							pos:   position{line: 257, col: 30, offset: 8155},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 257, col: 35, offset: 8160},
								expr: &seqExpr{
									pos: position{line: 257, col: 36, offset: 8161},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 257, col: 36, offset: 8161},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 257, col: 39, offset: 8164},
											name: "OperatorLow",
										},
										&ruleRefExpr{
											pos:  position{line: 257, col: 51, offset: 8176},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 257, col: 54, offset: 8179},
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
			pos:  position{line: 277, col: 1, offset: 8780},
			expr: &actionExpr{
				pos: position{line: 277, col: 13, offset: 8792},
				run: (*parser).callonBinOpHigh1,
				expr: &seqExpr{
					pos: position{line: 277, col: 13, offset: 8792},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 277, col: 13, offset: 8792},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 277, col: 15, offset: 8794},
							label: "first",
							expr: &ruleRefExpr{
								pos:  position{line: 277, col: 21, offset: 8800},
								name: "BinOpParens",
							},
						},
						&labeledExpr{
							pos:   position{line: 277, col: 33, offset: 8812},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 277, col: 38, offset: 8817},
								expr: &seqExpr{
									pos: position{line: 277, col: 39, offset: 8818},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 277, col: 39, offset: 8818},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 277, col: 42, offset: 8821},
											name: "OperatorHigh",
										},
										&ruleRefExpr{
											pos:  position{line: 277, col: 55, offset: 8834},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 277, col: 58, offset: 8837},
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
			pos:  position{line: 296, col: 1, offset: 9441},
			expr: &choiceExpr{
				pos: position{line: 296, col: 15, offset: 9455},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 296, col: 15, offset: 9455},
						run: (*parser).callonBinOpParens2,
						expr: &seqExpr{
							pos: position{line: 296, col: 15, offset: 9455},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 296, col: 15, offset: 9455},
									val:        "(",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 296, col: 19, offset: 9459},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 296, col: 21, offset: 9461},
									label: "first",
									expr: &ruleRefExpr{
										pos:  position{line: 296, col: 27, offset: 9467},
										name: "BinOp",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 296, col: 33, offset: 9473},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 296, col: 35, offset: 9475},
									val:        ")",
									ignoreCase: false,
								},
							},
						},
					},
					&ruleRefExpr{
						pos:  position{line: 299, col: 5, offset: 9624},
						name: "Value",
					},
				},
			},
		},
		{
			name: "Operator",
			pos:  position{line: 301, col: 1, offset: 9631},
			expr: &choiceExpr{
				pos: position{line: 301, col: 12, offset: 9642},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 301, col: 12, offset: 9642},
						name: "OperatorBoolean",
					},
					&ruleRefExpr{
						pos:  position{line: 301, col: 30, offset: 9660},
						name: "OperatorEquality",
					},
					&ruleRefExpr{
						pos:  position{line: 301, col: 49, offset: 9679},
						name: "OperatorHigh",
					},
					&ruleRefExpr{
						pos:  position{line: 301, col: 64, offset: 9694},
						name: "OperatorLow",
					},
				},
			},
		},
		{
			name: "OperatorBoolean",
			pos:  position{line: 303, col: 1, offset: 9707},
			expr: &actionExpr{
				pos: position{line: 303, col: 19, offset: 9725},
				run: (*parser).callonOperatorBoolean1,
				expr: &choiceExpr{
					pos: position{line: 303, col: 21, offset: 9727},
					alternatives: []interface{}{
						&litMatcher{
							pos:        position{line: 303, col: 21, offset: 9727},
							val:        "not",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 303, col: 29, offset: 9735},
							val:        "or",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 303, col: 36, offset: 9742},
							val:        "and",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "OperatorEquality",
			pos:  position{line: 307, col: 1, offset: 9841},
			expr: &actionExpr{
				pos: position{line: 307, col: 20, offset: 9860},
				run: (*parser).callonOperatorEquality1,
				expr: &choiceExpr{
					pos: position{line: 307, col: 22, offset: 9862},
					alternatives: []interface{}{
						&litMatcher{
							pos:        position{line: 307, col: 22, offset: 9862},
							val:        ">=",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 307, col: 29, offset: 9869},
							val:        "<=",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 307, col: 36, offset: 9876},
							val:        "<",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 307, col: 42, offset: 9882},
							val:        ">",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 307, col: 48, offset: 9888},
							val:        "===",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 307, col: 56, offset: 9896},
							val:        "==",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "OperatorHigh",
			pos:  position{line: 311, col: 1, offset: 10002},
			expr: &choiceExpr{
				pos: position{line: 311, col: 16, offset: 10017},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 311, col: 16, offset: 10017},
						run: (*parser).callonOperatorHigh2,
						expr: &choiceExpr{
							pos: position{line: 311, col: 18, offset: 10019},
							alternatives: []interface{}{
								&litMatcher{
									pos:        position{line: 311, col: 18, offset: 10019},
									val:        "/.",
									ignoreCase: false,
								},
								&litMatcher{
									pos:        position{line: 311, col: 25, offset: 10026},
									val:        "*.",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 314, col: 3, offset: 10132},
						run: (*parser).callonOperatorHigh6,
						expr: &choiceExpr{
							pos: position{line: 314, col: 5, offset: 10134},
							alternatives: []interface{}{
								&litMatcher{
									pos:        position{line: 314, col: 5, offset: 10134},
									val:        "*",
									ignoreCase: false,
								},
								&litMatcher{
									pos:        position{line: 314, col: 11, offset: 10140},
									val:        "/",
									ignoreCase: false,
								},
								&litMatcher{
									pos:        position{line: 314, col: 17, offset: 10146},
									val:        "^",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 317, col: 3, offset: 10249},
						run: (*parser).callonOperatorHigh11,
						expr: &litMatcher{
							pos:        position{line: 317, col: 3, offset: 10249},
							val:        "++",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "OperatorLow",
			pos:  position{line: 321, col: 1, offset: 10353},
			expr: &choiceExpr{
				pos: position{line: 321, col: 15, offset: 10367},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 321, col: 15, offset: 10367},
						run: (*parser).callonOperatorLow2,
						expr: &choiceExpr{
							pos: position{line: 321, col: 17, offset: 10369},
							alternatives: []interface{}{
								&litMatcher{
									pos:        position{line: 321, col: 17, offset: 10369},
									val:        "+.",
									ignoreCase: false,
								},
								&litMatcher{
									pos:        position{line: 321, col: 24, offset: 10376},
									val:        "-.",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 324, col: 3, offset: 10482},
						run: (*parser).callonOperatorLow6,
						expr: &choiceExpr{
							pos: position{line: 324, col: 5, offset: 10484},
							alternatives: []interface{}{
								&litMatcher{
									pos:        position{line: 324, col: 5, offset: 10484},
									val:        "+",
									ignoreCase: false,
								},
								&litMatcher{
									pos:        position{line: 324, col: 11, offset: 10490},
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
			pos:  position{line: 328, col: 1, offset: 10592},
			expr: &choiceExpr{
				pos: position{line: 328, col: 9, offset: 10600},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 328, col: 9, offset: 10600},
						name: "VariableName",
					},
					&ruleRefExpr{
						pos:  position{line: 328, col: 24, offset: 10615},
						name: "Const",
					},
				},
			},
		},
		{
			name: "Assignable",
			pos:  position{line: 330, col: 1, offset: 10622},
			expr: &choiceExpr{
				pos: position{line: 330, col: 14, offset: 10635},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 330, col: 14, offset: 10635},
						name: "VariableName",
					},
					&ruleRefExpr{
						pos:  position{line: 330, col: 29, offset: 10650},
						name: "Unused",
					},
				},
			},
		},
		{
			name: "Identifier",
			pos:  position{line: 332, col: 1, offset: 10658},
			expr: &choiceExpr{
				pos: position{line: 332, col: 14, offset: 10671},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 332, col: 14, offset: 10671},
						name: "VariableName",
					},
					&ruleRefExpr{
						pos:  position{line: 332, col: 29, offset: 10686},
						name: "ModuleName",
					},
				},
			},
		},
		{
			name: "ArrayLiteral",
			pos:  position{line: 334, col: 1, offset: 10698},
			expr: &actionExpr{
				pos: position{line: 334, col: 16, offset: 10713},
				run: (*parser).callonArrayLiteral1,
				expr: &seqExpr{
					pos: position{line: 334, col: 16, offset: 10713},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 334, col: 16, offset: 10713},
							val:        "[",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 334, col: 20, offset: 10717},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 334, col: 22, offset: 10719},
							label: "first",
							expr: &ruleRefExpr{
								pos:  position{line: 334, col: 28, offset: 10725},
								name: "Expr",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 334, col: 33, offset: 10730},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 334, col: 35, offset: 10732},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 334, col: 40, offset: 10737},
								expr: &seqExpr{
									pos: position{line: 334, col: 41, offset: 10738},
									exprs: []interface{}{
										&litMatcher{
											pos:        position{line: 334, col: 41, offset: 10738},
											val:        ",",
											ignoreCase: false,
										},
										&ruleRefExpr{
											pos:  position{line: 334, col: 45, offset: 10742},
											name: "_",
										},
										&ruleRefExpr{
											pos:  position{line: 334, col: 47, offset: 10744},
											name: "Expr",
										},
										&ruleRefExpr{
											pos:  position{line: 334, col: 52, offset: 10749},
											name: "_",
										},
									},
								},
							},
						},
						&zeroOrOneExpr{
							pos: position{line: 334, col: 56, offset: 10753},
							expr: &litMatcher{
								pos:        position{line: 334, col: 56, offset: 10753},
								val:        ",",
								ignoreCase: false,
							},
						},
						&ruleRefExpr{
							pos:  position{line: 334, col: 61, offset: 10758},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 334, col: 63, offset: 10760},
							val:        "]",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "TypeParameter",
			pos:  position{line: 350, col: 1, offset: 11241},
			expr: &actionExpr{
				pos: position{line: 350, col: 17, offset: 11257},
				run: (*parser).callonTypeParameter1,
				expr: &seqExpr{
					pos: position{line: 350, col: 17, offset: 11257},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 350, col: 17, offset: 11257},
							val:        "'",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 350, col: 22, offset: 11262},
							name: "VariableName",
						},
					},
				},
			},
		},
		{
			name: "VariableName",
			pos:  position{line: 354, col: 1, offset: 11370},
			expr: &actionExpr{
				pos: position{line: 354, col: 16, offset: 11385},
				run: (*parser).callonVariableName1,
				expr: &seqExpr{
					pos: position{line: 354, col: 16, offset: 11385},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 354, col: 16, offset: 11385},
							expr: &ruleRefExpr{
								pos:  position{line: 354, col: 17, offset: 11386},
								name: "Reserved",
							},
						},
						&seqExpr{
							pos: position{line: 354, col: 27, offset: 11396},
							exprs: []interface{}{
								&oneOrMoreExpr{
									pos: position{line: 354, col: 27, offset: 11396},
									expr: &charClassMatcher{
										pos:        position{line: 354, col: 27, offset: 11396},
										val:        "[a-z]",
										ranges:     []rune{'a', 'z'},
										ignoreCase: false,
										inverted:   false,
									},
								},
								&zeroOrMoreExpr{
									pos: position{line: 354, col: 34, offset: 11403},
									expr: &charClassMatcher{
										pos:        position{line: 354, col: 34, offset: 11403},
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
			pos:  position{line: 358, col: 1, offset: 11514},
			expr: &actionExpr{
				pos: position{line: 358, col: 14, offset: 11527},
				run: (*parser).callonModuleName1,
				expr: &seqExpr{
					pos: position{line: 358, col: 15, offset: 11528},
					exprs: []interface{}{
						&oneOrMoreExpr{
							pos: position{line: 358, col: 15, offset: 11528},
							expr: &charClassMatcher{
								pos:        position{line: 358, col: 15, offset: 11528},
								val:        "[A-Z]",
								ranges:     []rune{'A', 'Z'},
								ignoreCase: false,
								inverted:   false,
							},
						},
						&zeroOrMoreExpr{
							pos: position{line: 358, col: 22, offset: 11535},
							expr: &charClassMatcher{
								pos:        position{line: 358, col: 22, offset: 11535},
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
			pos:  position{line: 362, col: 1, offset: 11646},
			expr: &choiceExpr{
				pos: position{line: 362, col: 9, offset: 11654},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 362, col: 9, offset: 11654},
						run: (*parser).callonConst2,
						expr: &seqExpr{
							pos: position{line: 362, col: 9, offset: 11654},
							exprs: []interface{}{
								&zeroOrOneExpr{
									pos: position{line: 362, col: 9, offset: 11654},
									expr: &litMatcher{
										pos:        position{line: 362, col: 9, offset: 11654},
										val:        "-",
										ignoreCase: false,
									},
								},
								&oneOrMoreExpr{
									pos: position{line: 362, col: 14, offset: 11659},
									expr: &charClassMatcher{
										pos:        position{line: 362, col: 14, offset: 11659},
										val:        "[0-9]",
										ranges:     []rune{'0', '9'},
										ignoreCase: false,
										inverted:   false,
									},
								},
								&notExpr{
									pos: position{line: 362, col: 21, offset: 11666},
									expr: &litMatcher{
										pos:        position{line: 362, col: 22, offset: 11667},
										val:        ".",
										ignoreCase: false,
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 369, col: 3, offset: 11843},
						run: (*parser).callonConst10,
						expr: &seqExpr{
							pos: position{line: 369, col: 3, offset: 11843},
							exprs: []interface{}{
								&zeroOrOneExpr{
									pos: position{line: 369, col: 3, offset: 11843},
									expr: &litMatcher{
										pos:        position{line: 369, col: 3, offset: 11843},
										val:        "-",
										ignoreCase: false,
									},
								},
								&oneOrMoreExpr{
									pos: position{line: 369, col: 8, offset: 11848},
									expr: &charClassMatcher{
										pos:        position{line: 369, col: 8, offset: 11848},
										val:        "[0-9]",
										ranges:     []rune{'0', '9'},
										ignoreCase: false,
										inverted:   false,
									},
								},
								&litMatcher{
									pos:        position{line: 369, col: 15, offset: 11855},
									val:        ".",
									ignoreCase: false,
								},
								&oneOrMoreExpr{
									pos: position{line: 369, col: 19, offset: 11859},
									expr: &charClassMatcher{
										pos:        position{line: 369, col: 19, offset: 11859},
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
						pos:        position{line: 376, col: 3, offset: 12049},
						val:        "true",
						ignoreCase: false,
					},
					&actionExpr{
						pos: position{line: 376, col: 12, offset: 12058},
						run: (*parser).callonConst20,
						expr: &litMatcher{
							pos:        position{line: 376, col: 12, offset: 12058},
							val:        "false",
							ignoreCase: false,
						},
					},
					&actionExpr{
						pos: position{line: 382, col: 3, offset: 12259},
						run: (*parser).callonConst22,
						expr: &litMatcher{
							pos:        position{line: 382, col: 3, offset: 12259},
							val:        "()",
							ignoreCase: false,
						},
					},
					&actionExpr{
						pos: position{line: 385, col: 3, offset: 12322},
						run: (*parser).callonConst24,
						expr: &seqExpr{
							pos: position{line: 385, col: 3, offset: 12322},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 385, col: 3, offset: 12322},
									val:        "\"",
									ignoreCase: false,
								},
								&zeroOrMoreExpr{
									pos: position{line: 385, col: 7, offset: 12326},
									expr: &seqExpr{
										pos: position{line: 385, col: 8, offset: 12327},
										exprs: []interface{}{
											&notExpr{
												pos: position{line: 385, col: 8, offset: 12327},
												expr: &ruleRefExpr{
													pos:  position{line: 385, col: 9, offset: 12328},
													name: "EscapedChar",
												},
											},
											&anyMatcher{
												line: 385, col: 21, offset: 12340,
											},
										},
									},
								},
								&litMatcher{
									pos:        position{line: 385, col: 25, offset: 12344},
									val:        "\"",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 392, col: 3, offset: 12528},
						run: (*parser).callonConst33,
						expr: &seqExpr{
							pos: position{line: 392, col: 3, offset: 12528},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 392, col: 3, offset: 12528},
									val:        "'",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 392, col: 7, offset: 12532},
									label: "val",
									expr: &seqExpr{
										pos: position{line: 392, col: 12, offset: 12537},
										exprs: []interface{}{
											&notExpr{
												pos: position{line: 392, col: 12, offset: 12537},
												expr: &ruleRefExpr{
													pos:  position{line: 392, col: 13, offset: 12538},
													name: "EscapedChar",
												},
											},
											&anyMatcher{
												line: 392, col: 25, offset: 12550,
											},
										},
									},
								},
								&litMatcher{
									pos:        position{line: 392, col: 28, offset: 12553},
									val:        "'",
									ignoreCase: false,
								},
							},
						},
					},
					&ruleRefExpr{
						pos:  position{line: 394, col: 5, offset: 12645},
						name: "ArrayLiteral",
					},
				},
			},
		},
		{
			name: "Unused",
			pos:  position{line: 396, col: 1, offset: 12659},
			expr: &actionExpr{
				pos: position{line: 396, col: 10, offset: 12668},
				run: (*parser).callonUnused1,
				expr: &litMatcher{
					pos:        position{line: 396, col: 11, offset: 12669},
					val:        "_",
					ignoreCase: false,
				},
			},
		},
		{
			name: "Reserved",
			pos:  position{line: 400, col: 1, offset: 12770},
			expr: &seqExpr{
				pos: position{line: 400, col: 12, offset: 12781},
				exprs: []interface{}{
					&choiceExpr{
						pos: position{line: 400, col: 13, offset: 12782},
						alternatives: []interface{}{
							&litMatcher{
								pos:        position{line: 400, col: 13, offset: 12782},
								val:        "let",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 400, col: 21, offset: 12790},
								val:        "if",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 400, col: 28, offset: 12797},
								val:        "else",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 400, col: 37, offset: 12806},
								val:        "func",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 400, col: 46, offset: 12815},
								val:        "type",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 400, col: 55, offset: 12824},
								val:        "true",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 400, col: 64, offset: 12833},
								val:        "false",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 400, col: 74, offset: 12843},
								val:        "mutable",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 400, col: 86, offset: 12855},
								val:        "type",
								ignoreCase: false,
							},
							&ruleRefExpr{
								pos:  position{line: 400, col: 95, offset: 12864},
								name: "BaseType",
							},
						},
					},
					&notExpr{
						pos: position{line: 400, col: 105, offset: 12874},
						expr: &oneOrMoreExpr{
							pos: position{line: 400, col: 106, offset: 12875},
							expr: &charClassMatcher{
								pos:        position{line: 400, col: 106, offset: 12875},
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
			pos:  position{line: 402, col: 1, offset: 12883},
			expr: &actionExpr{
				pos: position{line: 402, col: 12, offset: 12894},
				run: (*parser).callonBaseType1,
				expr: &choiceExpr{
					pos: position{line: 402, col: 14, offset: 12896},
					alternatives: []interface{}{
						&litMatcher{
							pos:        position{line: 402, col: 14, offset: 12896},
							val:        "int",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 402, col: 22, offset: 12904},
							val:        "bool",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 402, col: 31, offset: 12913},
							val:        "string",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 402, col: 42, offset: 12924},
							val:        "byte",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 402, col: 51, offset: 12933},
							val:        "rune",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 402, col: 60, offset: 12942},
							val:        "float",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 402, col: 70, offset: 12952},
							val:        "list",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "EscapedChar",
			pos:  position{line: 406, col: 1, offset: 13051},
			expr: &charClassMatcher{
				pos:        position{line: 406, col: 15, offset: 13065},
				val:        "[\\x00-\\x1f\"\\\\]",
				chars:      []rune{'"', '\\'},
				ranges:     []rune{'\x00', '\x1f'},
				ignoreCase: false,
				inverted:   false,
			},
		},
		{
			name: "EscapeSequence",
			pos:  position{line: 408, col: 1, offset: 13081},
			expr: &choiceExpr{
				pos: position{line: 408, col: 18, offset: 13098},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 408, col: 18, offset: 13098},
						name: "SingleCharEscape",
					},
					&ruleRefExpr{
						pos:  position{line: 408, col: 37, offset: 13117},
						name: "UnicodeEscape",
					},
				},
			},
		},
		{
			name: "SingleCharEscape",
			pos:  position{line: 410, col: 1, offset: 13132},
			expr: &charClassMatcher{
				pos:        position{line: 410, col: 20, offset: 13151},
				val:        "[\"\\\\/bfnrt]",
				chars:      []rune{'"', '\\', '/', 'b', 'f', 'n', 'r', 't'},
				ignoreCase: false,
				inverted:   false,
			},
		},
		{
			name: "DecimalDigit",
			pos:  position{line: 412, col: 1, offset: 13164},
			expr: &charClassMatcher{
				pos:        position{line: 412, col: 16, offset: 13179},
				val:        "[0-9]",
				ranges:     []rune{'0', '9'},
				ignoreCase: false,
				inverted:   false,
			},
		},
		{
			name: "NonZeroDecimalDigit",
			pos:  position{line: 414, col: 1, offset: 13186},
			expr: &charClassMatcher{
				pos:        position{line: 414, col: 23, offset: 13208},
				val:        "[1-9]",
				ranges:     []rune{'1', '9'},
				ignoreCase: false,
				inverted:   false,
			},
		},
		{
			name: "HexDigit",
			pos:  position{line: 416, col: 1, offset: 13215},
			expr: &charClassMatcher{
				pos:        position{line: 416, col: 12, offset: 13226},
				val:        "[0-9a-f]i",
				ranges:     []rune{'0', '9', 'a', 'f'},
				ignoreCase: true,
				inverted:   false,
			},
		},
		{
			name:        "__",
			displayName: "\"sigwhitespace\"",
			pos:         position{line: 418, col: 1, offset: 13237},
			expr: &oneOrMoreExpr{
				pos: position{line: 418, col: 22, offset: 13258},
				expr: &charClassMatcher{
					pos:        position{line: 418, col: 22, offset: 13258},
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
			pos:         position{line: 420, col: 1, offset: 13270},
			expr: &zeroOrMoreExpr{
				pos: position{line: 420, col: 18, offset: 13287},
				expr: &charClassMatcher{
					pos:        position{line: 420, col: 18, offset: 13287},
					val:        "[ \\r\\n\\t]",
					chars:      []rune{' ', '\r', '\n', '\t'},
					ignoreCase: false,
					inverted:   false,
				},
			},
		},
		{
			name: "EOF",
			pos:  position{line: 422, col: 1, offset: 13299},
			expr: &notExpr{
				pos: position{line: 422, col: 7, offset: 13305},
				expr: &anyMatcher{
					line: 422, col: 8, offset: 13306,
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
	return AliasType{Name: name}, nil
}

func (p *parser) callonTypeDefn2() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onTypeDefn2(stack["name"], stack["params"], stack["types"])
}

func (c *current) onTypeDefn25(name, params, first, rest interface{}) (interface{}, error) {
	// record type
	return RecordType{Name: name}, nil
}

func (p *parser) callonTypeDefn25() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onTypeDefn25(stack["name"], stack["params"], stack["first"], stack["rest"])
}

func (c *current) onTypeDefn57(name, params, rest interface{}) (interface{}, error) {
	// variant type
	return VariantType{Name: name}, nil
}

func (p *parser) callonTypeDefn57() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onTypeDefn57(stack["name"], stack["params"], stack["rest"])
}

func (c *current) onVariantConstructor1(name, rest interface{}) (interface{}, error) {
	return VariantConstructor{}, nil
}

func (p *parser) callonVariantConstructor1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onVariantConstructor1(stack["name"], stack["rest"])
}

func (c *current) onRecordFieldDefn1(name, t interface{}) (interface{}, error) {
	return RecordField{}, nil
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
