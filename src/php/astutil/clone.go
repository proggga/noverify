// Code generated by `codegen.go`. DO NOT EDIT.
package astutil

import (
	"fmt"
	"github.com/VKCOM/noverify/src/php/parser/node"
	"github.com/VKCOM/noverify/src/php/parser/node/expr"
	"github.com/VKCOM/noverify/src/php/parser/node/expr/assign"
	"github.com/VKCOM/noverify/src/php/parser/node/expr/binary"
	"github.com/VKCOM/noverify/src/php/parser/node/expr/cast"
	"github.com/VKCOM/noverify/src/php/parser/node/name"
	"github.com/VKCOM/noverify/src/php/parser/node/scalar"
	"github.com/VKCOM/noverify/src/php/parser/node/stmt"
)

func NodeClone(x node.Node) node.Node {
	if x == nil {
		return nil
	}
	switch x := x.(type) {
	case *assign.Assign:
		clone := *x
		if x.Variable != nil {
			clone.Variable = NodeClone(x.Variable).(node.Node)
		}
		if x.Expression != nil {
			clone.Expression = NodeClone(x.Expression).(node.Node)
		}
		return &clone
	case *assign.BitwiseAnd:
		clone := *x
		if x.Variable != nil {
			clone.Variable = NodeClone(x.Variable).(node.Node)
		}
		if x.Expression != nil {
			clone.Expression = NodeClone(x.Expression).(node.Node)
		}
		return &clone
	case *assign.BitwiseOr:
		clone := *x
		if x.Variable != nil {
			clone.Variable = NodeClone(x.Variable).(node.Node)
		}
		if x.Expression != nil {
			clone.Expression = NodeClone(x.Expression).(node.Node)
		}
		return &clone
	case *assign.BitwiseXor:
		clone := *x
		if x.Variable != nil {
			clone.Variable = NodeClone(x.Variable).(node.Node)
		}
		if x.Expression != nil {
			clone.Expression = NodeClone(x.Expression).(node.Node)
		}
		return &clone
	case *assign.Concat:
		clone := *x
		if x.Variable != nil {
			clone.Variable = NodeClone(x.Variable).(node.Node)
		}
		if x.Expression != nil {
			clone.Expression = NodeClone(x.Expression).(node.Node)
		}
		return &clone
	case *assign.Div:
		clone := *x
		if x.Variable != nil {
			clone.Variable = NodeClone(x.Variable).(node.Node)
		}
		if x.Expression != nil {
			clone.Expression = NodeClone(x.Expression).(node.Node)
		}
		return &clone
	case *assign.Minus:
		clone := *x
		if x.Variable != nil {
			clone.Variable = NodeClone(x.Variable).(node.Node)
		}
		if x.Expression != nil {
			clone.Expression = NodeClone(x.Expression).(node.Node)
		}
		return &clone
	case *assign.Mod:
		clone := *x
		if x.Variable != nil {
			clone.Variable = NodeClone(x.Variable).(node.Node)
		}
		if x.Expression != nil {
			clone.Expression = NodeClone(x.Expression).(node.Node)
		}
		return &clone
	case *assign.Mul:
		clone := *x
		if x.Variable != nil {
			clone.Variable = NodeClone(x.Variable).(node.Node)
		}
		if x.Expression != nil {
			clone.Expression = NodeClone(x.Expression).(node.Node)
		}
		return &clone
	case *assign.Plus:
		clone := *x
		if x.Variable != nil {
			clone.Variable = NodeClone(x.Variable).(node.Node)
		}
		if x.Expression != nil {
			clone.Expression = NodeClone(x.Expression).(node.Node)
		}
		return &clone
	case *assign.Pow:
		clone := *x
		if x.Variable != nil {
			clone.Variable = NodeClone(x.Variable).(node.Node)
		}
		if x.Expression != nil {
			clone.Expression = NodeClone(x.Expression).(node.Node)
		}
		return &clone
	case *assign.Reference:
		clone := *x
		if x.Variable != nil {
			clone.Variable = NodeClone(x.Variable).(node.Node)
		}
		if x.Expression != nil {
			clone.Expression = NodeClone(x.Expression).(node.Node)
		}
		return &clone
	case *assign.ShiftLeft:
		clone := *x
		if x.Variable != nil {
			clone.Variable = NodeClone(x.Variable).(node.Node)
		}
		if x.Expression != nil {
			clone.Expression = NodeClone(x.Expression).(node.Node)
		}
		return &clone
	case *assign.ShiftRight:
		clone := *x
		if x.Variable != nil {
			clone.Variable = NodeClone(x.Variable).(node.Node)
		}
		if x.Expression != nil {
			clone.Expression = NodeClone(x.Expression).(node.Node)
		}
		return &clone
	case *binary.BitwiseAnd:
		clone := *x
		if x.Left != nil {
			clone.Left = NodeClone(x.Left).(node.Node)
		}
		if x.Right != nil {
			clone.Right = NodeClone(x.Right).(node.Node)
		}
		return &clone
	case *binary.BitwiseOr:
		clone := *x
		if x.Left != nil {
			clone.Left = NodeClone(x.Left).(node.Node)
		}
		if x.Right != nil {
			clone.Right = NodeClone(x.Right).(node.Node)
		}
		return &clone
	case *binary.BitwiseXor:
		clone := *x
		if x.Left != nil {
			clone.Left = NodeClone(x.Left).(node.Node)
		}
		if x.Right != nil {
			clone.Right = NodeClone(x.Right).(node.Node)
		}
		return &clone
	case *binary.BooleanAnd:
		clone := *x
		if x.Left != nil {
			clone.Left = NodeClone(x.Left).(node.Node)
		}
		if x.Right != nil {
			clone.Right = NodeClone(x.Right).(node.Node)
		}
		return &clone
	case *binary.BooleanOr:
		clone := *x
		if x.Left != nil {
			clone.Left = NodeClone(x.Left).(node.Node)
		}
		if x.Right != nil {
			clone.Right = NodeClone(x.Right).(node.Node)
		}
		return &clone
	case *binary.Coalesce:
		clone := *x
		if x.Left != nil {
			clone.Left = NodeClone(x.Left).(node.Node)
		}
		if x.Right != nil {
			clone.Right = NodeClone(x.Right).(node.Node)
		}
		return &clone
	case *binary.Concat:
		clone := *x
		if x.Left != nil {
			clone.Left = NodeClone(x.Left).(node.Node)
		}
		if x.Right != nil {
			clone.Right = NodeClone(x.Right).(node.Node)
		}
		return &clone
	case *binary.Div:
		clone := *x
		if x.Left != nil {
			clone.Left = NodeClone(x.Left).(node.Node)
		}
		if x.Right != nil {
			clone.Right = NodeClone(x.Right).(node.Node)
		}
		return &clone
	case *binary.Equal:
		clone := *x
		if x.Left != nil {
			clone.Left = NodeClone(x.Left).(node.Node)
		}
		if x.Right != nil {
			clone.Right = NodeClone(x.Right).(node.Node)
		}
		return &clone
	case *binary.Greater:
		clone := *x
		if x.Left != nil {
			clone.Left = NodeClone(x.Left).(node.Node)
		}
		if x.Right != nil {
			clone.Right = NodeClone(x.Right).(node.Node)
		}
		return &clone
	case *binary.GreaterOrEqual:
		clone := *x
		if x.Left != nil {
			clone.Left = NodeClone(x.Left).(node.Node)
		}
		if x.Right != nil {
			clone.Right = NodeClone(x.Right).(node.Node)
		}
		return &clone
	case *binary.Identical:
		clone := *x
		if x.Left != nil {
			clone.Left = NodeClone(x.Left).(node.Node)
		}
		if x.Right != nil {
			clone.Right = NodeClone(x.Right).(node.Node)
		}
		return &clone
	case *binary.LogicalAnd:
		clone := *x
		if x.Left != nil {
			clone.Left = NodeClone(x.Left).(node.Node)
		}
		if x.Right != nil {
			clone.Right = NodeClone(x.Right).(node.Node)
		}
		return &clone
	case *binary.LogicalOr:
		clone := *x
		if x.Left != nil {
			clone.Left = NodeClone(x.Left).(node.Node)
		}
		if x.Right != nil {
			clone.Right = NodeClone(x.Right).(node.Node)
		}
		return &clone
	case *binary.LogicalXor:
		clone := *x
		if x.Left != nil {
			clone.Left = NodeClone(x.Left).(node.Node)
		}
		if x.Right != nil {
			clone.Right = NodeClone(x.Right).(node.Node)
		}
		return &clone
	case *binary.Minus:
		clone := *x
		if x.Left != nil {
			clone.Left = NodeClone(x.Left).(node.Node)
		}
		if x.Right != nil {
			clone.Right = NodeClone(x.Right).(node.Node)
		}
		return &clone
	case *binary.Mod:
		clone := *x
		if x.Left != nil {
			clone.Left = NodeClone(x.Left).(node.Node)
		}
		if x.Right != nil {
			clone.Right = NodeClone(x.Right).(node.Node)
		}
		return &clone
	case *binary.Mul:
		clone := *x
		if x.Left != nil {
			clone.Left = NodeClone(x.Left).(node.Node)
		}
		if x.Right != nil {
			clone.Right = NodeClone(x.Right).(node.Node)
		}
		return &clone
	case *binary.NotEqual:
		clone := *x
		if x.Left != nil {
			clone.Left = NodeClone(x.Left).(node.Node)
		}
		if x.Right != nil {
			clone.Right = NodeClone(x.Right).(node.Node)
		}
		return &clone
	case *binary.NotIdentical:
		clone := *x
		if x.Left != nil {
			clone.Left = NodeClone(x.Left).(node.Node)
		}
		if x.Right != nil {
			clone.Right = NodeClone(x.Right).(node.Node)
		}
		return &clone
	case *binary.Plus:
		clone := *x
		if x.Left != nil {
			clone.Left = NodeClone(x.Left).(node.Node)
		}
		if x.Right != nil {
			clone.Right = NodeClone(x.Right).(node.Node)
		}
		return &clone
	case *binary.Pow:
		clone := *x
		if x.Left != nil {
			clone.Left = NodeClone(x.Left).(node.Node)
		}
		if x.Right != nil {
			clone.Right = NodeClone(x.Right).(node.Node)
		}
		return &clone
	case *binary.ShiftLeft:
		clone := *x
		if x.Left != nil {
			clone.Left = NodeClone(x.Left).(node.Node)
		}
		if x.Right != nil {
			clone.Right = NodeClone(x.Right).(node.Node)
		}
		return &clone
	case *binary.ShiftRight:
		clone := *x
		if x.Left != nil {
			clone.Left = NodeClone(x.Left).(node.Node)
		}
		if x.Right != nil {
			clone.Right = NodeClone(x.Right).(node.Node)
		}
		return &clone
	case *binary.Smaller:
		clone := *x
		if x.Left != nil {
			clone.Left = NodeClone(x.Left).(node.Node)
		}
		if x.Right != nil {
			clone.Right = NodeClone(x.Right).(node.Node)
		}
		return &clone
	case *binary.SmallerOrEqual:
		clone := *x
		if x.Left != nil {
			clone.Left = NodeClone(x.Left).(node.Node)
		}
		if x.Right != nil {
			clone.Right = NodeClone(x.Right).(node.Node)
		}
		return &clone
	case *binary.Spaceship:
		clone := *x
		if x.Left != nil {
			clone.Left = NodeClone(x.Left).(node.Node)
		}
		if x.Right != nil {
			clone.Right = NodeClone(x.Right).(node.Node)
		}
		return &clone
	case *cast.Array:
		clone := *x
		if x.Expr != nil {
			clone.Expr = NodeClone(x.Expr).(node.Node)
		}
		return &clone
	case *cast.Bool:
		clone := *x
		if x.Expr != nil {
			clone.Expr = NodeClone(x.Expr).(node.Node)
		}
		return &clone
	case *cast.Double:
		clone := *x
		if x.Expr != nil {
			clone.Expr = NodeClone(x.Expr).(node.Node)
		}
		return &clone
	case *cast.Int:
		clone := *x
		if x.Expr != nil {
			clone.Expr = NodeClone(x.Expr).(node.Node)
		}
		return &clone
	case *cast.Object:
		clone := *x
		if x.Expr != nil {
			clone.Expr = NodeClone(x.Expr).(node.Node)
		}
		return &clone
	case *cast.String:
		clone := *x
		if x.Expr != nil {
			clone.Expr = NodeClone(x.Expr).(node.Node)
		}
		return &clone
	case *cast.Unset:
		clone := *x
		if x.Expr != nil {
			clone.Expr = NodeClone(x.Expr).(node.Node)
		}
		return &clone
	case *expr.Array:
		clone := *x
		{
			sliceClone := make([]*expr.ArrayItem, len(x.Items))
			for i := range x.Items {
				sliceClone[i] = NodeClone(x.Items[i]).(*expr.ArrayItem)
			}
			clone.Items = sliceClone
		}
		return &clone
	case *expr.ArrayDimFetch:
		clone := *x
		if x.Variable != nil {
			clone.Variable = NodeClone(x.Variable).(node.Node)
		}
		if x.Dim != nil {
			clone.Dim = NodeClone(x.Dim).(node.Node)
		}
		return &clone
	case *expr.ArrayItem:
		clone := *x
		if x.Key != nil {
			clone.Key = NodeClone(x.Key).(node.Node)
		}
		if x.Val != nil {
			clone.Val = NodeClone(x.Val).(node.Node)
		}
		return &clone
	case *expr.BitwiseNot:
		clone := *x
		if x.Expr != nil {
			clone.Expr = NodeClone(x.Expr).(node.Node)
		}
		return &clone
	case *expr.BooleanNot:
		clone := *x
		if x.Expr != nil {
			clone.Expr = NodeClone(x.Expr).(node.Node)
		}
		return &clone
	case *expr.ClassConstFetch:
		clone := *x
		if x.Class != nil {
			clone.Class = NodeClone(x.Class).(node.Node)
		}
		if x.ConstantName != nil {
			clone.ConstantName = NodeClone(x.ConstantName).(*node.Identifier)
		}
		return &clone
	case *expr.Clone:
		clone := *x
		if x.Expr != nil {
			clone.Expr = NodeClone(x.Expr).(node.Node)
		}
		return &clone
	case *expr.Closure:
		clone := *x
		clone.Params = NodeSliceClone(x.Params)
		if x.ClosureUse != nil {
			clone.ClosureUse = NodeClone(x.ClosureUse).(*expr.ClosureUse)
		}
		if x.ReturnType != nil {
			clone.ReturnType = NodeClone(x.ReturnType).(node.Node)
		}
		clone.Stmts = NodeSliceClone(x.Stmts)
		return &clone
	case *expr.ClosureUse:
		clone := *x
		clone.Uses = NodeSliceClone(x.Uses)
		return &clone
	case *expr.ConstFetch:
		clone := *x
		if x.Constant != nil {
			clone.Constant = NodeClone(x.Constant).(node.Node)
		}
		return &clone
	case *expr.Empty:
		clone := *x
		if x.Expr != nil {
			clone.Expr = NodeClone(x.Expr).(node.Node)
		}
		return &clone
	case *expr.ErrorSuppress:
		clone := *x
		if x.Expr != nil {
			clone.Expr = NodeClone(x.Expr).(node.Node)
		}
		return &clone
	case *expr.Eval:
		clone := *x
		if x.Expr != nil {
			clone.Expr = NodeClone(x.Expr).(node.Node)
		}
		return &clone
	case *expr.Exit:
		clone := *x
		if x.Expr != nil {
			clone.Expr = NodeClone(x.Expr).(node.Node)
		}
		return &clone
	case *expr.FunctionCall:
		clone := *x
		if x.Function != nil {
			clone.Function = NodeClone(x.Function).(node.Node)
		}
		if x.ArgumentList != nil {
			clone.ArgumentList = NodeClone(x.ArgumentList).(*node.ArgumentList)
		}
		return &clone
	case *expr.Include:
		clone := *x
		if x.Expr != nil {
			clone.Expr = NodeClone(x.Expr).(node.Node)
		}
		return &clone
	case *expr.IncludeOnce:
		clone := *x
		if x.Expr != nil {
			clone.Expr = NodeClone(x.Expr).(node.Node)
		}
		return &clone
	case *expr.InstanceOf:
		clone := *x
		if x.Expr != nil {
			clone.Expr = NodeClone(x.Expr).(node.Node)
		}
		if x.Class != nil {
			clone.Class = NodeClone(x.Class).(node.Node)
		}
		return &clone
	case *expr.Isset:
		clone := *x
		clone.Variables = NodeSliceClone(x.Variables)
		return &clone
	case *expr.List:
		clone := *x
		{
			sliceClone := make([]*expr.ArrayItem, len(x.Items))
			for i := range x.Items {
				sliceClone[i] = NodeClone(x.Items[i]).(*expr.ArrayItem)
			}
			clone.Items = sliceClone
		}
		return &clone
	case *expr.MethodCall:
		clone := *x
		if x.Variable != nil {
			clone.Variable = NodeClone(x.Variable).(node.Node)
		}
		if x.Method != nil {
			clone.Method = NodeClone(x.Method).(node.Node)
		}
		if x.ArgumentList != nil {
			clone.ArgumentList = NodeClone(x.ArgumentList).(*node.ArgumentList)
		}
		return &clone
	case *expr.New:
		clone := *x
		if x.Class != nil {
			clone.Class = NodeClone(x.Class).(node.Node)
		}
		if x.ArgumentList != nil {
			clone.ArgumentList = NodeClone(x.ArgumentList).(*node.ArgumentList)
		}
		return &clone
	case *expr.Paren:
		clone := *x
		if x.Expr != nil {
			clone.Expr = NodeClone(x.Expr).(node.Node)
		}
		return &clone
	case *expr.PostDec:
		clone := *x
		if x.Variable != nil {
			clone.Variable = NodeClone(x.Variable).(node.Node)
		}
		return &clone
	case *expr.PostInc:
		clone := *x
		if x.Variable != nil {
			clone.Variable = NodeClone(x.Variable).(node.Node)
		}
		return &clone
	case *expr.PreDec:
		clone := *x
		if x.Variable != nil {
			clone.Variable = NodeClone(x.Variable).(node.Node)
		}
		return &clone
	case *expr.PreInc:
		clone := *x
		if x.Variable != nil {
			clone.Variable = NodeClone(x.Variable).(node.Node)
		}
		return &clone
	case *expr.Print:
		clone := *x
		if x.Expr != nil {
			clone.Expr = NodeClone(x.Expr).(node.Node)
		}
		return &clone
	case *expr.PropertyFetch:
		clone := *x
		if x.Variable != nil {
			clone.Variable = NodeClone(x.Variable).(node.Node)
		}
		if x.Property != nil {
			clone.Property = NodeClone(x.Property).(node.Node)
		}
		return &clone
	case *expr.Reference:
		clone := *x
		if x.Variable != nil {
			clone.Variable = NodeClone(x.Variable).(node.Node)
		}
		return &clone
	case *expr.Require:
		clone := *x
		if x.Expr != nil {
			clone.Expr = NodeClone(x.Expr).(node.Node)
		}
		return &clone
	case *expr.RequireOnce:
		clone := *x
		if x.Expr != nil {
			clone.Expr = NodeClone(x.Expr).(node.Node)
		}
		return &clone
	case *expr.ShellExec:
		clone := *x
		clone.Parts = NodeSliceClone(x.Parts)
		return &clone
	case *expr.StaticCall:
		clone := *x
		if x.Class != nil {
			clone.Class = NodeClone(x.Class).(node.Node)
		}
		if x.Call != nil {
			clone.Call = NodeClone(x.Call).(node.Node)
		}
		if x.ArgumentList != nil {
			clone.ArgumentList = NodeClone(x.ArgumentList).(*node.ArgumentList)
		}
		return &clone
	case *expr.StaticPropertyFetch:
		clone := *x
		if x.Class != nil {
			clone.Class = NodeClone(x.Class).(node.Node)
		}
		if x.Property != nil {
			clone.Property = NodeClone(x.Property).(node.Node)
		}
		return &clone
	case *expr.Ternary:
		clone := *x
		if x.Condition != nil {
			clone.Condition = NodeClone(x.Condition).(node.Node)
		}
		if x.IfTrue != nil {
			clone.IfTrue = NodeClone(x.IfTrue).(node.Node)
		}
		if x.IfFalse != nil {
			clone.IfFalse = NodeClone(x.IfFalse).(node.Node)
		}
		return &clone
	case *expr.UnaryMinus:
		clone := *x
		if x.Expr != nil {
			clone.Expr = NodeClone(x.Expr).(node.Node)
		}
		return &clone
	case *expr.UnaryPlus:
		clone := *x
		if x.Expr != nil {
			clone.Expr = NodeClone(x.Expr).(node.Node)
		}
		return &clone
	case *expr.Yield:
		clone := *x
		if x.Key != nil {
			clone.Key = NodeClone(x.Key).(node.Node)
		}
		if x.Value != nil {
			clone.Value = NodeClone(x.Value).(node.Node)
		}
		return &clone
	case *expr.YieldFrom:
		clone := *x
		if x.Expr != nil {
			clone.Expr = NodeClone(x.Expr).(node.Node)
		}
		return &clone
	case *name.FullyQualified:
		clone := *x
		clone.Parts = NodeSliceClone(x.Parts)
		return &clone
	case *name.Name:
		clone := *x
		clone.Parts = NodeSliceClone(x.Parts)
		return &clone
	case *name.NamePart:
		clone := *x
		return &clone
	case *name.Relative:
		clone := *x
		clone.Parts = NodeSliceClone(x.Parts)
		return &clone
	case *node.Argument:
		clone := *x
		if x.Expr != nil {
			clone.Expr = NodeClone(x.Expr).(node.Node)
		}
		return &clone
	case *node.ArgumentList:
		clone := *x
		{
			sliceClone := make([]node.Node, len(x.Arguments))
			for i := range x.Arguments {
				sliceClone[i] = NodeClone(x.Arguments[i]).(node.Node)
			}
			clone.Arguments = sliceClone
		}
		return &clone
	case *node.Identifier:
		clone := *x
		return &clone
	case *node.Nullable:
		clone := *x
		if x.Expr != nil {
			clone.Expr = NodeClone(x.Expr).(node.Node)
		}
		return &clone
	case *node.Parameter:
		clone := *x
		if x.VariableType != nil {
			clone.VariableType = NodeClone(x.VariableType).(node.Node)
		}
		if x.Variable != nil {
			clone.Variable = NodeClone(x.Variable).(*node.SimpleVar)
		}
		if x.DefaultValue != nil {
			clone.DefaultValue = NodeClone(x.DefaultValue).(node.Node)
		}
		return &clone
	case *node.Root:
		clone := *x
		{
			sliceClone := make([]node.Node, len(x.Stmts))
			for i := range x.Stmts {
				sliceClone[i] = NodeClone(x.Stmts[i]).(node.Node)
			}
			clone.Stmts = sliceClone
		}
		return &clone
	case *node.SimpleVar:
		clone := *x
		return &clone
	case *node.Var:
		clone := *x
		if x.Expr != nil {
			clone.Expr = NodeClone(x.Expr).(node.Node)
		}
		return &clone
	case *scalar.Dnumber:
		clone := *x
		return &clone
	case *scalar.Encapsed:
		clone := *x
		clone.Parts = NodeSliceClone(x.Parts)
		return &clone
	case *scalar.EncapsedStringPart:
		clone := *x
		return &clone
	case *scalar.Heredoc:
		clone := *x
		clone.Parts = NodeSliceClone(x.Parts)
		return &clone
	case *scalar.Lnumber:
		clone := *x
		return &clone
	case *scalar.MagicConstant:
		clone := *x
		return &clone
	case *scalar.String:
		clone := *x
		return &clone
	case *stmt.Break:
		clone := *x
		if x.Expr != nil {
			clone.Expr = NodeClone(x.Expr).(node.Node)
		}
		return &clone
	case *stmt.Case:
		clone := *x
		if x.Cond != nil {
			clone.Cond = NodeClone(x.Cond).(node.Node)
		}
		clone.Stmts = NodeSliceClone(x.Stmts)
		return &clone
	case *stmt.CaseList:
		clone := *x
		clone.Cases = NodeSliceClone(x.Cases)
		return &clone
	case *stmt.Catch:
		clone := *x
		clone.Types = NodeSliceClone(x.Types)
		if x.Variable != nil {
			clone.Variable = NodeClone(x.Variable).(*node.SimpleVar)
		}
		clone.Stmts = NodeSliceClone(x.Stmts)
		return &clone
	case *stmt.Class:
		clone := *x
		if x.ClassName != nil {
			clone.ClassName = NodeClone(x.ClassName).(*node.Identifier)
		}
		{
			sliceClone := make([]*node.Identifier, len(x.Modifiers))
			for i := range x.Modifiers {
				sliceClone[i] = NodeClone(x.Modifiers[i]).(*node.Identifier)
			}
			clone.Modifiers = sliceClone
		}
		if x.ArgumentList != nil {
			clone.ArgumentList = NodeClone(x.ArgumentList).(*node.ArgumentList)
		}
		if x.Extends != nil {
			clone.Extends = NodeClone(x.Extends).(*stmt.ClassExtends)
		}
		if x.Implements != nil {
			clone.Implements = NodeClone(x.Implements).(*stmt.ClassImplements)
		}
		clone.Stmts = NodeSliceClone(x.Stmts)
		return &clone
	case *stmt.ClassConstList:
		clone := *x
		{
			sliceClone := make([]*node.Identifier, len(x.Modifiers))
			for i := range x.Modifiers {
				sliceClone[i] = NodeClone(x.Modifiers[i]).(*node.Identifier)
			}
			clone.Modifiers = sliceClone
		}
		clone.Consts = NodeSliceClone(x.Consts)
		return &clone
	case *stmt.ClassExtends:
		clone := *x
		if x.ClassName != nil {
			clone.ClassName = NodeClone(x.ClassName).(node.Node)
		}
		return &clone
	case *stmt.ClassImplements:
		clone := *x
		clone.InterfaceNames = NodeSliceClone(x.InterfaceNames)
		return &clone
	case *stmt.ClassMethod:
		clone := *x
		if x.MethodName != nil {
			clone.MethodName = NodeClone(x.MethodName).(*node.Identifier)
		}
		{
			sliceClone := make([]*node.Identifier, len(x.Modifiers))
			for i := range x.Modifiers {
				sliceClone[i] = NodeClone(x.Modifiers[i]).(*node.Identifier)
			}
			clone.Modifiers = sliceClone
		}
		clone.Params = NodeSliceClone(x.Params)
		if x.ReturnType != nil {
			clone.ReturnType = NodeClone(x.ReturnType).(node.Node)
		}
		if x.Stmt != nil {
			clone.Stmt = NodeClone(x.Stmt).(node.Node)
		}
		return &clone
	case *stmt.ConstList:
		clone := *x
		clone.Consts = NodeSliceClone(x.Consts)
		return &clone
	case *stmt.Constant:
		clone := *x
		if x.ConstantName != nil {
			clone.ConstantName = NodeClone(x.ConstantName).(*node.Identifier)
		}
		if x.Expr != nil {
			clone.Expr = NodeClone(x.Expr).(node.Node)
		}
		return &clone
	case *stmt.Continue:
		clone := *x
		if x.Expr != nil {
			clone.Expr = NodeClone(x.Expr).(node.Node)
		}
		return &clone
	case *stmt.Declare:
		clone := *x
		clone.Consts = NodeSliceClone(x.Consts)
		if x.Stmt != nil {
			clone.Stmt = NodeClone(x.Stmt).(node.Node)
		}
		return &clone
	case *stmt.Default:
		clone := *x
		clone.Stmts = NodeSliceClone(x.Stmts)
		return &clone
	case *stmt.Do:
		clone := *x
		if x.Stmt != nil {
			clone.Stmt = NodeClone(x.Stmt).(node.Node)
		}
		if x.Cond != nil {
			clone.Cond = NodeClone(x.Cond).(node.Node)
		}
		return &clone
	case *stmt.Echo:
		clone := *x
		clone.Exprs = NodeSliceClone(x.Exprs)
		return &clone
	case *stmt.Else:
		clone := *x
		if x.Stmt != nil {
			clone.Stmt = NodeClone(x.Stmt).(node.Node)
		}
		return &clone
	case *stmt.ElseIf:
		clone := *x
		if x.Cond != nil {
			clone.Cond = NodeClone(x.Cond).(node.Node)
		}
		if x.Stmt != nil {
			clone.Stmt = NodeClone(x.Stmt).(node.Node)
		}
		return &clone
	case *stmt.Expression:
		clone := *x
		if x.Expr != nil {
			clone.Expr = NodeClone(x.Expr).(node.Node)
		}
		return &clone
	case *stmt.Finally:
		clone := *x
		clone.Stmts = NodeSliceClone(x.Stmts)
		return &clone
	case *stmt.For:
		clone := *x
		clone.Init = NodeSliceClone(x.Init)
		clone.Cond = NodeSliceClone(x.Cond)
		clone.Loop = NodeSliceClone(x.Loop)
		if x.Stmt != nil {
			clone.Stmt = NodeClone(x.Stmt).(node.Node)
		}
		return &clone
	case *stmt.Foreach:
		clone := *x
		if x.Expr != nil {
			clone.Expr = NodeClone(x.Expr).(node.Node)
		}
		if x.Key != nil {
			clone.Key = NodeClone(x.Key).(node.Node)
		}
		if x.Variable != nil {
			clone.Variable = NodeClone(x.Variable).(node.Node)
		}
		if x.Stmt != nil {
			clone.Stmt = NodeClone(x.Stmt).(node.Node)
		}
		return &clone
	case *stmt.Function:
		clone := *x
		if x.FunctionName != nil {
			clone.FunctionName = NodeClone(x.FunctionName).(*node.Identifier)
		}
		clone.Params = NodeSliceClone(x.Params)
		if x.ReturnType != nil {
			clone.ReturnType = NodeClone(x.ReturnType).(node.Node)
		}
		clone.Stmts = NodeSliceClone(x.Stmts)
		return &clone
	case *stmt.Global:
		clone := *x
		clone.Vars = NodeSliceClone(x.Vars)
		return &clone
	case *stmt.Goto:
		clone := *x
		if x.Label != nil {
			clone.Label = NodeClone(x.Label).(*node.Identifier)
		}
		return &clone
	case *stmt.GroupUse:
		clone := *x
		if x.UseType != nil {
			clone.UseType = NodeClone(x.UseType).(node.Node)
		}
		if x.Prefix != nil {
			clone.Prefix = NodeClone(x.Prefix).(node.Node)
		}
		clone.UseList = NodeSliceClone(x.UseList)
		return &clone
	case *stmt.HaltCompiler:
		clone := *x
		return &clone
	case *stmt.If:
		clone := *x
		if x.Cond != nil {
			clone.Cond = NodeClone(x.Cond).(node.Node)
		}
		if x.Stmt != nil {
			clone.Stmt = NodeClone(x.Stmt).(node.Node)
		}
		clone.ElseIf = NodeSliceClone(x.ElseIf)
		if x.Else != nil {
			clone.Else = NodeClone(x.Else).(node.Node)
		}
		return &clone
	case *stmt.InlineHtml:
		clone := *x
		return &clone
	case *stmt.Interface:
		clone := *x
		if x.InterfaceName != nil {
			clone.InterfaceName = NodeClone(x.InterfaceName).(*node.Identifier)
		}
		if x.Extends != nil {
			clone.Extends = NodeClone(x.Extends).(*stmt.InterfaceExtends)
		}
		clone.Stmts = NodeSliceClone(x.Stmts)
		return &clone
	case *stmt.InterfaceExtends:
		clone := *x
		clone.InterfaceNames = NodeSliceClone(x.InterfaceNames)
		return &clone
	case *stmt.Label:
		clone := *x
		if x.LabelName != nil {
			clone.LabelName = NodeClone(x.LabelName).(*node.Identifier)
		}
		return &clone
	case *stmt.Namespace:
		clone := *x
		if x.NamespaceName != nil {
			clone.NamespaceName = NodeClone(x.NamespaceName).(node.Node)
		}
		clone.Stmts = NodeSliceClone(x.Stmts)
		return &clone
	case *stmt.Nop:
		clone := *x
		return &clone
	case *stmt.Property:
		clone := *x
		if x.Variable != nil {
			clone.Variable = NodeClone(x.Variable).(*node.SimpleVar)
		}
		if x.Expr != nil {
			clone.Expr = NodeClone(x.Expr).(node.Node)
		}
		return &clone
	case *stmt.PropertyList:
		clone := *x
		{
			sliceClone := make([]*node.Identifier, len(x.Modifiers))
			for i := range x.Modifiers {
				sliceClone[i] = NodeClone(x.Modifiers[i]).(*node.Identifier)
			}
			clone.Modifiers = sliceClone
		}
		clone.Properties = NodeSliceClone(x.Properties)
		return &clone
	case *stmt.Return:
		clone := *x
		if x.Expr != nil {
			clone.Expr = NodeClone(x.Expr).(node.Node)
		}
		return &clone
	case *stmt.Static:
		clone := *x
		clone.Vars = NodeSliceClone(x.Vars)
		return &clone
	case *stmt.StaticVar:
		clone := *x
		if x.Variable != nil {
			clone.Variable = NodeClone(x.Variable).(*node.SimpleVar)
		}
		if x.Expr != nil {
			clone.Expr = NodeClone(x.Expr).(node.Node)
		}
		return &clone
	case *stmt.StmtList:
		clone := *x
		clone.Stmts = NodeSliceClone(x.Stmts)
		return &clone
	case *stmt.Switch:
		clone := *x
		if x.Cond != nil {
			clone.Cond = NodeClone(x.Cond).(node.Node)
		}
		if x.CaseList != nil {
			clone.CaseList = NodeClone(x.CaseList).(*stmt.CaseList)
		}
		return &clone
	case *stmt.Throw:
		clone := *x
		if x.Expr != nil {
			clone.Expr = NodeClone(x.Expr).(node.Node)
		}
		return &clone
	case *stmt.Trait:
		clone := *x
		if x.TraitName != nil {
			clone.TraitName = NodeClone(x.TraitName).(*node.Identifier)
		}
		clone.Stmts = NodeSliceClone(x.Stmts)
		return &clone
	case *stmt.TraitAdaptationList:
		clone := *x
		clone.Adaptations = NodeSliceClone(x.Adaptations)
		return &clone
	case *stmt.TraitMethodRef:
		clone := *x
		if x.Trait != nil {
			clone.Trait = NodeClone(x.Trait).(node.Node)
		}
		if x.Method != nil {
			clone.Method = NodeClone(x.Method).(*node.Identifier)
		}
		return &clone
	case *stmt.TraitUse:
		clone := *x
		clone.Traits = NodeSliceClone(x.Traits)
		if x.TraitAdaptationList != nil {
			clone.TraitAdaptationList = NodeClone(x.TraitAdaptationList).(node.Node)
		}
		return &clone
	case *stmt.TraitUseAlias:
		clone := *x
		if x.Ref != nil {
			clone.Ref = NodeClone(x.Ref).(node.Node)
		}
		if x.Modifier != nil {
			clone.Modifier = NodeClone(x.Modifier).(node.Node)
		}
		if x.Alias != nil {
			clone.Alias = NodeClone(x.Alias).(*node.Identifier)
		}
		return &clone
	case *stmt.TraitUsePrecedence:
		clone := *x
		if x.Ref != nil {
			clone.Ref = NodeClone(x.Ref).(node.Node)
		}
		clone.Insteadof = NodeSliceClone(x.Insteadof)
		return &clone
	case *stmt.Try:
		clone := *x
		clone.Stmts = NodeSliceClone(x.Stmts)
		clone.Catches = NodeSliceClone(x.Catches)
		if x.Finally != nil {
			clone.Finally = NodeClone(x.Finally).(node.Node)
		}
		return &clone
	case *stmt.Unset:
		clone := *x
		clone.Vars = NodeSliceClone(x.Vars)
		return &clone
	case *stmt.Use:
		clone := *x
		if x.UseType != nil {
			clone.UseType = NodeClone(x.UseType).(*node.Identifier)
		}
		if x.Use != nil {
			clone.Use = NodeClone(x.Use).(node.Node)
		}
		if x.Alias != nil {
			clone.Alias = NodeClone(x.Alias).(*node.Identifier)
		}
		return &clone
	case *stmt.UseList:
		clone := *x
		if x.UseType != nil {
			clone.UseType = NodeClone(x.UseType).(node.Node)
		}
		clone.Uses = NodeSliceClone(x.Uses)
		return &clone
	case *stmt.While:
		clone := *x
		if x.Cond != nil {
			clone.Cond = NodeClone(x.Cond).(node.Node)
		}
		if x.Stmt != nil {
			clone.Stmt = NodeClone(x.Stmt).(node.Node)
		}
		return &clone
	}
	panic(fmt.Sprintf(`unhandled type %T`, x))
}