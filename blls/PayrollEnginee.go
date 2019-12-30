package blls

import (
	"fmt"
	"payroll/utils"
	formula "github.com/yidane/formula"
)

type PayrollEnginee struct {
}

func (e *PayrollEnginee) Evaluate(FormulaText string, VarList map[string]interface{}) (float64, error) {
	expression := formula.NewExpression(FormulaText)
	for k, v := range VarList {
		err := expression.AddParameter(k, v)
		if err != nil {
			return 0, err
		}
	}

	result, err1 := expression.Evaluate()
	if err1 != nil {
		return 0, err1
	}

	fmt.Println(result.Value)
	return utils.interface2String(result.Value), nil

}
