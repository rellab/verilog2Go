package expression

import (
	"fmt"
	"testing"
)

func TestExpressionCompiler(t *testing.T) {
	fmt.Println(compileExpression("a+b*c+ d[2] + e"))
}
