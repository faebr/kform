// Code generated by "stringer -type=StatusPolicy -linecomment"; DO NOT EDIT.

package policy

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[StatusPolicyNone-0]
	_ = x[StatusPolicyAll-1]
}

const _StatusPolicy_name = "NoneAll"

var _StatusPolicy_index = [...]uint8{0, 4, 7}

func (i StatusPolicy) String() string {
	if i < 0 || i >= StatusPolicy(len(_StatusPolicy_index)-1) {
		return "StatusPolicy(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _StatusPolicy_name[_StatusPolicy_index[i]:_StatusPolicy_index[i+1]]
}