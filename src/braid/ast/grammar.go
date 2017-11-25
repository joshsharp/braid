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
				},
			},
		},
		{
			name: "Statement",
			pos:  position{line: 31, col: 1, offset: 775},
			expr: &choiceExpr{
				pos: position{line: 31, col: 13, offset: 787},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 31, col: 13, offset: 787},
						name: "FuncDefn",
					},
					&ruleRefExpr{
						pos:  position{line: 31, col: 24, offset: 798},
						name: "Assignment",
					},
					&ruleRefExpr{
						pos:  position{line: 31, col: 37, offset: 811},
						name: "ExprLine",
					},
				},
			},
		},
		{
			name: "ExprLine",
			pos:  position{line: 33, col: 1, offset: 821},
			expr: &actionExpr{
				pos: position{line: 33, col: 12, offset: 832},
				run: (*parser).callonExprLine1,
				expr: &seqExpr{
					pos: position{line: 33, col: 12, offset: 832},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 33, col: 12, offset: 832},
							label: "e",
							expr: &ruleRefExpr{
								pos:  position{line: 33, col: 14, offset: 834},
								name: "Expr",
							},
						},
						&andExpr{
							pos: position{line: 33, col: 19, offset: 839},
							expr: &litMatcher{
								pos:        position{line: 33, col: 20, offset: 840},
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
			pos:  position{line: 41, col: 1, offset: 1019},
			expr: &choiceExpr{
				pos: position{line: 41, col: 12, offset: 1030},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 41, col: 12, offset: 1030},
						run: (*parser).callonTypeDefn2,
						expr: &seqExpr{
							pos: position{line: 41, col: 12, offset: 1030},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 41, col: 12, offset: 1030},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 41, col: 14, offset: 1032},
									val:        "type",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 41, col: 21, offset: 1039},
									name: "__",
								},
								&labeledExpr{
									pos:   position{line: 41, col: 24, offset: 1042},
									label: "name",
									expr: &ruleRefExpr{
										pos:  position{line: 41, col: 29, offset: 1047},
										name: "ModuleName",
									},
								},
								&labeledExpr{
									pos:   position{line: 41, col: 40, offset: 1058},
									label: "params",
									expr: &zeroOrMoreExpr{
										pos: position{line: 41, col: 47, offset: 1065},
										expr: &seqExpr{
											pos: position{line: 41, col: 48, offset: 1066},
											exprs: []interface{}{
												&ruleRefExpr{
													pos:  position{line: 41, col: 48, offset: 1066},
													name: "__",
												},
												&ruleRefExpr{
													pos:  position{line: 41, col: 51, offset: 1069},
													name: "TypeParameter",
												},
											},
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 41, col: 67, offset: 1085},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 41, col: 69, offset: 1087},
									val:        "=",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 41, col: 73, offset: 1091},
									label: "types",
									expr: &oneOrMoreExpr{
										pos: position{line: 41, col: 79, offset: 1097},
										expr: &seqExpr{
											pos: position{line: 41, col: 80, offset: 1098},
											exprs: []interface{}{
												&ruleRefExpr{
													pos:  position{line: 41, col: 80, offset: 1098},
													name: "__",
												},
												&ruleRefExpr{
													pos:  position{line: 41, col: 83, offset: 1101},
													name: "AnyType",
												},
											},
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 41, col: 93, offset: 1111},
									name: "__",
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 60, col: 1, offset: 1607},
						run: (*parser).callonTypeDefn22,
						expr: &seqExpr{
							pos: position{line: 60, col: 1, offset: 1607},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 60, col: 1, offset: 1607},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 60, col: 3, offset: 1609},
									val:        "type",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 60, col: 10, offset: 1616},
									name: "__",
								},
								&labeledExpr{
									pos:   position{line: 60, col: 13, offset: 1619},
									label: "name",
									expr: &ruleRefExpr{
										pos:  position{line: 60, col: 18, offset: 1624},
										name: "ModuleName",
									},
								},
								&labeledExpr{
									pos:   position{line: 60, col: 29, offset: 1635},
									label: "params",
									expr: &zeroOrMoreExpr{
										pos: position{line: 60, col: 36, offset: 1642},
										expr: &seqExpr{
											pos: position{line: 60, col: 37, offset: 1643},
											exprs: []interface{}{
												&ruleRefExpr{
													pos:  position{line: 60, col: 37, offset: 1643},
													name: "__",
												},
												&ruleRefExpr{
													pos:  position{line: 60, col: 40, offset: 1646},
													name: "TypeParameter",
												},
											},
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 60, col: 56, offset: 1662},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 60, col: 58, offset: 1664},
									val:        "=",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 60, col: 62, offset: 1668},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 61, col: 5, offset: 1674},
									val:        "{",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 61, col: 9, offset: 1678},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 61, col: 11, offset: 1680},
									label: "first",
									expr: &ruleRefExpr{
										pos:  position{line: 61, col: 17, offset: 1686},
										name: "RecordFieldDefn",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 61, col: 33, offset: 1702},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 61, col: 35, offset: 1704},
									label: "rest",
									expr: &zeroOrMoreExpr{
										pos: position{line: 61, col: 40, offset: 1709},
										expr: &seqExpr{
											pos: position{line: 61, col: 41, offset: 1710},
											exprs: []interface{}{
												&litMatcher{
													pos:        position{line: 61, col: 41, offset: 1710},
													val:        ",",
													ignoreCase: false,
												},
												&ruleRefExpr{
													pos:  position{line: 61, col: 45, offset: 1714},
													name: "_",
												},
												&ruleRefExpr{
													pos:  position{line: 61, col: 47, offset: 1716},
													name: "RecordFieldDefn",
												},
												&ruleRefExpr{
													pos:  position{line: 61, col: 63, offset: 1732},
													name: "_",
												},
											},
										},
									},
								},
								&zeroOrOneExpr{
									pos: position{line: 61, col: 67, offset: 1736},
									expr: &litMatcher{
										pos:        position{line: 61, col: 67, offset: 1736},
										val:        ",",
										ignoreCase: false,
									},
								},
								&ruleRefExpr{
									pos:  position{line: 61, col: 72, offset: 1741},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 61, col: 74, offset: 1743},
									val:        "}",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 61, col: 78, offset: 1747},
									name: "__",
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 79, col: 1, offset: 2232},
						run: (*parser).callonTypeDefn54,
						expr: &seqExpr{
							pos: position{line: 79, col: 1, offset: 2232},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 79, col: 1, offset: 2232},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 79, col: 3, offset: 2234},
									val:        "type",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 79, col: 10, offset: 2241},
									name: "__",
								},
								&labeledExpr{
									pos:   position{line: 79, col: 13, offset: 2244},
									label: "name",
									expr: &ruleRefExpr{
										pos:  position{line: 79, col: 18, offset: 2249},
										name: "ModuleName",
									},
								},
								&labeledExpr{
									pos:   position{line: 79, col: 29, offset: 2260},
									label: "params",
									expr: &zeroOrMoreExpr{
										pos: position{line: 79, col: 36, offset: 2267},
										expr: &seqExpr{
											pos: position{line: 79, col: 37, offset: 2268},
											exprs: []interface{}{
												&ruleRefExpr{
													pos:  position{line: 79, col: 37, offset: 2268},
													name: "__",
												},
												&ruleRefExpr{
													pos:  position{line: 79, col: 40, offset: 2271},
													name: "TypeParameter",
												},
											},
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 79, col: 56, offset: 2287},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 79, col: 58, offset: 2289},
									val:        "=",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 79, col: 62, offset: 2293},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 79, col: 64, offset: 2295},
									label: "rest",
									expr: &oneOrMoreExpr{
										pos: position{line: 79, col: 69, offset: 2300},
										expr: &ruleRefExpr{
											pos:  position{line: 79, col: 70, offset: 2301},
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
			pos:  position{line: 94, col: 1, offset: 2708},
			expr: &actionExpr{
				pos: position{line: 94, col: 19, offset: 2726},
				run: (*parser).callonRecordFieldDefn1,
				expr: &seqExpr{
					pos: position{line: 94, col: 19, offset: 2726},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 94, col: 19, offset: 2726},
							label: "name",
							expr: &ruleRefExpr{
								pos:  position{line: 94, col: 24, offset: 2731},
								name: "VariableName",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 94, col: 37, offset: 2744},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 94, col: 39, offset: 2746},
							val:        ":",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 94, col: 43, offset: 2750},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 94, col: 45, offset: 2752},
							label: "t",
							expr: &ruleRefExpr{
								pos:  position{line: 94, col: 48, offset: 2755},
								name: "AnyType",
							},
						},
					},
				},
			},
		},
		{
			name: "VariantConstructor",
			pos:  position{line: 98, col: 1, offset: 2849},
			expr: &choiceExpr{
				pos: position{line: 98, col: 22, offset: 2870},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 98, col: 22, offset: 2870},
						run: (*parser).callonVariantConstructor2,
						expr: &seqExpr{
							pos: position{line: 98, col: 22, offset: 2870},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 98, col: 22, offset: 2870},
									val:        "|",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 98, col: 26, offset: 2874},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 98, col: 28, offset: 2876},
									label: "name",
									expr: &ruleRefExpr{
										pos:  position{line: 98, col: 33, offset: 2881},
										name: "ModuleName",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 98, col: 44, offset: 2892},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 98, col: 46, offset: 2894},
									val:        "{",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 98, col: 50, offset: 2898},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 98, col: 52, offset: 2900},
									label: "first",
									expr: &ruleRefExpr{
										pos:  position{line: 98, col: 58, offset: 2906},
										name: "RecordFieldDefn",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 98, col: 74, offset: 2922},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 98, col: 76, offset: 2924},
									label: "rest",
									expr: &zeroOrMoreExpr{
										pos: position{line: 98, col: 81, offset: 2929},
										expr: &seqExpr{
											pos: position{line: 98, col: 82, offset: 2930},
											exprs: []interface{}{
												&litMatcher{
													pos:        position{line: 98, col: 82, offset: 2930},
													val:        ",",
													ignoreCase: false,
												},
												&ruleRefExpr{
													pos:  position{line: 98, col: 86, offset: 2934},
													name: "_",
												},
												&ruleRefExpr{
													pos:  position{line: 98, col: 88, offset: 2936},
													name: "RecordFieldDefn",
												},
												&ruleRefExpr{
													pos:  position{line: 98, col: 104, offset: 2952},
													name: "_",
												},
											},
										},
									},
								},
								&zeroOrOneExpr{
									pos: position{line: 98, col: 108, offset: 2956},
									expr: &litMatcher{
										pos:        position{line: 98, col: 108, offset: 2956},
										val:        ",",
										ignoreCase: false,
									},
								},
								&ruleRefExpr{
									pos:  position{line: 98, col: 113, offset: 2961},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 98, col: 115, offset: 2963},
									val:        "}",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 98, col: 119, offset: 2967},
									name: "__",
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 117, col: 1, offset: 3572},
						run: (*parser).callonVariantConstructor26,
						expr: &seqExpr{
							pos: position{line: 117, col: 1, offset: 3572},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 117, col: 1, offset: 3572},
									val:        "|",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 117, col: 5, offset: 3576},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 117, col: 7, offset: 3578},
									label: "name",
									expr: &ruleRefExpr{
										pos:  position{line: 117, col: 12, offset: 3583},
										name: "ModuleName",
									},
								},
								&labeledExpr{
									pos:   position{line: 117, col: 23, offset: 3594},
									label: "rest",
									expr: &zeroOrMoreExpr{
										pos: position{line: 117, col: 28, offset: 3599},
										expr: &seqExpr{
											pos: position{line: 117, col: 29, offset: 3600},
											exprs: []interface{}{
												&ruleRefExpr{
													pos:  position{line: 117, col: 29, offset: 3600},
													name: "__",
												},
												&ruleRefExpr{
													pos:  position{line: 117, col: 32, offset: 3603},
													name: "AnyType",
												},
											},
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 117, col: 42, offset: 3613},
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
			pos:  position{line: 134, col: 1, offset: 4050},
			expr: &choiceExpr{
				pos: position{line: 134, col: 11, offset: 4060},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 134, col: 11, offset: 4060},
						name: "BaseType",
					},
					&ruleRefExpr{
						pos:  position{line: 134, col: 22, offset: 4071},
						name: "TypeParameter",
					},
				},
			},
		},
		{
			name: "Assignment",
			pos:  position{line: 136, col: 1, offset: 4086},
			expr: &choiceExpr{
				pos: position{line: 136, col: 14, offset: 4099},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 136, col: 14, offset: 4099},
						run: (*parser).callonAssignment2,
						expr: &seqExpr{
							pos: position{line: 136, col: 14, offset: 4099},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 136, col: 14, offset: 4099},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 136, col: 16, offset: 4101},
									val:        "let",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 136, col: 22, offset: 4107},
									name: "__",
								},
								&labeledExpr{
									pos:   position{line: 136, col: 25, offset: 4110},
									label: "i",
									expr: &ruleRefExpr{
										pos:  position{line: 136, col: 27, offset: 4112},
										name: "Assignable",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 136, col: 38, offset: 4123},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 136, col: 40, offset: 4125},
									val:        "=",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 136, col: 44, offset: 4129},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 136, col: 46, offset: 4131},
									label: "expr",
									expr: &ruleRefExpr{
										pos:  position{line: 136, col: 51, offset: 4136},
										name: "Expr",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 136, col: 56, offset: 4141},
									name: "_",
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 142, col: 1, offset: 4260},
						run: (*parser).callonAssignment15,
						expr: &seqExpr{
							pos: position{line: 142, col: 1, offset: 4260},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 142, col: 1, offset: 4260},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 142, col: 3, offset: 4262},
									val:        "let",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 142, col: 9, offset: 4268},
									name: "__",
								},
								&notExpr{
									pos: position{line: 142, col: 12, offset: 4271},
									expr: &ruleRefExpr{
										pos:  position{line: 142, col: 13, offset: 4272},
										name: "Assignable",
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 146, col: 1, offset: 4370},
						run: (*parser).callonAssignment22,
						expr: &seqExpr{
							pos: position{line: 146, col: 1, offset: 4370},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 146, col: 1, offset: 4370},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 146, col: 3, offset: 4372},
									val:        "let",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 146, col: 9, offset: 4378},
									name: "__",
								},
								&labeledExpr{
									pos:   position{line: 146, col: 12, offset: 4381},
									label: "i",
									expr: &ruleRefExpr{
										pos:  position{line: 146, col: 14, offset: 4383},
										name: "Assignable",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 146, col: 25, offset: 4394},
									name: "_",
								},
								&notExpr{
									pos: position{line: 146, col: 27, offset: 4396},
									expr: &litMatcher{
										pos:        position{line: 146, col: 28, offset: 4397},
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
			pos:  position{line: 150, col: 1, offset: 4491},
			expr: &actionExpr{
				pos: position{line: 150, col: 12, offset: 4502},
				run: (*parser).callonFuncDefn1,
				expr: &seqExpr{
					pos: position{line: 150, col: 12, offset: 4502},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 150, col: 12, offset: 4502},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 150, col: 14, offset: 4504},
							val:        "let",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 150, col: 20, offset: 4510},
							name: "__",
						},
						&labeledExpr{
							pos:   position{line: 150, col: 23, offset: 4513},
							label: "i",
							expr: &ruleRefExpr{
								pos:  position{line: 150, col: 25, offset: 4515},
								name: "VariableName",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 150, col: 38, offset: 4528},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 150, col: 40, offset: 4530},
							val:        "=",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 150, col: 44, offset: 4534},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 150, col: 46, offset: 4536},
							label: "ids",
							expr: &zeroOrOneExpr{
								pos: position{line: 150, col: 50, offset: 4540},
								expr: &seqExpr{
									pos: position{line: 150, col: 51, offset: 4541},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 150, col: 51, offset: 4541},
											name: "ArgsDefn",
										},
										&ruleRefExpr{
											pos:  position{line: 150, col: 60, offset: 4550},
											name: "_",
										},
									},
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 150, col: 64, offset: 4554},
							label: "ret",
							expr: &zeroOrOneExpr{
								pos: position{line: 150, col: 68, offset: 4558},
								expr: &seqExpr{
									pos: position{line: 150, col: 69, offset: 4559},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 150, col: 69, offset: 4559},
											name: "AnyType",
										},
										&ruleRefExpr{
											pos:  position{line: 150, col: 77, offset: 4567},
											name: "_",
										},
									},
								},
							},
						},
						&litMatcher{
							pos:        position{line: 150, col: 81, offset: 4571},
							val:        "{",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 150, col: 85, offset: 4575},
							name: "__",
						},
						&labeledExpr{
							pos:   position{line: 150, col: 88, offset: 4578},
							label: "statements",
							expr: &oneOrMoreExpr{
								pos: position{line: 150, col: 99, offset: 4589},
								expr: &ruleRefExpr{
									pos:  position{line: 150, col: 100, offset: 4590},
									name: "Statement",
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 150, col: 112, offset: 4602},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 150, col: 114, offset: 4604},
							val:        "}",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 150, col: 118, offset: 4608},
							name: "_",
						},
					},
				},
			},
		},
		{
			name: "Expr",
			pos:  position{line: 176, col: 1, offset: 5218},
			expr: &actionExpr{
				pos: position{line: 176, col: 8, offset: 5225},
				run: (*parser).callonExpr1,
				expr: &labeledExpr{
					pos:   position{line: 176, col: 8, offset: 5225},
					label: "ex",
					expr: &choiceExpr{
						pos: position{line: 176, col: 12, offset: 5229},
						alternatives: []interface{}{
							&ruleRefExpr{
								pos:  position{line: 176, col: 12, offset: 5229},
								name: "IfExpr",
							},
							&ruleRefExpr{
								pos:  position{line: 176, col: 21, offset: 5238},
								name: "BinOp",
							},
							&ruleRefExpr{
								pos:  position{line: 176, col: 29, offset: 5246},
								name: "Call",
							},
						},
					},
				},
			},
		},
		{
			name: "IfExpr",
			pos:  position{line: 182, col: 1, offset: 5355},
			expr: &choiceExpr{
				pos: position{line: 182, col: 10, offset: 5364},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 182, col: 10, offset: 5364},
						run: (*parser).callonIfExpr2,
						expr: &seqExpr{
							pos: position{line: 182, col: 10, offset: 5364},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 182, col: 10, offset: 5364},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 182, col: 12, offset: 5366},
									val:        "if",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 182, col: 17, offset: 5371},
									name: "__",
								},
								&labeledExpr{
									pos:   position{line: 182, col: 20, offset: 5374},
									label: "expr",
									expr: &ruleRefExpr{
										pos:  position{line: 182, col: 25, offset: 5379},
										name: "BinOpBool",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 182, col: 35, offset: 5389},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 182, col: 37, offset: 5391},
									val:        "{",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 182, col: 41, offset: 5395},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 182, col: 43, offset: 5397},
									label: "thens",
									expr: &oneOrMoreExpr{
										pos: position{line: 182, col: 49, offset: 5403},
										expr: &ruleRefExpr{
											pos:  position{line: 182, col: 50, offset: 5404},
											name: "Statement",
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 182, col: 62, offset: 5416},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 182, col: 64, offset: 5418},
									val:        "}",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 182, col: 68, offset: 5422},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 182, col: 70, offset: 5424},
									val:        "else",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 182, col: 77, offset: 5431},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 182, col: 79, offset: 5433},
									label: "elseifs",
									expr: &ruleRefExpr{
										pos:  position{line: 182, col: 87, offset: 5441},
										name: "IfExpr",
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 194, col: 1, offset: 5771},
						run: (*parser).callonIfExpr22,
						expr: &seqExpr{
							pos: position{line: 194, col: 1, offset: 5771},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 194, col: 1, offset: 5771},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 194, col: 3, offset: 5773},
									val:        "if",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 194, col: 8, offset: 5778},
									name: "__",
								},
								&labeledExpr{
									pos:   position{line: 194, col: 11, offset: 5781},
									label: "expr",
									expr: &ruleRefExpr{
										pos:  position{line: 194, col: 16, offset: 5786},
										name: "BinOpBool",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 194, col: 26, offset: 5796},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 194, col: 28, offset: 5798},
									val:        "{",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 194, col: 32, offset: 5802},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 194, col: 34, offset: 5804},
									label: "thens",
									expr: &oneOrMoreExpr{
										pos: position{line: 194, col: 40, offset: 5810},
										expr: &ruleRefExpr{
											pos:  position{line: 194, col: 41, offset: 5811},
											name: "Statement",
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 194, col: 53, offset: 5823},
									name: "__",
								},
								&litMatcher{
									pos:        position{line: 194, col: 56, offset: 5826},
									val:        "}",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 194, col: 60, offset: 5830},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 194, col: 62, offset: 5832},
									val:        "else",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 194, col: 69, offset: 5839},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 194, col: 71, offset: 5841},
									val:        "{",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 194, col: 75, offset: 5845},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 194, col: 77, offset: 5847},
									label: "elses",
									expr: &oneOrMoreExpr{
										pos: position{line: 194, col: 83, offset: 5853},
										expr: &ruleRefExpr{
											pos:  position{line: 194, col: 84, offset: 5854},
											name: "Statement",
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 194, col: 96, offset: 5866},
									name: "__",
								},
								&litMatcher{
									pos:        position{line: 194, col: 99, offset: 5869},
									val:        "}",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 213, col: 1, offset: 6372},
						run: (*parser).callonIfExpr47,
						expr: &seqExpr{
							pos: position{line: 213, col: 1, offset: 6372},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 213, col: 1, offset: 6372},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 213, col: 3, offset: 6374},
									val:        "if",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 213, col: 8, offset: 6379},
									name: "__",
								},
								&labeledExpr{
									pos:   position{line: 213, col: 11, offset: 6382},
									label: "expr",
									expr: &ruleRefExpr{
										pos:  position{line: 213, col: 16, offset: 6387},
										name: "BinOpBool",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 213, col: 26, offset: 6397},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 213, col: 28, offset: 6399},
									val:        "{",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 213, col: 32, offset: 6403},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 213, col: 34, offset: 6405},
									label: "thens",
									expr: &oneOrMoreExpr{
										pos: position{line: 213, col: 40, offset: 6411},
										expr: &ruleRefExpr{
											pos:  position{line: 213, col: 41, offset: 6412},
											name: "Statement",
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 213, col: 53, offset: 6424},
									name: "__",
								},
								&litMatcher{
									pos:        position{line: 213, col: 56, offset: 6427},
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
			pos:  position{line: 225, col: 1, offset: 6725},
			expr: &choiceExpr{
				pos: position{line: 225, col: 8, offset: 6732},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 225, col: 8, offset: 6732},
						run: (*parser).callonCall2,
						expr: &seqExpr{
							pos: position{line: 225, col: 8, offset: 6732},
							exprs: []interface{}{
								&labeledExpr{
									pos:   position{line: 225, col: 8, offset: 6732},
									label: "module",
									expr: &ruleRefExpr{
										pos:  position{line: 225, col: 15, offset: 6739},
										name: "ModuleName",
									},
								},
								&litMatcher{
									pos:        position{line: 225, col: 26, offset: 6750},
									val:        ".",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 225, col: 30, offset: 6754},
									label: "fn",
									expr: &ruleRefExpr{
										pos:  position{line: 225, col: 33, offset: 6757},
										name: "VariableName",
									},
								},
								&labeledExpr{
									pos:   position{line: 225, col: 46, offset: 6770},
									label: "args",
									expr: &ruleRefExpr{
										pos:  position{line: 225, col: 51, offset: 6775},
										name: "Arguments",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 225, col: 61, offset: 6785},
									name: "_",
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 239, col: 1, offset: 7101},
						run: (*parser).callonCall12,
						expr: &seqExpr{
							pos: position{line: 239, col: 1, offset: 7101},
							exprs: []interface{}{
								&labeledExpr{
									pos:   position{line: 239, col: 1, offset: 7101},
									label: "fn",
									expr: &ruleRefExpr{
										pos:  position{line: 239, col: 4, offset: 7104},
										name: "VariableName",
									},
								},
								&labeledExpr{
									pos:   position{line: 239, col: 17, offset: 7117},
									label: "args",
									expr: &ruleRefExpr{
										pos:  position{line: 239, col: 22, offset: 7122},
										name: "Arguments",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 239, col: 32, offset: 7132},
									name: "_",
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 253, col: 1, offset: 7441},
						run: (*parser).callonCall19,
						expr: &seqExpr{
							pos: position{line: 253, col: 1, offset: 7441},
							exprs: []interface{}{
								&labeledExpr{
									pos:   position{line: 253, col: 1, offset: 7441},
									label: "fn",
									expr: &ruleRefExpr{
										pos:  position{line: 253, col: 4, offset: 7444},
										name: "VariableName",
									},
								},
								&litMatcher{
									pos:        position{line: 253, col: 17, offset: 7457},
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
			pos:  position{line: 260, col: 1, offset: 7628},
			expr: &actionExpr{
				pos: position{line: 260, col: 12, offset: 7639},
				run: (*parser).callonArgsDefn1,
				expr: &seqExpr{
					pos: position{line: 260, col: 12, offset: 7639},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 260, col: 12, offset: 7639},
							val:        "(",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 260, col: 16, offset: 7643},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 260, col: 18, offset: 7645},
							label: "argument",
							expr: &ruleRefExpr{
								pos:  position{line: 260, col: 27, offset: 7654},
								name: "ArgDefn",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 260, col: 35, offset: 7662},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 260, col: 37, offset: 7664},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 260, col: 42, offset: 7669},
								expr: &seqExpr{
									pos: position{line: 260, col: 43, offset: 7670},
									exprs: []interface{}{
										&litMatcher{
											pos:        position{line: 260, col: 43, offset: 7670},
											val:        ",",
											ignoreCase: false,
										},
										&ruleRefExpr{
											pos:  position{line: 260, col: 47, offset: 7674},
											name: "_",
										},
										&ruleRefExpr{
											pos:  position{line: 260, col: 49, offset: 7676},
											name: "ArgDefn",
										},
									},
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 260, col: 59, offset: 7686},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 260, col: 61, offset: 7688},
							val:        ")",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "ArgDefn",
			pos:  position{line: 278, col: 1, offset: 8110},
			expr: &actionExpr{
				pos: position{line: 278, col: 11, offset: 8120},
				run: (*parser).callonArgDefn1,
				expr: &seqExpr{
					pos: position{line: 278, col: 11, offset: 8120},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 278, col: 11, offset: 8120},
							label: "name",
							expr: &ruleRefExpr{
								pos:  position{line: 278, col: 16, offset: 8125},
								name: "Identifier",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 278, col: 27, offset: 8136},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 278, col: 29, offset: 8138},
							label: "anno",
							expr: &zeroOrOneExpr{
								pos: position{line: 278, col: 34, offset: 8143},
								expr: &seqExpr{
									pos: position{line: 278, col: 35, offset: 8144},
									exprs: []interface{}{
										&litMatcher{
											pos:        position{line: 278, col: 35, offset: 8144},
											val:        ":",
											ignoreCase: false,
										},
										&ruleRefExpr{
											pos:  position{line: 278, col: 39, offset: 8148},
											name: "_",
										},
										&ruleRefExpr{
											pos:  position{line: 278, col: 41, offset: 8150},
											name: "AnyType",
										},
									},
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 278, col: 52, offset: 8161},
							name: "_",
						},
					},
				},
			},
		},
		{
			name: "Arguments",
			pos:  position{line: 298, col: 1, offset: 8643},
			expr: &choiceExpr{
				pos: position{line: 298, col: 13, offset: 8655},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 298, col: 13, offset: 8655},
						run: (*parser).callonArguments2,
						expr: &seqExpr{
							pos: position{line: 298, col: 13, offset: 8655},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 298, col: 13, offset: 8655},
									val:        "(",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 298, col: 17, offset: 8659},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 298, col: 19, offset: 8661},
									label: "argument",
									expr: &ruleRefExpr{
										pos:  position{line: 298, col: 28, offset: 8670},
										name: "BinOpParens",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 298, col: 40, offset: 8682},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 298, col: 42, offset: 8684},
									label: "rest",
									expr: &zeroOrMoreExpr{
										pos: position{line: 298, col: 47, offset: 8689},
										expr: &seqExpr{
											pos: position{line: 298, col: 48, offset: 8690},
											exprs: []interface{}{
												&litMatcher{
													pos:        position{line: 298, col: 48, offset: 8690},
													val:        ",",
													ignoreCase: false,
												},
												&ruleRefExpr{
													pos:  position{line: 298, col: 52, offset: 8694},
													name: "_",
												},
												&ruleRefExpr{
													pos:  position{line: 298, col: 54, offset: 8696},
													name: "BinOpParens",
												},
											},
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 298, col: 68, offset: 8710},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 298, col: 70, offset: 8712},
									val:        ")",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 315, col: 1, offset: 9134},
						run: (*parser).callonArguments17,
						expr: &seqExpr{
							pos: position{line: 315, col: 1, offset: 9134},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 315, col: 1, offset: 9134},
									val:        "(",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 315, col: 5, offset: 9138},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 315, col: 7, offset: 9140},
									label: "argument",
									expr: &ruleRefExpr{
										pos:  position{line: 315, col: 16, offset: 9149},
										name: "Expr",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 315, col: 21, offset: 9154},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 315, col: 23, offset: 9156},
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
			pos:  position{line: 320, col: 1, offset: 9261},
			expr: &actionExpr{
				pos: position{line: 320, col: 16, offset: 9276},
				run: (*parser).callonCompoundExpr1,
				expr: &seqExpr{
					pos: position{line: 320, col: 16, offset: 9276},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 320, col: 16, offset: 9276},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 320, col: 18, offset: 9278},
							label: "op",
							expr: &ruleRefExpr{
								pos:  position{line: 320, col: 21, offset: 9281},
								name: "BinOp",
							},
						},
						&labeledExpr{
							pos:   position{line: 320, col: 27, offset: 9287},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 320, col: 32, offset: 9292},
								expr: &seqExpr{
									pos: position{line: 320, col: 33, offset: 9293},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 320, col: 33, offset: 9293},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 320, col: 36, offset: 9296},
											name: "Operator",
										},
										&ruleRefExpr{
											pos:  position{line: 320, col: 45, offset: 9305},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 320, col: 48, offset: 9308},
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
			pos:  position{line: 340, col: 1, offset: 9914},
			expr: &choiceExpr{
				pos: position{line: 340, col: 9, offset: 9922},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 340, col: 9, offset: 9922},
						name: "BinOpBool",
					},
					&ruleRefExpr{
						pos:  position{line: 340, col: 21, offset: 9934},
						name: "BinOpEquality",
					},
					&ruleRefExpr{
						pos:  position{line: 340, col: 37, offset: 9950},
						name: "BinOpLow",
					},
					&ruleRefExpr{
						pos:  position{line: 340, col: 48, offset: 9961},
						name: "BinOpHigh",
					},
					&ruleRefExpr{
						pos:  position{line: 340, col: 60, offset: 9973},
						name: "BinOpParens",
					},
				},
			},
		},
		{
			name: "BinOpBool",
			pos:  position{line: 342, col: 1, offset: 9986},
			expr: &actionExpr{
				pos: position{line: 342, col: 13, offset: 9998},
				run: (*parser).callonBinOpBool1,
				expr: &seqExpr{
					pos: position{line: 342, col: 13, offset: 9998},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 342, col: 13, offset: 9998},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 342, col: 15, offset: 10000},
							label: "first",
							expr: &ruleRefExpr{
								pos:  position{line: 342, col: 21, offset: 10006},
								name: "BinOpEquality",
							},
						},
						&labeledExpr{
							pos:   position{line: 342, col: 35, offset: 10020},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 342, col: 40, offset: 10025},
								expr: &seqExpr{
									pos: position{line: 342, col: 41, offset: 10026},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 342, col: 41, offset: 10026},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 342, col: 44, offset: 10029},
											name: "OperatorBoolean",
										},
										&ruleRefExpr{
											pos:  position{line: 342, col: 60, offset: 10045},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 342, col: 63, offset: 10048},
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
			pos:  position{line: 375, col: 1, offset: 10941},
			expr: &actionExpr{
				pos: position{line: 375, col: 17, offset: 10957},
				run: (*parser).callonBinOpEquality1,
				expr: &seqExpr{
					pos: position{line: 375, col: 17, offset: 10957},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 375, col: 17, offset: 10957},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 375, col: 19, offset: 10959},
							label: "first",
							expr: &ruleRefExpr{
								pos:  position{line: 375, col: 25, offset: 10965},
								name: "BinOpLow",
							},
						},
						&labeledExpr{
							pos:   position{line: 375, col: 34, offset: 10974},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 375, col: 39, offset: 10979},
								expr: &seqExpr{
									pos: position{line: 375, col: 40, offset: 10980},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 375, col: 40, offset: 10980},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 375, col: 43, offset: 10983},
											name: "OperatorEquality",
										},
										&ruleRefExpr{
											pos:  position{line: 375, col: 60, offset: 11000},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 375, col: 63, offset: 11003},
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
			pos:  position{line: 407, col: 1, offset: 11890},
			expr: &actionExpr{
				pos: position{line: 407, col: 12, offset: 11901},
				run: (*parser).callonBinOpLow1,
				expr: &seqExpr{
					pos: position{line: 407, col: 12, offset: 11901},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 407, col: 12, offset: 11901},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 407, col: 14, offset: 11903},
							label: "first",
							expr: &ruleRefExpr{
								pos:  position{line: 407, col: 20, offset: 11909},
								name: "BinOpHigh",
							},
						},
						&labeledExpr{
							pos:   position{line: 407, col: 30, offset: 11919},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 407, col: 35, offset: 11924},
								expr: &seqExpr{
									pos: position{line: 407, col: 36, offset: 11925},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 407, col: 36, offset: 11925},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 407, col: 39, offset: 11928},
											name: "OperatorLow",
										},
										&ruleRefExpr{
											pos:  position{line: 407, col: 51, offset: 11940},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 407, col: 54, offset: 11943},
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
			pos:  position{line: 439, col: 1, offset: 12831},
			expr: &actionExpr{
				pos: position{line: 439, col: 13, offset: 12843},
				run: (*parser).callonBinOpHigh1,
				expr: &seqExpr{
					pos: position{line: 439, col: 13, offset: 12843},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 439, col: 13, offset: 12843},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 439, col: 15, offset: 12845},
							label: "first",
							expr: &ruleRefExpr{
								pos:  position{line: 439, col: 21, offset: 12851},
								name: "BinOpParens",
							},
						},
						&labeledExpr{
							pos:   position{line: 439, col: 33, offset: 12863},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 439, col: 38, offset: 12868},
								expr: &seqExpr{
									pos: position{line: 439, col: 39, offset: 12869},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 439, col: 39, offset: 12869},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 439, col: 42, offset: 12872},
											name: "OperatorHigh",
										},
										&ruleRefExpr{
											pos:  position{line: 439, col: 55, offset: 12885},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 439, col: 58, offset: 12888},
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
			pos:  position{line: 470, col: 1, offset: 13777},
			expr: &choiceExpr{
				pos: position{line: 470, col: 15, offset: 13791},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 470, col: 15, offset: 13791},
						run: (*parser).callonBinOpParens2,
						expr: &seqExpr{
							pos: position{line: 470, col: 15, offset: 13791},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 470, col: 15, offset: 13791},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 470, col: 17, offset: 13793},
									val:        "(",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 470, col: 21, offset: 13797},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 470, col: 23, offset: 13799},
									label: "first",
									expr: &ruleRefExpr{
										pos:  position{line: 470, col: 29, offset: 13805},
										name: "BinOp",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 470, col: 35, offset: 13811},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 470, col: 37, offset: 13813},
									val:        ")",
									ignoreCase: false,
								},
							},
						},
					},
					&ruleRefExpr{
						pos:  position{line: 473, col: 5, offset: 13936},
						name: "Value",
					},
				},
			},
		},
		{
			name: "Operator",
			pos:  position{line: 475, col: 1, offset: 13943},
			expr: &choiceExpr{
				pos: position{line: 475, col: 12, offset: 13954},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 475, col: 12, offset: 13954},
						name: "OperatorBoolean",
					},
					&ruleRefExpr{
						pos:  position{line: 475, col: 30, offset: 13972},
						name: "OperatorEquality",
					},
					&ruleRefExpr{
						pos:  position{line: 475, col: 49, offset: 13991},
						name: "OperatorHigh",
					},
					&ruleRefExpr{
						pos:  position{line: 475, col: 64, offset: 14006},
						name: "OperatorLow",
					},
				},
			},
		},
		{
			name: "OperatorBoolean",
			pos:  position{line: 477, col: 1, offset: 14019},
			expr: &actionExpr{
				pos: position{line: 477, col: 19, offset: 14037},
				run: (*parser).callonOperatorBoolean1,
				expr: &choiceExpr{
					pos: position{line: 477, col: 21, offset: 14039},
					alternatives: []interface{}{
						&litMatcher{
							pos:        position{line: 477, col: 21, offset: 14039},
							val:        "or",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 477, col: 28, offset: 14046},
							val:        "and",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "OperatorEquality",
			pos:  position{line: 481, col: 1, offset: 14128},
			expr: &actionExpr{
				pos: position{line: 481, col: 20, offset: 14147},
				run: (*parser).callonOperatorEquality1,
				expr: &choiceExpr{
					pos: position{line: 481, col: 22, offset: 14149},
					alternatives: []interface{}{
						&litMatcher{
							pos:        position{line: 481, col: 22, offset: 14149},
							val:        ">=",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 481, col: 29, offset: 14156},
							val:        "<=",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 481, col: 36, offset: 14163},
							val:        "<",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 481, col: 42, offset: 14169},
							val:        ">",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 481, col: 48, offset: 14175},
							val:        "===",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 481, col: 56, offset: 14183},
							val:        "==",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "OperatorHigh",
			pos:  position{line: 485, col: 1, offset: 14262},
			expr: &choiceExpr{
				pos: position{line: 485, col: 16, offset: 14277},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 485, col: 16, offset: 14277},
						run: (*parser).callonOperatorHigh2,
						expr: &choiceExpr{
							pos: position{line: 485, col: 18, offset: 14279},
							alternatives: []interface{}{
								&litMatcher{
									pos:        position{line: 485, col: 18, offset: 14279},
									val:        "/",
									ignoreCase: false,
								},
								&litMatcher{
									pos:        position{line: 485, col: 24, offset: 14285},
									val:        "*",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 488, col: 3, offset: 14368},
						run: (*parser).callonOperatorHigh6,
						expr: &litMatcher{
							pos:        position{line: 488, col: 5, offset: 14370},
							val:        "^",
							ignoreCase: false,
						},
					},
					&actionExpr{
						pos: position{line: 491, col: 3, offset: 14450},
						run: (*parser).callonOperatorHigh8,
						expr: &litMatcher{
							pos:        position{line: 491, col: 3, offset: 14450},
							val:        "++",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "OperatorLow",
			pos:  position{line: 495, col: 1, offset: 14531},
			expr: &actionExpr{
				pos: position{line: 495, col: 15, offset: 14545},
				run: (*parser).callonOperatorLow1,
				expr: &choiceExpr{
					pos: position{line: 495, col: 17, offset: 14547},
					alternatives: []interface{}{
						&litMatcher{
							pos:        position{line: 495, col: 17, offset: 14547},
							val:        "+",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 495, col: 23, offset: 14553},
							val:        "-",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "Value",
			pos:  position{line: 499, col: 1, offset: 14635},
			expr: &choiceExpr{
				pos: position{line: 499, col: 9, offset: 14643},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 499, col: 9, offset: 14643},
						name: "Call",
					},
					&ruleRefExpr{
						pos:  position{line: 499, col: 16, offset: 14650},
						name: "VariableName",
					},
					&ruleRefExpr{
						pos:  position{line: 499, col: 31, offset: 14665},
						name: "Const",
					},
				},
			},
		},
		{
			name: "Assignable",
			pos:  position{line: 501, col: 1, offset: 14672},
			expr: &choiceExpr{
				pos: position{line: 501, col: 14, offset: 14685},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 501, col: 14, offset: 14685},
						name: "VariableName",
					},
					&ruleRefExpr{
						pos:  position{line: 501, col: 29, offset: 14700},
						name: "Unused",
					},
				},
			},
		},
		{
			name: "Identifier",
			pos:  position{line: 503, col: 1, offset: 14708},
			expr: &choiceExpr{
				pos: position{line: 503, col: 14, offset: 14721},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 503, col: 14, offset: 14721},
						name: "VariableName",
					},
					&ruleRefExpr{
						pos:  position{line: 503, col: 29, offset: 14736},
						name: "ModuleName",
					},
				},
			},
		},
		{
			name: "ArrayLiteral",
			pos:  position{line: 505, col: 1, offset: 14748},
			expr: &actionExpr{
				pos: position{line: 505, col: 16, offset: 14763},
				run: (*parser).callonArrayLiteral1,
				expr: &seqExpr{
					pos: position{line: 505, col: 16, offset: 14763},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 505, col: 16, offset: 14763},
							val:        "[",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 505, col: 20, offset: 14767},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 505, col: 22, offset: 14769},
							label: "first",
							expr: &ruleRefExpr{
								pos:  position{line: 505, col: 28, offset: 14775},
								name: "Expr",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 505, col: 33, offset: 14780},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 505, col: 35, offset: 14782},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 505, col: 40, offset: 14787},
								expr: &seqExpr{
									pos: position{line: 505, col: 41, offset: 14788},
									exprs: []interface{}{
										&litMatcher{
											pos:        position{line: 505, col: 41, offset: 14788},
											val:        ",",
											ignoreCase: false,
										},
										&ruleRefExpr{
											pos:  position{line: 505, col: 45, offset: 14792},
											name: "_",
										},
										&ruleRefExpr{
											pos:  position{line: 505, col: 47, offset: 14794},
											name: "Expr",
										},
										&ruleRefExpr{
											pos:  position{line: 505, col: 52, offset: 14799},
											name: "_",
										},
									},
								},
							},
						},
						&zeroOrOneExpr{
							pos: position{line: 505, col: 56, offset: 14803},
							expr: &litMatcher{
								pos:        position{line: 505, col: 56, offset: 14803},
								val:        ",",
								ignoreCase: false,
							},
						},
						&ruleRefExpr{
							pos:  position{line: 505, col: 61, offset: 14808},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 505, col: 63, offset: 14810},
							val:        "]",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "VariantInstance",
			pos:  position{line: 521, col: 1, offset: 15255},
			expr: &actionExpr{
				pos: position{line: 521, col: 19, offset: 15273},
				run: (*parser).callonVariantInstance1,
				expr: &seqExpr{
					pos: position{line: 521, col: 19, offset: 15273},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 521, col: 19, offset: 15273},
							label: "name",
							expr: &ruleRefExpr{
								pos:  position{line: 521, col: 24, offset: 15278},
								name: "ModuleName",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 521, col: 35, offset: 15289},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 521, col: 37, offset: 15291},
							label: "args",
							expr: &ruleRefExpr{
								pos:  position{line: 521, col: 42, offset: 15296},
								name: "Arguments",
							},
						},
					},
				},
			},
		},
		{
			name: "RecordInstance",
			pos:  position{line: 534, col: 1, offset: 15566},
			expr: &actionExpr{
				pos: position{line: 534, col: 18, offset: 15583},
				run: (*parser).callonRecordInstance1,
				expr: &seqExpr{
					pos: position{line: 534, col: 18, offset: 15583},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 534, col: 18, offset: 15583},
							label: "name",
							expr: &ruleRefExpr{
								pos:  position{line: 534, col: 23, offset: 15588},
								name: "ModuleName",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 534, col: 34, offset: 15599},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 534, col: 36, offset: 15601},
							val:        "{",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 534, col: 40, offset: 15605},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 534, col: 42, offset: 15607},
							label: "firstName",
							expr: &ruleRefExpr{
								pos:  position{line: 534, col: 52, offset: 15617},
								name: "VariableName",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 534, col: 65, offset: 15630},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 534, col: 67, offset: 15632},
							val:        ":",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 534, col: 71, offset: 15636},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 534, col: 73, offset: 15638},
							label: "firstValue",
							expr: &ruleRefExpr{
								pos:  position{line: 534, col: 84, offset: 15649},
								name: "Expr",
							},
						},
						&labeledExpr{
							pos:   position{line: 534, col: 89, offset: 15654},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 534, col: 94, offset: 15659},
								expr: &seqExpr{
									pos: position{line: 534, col: 95, offset: 15660},
									exprs: []interface{}{
										&litMatcher{
											pos:        position{line: 534, col: 95, offset: 15660},
											val:        ",",
											ignoreCase: false,
										},
										&ruleRefExpr{
											pos:  position{line: 534, col: 99, offset: 15664},
											name: "_",
										},
										&ruleRefExpr{
											pos:  position{line: 534, col: 101, offset: 15666},
											name: "VariableName",
										},
										&ruleRefExpr{
											pos:  position{line: 534, col: 114, offset: 15679},
											name: "_",
										},
										&litMatcher{
											pos:        position{line: 534, col: 116, offset: 15681},
											val:        ":",
											ignoreCase: false,
										},
										&ruleRefExpr{
											pos:  position{line: 534, col: 120, offset: 15685},
											name: "_",
										},
										&ruleRefExpr{
											pos:  position{line: 534, col: 122, offset: 15687},
											name: "Expr",
										},
									},
								},
							},
						},
						&litMatcher{
							pos:        position{line: 534, col: 130, offset: 15695},
							val:        "}",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "TypeParameter",
			pos:  position{line: 554, col: 1, offset: 16279},
			expr: &actionExpr{
				pos: position{line: 554, col: 17, offset: 16295},
				run: (*parser).callonTypeParameter1,
				expr: &seqExpr{
					pos: position{line: 554, col: 17, offset: 16295},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 554, col: 17, offset: 16295},
							val:        "'",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 554, col: 22, offset: 16300},
							name: "VariableName",
						},
					},
				},
			},
		},
		{
			name: "VariableName",
			pos:  position{line: 558, col: 1, offset: 16373},
			expr: &actionExpr{
				pos: position{line: 558, col: 16, offset: 16388},
				run: (*parser).callonVariableName1,
				expr: &seqExpr{
					pos: position{line: 558, col: 16, offset: 16388},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 558, col: 16, offset: 16388},
							expr: &ruleRefExpr{
								pos:  position{line: 558, col: 17, offset: 16389},
								name: "Reserved",
							},
						},
						&seqExpr{
							pos: position{line: 558, col: 27, offset: 16399},
							exprs: []interface{}{
								&oneOrMoreExpr{
									pos: position{line: 558, col: 27, offset: 16399},
									expr: &charClassMatcher{
										pos:        position{line: 558, col: 27, offset: 16399},
										val:        "[a-z]",
										ranges:     []rune{'a', 'z'},
										ignoreCase: false,
										inverted:   false,
									},
								},
								&zeroOrMoreExpr{
									pos: position{line: 558, col: 34, offset: 16406},
									expr: &charClassMatcher{
										pos:        position{line: 558, col: 34, offset: 16406},
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
			pos:  position{line: 562, col: 1, offset: 16481},
			expr: &actionExpr{
				pos: position{line: 562, col: 14, offset: 16494},
				run: (*parser).callonModuleName1,
				expr: &seqExpr{
					pos: position{line: 562, col: 15, offset: 16495},
					exprs: []interface{}{
						&oneOrMoreExpr{
							pos: position{line: 562, col: 15, offset: 16495},
							expr: &charClassMatcher{
								pos:        position{line: 562, col: 15, offset: 16495},
								val:        "[A-Z]",
								ranges:     []rune{'A', 'Z'},
								ignoreCase: false,
								inverted:   false,
							},
						},
						&zeroOrMoreExpr{
							pos: position{line: 562, col: 22, offset: 16502},
							expr: &charClassMatcher{
								pos:        position{line: 562, col: 22, offset: 16502},
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
			pos:  position{line: 566, col: 1, offset: 16577},
			expr: &choiceExpr{
				pos: position{line: 566, col: 9, offset: 16585},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 566, col: 9, offset: 16585},
						run: (*parser).callonConst2,
						expr: &seqExpr{
							pos: position{line: 566, col: 9, offset: 16585},
							exprs: []interface{}{
								&zeroOrOneExpr{
									pos: position{line: 566, col: 9, offset: 16585},
									expr: &litMatcher{
										pos:        position{line: 566, col: 9, offset: 16585},
										val:        "-",
										ignoreCase: false,
									},
								},
								&oneOrMoreExpr{
									pos: position{line: 566, col: 14, offset: 16590},
									expr: &charClassMatcher{
										pos:        position{line: 566, col: 14, offset: 16590},
										val:        "[0-9]",
										ranges:     []rune{'0', '9'},
										ignoreCase: false,
										inverted:   false,
									},
								},
								&notExpr{
									pos: position{line: 566, col: 21, offset: 16597},
									expr: &litMatcher{
										pos:        position{line: 566, col: 22, offset: 16598},
										val:        ".",
										ignoreCase: false,
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 573, col: 3, offset: 16773},
						run: (*parser).callonConst10,
						expr: &seqExpr{
							pos: position{line: 573, col: 3, offset: 16773},
							exprs: []interface{}{
								&zeroOrOneExpr{
									pos: position{line: 573, col: 3, offset: 16773},
									expr: &litMatcher{
										pos:        position{line: 573, col: 3, offset: 16773},
										val:        "-",
										ignoreCase: false,
									},
								},
								&oneOrMoreExpr{
									pos: position{line: 573, col: 8, offset: 16778},
									expr: &charClassMatcher{
										pos:        position{line: 573, col: 8, offset: 16778},
										val:        "[0-9]",
										ranges:     []rune{'0', '9'},
										ignoreCase: false,
										inverted:   false,
									},
								},
								&litMatcher{
									pos:        position{line: 573, col: 15, offset: 16785},
									val:        ".",
									ignoreCase: false,
								},
								&oneOrMoreExpr{
									pos: position{line: 573, col: 19, offset: 16789},
									expr: &charClassMatcher{
										pos:        position{line: 573, col: 19, offset: 16789},
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
						pos: position{line: 580, col: 3, offset: 16978},
						run: (*parser).callonConst19,
						expr: &litMatcher{
							pos:        position{line: 580, col: 3, offset: 16978},
							val:        "true",
							ignoreCase: false,
						},
					},
					&actionExpr{
						pos: position{line: 584, col: 3, offset: 17063},
						run: (*parser).callonConst21,
						expr: &litMatcher{
							pos:        position{line: 584, col: 3, offset: 17063},
							val:        "false",
							ignoreCase: false,
						},
					},
					&ruleRefExpr{
						pos:  position{line: 587, col: 3, offset: 17149},
						name: "Unit",
					},
					&ruleRefExpr{
						pos:  position{line: 588, col: 3, offset: 17156},
						name: "StringLiteral",
					},
					&actionExpr{
						pos: position{line: 589, col: 3, offset: 17172},
						run: (*parser).callonConst25,
						expr: &seqExpr{
							pos: position{line: 589, col: 3, offset: 17172},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 589, col: 3, offset: 17172},
									val:        "'",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 589, col: 7, offset: 17176},
									label: "val",
									expr: &seqExpr{
										pos: position{line: 589, col: 12, offset: 17181},
										exprs: []interface{}{
											&notExpr{
												pos: position{line: 589, col: 12, offset: 17181},
												expr: &ruleRefExpr{
													pos:  position{line: 589, col: 13, offset: 17182},
													name: "EscapedChar",
												},
											},
											&anyMatcher{
												line: 589, col: 25, offset: 17194,
											},
										},
									},
								},
								&litMatcher{
									pos:        position{line: 589, col: 28, offset: 17197},
									val:        "'",
									ignoreCase: false,
								},
							},
						},
					},
					&ruleRefExpr{
						pos:  position{line: 591, col: 5, offset: 17289},
						name: "ArrayLiteral",
					},
					&ruleRefExpr{
						pos:  position{line: 591, col: 20, offset: 17304},
						name: "RecordInstance",
					},
					&ruleRefExpr{
						pos:  position{line: 591, col: 37, offset: 17321},
						name: "VariantInstance",
					},
				},
			},
		},
		{
			name: "Unit",
			pos:  position{line: 593, col: 1, offset: 17338},
			expr: &actionExpr{
				pos: position{line: 593, col: 8, offset: 17345},
				run: (*parser).callonUnit1,
				expr: &litMatcher{
					pos:        position{line: 593, col: 8, offset: 17345},
					val:        "()",
					ignoreCase: false,
				},
			},
		},
		{
			name: "StringLiteral",
			pos:  position{line: 597, col: 1, offset: 17408},
			expr: &actionExpr{
				pos: position{line: 597, col: 17, offset: 17424},
				run: (*parser).callonStringLiteral1,
				expr: &seqExpr{
					pos: position{line: 597, col: 17, offset: 17424},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 597, col: 17, offset: 17424},
							val:        "\"",
							ignoreCase: false,
						},
						&zeroOrMoreExpr{
							pos: position{line: 597, col: 21, offset: 17428},
							expr: &seqExpr{
								pos: position{line: 597, col: 22, offset: 17429},
								exprs: []interface{}{
									&notExpr{
										pos: position{line: 597, col: 22, offset: 17429},
										expr: &ruleRefExpr{
											pos:  position{line: 597, col: 23, offset: 17430},
											name: "EscapedChar",
										},
									},
									&anyMatcher{
										line: 597, col: 35, offset: 17442,
									},
								},
							},
						},
						&litMatcher{
							pos:        position{line: 597, col: 39, offset: 17446},
							val:        "\"",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "Unused",
			pos:  position{line: 605, col: 1, offset: 17629},
			expr: &actionExpr{
				pos: position{line: 605, col: 10, offset: 17638},
				run: (*parser).callonUnused1,
				expr: &litMatcher{
					pos:        position{line: 605, col: 11, offset: 17639},
					val:        "_",
					ignoreCase: false,
				},
			},
		},
		{
			name: "Reserved",
			pos:  position{line: 609, col: 1, offset: 17694},
			expr: &seqExpr{
				pos: position{line: 609, col: 12, offset: 17705},
				exprs: []interface{}{
					&choiceExpr{
						pos: position{line: 609, col: 13, offset: 17706},
						alternatives: []interface{}{
							&litMatcher{
								pos:        position{line: 609, col: 13, offset: 17706},
								val:        "let",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 609, col: 21, offset: 17714},
								val:        "if",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 609, col: 28, offset: 17721},
								val:        "else",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 609, col: 37, offset: 17730},
								val:        "func",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 609, col: 46, offset: 17739},
								val:        "type",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 609, col: 55, offset: 17748},
								val:        "true",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 609, col: 64, offset: 17757},
								val:        "false",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 609, col: 74, offset: 17767},
								val:        "mutable",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 609, col: 86, offset: 17779},
								val:        "type",
								ignoreCase: false,
							},
							&ruleRefExpr{
								pos:  position{line: 609, col: 95, offset: 17788},
								name: "BaseType",
							},
						},
					},
					&notExpr{
						pos: position{line: 609, col: 105, offset: 17798},
						expr: &oneOrMoreExpr{
							pos: position{line: 609, col: 106, offset: 17799},
							expr: &charClassMatcher{
								pos:        position{line: 609, col: 106, offset: 17799},
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
			pos:  position{line: 611, col: 1, offset: 17807},
			expr: &choiceExpr{
				pos: position{line: 611, col: 12, offset: 17818},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 611, col: 12, offset: 17818},
						run: (*parser).callonBaseType2,
						expr: &choiceExpr{
							pos: position{line: 611, col: 14, offset: 17820},
							alternatives: []interface{}{
								&litMatcher{
									pos:        position{line: 611, col: 14, offset: 17820},
									val:        "int",
									ignoreCase: false,
								},
								&litMatcher{
									pos:        position{line: 611, col: 22, offset: 17828},
									val:        "bool",
									ignoreCase: false,
								},
								&litMatcher{
									pos:        position{line: 611, col: 31, offset: 17837},
									val:        "string",
									ignoreCase: false,
								},
								&litMatcher{
									pos:        position{line: 611, col: 42, offset: 17848},
									val:        "byte",
									ignoreCase: false,
								},
								&litMatcher{
									pos:        position{line: 611, col: 51, offset: 17857},
									val:        "rune",
									ignoreCase: false,
								},
								&litMatcher{
									pos:        position{line: 611, col: 60, offset: 17866},
									val:        "float",
									ignoreCase: false,
								},
							},
						},
					},
					&ruleRefExpr{
						pos:  position{line: 614, col: 3, offset: 17966},
						name: "Unit",
					},
				},
			},
		},
		{
			name: "EscapedChar",
			pos:  position{line: 616, col: 1, offset: 17972},
			expr: &charClassMatcher{
				pos:        position{line: 616, col: 15, offset: 17986},
				val:        "[\\x00-\\x1f\"\\\\]",
				chars:      []rune{'"', '\\'},
				ranges:     []rune{'\x00', '\x1f'},
				ignoreCase: false,
				inverted:   false,
			},
		},
		{
			name: "EscapeSequence",
			pos:  position{line: 618, col: 1, offset: 18002},
			expr: &choiceExpr{
				pos: position{line: 618, col: 18, offset: 18019},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 618, col: 18, offset: 18019},
						name: "SingleCharEscape",
					},
					&ruleRefExpr{
						pos:  position{line: 618, col: 37, offset: 18038},
						name: "UnicodeEscape",
					},
				},
			},
		},
		{
			name: "SingleCharEscape",
			pos:  position{line: 620, col: 1, offset: 18053},
			expr: &charClassMatcher{
				pos:        position{line: 620, col: 20, offset: 18072},
				val:        "[\"\\\\/bfnrt]",
				chars:      []rune{'"', '\\', '/', 'b', 'f', 'n', 'r', 't'},
				ignoreCase: false,
				inverted:   false,
			},
		},
		{
			name: "DecimalDigit",
			pos:  position{line: 622, col: 1, offset: 18085},
			expr: &charClassMatcher{
				pos:        position{line: 622, col: 16, offset: 18100},
				val:        "[0-9]",
				ranges:     []rune{'0', '9'},
				ignoreCase: false,
				inverted:   false,
			},
		},
		{
			name: "NonZeroDecimalDigit",
			pos:  position{line: 624, col: 1, offset: 18107},
			expr: &charClassMatcher{
				pos:        position{line: 624, col: 23, offset: 18129},
				val:        "[1-9]",
				ranges:     []rune{'1', '9'},
				ignoreCase: false,
				inverted:   false,
			},
		},
		{
			name: "HexDigit",
			pos:  position{line: 626, col: 1, offset: 18136},
			expr: &charClassMatcher{
				pos:        position{line: 626, col: 12, offset: 18147},
				val:        "[0-9a-f]i",
				ranges:     []rune{'0', '9', 'a', 'f'},
				ignoreCase: true,
				inverted:   false,
			},
		},
		{
			name:        "__",
			displayName: "\"sigwhitespace\"",
			pos:         position{line: 628, col: 1, offset: 18158},
			expr: &choiceExpr{
				pos: position{line: 628, col: 22, offset: 18179},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 628, col: 22, offset: 18179},
						name: "Comment",
					},
					&oneOrMoreExpr{
						pos: position{line: 628, col: 32, offset: 18189},
						expr: &charClassMatcher{
							pos:        position{line: 628, col: 32, offset: 18189},
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
			pos:         position{line: 630, col: 1, offset: 18201},
			expr: &choiceExpr{
				pos: position{line: 630, col: 18, offset: 18218},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 630, col: 18, offset: 18218},
						name: "Comment",
					},
					&zeroOrMoreExpr{
						pos: position{line: 630, col: 28, offset: 18228},
						expr: &charClassMatcher{
							pos:        position{line: 630, col: 28, offset: 18228},
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
			pos:  position{line: 632, col: 1, offset: 18240},
			expr: &seqExpr{
				pos: position{line: 632, col: 11, offset: 18250},
				exprs: []interface{}{
					&zeroOrMoreExpr{
						pos: position{line: 632, col: 11, offset: 18250},
						expr: &charClassMatcher{
							pos:        position{line: 632, col: 11, offset: 18250},
							val:        "[ \\r\\n\\t]",
							chars:      []rune{' ', '\r', '\n', '\t'},
							ignoreCase: false,
							inverted:   false,
						},
					},
					&litMatcher{
						pos:        position{line: 632, col: 22, offset: 18261},
						val:        "#",
						ignoreCase: false,
					},
					&zeroOrMoreExpr{
						pos: position{line: 632, col: 26, offset: 18265},
						expr: &seqExpr{
							pos: position{line: 632, col: 27, offset: 18266},
							exprs: []interface{}{
								&notExpr{
									pos: position{line: 632, col: 27, offset: 18266},
									expr: &charClassMatcher{
										pos:        position{line: 632, col: 28, offset: 18267},
										val:        "[\\n]",
										chars:      []rune{'\n'},
										ignoreCase: false,
										inverted:   false,
									},
								},
								&anyMatcher{
									line: 632, col: 33, offset: 18272,
								},
							},
						},
					},
					&andExpr{
						pos: position{line: 632, col: 37, offset: 18276},
						expr: &litMatcher{
							pos:        position{line: 632, col: 38, offset: 18277},
							val:        "\n",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "EOF",
			pos:  position{line: 634, col: 1, offset: 18283},
			expr: &notExpr{
				pos: position{line: 634, col: 7, offset: 18289},
				expr: &anyMatcher{
					line: 634, col: 8, offset: 18290,
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

	return Call{Module: module.(Identifier), Function: fn.(Identifier), Arguments: arguments}, nil
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

	return Call{Module: Identifier{}, Function: fn.(Identifier), Arguments: arguments}, nil
}

func (p *parser) callonCall12() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onCall12(stack["fn"], stack["args"])
}

func (c *current) onCall19(fn interface{}) (interface{}, error) {
	//fmt.Println("call", string(c.text))
	arguments := []Ast{}

	return Call{Module: Identifier{}, Function: fn.(Identifier), Arguments: arguments}, nil
}

func (p *parser) callonCall19() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onCall19(stack["fn"])
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
	fmt.Println("parsing arg:", string(c.text))
	arg := name.(Identifier)

	if anno != nil {
		vals := anno.([]interface{})
		fmt.Println(vals)
		//restSl := toIfaceSlice(vals[0])

		switch vals[2].(type) {
		case BasicAst:
			arg.Annotation = vals[2].(BasicAst).StringValue
		case Identifier:
			arg.Annotation = vals[2].(Identifier).StringValue
		}
	}
	fmt.Println("parsed:", arg)
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
