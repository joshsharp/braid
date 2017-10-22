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
						name: "FuncDefn",
					},
					&ruleRefExpr{
						pos:  position{line: 28, col: 32, offset: 664},
						name: "TypeDefn",
					},
				},
			},
		},
		{
			name: "Statement",
			pos:  position{line: 30, col: 1, offset: 674},
			expr: &choiceExpr{
				pos: position{line: 30, col: 13, offset: 686},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 30, col: 13, offset: 686},
						name: "FuncDefn",
					},
					&ruleRefExpr{
						pos:  position{line: 30, col: 24, offset: 697},
						name: "Assignment",
					},
					&ruleRefExpr{
						pos:  position{line: 30, col: 37, offset: 710},
						name: "ExprLine",
					},
				},
			},
		},
		{
			name: "ExprLine",
			pos:  position{line: 32, col: 1, offset: 720},
			expr: &actionExpr{
				pos: position{line: 32, col: 12, offset: 731},
				run: (*parser).callonExprLine1,
				expr: &seqExpr{
					pos: position{line: 32, col: 12, offset: 731},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 32, col: 12, offset: 731},
							label: "e",
							expr: &ruleRefExpr{
								pos:  position{line: 32, col: 14, offset: 733},
								name: "Expr",
							},
						},
						&andExpr{
							pos: position{line: 32, col: 19, offset: 738},
							expr: &litMatcher{
								pos:        position{line: 32, col: 20, offset: 739},
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
			pos:  position{line: 37, col: 1, offset: 768},
			expr: &choiceExpr{
				pos: position{line: 37, col: 12, offset: 779},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 37, col: 12, offset: 779},
						run: (*parser).callonTypeDefn2,
						expr: &seqExpr{
							pos: position{line: 37, col: 12, offset: 779},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 37, col: 12, offset: 779},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 37, col: 14, offset: 781},
									val:        "type",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 37, col: 21, offset: 788},
									name: "__",
								},
								&labeledExpr{
									pos:   position{line: 37, col: 24, offset: 791},
									label: "name",
									expr: &ruleRefExpr{
										pos:  position{line: 37, col: 29, offset: 796},
										name: "ModuleName",
									},
								},
								&labeledExpr{
									pos:   position{line: 37, col: 40, offset: 807},
									label: "params",
									expr: &zeroOrMoreExpr{
										pos: position{line: 37, col: 47, offset: 814},
										expr: &seqExpr{
											pos: position{line: 37, col: 48, offset: 815},
											exprs: []interface{}{
												&ruleRefExpr{
													pos:  position{line: 37, col: 48, offset: 815},
													name: "__",
												},
												&ruleRefExpr{
													pos:  position{line: 37, col: 51, offset: 818},
													name: "TypeParameter",
												},
											},
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 37, col: 67, offset: 834},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 37, col: 69, offset: 836},
									val:        "=",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 37, col: 73, offset: 840},
									label: "types",
									expr: &oneOrMoreExpr{
										pos: position{line: 37, col: 79, offset: 846},
										expr: &seqExpr{
											pos: position{line: 37, col: 80, offset: 847},
											exprs: []interface{}{
												&ruleRefExpr{
													pos:  position{line: 37, col: 80, offset: 847},
													name: "__",
												},
												&ruleRefExpr{
													pos:  position{line: 37, col: 83, offset: 850},
													name: "AnyType",
												},
											},
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 37, col: 93, offset: 860},
									name: "__",
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 56, col: 1, offset: 1356},
						run: (*parser).callonTypeDefn22,
						expr: &seqExpr{
							pos: position{line: 56, col: 1, offset: 1356},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 56, col: 1, offset: 1356},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 56, col: 3, offset: 1358},
									val:        "type",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 56, col: 10, offset: 1365},
									name: "__",
								},
								&labeledExpr{
									pos:   position{line: 56, col: 13, offset: 1368},
									label: "name",
									expr: &ruleRefExpr{
										pos:  position{line: 56, col: 18, offset: 1373},
										name: "ModuleName",
									},
								},
								&labeledExpr{
									pos:   position{line: 56, col: 29, offset: 1384},
									label: "params",
									expr: &zeroOrMoreExpr{
										pos: position{line: 56, col: 36, offset: 1391},
										expr: &seqExpr{
											pos: position{line: 56, col: 37, offset: 1392},
											exprs: []interface{}{
												&ruleRefExpr{
													pos:  position{line: 56, col: 37, offset: 1392},
													name: "__",
												},
												&ruleRefExpr{
													pos:  position{line: 56, col: 40, offset: 1395},
													name: "TypeParameter",
												},
											},
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 56, col: 56, offset: 1411},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 56, col: 58, offset: 1413},
									val:        "=",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 56, col: 62, offset: 1417},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 57, col: 5, offset: 1423},
									val:        "{",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 57, col: 9, offset: 1427},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 57, col: 11, offset: 1429},
									label: "first",
									expr: &ruleRefExpr{
										pos:  position{line: 57, col: 17, offset: 1435},
										name: "RecordFieldDefn",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 57, col: 33, offset: 1451},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 57, col: 35, offset: 1453},
									label: "rest",
									expr: &zeroOrMoreExpr{
										pos: position{line: 57, col: 40, offset: 1458},
										expr: &seqExpr{
											pos: position{line: 57, col: 41, offset: 1459},
											exprs: []interface{}{
												&litMatcher{
													pos:        position{line: 57, col: 41, offset: 1459},
													val:        ",",
													ignoreCase: false,
												},
												&ruleRefExpr{
													pos:  position{line: 57, col: 45, offset: 1463},
													name: "_",
												},
												&ruleRefExpr{
													pos:  position{line: 57, col: 47, offset: 1465},
													name: "RecordFieldDefn",
												},
												&ruleRefExpr{
													pos:  position{line: 57, col: 63, offset: 1481},
													name: "_",
												},
											},
										},
									},
								},
								&zeroOrOneExpr{
									pos: position{line: 57, col: 67, offset: 1485},
									expr: &litMatcher{
										pos:        position{line: 57, col: 67, offset: 1485},
										val:        ",",
										ignoreCase: false,
									},
								},
								&ruleRefExpr{
									pos:  position{line: 57, col: 72, offset: 1490},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 57, col: 74, offset: 1492},
									val:        "}",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 57, col: 78, offset: 1496},
									name: "__",
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 75, col: 1, offset: 1981},
						run: (*parser).callonTypeDefn54,
						expr: &seqExpr{
							pos: position{line: 75, col: 1, offset: 1981},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 75, col: 1, offset: 1981},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 75, col: 3, offset: 1983},
									val:        "type",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 75, col: 10, offset: 1990},
									name: "__",
								},
								&labeledExpr{
									pos:   position{line: 75, col: 13, offset: 1993},
									label: "name",
									expr: &ruleRefExpr{
										pos:  position{line: 75, col: 18, offset: 1998},
										name: "ModuleName",
									},
								},
								&labeledExpr{
									pos:   position{line: 75, col: 29, offset: 2009},
									label: "params",
									expr: &zeroOrMoreExpr{
										pos: position{line: 75, col: 36, offset: 2016},
										expr: &seqExpr{
											pos: position{line: 75, col: 37, offset: 2017},
											exprs: []interface{}{
												&ruleRefExpr{
													pos:  position{line: 75, col: 37, offset: 2017},
													name: "__",
												},
												&ruleRefExpr{
													pos:  position{line: 75, col: 40, offset: 2020},
													name: "TypeParameter",
												},
											},
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 75, col: 56, offset: 2036},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 75, col: 58, offset: 2038},
									val:        "=",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 75, col: 62, offset: 2042},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 75, col: 64, offset: 2044},
									label: "rest",
									expr: &oneOrMoreExpr{
										pos: position{line: 75, col: 69, offset: 2049},
										expr: &ruleRefExpr{
											pos:  position{line: 75, col: 70, offset: 2050},
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
			pos:  position{line: 90, col: 1, offset: 2457},
			expr: &actionExpr{
				pos: position{line: 90, col: 19, offset: 2475},
				run: (*parser).callonRecordFieldDefn1,
				expr: &seqExpr{
					pos: position{line: 90, col: 19, offset: 2475},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 90, col: 19, offset: 2475},
							label: "name",
							expr: &ruleRefExpr{
								pos:  position{line: 90, col: 24, offset: 2480},
								name: "VariableName",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 90, col: 37, offset: 2493},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 90, col: 39, offset: 2495},
							val:        ":",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 90, col: 43, offset: 2499},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 90, col: 45, offset: 2501},
							label: "t",
							expr: &ruleRefExpr{
								pos:  position{line: 90, col: 48, offset: 2504},
								name: "AnyType",
							},
						},
					},
				},
			},
		},
		{
			name: "VariantConstructor",
			pos:  position{line: 94, col: 1, offset: 2598},
			expr: &choiceExpr{
				pos: position{line: 94, col: 22, offset: 2619},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 94, col: 22, offset: 2619},
						run: (*parser).callonVariantConstructor2,
						expr: &seqExpr{
							pos: position{line: 94, col: 22, offset: 2619},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 94, col: 22, offset: 2619},
									val:        "|",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 94, col: 26, offset: 2623},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 94, col: 28, offset: 2625},
									label: "name",
									expr: &ruleRefExpr{
										pos:  position{line: 94, col: 33, offset: 2630},
										name: "ModuleName",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 94, col: 44, offset: 2641},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 94, col: 46, offset: 2643},
									val:        "{",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 94, col: 50, offset: 2647},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 94, col: 52, offset: 2649},
									label: "first",
									expr: &ruleRefExpr{
										pos:  position{line: 94, col: 58, offset: 2655},
										name: "RecordFieldDefn",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 94, col: 74, offset: 2671},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 94, col: 76, offset: 2673},
									label: "rest",
									expr: &zeroOrMoreExpr{
										pos: position{line: 94, col: 81, offset: 2678},
										expr: &seqExpr{
											pos: position{line: 94, col: 82, offset: 2679},
											exprs: []interface{}{
												&litMatcher{
													pos:        position{line: 94, col: 82, offset: 2679},
													val:        ",",
													ignoreCase: false,
												},
												&ruleRefExpr{
													pos:  position{line: 94, col: 86, offset: 2683},
													name: "_",
												},
												&ruleRefExpr{
													pos:  position{line: 94, col: 88, offset: 2685},
													name: "RecordFieldDefn",
												},
												&ruleRefExpr{
													pos:  position{line: 94, col: 104, offset: 2701},
													name: "_",
												},
											},
										},
									},
								},
								&zeroOrOneExpr{
									pos: position{line: 94, col: 108, offset: 2705},
									expr: &litMatcher{
										pos:        position{line: 94, col: 108, offset: 2705},
										val:        ",",
										ignoreCase: false,
									},
								},
								&ruleRefExpr{
									pos:  position{line: 94, col: 113, offset: 2710},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 94, col: 115, offset: 2712},
									val:        "}",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 94, col: 119, offset: 2716},
									name: "__",
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 113, col: 1, offset: 3321},
						run: (*parser).callonVariantConstructor26,
						expr: &seqExpr{
							pos: position{line: 113, col: 1, offset: 3321},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 113, col: 1, offset: 3321},
									val:        "|",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 113, col: 5, offset: 3325},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 113, col: 7, offset: 3327},
									label: "name",
									expr: &ruleRefExpr{
										pos:  position{line: 113, col: 12, offset: 3332},
										name: "ModuleName",
									},
								},
								&labeledExpr{
									pos:   position{line: 113, col: 23, offset: 3343},
									label: "rest",
									expr: &zeroOrMoreExpr{
										pos: position{line: 113, col: 28, offset: 3348},
										expr: &seqExpr{
											pos: position{line: 113, col: 29, offset: 3349},
											exprs: []interface{}{
												&ruleRefExpr{
													pos:  position{line: 113, col: 29, offset: 3349},
													name: "__",
												},
												&ruleRefExpr{
													pos:  position{line: 113, col: 32, offset: 3352},
													name: "AnyType",
												},
											},
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 113, col: 42, offset: 3362},
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
			pos:  position{line: 130, col: 1, offset: 3799},
			expr: &choiceExpr{
				pos: position{line: 130, col: 11, offset: 3809},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 130, col: 11, offset: 3809},
						name: "BaseType",
					},
					&ruleRefExpr{
						pos:  position{line: 130, col: 22, offset: 3820},
						name: "TypeParameter",
					},
				},
			},
		},
		{
			name: "Assignment",
			pos:  position{line: 132, col: 1, offset: 3835},
			expr: &choiceExpr{
				pos: position{line: 132, col: 14, offset: 3848},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 132, col: 14, offset: 3848},
						run: (*parser).callonAssignment2,
						expr: &seqExpr{
							pos: position{line: 132, col: 14, offset: 3848},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 132, col: 14, offset: 3848},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 132, col: 16, offset: 3850},
									val:        "let",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 132, col: 22, offset: 3856},
									name: "__",
								},
								&labeledExpr{
									pos:   position{line: 132, col: 25, offset: 3859},
									label: "i",
									expr: &ruleRefExpr{
										pos:  position{line: 132, col: 27, offset: 3861},
										name: "Assignable",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 132, col: 38, offset: 3872},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 132, col: 40, offset: 3874},
									val:        "=",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 132, col: 44, offset: 3878},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 132, col: 46, offset: 3880},
									label: "expr",
									expr: &ruleRefExpr{
										pos:  position{line: 132, col: 51, offset: 3885},
										name: "Expr",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 132, col: 56, offset: 3890},
									name: "_",
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 138, col: 1, offset: 4009},
						run: (*parser).callonAssignment15,
						expr: &seqExpr{
							pos: position{line: 138, col: 1, offset: 4009},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 138, col: 1, offset: 4009},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 138, col: 3, offset: 4011},
									val:        "let",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 138, col: 9, offset: 4017},
									name: "__",
								},
								&notExpr{
									pos: position{line: 138, col: 12, offset: 4020},
									expr: &ruleRefExpr{
										pos:  position{line: 138, col: 13, offset: 4021},
										name: "Assignable",
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 142, col: 1, offset: 4119},
						run: (*parser).callonAssignment22,
						expr: &seqExpr{
							pos: position{line: 142, col: 1, offset: 4119},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 142, col: 1, offset: 4119},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 142, col: 3, offset: 4121},
									val:        "let",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 142, col: 9, offset: 4127},
									name: "__",
								},
								&labeledExpr{
									pos:   position{line: 142, col: 12, offset: 4130},
									label: "i",
									expr: &ruleRefExpr{
										pos:  position{line: 142, col: 14, offset: 4132},
										name: "Assignable",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 142, col: 25, offset: 4143},
									name: "_",
								},
								&notExpr{
									pos: position{line: 142, col: 27, offset: 4145},
									expr: &litMatcher{
										pos:        position{line: 142, col: 28, offset: 4146},
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
			pos:  position{line: 146, col: 1, offset: 4240},
			expr: &actionExpr{
				pos: position{line: 146, col: 12, offset: 4251},
				run: (*parser).callonFuncDefn1,
				expr: &seqExpr{
					pos: position{line: 146, col: 12, offset: 4251},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 146, col: 12, offset: 4251},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 146, col: 14, offset: 4253},
							val:        "let",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 146, col: 20, offset: 4259},
							name: "__",
						},
						&labeledExpr{
							pos:   position{line: 146, col: 23, offset: 4262},
							label: "i",
							expr: &ruleRefExpr{
								pos:  position{line: 146, col: 25, offset: 4264},
								name: "VariableName",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 146, col: 38, offset: 4277},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 146, col: 40, offset: 4279},
							val:        "=",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 146, col: 44, offset: 4283},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 146, col: 46, offset: 4285},
							val:        "func",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 146, col: 53, offset: 4292},
							name: "__",
						},
						&labeledExpr{
							pos:   position{line: 146, col: 56, offset: 4295},
							label: "ids",
							expr: &zeroOrMoreExpr{
								pos: position{line: 146, col: 60, offset: 4299},
								expr: &seqExpr{
									pos: position{line: 146, col: 61, offset: 4300},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 146, col: 61, offset: 4300},
											name: "VariableName",
										},
										&ruleRefExpr{
											pos:  position{line: 146, col: 74, offset: 4313},
											name: "__",
										},
									},
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 146, col: 79, offset: 4318},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 146, col: 81, offset: 4320},
							val:        "{",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 146, col: 85, offset: 4324},
							name: "__",
						},
						&labeledExpr{
							pos:   position{line: 146, col: 88, offset: 4327},
							label: "statements",
							expr: &oneOrMoreExpr{
								pos: position{line: 146, col: 99, offset: 4338},
								expr: &ruleRefExpr{
									pos:  position{line: 146, col: 100, offset: 4339},
									name: "Statement",
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 146, col: 112, offset: 4351},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 146, col: 114, offset: 4353},
							val:        "}",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 146, col: 118, offset: 4357},
							name: "_",
						},
					},
				},
			},
		},
		{
			name: "Expr",
			pos:  position{line: 169, col: 1, offset: 5013},
			expr: &actionExpr{
				pos: position{line: 169, col: 8, offset: 5020},
				run: (*parser).callonExpr1,
				expr: &labeledExpr{
					pos:   position{line: 169, col: 8, offset: 5020},
					label: "ex",
					expr: &choiceExpr{
						pos: position{line: 169, col: 12, offset: 5024},
						alternatives: []interface{}{
							&ruleRefExpr{
								pos:  position{line: 169, col: 12, offset: 5024},
								name: "IfExpr",
							},
							&ruleRefExpr{
								pos:  position{line: 169, col: 21, offset: 5033},
								name: "BinOp",
							},
							&ruleRefExpr{
								pos:  position{line: 169, col: 29, offset: 5041},
								name: "Call",
							},
						},
					},
				},
			},
		},
		{
			name: "IfExpr",
			pos:  position{line: 175, col: 1, offset: 5150},
			expr: &choiceExpr{
				pos: position{line: 175, col: 10, offset: 5159},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 175, col: 10, offset: 5159},
						run: (*parser).callonIfExpr2,
						expr: &seqExpr{
							pos: position{line: 175, col: 10, offset: 5159},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 175, col: 10, offset: 5159},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 175, col: 12, offset: 5161},
									val:        "if",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 175, col: 17, offset: 5166},
									name: "__",
								},
								&labeledExpr{
									pos:   position{line: 175, col: 20, offset: 5169},
									label: "expr",
									expr: &ruleRefExpr{
										pos:  position{line: 175, col: 25, offset: 5174},
										name: "BinOpBool",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 175, col: 35, offset: 5184},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 175, col: 37, offset: 5186},
									val:        "{",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 175, col: 41, offset: 5190},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 175, col: 43, offset: 5192},
									label: "thens",
									expr: &oneOrMoreExpr{
										pos: position{line: 175, col: 49, offset: 5198},
										expr: &ruleRefExpr{
											pos:  position{line: 175, col: 50, offset: 5199},
											name: "Statement",
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 175, col: 62, offset: 5211},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 175, col: 64, offset: 5213},
									val:        "}",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 175, col: 68, offset: 5217},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 175, col: 70, offset: 5219},
									val:        "else",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 175, col: 77, offset: 5226},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 175, col: 79, offset: 5228},
									label: "elseifs",
									expr: &ruleRefExpr{
										pos:  position{line: 175, col: 87, offset: 5236},
										name: "IfExpr",
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 187, col: 1, offset: 5566},
						run: (*parser).callonIfExpr22,
						expr: &seqExpr{
							pos: position{line: 187, col: 1, offset: 5566},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 187, col: 1, offset: 5566},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 187, col: 3, offset: 5568},
									val:        "if",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 187, col: 8, offset: 5573},
									name: "__",
								},
								&labeledExpr{
									pos:   position{line: 187, col: 11, offset: 5576},
									label: "expr",
									expr: &ruleRefExpr{
										pos:  position{line: 187, col: 16, offset: 5581},
										name: "BinOpBool",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 187, col: 26, offset: 5591},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 187, col: 28, offset: 5593},
									val:        "{",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 187, col: 32, offset: 5597},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 187, col: 34, offset: 5599},
									label: "thens",
									expr: &oneOrMoreExpr{
										pos: position{line: 187, col: 40, offset: 5605},
										expr: &ruleRefExpr{
											pos:  position{line: 187, col: 41, offset: 5606},
											name: "Statement",
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 187, col: 53, offset: 5618},
									name: "__",
								},
								&litMatcher{
									pos:        position{line: 187, col: 56, offset: 5621},
									val:        "}",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 187, col: 60, offset: 5625},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 187, col: 62, offset: 5627},
									val:        "else",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 187, col: 69, offset: 5634},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 187, col: 71, offset: 5636},
									val:        "{",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 187, col: 75, offset: 5640},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 187, col: 77, offset: 5642},
									label: "elses",
									expr: &oneOrMoreExpr{
										pos: position{line: 187, col: 83, offset: 5648},
										expr: &ruleRefExpr{
											pos:  position{line: 187, col: 84, offset: 5649},
											name: "Statement",
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 187, col: 96, offset: 5661},
									name: "__",
								},
								&litMatcher{
									pos:        position{line: 187, col: 99, offset: 5664},
									val:        "}",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 206, col: 1, offset: 6167},
						run: (*parser).callonIfExpr47,
						expr: &seqExpr{
							pos: position{line: 206, col: 1, offset: 6167},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 206, col: 1, offset: 6167},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 206, col: 3, offset: 6169},
									val:        "if",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 206, col: 8, offset: 6174},
									name: "__",
								},
								&labeledExpr{
									pos:   position{line: 206, col: 11, offset: 6177},
									label: "expr",
									expr: &ruleRefExpr{
										pos:  position{line: 206, col: 16, offset: 6182},
										name: "BinOpBool",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 206, col: 26, offset: 6192},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 206, col: 28, offset: 6194},
									val:        "{",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 206, col: 32, offset: 6198},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 206, col: 34, offset: 6200},
									label: "thens",
									expr: &oneOrMoreExpr{
										pos: position{line: 206, col: 40, offset: 6206},
										expr: &ruleRefExpr{
											pos:  position{line: 206, col: 41, offset: 6207},
											name: "Statement",
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 206, col: 53, offset: 6219},
									name: "__",
								},
								&litMatcher{
									pos:        position{line: 206, col: 56, offset: 6222},
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
			pos:  position{line: 218, col: 1, offset: 6520},
			expr: &choiceExpr{
				pos: position{line: 218, col: 8, offset: 6527},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 218, col: 8, offset: 6527},
						run: (*parser).callonCall2,
						expr: &seqExpr{
							pos: position{line: 218, col: 8, offset: 6527},
							exprs: []interface{}{
								&labeledExpr{
									pos:   position{line: 218, col: 8, offset: 6527},
									label: "module",
									expr: &ruleRefExpr{
										pos:  position{line: 218, col: 15, offset: 6534},
										name: "ModuleName",
									},
								},
								&litMatcher{
									pos:        position{line: 218, col: 26, offset: 6545},
									val:        ".",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 218, col: 30, offset: 6549},
									label: "fn",
									expr: &ruleRefExpr{
										pos:  position{line: 218, col: 33, offset: 6552},
										name: "VariableName",
									},
								},
								&labeledExpr{
									pos:   position{line: 218, col: 46, offset: 6565},
									label: "args",
									expr: &ruleRefExpr{
										pos:  position{line: 218, col: 51, offset: 6570},
										name: "Arguments",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 218, col: 61, offset: 6580},
									name: "_",
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 232, col: 1, offset: 6882},
						run: (*parser).callonCall12,
						expr: &seqExpr{
							pos: position{line: 232, col: 1, offset: 6882},
							exprs: []interface{}{
								&labeledExpr{
									pos:   position{line: 232, col: 1, offset: 6882},
									label: "fn",
									expr: &ruleRefExpr{
										pos:  position{line: 232, col: 4, offset: 6885},
										name: "VariableName",
									},
								},
								&labeledExpr{
									pos:   position{line: 232, col: 17, offset: 6898},
									label: "args",
									expr: &ruleRefExpr{
										pos:  position{line: 232, col: 22, offset: 6903},
										name: "Arguments",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 232, col: 32, offset: 6913},
									name: "_",
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 246, col: 1, offset: 7206},
						run: (*parser).callonCall19,
						expr: &seqExpr{
							pos: position{line: 246, col: 1, offset: 7206},
							exprs: []interface{}{
								&labeledExpr{
									pos:   position{line: 246, col: 1, offset: 7206},
									label: "fn",
									expr: &ruleRefExpr{
										pos:  position{line: 246, col: 4, offset: 7209},
										name: "VariableName",
									},
								},
								&litMatcher{
									pos:        position{line: 246, col: 17, offset: 7222},
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
			pos:  position{line: 254, col: 1, offset: 7378},
			expr: &choiceExpr{
				pos: position{line: 254, col: 13, offset: 7390},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 254, col: 13, offset: 7390},
						run: (*parser).callonArguments2,
						expr: &seqExpr{
							pos: position{line: 254, col: 13, offset: 7390},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 254, col: 13, offset: 7390},
									val:        "(",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 254, col: 17, offset: 7394},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 254, col: 19, offset: 7396},
									label: "argument",
									expr: &ruleRefExpr{
										pos:  position{line: 254, col: 28, offset: 7405},
										name: "BinOpParens",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 254, col: 40, offset: 7417},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 254, col: 42, offset: 7419},
									label: "rest",
									expr: &zeroOrMoreExpr{
										pos: position{line: 254, col: 47, offset: 7424},
										expr: &seqExpr{
											pos: position{line: 254, col: 48, offset: 7425},
											exprs: []interface{}{
												&litMatcher{
													pos:        position{line: 254, col: 48, offset: 7425},
													val:        ",",
													ignoreCase: false,
												},
												&ruleRefExpr{
													pos:  position{line: 254, col: 52, offset: 7429},
													name: "_",
												},
												&ruleRefExpr{
													pos:  position{line: 254, col: 54, offset: 7431},
													name: "BinOpParens",
												},
											},
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 254, col: 68, offset: 7445},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 254, col: 70, offset: 7447},
									val:        ")",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 271, col: 1, offset: 7869},
						run: (*parser).callonArguments17,
						expr: &seqExpr{
							pos: position{line: 271, col: 1, offset: 7869},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 271, col: 1, offset: 7869},
									val:        "(",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 271, col: 5, offset: 7873},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 271, col: 7, offset: 7875},
									label: "argument",
									expr: &ruleRefExpr{
										pos:  position{line: 271, col: 16, offset: 7884},
										name: "Expr",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 271, col: 21, offset: 7889},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 271, col: 23, offset: 7891},
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
			pos:  position{line: 276, col: 1, offset: 7996},
			expr: &actionExpr{
				pos: position{line: 276, col: 16, offset: 8011},
				run: (*parser).callonCompoundExpr1,
				expr: &seqExpr{
					pos: position{line: 276, col: 16, offset: 8011},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 276, col: 16, offset: 8011},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 276, col: 18, offset: 8013},
							label: "op",
							expr: &ruleRefExpr{
								pos:  position{line: 276, col: 21, offset: 8016},
								name: "BinOp",
							},
						},
						&labeledExpr{
							pos:   position{line: 276, col: 27, offset: 8022},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 276, col: 32, offset: 8027},
								expr: &seqExpr{
									pos: position{line: 276, col: 33, offset: 8028},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 276, col: 33, offset: 8028},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 276, col: 36, offset: 8031},
											name: "Operator",
										},
										&ruleRefExpr{
											pos:  position{line: 276, col: 45, offset: 8040},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 276, col: 48, offset: 8043},
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
			pos:  position{line: 296, col: 1, offset: 8649},
			expr: &choiceExpr{
				pos: position{line: 296, col: 9, offset: 8657},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 296, col: 9, offset: 8657},
						name: "BinOpBool",
					},
					&ruleRefExpr{
						pos:  position{line: 296, col: 21, offset: 8669},
						name: "BinOpEquality",
					},
					&ruleRefExpr{
						pos:  position{line: 296, col: 37, offset: 8685},
						name: "BinOpLow",
					},
					&ruleRefExpr{
						pos:  position{line: 296, col: 48, offset: 8696},
						name: "BinOpHigh",
					},
					&ruleRefExpr{
						pos:  position{line: 296, col: 60, offset: 8708},
						name: "BinOpParens",
					},
				},
			},
		},
		{
			name: "BinOpBool",
			pos:  position{line: 298, col: 1, offset: 8721},
			expr: &actionExpr{
				pos: position{line: 298, col: 13, offset: 8733},
				run: (*parser).callonBinOpBool1,
				expr: &seqExpr{
					pos: position{line: 298, col: 13, offset: 8733},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 298, col: 13, offset: 8733},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 298, col: 15, offset: 8735},
							label: "first",
							expr: &ruleRefExpr{
								pos:  position{line: 298, col: 21, offset: 8741},
								name: "BinOpEquality",
							},
						},
						&labeledExpr{
							pos:   position{line: 298, col: 35, offset: 8755},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 298, col: 40, offset: 8760},
								expr: &seqExpr{
									pos: position{line: 298, col: 41, offset: 8761},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 298, col: 41, offset: 8761},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 298, col: 44, offset: 8764},
											name: "OperatorBoolean",
										},
										&ruleRefExpr{
											pos:  position{line: 298, col: 60, offset: 8780},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 298, col: 63, offset: 8783},
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
			pos:  position{line: 331, col: 1, offset: 9676},
			expr: &actionExpr{
				pos: position{line: 331, col: 17, offset: 9692},
				run: (*parser).callonBinOpEquality1,
				expr: &seqExpr{
					pos: position{line: 331, col: 17, offset: 9692},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 331, col: 17, offset: 9692},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 331, col: 19, offset: 9694},
							label: "first",
							expr: &ruleRefExpr{
								pos:  position{line: 331, col: 25, offset: 9700},
								name: "BinOpLow",
							},
						},
						&labeledExpr{
							pos:   position{line: 331, col: 34, offset: 9709},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 331, col: 39, offset: 9714},
								expr: &seqExpr{
									pos: position{line: 331, col: 40, offset: 9715},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 331, col: 40, offset: 9715},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 331, col: 43, offset: 9718},
											name: "OperatorEquality",
										},
										&ruleRefExpr{
											pos:  position{line: 331, col: 60, offset: 9735},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 331, col: 63, offset: 9738},
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
			pos:  position{line: 363, col: 1, offset: 10625},
			expr: &actionExpr{
				pos: position{line: 363, col: 12, offset: 10636},
				run: (*parser).callonBinOpLow1,
				expr: &seqExpr{
					pos: position{line: 363, col: 12, offset: 10636},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 363, col: 12, offset: 10636},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 363, col: 14, offset: 10638},
							label: "first",
							expr: &ruleRefExpr{
								pos:  position{line: 363, col: 20, offset: 10644},
								name: "BinOpHigh",
							},
						},
						&labeledExpr{
							pos:   position{line: 363, col: 30, offset: 10654},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 363, col: 35, offset: 10659},
								expr: &seqExpr{
									pos: position{line: 363, col: 36, offset: 10660},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 363, col: 36, offset: 10660},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 363, col: 39, offset: 10663},
											name: "OperatorLow",
										},
										&ruleRefExpr{
											pos:  position{line: 363, col: 51, offset: 10675},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 363, col: 54, offset: 10678},
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
			pos:  position{line: 395, col: 1, offset: 11566},
			expr: &actionExpr{
				pos: position{line: 395, col: 13, offset: 11578},
				run: (*parser).callonBinOpHigh1,
				expr: &seqExpr{
					pos: position{line: 395, col: 13, offset: 11578},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 395, col: 13, offset: 11578},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 395, col: 15, offset: 11580},
							label: "first",
							expr: &ruleRefExpr{
								pos:  position{line: 395, col: 21, offset: 11586},
								name: "BinOpParens",
							},
						},
						&labeledExpr{
							pos:   position{line: 395, col: 33, offset: 11598},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 395, col: 38, offset: 11603},
								expr: &seqExpr{
									pos: position{line: 395, col: 39, offset: 11604},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 395, col: 39, offset: 11604},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 395, col: 42, offset: 11607},
											name: "OperatorHigh",
										},
										&ruleRefExpr{
											pos:  position{line: 395, col: 55, offset: 11620},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 395, col: 58, offset: 11623},
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
			pos:  position{line: 426, col: 1, offset: 12512},
			expr: &choiceExpr{
				pos: position{line: 426, col: 15, offset: 12526},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 426, col: 15, offset: 12526},
						run: (*parser).callonBinOpParens2,
						expr: &seqExpr{
							pos: position{line: 426, col: 15, offset: 12526},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 426, col: 15, offset: 12526},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 426, col: 17, offset: 12528},
									val:        "(",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 426, col: 21, offset: 12532},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 426, col: 23, offset: 12534},
									label: "first",
									expr: &ruleRefExpr{
										pos:  position{line: 426, col: 29, offset: 12540},
										name: "BinOp",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 426, col: 35, offset: 12546},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 426, col: 37, offset: 12548},
									val:        ")",
									ignoreCase: false,
								},
							},
						},
					},
					&ruleRefExpr{
						pos:  position{line: 429, col: 5, offset: 12671},
						name: "Value",
					},
				},
			},
		},
		{
			name: "Operator",
			pos:  position{line: 431, col: 1, offset: 12678},
			expr: &choiceExpr{
				pos: position{line: 431, col: 12, offset: 12689},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 431, col: 12, offset: 12689},
						name: "OperatorBoolean",
					},
					&ruleRefExpr{
						pos:  position{line: 431, col: 30, offset: 12707},
						name: "OperatorEquality",
					},
					&ruleRefExpr{
						pos:  position{line: 431, col: 49, offset: 12726},
						name: "OperatorHigh",
					},
					&ruleRefExpr{
						pos:  position{line: 431, col: 64, offset: 12741},
						name: "OperatorLow",
					},
				},
			},
		},
		{
			name: "OperatorBoolean",
			pos:  position{line: 433, col: 1, offset: 12754},
			expr: &actionExpr{
				pos: position{line: 433, col: 19, offset: 12772},
				run: (*parser).callonOperatorBoolean1,
				expr: &choiceExpr{
					pos: position{line: 433, col: 21, offset: 12774},
					alternatives: []interface{}{
						&litMatcher{
							pos:        position{line: 433, col: 21, offset: 12774},
							val:        "or",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 433, col: 28, offset: 12781},
							val:        "and",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "OperatorEquality",
			pos:  position{line: 437, col: 1, offset: 12863},
			expr: &actionExpr{
				pos: position{line: 437, col: 20, offset: 12882},
				run: (*parser).callonOperatorEquality1,
				expr: &choiceExpr{
					pos: position{line: 437, col: 22, offset: 12884},
					alternatives: []interface{}{
						&litMatcher{
							pos:        position{line: 437, col: 22, offset: 12884},
							val:        ">=",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 437, col: 29, offset: 12891},
							val:        "<=",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 437, col: 36, offset: 12898},
							val:        "<",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 437, col: 42, offset: 12904},
							val:        ">",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 437, col: 48, offset: 12910},
							val:        "===",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 437, col: 56, offset: 12918},
							val:        "==",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "OperatorHigh",
			pos:  position{line: 441, col: 1, offset: 12997},
			expr: &choiceExpr{
				pos: position{line: 441, col: 16, offset: 13012},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 441, col: 16, offset: 13012},
						run: (*parser).callonOperatorHigh2,
						expr: &choiceExpr{
							pos: position{line: 441, col: 18, offset: 13014},
							alternatives: []interface{}{
								&litMatcher{
									pos:        position{line: 441, col: 18, offset: 13014},
									val:        "/",
									ignoreCase: false,
								},
								&litMatcher{
									pos:        position{line: 441, col: 24, offset: 13020},
									val:        "*",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 444, col: 3, offset: 13103},
						run: (*parser).callonOperatorHigh6,
						expr: &litMatcher{
							pos:        position{line: 444, col: 5, offset: 13105},
							val:        "^",
							ignoreCase: false,
						},
					},
					&actionExpr{
						pos: position{line: 447, col: 3, offset: 13185},
						run: (*parser).callonOperatorHigh8,
						expr: &litMatcher{
							pos:        position{line: 447, col: 3, offset: 13185},
							val:        "++",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "OperatorLow",
			pos:  position{line: 451, col: 1, offset: 13266},
			expr: &actionExpr{
				pos: position{line: 451, col: 15, offset: 13280},
				run: (*parser).callonOperatorLow1,
				expr: &choiceExpr{
					pos: position{line: 451, col: 17, offset: 13282},
					alternatives: []interface{}{
						&litMatcher{
							pos:        position{line: 451, col: 17, offset: 13282},
							val:        "+",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 451, col: 23, offset: 13288},
							val:        "-",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "Value",
			pos:  position{line: 455, col: 1, offset: 13370},
			expr: &choiceExpr{
				pos: position{line: 455, col: 9, offset: 13378},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 455, col: 9, offset: 13378},
						name: "Call",
					},
					&ruleRefExpr{
						pos:  position{line: 455, col: 16, offset: 13385},
						name: "VariableName",
					},
					&ruleRefExpr{
						pos:  position{line: 455, col: 31, offset: 13400},
						name: "Const",
					},
				},
			},
		},
		{
			name: "Assignable",
			pos:  position{line: 457, col: 1, offset: 13407},
			expr: &choiceExpr{
				pos: position{line: 457, col: 14, offset: 13420},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 457, col: 14, offset: 13420},
						name: "VariableName",
					},
					&ruleRefExpr{
						pos:  position{line: 457, col: 29, offset: 13435},
						name: "Unused",
					},
				},
			},
		},
		{
			name: "Identifier",
			pos:  position{line: 459, col: 1, offset: 13443},
			expr: &choiceExpr{
				pos: position{line: 459, col: 14, offset: 13456},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 459, col: 14, offset: 13456},
						name: "VariableName",
					},
					&ruleRefExpr{
						pos:  position{line: 459, col: 29, offset: 13471},
						name: "ModuleName",
					},
				},
			},
		},
		{
			name: "ArrayLiteral",
			pos:  position{line: 461, col: 1, offset: 13483},
			expr: &actionExpr{
				pos: position{line: 461, col: 16, offset: 13498},
				run: (*parser).callonArrayLiteral1,
				expr: &seqExpr{
					pos: position{line: 461, col: 16, offset: 13498},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 461, col: 16, offset: 13498},
							val:        "[",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 461, col: 20, offset: 13502},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 461, col: 22, offset: 13504},
							label: "first",
							expr: &ruleRefExpr{
								pos:  position{line: 461, col: 28, offset: 13510},
								name: "Expr",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 461, col: 33, offset: 13515},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 461, col: 35, offset: 13517},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 461, col: 40, offset: 13522},
								expr: &seqExpr{
									pos: position{line: 461, col: 41, offset: 13523},
									exprs: []interface{}{
										&litMatcher{
											pos:        position{line: 461, col: 41, offset: 13523},
											val:        ",",
											ignoreCase: false,
										},
										&ruleRefExpr{
											pos:  position{line: 461, col: 45, offset: 13527},
											name: "_",
										},
										&ruleRefExpr{
											pos:  position{line: 461, col: 47, offset: 13529},
											name: "Expr",
										},
										&ruleRefExpr{
											pos:  position{line: 461, col: 52, offset: 13534},
											name: "_",
										},
									},
								},
							},
						},
						&zeroOrOneExpr{
							pos: position{line: 461, col: 56, offset: 13538},
							expr: &litMatcher{
								pos:        position{line: 461, col: 56, offset: 13538},
								val:        ",",
								ignoreCase: false,
							},
						},
						&ruleRefExpr{
							pos:  position{line: 461, col: 61, offset: 13543},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 461, col: 63, offset: 13545},
							val:        "]",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "VariantInstance",
			pos:  position{line: 477, col: 1, offset: 13990},
			expr: &actionExpr{
				pos: position{line: 477, col: 19, offset: 14008},
				run: (*parser).callonVariantInstance1,
				expr: &seqExpr{
					pos: position{line: 477, col: 19, offset: 14008},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 477, col: 19, offset: 14008},
							label: "name",
							expr: &ruleRefExpr{
								pos:  position{line: 477, col: 24, offset: 14013},
								name: "ModuleName",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 477, col: 35, offset: 14024},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 477, col: 37, offset: 14026},
							label: "args",
							expr: &ruleRefExpr{
								pos:  position{line: 477, col: 42, offset: 14031},
								name: "Arguments",
							},
						},
					},
				},
			},
		},
		{
			name: "RecordInstance",
			pos:  position{line: 490, col: 1, offset: 14301},
			expr: &actionExpr{
				pos: position{line: 490, col: 18, offset: 14318},
				run: (*parser).callonRecordInstance1,
				expr: &seqExpr{
					pos: position{line: 490, col: 18, offset: 14318},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 490, col: 18, offset: 14318},
							label: "name",
							expr: &ruleRefExpr{
								pos:  position{line: 490, col: 23, offset: 14323},
								name: "ModuleName",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 490, col: 34, offset: 14334},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 490, col: 36, offset: 14336},
							val:        "{",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 490, col: 40, offset: 14340},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 490, col: 42, offset: 14342},
							label: "firstName",
							expr: &ruleRefExpr{
								pos:  position{line: 490, col: 52, offset: 14352},
								name: "VariableName",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 490, col: 65, offset: 14365},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 490, col: 67, offset: 14367},
							val:        ":",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 490, col: 71, offset: 14371},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 490, col: 73, offset: 14373},
							label: "firstValue",
							expr: &ruleRefExpr{
								pos:  position{line: 490, col: 84, offset: 14384},
								name: "Expr",
							},
						},
						&labeledExpr{
							pos:   position{line: 490, col: 89, offset: 14389},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 490, col: 94, offset: 14394},
								expr: &seqExpr{
									pos: position{line: 490, col: 95, offset: 14395},
									exprs: []interface{}{
										&litMatcher{
											pos:        position{line: 490, col: 95, offset: 14395},
											val:        ",",
											ignoreCase: false,
										},
										&ruleRefExpr{
											pos:  position{line: 490, col: 99, offset: 14399},
											name: "_",
										},
										&ruleRefExpr{
											pos:  position{line: 490, col: 101, offset: 14401},
											name: "VariableName",
										},
										&ruleRefExpr{
											pos:  position{line: 490, col: 114, offset: 14414},
											name: "_",
										},
										&litMatcher{
											pos:        position{line: 490, col: 116, offset: 14416},
											val:        ":",
											ignoreCase: false,
										},
										&ruleRefExpr{
											pos:  position{line: 490, col: 120, offset: 14420},
											name: "_",
										},
										&ruleRefExpr{
											pos:  position{line: 490, col: 122, offset: 14422},
											name: "Expr",
										},
									},
								},
							},
						},
						&litMatcher{
							pos:        position{line: 490, col: 130, offset: 14430},
							val:        "}",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "TypeParameter",
			pos:  position{line: 510, col: 1, offset: 15014},
			expr: &actionExpr{
				pos: position{line: 510, col: 17, offset: 15030},
				run: (*parser).callonTypeParameter1,
				expr: &seqExpr{
					pos: position{line: 510, col: 17, offset: 15030},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 510, col: 17, offset: 15030},
							val:        "'",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 510, col: 22, offset: 15035},
							name: "VariableName",
						},
					},
				},
			},
		},
		{
			name: "VariableName",
			pos:  position{line: 514, col: 1, offset: 15108},
			expr: &actionExpr{
				pos: position{line: 514, col: 16, offset: 15123},
				run: (*parser).callonVariableName1,
				expr: &seqExpr{
					pos: position{line: 514, col: 16, offset: 15123},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 514, col: 16, offset: 15123},
							expr: &ruleRefExpr{
								pos:  position{line: 514, col: 17, offset: 15124},
								name: "Reserved",
							},
						},
						&seqExpr{
							pos: position{line: 514, col: 27, offset: 15134},
							exprs: []interface{}{
								&oneOrMoreExpr{
									pos: position{line: 514, col: 27, offset: 15134},
									expr: &charClassMatcher{
										pos:        position{line: 514, col: 27, offset: 15134},
										val:        "[a-z]",
										ranges:     []rune{'a', 'z'},
										ignoreCase: false,
										inverted:   false,
									},
								},
								&zeroOrMoreExpr{
									pos: position{line: 514, col: 34, offset: 15141},
									expr: &charClassMatcher{
										pos:        position{line: 514, col: 34, offset: 15141},
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
			pos:  position{line: 518, col: 1, offset: 15216},
			expr: &actionExpr{
				pos: position{line: 518, col: 14, offset: 15229},
				run: (*parser).callonModuleName1,
				expr: &seqExpr{
					pos: position{line: 518, col: 15, offset: 15230},
					exprs: []interface{}{
						&oneOrMoreExpr{
							pos: position{line: 518, col: 15, offset: 15230},
							expr: &charClassMatcher{
								pos:        position{line: 518, col: 15, offset: 15230},
								val:        "[A-Z]",
								ranges:     []rune{'A', 'Z'},
								ignoreCase: false,
								inverted:   false,
							},
						},
						&zeroOrMoreExpr{
							pos: position{line: 518, col: 22, offset: 15237},
							expr: &charClassMatcher{
								pos:        position{line: 518, col: 22, offset: 15237},
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
			pos:  position{line: 522, col: 1, offset: 15312},
			expr: &choiceExpr{
				pos: position{line: 522, col: 9, offset: 15320},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 522, col: 9, offset: 15320},
						run: (*parser).callonConst2,
						expr: &seqExpr{
							pos: position{line: 522, col: 9, offset: 15320},
							exprs: []interface{}{
								&zeroOrOneExpr{
									pos: position{line: 522, col: 9, offset: 15320},
									expr: &litMatcher{
										pos:        position{line: 522, col: 9, offset: 15320},
										val:        "-",
										ignoreCase: false,
									},
								},
								&oneOrMoreExpr{
									pos: position{line: 522, col: 14, offset: 15325},
									expr: &charClassMatcher{
										pos:        position{line: 522, col: 14, offset: 15325},
										val:        "[0-9]",
										ranges:     []rune{'0', '9'},
										ignoreCase: false,
										inverted:   false,
									},
								},
								&notExpr{
									pos: position{line: 522, col: 21, offset: 15332},
									expr: &litMatcher{
										pos:        position{line: 522, col: 22, offset: 15333},
										val:        ".",
										ignoreCase: false,
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 529, col: 3, offset: 15508},
						run: (*parser).callonConst10,
						expr: &seqExpr{
							pos: position{line: 529, col: 3, offset: 15508},
							exprs: []interface{}{
								&zeroOrOneExpr{
									pos: position{line: 529, col: 3, offset: 15508},
									expr: &litMatcher{
										pos:        position{line: 529, col: 3, offset: 15508},
										val:        "-",
										ignoreCase: false,
									},
								},
								&oneOrMoreExpr{
									pos: position{line: 529, col: 8, offset: 15513},
									expr: &charClassMatcher{
										pos:        position{line: 529, col: 8, offset: 15513},
										val:        "[0-9]",
										ranges:     []rune{'0', '9'},
										ignoreCase: false,
										inverted:   false,
									},
								},
								&litMatcher{
									pos:        position{line: 529, col: 15, offset: 15520},
									val:        ".",
									ignoreCase: false,
								},
								&oneOrMoreExpr{
									pos: position{line: 529, col: 19, offset: 15524},
									expr: &charClassMatcher{
										pos:        position{line: 529, col: 19, offset: 15524},
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
						pos: position{line: 536, col: 3, offset: 15713},
						run: (*parser).callonConst19,
						expr: &litMatcher{
							pos:        position{line: 536, col: 3, offset: 15713},
							val:        "true",
							ignoreCase: false,
						},
					},
					&actionExpr{
						pos: position{line: 540, col: 3, offset: 15798},
						run: (*parser).callonConst21,
						expr: &litMatcher{
							pos:        position{line: 540, col: 3, offset: 15798},
							val:        "false",
							ignoreCase: false,
						},
					},
					&ruleRefExpr{
						pos:  position{line: 543, col: 3, offset: 15884},
						name: "Unit",
					},
					&actionExpr{
						pos: position{line: 544, col: 3, offset: 15891},
						run: (*parser).callonConst24,
						expr: &seqExpr{
							pos: position{line: 544, col: 3, offset: 15891},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 544, col: 3, offset: 15891},
									val:        "\"",
									ignoreCase: false,
								},
								&zeroOrMoreExpr{
									pos: position{line: 544, col: 7, offset: 15895},
									expr: &seqExpr{
										pos: position{line: 544, col: 8, offset: 15896},
										exprs: []interface{}{
											&notExpr{
												pos: position{line: 544, col: 8, offset: 15896},
												expr: &ruleRefExpr{
													pos:  position{line: 544, col: 9, offset: 15897},
													name: "EscapedChar",
												},
											},
											&anyMatcher{
												line: 544, col: 21, offset: 15909,
											},
										},
									},
								},
								&litMatcher{
									pos:        position{line: 544, col: 25, offset: 15913},
									val:        "\"",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 551, col: 3, offset: 16097},
						run: (*parser).callonConst33,
						expr: &seqExpr{
							pos: position{line: 551, col: 3, offset: 16097},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 551, col: 3, offset: 16097},
									val:        "'",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 551, col: 7, offset: 16101},
									label: "val",
									expr: &seqExpr{
										pos: position{line: 551, col: 12, offset: 16106},
										exprs: []interface{}{
											&notExpr{
												pos: position{line: 551, col: 12, offset: 16106},
												expr: &ruleRefExpr{
													pos:  position{line: 551, col: 13, offset: 16107},
													name: "EscapedChar",
												},
											},
											&anyMatcher{
												line: 551, col: 25, offset: 16119,
											},
										},
									},
								},
								&litMatcher{
									pos:        position{line: 551, col: 28, offset: 16122},
									val:        "'",
									ignoreCase: false,
								},
							},
						},
					},
					&ruleRefExpr{
						pos:  position{line: 553, col: 5, offset: 16214},
						name: "ArrayLiteral",
					},
					&ruleRefExpr{
						pos:  position{line: 553, col: 20, offset: 16229},
						name: "RecordInstance",
					},
					&ruleRefExpr{
						pos:  position{line: 553, col: 37, offset: 16246},
						name: "VariantInstance",
					},
				},
			},
		},
		{
			name: "Unit",
			pos:  position{line: 555, col: 1, offset: 16263},
			expr: &actionExpr{
				pos: position{line: 555, col: 8, offset: 16270},
				run: (*parser).callonUnit1,
				expr: &litMatcher{
					pos:        position{line: 555, col: 8, offset: 16270},
					val:        "()",
					ignoreCase: false,
				},
			},
		},
		{
			name: "Unused",
			pos:  position{line: 559, col: 1, offset: 16333},
			expr: &actionExpr{
				pos: position{line: 559, col: 10, offset: 16342},
				run: (*parser).callonUnused1,
				expr: &litMatcher{
					pos:        position{line: 559, col: 11, offset: 16343},
					val:        "_",
					ignoreCase: false,
				},
			},
		},
		{
			name: "Reserved",
			pos:  position{line: 563, col: 1, offset: 16398},
			expr: &seqExpr{
				pos: position{line: 563, col: 12, offset: 16409},
				exprs: []interface{}{
					&choiceExpr{
						pos: position{line: 563, col: 13, offset: 16410},
						alternatives: []interface{}{
							&litMatcher{
								pos:        position{line: 563, col: 13, offset: 16410},
								val:        "let",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 563, col: 21, offset: 16418},
								val:        "if",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 563, col: 28, offset: 16425},
								val:        "else",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 563, col: 37, offset: 16434},
								val:        "func",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 563, col: 46, offset: 16443},
								val:        "type",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 563, col: 55, offset: 16452},
								val:        "true",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 563, col: 64, offset: 16461},
								val:        "false",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 563, col: 74, offset: 16471},
								val:        "mutable",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 563, col: 86, offset: 16483},
								val:        "type",
								ignoreCase: false,
							},
							&ruleRefExpr{
								pos:  position{line: 563, col: 95, offset: 16492},
								name: "BaseType",
							},
						},
					},
					&notExpr{
						pos: position{line: 563, col: 105, offset: 16502},
						expr: &oneOrMoreExpr{
							pos: position{line: 563, col: 106, offset: 16503},
							expr: &charClassMatcher{
								pos:        position{line: 563, col: 106, offset: 16503},
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
			pos:  position{line: 565, col: 1, offset: 16511},
			expr: &actionExpr{
				pos: position{line: 565, col: 12, offset: 16522},
				run: (*parser).callonBaseType1,
				expr: &choiceExpr{
					pos: position{line: 565, col: 14, offset: 16524},
					alternatives: []interface{}{
						&litMatcher{
							pos:        position{line: 565, col: 14, offset: 16524},
							val:        "int",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 565, col: 22, offset: 16532},
							val:        "bool",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 565, col: 31, offset: 16541},
							val:        "string",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 565, col: 42, offset: 16552},
							val:        "byte",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 565, col: 51, offset: 16561},
							val:        "rune",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 565, col: 60, offset: 16570},
							val:        "float",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 565, col: 70, offset: 16580},
							val:        "list",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "EscapedChar",
			pos:  position{line: 569, col: 1, offset: 16679},
			expr: &charClassMatcher{
				pos:        position{line: 569, col: 15, offset: 16693},
				val:        "[\\x00-\\x1f\"\\\\]",
				chars:      []rune{'"', '\\'},
				ranges:     []rune{'\x00', '\x1f'},
				ignoreCase: false,
				inverted:   false,
			},
		},
		{
			name: "EscapeSequence",
			pos:  position{line: 571, col: 1, offset: 16709},
			expr: &choiceExpr{
				pos: position{line: 571, col: 18, offset: 16726},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 571, col: 18, offset: 16726},
						name: "SingleCharEscape",
					},
					&ruleRefExpr{
						pos:  position{line: 571, col: 37, offset: 16745},
						name: "UnicodeEscape",
					},
				},
			},
		},
		{
			name: "SingleCharEscape",
			pos:  position{line: 573, col: 1, offset: 16760},
			expr: &charClassMatcher{
				pos:        position{line: 573, col: 20, offset: 16779},
				val:        "[\"\\\\/bfnrt]",
				chars:      []rune{'"', '\\', '/', 'b', 'f', 'n', 'r', 't'},
				ignoreCase: false,
				inverted:   false,
			},
		},
		{
			name: "DecimalDigit",
			pos:  position{line: 575, col: 1, offset: 16792},
			expr: &charClassMatcher{
				pos:        position{line: 575, col: 16, offset: 16807},
				val:        "[0-9]",
				ranges:     []rune{'0', '9'},
				ignoreCase: false,
				inverted:   false,
			},
		},
		{
			name: "NonZeroDecimalDigit",
			pos:  position{line: 577, col: 1, offset: 16814},
			expr: &charClassMatcher{
				pos:        position{line: 577, col: 23, offset: 16836},
				val:        "[1-9]",
				ranges:     []rune{'1', '9'},
				ignoreCase: false,
				inverted:   false,
			},
		},
		{
			name: "HexDigit",
			pos:  position{line: 579, col: 1, offset: 16843},
			expr: &charClassMatcher{
				pos:        position{line: 579, col: 12, offset: 16854},
				val:        "[0-9a-f]i",
				ranges:     []rune{'0', '9', 'a', 'f'},
				ignoreCase: true,
				inverted:   false,
			},
		},
		{
			name:        "__",
			displayName: "\"sigwhitespace\"",
			pos:         position{line: 581, col: 1, offset: 16865},
			expr: &choiceExpr{
				pos: position{line: 581, col: 22, offset: 16886},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 581, col: 22, offset: 16886},
						name: "Comment",
					},
					&oneOrMoreExpr{
						pos: position{line: 581, col: 32, offset: 16896},
						expr: &charClassMatcher{
							pos:        position{line: 581, col: 32, offset: 16896},
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
			pos:         position{line: 583, col: 1, offset: 16908},
			expr: &choiceExpr{
				pos: position{line: 583, col: 18, offset: 16925},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 583, col: 18, offset: 16925},
						name: "Comment",
					},
					&zeroOrMoreExpr{
						pos: position{line: 583, col: 28, offset: 16935},
						expr: &charClassMatcher{
							pos:        position{line: 583, col: 28, offset: 16935},
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
			name: "Comment",
			pos:  position{line: 585, col: 1, offset: 16947},
			expr: &seqExpr{
				pos: position{line: 585, col: 11, offset: 16957},
				exprs: []interface{}{
					&zeroOrMoreExpr{
						pos: position{line: 585, col: 11, offset: 16957},
						expr: &charClassMatcher{
							pos:        position{line: 585, col: 11, offset: 16957},
							val:        "[ \\r\\n\\t]",
							chars:      []rune{' ', '\r', '\n', '\t'},
							ignoreCase: false,
							inverted:   false,
						},
					},
					&litMatcher{
						pos:        position{line: 585, col: 22, offset: 16968},
						val:        "#",
						ignoreCase: false,
					},
					&zeroOrMoreExpr{
						pos: position{line: 585, col: 26, offset: 16972},
						expr: &seqExpr{
							pos: position{line: 585, col: 27, offset: 16973},
							exprs: []interface{}{
								&notExpr{
									pos: position{line: 585, col: 27, offset: 16973},
									expr: &charClassMatcher{
										pos:        position{line: 585, col: 28, offset: 16974},
										val:        "[\\n]",
										chars:      []rune{'\n'},
										ignoreCase: false,
										inverted:   false,
									},
								},
								&anyMatcher{
									line: 585, col: 33, offset: 16979,
								},
							},
						},
					},
					&andExpr{
						pos: position{line: 585, col: 37, offset: 16983},
						expr: &litMatcher{
							pos:        position{line: 585, col: 38, offset: 16984},
							val:        "\n",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "EOF",
			pos:  position{line: 587, col: 1, offset: 16990},
			expr: &notExpr{
				pos: position{line: 587, col: 7, offset: 16996},
				expr: &anyMatcher{
					line: 587, col: 8, offset: 16997,
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
