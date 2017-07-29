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
			pos:  position{line: 28, col: 1, offset: 634},
			expr: &choiceExpr{
				pos: position{line: 28, col: 21, offset: 654},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 28, col: 21, offset: 654},
						name: "Comment",
					},
					&ruleRefExpr{
						pos:  position{line: 28, col: 31, offset: 664},
						name: "FuncDefn",
					},
					&ruleRefExpr{
						pos:  position{line: 28, col: 42, offset: 675},
						name: "TypeDefn",
					},
				},
			},
		},
		{
			name: "Statement",
			pos:  position{line: 30, col: 1, offset: 685},
			expr: &choiceExpr{
				pos: position{line: 30, col: 13, offset: 697},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 30, col: 13, offset: 697},
						name: "Comment",
					},
					&ruleRefExpr{
						pos:  position{line: 30, col: 23, offset: 707},
						name: "FuncDefn",
					},
					&ruleRefExpr{
						pos:  position{line: 30, col: 34, offset: 718},
						name: "Assignment",
					},
					&ruleRefExpr{
						pos:  position{line: 30, col: 47, offset: 731},
						name: "Call",
					},
					&ruleRefExpr{
						pos:  position{line: 30, col: 54, offset: 738},
						name: "ExprLine",
					},
				},
			},
		},
		{
			name: "ExprLine",
			pos:  position{line: 32, col: 1, offset: 748},
			expr: &actionExpr{
				pos: position{line: 32, col: 12, offset: 759},
				run: (*parser).callonExprLine1,
				expr: &seqExpr{
					pos: position{line: 32, col: 12, offset: 759},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 32, col: 12, offset: 759},
							label: "e",
							expr: &ruleRefExpr{
								pos:  position{line: 32, col: 14, offset: 761},
								name: "Expr",
							},
						},
						&andExpr{
							pos: position{line: 32, col: 19, offset: 766},
							expr: &litMatcher{
								pos:        position{line: 32, col: 20, offset: 767},
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
			pos:  position{line: 36, col: 1, offset: 795},
			expr: &actionExpr{
				pos: position{line: 36, col: 11, offset: 805},
				run: (*parser).callonComment1,
				expr: &seqExpr{
					pos: position{line: 36, col: 11, offset: 805},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 36, col: 11, offset: 805},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 36, col: 13, offset: 807},
							val:        "#",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 36, col: 17, offset: 811},
							label: "comment",
							expr: &zeroOrMoreExpr{
								pos: position{line: 36, col: 25, offset: 819},
								expr: &seqExpr{
									pos: position{line: 36, col: 26, offset: 820},
									exprs: []interface{}{
										&notExpr{
											pos: position{line: 36, col: 26, offset: 820},
											expr: &ruleRefExpr{
												pos:  position{line: 36, col: 27, offset: 821},
												name: "EscapedChar",
											},
										},
										&anyMatcher{
											line: 36, col: 39, offset: 833,
										},
									},
								},
							},
						},
						&andExpr{
							pos: position{line: 36, col: 43, offset: 837},
							expr: &litMatcher{
								pos:        position{line: 36, col: 44, offset: 838},
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
			pos:  position{line: 41, col: 1, offset: 988},
			expr: &choiceExpr{
				pos: position{line: 41, col: 12, offset: 999},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 41, col: 12, offset: 999},
						run: (*parser).callonTypeDefn2,
						expr: &seqExpr{
							pos: position{line: 41, col: 12, offset: 999},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 41, col: 12, offset: 999},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 41, col: 14, offset: 1001},
									val:        "type",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 41, col: 21, offset: 1008},
									name: "__",
								},
								&labeledExpr{
									pos:   position{line: 41, col: 24, offset: 1011},
									label: "name",
									expr: &ruleRefExpr{
										pos:  position{line: 41, col: 29, offset: 1016},
										name: "ModuleName",
									},
								},
								&labeledExpr{
									pos:   position{line: 41, col: 40, offset: 1027},
									label: "params",
									expr: &zeroOrMoreExpr{
										pos: position{line: 41, col: 47, offset: 1034},
										expr: &seqExpr{
											pos: position{line: 41, col: 48, offset: 1035},
											exprs: []interface{}{
												&ruleRefExpr{
													pos:  position{line: 41, col: 48, offset: 1035},
													name: "__",
												},
												&ruleRefExpr{
													pos:  position{line: 41, col: 51, offset: 1038},
													name: "TypeParameter",
												},
											},
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 41, col: 67, offset: 1054},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 41, col: 69, offset: 1056},
									val:        "=",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 41, col: 73, offset: 1060},
									label: "types",
									expr: &oneOrMoreExpr{
										pos: position{line: 41, col: 79, offset: 1066},
										expr: &seqExpr{
											pos: position{line: 41, col: 80, offset: 1067},
											exprs: []interface{}{
												&ruleRefExpr{
													pos:  position{line: 41, col: 80, offset: 1067},
													name: "__",
												},
												&ruleRefExpr{
													pos:  position{line: 41, col: 83, offset: 1070},
													name: "AnyType",
												},
											},
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 41, col: 93, offset: 1080},
									name: "__",
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 60, col: 1, offset: 1574},
						run: (*parser).callonTypeDefn22,
						expr: &seqExpr{
							pos: position{line: 60, col: 1, offset: 1574},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 60, col: 1, offset: 1574},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 60, col: 3, offset: 1576},
									val:        "type",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 60, col: 10, offset: 1583},
									name: "__",
								},
								&labeledExpr{
									pos:   position{line: 60, col: 13, offset: 1586},
									label: "name",
									expr: &ruleRefExpr{
										pos:  position{line: 60, col: 18, offset: 1591},
										name: "ModuleName",
									},
								},
								&labeledExpr{
									pos:   position{line: 60, col: 29, offset: 1602},
									label: "params",
									expr: &zeroOrMoreExpr{
										pos: position{line: 60, col: 36, offset: 1609},
										expr: &seqExpr{
											pos: position{line: 60, col: 37, offset: 1610},
											exprs: []interface{}{
												&ruleRefExpr{
													pos:  position{line: 60, col: 37, offset: 1610},
													name: "__",
												},
												&ruleRefExpr{
													pos:  position{line: 60, col: 40, offset: 1613},
													name: "TypeParameter",
												},
											},
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 60, col: 56, offset: 1629},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 60, col: 58, offset: 1631},
									val:        "=",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 60, col: 62, offset: 1635},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 61, col: 5, offset: 1641},
									val:        "{",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 61, col: 9, offset: 1645},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 61, col: 11, offset: 1647},
									label: "first",
									expr: &ruleRefExpr{
										pos:  position{line: 61, col: 17, offset: 1653},
										name: "RecordFieldDefn",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 61, col: 33, offset: 1669},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 61, col: 35, offset: 1671},
									label: "rest",
									expr: &zeroOrMoreExpr{
										pos: position{line: 61, col: 40, offset: 1676},
										expr: &seqExpr{
											pos: position{line: 61, col: 41, offset: 1677},
											exprs: []interface{}{
												&litMatcher{
													pos:        position{line: 61, col: 41, offset: 1677},
													val:        ",",
													ignoreCase: false,
												},
												&ruleRefExpr{
													pos:  position{line: 61, col: 45, offset: 1681},
													name: "_",
												},
												&ruleRefExpr{
													pos:  position{line: 61, col: 47, offset: 1683},
													name: "RecordFieldDefn",
												},
												&ruleRefExpr{
													pos:  position{line: 61, col: 63, offset: 1699},
													name: "_",
												},
											},
										},
									},
								},
								&zeroOrOneExpr{
									pos: position{line: 61, col: 67, offset: 1703},
									expr: &litMatcher{
										pos:        position{line: 61, col: 67, offset: 1703},
										val:        ",",
										ignoreCase: false,
									},
								},
								&ruleRefExpr{
									pos:  position{line: 61, col: 72, offset: 1708},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 61, col: 74, offset: 1710},
									val:        "}",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 61, col: 78, offset: 1714},
									name: "__",
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 79, col: 1, offset: 2197},
						run: (*parser).callonTypeDefn54,
						expr: &seqExpr{
							pos: position{line: 79, col: 1, offset: 2197},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 79, col: 1, offset: 2197},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 79, col: 3, offset: 2199},
									val:        "type",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 79, col: 10, offset: 2206},
									name: "__",
								},
								&labeledExpr{
									pos:   position{line: 79, col: 13, offset: 2209},
									label: "name",
									expr: &ruleRefExpr{
										pos:  position{line: 79, col: 18, offset: 2214},
										name: "ModuleName",
									},
								},
								&labeledExpr{
									pos:   position{line: 79, col: 29, offset: 2225},
									label: "params",
									expr: &zeroOrMoreExpr{
										pos: position{line: 79, col: 36, offset: 2232},
										expr: &seqExpr{
											pos: position{line: 79, col: 37, offset: 2233},
											exprs: []interface{}{
												&ruleRefExpr{
													pos:  position{line: 79, col: 37, offset: 2233},
													name: "__",
												},
												&ruleRefExpr{
													pos:  position{line: 79, col: 40, offset: 2236},
													name: "TypeParameter",
												},
											},
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 79, col: 56, offset: 2252},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 79, col: 58, offset: 2254},
									val:        "=",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 79, col: 62, offset: 2258},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 79, col: 64, offset: 2260},
									label: "rest",
									expr: &oneOrMoreExpr{
										pos: position{line: 79, col: 69, offset: 2265},
										expr: &ruleRefExpr{
											pos:  position{line: 79, col: 70, offset: 2266},
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
			pos:  position{line: 94, col: 1, offset: 2671},
			expr: &actionExpr{
				pos: position{line: 94, col: 19, offset: 2689},
				run: (*parser).callonRecordFieldDefn1,
				expr: &seqExpr{
					pos: position{line: 94, col: 19, offset: 2689},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 94, col: 19, offset: 2689},
							label: "name",
							expr: &ruleRefExpr{
								pos:  position{line: 94, col: 24, offset: 2694},
								name: "VariableName",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 94, col: 37, offset: 2707},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 94, col: 39, offset: 2709},
							val:        ":",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 94, col: 43, offset: 2713},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 94, col: 45, offset: 2715},
							label: "t",
							expr: &ruleRefExpr{
								pos:  position{line: 94, col: 48, offset: 2718},
								name: "AnyType",
							},
						},
					},
				},
			},
		},
		{
			name: "VariantConstructor",
			pos:  position{line: 98, col: 1, offset: 2810},
			expr: &choiceExpr{
				pos: position{line: 98, col: 22, offset: 2831},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 98, col: 22, offset: 2831},
						run: (*parser).callonVariantConstructor2,
						expr: &seqExpr{
							pos: position{line: 98, col: 22, offset: 2831},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 98, col: 22, offset: 2831},
									val:        "|",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 98, col: 26, offset: 2835},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 98, col: 28, offset: 2837},
									label: "name",
									expr: &ruleRefExpr{
										pos:  position{line: 98, col: 33, offset: 2842},
										name: "ModuleName",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 98, col: 44, offset: 2853},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 98, col: 46, offset: 2855},
									val:        "{",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 98, col: 50, offset: 2859},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 98, col: 52, offset: 2861},
									label: "first",
									expr: &ruleRefExpr{
										pos:  position{line: 98, col: 58, offset: 2867},
										name: "RecordFieldDefn",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 98, col: 74, offset: 2883},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 98, col: 76, offset: 2885},
									label: "rest",
									expr: &zeroOrMoreExpr{
										pos: position{line: 98, col: 81, offset: 2890},
										expr: &seqExpr{
											pos: position{line: 98, col: 82, offset: 2891},
											exprs: []interface{}{
												&litMatcher{
													pos:        position{line: 98, col: 82, offset: 2891},
													val:        ",",
													ignoreCase: false,
												},
												&ruleRefExpr{
													pos:  position{line: 98, col: 86, offset: 2895},
													name: "_",
												},
												&ruleRefExpr{
													pos:  position{line: 98, col: 88, offset: 2897},
													name: "RecordFieldDefn",
												},
												&ruleRefExpr{
													pos:  position{line: 98, col: 104, offset: 2913},
													name: "_",
												},
											},
										},
									},
								},
								&zeroOrOneExpr{
									pos: position{line: 98, col: 108, offset: 2917},
									expr: &litMatcher{
										pos:        position{line: 98, col: 108, offset: 2917},
										val:        ",",
										ignoreCase: false,
									},
								},
								&ruleRefExpr{
									pos:  position{line: 98, col: 113, offset: 2922},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 98, col: 115, offset: 2924},
									val:        "}",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 98, col: 119, offset: 2928},
									name: "__",
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 117, col: 1, offset: 3529},
						run: (*parser).callonVariantConstructor26,
						expr: &seqExpr{
							pos: position{line: 117, col: 1, offset: 3529},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 117, col: 1, offset: 3529},
									val:        "|",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 117, col: 5, offset: 3533},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 117, col: 7, offset: 3535},
									label: "name",
									expr: &ruleRefExpr{
										pos:  position{line: 117, col: 12, offset: 3540},
										name: "ModuleName",
									},
								},
								&labeledExpr{
									pos:   position{line: 117, col: 23, offset: 3551},
									label: "rest",
									expr: &zeroOrMoreExpr{
										pos: position{line: 117, col: 28, offset: 3556},
										expr: &seqExpr{
											pos: position{line: 117, col: 29, offset: 3557},
											exprs: []interface{}{
												&ruleRefExpr{
													pos:  position{line: 117, col: 29, offset: 3557},
													name: "__",
												},
												&ruleRefExpr{
													pos:  position{line: 117, col: 32, offset: 3560},
													name: "AnyType",
												},
											},
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 117, col: 42, offset: 3570},
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
			pos:  position{line: 134, col: 1, offset: 4005},
			expr: &choiceExpr{
				pos: position{line: 134, col: 11, offset: 4015},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 134, col: 11, offset: 4015},
						name: "BaseType",
					},
					&ruleRefExpr{
						pos:  position{line: 134, col: 22, offset: 4026},
						name: "TypeParameter",
					},
				},
			},
		},
		{
			name: "Assignment",
			pos:  position{line: 136, col: 1, offset: 4041},
			expr: &choiceExpr{
				pos: position{line: 136, col: 14, offset: 4054},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 136, col: 14, offset: 4054},
						run: (*parser).callonAssignment2,
						expr: &seqExpr{
							pos: position{line: 136, col: 14, offset: 4054},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 136, col: 14, offset: 4054},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 136, col: 16, offset: 4056},
									val:        "let",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 136, col: 22, offset: 4062},
									name: "__",
								},
								&labeledExpr{
									pos:   position{line: 136, col: 25, offset: 4065},
									label: "i",
									expr: &ruleRefExpr{
										pos:  position{line: 136, col: 27, offset: 4067},
										name: "Assignable",
									},
								},
								&labeledExpr{
									pos:   position{line: 136, col: 38, offset: 4078},
									label: "rest",
									expr: &zeroOrMoreExpr{
										pos: position{line: 136, col: 43, offset: 4083},
										expr: &seqExpr{
											pos: position{line: 136, col: 44, offset: 4084},
											exprs: []interface{}{
												&litMatcher{
													pos:        position{line: 136, col: 44, offset: 4084},
													val:        ",",
													ignoreCase: false,
												},
												&ruleRefExpr{
													pos:  position{line: 136, col: 48, offset: 4088},
													name: "_",
												},
												&ruleRefExpr{
													pos:  position{line: 136, col: 50, offset: 4090},
													name: "Assignable",
												},
											},
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 136, col: 63, offset: 4103},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 136, col: 65, offset: 4105},
									val:        "=",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 136, col: 69, offset: 4109},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 136, col: 71, offset: 4111},
									label: "expr",
									expr: &ruleRefExpr{
										pos:  position{line: 136, col: 76, offset: 4116},
										name: "Expr",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 136, col: 81, offset: 4121},
									name: "_",
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 151, col: 1, offset: 4560},
						run: (*parser).callonAssignment21,
						expr: &seqExpr{
							pos: position{line: 151, col: 1, offset: 4560},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 151, col: 1, offset: 4560},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 151, col: 3, offset: 4562},
									val:        "let",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 151, col: 9, offset: 4568},
									name: "__",
								},
								&notExpr{
									pos: position{line: 151, col: 12, offset: 4571},
									expr: &ruleRefExpr{
										pos:  position{line: 151, col: 13, offset: 4572},
										name: "Assignable",
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 155, col: 1, offset: 4680},
						run: (*parser).callonAssignment28,
						expr: &seqExpr{
							pos: position{line: 155, col: 1, offset: 4680},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 155, col: 1, offset: 4680},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 155, col: 3, offset: 4682},
									val:        "let",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 155, col: 9, offset: 4688},
									name: "__",
								},
								&labeledExpr{
									pos:   position{line: 155, col: 12, offset: 4691},
									label: "i",
									expr: &ruleRefExpr{
										pos:  position{line: 155, col: 14, offset: 4693},
										name: "Assignable",
									},
								},
								&labeledExpr{
									pos:   position{line: 155, col: 25, offset: 4704},
									label: "rest",
									expr: &zeroOrMoreExpr{
										pos: position{line: 155, col: 30, offset: 4709},
										expr: &seqExpr{
											pos: position{line: 155, col: 31, offset: 4710},
											exprs: []interface{}{
												&litMatcher{
													pos:        position{line: 155, col: 31, offset: 4710},
													val:        ",",
													ignoreCase: false,
												},
												&ruleRefExpr{
													pos:  position{line: 155, col: 35, offset: 4714},
													name: "_",
												},
												&ruleRefExpr{
													pos:  position{line: 155, col: 37, offset: 4716},
													name: "Assignable",
												},
											},
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 155, col: 50, offset: 4729},
									name: "_",
								},
								&notExpr{
									pos: position{line: 155, col: 52, offset: 4731},
									expr: &litMatcher{
										pos:        position{line: 155, col: 53, offset: 4732},
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
			pos:  position{line: 159, col: 1, offset: 4826},
			expr: &actionExpr{
				pos: position{line: 159, col: 12, offset: 4837},
				run: (*parser).callonFuncDefn1,
				expr: &seqExpr{
					pos: position{line: 159, col: 12, offset: 4837},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 159, col: 12, offset: 4837},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 159, col: 14, offset: 4839},
							val:        "let",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 159, col: 20, offset: 4845},
							name: "__",
						},
						&labeledExpr{
							pos:   position{line: 159, col: 23, offset: 4848},
							label: "i",
							expr: &ruleRefExpr{
								pos:  position{line: 159, col: 25, offset: 4850},
								name: "VariableName",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 159, col: 38, offset: 4863},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 159, col: 40, offset: 4865},
							val:        "=",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 159, col: 44, offset: 4869},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 159, col: 46, offset: 4871},
							val:        "func",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 159, col: 53, offset: 4878},
							name: "__",
						},
						&labeledExpr{
							pos:   position{line: 159, col: 56, offset: 4881},
							label: "ids",
							expr: &zeroOrMoreExpr{
								pos: position{line: 159, col: 60, offset: 4885},
								expr: &seqExpr{
									pos: position{line: 159, col: 61, offset: 4886},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 159, col: 61, offset: 4886},
											name: "VariableName",
										},
										&ruleRefExpr{
											pos:  position{line: 159, col: 74, offset: 4899},
											name: "__",
										},
									},
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 159, col: 79, offset: 4904},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 159, col: 81, offset: 4906},
							val:        "{",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 159, col: 85, offset: 4910},
							name: "__",
						},
						&labeledExpr{
							pos:   position{line: 159, col: 88, offset: 4913},
							label: "statements",
							expr: &oneOrMoreExpr{
								pos: position{line: 159, col: 99, offset: 4924},
								expr: &ruleRefExpr{
									pos:  position{line: 159, col: 100, offset: 4925},
									name: "Statement",
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 159, col: 112, offset: 4937},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 159, col: 114, offset: 4939},
							val:        "}",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 159, col: 118, offset: 4943},
							name: "_",
						},
					},
				},
			},
		},
		{
			name: "Expr",
			pos:  position{line: 182, col: 1, offset: 5597},
			expr: &actionExpr{
				pos: position{line: 182, col: 8, offset: 5604},
				run: (*parser).callonExpr1,
				expr: &labeledExpr{
					pos:   position{line: 182, col: 8, offset: 5604},
					label: "ex",
					expr: &choiceExpr{
						pos: position{line: 182, col: 12, offset: 5608},
						alternatives: []interface{}{
							&ruleRefExpr{
								pos:  position{line: 182, col: 12, offset: 5608},
								name: "IfExpr",
							},
							&ruleRefExpr{
								pos:  position{line: 182, col: 21, offset: 5617},
								name: "Call",
							},
							&ruleRefExpr{
								pos:  position{line: 182, col: 28, offset: 5624},
								name: "CompoundExpr",
							},
						},
					},
				},
			},
		},
		{
			name: "IfExpr",
			pos:  position{line: 188, col: 1, offset: 5741},
			expr: &choiceExpr{
				pos: position{line: 188, col: 10, offset: 5750},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 188, col: 10, offset: 5750},
						run: (*parser).callonIfExpr2,
						expr: &seqExpr{
							pos: position{line: 188, col: 10, offset: 5750},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 188, col: 10, offset: 5750},
									val:        "if",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 188, col: 15, offset: 5755},
									name: "__",
								},
								&labeledExpr{
									pos:   position{line: 188, col: 18, offset: 5758},
									label: "expr",
									expr: &ruleRefExpr{
										pos:  position{line: 188, col: 23, offset: 5763},
										name: "BinOpBool",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 188, col: 33, offset: 5773},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 188, col: 35, offset: 5775},
									val:        "{",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 188, col: 39, offset: 5779},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 188, col: 41, offset: 5781},
									label: "thens",
									expr: &oneOrMoreExpr{
										pos: position{line: 188, col: 47, offset: 5787},
										expr: &ruleRefExpr{
											pos:  position{line: 188, col: 48, offset: 5788},
											name: "Statement",
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 188, col: 60, offset: 5800},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 188, col: 62, offset: 5802},
									val:        "}",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 188, col: 66, offset: 5806},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 188, col: 68, offset: 5808},
									val:        "else",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 188, col: 75, offset: 5815},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 188, col: 77, offset: 5817},
									label: "elseifs",
									expr: &ruleRefExpr{
										pos:  position{line: 188, col: 85, offset: 5825},
										name: "IfExpr",
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 200, col: 1, offset: 6155},
						run: (*parser).callonIfExpr21,
						expr: &seqExpr{
							pos: position{line: 200, col: 1, offset: 6155},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 200, col: 1, offset: 6155},
									val:        "if",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 200, col: 6, offset: 6160},
									name: "__",
								},
								&labeledExpr{
									pos:   position{line: 200, col: 9, offset: 6163},
									label: "expr",
									expr: &ruleRefExpr{
										pos:  position{line: 200, col: 14, offset: 6168},
										name: "BinOpBool",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 200, col: 24, offset: 6178},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 200, col: 26, offset: 6180},
									val:        "{",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 200, col: 30, offset: 6184},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 200, col: 32, offset: 6186},
									label: "thens",
									expr: &oneOrMoreExpr{
										pos: position{line: 200, col: 38, offset: 6192},
										expr: &ruleRefExpr{
											pos:  position{line: 200, col: 39, offset: 6193},
											name: "Statement",
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 200, col: 51, offset: 6205},
									name: "__",
								},
								&litMatcher{
									pos:        position{line: 200, col: 54, offset: 6208},
									val:        "}",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 200, col: 58, offset: 6212},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 200, col: 60, offset: 6214},
									val:        "else",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 200, col: 67, offset: 6221},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 200, col: 69, offset: 6223},
									val:        "{",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 200, col: 73, offset: 6227},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 200, col: 75, offset: 6229},
									label: "elses",
									expr: &oneOrMoreExpr{
										pos: position{line: 200, col: 81, offset: 6235},
										expr: &ruleRefExpr{
											pos:  position{line: 200, col: 82, offset: 6236},
											name: "Statement",
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 200, col: 94, offset: 6248},
									name: "__",
								},
								&litMatcher{
									pos:        position{line: 200, col: 97, offset: 6251},
									val:        "}",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 219, col: 1, offset: 6754},
						run: (*parser).callonIfExpr45,
						expr: &seqExpr{
							pos: position{line: 219, col: 1, offset: 6754},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 219, col: 1, offset: 6754},
									val:        "if",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 219, col: 6, offset: 6759},
									name: "__",
								},
								&labeledExpr{
									pos:   position{line: 219, col: 9, offset: 6762},
									label: "expr",
									expr: &ruleRefExpr{
										pos:  position{line: 219, col: 14, offset: 6767},
										name: "BinOpBool",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 219, col: 24, offset: 6777},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 219, col: 26, offset: 6779},
									val:        "{",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 219, col: 30, offset: 6783},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 219, col: 32, offset: 6785},
									label: "thens",
									expr: &oneOrMoreExpr{
										pos: position{line: 219, col: 38, offset: 6791},
										expr: &ruleRefExpr{
											pos:  position{line: 219, col: 39, offset: 6792},
											name: "Statement",
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 219, col: 51, offset: 6804},
									name: "__",
								},
								&litMatcher{
									pos:        position{line: 219, col: 54, offset: 6807},
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
			pos:  position{line: 231, col: 1, offset: 7105},
			expr: &choiceExpr{
				pos: position{line: 231, col: 8, offset: 7112},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 231, col: 8, offset: 7112},
						run: (*parser).callonCall2,
						expr: &seqExpr{
							pos: position{line: 231, col: 8, offset: 7112},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 231, col: 8, offset: 7112},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 231, col: 10, offset: 7114},
									label: "module",
									expr: &ruleRefExpr{
										pos:  position{line: 231, col: 17, offset: 7121},
										name: "ModuleName",
									},
								},
								&litMatcher{
									pos:        position{line: 231, col: 28, offset: 7132},
									val:        ".",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 231, col: 32, offset: 7136},
									label: "fn",
									expr: &ruleRefExpr{
										pos:  position{line: 231, col: 35, offset: 7139},
										name: "VariableName",
									},
								},
								&labeledExpr{
									pos:   position{line: 231, col: 48, offset: 7152},
									label: "args",
									expr: &ruleRefExpr{
										pos:  position{line: 231, col: 53, offset: 7157},
										name: "Arguments",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 231, col: 63, offset: 7167},
									name: "_",
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 243, col: 1, offset: 7425},
						run: (*parser).callonCall13,
						expr: &seqExpr{
							pos: position{line: 243, col: 1, offset: 7425},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 243, col: 1, offset: 7425},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 243, col: 3, offset: 7427},
									label: "fn",
									expr: &ruleRefExpr{
										pos:  position{line: 243, col: 6, offset: 7430},
										name: "VariableName",
									},
								},
								&labeledExpr{
									pos:   position{line: 243, col: 19, offset: 7443},
									label: "args",
									expr: &ruleRefExpr{
										pos:  position{line: 243, col: 24, offset: 7448},
										name: "Arguments",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 243, col: 34, offset: 7458},
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
			pos:  position{line: 255, col: 1, offset: 7706},
			expr: &choiceExpr{
				pos: position{line: 255, col: 13, offset: 7718},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 255, col: 13, offset: 7718},
						run: (*parser).callonArguments2,
						expr: &seqExpr{
							pos: position{line: 255, col: 13, offset: 7718},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 255, col: 13, offset: 7718},
									val:        "(",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 255, col: 17, offset: 7722},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 255, col: 19, offset: 7724},
									label: "argument",
									expr: &ruleRefExpr{
										pos:  position{line: 255, col: 28, offset: 7733},
										name: "BinOpParens",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 255, col: 40, offset: 7745},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 255, col: 42, offset: 7747},
									label: "rest",
									expr: &zeroOrMoreExpr{
										pos: position{line: 255, col: 47, offset: 7752},
										expr: &seqExpr{
											pos: position{line: 255, col: 48, offset: 7753},
											exprs: []interface{}{
												&litMatcher{
													pos:        position{line: 255, col: 48, offset: 7753},
													val:        ",",
													ignoreCase: false,
												},
												&ruleRefExpr{
													pos:  position{line: 255, col: 52, offset: 7757},
													name: "_",
												},
												&ruleRefExpr{
													pos:  position{line: 255, col: 54, offset: 7759},
													name: "BinOpParens",
												},
											},
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 255, col: 68, offset: 7773},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 255, col: 70, offset: 7775},
									val:        ")",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 272, col: 1, offset: 8218},
						run: (*parser).callonArguments17,
						expr: &seqExpr{
							pos: position{line: 272, col: 1, offset: 8218},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 272, col: 1, offset: 8218},
									val:        "(",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 272, col: 5, offset: 8222},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 272, col: 7, offset: 8224},
									label: "argument",
									expr: &ruleRefExpr{
										pos:  position{line: 272, col: 16, offset: 8233},
										name: "Expr",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 272, col: 21, offset: 8238},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 272, col: 23, offset: 8240},
									val:        ")",
									ignoreCase: false,
								},
							},
						},
					},
					&ruleRefExpr{
						pos:  position{line: 277, col: 1, offset: 8367},
						name: "Unit",
					},
				},
			},
		},
		{
			name: "CompoundExpr",
			pos:  position{line: 279, col: 1, offset: 8373},
			expr: &actionExpr{
				pos: position{line: 279, col: 16, offset: 8388},
				run: (*parser).callonCompoundExpr1,
				expr: &seqExpr{
					pos: position{line: 279, col: 16, offset: 8388},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 279, col: 16, offset: 8388},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 279, col: 18, offset: 8390},
							label: "op",
							expr: &ruleRefExpr{
								pos:  position{line: 279, col: 21, offset: 8393},
								name: "BinOp",
							},
						},
						&labeledExpr{
							pos:   position{line: 279, col: 27, offset: 8399},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 279, col: 32, offset: 8404},
								expr: &seqExpr{
									pos: position{line: 279, col: 33, offset: 8405},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 279, col: 33, offset: 8405},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 279, col: 36, offset: 8408},
											name: "Operator",
										},
										&ruleRefExpr{
											pos:  position{line: 279, col: 45, offset: 8417},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 279, col: 48, offset: 8420},
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
			pos:  position{line: 299, col: 1, offset: 9084},
			expr: &choiceExpr{
				pos: position{line: 299, col: 9, offset: 9092},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 299, col: 9, offset: 9092},
						name: "BinOpBool",
					},
					&ruleRefExpr{
						pos:  position{line: 299, col: 21, offset: 9104},
						name: "BinOpEquality",
					},
					&ruleRefExpr{
						pos:  position{line: 299, col: 37, offset: 9120},
						name: "BinOpLow",
					},
					&ruleRefExpr{
						pos:  position{line: 299, col: 48, offset: 9131},
						name: "BinOpHigh",
					},
					&ruleRefExpr{
						pos:  position{line: 299, col: 60, offset: 9143},
						name: "BinOpParens",
					},
				},
			},
		},
		{
			name: "BinOpBool",
			pos:  position{line: 301, col: 1, offset: 9156},
			expr: &actionExpr{
				pos: position{line: 301, col: 13, offset: 9168},
				run: (*parser).callonBinOpBool1,
				expr: &seqExpr{
					pos: position{line: 301, col: 13, offset: 9168},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 301, col: 13, offset: 9168},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 301, col: 15, offset: 9170},
							label: "first",
							expr: &ruleRefExpr{
								pos:  position{line: 301, col: 21, offset: 9176},
								name: "BinOpEquality",
							},
						},
						&labeledExpr{
							pos:   position{line: 301, col: 35, offset: 9190},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 301, col: 40, offset: 9195},
								expr: &seqExpr{
									pos: position{line: 301, col: 41, offset: 9196},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 301, col: 41, offset: 9196},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 301, col: 44, offset: 9199},
											name: "OperatorBoolean",
										},
										&ruleRefExpr{
											pos:  position{line: 301, col: 60, offset: 9215},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 301, col: 63, offset: 9218},
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
			pos:  position{line: 320, col: 1, offset: 9824},
			expr: &actionExpr{
				pos: position{line: 320, col: 17, offset: 9840},
				run: (*parser).callonBinOpEquality1,
				expr: &seqExpr{
					pos: position{line: 320, col: 17, offset: 9840},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 320, col: 17, offset: 9840},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 320, col: 19, offset: 9842},
							label: "first",
							expr: &ruleRefExpr{
								pos:  position{line: 320, col: 25, offset: 9848},
								name: "BinOpLow",
							},
						},
						&labeledExpr{
							pos:   position{line: 320, col: 34, offset: 9857},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 320, col: 39, offset: 9862},
								expr: &seqExpr{
									pos: position{line: 320, col: 40, offset: 9863},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 320, col: 40, offset: 9863},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 320, col: 43, offset: 9866},
											name: "OperatorEquality",
										},
										&ruleRefExpr{
											pos:  position{line: 320, col: 60, offset: 9883},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 320, col: 63, offset: 9886},
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
			pos:  position{line: 340, col: 1, offset: 10490},
			expr: &actionExpr{
				pos: position{line: 340, col: 12, offset: 10501},
				run: (*parser).callonBinOpLow1,
				expr: &seqExpr{
					pos: position{line: 340, col: 12, offset: 10501},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 340, col: 12, offset: 10501},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 340, col: 14, offset: 10503},
							label: "first",
							expr: &ruleRefExpr{
								pos:  position{line: 340, col: 20, offset: 10509},
								name: "BinOpHigh",
							},
						},
						&labeledExpr{
							pos:   position{line: 340, col: 30, offset: 10519},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 340, col: 35, offset: 10524},
								expr: &seqExpr{
									pos: position{line: 340, col: 36, offset: 10525},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 340, col: 36, offset: 10525},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 340, col: 39, offset: 10528},
											name: "OperatorLow",
										},
										&ruleRefExpr{
											pos:  position{line: 340, col: 51, offset: 10540},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 340, col: 54, offset: 10543},
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
			pos:  position{line: 360, col: 1, offset: 11144},
			expr: &actionExpr{
				pos: position{line: 360, col: 13, offset: 11156},
				run: (*parser).callonBinOpHigh1,
				expr: &seqExpr{
					pos: position{line: 360, col: 13, offset: 11156},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 360, col: 13, offset: 11156},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 360, col: 15, offset: 11158},
							label: "first",
							expr: &ruleRefExpr{
								pos:  position{line: 360, col: 21, offset: 11164},
								name: "BinOpParens",
							},
						},
						&labeledExpr{
							pos:   position{line: 360, col: 33, offset: 11176},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 360, col: 38, offset: 11181},
								expr: &seqExpr{
									pos: position{line: 360, col: 39, offset: 11182},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 360, col: 39, offset: 11182},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 360, col: 42, offset: 11185},
											name: "OperatorHigh",
										},
										&ruleRefExpr{
											pos:  position{line: 360, col: 55, offset: 11198},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 360, col: 58, offset: 11201},
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
			pos:  position{line: 379, col: 1, offset: 11805},
			expr: &choiceExpr{
				pos: position{line: 379, col: 15, offset: 11819},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 379, col: 15, offset: 11819},
						run: (*parser).callonBinOpParens2,
						expr: &seqExpr{
							pos: position{line: 379, col: 15, offset: 11819},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 379, col: 15, offset: 11819},
									val:        "(",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 379, col: 19, offset: 11823},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 379, col: 21, offset: 11825},
									label: "first",
									expr: &ruleRefExpr{
										pos:  position{line: 379, col: 27, offset: 11831},
										name: "BinOp",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 379, col: 33, offset: 11837},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 379, col: 35, offset: 11839},
									val:        ")",
									ignoreCase: false,
								},
							},
						},
					},
					&ruleRefExpr{
						pos:  position{line: 382, col: 5, offset: 11988},
						name: "Value",
					},
				},
			},
		},
		{
			name: "Operator",
			pos:  position{line: 384, col: 1, offset: 11995},
			expr: &choiceExpr{
				pos: position{line: 384, col: 12, offset: 12006},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 384, col: 12, offset: 12006},
						name: "OperatorBoolean",
					},
					&ruleRefExpr{
						pos:  position{line: 384, col: 30, offset: 12024},
						name: "OperatorEquality",
					},
					&ruleRefExpr{
						pos:  position{line: 384, col: 49, offset: 12043},
						name: "OperatorHigh",
					},
					&ruleRefExpr{
						pos:  position{line: 384, col: 64, offset: 12058},
						name: "OperatorLow",
					},
				},
			},
		},
		{
			name: "OperatorBoolean",
			pos:  position{line: 386, col: 1, offset: 12071},
			expr: &actionExpr{
				pos: position{line: 386, col: 19, offset: 12089},
				run: (*parser).callonOperatorBoolean1,
				expr: &choiceExpr{
					pos: position{line: 386, col: 21, offset: 12091},
					alternatives: []interface{}{
						&litMatcher{
							pos:        position{line: 386, col: 21, offset: 12091},
							val:        "not",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 386, col: 29, offset: 12099},
							val:        "or",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 386, col: 36, offset: 12106},
							val:        "and",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "OperatorEquality",
			pos:  position{line: 390, col: 1, offset: 12205},
			expr: &actionExpr{
				pos: position{line: 390, col: 20, offset: 12224},
				run: (*parser).callonOperatorEquality1,
				expr: &choiceExpr{
					pos: position{line: 390, col: 22, offset: 12226},
					alternatives: []interface{}{
						&litMatcher{
							pos:        position{line: 390, col: 22, offset: 12226},
							val:        ">=",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 390, col: 29, offset: 12233},
							val:        "<=",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 390, col: 36, offset: 12240},
							val:        "<",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 390, col: 42, offset: 12246},
							val:        ">",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 390, col: 48, offset: 12252},
							val:        "===",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 390, col: 56, offset: 12260},
							val:        "==",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "OperatorHigh",
			pos:  position{line: 394, col: 1, offset: 12366},
			expr: &choiceExpr{
				pos: position{line: 394, col: 16, offset: 12381},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 394, col: 16, offset: 12381},
						run: (*parser).callonOperatorHigh2,
						expr: &choiceExpr{
							pos: position{line: 394, col: 18, offset: 12383},
							alternatives: []interface{}{
								&litMatcher{
									pos:        position{line: 394, col: 18, offset: 12383},
									val:        "/.",
									ignoreCase: false,
								},
								&litMatcher{
									pos:        position{line: 394, col: 25, offset: 12390},
									val:        "*.",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 397, col: 3, offset: 12496},
						run: (*parser).callonOperatorHigh6,
						expr: &choiceExpr{
							pos: position{line: 397, col: 5, offset: 12498},
							alternatives: []interface{}{
								&litMatcher{
									pos:        position{line: 397, col: 5, offset: 12498},
									val:        "*",
									ignoreCase: false,
								},
								&litMatcher{
									pos:        position{line: 397, col: 11, offset: 12504},
									val:        "/",
									ignoreCase: false,
								},
								&litMatcher{
									pos:        position{line: 397, col: 17, offset: 12510},
									val:        "^",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 400, col: 3, offset: 12613},
						run: (*parser).callonOperatorHigh11,
						expr: &litMatcher{
							pos:        position{line: 400, col: 3, offset: 12613},
							val:        "++",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "OperatorLow",
			pos:  position{line: 404, col: 1, offset: 12717},
			expr: &choiceExpr{
				pos: position{line: 404, col: 15, offset: 12731},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 404, col: 15, offset: 12731},
						run: (*parser).callonOperatorLow2,
						expr: &choiceExpr{
							pos: position{line: 404, col: 17, offset: 12733},
							alternatives: []interface{}{
								&litMatcher{
									pos:        position{line: 404, col: 17, offset: 12733},
									val:        "+.",
									ignoreCase: false,
								},
								&litMatcher{
									pos:        position{line: 404, col: 24, offset: 12740},
									val:        "-.",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 407, col: 3, offset: 12846},
						run: (*parser).callonOperatorLow6,
						expr: &choiceExpr{
							pos: position{line: 407, col: 5, offset: 12848},
							alternatives: []interface{}{
								&litMatcher{
									pos:        position{line: 407, col: 5, offset: 12848},
									val:        "+",
									ignoreCase: false,
								},
								&litMatcher{
									pos:        position{line: 407, col: 11, offset: 12854},
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
			pos:  position{line: 411, col: 1, offset: 12956},
			expr: &choiceExpr{
				pos: position{line: 411, col: 9, offset: 12964},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 411, col: 9, offset: 12964},
						name: "VariableName",
					},
					&ruleRefExpr{
						pos:  position{line: 411, col: 24, offset: 12979},
						name: "Const",
					},
				},
			},
		},
		{
			name: "Assignable",
			pos:  position{line: 413, col: 1, offset: 12986},
			expr: &choiceExpr{
				pos: position{line: 413, col: 14, offset: 12999},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 413, col: 14, offset: 12999},
						name: "VariableName",
					},
					&ruleRefExpr{
						pos:  position{line: 413, col: 29, offset: 13014},
						name: "Unused",
					},
				},
			},
		},
		{
			name: "Identifier",
			pos:  position{line: 415, col: 1, offset: 13022},
			expr: &choiceExpr{
				pos: position{line: 415, col: 14, offset: 13035},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 415, col: 14, offset: 13035},
						name: "VariableName",
					},
					&ruleRefExpr{
						pos:  position{line: 415, col: 29, offset: 13050},
						name: "ModuleName",
					},
				},
			},
		},
		{
			name: "ArrayLiteral",
			pos:  position{line: 417, col: 1, offset: 13062},
			expr: &actionExpr{
				pos: position{line: 417, col: 16, offset: 13077},
				run: (*parser).callonArrayLiteral1,
				expr: &seqExpr{
					pos: position{line: 417, col: 16, offset: 13077},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 417, col: 16, offset: 13077},
							val:        "[",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 417, col: 20, offset: 13081},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 417, col: 22, offset: 13083},
							label: "first",
							expr: &ruleRefExpr{
								pos:  position{line: 417, col: 28, offset: 13089},
								name: "Expr",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 417, col: 33, offset: 13094},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 417, col: 35, offset: 13096},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 417, col: 40, offset: 13101},
								expr: &seqExpr{
									pos: position{line: 417, col: 41, offset: 13102},
									exprs: []interface{}{
										&litMatcher{
											pos:        position{line: 417, col: 41, offset: 13102},
											val:        ",",
											ignoreCase: false,
										},
										&ruleRefExpr{
											pos:  position{line: 417, col: 45, offset: 13106},
											name: "_",
										},
										&ruleRefExpr{
											pos:  position{line: 417, col: 47, offset: 13108},
											name: "Expr",
										},
										&ruleRefExpr{
											pos:  position{line: 417, col: 52, offset: 13113},
											name: "_",
										},
									},
								},
							},
						},
						&zeroOrOneExpr{
							pos: position{line: 417, col: 56, offset: 13117},
							expr: &litMatcher{
								pos:        position{line: 417, col: 56, offset: 13117},
								val:        ",",
								ignoreCase: false,
							},
						},
						&ruleRefExpr{
							pos:  position{line: 417, col: 61, offset: 13122},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 417, col: 63, offset: 13124},
							val:        "]",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "VariantInstance",
			pos:  position{line: 433, col: 1, offset: 13605},
			expr: &actionExpr{
				pos: position{line: 433, col: 19, offset: 13623},
				run: (*parser).callonVariantInstance1,
				expr: &seqExpr{
					pos: position{line: 433, col: 19, offset: 13623},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 433, col: 19, offset: 13623},
							label: "name",
							expr: &ruleRefExpr{
								pos:  position{line: 433, col: 24, offset: 13628},
								name: "ModuleName",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 433, col: 35, offset: 13639},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 433, col: 37, offset: 13641},
							label: "args",
							expr: &ruleRefExpr{
								pos:  position{line: 433, col: 42, offset: 13646},
								name: "Arguments",
							},
						},
					},
				},
			},
		},
		{
			name: "RecordInstance",
			pos:  position{line: 443, col: 1, offset: 13871},
			expr: &actionExpr{
				pos: position{line: 443, col: 18, offset: 13888},
				run: (*parser).callonRecordInstance1,
				expr: &seqExpr{
					pos: position{line: 443, col: 18, offset: 13888},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 443, col: 18, offset: 13888},
							label: "name",
							expr: &ruleRefExpr{
								pos:  position{line: 443, col: 23, offset: 13893},
								name: "ModuleName",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 443, col: 34, offset: 13904},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 443, col: 36, offset: 13906},
							val:        "{",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 443, col: 40, offset: 13910},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 443, col: 42, offset: 13912},
							label: "firstName",
							expr: &ruleRefExpr{
								pos:  position{line: 443, col: 52, offset: 13922},
								name: "VariableName",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 443, col: 65, offset: 13935},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 443, col: 67, offset: 13937},
							val:        ":",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 443, col: 71, offset: 13941},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 443, col: 73, offset: 13943},
							label: "firstValue",
							expr: &ruleRefExpr{
								pos:  position{line: 443, col: 84, offset: 13954},
								name: "Expr",
							},
						},
						&labeledExpr{
							pos:   position{line: 443, col: 89, offset: 13959},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 443, col: 94, offset: 13964},
								expr: &seqExpr{
									pos: position{line: 443, col: 95, offset: 13965},
									exprs: []interface{}{
										&litMatcher{
											pos:        position{line: 443, col: 95, offset: 13965},
											val:        ",",
											ignoreCase: false,
										},
										&ruleRefExpr{
											pos:  position{line: 443, col: 99, offset: 13969},
											name: "_",
										},
										&ruleRefExpr{
											pos:  position{line: 443, col: 101, offset: 13971},
											name: "VariableName",
										},
										&ruleRefExpr{
											pos:  position{line: 443, col: 114, offset: 13984},
											name: "_",
										},
										&litMatcher{
											pos:        position{line: 443, col: 116, offset: 13986},
											val:        ":",
											ignoreCase: false,
										},
										&ruleRefExpr{
											pos:  position{line: 443, col: 120, offset: 13990},
											name: "_",
										},
										&ruleRefExpr{
											pos:  position{line: 443, col: 122, offset: 13992},
											name: "Expr",
										},
									},
								},
							},
						},
						&litMatcher{
							pos:        position{line: 443, col: 130, offset: 14000},
							val:        "}",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "TypeParameter",
			pos:  position{line: 463, col: 1, offset: 14584},
			expr: &actionExpr{
				pos: position{line: 463, col: 17, offset: 14600},
				run: (*parser).callonTypeParameter1,
				expr: &seqExpr{
					pos: position{line: 463, col: 17, offset: 14600},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 463, col: 17, offset: 14600},
							val:        "'",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 463, col: 22, offset: 14605},
							name: "VariableName",
						},
					},
				},
			},
		},
		{
			name: "VariableName",
			pos:  position{line: 467, col: 1, offset: 14713},
			expr: &actionExpr{
				pos: position{line: 467, col: 16, offset: 14728},
				run: (*parser).callonVariableName1,
				expr: &seqExpr{
					pos: position{line: 467, col: 16, offset: 14728},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 467, col: 16, offset: 14728},
							expr: &ruleRefExpr{
								pos:  position{line: 467, col: 17, offset: 14729},
								name: "Reserved",
							},
						},
						&seqExpr{
							pos: position{line: 467, col: 27, offset: 14739},
							exprs: []interface{}{
								&oneOrMoreExpr{
									pos: position{line: 467, col: 27, offset: 14739},
									expr: &charClassMatcher{
										pos:        position{line: 467, col: 27, offset: 14739},
										val:        "[a-z]",
										ranges:     []rune{'a', 'z'},
										ignoreCase: false,
										inverted:   false,
									},
								},
								&zeroOrMoreExpr{
									pos: position{line: 467, col: 34, offset: 14746},
									expr: &charClassMatcher{
										pos:        position{line: 467, col: 34, offset: 14746},
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
			pos:  position{line: 471, col: 1, offset: 14857},
			expr: &actionExpr{
				pos: position{line: 471, col: 14, offset: 14870},
				run: (*parser).callonModuleName1,
				expr: &seqExpr{
					pos: position{line: 471, col: 15, offset: 14871},
					exprs: []interface{}{
						&oneOrMoreExpr{
							pos: position{line: 471, col: 15, offset: 14871},
							expr: &charClassMatcher{
								pos:        position{line: 471, col: 15, offset: 14871},
								val:        "[A-Z]",
								ranges:     []rune{'A', 'Z'},
								ignoreCase: false,
								inverted:   false,
							},
						},
						&zeroOrMoreExpr{
							pos: position{line: 471, col: 22, offset: 14878},
							expr: &charClassMatcher{
								pos:        position{line: 471, col: 22, offset: 14878},
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
			pos:  position{line: 475, col: 1, offset: 14989},
			expr: &choiceExpr{
				pos: position{line: 475, col: 9, offset: 14997},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 475, col: 9, offset: 14997},
						run: (*parser).callonConst2,
						expr: &seqExpr{
							pos: position{line: 475, col: 9, offset: 14997},
							exprs: []interface{}{
								&zeroOrOneExpr{
									pos: position{line: 475, col: 9, offset: 14997},
									expr: &litMatcher{
										pos:        position{line: 475, col: 9, offset: 14997},
										val:        "-",
										ignoreCase: false,
									},
								},
								&oneOrMoreExpr{
									pos: position{line: 475, col: 14, offset: 15002},
									expr: &charClassMatcher{
										pos:        position{line: 475, col: 14, offset: 15002},
										val:        "[0-9]",
										ranges:     []rune{'0', '9'},
										ignoreCase: false,
										inverted:   false,
									},
								},
								&notExpr{
									pos: position{line: 475, col: 21, offset: 15009},
									expr: &litMatcher{
										pos:        position{line: 475, col: 22, offset: 15010},
										val:        ".",
										ignoreCase: false,
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 482, col: 3, offset: 15186},
						run: (*parser).callonConst10,
						expr: &seqExpr{
							pos: position{line: 482, col: 3, offset: 15186},
							exprs: []interface{}{
								&zeroOrOneExpr{
									pos: position{line: 482, col: 3, offset: 15186},
									expr: &litMatcher{
										pos:        position{line: 482, col: 3, offset: 15186},
										val:        "-",
										ignoreCase: false,
									},
								},
								&oneOrMoreExpr{
									pos: position{line: 482, col: 8, offset: 15191},
									expr: &charClassMatcher{
										pos:        position{line: 482, col: 8, offset: 15191},
										val:        "[0-9]",
										ranges:     []rune{'0', '9'},
										ignoreCase: false,
										inverted:   false,
									},
								},
								&litMatcher{
									pos:        position{line: 482, col: 15, offset: 15198},
									val:        ".",
									ignoreCase: false,
								},
								&oneOrMoreExpr{
									pos: position{line: 482, col: 19, offset: 15202},
									expr: &charClassMatcher{
										pos:        position{line: 482, col: 19, offset: 15202},
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
						pos:        position{line: 489, col: 3, offset: 15392},
						val:        "true",
						ignoreCase: false,
					},
					&actionExpr{
						pos: position{line: 489, col: 12, offset: 15401},
						run: (*parser).callonConst20,
						expr: &litMatcher{
							pos:        position{line: 489, col: 12, offset: 15401},
							val:        "false",
							ignoreCase: false,
						},
					},
					&ruleRefExpr{
						pos:  position{line: 495, col: 3, offset: 15602},
						name: "Unit",
					},
					&actionExpr{
						pos: position{line: 496, col: 3, offset: 15609},
						run: (*parser).callonConst23,
						expr: &seqExpr{
							pos: position{line: 496, col: 3, offset: 15609},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 496, col: 3, offset: 15609},
									val:        "\"",
									ignoreCase: false,
								},
								&zeroOrMoreExpr{
									pos: position{line: 496, col: 7, offset: 15613},
									expr: &seqExpr{
										pos: position{line: 496, col: 8, offset: 15614},
										exprs: []interface{}{
											&notExpr{
												pos: position{line: 496, col: 8, offset: 15614},
												expr: &ruleRefExpr{
													pos:  position{line: 496, col: 9, offset: 15615},
													name: "EscapedChar",
												},
											},
											&anyMatcher{
												line: 496, col: 21, offset: 15627,
											},
										},
									},
								},
								&litMatcher{
									pos:        position{line: 496, col: 25, offset: 15631},
									val:        "\"",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 503, col: 3, offset: 15815},
						run: (*parser).callonConst32,
						expr: &seqExpr{
							pos: position{line: 503, col: 3, offset: 15815},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 503, col: 3, offset: 15815},
									val:        "'",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 503, col: 7, offset: 15819},
									label: "val",
									expr: &seqExpr{
										pos: position{line: 503, col: 12, offset: 15824},
										exprs: []interface{}{
											&notExpr{
												pos: position{line: 503, col: 12, offset: 15824},
												expr: &ruleRefExpr{
													pos:  position{line: 503, col: 13, offset: 15825},
													name: "EscapedChar",
												},
											},
											&anyMatcher{
												line: 503, col: 25, offset: 15837,
											},
										},
									},
								},
								&litMatcher{
									pos:        position{line: 503, col: 28, offset: 15840},
									val:        "'",
									ignoreCase: false,
								},
							},
						},
					},
					&ruleRefExpr{
						pos:  position{line: 505, col: 5, offset: 15932},
						name: "ArrayLiteral",
					},
					&ruleRefExpr{
						pos:  position{line: 505, col: 20, offset: 15947},
						name: "RecordInstance",
					},
					&ruleRefExpr{
						pos:  position{line: 505, col: 37, offset: 15964},
						name: "VariantInstance",
					},
				},
			},
		},
		{
			name: "Unit",
			pos:  position{line: 507, col: 1, offset: 15981},
			expr: &actionExpr{
				pos: position{line: 507, col: 8, offset: 15988},
				run: (*parser).callonUnit1,
				expr: &litMatcher{
					pos:        position{line: 507, col: 8, offset: 15988},
					val:        "()",
					ignoreCase: false,
				},
			},
		},
		{
			name: "Unused",
			pos:  position{line: 511, col: 1, offset: 16050},
			expr: &actionExpr{
				pos: position{line: 511, col: 10, offset: 16059},
				run: (*parser).callonUnused1,
				expr: &litMatcher{
					pos:        position{line: 511, col: 11, offset: 16060},
					val:        "_",
					ignoreCase: false,
				},
			},
		},
		{
			name: "Reserved",
			pos:  position{line: 515, col: 1, offset: 16161},
			expr: &seqExpr{
				pos: position{line: 515, col: 12, offset: 16172},
				exprs: []interface{}{
					&choiceExpr{
						pos: position{line: 515, col: 13, offset: 16173},
						alternatives: []interface{}{
							&litMatcher{
								pos:        position{line: 515, col: 13, offset: 16173},
								val:        "let",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 515, col: 21, offset: 16181},
								val:        "if",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 515, col: 28, offset: 16188},
								val:        "else",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 515, col: 37, offset: 16197},
								val:        "func",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 515, col: 46, offset: 16206},
								val:        "type",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 515, col: 55, offset: 16215},
								val:        "true",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 515, col: 64, offset: 16224},
								val:        "false",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 515, col: 74, offset: 16234},
								val:        "mutable",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 515, col: 86, offset: 16246},
								val:        "type",
								ignoreCase: false,
							},
							&ruleRefExpr{
								pos:  position{line: 515, col: 95, offset: 16255},
								name: "BaseType",
							},
						},
					},
					&notExpr{
						pos: position{line: 515, col: 105, offset: 16265},
						expr: &oneOrMoreExpr{
							pos: position{line: 515, col: 106, offset: 16266},
							expr: &charClassMatcher{
								pos:        position{line: 515, col: 106, offset: 16266},
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
			pos:  position{line: 517, col: 1, offset: 16274},
			expr: &actionExpr{
				pos: position{line: 517, col: 12, offset: 16285},
				run: (*parser).callonBaseType1,
				expr: &choiceExpr{
					pos: position{line: 517, col: 14, offset: 16287},
					alternatives: []interface{}{
						&litMatcher{
							pos:        position{line: 517, col: 14, offset: 16287},
							val:        "int",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 517, col: 22, offset: 16295},
							val:        "bool",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 517, col: 31, offset: 16304},
							val:        "string",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 517, col: 42, offset: 16315},
							val:        "byte",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 517, col: 51, offset: 16324},
							val:        "rune",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 517, col: 60, offset: 16333},
							val:        "float",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 517, col: 70, offset: 16343},
							val:        "list",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "EscapedChar",
			pos:  position{line: 521, col: 1, offset: 16442},
			expr: &charClassMatcher{
				pos:        position{line: 521, col: 15, offset: 16456},
				val:        "[\\x00-\\x1f\"\\\\]",
				chars:      []rune{'"', '\\'},
				ranges:     []rune{'\x00', '\x1f'},
				ignoreCase: false,
				inverted:   false,
			},
		},
		{
			name: "EscapeSequence",
			pos:  position{line: 523, col: 1, offset: 16472},
			expr: &choiceExpr{
				pos: position{line: 523, col: 18, offset: 16489},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 523, col: 18, offset: 16489},
						name: "SingleCharEscape",
					},
					&ruleRefExpr{
						pos:  position{line: 523, col: 37, offset: 16508},
						name: "UnicodeEscape",
					},
				},
			},
		},
		{
			name: "SingleCharEscape",
			pos:  position{line: 525, col: 1, offset: 16523},
			expr: &charClassMatcher{
				pos:        position{line: 525, col: 20, offset: 16542},
				val:        "[\"\\\\/bfnrt]",
				chars:      []rune{'"', '\\', '/', 'b', 'f', 'n', 'r', 't'},
				ignoreCase: false,
				inverted:   false,
			},
		},
		{
			name: "DecimalDigit",
			pos:  position{line: 527, col: 1, offset: 16555},
			expr: &charClassMatcher{
				pos:        position{line: 527, col: 16, offset: 16570},
				val:        "[0-9]",
				ranges:     []rune{'0', '9'},
				ignoreCase: false,
				inverted:   false,
			},
		},
		{
			name: "NonZeroDecimalDigit",
			pos:  position{line: 529, col: 1, offset: 16577},
			expr: &charClassMatcher{
				pos:        position{line: 529, col: 23, offset: 16599},
				val:        "[1-9]",
				ranges:     []rune{'1', '9'},
				ignoreCase: false,
				inverted:   false,
			},
		},
		{
			name: "HexDigit",
			pos:  position{line: 531, col: 1, offset: 16606},
			expr: &charClassMatcher{
				pos:        position{line: 531, col: 12, offset: 16617},
				val:        "[0-9a-f]i",
				ranges:     []rune{'0', '9', 'a', 'f'},
				ignoreCase: true,
				inverted:   false,
			},
		},
		{
			name:        "__",
			displayName: "\"sigwhitespace\"",
			pos:         position{line: 533, col: 1, offset: 16628},
			expr: &choiceExpr{
				pos: position{line: 533, col: 22, offset: 16649},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 533, col: 22, offset: 16649},
						name: "Comment",
					},
					&oneOrMoreExpr{
						pos: position{line: 533, col: 32, offset: 16659},
						expr: &charClassMatcher{
							pos:        position{line: 533, col: 32, offset: 16659},
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
			pos:         position{line: 535, col: 1, offset: 16671},
			expr: &zeroOrMoreExpr{
				pos: position{line: 535, col: 18, offset: 16688},
				expr: &charClassMatcher{
					pos:        position{line: 535, col: 18, offset: 16688},
					val:        "[ \\r\\n\\t]",
					chars:      []rune{' ', '\r', '\n', '\t'},
					ignoreCase: false,
					inverted:   false,
				},
			},
		},
		{
			name: "EOF",
			pos:  position{line: 537, col: 1, offset: 16700},
			expr: &notExpr{
				pos: position{line: 537, col: 7, offset: 16706},
				expr: &anyMatcher{
					line: 537, col: 8, offset: 16707,
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

	if args.(BasicAst).ValueType != NIL {
		arguments = args.(BasicAst).Subvalues

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

	if args.(BasicAst).ValueType != NIL {
		arguments = args.(BasicAst).Subvalues

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

	return BasicAst{Type: "Arguments", Subvalues: args, ValueType: CONTAINER}, nil
}

func (p *parser) callonArguments2() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onArguments2(stack["argument"], stack["rest"])
}

func (c *current) onArguments17(argument interface{}) (interface{}, error) {
	args := []Ast{argument.(Ast)}
	return BasicAst{Type: "Arguments", Subvalues: args, ValueType: CONTAINER}, nil
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

func (c *current) onVariantInstance1(name, args interface{}) (interface{}, error) {
	arguments := []Ast{}

	if args.(BasicAst).ValueType != NIL {
		arguments = args.(BasicAst).Subvalues

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
