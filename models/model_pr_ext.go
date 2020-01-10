package models

type PayrollDomain struct {
}

func (o *PayrollDomain) GetMonthAttendance(MonthId, TenantId string) []MonthAttendanceMapDto {

	return nil
}

type MonthAttendanceMapDto struct {
	FullName string
	PersonId string
	AttId    string //考勤Id
}
