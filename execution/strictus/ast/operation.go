package ast

import "bamboo-runtime/execution/strictus/errors"

//go:generate stringer -type=Operation

type Operation int

const (
	OperationOr Operation = iota
	OperationAnd
	OperationEqual
	OperationUnequal
	OperationLess
	OperationGreater
	OperationLessEqual
	OperationGreaterEqual
	OperationPlus
	OperationMinus
	OperationMul
	OperationDiv
	OperationMod
	OperationNegate
)

func (s Operation) Symbol() string {
	switch s {
	case OperationOr:
		return "||"
	case OperationAnd:
		return "&&"
	case OperationEqual:
		return "=="
	case OperationUnequal:
		return "!="
	case OperationLess:
		return "<"
	case OperationGreater:
		return ">"
	case OperationLessEqual:
		return "<="
	case OperationGreaterEqual:
		return ">="
	case OperationPlus:
		return "+"
	case OperationMinus:
		return "-"
	case OperationMul:
		return "*"
	case OperationDiv:
		return "/"
	case OperationMod:
		return "%"
	case OperationNegate:
		return "!"
	default:
		panic(&errors.UnreachableError{})
	}
}
