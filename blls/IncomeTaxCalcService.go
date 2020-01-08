package blls

import (
	"fmt"
)

type TaxInfo struct {
	Level          float32
	TaxRate        float32
	QuickDeduction float32
}

type TaxCalcService struct {
	taxs []TaxInfo
}

func main() {
	serv := CreateTaxService()
	num := serv.GetTax(3, 20000, 40000, 300, 600, 2000, 3000, 2300)
	fmt.Println(num)
}

func CreateTaxService() *TaxCalcService {
	entity := &TaxCalcService{}

	entity.taxs = make([]TaxInfo, 8)
	entity.taxs = append(entity.taxs, TaxInfo{Level: 36000, TaxRate: 0.03, QuickDeduction: 0})
	entity.taxs = append(entity.taxs, TaxInfo{Level: 144000, TaxRate: 0.10, QuickDeduction: 2520})
	entity.taxs = append(entity.taxs, TaxInfo{Level: 300000, TaxRate: 0.20, QuickDeduction: 16920})
	entity.taxs = append(entity.taxs, TaxInfo{Level: 420000, TaxRate: 0.25, QuickDeduction: 31920})
	entity.taxs = append(entity.taxs, TaxInfo{Level: 660000, TaxRate: 0.30, QuickDeduction: 52920})
	entity.taxs = append(entity.taxs, TaxInfo{Level: 960000, TaxRate: 0.35, QuickDeduction: 85920})
	entity.taxs = append(entity.taxs, TaxInfo{Level: 999999999, TaxRate: 0.45, QuickDeduction: 181920})

	return entity
}

//MonthIndex
//Sallary,
//AccumulatedSallary,
//ShebaoFee,
//AccumulatedShebaoFee,
//DeductionAmount,
//AccumulatedDeductionAmount,
//AccumulatedTaxFee
func (o *TaxCalcService) GetTax(MonthIndex,
	Sallary, AccumulatedSallary,
	ShebaoFee, AccumulatedShebaoFee,
	DeductionAmount, AccumulatedDeductionAmount,
	AccumulatedTaxFee float32) float32 {
	amount := (Sallary + AccumulatedSallary) - 5000*(MonthIndex) - ShebaoFee - AccumulatedShebaoFee - DeductionAmount - AccumulatedDeductionAmount
	refTax := &TaxInfo{}
	fmt.Println(amount)

	for _, v := range o.taxs {
		if v.Level > amount {
			refTax = &v
			break
		}
	}
	fmt.Println("rate", refTax.TaxRate, "QuickDeduction=>", refTax.QuickDeduction, "AccumulatedTaxFee=>", AccumulatedTaxFee)
	result := amount*refTax.TaxRate - refTax.QuickDeduction - AccumulatedTaxFee
	if result < 0 {
		return 0
	}
	return result

}
