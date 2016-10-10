package model

import "time"

// BaseEntity Entity 基类
type BaseEntity struct {
	InsRep  string    `json:"insRep"`
	InsDate time.Time `json:"insDate"`
	UpdRep  string    `json:"updReq"`
	UpdDate time.Time `json:"updDate"`
	DelFlg  string    `json:"delFlg"`
	Version int       `json:"version"`
}
