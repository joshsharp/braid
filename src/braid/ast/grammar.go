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
					&ruleRefExpr{
						pos:  position{line: 28, col: 56, offset: 774},
						name: "ExternType",
					},
				},
			},
		},
		{
			name: "Statement",
			pos:  position{line: 30, col: 1, offset: 786},
			expr: &choiceExpr{
				pos: position{line: 30, col: 13, offset: 798},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 30, col: 13, offset: 798},
						name: "FuncDefn",
					},
					&ruleRefExpr{
						pos:  position{line: 30, col: 24, offset: 809},
						name: "Assignment",
					},
					&ruleRefExpr{
						pos:  position{line: 30, col: 37, offset: 822},
						name: "ExprLine",
					},
				},
			},
		},
		{
			name: "ExprLine",
			pos:  position{line: 32, col: 1, offset: 832},
			expr: &actionExpr{
				pos: position{line: 32, col: 12, offset: 843},
				run: (*parser).callonExprLine1,
				expr: &seqExpr{
					pos: position{line: 32, col: 12, offset: 843},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 32, col: 12, offset: 843},
							label: "e",
							expr: &ruleRefExpr{
								pos:  position{line: 32, col: 14, offset: 845},
								name: "Expr",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 32, col: 19, offset: 850},
							name: "N",
						},
					},
				},
			},
		},
		{
			name: "ExternFunc",
			pos:  position{line: 46, col: 1, offset: 1181},
			expr: &actionExpr{
				pos: position{line: 46, col: 14, offset: 1194},
				run: (*parser).callonExternFunc1,
				expr: &seqExpr{
					pos: position{line: 46, col: 14, offset: 1194},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 46, col: 14, offset: 1194},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 46, col: 16, offset: 1196},
							val:        "extern",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 46, col: 25, offset: 1205},
							name: "__N",
						},
						&litMatcher{
							pos:        position{line: 46, col: 29, offset: 1209},
							val:        "func",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 46, col: 36, offset: 1216},
							name: "__N",
						},
						&labeledExpr{
							pos:   position{line: 46, col: 40, offset: 1220},
							label: "name",
							expr: &ruleRefExpr{
								pos:  position{line: 46, col: 45, offset: 1225},
								name: "Identifier",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 46, col: 56, offset: 1236},
							name: "_N",
						},
						&litMatcher{
							pos:        position{line: 46, col: 59, offset: 1239},
							val:        "=",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 46, col: 63, offset: 1243},
							name: "_N",
						},
						&labeledExpr{
							pos:   position{line: 46, col: 66, offset: 1246},
							label: "importName",
							expr: &ruleRefExpr{
								pos:  position{line: 46, col: 77, offset: 1257},
								name: "StringLiteral",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 46, col: 91, offset: 1271},
							name: "_N",
						},
						&labeledExpr{
							pos:   position{line: 46, col: 94, offset: 1274},
							label: "args",
							expr: &ruleRefExpr{
								pos:  position{line: 46, col: 99, offset: 1279},
								name: "ArgsDefn",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 46, col: 108, offset: 1288},
							name: "_N",
						},
						&labeledExpr{
							pos:   position{line: 46, col: 111, offset: 1291},
							label: "ret",
							expr: &ruleRefExpr{
								pos:  position{line: 46, col: 115, offset: 1295},
								name: "BaseType",
							},
						},
					},
				},
			},
		},
		{
			name: "ExternType",
			pos:  position{line: 52, col: 1, offset: 1511},
			expr: &actionExpr{
				pos: position{line: 52, col: 14, offset: 1524},
				run: (*parser).callonExternType1,
				expr: &seqExpr{
					pos: position{line: 52, col: 14, offset: 1524},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 52, col: 14, offset: 1524},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 52, col: 16, offset: 1526},
							val:        "extern",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 52, col: 25, offset: 1535},
							name: "__N",
						},
						&litMatcher{
							pos:        position{line: 52, col: 29, offset: 1539},
							val:        "type",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 52, col: 36, offset: 1546},
							name: "__",
						},
						&labeledExpr{
							pos:   position{line: 52, col: 39, offset: 1549},
							label: "name",
							expr: &ruleRefExpr{
								pos:  position{line: 52, col: 44, offset: 1554},
								name: "ModuleName",
							},
						},
						&labeledExpr{
							pos:   position{line: 52, col: 55, offset: 1565},
							label: "params",
							expr: &zeroOrMoreExpr{
								pos: position{line: 52, col: 62, offset: 1572},
								expr: &seqExpr{
									pos: position{line: 52, col: 63, offset: 1573},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 52, col: 63, offset: 1573},
											name: "__",
										},
										&ruleRefExpr{
											pos:  position{line: 52, col: 66, offset: 1576},
											name: "TypeParameter",
										},
									},
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 52, col: 82, offset: 1592},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 52, col: 84, offset: 1594},
							val:        "=",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 52, col: 88, offset: 1598},
							name: "_",
						},
						&ruleRefExpr{
							pos:  position{line: 53, col: 5, offset: 1604},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 53, col: 7, offset: 1606},
							label: "importName",
							expr: &ruleRefExpr{
								pos:  position{line: 53, col: 18, offset: 1617},
								name: "StringLiteral",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 53, col: 32, offset: 1631},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 53, col: 34, offset: 1633},
							val:        "{",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 53, col: 38, offset: 1637},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 53, col: 40, offset: 1639},
							label: "first",
							expr: &ruleRefExpr{
								pos:  position{line: 53, col: 46, offset: 1645},
								name: "RecordFieldDefn",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 53, col: 62, offset: 1661},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 53, col: 64, offset: 1663},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 53, col: 69, offset: 1668},
								expr: &seqExpr{
									pos: position{line: 53, col: 70, offset: 1669},
									exprs: []interface{}{
										&litMatcher{
											pos:        position{line: 53, col: 70, offset: 1669},
											val:        ",",
											ignoreCase: false,
										},
										&ruleRefExpr{
											pos:  position{line: 53, col: 74, offset: 1673},
											name: "_",
										},
										&ruleRefExpr{
											pos:  position{line: 53, col: 76, offset: 1675},
											name: "RecordFieldDefn",
										},
										&ruleRefExpr{
											pos:  position{line: 53, col: 92, offset: 1691},
											name: "_",
										},
									},
								},
							},
						},
						&zeroOrOneExpr{
							pos: position{line: 53, col: 96, offset: 1695},
							expr: &litMatcher{
								pos:        position{line: 53, col: 96, offset: 1695},
								val:        ",",
								ignoreCase: false,
							},
						},
						&ruleRefExpr{
							pos:  position{line: 53, col: 101, offset: 1700},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 53, col: 103, offset: 1702},
							val:        "}",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 53, col: 107, offset: 1706},
							name: "__",
						},
					},
				},
			},
		},
		{
			name: "TypeDefn",
			pos:  position{line: 72, col: 1, offset: 2254},
			expr: &choiceExpr{
				pos: position{line: 72, col: 12, offset: 2265},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 72, col: 12, offset: 2265},
						run: (*parser).callonTypeDefn2,
						expr: &seqExpr{
							pos: position{line: 72, col: 12, offset: 2265},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 72, col: 12, offset: 2265},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 72, col: 14, offset: 2267},
									val:        "type",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 72, col: 21, offset: 2274},
									name: "__",
								},
								&labeledExpr{
									pos:   position{line: 72, col: 24, offset: 2277},
									label: "name",
									expr: &ruleRefExpr{
										pos:  position{line: 72, col: 29, offset: 2282},
										name: "ModuleName",
									},
								},
								&labeledExpr{
									pos:   position{line: 72, col: 40, offset: 2293},
									label: "params",
									expr: &zeroOrMoreExpr{
										pos: position{line: 72, col: 47, offset: 2300},
										expr: &seqExpr{
											pos: position{line: 72, col: 48, offset: 2301},
											exprs: []interface{}{
												&ruleRefExpr{
													pos:  position{line: 72, col: 48, offset: 2301},
													name: "__",
												},
												&ruleRefExpr{
													pos:  position{line: 72, col: 51, offset: 2304},
													name: "TypeParameter",
												},
											},
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 72, col: 67, offset: 2320},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 72, col: 69, offset: 2322},
									val:        "=",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 72, col: 73, offset: 2326},
									label: "types",
									expr: &oneOrMoreExpr{
										pos: position{line: 72, col: 79, offset: 2332},
										expr: &seqExpr{
											pos: position{line: 72, col: 80, offset: 2333},
											exprs: []interface{}{
												&ruleRefExpr{
													pos:  position{line: 72, col: 80, offset: 2333},
													name: "__",
												},
												&ruleRefExpr{
													pos:  position{line: 72, col: 83, offset: 2336},
													name: "AnyType",
												},
											},
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 72, col: 93, offset: 2346},
									name: "__",
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 91, col: 1, offset: 2842},
						run: (*parser).callonTypeDefn22,
						expr: &seqExpr{
							pos: position{line: 91, col: 1, offset: 2842},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 91, col: 1, offset: 2842},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 91, col: 3, offset: 2844},
									val:        "type",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 91, col: 10, offset: 2851},
									name: "__",
								},
								&labeledExpr{
									pos:   position{line: 91, col: 13, offset: 2854},
									label: "name",
									expr: &ruleRefExpr{
										pos:  position{line: 91, col: 18, offset: 2859},
										name: "ModuleName",
									},
								},
								&labeledExpr{
									pos:   position{line: 91, col: 29, offset: 2870},
									label: "params",
									expr: &zeroOrMoreExpr{
										pos: position{line: 91, col: 36, offset: 2877},
										expr: &seqExpr{
											pos: position{line: 91, col: 37, offset: 2878},
											exprs: []interface{}{
												&ruleRefExpr{
													pos:  position{line: 91, col: 37, offset: 2878},
													name: "__",
												},
												&ruleRefExpr{
													pos:  position{line: 91, col: 40, offset: 2881},
													name: "TypeParameter",
												},
											},
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 91, col: 56, offset: 2897},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 91, col: 58, offset: 2899},
									val:        "=",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 91, col: 62, offset: 2903},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 92, col: 5, offset: 2909},
									val:        "{",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 92, col: 9, offset: 2913},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 92, col: 11, offset: 2915},
									label: "first",
									expr: &ruleRefExpr{
										pos:  position{line: 92, col: 17, offset: 2921},
										name: "RecordFieldDefn",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 92, col: 33, offset: 2937},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 92, col: 35, offset: 2939},
									label: "rest",
									expr: &zeroOrMoreExpr{
										pos: position{line: 92, col: 40, offset: 2944},
										expr: &seqExpr{
											pos: position{line: 92, col: 41, offset: 2945},
											exprs: []interface{}{
												&litMatcher{
													pos:        position{line: 92, col: 41, offset: 2945},
													val:        ",",
													ignoreCase: false,
												},
												&ruleRefExpr{
													pos:  position{line: 92, col: 45, offset: 2949},
													name: "_",
												},
												&ruleRefExpr{
													pos:  position{line: 92, col: 47, offset: 2951},
													name: "RecordFieldDefn",
												},
												&ruleRefExpr{
													pos:  position{line: 92, col: 63, offset: 2967},
													name: "_",
												},
											},
										},
									},
								},
								&zeroOrOneExpr{
									pos: position{line: 92, col: 67, offset: 2971},
									expr: &litMatcher{
										pos:        position{line: 92, col: 67, offset: 2971},
										val:        ",",
										ignoreCase: false,
									},
								},
								&ruleRefExpr{
									pos:  position{line: 92, col: 72, offset: 2976},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 92, col: 74, offset: 2978},
									val:        "}",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 92, col: 78, offset: 2982},
									name: "__",
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 110, col: 1, offset: 3467},
						run: (*parser).callonTypeDefn54,
						expr: &seqExpr{
							pos: position{line: 110, col: 1, offset: 3467},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 110, col: 1, offset: 3467},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 110, col: 3, offset: 3469},
									val:        "type",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 110, col: 10, offset: 3476},
									name: "__",
								},
								&labeledExpr{
									pos:   position{line: 110, col: 13, offset: 3479},
									label: "name",
									expr: &ruleRefExpr{
										pos:  position{line: 110, col: 18, offset: 3484},
										name: "ModuleName",
									},
								},
								&labeledExpr{
									pos:   position{line: 110, col: 29, offset: 3495},
									label: "params",
									expr: &zeroOrMoreExpr{
										pos: position{line: 110, col: 36, offset: 3502},
										expr: &seqExpr{
											pos: position{line: 110, col: 37, offset: 3503},
											exprs: []interface{}{
												&ruleRefExpr{
													pos:  position{line: 110, col: 37, offset: 3503},
													name: "__",
												},
												&ruleRefExpr{
													pos:  position{line: 110, col: 40, offset: 3506},
													name: "TypeParameter",
												},
											},
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 110, col: 56, offset: 3522},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 110, col: 58, offset: 3524},
									val:        "=",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 110, col: 62, offset: 3528},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 110, col: 64, offset: 3530},
									label: "rest",
									expr: &oneOrMoreExpr{
										pos: position{line: 110, col: 69, offset: 3535},
										expr: &ruleRefExpr{
											pos:  position{line: 110, col: 70, offset: 3536},
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
			pos:  position{line: 125, col: 1, offset: 3943},
			expr: &actionExpr{
				pos: position{line: 125, col: 19, offset: 3961},
				run: (*parser).callonRecordFieldDefn1,
				expr: &seqExpr{
					pos: position{line: 125, col: 19, offset: 3961},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 125, col: 19, offset: 3961},
							label: "name",
							expr: &ruleRefExpr{
								pos:  position{line: 125, col: 24, offset: 3966},
								name: "VariableName",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 125, col: 37, offset: 3979},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 125, col: 39, offset: 3981},
							val:        ":",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 125, col: 43, offset: 3985},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 125, col: 45, offset: 3987},
							label: "t",
							expr: &ruleRefExpr{
								pos:  position{line: 125, col: 48, offset: 3990},
								name: "AnyType",
							},
						},
					},
				},
			},
		},
		{
			name: "VariantConstructor",
			pos:  position{line: 129, col: 1, offset: 4084},
			expr: &choiceExpr{
				pos: position{line: 129, col: 22, offset: 4105},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 129, col: 22, offset: 4105},
						run: (*parser).callonVariantConstructor2,
						expr: &seqExpr{
							pos: position{line: 129, col: 22, offset: 4105},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 129, col: 22, offset: 4105},
									val:        "|",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 129, col: 26, offset: 4109},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 129, col: 28, offset: 4111},
									label: "name",
									expr: &ruleRefExpr{
										pos:  position{line: 129, col: 33, offset: 4116},
										name: "ModuleName",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 129, col: 44, offset: 4127},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 129, col: 46, offset: 4129},
									val:        "{",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 129, col: 50, offset: 4133},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 129, col: 52, offset: 4135},
									label: "first",
									expr: &ruleRefExpr{
										pos:  position{line: 129, col: 58, offset: 4141},
										name: "RecordFieldDefn",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 129, col: 74, offset: 4157},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 129, col: 76, offset: 4159},
									label: "rest",
									expr: &zeroOrMoreExpr{
										pos: position{line: 129, col: 81, offset: 4164},
										expr: &seqExpr{
											pos: position{line: 129, col: 82, offset: 4165},
											exprs: []interface{}{
												&litMatcher{
													pos:        position{line: 129, col: 82, offset: 4165},
													val:        ",",
													ignoreCase: false,
												},
												&ruleRefExpr{
													pos:  position{line: 129, col: 86, offset: 4169},
													name: "_",
												},
												&ruleRefExpr{
													pos:  position{line: 129, col: 88, offset: 4171},
													name: "RecordFieldDefn",
												},
												&ruleRefExpr{
													pos:  position{line: 129, col: 104, offset: 4187},
													name: "_",
												},
											},
										},
									},
								},
								&zeroOrOneExpr{
									pos: position{line: 129, col: 108, offset: 4191},
									expr: &litMatcher{
										pos:        position{line: 129, col: 108, offset: 4191},
										val:        ",",
										ignoreCase: false,
									},
								},
								&ruleRefExpr{
									pos:  position{line: 129, col: 113, offset: 4196},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 129, col: 115, offset: 4198},
									val:        "}",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 129, col: 119, offset: 4202},
									name: "__",
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 148, col: 1, offset: 4807},
						run: (*parser).callonVariantConstructor26,
						expr: &seqExpr{
							pos: position{line: 148, col: 1, offset: 4807},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 148, col: 1, offset: 4807},
									val:        "|",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 148, col: 5, offset: 4811},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 148, col: 7, offset: 4813},
									label: "name",
									expr: &ruleRefExpr{
										pos:  position{line: 148, col: 12, offset: 4818},
										name: "ModuleName",
									},
								},
								&labeledExpr{
									pos:   position{line: 148, col: 23, offset: 4829},
									label: "rest",
									expr: &zeroOrMoreExpr{
										pos: position{line: 148, col: 28, offset: 4834},
										expr: &seqExpr{
											pos: position{line: 148, col: 29, offset: 4835},
											exprs: []interface{}{
												&ruleRefExpr{
													pos:  position{line: 148, col: 29, offset: 4835},
													name: "__",
												},
												&ruleRefExpr{
													pos:  position{line: 148, col: 32, offset: 4838},
													name: "AnyType",
												},
											},
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 148, col: 42, offset: 4848},
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
			pos:  position{line: 165, col: 1, offset: 5285},
			expr: &choiceExpr{
				pos: position{line: 165, col: 11, offset: 5295},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 165, col: 11, offset: 5295},
						name: "BaseType",
					},
					&ruleRefExpr{
						pos:  position{line: 165, col: 22, offset: 5306},
						name: "TypeParameter",
					},
				},
			},
		},
		{
			name: "Assignment",
			pos:  position{line: 167, col: 1, offset: 5321},
			expr: &choiceExpr{
				pos: position{line: 167, col: 14, offset: 5334},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 167, col: 14, offset: 5334},
						run: (*parser).callonAssignment2,
						expr: &seqExpr{
							pos: position{line: 167, col: 14, offset: 5334},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 167, col: 14, offset: 5334},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 167, col: 16, offset: 5336},
									val:        "let",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 167, col: 22, offset: 5342},
									name: "__N",
								},
								&labeledExpr{
									pos:   position{line: 167, col: 26, offset: 5346},
									label: "i",
									expr: &ruleRefExpr{
										pos:  position{line: 167, col: 28, offset: 5348},
										name: "Assignable",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 167, col: 39, offset: 5359},
									name: "_N",
								},
								&litMatcher{
									pos:        position{line: 167, col: 42, offset: 5362},
									val:        "=",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 167, col: 46, offset: 5366},
									name: "_N",
								},
								&labeledExpr{
									pos:   position{line: 167, col: 49, offset: 5369},
									label: "expr",
									expr: &ruleRefExpr{
										pos:  position{line: 167, col: 54, offset: 5374},
										name: "Expr",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 167, col: 59, offset: 5379},
									name: "N",
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 173, col: 1, offset: 5498},
						run: (*parser).callonAssignment15,
						expr: &seqExpr{
							pos: position{line: 173, col: 1, offset: 5498},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 173, col: 1, offset: 5498},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 173, col: 3, offset: 5500},
									val:        "let",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 173, col: 9, offset: 5506},
									name: "__N",
								},
								&notExpr{
									pos: position{line: 173, col: 13, offset: 5510},
									expr: &ruleRefExpr{
										pos:  position{line: 173, col: 14, offset: 5511},
										name: "Assignable",
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 177, col: 1, offset: 5609},
						run: (*parser).callonAssignment22,
						expr: &seqExpr{
							pos: position{line: 177, col: 1, offset: 5609},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 177, col: 1, offset: 5609},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 177, col: 3, offset: 5611},
									val:        "let",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 177, col: 9, offset: 5617},
									name: "__N",
								},
								&labeledExpr{
									pos:   position{line: 177, col: 13, offset: 5621},
									label: "i",
									expr: &ruleRefExpr{
										pos:  position{line: 177, col: 15, offset: 5623},
										name: "Assignable",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 177, col: 26, offset: 5634},
									name: "_N",
								},
								&notExpr{
									pos: position{line: 177, col: 29, offset: 5637},
									expr: &litMatcher{
										pos:        position{line: 177, col: 30, offset: 5638},
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
			pos:  position{line: 181, col: 1, offset: 5732},
			expr: &actionExpr{
				pos: position{line: 181, col: 12, offset: 5743},
				run: (*parser).callonFuncDefn1,
				expr: &seqExpr{
					pos: position{line: 181, col: 12, offset: 5743},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 181, col: 12, offset: 5743},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 181, col: 14, offset: 5745},
							val:        "let",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 181, col: 20, offset: 5751},
							name: "__N",
						},
						&labeledExpr{
							pos:   position{line: 181, col: 24, offset: 5755},
							label: "i",
							expr: &ruleRefExpr{
								pos:  position{line: 181, col: 26, offset: 5757},
								name: "VariableName",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 181, col: 39, offset: 5770},
							name: "_N",
						},
						&litMatcher{
							pos:        position{line: 181, col: 42, offset: 5773},
							val:        "=",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 181, col: 46, offset: 5777},
							name: "_N",
						},
						&labeledExpr{
							pos:   position{line: 181, col: 49, offset: 5780},
							label: "ids",
							expr: &zeroOrOneExpr{
								pos: position{line: 181, col: 53, offset: 5784},
								expr: &seqExpr{
									pos: position{line: 181, col: 54, offset: 5785},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 181, col: 54, offset: 5785},
											name: "ArgsDefn",
										},
										&ruleRefExpr{
											pos:  position{line: 181, col: 63, offset: 5794},
											name: "_",
										},
									},
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 181, col: 67, offset: 5798},
							label: "ret",
							expr: &zeroOrOneExpr{
								pos: position{line: 181, col: 71, offset: 5802},
								expr: &seqExpr{
									pos: position{line: 181, col: 72, offset: 5803},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 181, col: 72, offset: 5803},
											name: "AnyType",
										},
										&ruleRefExpr{
											pos:  position{line: 181, col: 80, offset: 5811},
											name: "_",
										},
									},
								},
							},
						},
						&litMatcher{
							pos:        position{line: 181, col: 84, offset: 5815},
							val:        "{",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 181, col: 88, offset: 5819},
							name: "__",
						},
						&labeledExpr{
							pos:   position{line: 181, col: 91, offset: 5822},
							label: "statements",
							expr: &oneOrMoreExpr{
								pos: position{line: 181, col: 102, offset: 5833},
								expr: &ruleRefExpr{
									pos:  position{line: 181, col: 103, offset: 5834},
									name: "Statement",
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 181, col: 115, offset: 5846},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 181, col: 117, offset: 5848},
							val:        "}",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 181, col: 121, offset: 5852},
							name: "_",
						},
					},
				},
			},
		},
		{
			name: "Expr",
			pos:  position{line: 207, col: 1, offset: 6462},
			expr: &actionExpr{
				pos: position{line: 207, col: 8, offset: 6469},
				run: (*parser).callonExpr1,
				expr: &labeledExpr{
					pos:   position{line: 207, col: 8, offset: 6469},
					label: "ex",
					expr: &choiceExpr{
						pos: position{line: 207, col: 12, offset: 6473},
						alternatives: []interface{}{
							&ruleRefExpr{
								pos:  position{line: 207, col: 12, offset: 6473},
								name: "IfExpr",
							},
							&ruleRefExpr{
								pos:  position{line: 207, col: 21, offset: 6482},
								name: "CompoundExpr",
							},
						},
					},
				},
			},
		},
		{
			name: "IfExpr",
			pos:  position{line: 213, col: 1, offset: 6599},
			expr: &choiceExpr{
				pos: position{line: 213, col: 10, offset: 6608},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 213, col: 10, offset: 6608},
						run: (*parser).callonIfExpr2,
						expr: &seqExpr{
							pos: position{line: 213, col: 10, offset: 6608},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 213, col: 10, offset: 6608},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 213, col: 12, offset: 6610},
									val:        "if",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 213, col: 17, offset: 6615},
									name: "__N",
								},
								&labeledExpr{
									pos:   position{line: 213, col: 21, offset: 6619},
									label: "expr",
									expr: &ruleRefExpr{
										pos:  position{line: 213, col: 26, offset: 6624},
										name: "BinOpBool",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 213, col: 36, offset: 6634},
									name: "_N",
								},
								&litMatcher{
									pos:        position{line: 213, col: 39, offset: 6637},
									val:        "{",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 213, col: 43, offset: 6641},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 213, col: 45, offset: 6643},
									label: "thens",
									expr: &oneOrMoreExpr{
										pos: position{line: 213, col: 51, offset: 6649},
										expr: &ruleRefExpr{
											pos:  position{line: 213, col: 52, offset: 6650},
											name: "Statement",
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 213, col: 64, offset: 6662},
									name: "__",
								},
								&litMatcher{
									pos:        position{line: 213, col: 67, offset: 6665},
									val:        "}",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 213, col: 71, offset: 6669},
									name: "_N",
								},
								&litMatcher{
									pos:        position{line: 213, col: 74, offset: 6672},
									val:        "else",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 213, col: 81, offset: 6679},
									name: "_N",
								},
								&litMatcher{
									pos:        position{line: 213, col: 84, offset: 6682},
									val:        "{",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 213, col: 88, offset: 6686},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 213, col: 90, offset: 6688},
									label: "elses",
									expr: &oneOrMoreExpr{
										pos: position{line: 213, col: 96, offset: 6694},
										expr: &ruleRefExpr{
											pos:  position{line: 213, col: 97, offset: 6695},
											name: "Statement",
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 213, col: 109, offset: 6707},
									name: "__",
								},
								&litMatcher{
									pos:        position{line: 213, col: 112, offset: 6710},
									val:        "}",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 232, col: 1, offset: 7213},
						run: (*parser).callonIfExpr27,
						expr: &seqExpr{
							pos: position{line: 232, col: 1, offset: 7213},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 232, col: 1, offset: 7213},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 232, col: 3, offset: 7215},
									val:        "if",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 232, col: 8, offset: 7220},
									name: "__N",
								},
								&labeledExpr{
									pos:   position{line: 232, col: 12, offset: 7224},
									label: "expr",
									expr: &ruleRefExpr{
										pos:  position{line: 232, col: 17, offset: 7229},
										name: "BinOpBool",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 232, col: 27, offset: 7239},
									name: "_N",
								},
								&litMatcher{
									pos:        position{line: 232, col: 30, offset: 7242},
									val:        "{",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 232, col: 34, offset: 7246},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 232, col: 36, offset: 7248},
									label: "thens",
									expr: &oneOrMoreExpr{
										pos: position{line: 232, col: 42, offset: 7254},
										expr: &ruleRefExpr{
											pos:  position{line: 232, col: 43, offset: 7255},
											name: "Statement",
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 232, col: 55, offset: 7267},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 232, col: 57, offset: 7269},
									val:        "}",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 232, col: 61, offset: 7273},
									name: "_N",
								},
								&litMatcher{
									pos:        position{line: 232, col: 64, offset: 7276},
									val:        "else",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 232, col: 71, offset: 7283},
									label: "elseifs",
									expr: &ruleRefExpr{
										pos:  position{line: 232, col: 79, offset: 7291},
										name: "IfExpr",
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 244, col: 1, offset: 7621},
						run: (*parser).callonIfExpr46,
						expr: &seqExpr{
							pos: position{line: 244, col: 1, offset: 7621},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 244, col: 1, offset: 7621},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 244, col: 3, offset: 7623},
									val:        "if",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 244, col: 8, offset: 7628},
									name: "__N",
								},
								&labeledExpr{
									pos:   position{line: 244, col: 12, offset: 7632},
									label: "expr",
									expr: &ruleRefExpr{
										pos:  position{line: 244, col: 17, offset: 7637},
										name: "BinOpBool",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 244, col: 27, offset: 7647},
									name: "_N",
								},
								&litMatcher{
									pos:        position{line: 244, col: 30, offset: 7650},
									val:        "{",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 244, col: 34, offset: 7654},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 244, col: 36, offset: 7656},
									label: "thens",
									expr: &oneOrMoreExpr{
										pos: position{line: 244, col: 42, offset: 7662},
										expr: &ruleRefExpr{
											pos:  position{line: 244, col: 43, offset: 7663},
											name: "Statement",
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 244, col: 55, offset: 7675},
									name: "__",
								},
								&litMatcher{
									pos:        position{line: 244, col: 58, offset: 7678},
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
			pos:  position{line: 256, col: 1, offset: 7976},
			expr: &choiceExpr{
				pos: position{line: 256, col: 8, offset: 7983},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 256, col: 8, offset: 7983},
						run: (*parser).callonCall2,
						expr: &seqExpr{
							pos: position{line: 256, col: 8, offset: 7983},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 256, col: 8, offset: 7983},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 256, col: 10, offset: 7985},
									label: "module",
									expr: &ruleRefExpr{
										pos:  position{line: 256, col: 17, offset: 7992},
										name: "ModuleName",
									},
								},
								&litMatcher{
									pos:        position{line: 256, col: 28, offset: 8003},
									val:        ".",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 256, col: 32, offset: 8007},
									label: "fn",
									expr: &ruleRefExpr{
										pos:  position{line: 256, col: 35, offset: 8010},
										name: "VariableName",
									},
								},
								&labeledExpr{
									pos:   position{line: 256, col: 48, offset: 8023},
									label: "args",
									expr: &ruleRefExpr{
										pos:  position{line: 256, col: 53, offset: 8028},
										name: "Arguments",
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 270, col: 1, offset: 8352},
						run: (*parser).callonCall12,
						expr: &seqExpr{
							pos: position{line: 270, col: 1, offset: 8352},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 270, col: 1, offset: 8352},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 270, col: 3, offset: 8354},
									label: "fn",
									expr: &ruleRefExpr{
										pos:  position{line: 270, col: 6, offset: 8357},
										name: "VariableName",
									},
								},
								&labeledExpr{
									pos:   position{line: 270, col: 19, offset: 8370},
									label: "args",
									expr: &ruleRefExpr{
										pos:  position{line: 270, col: 24, offset: 8375},
										name: "Arguments",
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 284, col: 1, offset: 8692},
						run: (*parser).callonCall19,
						expr: &seqExpr{
							pos: position{line: 284, col: 1, offset: 8692},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 284, col: 1, offset: 8692},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 284, col: 3, offset: 8694},
									label: "fn",
									expr: &ruleRefExpr{
										pos:  position{line: 284, col: 6, offset: 8697},
										name: "VariableName",
									},
								},
								&litMatcher{
									pos:        position{line: 284, col: 19, offset: 8710},
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
			pos:  position{line: 291, col: 1, offset: 8881},
			expr: &actionExpr{
				pos: position{line: 291, col: 16, offset: 8896},
				run: (*parser).callonRecordAccess1,
				expr: &seqExpr{
					pos: position{line: 291, col: 16, offset: 8896},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 291, col: 16, offset: 8896},
							label: "record",
							expr: &ruleRefExpr{
								pos:  position{line: 291, col: 23, offset: 8903},
								name: "VariableName",
							},
						},
						&labeledExpr{
							pos:   position{line: 291, col: 36, offset: 8916},
							label: "rest",
							expr: &oneOrMoreExpr{
								pos: position{line: 291, col: 41, offset: 8921},
								expr: &seqExpr{
									pos: position{line: 291, col: 42, offset: 8922},
									exprs: []interface{}{
										&litMatcher{
											pos:        position{line: 291, col: 42, offset: 8922},
											val:        ".",
											ignoreCase: false,
										},
										&ruleRefExpr{
											pos:  position{line: 291, col: 46, offset: 8926},
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
			pos:  position{line: 308, col: 1, offset: 9363},
			expr: &actionExpr{
				pos: position{line: 308, col: 12, offset: 9374},
				run: (*parser).callonArgsDefn1,
				expr: &seqExpr{
					pos: position{line: 308, col: 12, offset: 9374},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 308, col: 12, offset: 9374},
							val:        "(",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 308, col: 16, offset: 9378},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 308, col: 18, offset: 9380},
							label: "argument",
							expr: &ruleRefExpr{
								pos:  position{line: 308, col: 27, offset: 9389},
								name: "ArgDefn",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 308, col: 35, offset: 9397},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 308, col: 37, offset: 9399},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 308, col: 42, offset: 9404},
								expr: &seqExpr{
									pos: position{line: 308, col: 43, offset: 9405},
									exprs: []interface{}{
										&litMatcher{
											pos:        position{line: 308, col: 43, offset: 9405},
											val:        ",",
											ignoreCase: false,
										},
										&ruleRefExpr{
											pos:  position{line: 308, col: 47, offset: 9409},
											name: "_",
										},
										&ruleRefExpr{
											pos:  position{line: 308, col: 49, offset: 9411},
											name: "ArgDefn",
										},
									},
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 308, col: 59, offset: 9421},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 308, col: 61, offset: 9423},
							val:        ")",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "ArgDefn",
			pos:  position{line: 326, col: 1, offset: 9845},
			expr: &actionExpr{
				pos: position{line: 326, col: 11, offset: 9855},
				run: (*parser).callonArgDefn1,
				expr: &seqExpr{
					pos: position{line: 326, col: 11, offset: 9855},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 326, col: 11, offset: 9855},
							label: "name",
							expr: &ruleRefExpr{
								pos:  position{line: 326, col: 16, offset: 9860},
								name: "Identifier",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 326, col: 27, offset: 9871},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 326, col: 29, offset: 9873},
							label: "anno",
							expr: &zeroOrOneExpr{
								pos: position{line: 326, col: 34, offset: 9878},
								expr: &seqExpr{
									pos: position{line: 326, col: 35, offset: 9879},
									exprs: []interface{}{
										&litMatcher{
											pos:        position{line: 326, col: 35, offset: 9879},
											val:        ":",
											ignoreCase: false,
										},
										&ruleRefExpr{
											pos:  position{line: 326, col: 39, offset: 9883},
											name: "_",
										},
										&ruleRefExpr{
											pos:  position{line: 326, col: 41, offset: 9885},
											name: "AnyType",
										},
									},
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 326, col: 52, offset: 9896},
							name: "_",
						},
					},
				},
			},
		},
		{
			name: "Arguments",
			pos:  position{line: 346, col: 1, offset: 10384},
			expr: &choiceExpr{
				pos: position{line: 346, col: 13, offset: 10396},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 346, col: 13, offset: 10396},
						run: (*parser).callonArguments2,
						expr: &seqExpr{
							pos: position{line: 346, col: 13, offset: 10396},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 346, col: 13, offset: 10396},
									val:        "(",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 346, col: 17, offset: 10400},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 346, col: 19, offset: 10402},
									label: "argument",
									expr: &ruleRefExpr{
										pos:  position{line: 346, col: 28, offset: 10411},
										name: "BinOpParens",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 346, col: 40, offset: 10423},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 346, col: 42, offset: 10425},
									label: "rest",
									expr: &zeroOrMoreExpr{
										pos: position{line: 346, col: 47, offset: 10430},
										expr: &seqExpr{
											pos: position{line: 346, col: 48, offset: 10431},
											exprs: []interface{}{
												&litMatcher{
													pos:        position{line: 346, col: 48, offset: 10431},
													val:        ",",
													ignoreCase: false,
												},
												&ruleRefExpr{
													pos:  position{line: 346, col: 52, offset: 10435},
													name: "_",
												},
												&ruleRefExpr{
													pos:  position{line: 346, col: 54, offset: 10437},
													name: "BinOpParens",
												},
											},
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 346, col: 68, offset: 10451},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 346, col: 70, offset: 10453},
									val:        ")",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 363, col: 1, offset: 10875},
						run: (*parser).callonArguments17,
						expr: &seqExpr{
							pos: position{line: 363, col: 1, offset: 10875},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 363, col: 1, offset: 10875},
									val:        "(",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 363, col: 5, offset: 10879},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 363, col: 7, offset: 10881},
									label: "argument",
									expr: &ruleRefExpr{
										pos:  position{line: 363, col: 16, offset: 10890},
										name: "Expr",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 363, col: 21, offset: 10895},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 363, col: 23, offset: 10897},
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
			pos:  position{line: 368, col: 1, offset: 11002},
			expr: &actionExpr{
				pos: position{line: 368, col: 16, offset: 11017},
				run: (*parser).callonCompoundExpr1,
				expr: &seqExpr{
					pos: position{line: 368, col: 16, offset: 11017},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 368, col: 16, offset: 11017},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 368, col: 18, offset: 11019},
							label: "op",
							expr: &ruleRefExpr{
								pos:  position{line: 368, col: 21, offset: 11022},
								name: "BinOp",
							},
						},
						&labeledExpr{
							pos:   position{line: 368, col: 27, offset: 11028},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 368, col: 32, offset: 11033},
								expr: &seqExpr{
									pos: position{line: 368, col: 33, offset: 11034},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 368, col: 33, offset: 11034},
											name: "__N",
										},
										&ruleRefExpr{
											pos:  position{line: 368, col: 37, offset: 11038},
											name: "Operator",
										},
										&ruleRefExpr{
											pos:  position{line: 368, col: 46, offset: 11047},
											name: "__N",
										},
										&ruleRefExpr{
											pos:  position{line: 368, col: 50, offset: 11051},
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
			pos:  position{line: 388, col: 1, offset: 11657},
			expr: &choiceExpr{
				pos: position{line: 388, col: 9, offset: 11665},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 388, col: 9, offset: 11665},
						name: "BinOpBool",
					},
					&ruleRefExpr{
						pos:  position{line: 388, col: 21, offset: 11677},
						name: "BinOpEquality",
					},
					&ruleRefExpr{
						pos:  position{line: 388, col: 37, offset: 11693},
						name: "BinOpLow",
					},
					&ruleRefExpr{
						pos:  position{line: 388, col: 48, offset: 11704},
						name: "BinOpHigh",
					},
					&ruleRefExpr{
						pos:  position{line: 388, col: 60, offset: 11716},
						name: "BinOpParens",
					},
				},
			},
		},
		{
			name: "BinOpBool",
			pos:  position{line: 390, col: 1, offset: 11729},
			expr: &actionExpr{
				pos: position{line: 390, col: 13, offset: 11741},
				run: (*parser).callonBinOpBool1,
				expr: &seqExpr{
					pos: position{line: 390, col: 13, offset: 11741},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 390, col: 13, offset: 11741},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 390, col: 15, offset: 11743},
							label: "first",
							expr: &ruleRefExpr{
								pos:  position{line: 390, col: 21, offset: 11749},
								name: "BinOpEquality",
							},
						},
						&labeledExpr{
							pos:   position{line: 390, col: 35, offset: 11763},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 390, col: 40, offset: 11768},
								expr: &seqExpr{
									pos: position{line: 390, col: 41, offset: 11769},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 390, col: 41, offset: 11769},
											name: "__N",
										},
										&ruleRefExpr{
											pos:  position{line: 390, col: 45, offset: 11773},
											name: "OperatorBoolean",
										},
										&ruleRefExpr{
											pos:  position{line: 390, col: 61, offset: 11789},
											name: "__N",
										},
										&ruleRefExpr{
											pos:  position{line: 390, col: 65, offset: 11793},
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
			pos:  position{line: 423, col: 1, offset: 12686},
			expr: &actionExpr{
				pos: position{line: 423, col: 17, offset: 12702},
				run: (*parser).callonBinOpEquality1,
				expr: &seqExpr{
					pos: position{line: 423, col: 17, offset: 12702},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 423, col: 17, offset: 12702},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 423, col: 19, offset: 12704},
							label: "first",
							expr: &ruleRefExpr{
								pos:  position{line: 423, col: 25, offset: 12710},
								name: "BinOpLow",
							},
						},
						&labeledExpr{
							pos:   position{line: 423, col: 34, offset: 12719},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 423, col: 39, offset: 12724},
								expr: &seqExpr{
									pos: position{line: 423, col: 40, offset: 12725},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 423, col: 40, offset: 12725},
											name: "__N",
										},
										&ruleRefExpr{
											pos:  position{line: 423, col: 44, offset: 12729},
											name: "OperatorEquality",
										},
										&ruleRefExpr{
											pos:  position{line: 423, col: 61, offset: 12746},
											name: "__N",
										},
										&ruleRefExpr{
											pos:  position{line: 423, col: 65, offset: 12750},
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
			pos:  position{line: 455, col: 1, offset: 13637},
			expr: &actionExpr{
				pos: position{line: 455, col: 12, offset: 13648},
				run: (*parser).callonBinOpLow1,
				expr: &seqExpr{
					pos: position{line: 455, col: 12, offset: 13648},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 455, col: 12, offset: 13648},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 455, col: 14, offset: 13650},
							label: "first",
							expr: &ruleRefExpr{
								pos:  position{line: 455, col: 20, offset: 13656},
								name: "BinOpHigh",
							},
						},
						&labeledExpr{
							pos:   position{line: 455, col: 30, offset: 13666},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 455, col: 35, offset: 13671},
								expr: &seqExpr{
									pos: position{line: 455, col: 36, offset: 13672},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 455, col: 36, offset: 13672},
											name: "__N",
										},
										&ruleRefExpr{
											pos:  position{line: 455, col: 40, offset: 13676},
											name: "OperatorLow",
										},
										&ruleRefExpr{
											pos:  position{line: 455, col: 52, offset: 13688},
											name: "__N",
										},
										&ruleRefExpr{
											pos:  position{line: 455, col: 56, offset: 13692},
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
			pos:  position{line: 487, col: 1, offset: 14580},
			expr: &actionExpr{
				pos: position{line: 487, col: 13, offset: 14592},
				run: (*parser).callonBinOpHigh1,
				expr: &seqExpr{
					pos: position{line: 487, col: 13, offset: 14592},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 487, col: 13, offset: 14592},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 487, col: 15, offset: 14594},
							label: "first",
							expr: &ruleRefExpr{
								pos:  position{line: 487, col: 21, offset: 14600},
								name: "BinOpParens",
							},
						},
						&labeledExpr{
							pos:   position{line: 487, col: 33, offset: 14612},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 487, col: 38, offset: 14617},
								expr: &seqExpr{
									pos: position{line: 487, col: 39, offset: 14618},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 487, col: 39, offset: 14618},
											name: "__N",
										},
										&ruleRefExpr{
											pos:  position{line: 487, col: 43, offset: 14622},
											name: "OperatorHigh",
										},
										&ruleRefExpr{
											pos:  position{line: 487, col: 56, offset: 14635},
											name: "__N",
										},
										&ruleRefExpr{
											pos:  position{line: 487, col: 60, offset: 14639},
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
			pos:  position{line: 518, col: 1, offset: 15528},
			expr: &choiceExpr{
				pos: position{line: 518, col: 15, offset: 15542},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 518, col: 15, offset: 15542},
						run: (*parser).callonBinOpParens2,
						expr: &seqExpr{
							pos: position{line: 518, col: 15, offset: 15542},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 518, col: 15, offset: 15542},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 518, col: 17, offset: 15544},
									val:        "(",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 518, col: 21, offset: 15548},
									name: "_N",
								},
								&labeledExpr{
									pos:   position{line: 518, col: 24, offset: 15551},
									label: "first",
									expr: &ruleRefExpr{
										pos:  position{line: 518, col: 30, offset: 15557},
										name: "BinOp",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 518, col: 36, offset: 15563},
									name: "_N",
								},
								&litMatcher{
									pos:        position{line: 518, col: 39, offset: 15566},
									val:        ")",
									ignoreCase: false,
								},
							},
						},
					},
					&ruleRefExpr{
						pos:  position{line: 521, col: 5, offset: 15689},
						name: "Value",
					},
				},
			},
		},
		{
			name: "Operator",
			pos:  position{line: 523, col: 1, offset: 15696},
			expr: &choiceExpr{
				pos: position{line: 523, col: 12, offset: 15707},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 523, col: 12, offset: 15707},
						name: "OperatorBoolean",
					},
					&ruleRefExpr{
						pos:  position{line: 523, col: 30, offset: 15725},
						name: "OperatorEquality",
					},
					&ruleRefExpr{
						pos:  position{line: 523, col: 49, offset: 15744},
						name: "OperatorHigh",
					},
					&ruleRefExpr{
						pos:  position{line: 523, col: 64, offset: 15759},
						name: "OperatorLow",
					},
				},
			},
		},
		{
			name: "OperatorBoolean",
			pos:  position{line: 525, col: 1, offset: 15772},
			expr: &actionExpr{
				pos: position{line: 525, col: 19, offset: 15790},
				run: (*parser).callonOperatorBoolean1,
				expr: &choiceExpr{
					pos: position{line: 525, col: 21, offset: 15792},
					alternatives: []interface{}{
						&litMatcher{
							pos:        position{line: 525, col: 21, offset: 15792},
							val:        "or",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 525, col: 28, offset: 15799},
							val:        "and",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "OperatorEquality",
			pos:  position{line: 529, col: 1, offset: 15881},
			expr: &actionExpr{
				pos: position{line: 529, col: 20, offset: 15900},
				run: (*parser).callonOperatorEquality1,
				expr: &choiceExpr{
					pos: position{line: 529, col: 22, offset: 15902},
					alternatives: []interface{}{
						&litMatcher{
							pos:        position{line: 529, col: 22, offset: 15902},
							val:        ">=",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 529, col: 29, offset: 15909},
							val:        "<=",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 529, col: 36, offset: 15916},
							val:        "<",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 529, col: 42, offset: 15922},
							val:        ">",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 529, col: 48, offset: 15928},
							val:        "===",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 529, col: 56, offset: 15936},
							val:        "==",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "OperatorHigh",
			pos:  position{line: 533, col: 1, offset: 16015},
			expr: &choiceExpr{
				pos: position{line: 533, col: 16, offset: 16030},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 533, col: 16, offset: 16030},
						run: (*parser).callonOperatorHigh2,
						expr: &choiceExpr{
							pos: position{line: 533, col: 18, offset: 16032},
							alternatives: []interface{}{
								&litMatcher{
									pos:        position{line: 533, col: 18, offset: 16032},
									val:        "/",
									ignoreCase: false,
								},
								&litMatcher{
									pos:        position{line: 533, col: 24, offset: 16038},
									val:        "*",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 536, col: 3, offset: 16121},
						run: (*parser).callonOperatorHigh6,
						expr: &litMatcher{
							pos:        position{line: 536, col: 5, offset: 16123},
							val:        "^",
							ignoreCase: false,
						},
					},
					&actionExpr{
						pos: position{line: 539, col: 3, offset: 16203},
						run: (*parser).callonOperatorHigh8,
						expr: &litMatcher{
							pos:        position{line: 539, col: 3, offset: 16203},
							val:        "++",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "OperatorLow",
			pos:  position{line: 543, col: 1, offset: 16284},
			expr: &actionExpr{
				pos: position{line: 543, col: 15, offset: 16298},
				run: (*parser).callonOperatorLow1,
				expr: &choiceExpr{
					pos: position{line: 543, col: 17, offset: 16300},
					alternatives: []interface{}{
						&litMatcher{
							pos:        position{line: 543, col: 17, offset: 16300},
							val:        "+",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 543, col: 23, offset: 16306},
							val:        "-",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "Value",
			pos:  position{line: 547, col: 1, offset: 16388},
			expr: &choiceExpr{
				pos: position{line: 547, col: 9, offset: 16396},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 547, col: 9, offset: 16396},
						name: "Call",
					},
					&ruleRefExpr{
						pos:  position{line: 547, col: 16, offset: 16403},
						name: "RecordAccess",
					},
					&ruleRefExpr{
						pos:  position{line: 547, col: 31, offset: 16418},
						name: "VariableName",
					},
					&ruleRefExpr{
						pos:  position{line: 547, col: 46, offset: 16433},
						name: "Const",
					},
				},
			},
		},
		{
			name: "Assignable",
			pos:  position{line: 549, col: 1, offset: 16440},
			expr: &choiceExpr{
				pos: position{line: 549, col: 14, offset: 16453},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 549, col: 14, offset: 16453},
						name: "VariableName",
					},
					&ruleRefExpr{
						pos:  position{line: 549, col: 29, offset: 16468},
						name: "Unused",
					},
				},
			},
		},
		{
			name: "Identifier",
			pos:  position{line: 551, col: 1, offset: 16476},
			expr: &choiceExpr{
				pos: position{line: 551, col: 14, offset: 16489},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 551, col: 14, offset: 16489},
						name: "VariableName",
					},
					&ruleRefExpr{
						pos:  position{line: 551, col: 29, offset: 16504},
						name: "ModuleName",
					},
				},
			},
		},
		{
			name: "ArrayLiteral",
			pos:  position{line: 553, col: 1, offset: 16516},
			expr: &actionExpr{
				pos: position{line: 553, col: 16, offset: 16531},
				run: (*parser).callonArrayLiteral1,
				expr: &seqExpr{
					pos: position{line: 553, col: 16, offset: 16531},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 553, col: 16, offset: 16531},
							val:        "[",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 553, col: 20, offset: 16535},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 553, col: 22, offset: 16537},
							label: "first",
							expr: &ruleRefExpr{
								pos:  position{line: 553, col: 28, offset: 16543},
								name: "Expr",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 553, col: 33, offset: 16548},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 553, col: 35, offset: 16550},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 553, col: 40, offset: 16555},
								expr: &seqExpr{
									pos: position{line: 553, col: 41, offset: 16556},
									exprs: []interface{}{
										&litMatcher{
											pos:        position{line: 553, col: 41, offset: 16556},
											val:        ",",
											ignoreCase: false,
										},
										&ruleRefExpr{
											pos:  position{line: 553, col: 45, offset: 16560},
											name: "_",
										},
										&ruleRefExpr{
											pos:  position{line: 553, col: 47, offset: 16562},
											name: "Expr",
										},
										&ruleRefExpr{
											pos:  position{line: 553, col: 52, offset: 16567},
											name: "_",
										},
									},
								},
							},
						},
						&zeroOrOneExpr{
							pos: position{line: 553, col: 56, offset: 16571},
							expr: &litMatcher{
								pos:        position{line: 553, col: 56, offset: 16571},
								val:        ",",
								ignoreCase: false,
							},
						},
						&ruleRefExpr{
							pos:  position{line: 553, col: 61, offset: 16576},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 553, col: 63, offset: 16578},
							val:        "]",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "VariantInstance",
			pos:  position{line: 569, col: 1, offset: 17023},
			expr: &actionExpr{
				pos: position{line: 569, col: 19, offset: 17041},
				run: (*parser).callonVariantInstance1,
				expr: &seqExpr{
					pos: position{line: 569, col: 19, offset: 17041},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 569, col: 19, offset: 17041},
							label: "name",
							expr: &ruleRefExpr{
								pos:  position{line: 569, col: 24, offset: 17046},
								name: "ModuleName",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 569, col: 35, offset: 17057},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 569, col: 37, offset: 17059},
							label: "args",
							expr: &ruleRefExpr{
								pos:  position{line: 569, col: 42, offset: 17064},
								name: "Arguments",
							},
						},
					},
				},
			},
		},
		{
			name: "RecordInstance",
			pos:  position{line: 582, col: 1, offset: 17334},
			expr: &actionExpr{
				pos: position{line: 582, col: 18, offset: 17351},
				run: (*parser).callonRecordInstance1,
				expr: &seqExpr{
					pos: position{line: 582, col: 18, offset: 17351},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 582, col: 18, offset: 17351},
							label: "name",
							expr: &ruleRefExpr{
								pos:  position{line: 582, col: 23, offset: 17356},
								name: "ModuleName",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 582, col: 34, offset: 17367},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 582, col: 36, offset: 17369},
							val:        "{",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 582, col: 40, offset: 17373},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 582, col: 42, offset: 17375},
							label: "firstName",
							expr: &ruleRefExpr{
								pos:  position{line: 582, col: 52, offset: 17385},
								name: "VariableName",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 582, col: 65, offset: 17398},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 582, col: 67, offset: 17400},
							val:        ":",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 582, col: 71, offset: 17404},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 582, col: 73, offset: 17406},
							label: "firstValue",
							expr: &ruleRefExpr{
								pos:  position{line: 582, col: 84, offset: 17417},
								name: "Expr",
							},
						},
						&labeledExpr{
							pos:   position{line: 582, col: 89, offset: 17422},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 582, col: 94, offset: 17427},
								expr: &seqExpr{
									pos: position{line: 582, col: 95, offset: 17428},
									exprs: []interface{}{
										&litMatcher{
											pos:        position{line: 582, col: 95, offset: 17428},
											val:        ",",
											ignoreCase: false,
										},
										&ruleRefExpr{
											pos:  position{line: 582, col: 99, offset: 17432},
											name: "_",
										},
										&ruleRefExpr{
											pos:  position{line: 582, col: 101, offset: 17434},
											name: "VariableName",
										},
										&ruleRefExpr{
											pos:  position{line: 582, col: 114, offset: 17447},
											name: "_",
										},
										&litMatcher{
											pos:        position{line: 582, col: 116, offset: 17449},
											val:        ":",
											ignoreCase: false,
										},
										&ruleRefExpr{
											pos:  position{line: 582, col: 120, offset: 17453},
											name: "_",
										},
										&ruleRefExpr{
											pos:  position{line: 582, col: 122, offset: 17455},
											name: "Expr",
										},
									},
								},
							},
						},
						&litMatcher{
							pos:        position{line: 582, col: 130, offset: 17463},
							val:        "}",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "TypeParameter",
			pos:  position{line: 602, col: 1, offset: 18047},
			expr: &actionExpr{
				pos: position{line: 602, col: 17, offset: 18063},
				run: (*parser).callonTypeParameter1,
				expr: &seqExpr{
					pos: position{line: 602, col: 17, offset: 18063},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 602, col: 17, offset: 18063},
							val:        "'",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 602, col: 22, offset: 18068},
							name: "VariableName",
						},
					},
				},
			},
		},
		{
			name: "VariableName",
			pos:  position{line: 606, col: 1, offset: 18141},
			expr: &actionExpr{
				pos: position{line: 606, col: 16, offset: 18156},
				run: (*parser).callonVariableName1,
				expr: &seqExpr{
					pos: position{line: 606, col: 16, offset: 18156},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 606, col: 16, offset: 18156},
							expr: &ruleRefExpr{
								pos:  position{line: 606, col: 17, offset: 18157},
								name: "Reserved",
							},
						},
						&seqExpr{
							pos: position{line: 606, col: 27, offset: 18167},
							exprs: []interface{}{
								&oneOrMoreExpr{
									pos: position{line: 606, col: 27, offset: 18167},
									expr: &charClassMatcher{
										pos:        position{line: 606, col: 27, offset: 18167},
										val:        "[a-z]",
										ranges:     []rune{'a', 'z'},
										ignoreCase: false,
										inverted:   false,
									},
								},
								&zeroOrMoreExpr{
									pos: position{line: 606, col: 34, offset: 18174},
									expr: &charClassMatcher{
										pos:        position{line: 606, col: 34, offset: 18174},
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
			pos:  position{line: 610, col: 1, offset: 18249},
			expr: &actionExpr{
				pos: position{line: 610, col: 14, offset: 18262},
				run: (*parser).callonModuleName1,
				expr: &seqExpr{
					pos: position{line: 610, col: 15, offset: 18263},
					exprs: []interface{}{
						&oneOrMoreExpr{
							pos: position{line: 610, col: 15, offset: 18263},
							expr: &charClassMatcher{
								pos:        position{line: 610, col: 15, offset: 18263},
								val:        "[A-Z]",
								ranges:     []rune{'A', 'Z'},
								ignoreCase: false,
								inverted:   false,
							},
						},
						&zeroOrMoreExpr{
							pos: position{line: 610, col: 22, offset: 18270},
							expr: &charClassMatcher{
								pos:        position{line: 610, col: 22, offset: 18270},
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
			pos:  position{line: 614, col: 1, offset: 18345},
			expr: &choiceExpr{
				pos: position{line: 614, col: 9, offset: 18353},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 614, col: 9, offset: 18353},
						run: (*parser).callonConst2,
						expr: &seqExpr{
							pos: position{line: 614, col: 9, offset: 18353},
							exprs: []interface{}{
								&zeroOrOneExpr{
									pos: position{line: 614, col: 9, offset: 18353},
									expr: &litMatcher{
										pos:        position{line: 614, col: 9, offset: 18353},
										val:        "-",
										ignoreCase: false,
									},
								},
								&oneOrMoreExpr{
									pos: position{line: 614, col: 14, offset: 18358},
									expr: &charClassMatcher{
										pos:        position{line: 614, col: 14, offset: 18358},
										val:        "[0-9]",
										ranges:     []rune{'0', '9'},
										ignoreCase: false,
										inverted:   false,
									},
								},
								&notExpr{
									pos: position{line: 614, col: 21, offset: 18365},
									expr: &litMatcher{
										pos:        position{line: 614, col: 22, offset: 18366},
										val:        ".",
										ignoreCase: false,
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 621, col: 3, offset: 18541},
						run: (*parser).callonConst10,
						expr: &seqExpr{
							pos: position{line: 621, col: 3, offset: 18541},
							exprs: []interface{}{
								&zeroOrOneExpr{
									pos: position{line: 621, col: 3, offset: 18541},
									expr: &litMatcher{
										pos:        position{line: 621, col: 3, offset: 18541},
										val:        "-",
										ignoreCase: false,
									},
								},
								&oneOrMoreExpr{
									pos: position{line: 621, col: 8, offset: 18546},
									expr: &charClassMatcher{
										pos:        position{line: 621, col: 8, offset: 18546},
										val:        "[0-9]",
										ranges:     []rune{'0', '9'},
										ignoreCase: false,
										inverted:   false,
									},
								},
								&litMatcher{
									pos:        position{line: 621, col: 15, offset: 18553},
									val:        ".",
									ignoreCase: false,
								},
								&oneOrMoreExpr{
									pos: position{line: 621, col: 19, offset: 18557},
									expr: &charClassMatcher{
										pos:        position{line: 621, col: 19, offset: 18557},
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
						pos: position{line: 628, col: 3, offset: 18746},
						run: (*parser).callonConst19,
						expr: &litMatcher{
							pos:        position{line: 628, col: 3, offset: 18746},
							val:        "true",
							ignoreCase: false,
						},
					},
					&actionExpr{
						pos: position{line: 632, col: 3, offset: 18831},
						run: (*parser).callonConst21,
						expr: &litMatcher{
							pos:        position{line: 632, col: 3, offset: 18831},
							val:        "false",
							ignoreCase: false,
						},
					},
					&ruleRefExpr{
						pos:  position{line: 635, col: 3, offset: 18917},
						name: "Unit",
					},
					&ruleRefExpr{
						pos:  position{line: 636, col: 3, offset: 18924},
						name: "StringLiteral",
					},
					&actionExpr{
						pos: position{line: 637, col: 3, offset: 18940},
						run: (*parser).callonConst25,
						expr: &seqExpr{
							pos: position{line: 637, col: 3, offset: 18940},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 637, col: 3, offset: 18940},
									val:        "'",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 637, col: 7, offset: 18944},
									label: "val",
									expr: &seqExpr{
										pos: position{line: 637, col: 12, offset: 18949},
										exprs: []interface{}{
											&notExpr{
												pos: position{line: 637, col: 12, offset: 18949},
												expr: &ruleRefExpr{
													pos:  position{line: 637, col: 13, offset: 18950},
													name: "EscapedChar",
												},
											},
											&anyMatcher{
												line: 637, col: 25, offset: 18962,
											},
										},
									},
								},
								&litMatcher{
									pos:        position{line: 637, col: 28, offset: 18965},
									val:        "'",
									ignoreCase: false,
								},
							},
						},
					},
					&ruleRefExpr{
						pos:  position{line: 639, col: 5, offset: 19057},
						name: "ArrayLiteral",
					},
					&ruleRefExpr{
						pos:  position{line: 639, col: 20, offset: 19072},
						name: "RecordInstance",
					},
					&ruleRefExpr{
						pos:  position{line: 639, col: 37, offset: 19089},
						name: "VariantInstance",
					},
				},
			},
		},
		{
			name: "Unit",
			pos:  position{line: 641, col: 1, offset: 19106},
			expr: &actionExpr{
				pos: position{line: 641, col: 8, offset: 19113},
				run: (*parser).callonUnit1,
				expr: &litMatcher{
					pos:        position{line: 641, col: 8, offset: 19113},
					val:        "()",
					ignoreCase: false,
				},
			},
		},
		{
			name: "StringLiteral",
			pos:  position{line: 645, col: 1, offset: 19176},
			expr: &actionExpr{
				pos: position{line: 645, col: 17, offset: 19192},
				run: (*parser).callonStringLiteral1,
				expr: &seqExpr{
					pos: position{line: 645, col: 17, offset: 19192},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 645, col: 17, offset: 19192},
							val:        "\"",
							ignoreCase: false,
						},
						&zeroOrMoreExpr{
							pos: position{line: 645, col: 21, offset: 19196},
							expr: &seqExpr{
								pos: position{line: 645, col: 22, offset: 19197},
								exprs: []interface{}{
									&notExpr{
										pos: position{line: 645, col: 22, offset: 19197},
										expr: &ruleRefExpr{
											pos:  position{line: 645, col: 23, offset: 19198},
											name: "EscapedChar",
										},
									},
									&anyMatcher{
										line: 645, col: 35, offset: 19210,
									},
								},
							},
						},
						&litMatcher{
							pos:        position{line: 645, col: 39, offset: 19214},
							val:        "\"",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "Unused",
			pos:  position{line: 653, col: 1, offset: 19397},
			expr: &actionExpr{
				pos: position{line: 653, col: 10, offset: 19406},
				run: (*parser).callonUnused1,
				expr: &litMatcher{
					pos:        position{line: 653, col: 11, offset: 19407},
					val:        "_",
					ignoreCase: false,
				},
			},
		},
		{
			name: "Reserved",
			pos:  position{line: 657, col: 1, offset: 19462},
			expr: &seqExpr{
				pos: position{line: 657, col: 12, offset: 19473},
				exprs: []interface{}{
					&choiceExpr{
						pos: position{line: 657, col: 13, offset: 19474},
						alternatives: []interface{}{
							&litMatcher{
								pos:        position{line: 657, col: 13, offset: 19474},
								val:        "let",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 657, col: 21, offset: 19482},
								val:        "if",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 657, col: 28, offset: 19489},
								val:        "else",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 657, col: 37, offset: 19498},
								val:        "extern",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 657, col: 48, offset: 19509},
								val:        "type",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 657, col: 57, offset: 19518},
								val:        "true",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 657, col: 66, offset: 19527},
								val:        "false",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 657, col: 76, offset: 19537},
								val:        "mutable",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 657, col: 88, offset: 19549},
								val:        "type",
								ignoreCase: false,
							},
							&ruleRefExpr{
								pos:  position{line: 657, col: 97, offset: 19558},
								name: "BaseType",
							},
						},
					},
					&notExpr{
						pos: position{line: 657, col: 107, offset: 19568},
						expr: &oneOrMoreExpr{
							pos: position{line: 657, col: 108, offset: 19569},
							expr: &charClassMatcher{
								pos:        position{line: 657, col: 108, offset: 19569},
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
			pos:  position{line: 659, col: 1, offset: 19577},
			expr: &choiceExpr{
				pos: position{line: 659, col: 12, offset: 19588},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 659, col: 12, offset: 19588},
						run: (*parser).callonBaseType2,
						expr: &choiceExpr{
							pos: position{line: 659, col: 14, offset: 19590},
							alternatives: []interface{}{
								&litMatcher{
									pos:        position{line: 659, col: 14, offset: 19590},
									val:        "int64",
									ignoreCase: false,
								},
								&litMatcher{
									pos:        position{line: 659, col: 24, offset: 19600},
									val:        "int",
									ignoreCase: false,
								},
								&litMatcher{
									pos:        position{line: 659, col: 32, offset: 19608},
									val:        "bool",
									ignoreCase: false,
								},
								&litMatcher{
									pos:        position{line: 659, col: 41, offset: 19617},
									val:        "string",
									ignoreCase: false,
								},
								&litMatcher{
									pos:        position{line: 659, col: 52, offset: 19628},
									val:        "byte",
									ignoreCase: false,
								},
								&litMatcher{
									pos:        position{line: 659, col: 61, offset: 19637},
									val:        "rune",
									ignoreCase: false,
								},
								&litMatcher{
									pos:        position{line: 659, col: 70, offset: 19646},
									val:        "float",
									ignoreCase: false,
								},
								&litMatcher{
									pos:        position{line: 659, col: 80, offset: 19656},
									val:        "()",
									ignoreCase: false,
								},
								&litMatcher{
									pos:        position{line: 659, col: 87, offset: 19663},
									val:        "func",
									ignoreCase: false,
								},
							},
						},
					},
					&ruleRefExpr{
						pos:  position{line: 662, col: 3, offset: 19763},
						name: "Unit",
					},
				},
			},
		},
		{
			name: "EscapedChar",
			pos:  position{line: 664, col: 1, offset: 19769},
			expr: &charClassMatcher{
				pos:        position{line: 664, col: 15, offset: 19783},
				val:        "[\\x00-\\x1f\"\\\\]",
				chars:      []rune{'"', '\\'},
				ranges:     []rune{'\x00', '\x1f'},
				ignoreCase: false,
				inverted:   false,
			},
		},
		{
			name: "EscapeSequence",
			pos:  position{line: 666, col: 1, offset: 19799},
			expr: &choiceExpr{
				pos: position{line: 666, col: 18, offset: 19816},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 666, col: 18, offset: 19816},
						name: "SingleCharEscape",
					},
					&ruleRefExpr{
						pos:  position{line: 666, col: 37, offset: 19835},
						name: "UnicodeEscape",
					},
				},
			},
		},
		{
			name: "SingleCharEscape",
			pos:  position{line: 668, col: 1, offset: 19850},
			expr: &charClassMatcher{
				pos:        position{line: 668, col: 20, offset: 19869},
				val:        "[\"\\\\/bfnrt]",
				chars:      []rune{'"', '\\', '/', 'b', 'f', 'n', 'r', 't'},
				ignoreCase: false,
				inverted:   false,
			},
		},
		{
			name: "DecimalDigit",
			pos:  position{line: 670, col: 1, offset: 19882},
			expr: &charClassMatcher{
				pos:        position{line: 670, col: 16, offset: 19897},
				val:        "[0-9]",
				ranges:     []rune{'0', '9'},
				ignoreCase: false,
				inverted:   false,
			},
		},
		{
			name: "NonZeroDecimalDigit",
			pos:  position{line: 672, col: 1, offset: 19904},
			expr: &charClassMatcher{
				pos:        position{line: 672, col: 23, offset: 19926},
				val:        "[1-9]",
				ranges:     []rune{'1', '9'},
				ignoreCase: false,
				inverted:   false,
			},
		},
		{
			name: "HexDigit",
			pos:  position{line: 674, col: 1, offset: 19933},
			expr: &charClassMatcher{
				pos:        position{line: 674, col: 12, offset: 19944},
				val:        "[0-9a-f]i",
				ranges:     []rune{'0', '9', 'a', 'f'},
				ignoreCase: true,
				inverted:   false,
			},
		},
		{
			name:        "__",
			displayName: "\"reqwhitespace\"",
			pos:         position{line: 676, col: 1, offset: 19955},
			expr: &choiceExpr{
				pos: position{line: 676, col: 22, offset: 19976},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 676, col: 22, offset: 19976},
						name: "Comments",
					},
					&oneOrMoreExpr{
						pos: position{line: 676, col: 33, offset: 19987},
						expr: &charClassMatcher{
							pos:        position{line: 676, col: 33, offset: 19987},
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
			pos:         position{line: 678, col: 1, offset: 19999},
			expr: &choiceExpr{
				pos: position{line: 678, col: 21, offset: 20019},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 678, col: 21, offset: 20019},
						name: "Comments",
					},
					&zeroOrMoreExpr{
						pos: position{line: 678, col: 32, offset: 20030},
						expr: &charClassMatcher{
							pos:        position{line: 678, col: 32, offset: 20030},
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
			pos:         position{line: 680, col: 1, offset: 20042},
			expr: &oneOrMoreExpr{
				pos: position{line: 680, col: 33, offset: 20074},
				expr: &charClassMatcher{
					pos:        position{line: 680, col: 33, offset: 20074},
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
			pos:         position{line: 682, col: 1, offset: 20082},
			expr: &zeroOrMoreExpr{
				pos: position{line: 682, col: 32, offset: 20113},
				expr: &charClassMatcher{
					pos:        position{line: 682, col: 32, offset: 20113},
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
			pos:         position{line: 684, col: 1, offset: 20121},
			expr: &choiceExpr{
				pos: position{line: 684, col: 15, offset: 20135},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 684, col: 15, offset: 20135},
						name: "Comments",
					},
					&seqExpr{
						pos: position{line: 684, col: 26, offset: 20146},
						exprs: []interface{}{
							&zeroOrMoreExpr{
								pos: position{line: 684, col: 26, offset: 20146},
								expr: &charClassMatcher{
									pos:        position{line: 684, col: 26, offset: 20146},
									val:        "[ \\r\\t]",
									chars:      []rune{' ', '\r', '\t'},
									ignoreCase: false,
									inverted:   false,
								},
							},
							&litMatcher{
								pos:        position{line: 684, col: 35, offset: 20155},
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
			pos:  position{line: 686, col: 1, offset: 20161},
			expr: &oneOrMoreExpr{
				pos: position{line: 686, col: 12, offset: 20172},
				expr: &ruleRefExpr{
					pos:  position{line: 686, col: 13, offset: 20173},
					name: "Comment",
				},
			},
		},
		{
			name: "Comment",
			pos:  position{line: 688, col: 1, offset: 20184},
			expr: &choiceExpr{
				pos: position{line: 688, col: 11, offset: 20194},
				alternatives: []interface{}{
					&seqExpr{
						pos: position{line: 688, col: 11, offset: 20194},
						exprs: []interface{}{
							&zeroOrMoreExpr{
								pos: position{line: 688, col: 11, offset: 20194},
								expr: &charClassMatcher{
									pos:        position{line: 688, col: 11, offset: 20194},
									val:        "[ \\r\\n\\t]",
									chars:      []rune{' ', '\r', '\n', '\t'},
									ignoreCase: false,
									inverted:   false,
								},
							},
							&litMatcher{
								pos:        position{line: 688, col: 22, offset: 20205},
								val:        "//",
								ignoreCase: false,
							},
							&zeroOrMoreExpr{
								pos: position{line: 688, col: 27, offset: 20210},
								expr: &seqExpr{
									pos: position{line: 688, col: 28, offset: 20211},
									exprs: []interface{}{
										&notExpr{
											pos: position{line: 688, col: 28, offset: 20211},
											expr: &charClassMatcher{
												pos:        position{line: 688, col: 29, offset: 20212},
												val:        "[\\n]",
												chars:      []rune{'\n'},
												ignoreCase: false,
												inverted:   false,
											},
										},
										&anyMatcher{
											line: 688, col: 34, offset: 20217,
										},
									},
								},
							},
							&andExpr{
								pos: position{line: 688, col: 38, offset: 20221},
								expr: &litMatcher{
									pos:        position{line: 688, col: 39, offset: 20222},
									val:        "\n",
									ignoreCase: false,
								},
							},
						},
					},
					&seqExpr{
						pos: position{line: 688, col: 46, offset: 20229},
						exprs: []interface{}{
							&zeroOrMoreExpr{
								pos: position{line: 688, col: 46, offset: 20229},
								expr: &charClassMatcher{
									pos:        position{line: 688, col: 46, offset: 20229},
									val:        "[ \\r\\n\\t]",
									chars:      []rune{' ', '\r', '\n', '\t'},
									ignoreCase: false,
									inverted:   false,
								},
							},
							&litMatcher{
								pos:        position{line: 688, col: 57, offset: 20240},
								val:        "/*",
								ignoreCase: false,
							},
							&zeroOrMoreExpr{
								pos: position{line: 688, col: 62, offset: 20245},
								expr: &seqExpr{
									pos: position{line: 688, col: 63, offset: 20246},
									exprs: []interface{}{
										&notExpr{
											pos: position{line: 688, col: 63, offset: 20246},
											expr: &litMatcher{
												pos:        position{line: 688, col: 64, offset: 20247},
												val:        "*/",
												ignoreCase: false,
											},
										},
										&anyMatcher{
											line: 688, col: 69, offset: 20252,
										},
									},
								},
							},
							&litMatcher{
								pos:        position{line: 688, col: 73, offset: 20256},
								val:        "*/",
								ignoreCase: false,
							},
							&zeroOrMoreExpr{
								pos: position{line: 688, col: 78, offset: 20261},
								expr: &charClassMatcher{
									pos:        position{line: 688, col: 78, offset: 20261},
									val:        "[\\r]",
									chars:      []rune{'\r'},
									ignoreCase: false,
									inverted:   false,
								},
							},
							&litMatcher{
								pos:        position{line: 688, col: 84, offset: 20267},
								val:        "\n",
								ignoreCase: false,
							},
						},
					},
				},
			},
		},
		{
			name: "EOF",
			pos:  position{line: 690, col: 1, offset: 20273},
			expr: &notExpr{
				pos: position{line: 690, col: 7, offset: 20279},
				expr: &anyMatcher{
					line: 690, col: 8, offset: 20280,
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

func (c *current) onExternType1(name, params, importName, first, rest interface{}) (interface{}, error) {
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

	return ExternRecordType{Name: name.(Identifier).StringValue, Import: importName.(BasicAst).StringValue,
		Fields: fields}, nil
}

func (p *parser) callonExternType1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onExternType1(stack["name"], stack["params"], stack["importName"], stack["first"], stack["rest"])
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
