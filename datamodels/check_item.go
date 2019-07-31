package datamodels

import "time"

type CheckItem struct {
	Date          time.Time //日期
	CheckId       string    //业务编号
	Company       string    //公司名称
	CustomType    int       //客户类型
	Address       string    //联系地址
	Contract      string    //联系人
	Phone         string    //联系电话
	Total         float64   //总计金额
	ReturnedMoney float64   //回款
	PayType       int       //回款
	IsPayed       int       //是否付款
	TicketNumber  string    //票号
}
