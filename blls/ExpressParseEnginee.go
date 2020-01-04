package blls

import (
	//"strconv"

	"fmt"
	"regexp"
	"strings"

	formula "github.com/levin9/formula"
)

type ExpressParseEnginee struct {
	VarList  map[string]string
	TipsList map[string]string
}

//计算公式
func (o *ExpressParseEnginee) Evaluate(formula string) (float64, string, error) {
	temp_formula := formula
	formula = strings.ToLower(formula)
	r, _ := regexp.Compile(`\[(.*?)]`)
	arrs := r.FindAllString(formula, -1)
	for _, v := range arrs {
		if v != `<nil>` {
			v1 := strings.TrimRight(strings.TrimLeft(v, "["), "]")
			v1Val := o.VarList[v1]
			if v1Val == "" {
				fmt.Println("替换为空", v1, "=>", v1Val)
			}
			formula = strings.ReplaceAll(formula, v, v1Val)
		}
		//fmt.Println(v)
	}
	result, err1 := Calc(formula)
	if err1 != nil {
		fmt.Println(err1)
		fmt.Println("失败：", temp_formula, "=>", formula)
		return -9999, formula, err1
	}

	//fmt.Println(result)
	return result, formula, nil
}

func Calc(formulaText string) (float64, error) {
	expression := formula.NewExpression(formulaText)
	result, err := expression.Evaluate()
	if err != nil {
		//handle err
		fmt.Println(err)
		return -999, err
	}

	v, err := result.Float64()
	if err != nil {
		//handle err
		return -999, err
	}
	return v, nil
}
