package global

import (
	"gorm.io/gorm"
)

var (
	GLOBAL_DB *gorm.DB
	//GLOBAL_Timer time.Timer = time.NewTimer()
	MonthMap = map[string]int{"January": 1, "February": 2, "March": 3, "April": 4, "May": 5, "June": 6, "July": 7, "August": 8, "September": 9, "October": 10, "November": 11, "December": 12}
	WeekMap  = map[string]string{"1": "星期一", "2": "星期二", "3": "星期三", "4": "星期四", "5": "星期五", "6": "星期六", "7": "星期天"}
)
