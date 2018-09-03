package models

import (
	"github.com/mnhkahn/maodou/cygo"
)

type Result struct {
	Id          string      `xorm:"not null pk autoincr INT(11)" json:"id"`
	Title       string      `xorm:"not null VARCHAR(200)" json:"title"`
	CreateTime  cygo.CyTime `xorm:"not null DATETIME" json:"create_time"`
	Author      string      `xorm:"VARCHAR(45)" json:"author"`
	Detail      string      `xorm:"not null LONGTEXT" json:"detail"`
	Category    string      `xorm:"VARCHAR(45)" json:"category"`
	Tags        string      `xorm:"VARCHAR(45)" json:"tags"`
	Figure      string      `xorm:"VARCHAR(100)" json:"figure"`
	Description string      `xorm:"TINYTEXT" json:"description"`
	Link        string      `xorm:"not null VARCHAR(100)" json:"link"`
	Source      string      `xorm:"VARCHAR(40)" json:"source"`
	ParseDate   cygo.CyTime `xorm:"updated" json:"parse_date"`
}
