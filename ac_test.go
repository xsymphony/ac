package ac

import (
	"fmt"
	"testing"
)

func TestNewAutomaton(t *testing.T) {
	ac := NewAutomaton()
	ac.Add("你好")
	ac.Add("世界")
	ac.Add("蛤")
	ac.Add("😊")
	ac.Build()
	fmt.Println(ac.Find("你好世界我不知道😊"))
}
