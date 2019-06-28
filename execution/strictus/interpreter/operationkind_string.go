// Code generated by "stringer -type=OperationKind"; DO NOT EDIT.

package interpreter

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[OperationKindUnary-0]
	_ = x[OperationKindBinary-1]
	_ = x[OperationKindTernary-2]
}

const _OperationKind_name = "OperationKindUnaryOperationKindBinaryOperationKindTernary"

var _OperationKind_index = [...]uint8{0, 18, 37, 57}

func (i OperationKind) String() string {
	if i < 0 || i >= OperationKind(len(_OperationKind_index)-1) {
		return "OperationKind(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _OperationKind_name[_OperationKind_index[i]:_OperationKind_index[i+1]]
}
