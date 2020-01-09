module payroll

go 1.13

require (
	github.com/360EntSecGroup-Skylar/excelize v1.4.1
	github.com/antlr/antlr4 v0.0.0-20191217191749-ff67971f8580 // indirect
	github.com/gin-contrib/cors v1.3.0
	github.com/gin-contrib/static v0.0.0-20191128031702-f81c604d8ac2
	github.com/gin-gonic/autotls v0.0.0-20191129055149-ffaac874b99f
	github.com/gin-gonic/gin v1.5.0
	github.com/go-redis/redis v6.15.6+incompatible
	github.com/go-sql-driver/mysql v1.4.1
	github.com/jinzhu/gorm v1.9.11
	github.com/levin9/payroll/blls v0.0.0
	github.com/levin9/payroll/utils v0.0.0-20200108162034-9538e8d09262 // indirect
	github.com/sirupsen/logrus v1.4.2
	github.com/spf13/viper v1.6.1
	github.com/yidane/formula v0.0.0-20190507083929-d272721f4dd9 // indirect
)

replace github.com/levin9/payroll/blls => ../payroll/blls
