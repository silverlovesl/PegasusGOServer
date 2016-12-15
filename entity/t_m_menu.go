package entity

import "time"

type T_M_Menu struct {
	MenuID       string
	MenuCategory string
	TypeFor      string
	MenuName     string
	IconName     *string
	SortBy       int
	InsRep       string
	InsDate      time.Time
	UpdRep       string
	UpdDate      time.Time
	DelFlg       string
	Version      int
}
