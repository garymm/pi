// Code generated by "stringer -type=Matches"; DO NOT EDIT.

package lex

import (
	"errors"
	"strconv"
)

var _ = errors.New("dummy error")

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[String-0]
	_ = x[StrName-1]
	_ = x[Letter-2]
	_ = x[Digit-3]
	_ = x[WhiteSpace-4]
	_ = x[CurState-5]
	_ = x[AnyRune-6]
	_ = x[MatchesN-7]
}

const _Matches_name = "StringStrNameLetterDigitWhiteSpaceCurStateAnyRuneMatchesN"

var _Matches_index = [...]uint8{0, 6, 13, 19, 24, 34, 42, 49, 57}

func (i Matches) String() string {
	if i < 0 || i >= Matches(len(_Matches_index)-1) {
		return "Matches(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Matches_name[_Matches_index[i]:_Matches_index[i+1]]
}

func (i *Matches) FromString(s string) error {
	for j := 0; j < len(_Matches_index)-1; j++ {
		if s == _Matches_name[_Matches_index[j]:_Matches_index[j+1]] {
			*i = Matches(j)
			return nil
		}
	}
	return errors.New("String: " + s + " is not a valid option for type: Matches")
}
