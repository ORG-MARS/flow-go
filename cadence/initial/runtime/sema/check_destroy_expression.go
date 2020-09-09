/*
 * Cadence - The resource-oriented smart contract programming language
 *
 * Copyright 2019-2020 Dapper Labs, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package sema

import (
	"example.com/cadence-initial/runtime/ast"
)

func (checker *Checker) VisitDestroyExpression(expression *ast.DestroyExpression) (resultType ast.Repr) {
	resultType = &VoidType{}

	valueType := expression.Expression.Accept(checker).(Type)

	checker.recordResourceInvalidation(
		expression.Expression,
		valueType,
		ResourceInvalidationKindDestroy,
	)

	// The destruction of any resource type (even compound resource types)

	if valueType.IsInvalidType() {
		return
	}

	if !valueType.IsResourceType() {

		checker.report(
			&InvalidDestructionError{
				Range: ast.NewRangeFromPositioned(expression.Expression),
			},
		)

		return
	}

	return
}