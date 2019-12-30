package blls

import (
	"fmt"
)

type ClaculateHandler struct {
}

//CalculateMonthPayroll
func (h *ClaculateHandler) CalculateMonthPayroll(TenantId, MonthId string) {
	fmt.Println(TenantId, MonthId)
}
