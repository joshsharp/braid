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
			pos:  position{line: 43, col: 1, offset: 1072},
			expr: &actionExpr{
				pos: position{line: 43, col: 14, offset: 1085},
				run: (*parser).callonExternDefn1,
				expr: &seqExpr{
					pos: position{line: 43, col: 14, offset: 1085},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 43, col: 14, offset: 1085},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 43, col: 16, offset: 1087},
							val:        "extern",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 43, col: 25, offset: 1096},
							name: "__",
						},
						&labeledExpr{
							pos:   position{line: 43, col: 28, offset: 1099},
							label: "name",
							expr: &ruleRefExpr{
								pos:  position{line: 43, col: 33, offset: 1104},
								name: "Identifier",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 43, col: 44, offset: 1115},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 43, col: 46, offset: 1117},
							val:        "=",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 43, col: 50, offset: 1121},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 43, col: 52, offset: 1123},
							label: "importName",
							expr: &ruleRefExpr{
								pos:  position{line: 43, col: 63, offset: 1134},
								name: "StringLiteral",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 43, col: 77, offset: 1148},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 43, col: 79, offset: 1150},
							label: "args",
							expr: &ruleRefExpr{
								pos:  position{line: 43, col: 84, offset: 1155},
								name: "ArgsDefn",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 43, col: 93, offset: 1164},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 43, col: 95, offset: 1166},
							label: "ret",
							expr: &ruleRefExpr{
								pos:  position{line: 43, col: 99, offset: 1170},
								name: "BaseType",
							},
						},
					},
				},
			},
		},
		{
			name: "TypeDefn",
			pos:  position{line: 49, col: 1, offset: 1382},
			expr: &choiceExpr{
				pos: position{line: 49, col: 12, offset: 1393},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 49, col: 12, offset: 1393},
						run: (*parser).callonTypeDefn2,
						expr: &seqExpr{
							pos: position{line: 49, col: 12, offset: 1393},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 49, col: 12, offset: 1393},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 49, col: 14, offset: 1395},
									val:        "type",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 49, col: 21, offset: 1402},
									name: "__",
								},
								&labeledExpr{
									pos:   position{line: 49, col: 24, offset: 1405},
									label: "name",
									expr: &ruleRefExpr{
										pos:  position{line: 49, col: 29, offset: 1410},
										name: "ModuleName",
									},
								},
								&labeledExpr{
									pos:   position{line: 49, col: 40, offset: 1421},
									label: "params",
									expr: &zeroOrMoreExpr{
										pos: position{line: 49, col: 47, offset: 1428},
										expr: &seqExpr{
											pos: position{line: 49, col: 48, offset: 1429},
											exprs: []interface{}{
												&ruleRefExpr{
													pos:  position{line: 49, col: 48, offset: 1429},
													name: "__",
												},
												&ruleRefExpr{
													pos:  position{line: 49, col: 51, offset: 1432},
													name: "TypeParameter",
												},
											},
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 49, col: 67, offset: 1448},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 49, col: 69, offset: 1450},
									val:        "=",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 49, col: 73, offset: 1454},
									label: "types",
									expr: &oneOrMoreExpr{
										pos: position{line: 49, col: 79, offset: 1460},
										expr: &seqExpr{
											pos: position{line: 49, col: 80, offset: 1461},
											exprs: []interface{}{
												&ruleRefExpr{
													pos:  position{line: 49, col: 80, offset: 1461},
													name: "__",
												},
												&ruleRefExpr{
													pos:  position{line: 49, col: 83, offset: 1464},
													name: "AnyType",
												},
											},
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 49, col: 93, offset: 1474},
									name: "__",
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 68, col: 1, offset: 1970},
						run: (*parser).callonTypeDefn22,
						expr: &seqExpr{
							pos: position{line: 68, col: 1, offset: 1970},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 68, col: 1, offset: 1970},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 68, col: 3, offset: 1972},
									val:        "type",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 68, col: 10, offset: 1979},
									name: "__",
								},
								&labeledExpr{
									pos:   position{line: 68, col: 13, offset: 1982},
									label: "name",
									expr: &ruleRefExpr{
										pos:  position{line: 68, col: 18, offset: 1987},
										name: "ModuleName",
									},
								},
								&labeledExpr{
									pos:   position{line: 68, col: 29, offset: 1998},
									label: "params",
									expr: &zeroOrMoreExpr{
										pos: position{line: 68, col: 36, offset: 2005},
										expr: &seqExpr{
											pos: position{line: 68, col: 37, offset: 2006},
											exprs: []interface{}{
												&ruleRefExpr{
													pos:  position{line: 68, col: 37, offset: 2006},
													name: "__",
												},
												&ruleRefExpr{
													pos:  position{line: 68, col: 40, offset: 2009},
													name: "TypeParameter",
												},
											},
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 68, col: 56, offset: 2025},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 68, col: 58, offset: 2027},
									val:        "=",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 68, col: 62, offset: 2031},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 69, col: 5, offset: 2037},
									val:        "{",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 69, col: 9, offset: 2041},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 69, col: 11, offset: 2043},
									label: "first",
									expr: &ruleRefExpr{
										pos:  position{line: 69, col: 17, offset: 2049},
										name: "RecordFieldDefn",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 69, col: 33, offset: 2065},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 69, col: 35, offset: 2067},
									label: "rest",
									expr: &zeroOrMoreExpr{
										pos: position{line: 69, col: 40, offset: 2072},
										expr: &seqExpr{
											pos: position{line: 69, col: 41, offset: 2073},
											exprs: []interface{}{
												&litMatcher{
													pos:        position{line: 69, col: 41, offset: 2073},
													val:        ",",
													ignoreCase: false,
												},
												&ruleRefExpr{
													pos:  position{line: 69, col: 45, offset: 2077},
													name: "_",
												},
												&ruleRefExpr{
													pos:  position{line: 69, col: 47, offset: 2079},
													name: "RecordFieldDefn",
												},
												&ruleRefExpr{
													pos:  position{line: 69, col: 63, offset: 2095},
													name: "_",
												},
											},
										},
									},
								},
								&zeroOrOneExpr{
									pos: position{line: 69, col: 67, offset: 2099},
									expr: &litMatcher{
										pos:        position{line: 69, col: 67, offset: 2099},
										val:        ",",
										ignoreCase: false,
									},
								},
								&ruleRefExpr{
									pos:  position{line: 69, col: 72, offset: 2104},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 69, col: 74, offset: 2106},
									val:        "}",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 69, col: 78, offset: 2110},
									name: "__",
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 87, col: 1, offset: 2595},
						run: (*parser).callonTypeDefn54,
						expr: &seqExpr{
							pos: position{line: 87, col: 1, offset: 2595},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 87, col: 1, offset: 2595},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 87, col: 3, offset: 2597},
									val:        "type",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 87, col: 10, offset: 2604},
									name: "__",
								},
								&labeledExpr{
									pos:   position{line: 87, col: 13, offset: 2607},
									label: "name",
									expr: &ruleRefExpr{
										pos:  position{line: 87, col: 18, offset: 2612},
										name: "ModuleName",
									},
								},
								&labeledExpr{
									pos:   position{line: 87, col: 29, offset: 2623},
									label: "params",
									expr: &zeroOrMoreExpr{
										pos: position{line: 87, col: 36, offset: 2630},
										expr: &seqExpr{
											pos: position{line: 87, col: 37, offset: 2631},
											exprs: []interface{}{
												&ruleRefExpr{
													pos:  position{line: 87, col: 37, offset: 2631},
													name: "__",
												},
												&ruleRefExpr{
													pos:  position{line: 87, col: 40, offset: 2634},
													name: "TypeParameter",
												},
											},
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 87, col: 56, offset: 2650},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 87, col: 58, offset: 2652},
									val:        "=",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 87, col: 62, offset: 2656},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 87, col: 64, offset: 2658},
									label: "rest",
									expr: &oneOrMoreExpr{
										pos: position{line: 87, col: 69, offset: 2663},
										expr: &ruleRefExpr{
											pos:  position{line: 87, col: 70, offset: 2664},
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
			pos:  position{line: 102, col: 1, offset: 3071},
			expr: &actionExpr{
				pos: position{line: 102, col: 19, offset: 3089},
				run: (*parser).callonRecordFieldDefn1,
				expr: &seqExpr{
					pos: position{line: 102, col: 19, offset: 3089},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 102, col: 19, offset: 3089},
							label: "name",
							expr: &ruleRefExpr{
								pos:  position{line: 102, col: 24, offset: 3094},
								name: "VariableName",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 102, col: 37, offset: 3107},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 102, col: 39, offset: 3109},
							val:        ":",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 102, col: 43, offset: 3113},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 102, col: 45, offset: 3115},
							label: "t",
							expr: &ruleRefExpr{
								pos:  position{line: 102, col: 48, offset: 3118},
								name: "AnyType",
							},
						},
					},
				},
			},
		},
		{
			name: "VariantConstructor",
			pos:  position{line: 106, col: 1, offset: 3212},
			expr: &choiceExpr{
				pos: position{line: 106, col: 22, offset: 3233},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 106, col: 22, offset: 3233},
						run: (*parser).callonVariantConstructor2,
						expr: &seqExpr{
							pos: position{line: 106, col: 22, offset: 3233},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 106, col: 22, offset: 3233},
									val:        "|",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 106, col: 26, offset: 3237},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 106, col: 28, offset: 3239},
									label: "name",
									expr: &ruleRefExpr{
										pos:  position{line: 106, col: 33, offset: 3244},
										name: "ModuleName",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 106, col: 44, offset: 3255},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 106, col: 46, offset: 3257},
									val:        "{",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 106, col: 50, offset: 3261},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 106, col: 52, offset: 3263},
									label: "first",
									expr: &ruleRefExpr{
										pos:  position{line: 106, col: 58, offset: 3269},
										name: "RecordFieldDefn",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 106, col: 74, offset: 3285},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 106, col: 76, offset: 3287},
									label: "rest",
									expr: &zeroOrMoreExpr{
										pos: position{line: 106, col: 81, offset: 3292},
										expr: &seqExpr{
											pos: position{line: 106, col: 82, offset: 3293},
											exprs: []interface{}{
												&litMatcher{
													pos:        position{line: 106, col: 82, offset: 3293},
													val:        ",",
													ignoreCase: false,
												},
												&ruleRefExpr{
													pos:  position{line: 106, col: 86, offset: 3297},
													name: "_",
												},
												&ruleRefExpr{
													pos:  position{line: 106, col: 88, offset: 3299},
													name: "RecordFieldDefn",
												},
												&ruleRefExpr{
													pos:  position{line: 106, col: 104, offset: 3315},
													name: "_",
												},
											},
										},
									},
								},
								&zeroOrOneExpr{
									pos: position{line: 106, col: 108, offset: 3319},
									expr: &litMatcher{
										pos:        position{line: 106, col: 108, offset: 3319},
										val:        ",",
										ignoreCase: false,
									},
								},
								&ruleRefExpr{
									pos:  position{line: 106, col: 113, offset: 3324},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 106, col: 115, offset: 3326},
									val:        "}",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 106, col: 119, offset: 3330},
									name: "__",
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 125, col: 1, offset: 3935},
						run: (*parser).callonVariantConstructor26,
						expr: &seqExpr{
							pos: position{line: 125, col: 1, offset: 3935},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 125, col: 1, offset: 3935},
									val:        "|",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 125, col: 5, offset: 3939},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 125, col: 7, offset: 3941},
									label: "name",
									expr: &ruleRefExpr{
										pos:  position{line: 125, col: 12, offset: 3946},
										name: "ModuleName",
									},
								},
								&labeledExpr{
									pos:   position{line: 125, col: 23, offset: 3957},
									label: "rest",
									expr: &zeroOrMoreExpr{
										pos: position{line: 125, col: 28, offset: 3962},
										expr: &seqExpr{
											pos: position{line: 125, col: 29, offset: 3963},
											exprs: []interface{}{
												&ruleRefExpr{
													pos:  position{line: 125, col: 29, offset: 3963},
													name: "__",
												},
												&ruleRefExpr{
													pos:  position{line: 125, col: 32, offset: 3966},
													name: "AnyType",
												},
											},
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 125, col: 42, offset: 3976},
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
			pos:  position{line: 142, col: 1, offset: 4413},
			expr: &choiceExpr{
				pos: position{line: 142, col: 11, offset: 4423},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 142, col: 11, offset: 4423},
						name: "BaseType",
					},
					&ruleRefExpr{
						pos:  position{line: 142, col: 22, offset: 4434},
						name: "TypeParameter",
					},
				},
			},
		},
		{
			name: "Assignment",
			pos:  position{line: 144, col: 1, offset: 4449},
			expr: &choiceExpr{
				pos: position{line: 144, col: 14, offset: 4462},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 144, col: 14, offset: 4462},
						run: (*parser).callonAssignment2,
						expr: &seqExpr{
							pos: position{line: 144, col: 14, offset: 4462},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 144, col: 14, offset: 4462},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 144, col: 16, offset: 4464},
									val:        "let",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 144, col: 22, offset: 4470},
									name: "__",
								},
								&labeledExpr{
									pos:   position{line: 144, col: 25, offset: 4473},
									label: "i",
									expr: &ruleRefExpr{
										pos:  position{line: 144, col: 27, offset: 4475},
										name: "Assignable",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 144, col: 38, offset: 4486},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 144, col: 40, offset: 4488},
									val:        "=",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 144, col: 44, offset: 4492},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 144, col: 46, offset: 4494},
									label: "expr",
									expr: &ruleRefExpr{
										pos:  position{line: 144, col: 51, offset: 4499},
										name: "Expr",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 144, col: 56, offset: 4504},
									name: "_",
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 150, col: 1, offset: 4623},
						run: (*parser).callonAssignment15,
						expr: &seqExpr{
							pos: position{line: 150, col: 1, offset: 4623},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 150, col: 1, offset: 4623},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 150, col: 3, offset: 4625},
									val:        "let",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 150, col: 9, offset: 4631},
									name: "__",
								},
								&notExpr{
									pos: position{line: 150, col: 12, offset: 4634},
									expr: &ruleRefExpr{
										pos:  position{line: 150, col: 13, offset: 4635},
										name: "Assignable",
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 154, col: 1, offset: 4733},
						run: (*parser).callonAssignment22,
						expr: &seqExpr{
							pos: position{line: 154, col: 1, offset: 4733},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 154, col: 1, offset: 4733},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 154, col: 3, offset: 4735},
									val:        "let",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 154, col: 9, offset: 4741},
									name: "__",
								},
								&labeledExpr{
									pos:   position{line: 154, col: 12, offset: 4744},
									label: "i",
									expr: &ruleRefExpr{
										pos:  position{line: 154, col: 14, offset: 4746},
										name: "Assignable",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 154, col: 25, offset: 4757},
									name: "_",
								},
								&notExpr{
									pos: position{line: 154, col: 27, offset: 4759},
									expr: &litMatcher{
										pos:        position{line: 154, col: 28, offset: 4760},
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
			pos:  position{line: 158, col: 1, offset: 4854},
			expr: &actionExpr{
				pos: position{line: 158, col: 12, offset: 4865},
				run: (*parser).callonFuncDefn1,
				expr: &seqExpr{
					pos: position{line: 158, col: 12, offset: 4865},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 158, col: 12, offset: 4865},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 158, col: 14, offset: 4867},
							val:        "let",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 158, col: 20, offset: 4873},
							name: "__",
						},
						&labeledExpr{
							pos:   position{line: 158, col: 23, offset: 4876},
							label: "i",
							expr: &ruleRefExpr{
								pos:  position{line: 158, col: 25, offset: 4878},
								name: "VariableName",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 158, col: 38, offset: 4891},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 158, col: 40, offset: 4893},
							val:        "=",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 158, col: 44, offset: 4897},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 158, col: 46, offset: 4899},
							label: "ids",
							expr: &zeroOrOneExpr{
								pos: position{line: 158, col: 50, offset: 4903},
								expr: &seqExpr{
									pos: position{line: 158, col: 51, offset: 4904},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 158, col: 51, offset: 4904},
											name: "ArgsDefn",
										},
										&ruleRefExpr{
											pos:  position{line: 158, col: 60, offset: 4913},
											name: "_",
										},
									},
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 158, col: 64, offset: 4917},
							label: "ret",
							expr: &zeroOrOneExpr{
								pos: position{line: 158, col: 68, offset: 4921},
								expr: &seqExpr{
									pos: position{line: 158, col: 69, offset: 4922},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 158, col: 69, offset: 4922},
											name: "AnyType",
										},
										&ruleRefExpr{
											pos:  position{line: 158, col: 77, offset: 4930},
											name: "_",
										},
									},
								},
							},
						},
						&litMatcher{
							pos:        position{line: 158, col: 81, offset: 4934},
							val:        "{",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 158, col: 85, offset: 4938},
							name: "__",
						},
						&labeledExpr{
							pos:   position{line: 158, col: 88, offset: 4941},
							label: "statements",
							expr: &oneOrMoreExpr{
								pos: position{line: 158, col: 99, offset: 4952},
								expr: &ruleRefExpr{
									pos:  position{line: 158, col: 100, offset: 4953},
									name: "Statement",
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 158, col: 112, offset: 4965},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 158, col: 114, offset: 4967},
							val:        "}",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 158, col: 118, offset: 4971},
							name: "_",
						},
					},
				},
			},
		},
		{
			name: "Expr",
			pos:  position{line: 184, col: 1, offset: 5581},
			expr: &actionExpr{
				pos: position{line: 184, col: 8, offset: 5588},
				run: (*parser).callonExpr1,
				expr: &labeledExpr{
					pos:   position{line: 184, col: 8, offset: 5588},
					label: "ex",
					expr: &choiceExpr{
						pos: position{line: 184, col: 12, offset: 5592},
						alternatives: []interface{}{
							&ruleRefExpr{
								pos:  position{line: 184, col: 12, offset: 5592},
								name: "IfExpr",
							},
							&ruleRefExpr{
								pos:  position{line: 184, col: 21, offset: 5601},
								name: "BinOp",
							},
							&ruleRefExpr{
								pos:  position{line: 184, col: 29, offset: 5609},
								name: "Call",
							},
						},
					},
				},
			},
		},
		{
			name: "IfExpr",
			pos:  position{line: 190, col: 1, offset: 5718},
			expr: &choiceExpr{
				pos: position{line: 190, col: 10, offset: 5727},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 190, col: 10, offset: 5727},
						run: (*parser).callonIfExpr2,
						expr: &seqExpr{
							pos: position{line: 190, col: 10, offset: 5727},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 190, col: 10, offset: 5727},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 190, col: 12, offset: 5729},
									val:        "if",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 190, col: 17, offset: 5734},
									name: "__",
								},
								&labeledExpr{
									pos:   position{line: 190, col: 20, offset: 5737},
									label: "expr",
									expr: &ruleRefExpr{
										pos:  position{line: 190, col: 25, offset: 5742},
										name: "BinOpBool",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 190, col: 35, offset: 5752},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 190, col: 37, offset: 5754},
									val:        "{",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 190, col: 41, offset: 5758},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 190, col: 43, offset: 5760},
									label: "thens",
									expr: &oneOrMoreExpr{
										pos: position{line: 190, col: 49, offset: 5766},
										expr: &ruleRefExpr{
											pos:  position{line: 190, col: 50, offset: 5767},
											name: "Statement",
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 190, col: 62, offset: 5779},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 190, col: 64, offset: 5781},
									val:        "}",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 190, col: 68, offset: 5785},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 190, col: 70, offset: 5787},
									val:        "else",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 190, col: 77, offset: 5794},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 190, col: 79, offset: 5796},
									label: "elseifs",
									expr: &ruleRefExpr{
										pos:  position{line: 190, col: 87, offset: 5804},
										name: "IfExpr",
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 202, col: 1, offset: 6134},
						run: (*parser).callonIfExpr22,
						expr: &seqExpr{
							pos: position{line: 202, col: 1, offset: 6134},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 202, col: 1, offset: 6134},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 202, col: 3, offset: 6136},
									val:        "if",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 202, col: 8, offset: 6141},
									name: "__",
								},
								&labeledExpr{
									pos:   position{line: 202, col: 11, offset: 6144},
									label: "expr",
									expr: &ruleRefExpr{
										pos:  position{line: 202, col: 16, offset: 6149},
										name: "BinOpBool",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 202, col: 26, offset: 6159},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 202, col: 28, offset: 6161},
									val:        "{",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 202, col: 32, offset: 6165},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 202, col: 34, offset: 6167},
									label: "thens",
									expr: &oneOrMoreExpr{
										pos: position{line: 202, col: 40, offset: 6173},
										expr: &ruleRefExpr{
											pos:  position{line: 202, col: 41, offset: 6174},
											name: "Statement",
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 202, col: 53, offset: 6186},
									name: "__",
								},
								&litMatcher{
									pos:        position{line: 202, col: 56, offset: 6189},
									val:        "}",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 202, col: 60, offset: 6193},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 202, col: 62, offset: 6195},
									val:        "else",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 202, col: 69, offset: 6202},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 202, col: 71, offset: 6204},
									val:        "{",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 202, col: 75, offset: 6208},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 202, col: 77, offset: 6210},
									label: "elses",
									expr: &oneOrMoreExpr{
										pos: position{line: 202, col: 83, offset: 6216},
										expr: &ruleRefExpr{
											pos:  position{line: 202, col: 84, offset: 6217},
											name: "Statement",
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 202, col: 96, offset: 6229},
									name: "__",
								},
								&litMatcher{
									pos:        position{line: 202, col: 99, offset: 6232},
									val:        "}",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 221, col: 1, offset: 6735},
						run: (*parser).callonIfExpr47,
						expr: &seqExpr{
							pos: position{line: 221, col: 1, offset: 6735},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 221, col: 1, offset: 6735},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 221, col: 3, offset: 6737},
									val:        "if",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 221, col: 8, offset: 6742},
									name: "__",
								},
								&labeledExpr{
									pos:   position{line: 221, col: 11, offset: 6745},
									label: "expr",
									expr: &ruleRefExpr{
										pos:  position{line: 221, col: 16, offset: 6750},
										name: "BinOpBool",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 221, col: 26, offset: 6760},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 221, col: 28, offset: 6762},
									val:        "{",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 221, col: 32, offset: 6766},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 221, col: 34, offset: 6768},
									label: "thens",
									expr: &oneOrMoreExpr{
										pos: position{line: 221, col: 40, offset: 6774},
										expr: &ruleRefExpr{
											pos:  position{line: 221, col: 41, offset: 6775},
											name: "Statement",
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 221, col: 53, offset: 6787},
									name: "__",
								},
								&litMatcher{
									pos:        position{line: 221, col: 56, offset: 6790},
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
			pos:  position{line: 233, col: 1, offset: 7088},
			expr: &choiceExpr{
				pos: position{line: 233, col: 8, offset: 7095},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 233, col: 8, offset: 7095},
						run: (*parser).callonCall2,
						expr: &seqExpr{
							pos: position{line: 233, col: 8, offset: 7095},
							exprs: []interface{}{
								&labeledExpr{
									pos:   position{line: 233, col: 8, offset: 7095},
									label: "module",
									expr: &ruleRefExpr{
										pos:  position{line: 233, col: 15, offset: 7102},
										name: "ModuleName",
									},
								},
								&litMatcher{
									pos:        position{line: 233, col: 26, offset: 7113},
									val:        ".",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 233, col: 30, offset: 7117},
									label: "fn",
									expr: &ruleRefExpr{
										pos:  position{line: 233, col: 33, offset: 7120},
										name: "VariableName",
									},
								},
								&labeledExpr{
									pos:   position{line: 233, col: 46, offset: 7133},
									label: "args",
									expr: &ruleRefExpr{
										pos:  position{line: 233, col: 51, offset: 7138},
										name: "Arguments",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 233, col: 61, offset: 7148},
									name: "_",
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 247, col: 1, offset: 7464},
						run: (*parser).callonCall12,
						expr: &seqExpr{
							pos: position{line: 247, col: 1, offset: 7464},
							exprs: []interface{}{
								&labeledExpr{
									pos:   position{line: 247, col: 1, offset: 7464},
									label: "fn",
									expr: &ruleRefExpr{
										pos:  position{line: 247, col: 4, offset: 7467},
										name: "VariableName",
									},
								},
								&labeledExpr{
									pos:   position{line: 247, col: 17, offset: 7480},
									label: "args",
									expr: &ruleRefExpr{
										pos:  position{line: 247, col: 22, offset: 7485},
										name: "Arguments",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 247, col: 32, offset: 7495},
									name: "_",
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 261, col: 1, offset: 7804},
						run: (*parser).callonCall19,
						expr: &seqExpr{
							pos: position{line: 261, col: 1, offset: 7804},
							exprs: []interface{}{
								&labeledExpr{
									pos:   position{line: 261, col: 1, offset: 7804},
									label: "fn",
									expr: &ruleRefExpr{
										pos:  position{line: 261, col: 4, offset: 7807},
										name: "VariableName",
									},
								},
								&litMatcher{
									pos:        position{line: 261, col: 17, offset: 7820},
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
			pos:  position{line: 268, col: 1, offset: 7991},
			expr: &actionExpr{
				pos: position{line: 268, col: 12, offset: 8002},
				run: (*parser).callonArgsDefn1,
				expr: &seqExpr{
					pos: position{line: 268, col: 12, offset: 8002},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 268, col: 12, offset: 8002},
							val:        "(",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 268, col: 16, offset: 8006},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 268, col: 18, offset: 8008},
							label: "argument",
							expr: &ruleRefExpr{
								pos:  position{line: 268, col: 27, offset: 8017},
								name: "ArgDefn",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 268, col: 35, offset: 8025},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 268, col: 37, offset: 8027},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 268, col: 42, offset: 8032},
								expr: &seqExpr{
									pos: position{line: 268, col: 43, offset: 8033},
									exprs: []interface{}{
										&litMatcher{
											pos:        position{line: 268, col: 43, offset: 8033},
											val:        ",",
											ignoreCase: false,
										},
										&ruleRefExpr{
											pos:  position{line: 268, col: 47, offset: 8037},
											name: "_",
										},
										&ruleRefExpr{
											pos:  position{line: 268, col: 49, offset: 8039},
											name: "ArgDefn",
										},
									},
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 268, col: 59, offset: 8049},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 268, col: 61, offset: 8051},
							val:        ")",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "ArgDefn",
			pos:  position{line: 286, col: 1, offset: 8473},
			expr: &actionExpr{
				pos: position{line: 286, col: 11, offset: 8483},
				run: (*parser).callonArgDefn1,
				expr: &seqExpr{
					pos: position{line: 286, col: 11, offset: 8483},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 286, col: 11, offset: 8483},
							label: "name",
							expr: &ruleRefExpr{
								pos:  position{line: 286, col: 16, offset: 8488},
								name: "Identifier",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 286, col: 27, offset: 8499},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 286, col: 29, offset: 8501},
							label: "anno",
							expr: &zeroOrOneExpr{
								pos: position{line: 286, col: 34, offset: 8506},
								expr: &seqExpr{
									pos: position{line: 286, col: 35, offset: 8507},
									exprs: []interface{}{
										&litMatcher{
											pos:        position{line: 286, col: 35, offset: 8507},
											val:        ":",
											ignoreCase: false,
										},
										&ruleRefExpr{
											pos:  position{line: 286, col: 39, offset: 8511},
											name: "_",
										},
										&ruleRefExpr{
											pos:  position{line: 286, col: 41, offset: 8513},
											name: "AnyType",
										},
									},
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 286, col: 52, offset: 8524},
							name: "_",
						},
					},
				},
			},
		},
		{
			name: "Arguments",
			pos:  position{line: 306, col: 1, offset: 9012},
			expr: &choiceExpr{
				pos: position{line: 306, col: 13, offset: 9024},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 306, col: 13, offset: 9024},
						run: (*parser).callonArguments2,
						expr: &seqExpr{
							pos: position{line: 306, col: 13, offset: 9024},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 306, col: 13, offset: 9024},
									val:        "(",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 306, col: 17, offset: 9028},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 306, col: 19, offset: 9030},
									label: "argument",
									expr: &ruleRefExpr{
										pos:  position{line: 306, col: 28, offset: 9039},
										name: "BinOpParens",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 306, col: 40, offset: 9051},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 306, col: 42, offset: 9053},
									label: "rest",
									expr: &zeroOrMoreExpr{
										pos: position{line: 306, col: 47, offset: 9058},
										expr: &seqExpr{
											pos: position{line: 306, col: 48, offset: 9059},
											exprs: []interface{}{
												&litMatcher{
													pos:        position{line: 306, col: 48, offset: 9059},
													val:        ",",
													ignoreCase: false,
												},
												&ruleRefExpr{
													pos:  position{line: 306, col: 52, offset: 9063},
													name: "_",
												},
												&ruleRefExpr{
													pos:  position{line: 306, col: 54, offset: 9065},
													name: "BinOpParens",
												},
											},
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 306, col: 68, offset: 9079},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 306, col: 70, offset: 9081},
									val:        ")",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 323, col: 1, offset: 9503},
						run: (*parser).callonArguments17,
						expr: &seqExpr{
							pos: position{line: 323, col: 1, offset: 9503},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 323, col: 1, offset: 9503},
									val:        "(",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 323, col: 5, offset: 9507},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 323, col: 7, offset: 9509},
									label: "argument",
									expr: &ruleRefExpr{
										pos:  position{line: 323, col: 16, offset: 9518},
										name: "Expr",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 323, col: 21, offset: 9523},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 323, col: 23, offset: 9525},
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
			pos:  position{line: 328, col: 1, offset: 9630},
			expr: &actionExpr{
				pos: position{line: 328, col: 16, offset: 9645},
				run: (*parser).callonCompoundExpr1,
				expr: &seqExpr{
					pos: position{line: 328, col: 16, offset: 9645},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 328, col: 16, offset: 9645},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 328, col: 18, offset: 9647},
							label: "op",
							expr: &ruleRefExpr{
								pos:  position{line: 328, col: 21, offset: 9650},
								name: "BinOp",
							},
						},
						&labeledExpr{
							pos:   position{line: 328, col: 27, offset: 9656},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 328, col: 32, offset: 9661},
								expr: &seqExpr{
									pos: position{line: 328, col: 33, offset: 9662},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 328, col: 33, offset: 9662},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 328, col: 36, offset: 9665},
											name: "Operator",
										},
										&ruleRefExpr{
											pos:  position{line: 328, col: 45, offset: 9674},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 328, col: 48, offset: 9677},
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
			pos:  position{line: 348, col: 1, offset: 10283},
			expr: &choiceExpr{
				pos: position{line: 348, col: 9, offset: 10291},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 348, col: 9, offset: 10291},
						name: "BinOpBool",
					},
					&ruleRefExpr{
						pos:  position{line: 348, col: 21, offset: 10303},
						name: "BinOpEquality",
					},
					&ruleRefExpr{
						pos:  position{line: 348, col: 37, offset: 10319},
						name: "BinOpLow",
					},
					&ruleRefExpr{
						pos:  position{line: 348, col: 48, offset: 10330},
						name: "BinOpHigh",
					},
					&ruleRefExpr{
						pos:  position{line: 348, col: 60, offset: 10342},
						name: "BinOpParens",
					},
				},
			},
		},
		{
			name: "BinOpBool",
			pos:  position{line: 350, col: 1, offset: 10355},
			expr: &actionExpr{
				pos: position{line: 350, col: 13, offset: 10367},
				run: (*parser).callonBinOpBool1,
				expr: &seqExpr{
					pos: position{line: 350, col: 13, offset: 10367},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 350, col: 13, offset: 10367},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 350, col: 15, offset: 10369},
							label: "first",
							expr: &ruleRefExpr{
								pos:  position{line: 350, col: 21, offset: 10375},
								name: "BinOpEquality",
							},
						},
						&labeledExpr{
							pos:   position{line: 350, col: 35, offset: 10389},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 350, col: 40, offset: 10394},
								expr: &seqExpr{
									pos: position{line: 350, col: 41, offset: 10395},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 350, col: 41, offset: 10395},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 350, col: 44, offset: 10398},
											name: "OperatorBoolean",
										},
										&ruleRefExpr{
											pos:  position{line: 350, col: 60, offset: 10414},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 350, col: 63, offset: 10417},
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
			pos:  position{line: 383, col: 1, offset: 11310},
			expr: &actionExpr{
				pos: position{line: 383, col: 17, offset: 11326},
				run: (*parser).callonBinOpEquality1,
				expr: &seqExpr{
					pos: position{line: 383, col: 17, offset: 11326},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 383, col: 17, offset: 11326},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 383, col: 19, offset: 11328},
							label: "first",
							expr: &ruleRefExpr{
								pos:  position{line: 383, col: 25, offset: 11334},
								name: "BinOpLow",
							},
						},
						&labeledExpr{
							pos:   position{line: 383, col: 34, offset: 11343},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 383, col: 39, offset: 11348},
								expr: &seqExpr{
									pos: position{line: 383, col: 40, offset: 11349},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 383, col: 40, offset: 11349},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 383, col: 43, offset: 11352},
											name: "OperatorEquality",
										},
										&ruleRefExpr{
											pos:  position{line: 383, col: 60, offset: 11369},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 383, col: 63, offset: 11372},
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
			pos:  position{line: 415, col: 1, offset: 12259},
			expr: &actionExpr{
				pos: position{line: 415, col: 12, offset: 12270},
				run: (*parser).callonBinOpLow1,
				expr: &seqExpr{
					pos: position{line: 415, col: 12, offset: 12270},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 415, col: 12, offset: 12270},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 415, col: 14, offset: 12272},
							label: "first",
							expr: &ruleRefExpr{
								pos:  position{line: 415, col: 20, offset: 12278},
								name: "BinOpHigh",
							},
						},
						&labeledExpr{
							pos:   position{line: 415, col: 30, offset: 12288},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 415, col: 35, offset: 12293},
								expr: &seqExpr{
									pos: position{line: 415, col: 36, offset: 12294},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 415, col: 36, offset: 12294},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 415, col: 39, offset: 12297},
											name: "OperatorLow",
										},
										&ruleRefExpr{
											pos:  position{line: 415, col: 51, offset: 12309},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 415, col: 54, offset: 12312},
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
			pos:  position{line: 447, col: 1, offset: 13200},
			expr: &actionExpr{
				pos: position{line: 447, col: 13, offset: 13212},
				run: (*parser).callonBinOpHigh1,
				expr: &seqExpr{
					pos: position{line: 447, col: 13, offset: 13212},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 447, col: 13, offset: 13212},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 447, col: 15, offset: 13214},
							label: "first",
							expr: &ruleRefExpr{
								pos:  position{line: 447, col: 21, offset: 13220},
								name: "BinOpParens",
							},
						},
						&labeledExpr{
							pos:   position{line: 447, col: 33, offset: 13232},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 447, col: 38, offset: 13237},
								expr: &seqExpr{
									pos: position{line: 447, col: 39, offset: 13238},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 447, col: 39, offset: 13238},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 447, col: 42, offset: 13241},
											name: "OperatorHigh",
										},
										&ruleRefExpr{
											pos:  position{line: 447, col: 55, offset: 13254},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 447, col: 58, offset: 13257},
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
			pos:  position{line: 478, col: 1, offset: 14146},
			expr: &choiceExpr{
				pos: position{line: 478, col: 15, offset: 14160},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 478, col: 15, offset: 14160},
						run: (*parser).callonBinOpParens2,
						expr: &seqExpr{
							pos: position{line: 478, col: 15, offset: 14160},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 478, col: 15, offset: 14160},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 478, col: 17, offset: 14162},
									val:        "(",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 478, col: 21, offset: 14166},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 478, col: 23, offset: 14168},
									label: "first",
									expr: &ruleRefExpr{
										pos:  position{line: 478, col: 29, offset: 14174},
										name: "BinOp",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 478, col: 35, offset: 14180},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 478, col: 37, offset: 14182},
									val:        ")",
									ignoreCase: false,
								},
							},
						},
					},
					&ruleRefExpr{
						pos:  position{line: 481, col: 5, offset: 14305},
						name: "Value",
					},
				},
			},
		},
		{
			name: "Operator",
			pos:  position{line: 483, col: 1, offset: 14312},
			expr: &choiceExpr{
				pos: position{line: 483, col: 12, offset: 14323},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 483, col: 12, offset: 14323},
						name: "OperatorBoolean",
					},
					&ruleRefExpr{
						pos:  position{line: 483, col: 30, offset: 14341},
						name: "OperatorEquality",
					},
					&ruleRefExpr{
						pos:  position{line: 483, col: 49, offset: 14360},
						name: "OperatorHigh",
					},
					&ruleRefExpr{
						pos:  position{line: 483, col: 64, offset: 14375},
						name: "OperatorLow",
					},
				},
			},
		},
		{
			name: "OperatorBoolean",
			pos:  position{line: 485, col: 1, offset: 14388},
			expr: &actionExpr{
				pos: position{line: 485, col: 19, offset: 14406},
				run: (*parser).callonOperatorBoolean1,
				expr: &choiceExpr{
					pos: position{line: 485, col: 21, offset: 14408},
					alternatives: []interface{}{
						&litMatcher{
							pos:        position{line: 485, col: 21, offset: 14408},
							val:        "or",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 485, col: 28, offset: 14415},
							val:        "and",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "OperatorEquality",
			pos:  position{line: 489, col: 1, offset: 14497},
			expr: &actionExpr{
				pos: position{line: 489, col: 20, offset: 14516},
				run: (*parser).callonOperatorEquality1,
				expr: &choiceExpr{
					pos: position{line: 489, col: 22, offset: 14518},
					alternatives: []interface{}{
						&litMatcher{
							pos:        position{line: 489, col: 22, offset: 14518},
							val:        ">=",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 489, col: 29, offset: 14525},
							val:        "<=",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 489, col: 36, offset: 14532},
							val:        "<",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 489, col: 42, offset: 14538},
							val:        ">",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 489, col: 48, offset: 14544},
							val:        "===",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 489, col: 56, offset: 14552},
							val:        "==",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "OperatorHigh",
			pos:  position{line: 493, col: 1, offset: 14631},
			expr: &choiceExpr{
				pos: position{line: 493, col: 16, offset: 14646},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 493, col: 16, offset: 14646},
						run: (*parser).callonOperatorHigh2,
						expr: &choiceExpr{
							pos: position{line: 493, col: 18, offset: 14648},
							alternatives: []interface{}{
								&litMatcher{
									pos:        position{line: 493, col: 18, offset: 14648},
									val:        "/",
									ignoreCase: false,
								},
								&litMatcher{
									pos:        position{line: 493, col: 24, offset: 14654},
									val:        "*",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 496, col: 3, offset: 14737},
						run: (*parser).callonOperatorHigh6,
						expr: &litMatcher{
							pos:        position{line: 496, col: 5, offset: 14739},
							val:        "^",
							ignoreCase: false,
						},
					},
					&actionExpr{
						pos: position{line: 499, col: 3, offset: 14819},
						run: (*parser).callonOperatorHigh8,
						expr: &litMatcher{
							pos:        position{line: 499, col: 3, offset: 14819},
							val:        "++",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "OperatorLow",
			pos:  position{line: 503, col: 1, offset: 14900},
			expr: &actionExpr{
				pos: position{line: 503, col: 15, offset: 14914},
				run: (*parser).callonOperatorLow1,
				expr: &choiceExpr{
					pos: position{line: 503, col: 17, offset: 14916},
					alternatives: []interface{}{
						&litMatcher{
							pos:        position{line: 503, col: 17, offset: 14916},
							val:        "+",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 503, col: 23, offset: 14922},
							val:        "-",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "Value",
			pos:  position{line: 507, col: 1, offset: 15004},
			expr: &choiceExpr{
				pos: position{line: 507, col: 9, offset: 15012},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 507, col: 9, offset: 15012},
						name: "Call",
					},
					&ruleRefExpr{
						pos:  position{line: 507, col: 16, offset: 15019},
						name: "VariableName",
					},
					&ruleRefExpr{
						pos:  position{line: 507, col: 31, offset: 15034},
						name: "Const",
					},
				},
			},
		},
		{
			name: "Assignable",
			pos:  position{line: 509, col: 1, offset: 15041},
			expr: &choiceExpr{
				pos: position{line: 509, col: 14, offset: 15054},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 509, col: 14, offset: 15054},
						name: "VariableName",
					},
					&ruleRefExpr{
						pos:  position{line: 509, col: 29, offset: 15069},
						name: "Unused",
					},
				},
			},
		},
		{
			name: "Identifier",
			pos:  position{line: 511, col: 1, offset: 15077},
			expr: &choiceExpr{
				pos: position{line: 511, col: 14, offset: 15090},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 511, col: 14, offset: 15090},
						name: "VariableName",
					},
					&ruleRefExpr{
						pos:  position{line: 511, col: 29, offset: 15105},
						name: "ModuleName",
					},
				},
			},
		},
		{
			name: "ArrayLiteral",
			pos:  position{line: 513, col: 1, offset: 15117},
			expr: &actionExpr{
				pos: position{line: 513, col: 16, offset: 15132},
				run: (*parser).callonArrayLiteral1,
				expr: &seqExpr{
					pos: position{line: 513, col: 16, offset: 15132},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 513, col: 16, offset: 15132},
							val:        "[",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 513, col: 20, offset: 15136},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 513, col: 22, offset: 15138},
							label: "first",
							expr: &ruleRefExpr{
								pos:  position{line: 513, col: 28, offset: 15144},
								name: "Expr",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 513, col: 33, offset: 15149},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 513, col: 35, offset: 15151},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 513, col: 40, offset: 15156},
								expr: &seqExpr{
									pos: position{line: 513, col: 41, offset: 15157},
									exprs: []interface{}{
										&litMatcher{
											pos:        position{line: 513, col: 41, offset: 15157},
											val:        ",",
											ignoreCase: false,
										},
										&ruleRefExpr{
											pos:  position{line: 513, col: 45, offset: 15161},
											name: "_",
										},
										&ruleRefExpr{
											pos:  position{line: 513, col: 47, offset: 15163},
											name: "Expr",
										},
										&ruleRefExpr{
											pos:  position{line: 513, col: 52, offset: 15168},
											name: "_",
										},
									},
								},
							},
						},
						&zeroOrOneExpr{
							pos: position{line: 513, col: 56, offset: 15172},
							expr: &litMatcher{
								pos:        position{line: 513, col: 56, offset: 15172},
								val:        ",",
								ignoreCase: false,
							},
						},
						&ruleRefExpr{
							pos:  position{line: 513, col: 61, offset: 15177},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 513, col: 63, offset: 15179},
							val:        "]",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "VariantInstance",
			pos:  position{line: 529, col: 1, offset: 15624},
			expr: &actionExpr{
				pos: position{line: 529, col: 19, offset: 15642},
				run: (*parser).callonVariantInstance1,
				expr: &seqExpr{
					pos: position{line: 529, col: 19, offset: 15642},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 529, col: 19, offset: 15642},
							label: "name",
							expr: &ruleRefExpr{
								pos:  position{line: 529, col: 24, offset: 15647},
								name: "ModuleName",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 529, col: 35, offset: 15658},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 529, col: 37, offset: 15660},
							label: "args",
							expr: &ruleRefExpr{
								pos:  position{line: 529, col: 42, offset: 15665},
								name: "Arguments",
							},
						},
					},
				},
			},
		},
		{
			name: "RecordInstance",
			pos:  position{line: 542, col: 1, offset: 15935},
			expr: &actionExpr{
				pos: position{line: 542, col: 18, offset: 15952},
				run: (*parser).callonRecordInstance1,
				expr: &seqExpr{
					pos: position{line: 542, col: 18, offset: 15952},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 542, col: 18, offset: 15952},
							label: "name",
							expr: &ruleRefExpr{
								pos:  position{line: 542, col: 23, offset: 15957},
								name: "ModuleName",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 542, col: 34, offset: 15968},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 542, col: 36, offset: 15970},
							val:        "{",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 542, col: 40, offset: 15974},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 542, col: 42, offset: 15976},
							label: "firstName",
							expr: &ruleRefExpr{
								pos:  position{line: 542, col: 52, offset: 15986},
								name: "VariableName",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 542, col: 65, offset: 15999},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 542, col: 67, offset: 16001},
							val:        ":",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 542, col: 71, offset: 16005},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 542, col: 73, offset: 16007},
							label: "firstValue",
							expr: &ruleRefExpr{
								pos:  position{line: 542, col: 84, offset: 16018},
								name: "Expr",
							},
						},
						&labeledExpr{
							pos:   position{line: 542, col: 89, offset: 16023},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 542, col: 94, offset: 16028},
								expr: &seqExpr{
									pos: position{line: 542, col: 95, offset: 16029},
									exprs: []interface{}{
										&litMatcher{
											pos:        position{line: 542, col: 95, offset: 16029},
											val:        ",",
											ignoreCase: false,
										},
										&ruleRefExpr{
											pos:  position{line: 542, col: 99, offset: 16033},
											name: "_",
										},
										&ruleRefExpr{
											pos:  position{line: 542, col: 101, offset: 16035},
											name: "VariableName",
										},
										&ruleRefExpr{
											pos:  position{line: 542, col: 114, offset: 16048},
											name: "_",
										},
										&litMatcher{
											pos:        position{line: 542, col: 116, offset: 16050},
											val:        ":",
											ignoreCase: false,
										},
										&ruleRefExpr{
											pos:  position{line: 542, col: 120, offset: 16054},
											name: "_",
										},
										&ruleRefExpr{
											pos:  position{line: 542, col: 122, offset: 16056},
											name: "Expr",
										},
									},
								},
							},
						},
						&litMatcher{
							pos:        position{line: 542, col: 130, offset: 16064},
							val:        "}",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "TypeParameter",
			pos:  position{line: 562, col: 1, offset: 16648},
			expr: &actionExpr{
				pos: position{line: 562, col: 17, offset: 16664},
				run: (*parser).callonTypeParameter1,
				expr: &seqExpr{
					pos: position{line: 562, col: 17, offset: 16664},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 562, col: 17, offset: 16664},
							val:        "'",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 562, col: 22, offset: 16669},
							name: "VariableName",
						},
					},
				},
			},
		},
		{
			name: "VariableName",
			pos:  position{line: 566, col: 1, offset: 16742},
			expr: &actionExpr{
				pos: position{line: 566, col: 16, offset: 16757},
				run: (*parser).callonVariableName1,
				expr: &seqExpr{
					pos: position{line: 566, col: 16, offset: 16757},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 566, col: 16, offset: 16757},
							expr: &ruleRefExpr{
								pos:  position{line: 566, col: 17, offset: 16758},
								name: "Reserved",
							},
						},
						&seqExpr{
							pos: position{line: 566, col: 27, offset: 16768},
							exprs: []interface{}{
								&oneOrMoreExpr{
									pos: position{line: 566, col: 27, offset: 16768},
									expr: &charClassMatcher{
										pos:        position{line: 566, col: 27, offset: 16768},
										val:        "[a-z]",
										ranges:     []rune{'a', 'z'},
										ignoreCase: false,
										inverted:   false,
									},
								},
								&zeroOrMoreExpr{
									pos: position{line: 566, col: 34, offset: 16775},
									expr: &charClassMatcher{
										pos:        position{line: 566, col: 34, offset: 16775},
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
			pos:  position{line: 570, col: 1, offset: 16850},
			expr: &actionExpr{
				pos: position{line: 570, col: 14, offset: 16863},
				run: (*parser).callonModuleName1,
				expr: &seqExpr{
					pos: position{line: 570, col: 15, offset: 16864},
					exprs: []interface{}{
						&oneOrMoreExpr{
							pos: position{line: 570, col: 15, offset: 16864},
							expr: &charClassMatcher{
								pos:        position{line: 570, col: 15, offset: 16864},
								val:        "[A-Z]",
								ranges:     []rune{'A', 'Z'},
								ignoreCase: false,
								inverted:   false,
							},
						},
						&zeroOrMoreExpr{
							pos: position{line: 570, col: 22, offset: 16871},
							expr: &charClassMatcher{
								pos:        position{line: 570, col: 22, offset: 16871},
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
			pos:  position{line: 574, col: 1, offset: 16946},
			expr: &choiceExpr{
				pos: position{line: 574, col: 9, offset: 16954},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 574, col: 9, offset: 16954},
						run: (*parser).callonConst2,
						expr: &seqExpr{
							pos: position{line: 574, col: 9, offset: 16954},
							exprs: []interface{}{
								&zeroOrOneExpr{
									pos: position{line: 574, col: 9, offset: 16954},
									expr: &litMatcher{
										pos:        position{line: 574, col: 9, offset: 16954},
										val:        "-",
										ignoreCase: false,
									},
								},
								&oneOrMoreExpr{
									pos: position{line: 574, col: 14, offset: 16959},
									expr: &charClassMatcher{
										pos:        position{line: 574, col: 14, offset: 16959},
										val:        "[0-9]",
										ranges:     []rune{'0', '9'},
										ignoreCase: false,
										inverted:   false,
									},
								},
								&notExpr{
									pos: position{line: 574, col: 21, offset: 16966},
									expr: &litMatcher{
										pos:        position{line: 574, col: 22, offset: 16967},
										val:        ".",
										ignoreCase: false,
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 581, col: 3, offset: 17142},
						run: (*parser).callonConst10,
						expr: &seqExpr{
							pos: position{line: 581, col: 3, offset: 17142},
							exprs: []interface{}{
								&zeroOrOneExpr{
									pos: position{line: 581, col: 3, offset: 17142},
									expr: &litMatcher{
										pos:        position{line: 581, col: 3, offset: 17142},
										val:        "-",
										ignoreCase: false,
									},
								},
								&oneOrMoreExpr{
									pos: position{line: 581, col: 8, offset: 17147},
									expr: &charClassMatcher{
										pos:        position{line: 581, col: 8, offset: 17147},
										val:        "[0-9]",
										ranges:     []rune{'0', '9'},
										ignoreCase: false,
										inverted:   false,
									},
								},
								&litMatcher{
									pos:        position{line: 581, col: 15, offset: 17154},
									val:        ".",
									ignoreCase: false,
								},
								&oneOrMoreExpr{
									pos: position{line: 581, col: 19, offset: 17158},
									expr: &charClassMatcher{
										pos:        position{line: 581, col: 19, offset: 17158},
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
						pos: position{line: 588, col: 3, offset: 17347},
						run: (*parser).callonConst19,
						expr: &litMatcher{
							pos:        position{line: 588, col: 3, offset: 17347},
							val:        "true",
							ignoreCase: false,
						},
					},
					&actionExpr{
						pos: position{line: 592, col: 3, offset: 17432},
						run: (*parser).callonConst21,
						expr: &litMatcher{
							pos:        position{line: 592, col: 3, offset: 17432},
							val:        "false",
							ignoreCase: false,
						},
					},
					&ruleRefExpr{
						pos:  position{line: 595, col: 3, offset: 17518},
						name: "Unit",
					},
					&ruleRefExpr{
						pos:  position{line: 596, col: 3, offset: 17525},
						name: "StringLiteral",
					},
					&actionExpr{
						pos: position{line: 597, col: 3, offset: 17541},
						run: (*parser).callonConst25,
						expr: &seqExpr{
							pos: position{line: 597, col: 3, offset: 17541},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 597, col: 3, offset: 17541},
									val:        "'",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 597, col: 7, offset: 17545},
									label: "val",
									expr: &seqExpr{
										pos: position{line: 597, col: 12, offset: 17550},
										exprs: []interface{}{
											&notExpr{
												pos: position{line: 597, col: 12, offset: 17550},
												expr: &ruleRefExpr{
													pos:  position{line: 597, col: 13, offset: 17551},
													name: "EscapedChar",
												},
											},
											&anyMatcher{
												line: 597, col: 25, offset: 17563,
											},
										},
									},
								},
								&litMatcher{
									pos:        position{line: 597, col: 28, offset: 17566},
									val:        "'",
									ignoreCase: false,
								},
							},
						},
					},
					&ruleRefExpr{
						pos:  position{line: 599, col: 5, offset: 17658},
						name: "ArrayLiteral",
					},
					&ruleRefExpr{
						pos:  position{line: 599, col: 20, offset: 17673},
						name: "RecordInstance",
					},
					&ruleRefExpr{
						pos:  position{line: 599, col: 37, offset: 17690},
						name: "VariantInstance",
					},
				},
			},
		},
		{
			name: "Unit",
			pos:  position{line: 601, col: 1, offset: 17707},
			expr: &actionExpr{
				pos: position{line: 601, col: 8, offset: 17714},
				run: (*parser).callonUnit1,
				expr: &litMatcher{
					pos:        position{line: 601, col: 8, offset: 17714},
					val:        "()",
					ignoreCase: false,
				},
			},
		},
		{
			name: "StringLiteral",
			pos:  position{line: 605, col: 1, offset: 17777},
			expr: &actionExpr{
				pos: position{line: 605, col: 17, offset: 17793},
				run: (*parser).callonStringLiteral1,
				expr: &seqExpr{
					pos: position{line: 605, col: 17, offset: 17793},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 605, col: 17, offset: 17793},
							val:        "\"",
							ignoreCase: false,
						},
						&zeroOrMoreExpr{
							pos: position{line: 605, col: 21, offset: 17797},
							expr: &seqExpr{
								pos: position{line: 605, col: 22, offset: 17798},
								exprs: []interface{}{
									&notExpr{
										pos: position{line: 605, col: 22, offset: 17798},
										expr: &ruleRefExpr{
											pos:  position{line: 605, col: 23, offset: 17799},
											name: "EscapedChar",
										},
									},
									&anyMatcher{
										line: 605, col: 35, offset: 17811,
									},
								},
							},
						},
						&litMatcher{
							pos:        position{line: 605, col: 39, offset: 17815},
							val:        "\"",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "Unused",
			pos:  position{line: 613, col: 1, offset: 17998},
			expr: &actionExpr{
				pos: position{line: 613, col: 10, offset: 18007},
				run: (*parser).callonUnused1,
				expr: &litMatcher{
					pos:        position{line: 613, col: 11, offset: 18008},
					val:        "_",
					ignoreCase: false,
				},
			},
		},
		{
			name: "Reserved",
			pos:  position{line: 617, col: 1, offset: 18063},
			expr: &seqExpr{
				pos: position{line: 617, col: 12, offset: 18074},
				exprs: []interface{}{
					&choiceExpr{
						pos: position{line: 617, col: 13, offset: 18075},
						alternatives: []interface{}{
							&litMatcher{
								pos:        position{line: 617, col: 13, offset: 18075},
								val:        "let",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 617, col: 21, offset: 18083},
								val:        "if",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 617, col: 28, offset: 18090},
								val:        "else",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 617, col: 37, offset: 18099},
								val:        "func",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 617, col: 46, offset: 18108},
								val:        "type",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 617, col: 55, offset: 18117},
								val:        "true",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 617, col: 64, offset: 18126},
								val:        "false",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 617, col: 74, offset: 18136},
								val:        "mutable",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 617, col: 86, offset: 18148},
								val:        "type",
								ignoreCase: false,
							},
							&ruleRefExpr{
								pos:  position{line: 617, col: 95, offset: 18157},
								name: "BaseType",
							},
						},
					},
					&notExpr{
						pos: position{line: 617, col: 105, offset: 18167},
						expr: &oneOrMoreExpr{
							pos: position{line: 617, col: 106, offset: 18168},
							expr: &charClassMatcher{
								pos:        position{line: 617, col: 106, offset: 18168},
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
			pos:  position{line: 619, col: 1, offset: 18176},
			expr: &choiceExpr{
				pos: position{line: 619, col: 12, offset: 18187},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 619, col: 12, offset: 18187},
						run: (*parser).callonBaseType2,
						expr: &choiceExpr{
							pos: position{line: 619, col: 14, offset: 18189},
							alternatives: []interface{}{
								&litMatcher{
									pos:        position{line: 619, col: 14, offset: 18189},
									val:        "int",
									ignoreCase: false,
								},
								&litMatcher{
									pos:        position{line: 619, col: 22, offset: 18197},
									val:        "bool",
									ignoreCase: false,
								},
								&litMatcher{
									pos:        position{line: 619, col: 31, offset: 18206},
									val:        "string",
									ignoreCase: false,
								},
								&litMatcher{
									pos:        position{line: 619, col: 42, offset: 18217},
									val:        "byte",
									ignoreCase: false,
								},
								&litMatcher{
									pos:        position{line: 619, col: 51, offset: 18226},
									val:        "rune",
									ignoreCase: false,
								},
								&litMatcher{
									pos:        position{line: 619, col: 60, offset: 18235},
									val:        "float",
									ignoreCase: false,
								},
								&litMatcher{
									pos:        position{line: 619, col: 70, offset: 18245},
									val:        "()",
									ignoreCase: false,
								},
							},
						},
					},
					&ruleRefExpr{
						pos:  position{line: 622, col: 3, offset: 18342},
						name: "Unit",
					},
				},
			},
		},
		{
			name: "EscapedChar",
			pos:  position{line: 624, col: 1, offset: 18348},
			expr: &charClassMatcher{
				pos:        position{line: 624, col: 15, offset: 18362},
				val:        "[\\x00-\\x1f\"\\\\]",
				chars:      []rune{'"', '\\'},
				ranges:     []rune{'\x00', '\x1f'},
				ignoreCase: false,
				inverted:   false,
			},
		},
		{
			name: "EscapeSequence",
			pos:  position{line: 626, col: 1, offset: 18378},
			expr: &choiceExpr{
				pos: position{line: 626, col: 18, offset: 18395},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 626, col: 18, offset: 18395},
						name: "SingleCharEscape",
					},
					&ruleRefExpr{
						pos:  position{line: 626, col: 37, offset: 18414},
						name: "UnicodeEscape",
					},
				},
			},
		},
		{
			name: "SingleCharEscape",
			pos:  position{line: 628, col: 1, offset: 18429},
			expr: &charClassMatcher{
				pos:        position{line: 628, col: 20, offset: 18448},
				val:        "[\"\\\\/bfnrt]",
				chars:      []rune{'"', '\\', '/', 'b', 'f', 'n', 'r', 't'},
				ignoreCase: false,
				inverted:   false,
			},
		},
		{
			name: "DecimalDigit",
			pos:  position{line: 630, col: 1, offset: 18461},
			expr: &charClassMatcher{
				pos:        position{line: 630, col: 16, offset: 18476},
				val:        "[0-9]",
				ranges:     []rune{'0', '9'},
				ignoreCase: false,
				inverted:   false,
			},
		},
		{
			name: "NonZeroDecimalDigit",
			pos:  position{line: 632, col: 1, offset: 18483},
			expr: &charClassMatcher{
				pos:        position{line: 632, col: 23, offset: 18505},
				val:        "[1-9]",
				ranges:     []rune{'1', '9'},
				ignoreCase: false,
				inverted:   false,
			},
		},
		{
			name: "HexDigit",
			pos:  position{line: 634, col: 1, offset: 18512},
			expr: &charClassMatcher{
				pos:        position{line: 634, col: 12, offset: 18523},
				val:        "[0-9a-f]i",
				ranges:     []rune{'0', '9', 'a', 'f'},
				ignoreCase: true,
				inverted:   false,
			},
		},
		{
			name:        "__",
			displayName: "\"sigwhitespace\"",
			pos:         position{line: 636, col: 1, offset: 18534},
			expr: &choiceExpr{
				pos: position{line: 636, col: 22, offset: 18555},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 636, col: 22, offset: 18555},
						name: "Comment",
					},
					&oneOrMoreExpr{
						pos: position{line: 636, col: 32, offset: 18565},
						expr: &charClassMatcher{
							pos:        position{line: 636, col: 32, offset: 18565},
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
			pos:         position{line: 638, col: 1, offset: 18577},
			expr: &choiceExpr{
				pos: position{line: 638, col: 18, offset: 18594},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 638, col: 18, offset: 18594},
						name: "Comment",
					},
					&zeroOrMoreExpr{
						pos: position{line: 638, col: 28, offset: 18604},
						expr: &charClassMatcher{
							pos:        position{line: 638, col: 28, offset: 18604},
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
			pos:  position{line: 640, col: 1, offset: 18616},
			expr: &seqExpr{
				pos: position{line: 640, col: 11, offset: 18626},
				exprs: []interface{}{
					&zeroOrMoreExpr{
						pos: position{line: 640, col: 11, offset: 18626},
						expr: &charClassMatcher{
							pos:        position{line: 640, col: 11, offset: 18626},
							val:        "[ \\r\\n\\t]",
							chars:      []rune{' ', '\r', '\n', '\t'},
							ignoreCase: false,
							inverted:   false,
						},
					},
					&litMatcher{
						pos:        position{line: 640, col: 22, offset: 18637},
						val:        "#",
						ignoreCase: false,
					},
					&zeroOrMoreExpr{
						pos: position{line: 640, col: 26, offset: 18641},
						expr: &seqExpr{
							pos: position{line: 640, col: 27, offset: 18642},
							exprs: []interface{}{
								&notExpr{
									pos: position{line: 640, col: 27, offset: 18642},
									expr: &charClassMatcher{
										pos:        position{line: 640, col: 28, offset: 18643},
										val:        "[\\n]",
										chars:      []rune{'\n'},
										ignoreCase: false,
										inverted:   false,
									},
								},
								&anyMatcher{
									line: 640, col: 33, offset: 18648,
								},
							},
						},
					},
					&andExpr{
						pos: position{line: 640, col: 37, offset: 18652},
						expr: &litMatcher{
							pos:        position{line: 640, col: 38, offset: 18653},
							val:        "\n",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "EOF",
			pos:  position{line: 642, col: 1, offset: 18659},
			expr: &notExpr{
				pos: position{line: 642, col: 7, offset: 18665},
				expr: &anyMatcher{
					line: 642, col: 8, offset: 18666,
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
	//fmt.Println(e)
	// wrap calls as statements in an expr
	switch e.(type) {
	case Call:
		ex := Expr{Subvalues: []Ast{e.(Call)}, AsStatement: true}
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
