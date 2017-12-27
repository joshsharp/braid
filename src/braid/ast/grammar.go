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
							name: "__N",
						},
						&labeledExpr{
							pos:   position{line: 13, col: 25, offset: 166},
							label: "name",
							expr: &ruleRefExpr{
								pos:  position{line: 13, col: 30, offset: 171},
								name: "ModuleName",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 13, col: 41, offset: 182},
							name: "__",
						},
						&labeledExpr{
							pos:   position{line: 13, col: 44, offset: 185},
							label: "stat",
							expr: &ruleRefExpr{
								pos:  position{line: 13, col: 49, offset: 190},
								name: "TopLevelStatement",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 13, col: 67, offset: 208},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 13, col: 69, offset: 210},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 13, col: 74, offset: 215},
								expr: &ruleRefExpr{
									pos:  position{line: 13, col: 75, offset: 216},
									name: "TopLevelStatement",
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 13, col: 95, offset: 236},
							name: "_",
						},
						&ruleRefExpr{
							pos:  position{line: 13, col: 97, offset: 238},
							name: "EOF",
						},
					},
				},
			},
		},
		{
			name: "TopLevelStatement",
			pos:  position{line: 28, col: 1, offset: 719},
			expr: &choiceExpr{
				pos: position{line: 28, col: 21, offset: 739},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 28, col: 21, offset: 739},
						name: "FuncDefn",
					},
					&ruleRefExpr{
						pos:  position{line: 28, col: 32, offset: 750},
						name: "TypeDefn",
					},
					&ruleRefExpr{
						pos:  position{line: 28, col: 43, offset: 761},
						name: "ExternFunc",
					},
				},
			},
		},
		{
			name: "Statement",
			pos:  position{line: 30, col: 1, offset: 773},
			expr: &choiceExpr{
				pos: position{line: 30, col: 13, offset: 785},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 30, col: 13, offset: 785},
						name: "FuncDefn",
					},
					&ruleRefExpr{
						pos:  position{line: 30, col: 24, offset: 796},
						name: "Assignment",
					},
					&ruleRefExpr{
						pos:  position{line: 30, col: 37, offset: 809},
						name: "ExprLine",
					},
				},
			},
		},
		{
			name: "ExprLine",
			pos:  position{line: 32, col: 1, offset: 819},
			expr: &actionExpr{
				pos: position{line: 32, col: 12, offset: 830},
				run: (*parser).callonExprLine1,
				expr: &seqExpr{
					pos: position{line: 32, col: 12, offset: 830},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 32, col: 12, offset: 830},
							label: "e",
							expr: &ruleRefExpr{
								pos:  position{line: 32, col: 14, offset: 832},
								name: "Expr",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 32, col: 19, offset: 837},
							name: "N",
						},
					},
				},
			},
		},
		{
			name: "ExternFunc",
			pos:  position{line: 46, col: 1, offset: 1168},
			expr: &actionExpr{
				pos: position{line: 46, col: 14, offset: 1181},
				run: (*parser).callonExternFunc1,
				expr: &seqExpr{
					pos: position{line: 46, col: 14, offset: 1181},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 46, col: 14, offset: 1181},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 46, col: 16, offset: 1183},
							val:        "extern",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 46, col: 25, offset: 1192},
							name: "__N",
						},
						&litMatcher{
							pos:        position{line: 46, col: 29, offset: 1196},
							val:        "func",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 46, col: 36, offset: 1203},
							name: "__N",
						},
						&labeledExpr{
							pos:   position{line: 46, col: 40, offset: 1207},
							label: "name",
							expr: &ruleRefExpr{
								pos:  position{line: 46, col: 45, offset: 1212},
								name: "Identifier",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 46, col: 56, offset: 1223},
							name: "_N",
						},
						&litMatcher{
							pos:        position{line: 46, col: 59, offset: 1226},
							val:        "=",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 46, col: 63, offset: 1230},
							name: "_N",
						},
						&labeledExpr{
							pos:   position{line: 46, col: 66, offset: 1233},
							label: "importName",
							expr: &ruleRefExpr{
								pos:  position{line: 46, col: 77, offset: 1244},
								name: "StringLiteral",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 46, col: 91, offset: 1258},
							name: "_N",
						},
						&labeledExpr{
							pos:   position{line: 46, col: 94, offset: 1261},
							label: "args",
							expr: &ruleRefExpr{
								pos:  position{line: 46, col: 99, offset: 1266},
								name: "ArgsDefn",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 46, col: 108, offset: 1275},
							name: "_N",
						},
						&labeledExpr{
							pos:   position{line: 46, col: 111, offset: 1278},
							label: "ret",
							expr: &ruleRefExpr{
								pos:  position{line: 46, col: 115, offset: 1282},
								name: "BaseType",
							},
						},
					},
				},
			},
		},
		{
			name: "TypeDefn",
			pos:  position{line: 52, col: 1, offset: 1498},
			expr: &choiceExpr{
				pos: position{line: 52, col: 12, offset: 1509},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 52, col: 12, offset: 1509},
						run: (*parser).callonTypeDefn2,
						expr: &seqExpr{
							pos: position{line: 52, col: 12, offset: 1509},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 52, col: 12, offset: 1509},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 52, col: 14, offset: 1511},
									val:        "type",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 52, col: 21, offset: 1518},
									name: "__",
								},
								&labeledExpr{
									pos:   position{line: 52, col: 24, offset: 1521},
									label: "name",
									expr: &ruleRefExpr{
										pos:  position{line: 52, col: 29, offset: 1526},
										name: "ModuleName",
									},
								},
								&labeledExpr{
									pos:   position{line: 52, col: 40, offset: 1537},
									label: "params",
									expr: &zeroOrMoreExpr{
										pos: position{line: 52, col: 47, offset: 1544},
										expr: &seqExpr{
											pos: position{line: 52, col: 48, offset: 1545},
											exprs: []interface{}{
												&ruleRefExpr{
													pos:  position{line: 52, col: 48, offset: 1545},
													name: "__",
												},
												&ruleRefExpr{
													pos:  position{line: 52, col: 51, offset: 1548},
													name: "TypeParameter",
												},
											},
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 52, col: 67, offset: 1564},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 52, col: 69, offset: 1566},
									val:        "=",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 52, col: 73, offset: 1570},
									label: "types",
									expr: &oneOrMoreExpr{
										pos: position{line: 52, col: 79, offset: 1576},
										expr: &seqExpr{
											pos: position{line: 52, col: 80, offset: 1577},
											exprs: []interface{}{
												&ruleRefExpr{
													pos:  position{line: 52, col: 80, offset: 1577},
													name: "__",
												},
												&ruleRefExpr{
													pos:  position{line: 52, col: 83, offset: 1580},
													name: "AnyType",
												},
											},
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 52, col: 93, offset: 1590},
									name: "__",
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 71, col: 1, offset: 2086},
						run: (*parser).callonTypeDefn22,
						expr: &seqExpr{
							pos: position{line: 71, col: 1, offset: 2086},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 71, col: 1, offset: 2086},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 71, col: 3, offset: 2088},
									val:        "type",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 71, col: 10, offset: 2095},
									name: "__",
								},
								&labeledExpr{
									pos:   position{line: 71, col: 13, offset: 2098},
									label: "name",
									expr: &ruleRefExpr{
										pos:  position{line: 71, col: 18, offset: 2103},
										name: "ModuleName",
									},
								},
								&labeledExpr{
									pos:   position{line: 71, col: 29, offset: 2114},
									label: "params",
									expr: &zeroOrMoreExpr{
										pos: position{line: 71, col: 36, offset: 2121},
										expr: &seqExpr{
											pos: position{line: 71, col: 37, offset: 2122},
											exprs: []interface{}{
												&ruleRefExpr{
													pos:  position{line: 71, col: 37, offset: 2122},
													name: "__",
												},
												&ruleRefExpr{
													pos:  position{line: 71, col: 40, offset: 2125},
													name: "TypeParameter",
												},
											},
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 71, col: 56, offset: 2141},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 71, col: 58, offset: 2143},
									val:        "=",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 71, col: 62, offset: 2147},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 72, col: 5, offset: 2153},
									val:        "{",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 72, col: 9, offset: 2157},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 72, col: 11, offset: 2159},
									label: "first",
									expr: &ruleRefExpr{
										pos:  position{line: 72, col: 17, offset: 2165},
										name: "RecordFieldDefn",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 72, col: 33, offset: 2181},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 72, col: 35, offset: 2183},
									label: "rest",
									expr: &zeroOrMoreExpr{
										pos: position{line: 72, col: 40, offset: 2188},
										expr: &seqExpr{
											pos: position{line: 72, col: 41, offset: 2189},
											exprs: []interface{}{
												&litMatcher{
													pos:        position{line: 72, col: 41, offset: 2189},
													val:        ",",
													ignoreCase: false,
												},
												&ruleRefExpr{
													pos:  position{line: 72, col: 45, offset: 2193},
													name: "_",
												},
												&ruleRefExpr{
													pos:  position{line: 72, col: 47, offset: 2195},
													name: "RecordFieldDefn",
												},
												&ruleRefExpr{
													pos:  position{line: 72, col: 63, offset: 2211},
													name: "_",
												},
											},
										},
									},
								},
								&zeroOrOneExpr{
									pos: position{line: 72, col: 67, offset: 2215},
									expr: &litMatcher{
										pos:        position{line: 72, col: 67, offset: 2215},
										val:        ",",
										ignoreCase: false,
									},
								},
								&ruleRefExpr{
									pos:  position{line: 72, col: 72, offset: 2220},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 72, col: 74, offset: 2222},
									val:        "}",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 72, col: 78, offset: 2226},
									name: "__",
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 90, col: 1, offset: 2711},
						run: (*parser).callonTypeDefn54,
						expr: &seqExpr{
							pos: position{line: 90, col: 1, offset: 2711},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 90, col: 1, offset: 2711},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 90, col: 3, offset: 2713},
									val:        "type",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 90, col: 10, offset: 2720},
									name: "__",
								},
								&labeledExpr{
									pos:   position{line: 90, col: 13, offset: 2723},
									label: "name",
									expr: &ruleRefExpr{
										pos:  position{line: 90, col: 18, offset: 2728},
										name: "ModuleName",
									},
								},
								&labeledExpr{
									pos:   position{line: 90, col: 29, offset: 2739},
									label: "params",
									expr: &zeroOrMoreExpr{
										pos: position{line: 90, col: 36, offset: 2746},
										expr: &seqExpr{
											pos: position{line: 90, col: 37, offset: 2747},
											exprs: []interface{}{
												&ruleRefExpr{
													pos:  position{line: 90, col: 37, offset: 2747},
													name: "__",
												},
												&ruleRefExpr{
													pos:  position{line: 90, col: 40, offset: 2750},
													name: "TypeParameter",
												},
											},
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 90, col: 56, offset: 2766},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 90, col: 58, offset: 2768},
									val:        "=",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 90, col: 62, offset: 2772},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 90, col: 64, offset: 2774},
									label: "rest",
									expr: &oneOrMoreExpr{
										pos: position{line: 90, col: 69, offset: 2779},
										expr: &ruleRefExpr{
											pos:  position{line: 90, col: 70, offset: 2780},
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
			pos:  position{line: 105, col: 1, offset: 3187},
			expr: &actionExpr{
				pos: position{line: 105, col: 19, offset: 3205},
				run: (*parser).callonRecordFieldDefn1,
				expr: &seqExpr{
					pos: position{line: 105, col: 19, offset: 3205},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 105, col: 19, offset: 3205},
							label: "name",
							expr: &ruleRefExpr{
								pos:  position{line: 105, col: 24, offset: 3210},
								name: "VariableName",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 105, col: 37, offset: 3223},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 105, col: 39, offset: 3225},
							val:        ":",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 105, col: 43, offset: 3229},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 105, col: 45, offset: 3231},
							label: "t",
							expr: &ruleRefExpr{
								pos:  position{line: 105, col: 48, offset: 3234},
								name: "AnyType",
							},
						},
					},
				},
			},
		},
		{
			name: "VariantConstructor",
			pos:  position{line: 109, col: 1, offset: 3328},
			expr: &choiceExpr{
				pos: position{line: 109, col: 22, offset: 3349},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 109, col: 22, offset: 3349},
						run: (*parser).callonVariantConstructor2,
						expr: &seqExpr{
							pos: position{line: 109, col: 22, offset: 3349},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 109, col: 22, offset: 3349},
									val:        "|",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 109, col: 26, offset: 3353},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 109, col: 28, offset: 3355},
									label: "name",
									expr: &ruleRefExpr{
										pos:  position{line: 109, col: 33, offset: 3360},
										name: "ModuleName",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 109, col: 44, offset: 3371},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 109, col: 46, offset: 3373},
									val:        "{",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 109, col: 50, offset: 3377},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 109, col: 52, offset: 3379},
									label: "first",
									expr: &ruleRefExpr{
										pos:  position{line: 109, col: 58, offset: 3385},
										name: "RecordFieldDefn",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 109, col: 74, offset: 3401},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 109, col: 76, offset: 3403},
									label: "rest",
									expr: &zeroOrMoreExpr{
										pos: position{line: 109, col: 81, offset: 3408},
										expr: &seqExpr{
											pos: position{line: 109, col: 82, offset: 3409},
											exprs: []interface{}{
												&litMatcher{
													pos:        position{line: 109, col: 82, offset: 3409},
													val:        ",",
													ignoreCase: false,
												},
												&ruleRefExpr{
													pos:  position{line: 109, col: 86, offset: 3413},
													name: "_",
												},
												&ruleRefExpr{
													pos:  position{line: 109, col: 88, offset: 3415},
													name: "RecordFieldDefn",
												},
												&ruleRefExpr{
													pos:  position{line: 109, col: 104, offset: 3431},
													name: "_",
												},
											},
										},
									},
								},
								&zeroOrOneExpr{
									pos: position{line: 109, col: 108, offset: 3435},
									expr: &litMatcher{
										pos:        position{line: 109, col: 108, offset: 3435},
										val:        ",",
										ignoreCase: false,
									},
								},
								&ruleRefExpr{
									pos:  position{line: 109, col: 113, offset: 3440},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 109, col: 115, offset: 3442},
									val:        "}",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 109, col: 119, offset: 3446},
									name: "__",
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 128, col: 1, offset: 4051},
						run: (*parser).callonVariantConstructor26,
						expr: &seqExpr{
							pos: position{line: 128, col: 1, offset: 4051},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 128, col: 1, offset: 4051},
									val:        "|",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 128, col: 5, offset: 4055},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 128, col: 7, offset: 4057},
									label: "name",
									expr: &ruleRefExpr{
										pos:  position{line: 128, col: 12, offset: 4062},
										name: "ModuleName",
									},
								},
								&labeledExpr{
									pos:   position{line: 128, col: 23, offset: 4073},
									label: "rest",
									expr: &zeroOrMoreExpr{
										pos: position{line: 128, col: 28, offset: 4078},
										expr: &seqExpr{
											pos: position{line: 128, col: 29, offset: 4079},
											exprs: []interface{}{
												&ruleRefExpr{
													pos:  position{line: 128, col: 29, offset: 4079},
													name: "__",
												},
												&ruleRefExpr{
													pos:  position{line: 128, col: 32, offset: 4082},
													name: "AnyType",
												},
											},
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 128, col: 42, offset: 4092},
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
			pos:  position{line: 145, col: 1, offset: 4529},
			expr: &choiceExpr{
				pos: position{line: 145, col: 11, offset: 4539},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 145, col: 11, offset: 4539},
						name: "BaseType",
					},
					&ruleRefExpr{
						pos:  position{line: 145, col: 22, offset: 4550},
						name: "TypeParameter",
					},
				},
			},
		},
		{
			name: "Assignment",
			pos:  position{line: 147, col: 1, offset: 4565},
			expr: &choiceExpr{
				pos: position{line: 147, col: 14, offset: 4578},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 147, col: 14, offset: 4578},
						run: (*parser).callonAssignment2,
						expr: &seqExpr{
							pos: position{line: 147, col: 14, offset: 4578},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 147, col: 14, offset: 4578},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 147, col: 16, offset: 4580},
									val:        "let",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 147, col: 22, offset: 4586},
									name: "__N",
								},
								&labeledExpr{
									pos:   position{line: 147, col: 26, offset: 4590},
									label: "i",
									expr: &ruleRefExpr{
										pos:  position{line: 147, col: 28, offset: 4592},
										name: "Assignable",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 147, col: 39, offset: 4603},
									name: "_N",
								},
								&litMatcher{
									pos:        position{line: 147, col: 42, offset: 4606},
									val:        "=",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 147, col: 46, offset: 4610},
									name: "_N",
								},
								&labeledExpr{
									pos:   position{line: 147, col: 49, offset: 4613},
									label: "expr",
									expr: &ruleRefExpr{
										pos:  position{line: 147, col: 54, offset: 4618},
										name: "Expr",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 147, col: 59, offset: 4623},
									name: "N",
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 153, col: 1, offset: 4742},
						run: (*parser).callonAssignment15,
						expr: &seqExpr{
							pos: position{line: 153, col: 1, offset: 4742},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 153, col: 1, offset: 4742},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 153, col: 3, offset: 4744},
									val:        "let",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 153, col: 9, offset: 4750},
									name: "__N",
								},
								&notExpr{
									pos: position{line: 153, col: 13, offset: 4754},
									expr: &ruleRefExpr{
										pos:  position{line: 153, col: 14, offset: 4755},
										name: "Assignable",
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 157, col: 1, offset: 4853},
						run: (*parser).callonAssignment22,
						expr: &seqExpr{
							pos: position{line: 157, col: 1, offset: 4853},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 157, col: 1, offset: 4853},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 157, col: 3, offset: 4855},
									val:        "let",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 157, col: 9, offset: 4861},
									name: "__N",
								},
								&labeledExpr{
									pos:   position{line: 157, col: 13, offset: 4865},
									label: "i",
									expr: &ruleRefExpr{
										pos:  position{line: 157, col: 15, offset: 4867},
										name: "Assignable",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 157, col: 26, offset: 4878},
									name: "_N",
								},
								&notExpr{
									pos: position{line: 157, col: 29, offset: 4881},
									expr: &litMatcher{
										pos:        position{line: 157, col: 30, offset: 4882},
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
			pos:  position{line: 161, col: 1, offset: 4976},
			expr: &actionExpr{
				pos: position{line: 161, col: 12, offset: 4987},
				run: (*parser).callonFuncDefn1,
				expr: &seqExpr{
					pos: position{line: 161, col: 12, offset: 4987},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 161, col: 12, offset: 4987},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 161, col: 14, offset: 4989},
							val:        "let",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 161, col: 20, offset: 4995},
							name: "__N",
						},
						&labeledExpr{
							pos:   position{line: 161, col: 24, offset: 4999},
							label: "i",
							expr: &ruleRefExpr{
								pos:  position{line: 161, col: 26, offset: 5001},
								name: "VariableName",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 161, col: 39, offset: 5014},
							name: "_N",
						},
						&litMatcher{
							pos:        position{line: 161, col: 42, offset: 5017},
							val:        "=",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 161, col: 46, offset: 5021},
							name: "_N",
						},
						&labeledExpr{
							pos:   position{line: 161, col: 49, offset: 5024},
							label: "ids",
							expr: &zeroOrOneExpr{
								pos: position{line: 161, col: 53, offset: 5028},
								expr: &seqExpr{
									pos: position{line: 161, col: 54, offset: 5029},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 161, col: 54, offset: 5029},
											name: "ArgsDefn",
										},
										&ruleRefExpr{
											pos:  position{line: 161, col: 63, offset: 5038},
											name: "_",
										},
									},
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 161, col: 67, offset: 5042},
							label: "ret",
							expr: &zeroOrOneExpr{
								pos: position{line: 161, col: 71, offset: 5046},
								expr: &seqExpr{
									pos: position{line: 161, col: 72, offset: 5047},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 161, col: 72, offset: 5047},
											name: "AnyType",
										},
										&ruleRefExpr{
											pos:  position{line: 161, col: 80, offset: 5055},
											name: "_",
										},
									},
								},
							},
						},
						&litMatcher{
							pos:        position{line: 161, col: 84, offset: 5059},
							val:        "{",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 161, col: 88, offset: 5063},
							name: "__",
						},
						&labeledExpr{
							pos:   position{line: 161, col: 91, offset: 5066},
							label: "statements",
							expr: &oneOrMoreExpr{
								pos: position{line: 161, col: 102, offset: 5077},
								expr: &ruleRefExpr{
									pos:  position{line: 161, col: 103, offset: 5078},
									name: "Statement",
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 161, col: 115, offset: 5090},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 161, col: 117, offset: 5092},
							val:        "}",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 161, col: 121, offset: 5096},
							name: "_",
						},
					},
				},
			},
		},
		{
			name: "Expr",
			pos:  position{line: 187, col: 1, offset: 5706},
			expr: &actionExpr{
				pos: position{line: 187, col: 8, offset: 5713},
				run: (*parser).callonExpr1,
				expr: &labeledExpr{
					pos:   position{line: 187, col: 8, offset: 5713},
					label: "ex",
					expr: &choiceExpr{
						pos: position{line: 187, col: 12, offset: 5717},
						alternatives: []interface{}{
							&ruleRefExpr{
								pos:  position{line: 187, col: 12, offset: 5717},
								name: "IfExpr",
							},
							&ruleRefExpr{
								pos:  position{line: 187, col: 21, offset: 5726},
								name: "CompoundExpr",
							},
						},
					},
				},
			},
		},
		{
			name: "IfExpr",
			pos:  position{line: 193, col: 1, offset: 5843},
			expr: &choiceExpr{
				pos: position{line: 193, col: 10, offset: 5852},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 193, col: 10, offset: 5852},
						run: (*parser).callonIfExpr2,
						expr: &seqExpr{
							pos: position{line: 193, col: 10, offset: 5852},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 193, col: 10, offset: 5852},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 193, col: 12, offset: 5854},
									val:        "if",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 193, col: 17, offset: 5859},
									name: "__N",
								},
								&labeledExpr{
									pos:   position{line: 193, col: 21, offset: 5863},
									label: "expr",
									expr: &ruleRefExpr{
										pos:  position{line: 193, col: 26, offset: 5868},
										name: "BinOpBool",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 193, col: 36, offset: 5878},
									name: "_N",
								},
								&litMatcher{
									pos:        position{line: 193, col: 39, offset: 5881},
									val:        "{",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 193, col: 43, offset: 5885},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 193, col: 45, offset: 5887},
									label: "thens",
									expr: &oneOrMoreExpr{
										pos: position{line: 193, col: 51, offset: 5893},
										expr: &ruleRefExpr{
											pos:  position{line: 193, col: 52, offset: 5894},
											name: "Statement",
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 193, col: 64, offset: 5906},
									name: "__",
								},
								&litMatcher{
									pos:        position{line: 193, col: 67, offset: 5909},
									val:        "}",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 193, col: 71, offset: 5913},
									name: "_N",
								},
								&litMatcher{
									pos:        position{line: 193, col: 74, offset: 5916},
									val:        "else",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 193, col: 81, offset: 5923},
									name: "_N",
								},
								&litMatcher{
									pos:        position{line: 193, col: 84, offset: 5926},
									val:        "{",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 193, col: 88, offset: 5930},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 193, col: 90, offset: 5932},
									label: "elses",
									expr: &oneOrMoreExpr{
										pos: position{line: 193, col: 96, offset: 5938},
										expr: &ruleRefExpr{
											pos:  position{line: 193, col: 97, offset: 5939},
											name: "Statement",
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 193, col: 109, offset: 5951},
									name: "__",
								},
								&litMatcher{
									pos:        position{line: 193, col: 112, offset: 5954},
									val:        "}",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 212, col: 1, offset: 6457},
						run: (*parser).callonIfExpr27,
						expr: &seqExpr{
							pos: position{line: 212, col: 1, offset: 6457},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 212, col: 1, offset: 6457},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 212, col: 3, offset: 6459},
									val:        "if",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 212, col: 8, offset: 6464},
									name: "__N",
								},
								&labeledExpr{
									pos:   position{line: 212, col: 12, offset: 6468},
									label: "expr",
									expr: &ruleRefExpr{
										pos:  position{line: 212, col: 17, offset: 6473},
										name: "BinOpBool",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 212, col: 27, offset: 6483},
									name: "_N",
								},
								&litMatcher{
									pos:        position{line: 212, col: 30, offset: 6486},
									val:        "{",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 212, col: 34, offset: 6490},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 212, col: 36, offset: 6492},
									label: "thens",
									expr: &oneOrMoreExpr{
										pos: position{line: 212, col: 42, offset: 6498},
										expr: &ruleRefExpr{
											pos:  position{line: 212, col: 43, offset: 6499},
											name: "Statement",
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 212, col: 55, offset: 6511},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 212, col: 57, offset: 6513},
									val:        "}",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 212, col: 61, offset: 6517},
									name: "_N",
								},
								&litMatcher{
									pos:        position{line: 212, col: 64, offset: 6520},
									val:        "else",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 212, col: 71, offset: 6527},
									label: "elseifs",
									expr: &ruleRefExpr{
										pos:  position{line: 212, col: 79, offset: 6535},
										name: "IfExpr",
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 224, col: 1, offset: 6865},
						run: (*parser).callonIfExpr46,
						expr: &seqExpr{
							pos: position{line: 224, col: 1, offset: 6865},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 224, col: 1, offset: 6865},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 224, col: 3, offset: 6867},
									val:        "if",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 224, col: 8, offset: 6872},
									name: "__N",
								},
								&labeledExpr{
									pos:   position{line: 224, col: 12, offset: 6876},
									label: "expr",
									expr: &ruleRefExpr{
										pos:  position{line: 224, col: 17, offset: 6881},
										name: "BinOpBool",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 224, col: 27, offset: 6891},
									name: "_N",
								},
								&litMatcher{
									pos:        position{line: 224, col: 30, offset: 6894},
									val:        "{",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 224, col: 34, offset: 6898},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 224, col: 36, offset: 6900},
									label: "thens",
									expr: &oneOrMoreExpr{
										pos: position{line: 224, col: 42, offset: 6906},
										expr: &ruleRefExpr{
											pos:  position{line: 224, col: 43, offset: 6907},
											name: "Statement",
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 224, col: 55, offset: 6919},
									name: "__",
								},
								&litMatcher{
									pos:        position{line: 224, col: 58, offset: 6922},
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
			pos:  position{line: 236, col: 1, offset: 7220},
			expr: &choiceExpr{
				pos: position{line: 236, col: 8, offset: 7227},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 236, col: 8, offset: 7227},
						run: (*parser).callonCall2,
						expr: &seqExpr{
							pos: position{line: 236, col: 8, offset: 7227},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 236, col: 8, offset: 7227},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 236, col: 10, offset: 7229},
									label: "module",
									expr: &ruleRefExpr{
										pos:  position{line: 236, col: 17, offset: 7236},
										name: "ModuleName",
									},
								},
								&litMatcher{
									pos:        position{line: 236, col: 28, offset: 7247},
									val:        ".",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 236, col: 32, offset: 7251},
									label: "fn",
									expr: &ruleRefExpr{
										pos:  position{line: 236, col: 35, offset: 7254},
										name: "VariableName",
									},
								},
								&labeledExpr{
									pos:   position{line: 236, col: 48, offset: 7267},
									label: "args",
									expr: &ruleRefExpr{
										pos:  position{line: 236, col: 53, offset: 7272},
										name: "Arguments",
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 250, col: 1, offset: 7596},
						run: (*parser).callonCall12,
						expr: &seqExpr{
							pos: position{line: 250, col: 1, offset: 7596},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 250, col: 1, offset: 7596},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 250, col: 3, offset: 7598},
									label: "fn",
									expr: &ruleRefExpr{
										pos:  position{line: 250, col: 6, offset: 7601},
										name: "VariableName",
									},
								},
								&labeledExpr{
									pos:   position{line: 250, col: 19, offset: 7614},
									label: "args",
									expr: &ruleRefExpr{
										pos:  position{line: 250, col: 24, offset: 7619},
										name: "Arguments",
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 264, col: 1, offset: 7936},
						run: (*parser).callonCall19,
						expr: &seqExpr{
							pos: position{line: 264, col: 1, offset: 7936},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 264, col: 1, offset: 7936},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 264, col: 3, offset: 7938},
									label: "fn",
									expr: &ruleRefExpr{
										pos:  position{line: 264, col: 6, offset: 7941},
										name: "VariableName",
									},
								},
								&litMatcher{
									pos:        position{line: 264, col: 19, offset: 7954},
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
			name: "RecordAccess",
			pos:  position{line: 271, col: 1, offset: 8125},
			expr: &actionExpr{
				pos: position{line: 271, col: 16, offset: 8140},
				run: (*parser).callonRecordAccess1,
				expr: &seqExpr{
					pos: position{line: 271, col: 16, offset: 8140},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 271, col: 16, offset: 8140},
							label: "record",
							expr: &ruleRefExpr{
								pos:  position{line: 271, col: 23, offset: 8147},
								name: "VariableName",
							},
						},
						&labeledExpr{
							pos:   position{line: 271, col: 36, offset: 8160},
							label: "rest",
							expr: &oneOrMoreExpr{
								pos: position{line: 271, col: 41, offset: 8165},
								expr: &seqExpr{
									pos: position{line: 271, col: 42, offset: 8166},
									exprs: []interface{}{
										&litMatcher{
											pos:        position{line: 271, col: 42, offset: 8166},
											val:        ".",
											ignoreCase: false,
										},
										&ruleRefExpr{
											pos:  position{line: 271, col: 46, offset: 8170},
											name: "VariableName",
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
			name: "ArgsDefn",
			pos:  position{line: 288, col: 1, offset: 8607},
			expr: &actionExpr{
				pos: position{line: 288, col: 12, offset: 8618},
				run: (*parser).callonArgsDefn1,
				expr: &seqExpr{
					pos: position{line: 288, col: 12, offset: 8618},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 288, col: 12, offset: 8618},
							val:        "(",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 288, col: 16, offset: 8622},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 288, col: 18, offset: 8624},
							label: "argument",
							expr: &ruleRefExpr{
								pos:  position{line: 288, col: 27, offset: 8633},
								name: "ArgDefn",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 288, col: 35, offset: 8641},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 288, col: 37, offset: 8643},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 288, col: 42, offset: 8648},
								expr: &seqExpr{
									pos: position{line: 288, col: 43, offset: 8649},
									exprs: []interface{}{
										&litMatcher{
											pos:        position{line: 288, col: 43, offset: 8649},
											val:        ",",
											ignoreCase: false,
										},
										&ruleRefExpr{
											pos:  position{line: 288, col: 47, offset: 8653},
											name: "_",
										},
										&ruleRefExpr{
											pos:  position{line: 288, col: 49, offset: 8655},
											name: "ArgDefn",
										},
									},
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 288, col: 59, offset: 8665},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 288, col: 61, offset: 8667},
							val:        ")",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "ArgDefn",
			pos:  position{line: 306, col: 1, offset: 9089},
			expr: &actionExpr{
				pos: position{line: 306, col: 11, offset: 9099},
				run: (*parser).callonArgDefn1,
				expr: &seqExpr{
					pos: position{line: 306, col: 11, offset: 9099},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 306, col: 11, offset: 9099},
							label: "name",
							expr: &ruleRefExpr{
								pos:  position{line: 306, col: 16, offset: 9104},
								name: "Identifier",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 306, col: 27, offset: 9115},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 306, col: 29, offset: 9117},
							label: "anno",
							expr: &zeroOrOneExpr{
								pos: position{line: 306, col: 34, offset: 9122},
								expr: &seqExpr{
									pos: position{line: 306, col: 35, offset: 9123},
									exprs: []interface{}{
										&litMatcher{
											pos:        position{line: 306, col: 35, offset: 9123},
											val:        ":",
											ignoreCase: false,
										},
										&ruleRefExpr{
											pos:  position{line: 306, col: 39, offset: 9127},
											name: "_",
										},
										&ruleRefExpr{
											pos:  position{line: 306, col: 41, offset: 9129},
											name: "AnyType",
										},
									},
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 306, col: 52, offset: 9140},
							name: "_",
						},
					},
				},
			},
		},
		{
			name: "Arguments",
			pos:  position{line: 326, col: 1, offset: 9628},
			expr: &choiceExpr{
				pos: position{line: 326, col: 13, offset: 9640},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 326, col: 13, offset: 9640},
						run: (*parser).callonArguments2,
						expr: &seqExpr{
							pos: position{line: 326, col: 13, offset: 9640},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 326, col: 13, offset: 9640},
									val:        "(",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 326, col: 17, offset: 9644},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 326, col: 19, offset: 9646},
									label: "argument",
									expr: &ruleRefExpr{
										pos:  position{line: 326, col: 28, offset: 9655},
										name: "BinOpParens",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 326, col: 40, offset: 9667},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 326, col: 42, offset: 9669},
									label: "rest",
									expr: &zeroOrMoreExpr{
										pos: position{line: 326, col: 47, offset: 9674},
										expr: &seqExpr{
											pos: position{line: 326, col: 48, offset: 9675},
											exprs: []interface{}{
												&litMatcher{
													pos:        position{line: 326, col: 48, offset: 9675},
													val:        ",",
													ignoreCase: false,
												},
												&ruleRefExpr{
													pos:  position{line: 326, col: 52, offset: 9679},
													name: "_",
												},
												&ruleRefExpr{
													pos:  position{line: 326, col: 54, offset: 9681},
													name: "BinOpParens",
												},
											},
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 326, col: 68, offset: 9695},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 326, col: 70, offset: 9697},
									val:        ")",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 343, col: 1, offset: 10119},
						run: (*parser).callonArguments17,
						expr: &seqExpr{
							pos: position{line: 343, col: 1, offset: 10119},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 343, col: 1, offset: 10119},
									val:        "(",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 343, col: 5, offset: 10123},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 343, col: 7, offset: 10125},
									label: "argument",
									expr: &ruleRefExpr{
										pos:  position{line: 343, col: 16, offset: 10134},
										name: "Expr",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 343, col: 21, offset: 10139},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 343, col: 23, offset: 10141},
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
			pos:  position{line: 348, col: 1, offset: 10246},
			expr: &actionExpr{
				pos: position{line: 348, col: 16, offset: 10261},
				run: (*parser).callonCompoundExpr1,
				expr: &seqExpr{
					pos: position{line: 348, col: 16, offset: 10261},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 348, col: 16, offset: 10261},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 348, col: 18, offset: 10263},
							label: "op",
							expr: &ruleRefExpr{
								pos:  position{line: 348, col: 21, offset: 10266},
								name: "BinOp",
							},
						},
						&labeledExpr{
							pos:   position{line: 348, col: 27, offset: 10272},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 348, col: 32, offset: 10277},
								expr: &seqExpr{
									pos: position{line: 348, col: 33, offset: 10278},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 348, col: 33, offset: 10278},
											name: "__N",
										},
										&ruleRefExpr{
											pos:  position{line: 348, col: 37, offset: 10282},
											name: "Operator",
										},
										&ruleRefExpr{
											pos:  position{line: 348, col: 46, offset: 10291},
											name: "__N",
										},
										&ruleRefExpr{
											pos:  position{line: 348, col: 50, offset: 10295},
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
			pos:  position{line: 368, col: 1, offset: 10901},
			expr: &choiceExpr{
				pos: position{line: 368, col: 9, offset: 10909},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 368, col: 9, offset: 10909},
						name: "BinOpBool",
					},
					&ruleRefExpr{
						pos:  position{line: 368, col: 21, offset: 10921},
						name: "BinOpEquality",
					},
					&ruleRefExpr{
						pos:  position{line: 368, col: 37, offset: 10937},
						name: "BinOpLow",
					},
					&ruleRefExpr{
						pos:  position{line: 368, col: 48, offset: 10948},
						name: "BinOpHigh",
					},
					&ruleRefExpr{
						pos:  position{line: 368, col: 60, offset: 10960},
						name: "BinOpParens",
					},
				},
			},
		},
		{
			name: "BinOpBool",
			pos:  position{line: 370, col: 1, offset: 10973},
			expr: &actionExpr{
				pos: position{line: 370, col: 13, offset: 10985},
				run: (*parser).callonBinOpBool1,
				expr: &seqExpr{
					pos: position{line: 370, col: 13, offset: 10985},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 370, col: 13, offset: 10985},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 370, col: 15, offset: 10987},
							label: "first",
							expr: &ruleRefExpr{
								pos:  position{line: 370, col: 21, offset: 10993},
								name: "BinOpEquality",
							},
						},
						&labeledExpr{
							pos:   position{line: 370, col: 35, offset: 11007},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 370, col: 40, offset: 11012},
								expr: &seqExpr{
									pos: position{line: 370, col: 41, offset: 11013},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 370, col: 41, offset: 11013},
											name: "__N",
										},
										&ruleRefExpr{
											pos:  position{line: 370, col: 45, offset: 11017},
											name: "OperatorBoolean",
										},
										&ruleRefExpr{
											pos:  position{line: 370, col: 61, offset: 11033},
											name: "__N",
										},
										&ruleRefExpr{
											pos:  position{line: 370, col: 65, offset: 11037},
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
			pos:  position{line: 403, col: 1, offset: 11930},
			expr: &actionExpr{
				pos: position{line: 403, col: 17, offset: 11946},
				run: (*parser).callonBinOpEquality1,
				expr: &seqExpr{
					pos: position{line: 403, col: 17, offset: 11946},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 403, col: 17, offset: 11946},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 403, col: 19, offset: 11948},
							label: "first",
							expr: &ruleRefExpr{
								pos:  position{line: 403, col: 25, offset: 11954},
								name: "BinOpLow",
							},
						},
						&labeledExpr{
							pos:   position{line: 403, col: 34, offset: 11963},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 403, col: 39, offset: 11968},
								expr: &seqExpr{
									pos: position{line: 403, col: 40, offset: 11969},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 403, col: 40, offset: 11969},
											name: "__N",
										},
										&ruleRefExpr{
											pos:  position{line: 403, col: 44, offset: 11973},
											name: "OperatorEquality",
										},
										&ruleRefExpr{
											pos:  position{line: 403, col: 61, offset: 11990},
											name: "__N",
										},
										&ruleRefExpr{
											pos:  position{line: 403, col: 65, offset: 11994},
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
			pos:  position{line: 435, col: 1, offset: 12881},
			expr: &actionExpr{
				pos: position{line: 435, col: 12, offset: 12892},
				run: (*parser).callonBinOpLow1,
				expr: &seqExpr{
					pos: position{line: 435, col: 12, offset: 12892},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 435, col: 12, offset: 12892},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 435, col: 14, offset: 12894},
							label: "first",
							expr: &ruleRefExpr{
								pos:  position{line: 435, col: 20, offset: 12900},
								name: "BinOpHigh",
							},
						},
						&labeledExpr{
							pos:   position{line: 435, col: 30, offset: 12910},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 435, col: 35, offset: 12915},
								expr: &seqExpr{
									pos: position{line: 435, col: 36, offset: 12916},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 435, col: 36, offset: 12916},
											name: "__N",
										},
										&ruleRefExpr{
											pos:  position{line: 435, col: 40, offset: 12920},
											name: "OperatorLow",
										},
										&ruleRefExpr{
											pos:  position{line: 435, col: 52, offset: 12932},
											name: "__N",
										},
										&ruleRefExpr{
											pos:  position{line: 435, col: 56, offset: 12936},
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
			pos:  position{line: 467, col: 1, offset: 13824},
			expr: &actionExpr{
				pos: position{line: 467, col: 13, offset: 13836},
				run: (*parser).callonBinOpHigh1,
				expr: &seqExpr{
					pos: position{line: 467, col: 13, offset: 13836},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 467, col: 13, offset: 13836},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 467, col: 15, offset: 13838},
							label: "first",
							expr: &ruleRefExpr{
								pos:  position{line: 467, col: 21, offset: 13844},
								name: "BinOpParens",
							},
						},
						&labeledExpr{
							pos:   position{line: 467, col: 33, offset: 13856},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 467, col: 38, offset: 13861},
								expr: &seqExpr{
									pos: position{line: 467, col: 39, offset: 13862},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 467, col: 39, offset: 13862},
											name: "__N",
										},
										&ruleRefExpr{
											pos:  position{line: 467, col: 43, offset: 13866},
											name: "OperatorHigh",
										},
										&ruleRefExpr{
											pos:  position{line: 467, col: 56, offset: 13879},
											name: "__N",
										},
										&ruleRefExpr{
											pos:  position{line: 467, col: 60, offset: 13883},
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
			pos:  position{line: 498, col: 1, offset: 14772},
			expr: &choiceExpr{
				pos: position{line: 498, col: 15, offset: 14786},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 498, col: 15, offset: 14786},
						run: (*parser).callonBinOpParens2,
						expr: &seqExpr{
							pos: position{line: 498, col: 15, offset: 14786},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 498, col: 15, offset: 14786},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 498, col: 17, offset: 14788},
									val:        "(",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 498, col: 21, offset: 14792},
									name: "_N",
								},
								&labeledExpr{
									pos:   position{line: 498, col: 24, offset: 14795},
									label: "first",
									expr: &ruleRefExpr{
										pos:  position{line: 498, col: 30, offset: 14801},
										name: "BinOp",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 498, col: 36, offset: 14807},
									name: "_N",
								},
								&litMatcher{
									pos:        position{line: 498, col: 39, offset: 14810},
									val:        ")",
									ignoreCase: false,
								},
							},
						},
					},
					&ruleRefExpr{
						pos:  position{line: 501, col: 5, offset: 14933},
						name: "Value",
					},
				},
			},
		},
		{
			name: "Operator",
			pos:  position{line: 503, col: 1, offset: 14940},
			expr: &choiceExpr{
				pos: position{line: 503, col: 12, offset: 14951},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 503, col: 12, offset: 14951},
						name: "OperatorBoolean",
					},
					&ruleRefExpr{
						pos:  position{line: 503, col: 30, offset: 14969},
						name: "OperatorEquality",
					},
					&ruleRefExpr{
						pos:  position{line: 503, col: 49, offset: 14988},
						name: "OperatorHigh",
					},
					&ruleRefExpr{
						pos:  position{line: 503, col: 64, offset: 15003},
						name: "OperatorLow",
					},
				},
			},
		},
		{
			name: "OperatorBoolean",
			pos:  position{line: 505, col: 1, offset: 15016},
			expr: &actionExpr{
				pos: position{line: 505, col: 19, offset: 15034},
				run: (*parser).callonOperatorBoolean1,
				expr: &choiceExpr{
					pos: position{line: 505, col: 21, offset: 15036},
					alternatives: []interface{}{
						&litMatcher{
							pos:        position{line: 505, col: 21, offset: 15036},
							val:        "or",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 505, col: 28, offset: 15043},
							val:        "and",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "OperatorEquality",
			pos:  position{line: 509, col: 1, offset: 15125},
			expr: &actionExpr{
				pos: position{line: 509, col: 20, offset: 15144},
				run: (*parser).callonOperatorEquality1,
				expr: &choiceExpr{
					pos: position{line: 509, col: 22, offset: 15146},
					alternatives: []interface{}{
						&litMatcher{
							pos:        position{line: 509, col: 22, offset: 15146},
							val:        ">=",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 509, col: 29, offset: 15153},
							val:        "<=",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 509, col: 36, offset: 15160},
							val:        "<",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 509, col: 42, offset: 15166},
							val:        ">",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 509, col: 48, offset: 15172},
							val:        "===",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 509, col: 56, offset: 15180},
							val:        "==",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "OperatorHigh",
			pos:  position{line: 513, col: 1, offset: 15259},
			expr: &choiceExpr{
				pos: position{line: 513, col: 16, offset: 15274},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 513, col: 16, offset: 15274},
						run: (*parser).callonOperatorHigh2,
						expr: &choiceExpr{
							pos: position{line: 513, col: 18, offset: 15276},
							alternatives: []interface{}{
								&litMatcher{
									pos:        position{line: 513, col: 18, offset: 15276},
									val:        "/",
									ignoreCase: false,
								},
								&litMatcher{
									pos:        position{line: 513, col: 24, offset: 15282},
									val:        "*",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 516, col: 3, offset: 15365},
						run: (*parser).callonOperatorHigh6,
						expr: &litMatcher{
							pos:        position{line: 516, col: 5, offset: 15367},
							val:        "^",
							ignoreCase: false,
						},
					},
					&actionExpr{
						pos: position{line: 519, col: 3, offset: 15447},
						run: (*parser).callonOperatorHigh8,
						expr: &litMatcher{
							pos:        position{line: 519, col: 3, offset: 15447},
							val:        "++",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "OperatorLow",
			pos:  position{line: 523, col: 1, offset: 15528},
			expr: &actionExpr{
				pos: position{line: 523, col: 15, offset: 15542},
				run: (*parser).callonOperatorLow1,
				expr: &choiceExpr{
					pos: position{line: 523, col: 17, offset: 15544},
					alternatives: []interface{}{
						&litMatcher{
							pos:        position{line: 523, col: 17, offset: 15544},
							val:        "+",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 523, col: 23, offset: 15550},
							val:        "-",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "Value",
			pos:  position{line: 527, col: 1, offset: 15632},
			expr: &choiceExpr{
				pos: position{line: 527, col: 9, offset: 15640},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 527, col: 9, offset: 15640},
						name: "Call",
					},
					&ruleRefExpr{
						pos:  position{line: 527, col: 16, offset: 15647},
						name: "RecordAccess",
					},
					&ruleRefExpr{
						pos:  position{line: 527, col: 31, offset: 15662},
						name: "VariableName",
					},
					&ruleRefExpr{
						pos:  position{line: 527, col: 46, offset: 15677},
						name: "Const",
					},
				},
			},
		},
		{
			name: "Assignable",
			pos:  position{line: 529, col: 1, offset: 15684},
			expr: &choiceExpr{
				pos: position{line: 529, col: 14, offset: 15697},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 529, col: 14, offset: 15697},
						name: "VariableName",
					},
					&ruleRefExpr{
						pos:  position{line: 529, col: 29, offset: 15712},
						name: "Unused",
					},
				},
			},
		},
		{
			name: "Identifier",
			pos:  position{line: 531, col: 1, offset: 15720},
			expr: &choiceExpr{
				pos: position{line: 531, col: 14, offset: 15733},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 531, col: 14, offset: 15733},
						name: "VariableName",
					},
					&ruleRefExpr{
						pos:  position{line: 531, col: 29, offset: 15748},
						name: "ModuleName",
					},
				},
			},
		},
		{
			name: "ArrayLiteral",
			pos:  position{line: 533, col: 1, offset: 15760},
			expr: &actionExpr{
				pos: position{line: 533, col: 16, offset: 15775},
				run: (*parser).callonArrayLiteral1,
				expr: &seqExpr{
					pos: position{line: 533, col: 16, offset: 15775},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 533, col: 16, offset: 15775},
							val:        "[",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 533, col: 20, offset: 15779},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 533, col: 22, offset: 15781},
							label: "first",
							expr: &ruleRefExpr{
								pos:  position{line: 533, col: 28, offset: 15787},
								name: "Expr",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 533, col: 33, offset: 15792},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 533, col: 35, offset: 15794},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 533, col: 40, offset: 15799},
								expr: &seqExpr{
									pos: position{line: 533, col: 41, offset: 15800},
									exprs: []interface{}{
										&litMatcher{
											pos:        position{line: 533, col: 41, offset: 15800},
											val:        ",",
											ignoreCase: false,
										},
										&ruleRefExpr{
											pos:  position{line: 533, col: 45, offset: 15804},
											name: "_",
										},
										&ruleRefExpr{
											pos:  position{line: 533, col: 47, offset: 15806},
											name: "Expr",
										},
										&ruleRefExpr{
											pos:  position{line: 533, col: 52, offset: 15811},
											name: "_",
										},
									},
								},
							},
						},
						&zeroOrOneExpr{
							pos: position{line: 533, col: 56, offset: 15815},
							expr: &litMatcher{
								pos:        position{line: 533, col: 56, offset: 15815},
								val:        ",",
								ignoreCase: false,
							},
						},
						&ruleRefExpr{
							pos:  position{line: 533, col: 61, offset: 15820},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 533, col: 63, offset: 15822},
							val:        "]",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "VariantInstance",
			pos:  position{line: 549, col: 1, offset: 16267},
			expr: &actionExpr{
				pos: position{line: 549, col: 19, offset: 16285},
				run: (*parser).callonVariantInstance1,
				expr: &seqExpr{
					pos: position{line: 549, col: 19, offset: 16285},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 549, col: 19, offset: 16285},
							label: "name",
							expr: &ruleRefExpr{
								pos:  position{line: 549, col: 24, offset: 16290},
								name: "ModuleName",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 549, col: 35, offset: 16301},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 549, col: 37, offset: 16303},
							label: "args",
							expr: &ruleRefExpr{
								pos:  position{line: 549, col: 42, offset: 16308},
								name: "Arguments",
							},
						},
					},
				},
			},
		},
		{
			name: "RecordInstance",
			pos:  position{line: 562, col: 1, offset: 16578},
			expr: &actionExpr{
				pos: position{line: 562, col: 18, offset: 16595},
				run: (*parser).callonRecordInstance1,
				expr: &seqExpr{
					pos: position{line: 562, col: 18, offset: 16595},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 562, col: 18, offset: 16595},
							label: "name",
							expr: &ruleRefExpr{
								pos:  position{line: 562, col: 23, offset: 16600},
								name: "ModuleName",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 562, col: 34, offset: 16611},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 562, col: 36, offset: 16613},
							val:        "{",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 562, col: 40, offset: 16617},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 562, col: 42, offset: 16619},
							label: "firstName",
							expr: &ruleRefExpr{
								pos:  position{line: 562, col: 52, offset: 16629},
								name: "VariableName",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 562, col: 65, offset: 16642},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 562, col: 67, offset: 16644},
							val:        ":",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 562, col: 71, offset: 16648},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 562, col: 73, offset: 16650},
							label: "firstValue",
							expr: &ruleRefExpr{
								pos:  position{line: 562, col: 84, offset: 16661},
								name: "Expr",
							},
						},
						&labeledExpr{
							pos:   position{line: 562, col: 89, offset: 16666},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 562, col: 94, offset: 16671},
								expr: &seqExpr{
									pos: position{line: 562, col: 95, offset: 16672},
									exprs: []interface{}{
										&litMatcher{
											pos:        position{line: 562, col: 95, offset: 16672},
											val:        ",",
											ignoreCase: false,
										},
										&ruleRefExpr{
											pos:  position{line: 562, col: 99, offset: 16676},
											name: "_",
										},
										&ruleRefExpr{
											pos:  position{line: 562, col: 101, offset: 16678},
											name: "VariableName",
										},
										&ruleRefExpr{
											pos:  position{line: 562, col: 114, offset: 16691},
											name: "_",
										},
										&litMatcher{
											pos:        position{line: 562, col: 116, offset: 16693},
											val:        ":",
											ignoreCase: false,
										},
										&ruleRefExpr{
											pos:  position{line: 562, col: 120, offset: 16697},
											name: "_",
										},
										&ruleRefExpr{
											pos:  position{line: 562, col: 122, offset: 16699},
											name: "Expr",
										},
									},
								},
							},
						},
						&litMatcher{
							pos:        position{line: 562, col: 130, offset: 16707},
							val:        "}",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "TypeParameter",
			pos:  position{line: 582, col: 1, offset: 17291},
			expr: &actionExpr{
				pos: position{line: 582, col: 17, offset: 17307},
				run: (*parser).callonTypeParameter1,
				expr: &seqExpr{
					pos: position{line: 582, col: 17, offset: 17307},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 582, col: 17, offset: 17307},
							val:        "'",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 582, col: 22, offset: 17312},
							name: "VariableName",
						},
					},
				},
			},
		},
		{
			name: "VariableName",
			pos:  position{line: 586, col: 1, offset: 17385},
			expr: &actionExpr{
				pos: position{line: 586, col: 16, offset: 17400},
				run: (*parser).callonVariableName1,
				expr: &seqExpr{
					pos: position{line: 586, col: 16, offset: 17400},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 586, col: 16, offset: 17400},
							expr: &ruleRefExpr{
								pos:  position{line: 586, col: 17, offset: 17401},
								name: "Reserved",
							},
						},
						&seqExpr{
							pos: position{line: 586, col: 27, offset: 17411},
							exprs: []interface{}{
								&oneOrMoreExpr{
									pos: position{line: 586, col: 27, offset: 17411},
									expr: &charClassMatcher{
										pos:        position{line: 586, col: 27, offset: 17411},
										val:        "[a-z]",
										ranges:     []rune{'a', 'z'},
										ignoreCase: false,
										inverted:   false,
									},
								},
								&zeroOrMoreExpr{
									pos: position{line: 586, col: 34, offset: 17418},
									expr: &charClassMatcher{
										pos:        position{line: 586, col: 34, offset: 17418},
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
			pos:  position{line: 590, col: 1, offset: 17493},
			expr: &actionExpr{
				pos: position{line: 590, col: 14, offset: 17506},
				run: (*parser).callonModuleName1,
				expr: &seqExpr{
					pos: position{line: 590, col: 15, offset: 17507},
					exprs: []interface{}{
						&oneOrMoreExpr{
							pos: position{line: 590, col: 15, offset: 17507},
							expr: &charClassMatcher{
								pos:        position{line: 590, col: 15, offset: 17507},
								val:        "[A-Z]",
								ranges:     []rune{'A', 'Z'},
								ignoreCase: false,
								inverted:   false,
							},
						},
						&zeroOrMoreExpr{
							pos: position{line: 590, col: 22, offset: 17514},
							expr: &charClassMatcher{
								pos:        position{line: 590, col: 22, offset: 17514},
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
			pos:  position{line: 594, col: 1, offset: 17589},
			expr: &choiceExpr{
				pos: position{line: 594, col: 9, offset: 17597},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 594, col: 9, offset: 17597},
						run: (*parser).callonConst2,
						expr: &seqExpr{
							pos: position{line: 594, col: 9, offset: 17597},
							exprs: []interface{}{
								&zeroOrOneExpr{
									pos: position{line: 594, col: 9, offset: 17597},
									expr: &litMatcher{
										pos:        position{line: 594, col: 9, offset: 17597},
										val:        "-",
										ignoreCase: false,
									},
								},
								&oneOrMoreExpr{
									pos: position{line: 594, col: 14, offset: 17602},
									expr: &charClassMatcher{
										pos:        position{line: 594, col: 14, offset: 17602},
										val:        "[0-9]",
										ranges:     []rune{'0', '9'},
										ignoreCase: false,
										inverted:   false,
									},
								},
								&notExpr{
									pos: position{line: 594, col: 21, offset: 17609},
									expr: &litMatcher{
										pos:        position{line: 594, col: 22, offset: 17610},
										val:        ".",
										ignoreCase: false,
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 601, col: 3, offset: 17785},
						run: (*parser).callonConst10,
						expr: &seqExpr{
							pos: position{line: 601, col: 3, offset: 17785},
							exprs: []interface{}{
								&zeroOrOneExpr{
									pos: position{line: 601, col: 3, offset: 17785},
									expr: &litMatcher{
										pos:        position{line: 601, col: 3, offset: 17785},
										val:        "-",
										ignoreCase: false,
									},
								},
								&oneOrMoreExpr{
									pos: position{line: 601, col: 8, offset: 17790},
									expr: &charClassMatcher{
										pos:        position{line: 601, col: 8, offset: 17790},
										val:        "[0-9]",
										ranges:     []rune{'0', '9'},
										ignoreCase: false,
										inverted:   false,
									},
								},
								&litMatcher{
									pos:        position{line: 601, col: 15, offset: 17797},
									val:        ".",
									ignoreCase: false,
								},
								&oneOrMoreExpr{
									pos: position{line: 601, col: 19, offset: 17801},
									expr: &charClassMatcher{
										pos:        position{line: 601, col: 19, offset: 17801},
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
						pos: position{line: 608, col: 3, offset: 17990},
						run: (*parser).callonConst19,
						expr: &litMatcher{
							pos:        position{line: 608, col: 3, offset: 17990},
							val:        "true",
							ignoreCase: false,
						},
					},
					&actionExpr{
						pos: position{line: 612, col: 3, offset: 18075},
						run: (*parser).callonConst21,
						expr: &litMatcher{
							pos:        position{line: 612, col: 3, offset: 18075},
							val:        "false",
							ignoreCase: false,
						},
					},
					&ruleRefExpr{
						pos:  position{line: 615, col: 3, offset: 18161},
						name: "Unit",
					},
					&ruleRefExpr{
						pos:  position{line: 616, col: 3, offset: 18168},
						name: "StringLiteral",
					},
					&actionExpr{
						pos: position{line: 617, col: 3, offset: 18184},
						run: (*parser).callonConst25,
						expr: &seqExpr{
							pos: position{line: 617, col: 3, offset: 18184},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 617, col: 3, offset: 18184},
									val:        "'",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 617, col: 7, offset: 18188},
									label: "val",
									expr: &seqExpr{
										pos: position{line: 617, col: 12, offset: 18193},
										exprs: []interface{}{
											&notExpr{
												pos: position{line: 617, col: 12, offset: 18193},
												expr: &ruleRefExpr{
													pos:  position{line: 617, col: 13, offset: 18194},
													name: "EscapedChar",
												},
											},
											&anyMatcher{
												line: 617, col: 25, offset: 18206,
											},
										},
									},
								},
								&litMatcher{
									pos:        position{line: 617, col: 28, offset: 18209},
									val:        "'",
									ignoreCase: false,
								},
							},
						},
					},
					&ruleRefExpr{
						pos:  position{line: 619, col: 5, offset: 18301},
						name: "ArrayLiteral",
					},
					&ruleRefExpr{
						pos:  position{line: 619, col: 20, offset: 18316},
						name: "RecordInstance",
					},
					&ruleRefExpr{
						pos:  position{line: 619, col: 37, offset: 18333},
						name: "VariantInstance",
					},
				},
			},
		},
		{
			name: "Unit",
			pos:  position{line: 621, col: 1, offset: 18350},
			expr: &actionExpr{
				pos: position{line: 621, col: 8, offset: 18357},
				run: (*parser).callonUnit1,
				expr: &litMatcher{
					pos:        position{line: 621, col: 8, offset: 18357},
					val:        "()",
					ignoreCase: false,
				},
			},
		},
		{
			name: "StringLiteral",
			pos:  position{line: 625, col: 1, offset: 18420},
			expr: &actionExpr{
				pos: position{line: 625, col: 17, offset: 18436},
				run: (*parser).callonStringLiteral1,
				expr: &seqExpr{
					pos: position{line: 625, col: 17, offset: 18436},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 625, col: 17, offset: 18436},
							val:        "\"",
							ignoreCase: false,
						},
						&zeroOrMoreExpr{
							pos: position{line: 625, col: 21, offset: 18440},
							expr: &seqExpr{
								pos: position{line: 625, col: 22, offset: 18441},
								exprs: []interface{}{
									&notExpr{
										pos: position{line: 625, col: 22, offset: 18441},
										expr: &ruleRefExpr{
											pos:  position{line: 625, col: 23, offset: 18442},
											name: "EscapedChar",
										},
									},
									&anyMatcher{
										line: 625, col: 35, offset: 18454,
									},
								},
							},
						},
						&litMatcher{
							pos:        position{line: 625, col: 39, offset: 18458},
							val:        "\"",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "Unused",
			pos:  position{line: 633, col: 1, offset: 18641},
			expr: &actionExpr{
				pos: position{line: 633, col: 10, offset: 18650},
				run: (*parser).callonUnused1,
				expr: &litMatcher{
					pos:        position{line: 633, col: 11, offset: 18651},
					val:        "_",
					ignoreCase: false,
				},
			},
		},
		{
			name: "Reserved",
			pos:  position{line: 637, col: 1, offset: 18706},
			expr: &seqExpr{
				pos: position{line: 637, col: 12, offset: 18717},
				exprs: []interface{}{
					&choiceExpr{
						pos: position{line: 637, col: 13, offset: 18718},
						alternatives: []interface{}{
							&litMatcher{
								pos:        position{line: 637, col: 13, offset: 18718},
								val:        "let",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 637, col: 21, offset: 18726},
								val:        "if",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 637, col: 28, offset: 18733},
								val:        "else",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 637, col: 37, offset: 18742},
								val:        "extern",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 637, col: 48, offset: 18753},
								val:        "type",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 637, col: 57, offset: 18762},
								val:        "true",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 637, col: 66, offset: 18771},
								val:        "false",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 637, col: 76, offset: 18781},
								val:        "mutable",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 637, col: 88, offset: 18793},
								val:        "type",
								ignoreCase: false,
							},
							&ruleRefExpr{
								pos:  position{line: 637, col: 97, offset: 18802},
								name: "BaseType",
							},
						},
					},
					&notExpr{
						pos: position{line: 637, col: 107, offset: 18812},
						expr: &oneOrMoreExpr{
							pos: position{line: 637, col: 108, offset: 18813},
							expr: &charClassMatcher{
								pos:        position{line: 637, col: 108, offset: 18813},
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
			pos:  position{line: 639, col: 1, offset: 18821},
			expr: &choiceExpr{
				pos: position{line: 639, col: 12, offset: 18832},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 639, col: 12, offset: 18832},
						run: (*parser).callonBaseType2,
						expr: &choiceExpr{
							pos: position{line: 639, col: 14, offset: 18834},
							alternatives: []interface{}{
								&litMatcher{
									pos:        position{line: 639, col: 14, offset: 18834},
									val:        "int64",
									ignoreCase: false,
								},
								&litMatcher{
									pos:        position{line: 639, col: 24, offset: 18844},
									val:        "int",
									ignoreCase: false,
								},
								&litMatcher{
									pos:        position{line: 639, col: 32, offset: 18852},
									val:        "bool",
									ignoreCase: false,
								},
								&litMatcher{
									pos:        position{line: 639, col: 41, offset: 18861},
									val:        "string",
									ignoreCase: false,
								},
								&litMatcher{
									pos:        position{line: 639, col: 52, offset: 18872},
									val:        "byte",
									ignoreCase: false,
								},
								&litMatcher{
									pos:        position{line: 639, col: 61, offset: 18881},
									val:        "rune",
									ignoreCase: false,
								},
								&litMatcher{
									pos:        position{line: 639, col: 70, offset: 18890},
									val:        "float",
									ignoreCase: false,
								},
								&litMatcher{
									pos:        position{line: 639, col: 80, offset: 18900},
									val:        "()",
									ignoreCase: false,
								},
								&litMatcher{
									pos:        position{line: 639, col: 87, offset: 18907},
									val:        "func",
									ignoreCase: false,
								},
							},
						},
					},
					&ruleRefExpr{
						pos:  position{line: 642, col: 3, offset: 19007},
						name: "Unit",
					},
				},
			},
		},
		{
			name: "EscapedChar",
			pos:  position{line: 644, col: 1, offset: 19013},
			expr: &charClassMatcher{
				pos:        position{line: 644, col: 15, offset: 19027},
				val:        "[\\x00-\\x1f\"\\\\]",
				chars:      []rune{'"', '\\'},
				ranges:     []rune{'\x00', '\x1f'},
				ignoreCase: false,
				inverted:   false,
			},
		},
		{
			name: "EscapeSequence",
			pos:  position{line: 646, col: 1, offset: 19043},
			expr: &choiceExpr{
				pos: position{line: 646, col: 18, offset: 19060},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 646, col: 18, offset: 19060},
						name: "SingleCharEscape",
					},
					&ruleRefExpr{
						pos:  position{line: 646, col: 37, offset: 19079},
						name: "UnicodeEscape",
					},
				},
			},
		},
		{
			name: "SingleCharEscape",
			pos:  position{line: 648, col: 1, offset: 19094},
			expr: &charClassMatcher{
				pos:        position{line: 648, col: 20, offset: 19113},
				val:        "[\"\\\\/bfnrt]",
				chars:      []rune{'"', '\\', '/', 'b', 'f', 'n', 'r', 't'},
				ignoreCase: false,
				inverted:   false,
			},
		},
		{
			name: "DecimalDigit",
			pos:  position{line: 650, col: 1, offset: 19126},
			expr: &charClassMatcher{
				pos:        position{line: 650, col: 16, offset: 19141},
				val:        "[0-9]",
				ranges:     []rune{'0', '9'},
				ignoreCase: false,
				inverted:   false,
			},
		},
		{
			name: "NonZeroDecimalDigit",
			pos:  position{line: 652, col: 1, offset: 19148},
			expr: &charClassMatcher{
				pos:        position{line: 652, col: 23, offset: 19170},
				val:        "[1-9]",
				ranges:     []rune{'1', '9'},
				ignoreCase: false,
				inverted:   false,
			},
		},
		{
			name: "HexDigit",
			pos:  position{line: 654, col: 1, offset: 19177},
			expr: &charClassMatcher{
				pos:        position{line: 654, col: 12, offset: 19188},
				val:        "[0-9a-f]i",
				ranges:     []rune{'0', '9', 'a', 'f'},
				ignoreCase: true,
				inverted:   false,
			},
		},
		{
			name:        "__",
			displayName: "\"reqwhitespace\"",
			pos:         position{line: 656, col: 1, offset: 19199},
			expr: &choiceExpr{
				pos: position{line: 656, col: 22, offset: 19220},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 656, col: 22, offset: 19220},
						name: "Comments",
					},
					&oneOrMoreExpr{
						pos: position{line: 656, col: 33, offset: 19231},
						expr: &charClassMatcher{
							pos:        position{line: 656, col: 33, offset: 19231},
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
			displayName: "\"optwhitespace\"",
			pos:         position{line: 658, col: 1, offset: 19243},
			expr: &choiceExpr{
				pos: position{line: 658, col: 21, offset: 19263},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 658, col: 21, offset: 19263},
						name: "Comments",
					},
					&zeroOrMoreExpr{
						pos: position{line: 658, col: 32, offset: 19274},
						expr: &charClassMatcher{
							pos:        position{line: 658, col: 32, offset: 19274},
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
			name:        "__N",
			displayName: "\"singleline_reqwhitepace\"",
			pos:         position{line: 660, col: 1, offset: 19286},
			expr: &oneOrMoreExpr{
				pos: position{line: 660, col: 33, offset: 19318},
				expr: &charClassMatcher{
					pos:        position{line: 660, col: 33, offset: 19318},
					val:        "[ \\t]",
					chars:      []rune{' ', '\t'},
					ignoreCase: false,
					inverted:   false,
				},
			},
		},
		{
			name:        "_N",
			displayName: "\"singleline_optwhitepace\"",
			pos:         position{line: 662, col: 1, offset: 19326},
			expr: &zeroOrMoreExpr{
				pos: position{line: 662, col: 32, offset: 19357},
				expr: &charClassMatcher{
					pos:        position{line: 662, col: 32, offset: 19357},
					val:        "[ \\t]",
					chars:      []rune{' ', '\t'},
					ignoreCase: false,
					inverted:   false,
				},
			},
		},
		{
			name:        "N",
			displayName: "\"newline\"",
			pos:         position{line: 664, col: 1, offset: 19365},
			expr: &choiceExpr{
				pos: position{line: 664, col: 15, offset: 19379},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 664, col: 15, offset: 19379},
						name: "Comments",
					},
					&seqExpr{
						pos: position{line: 664, col: 26, offset: 19390},
						exprs: []interface{}{
							&zeroOrMoreExpr{
								pos: position{line: 664, col: 26, offset: 19390},
								expr: &charClassMatcher{
									pos:        position{line: 664, col: 26, offset: 19390},
									val:        "[ \\r\\t]",
									chars:      []rune{' ', '\r', '\t'},
									ignoreCase: false,
									inverted:   false,
								},
							},
							&litMatcher{
								pos:        position{line: 664, col: 35, offset: 19399},
								val:        "\n",
								ignoreCase: false,
							},
						},
					},
				},
			},
		},
		{
			name: "Comments",
			pos:  position{line: 666, col: 1, offset: 19405},
			expr: &oneOrMoreExpr{
				pos: position{line: 666, col: 12, offset: 19416},
				expr: &ruleRefExpr{
					pos:  position{line: 666, col: 13, offset: 19417},
					name: "Comment",
				},
			},
		},
		{
			name: "Comment",
			pos:  position{line: 668, col: 1, offset: 19428},
			expr: &seqExpr{
				pos: position{line: 668, col: 11, offset: 19438},
				exprs: []interface{}{
					&zeroOrMoreExpr{
						pos: position{line: 668, col: 11, offset: 19438},
						expr: &charClassMatcher{
							pos:        position{line: 668, col: 11, offset: 19438},
							val:        "[ \\r\\n\\t]",
							chars:      []rune{' ', '\r', '\n', '\t'},
							ignoreCase: false,
							inverted:   false,
						},
					},
					&litMatcher{
						pos:        position{line: 668, col: 22, offset: 19449},
						val:        "#",
						ignoreCase: false,
					},
					&zeroOrMoreExpr{
						pos: position{line: 668, col: 26, offset: 19453},
						expr: &seqExpr{
							pos: position{line: 668, col: 27, offset: 19454},
							exprs: []interface{}{
								&notExpr{
									pos: position{line: 668, col: 27, offset: 19454},
									expr: &charClassMatcher{
										pos:        position{line: 668, col: 28, offset: 19455},
										val:        "[\\n]",
										chars:      []rune{'\n'},
										ignoreCase: false,
										inverted:   false,
									},
								},
								&anyMatcher{
									line: 668, col: 33, offset: 19460,
								},
							},
						},
					},
					&andExpr{
						pos: position{line: 668, col: 37, offset: 19464},
						expr: &litMatcher{
							pos:        position{line: 668, col: 38, offset: 19465},
							val:        "\n",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "EOF",
			pos:  position{line: 670, col: 1, offset: 19471},
			expr: &notExpr{
				pos: position{line: 670, col: 7, offset: 19477},
				expr: &anyMatcher{
					line: 670, col: 8, offset: 19478,
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

func (c *current) onExternFunc1(name, importName, args, ret interface{}) (interface{}, error) {

	return ExternFunc{Name: name.(Identifier).StringValue, Import: importName.(BasicAst).StringValue,
		Arguments: args.(Container).Subvalues, ReturnAnnotation: ret.(BasicAst).StringValue}, nil
}

func (p *parser) callonExternFunc1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onExternFunc1(stack["name"], stack["importName"], stack["args"], stack["ret"])
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

func (c *current) onRecordAccess1(record, rest interface{}) (interface{}, error) {
	args := []Identifier{record.(Identifier)}

	vals := rest.([]interface{})
	if len(vals) > 0 {
		restSl := toIfaceSlice(rest)
		for _, v := range restSl {
			// we can get each item in the grammar by index
			restExpr := toIfaceSlice(v)
			v := restExpr[1].(Identifier)
			args = append(args, v)
		}
	}

	return RecordAccess{Identifiers: args}, nil
}

func (p *parser) callonRecordAccess1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onRecordAccess1(stack["record"], stack["rest"])
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
