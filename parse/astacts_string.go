// Code generated by "stringer -type=AstActs"; DO NOT EDIT.

package parse

import (
	"errors"
	"strconv"
)

var _ = errors.New("dummy error")

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[NoAst-0]
	_ = x[AddAst-1]
	_ = x[SubAst-2]
	_ = x[AnchorAst-3]
	_ = x[AnchorFirstAst-4]
	_ = x[AstActsN-5]
}

const _AstActs_name = "NoAstAddAstSubAstAnchorAstAnchorFirstAstAstActsN"

var _AstActs_index = [...]uint8{0, 5, 11, 17, 26, 40, 48}

func (i AstActs) String() string {
	if i < 0 || i >= AstActs(len(_AstActs_index)-1) {
		return "AstActs(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _AstActs_name[_AstActs_index[i]:_AstActs_index[i+1]]
}

func (i *AstActs) FromString(s string) error {
	for j := 0; j < len(_AstActs_index)-1; j++ {
		if s == _AstActs_name[_AstActs_index[j]:_AstActs_index[j+1]] {
			*i = AstActs(j)
			return nil
		}
	}
	return errors.New("String: " + s + " is not a valid option for type: AstActs")
}
