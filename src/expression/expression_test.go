package expression

import (
	"fmt"
	"testing"
)

func TestExpressionCompiler(t *testing.T) {
	fmt.Println(CompileExpression("1'b0", "adder"))
}
