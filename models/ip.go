package models

import (
	"time"

	"gorm.io/gorm"
)

type IPModel struct {
	gorm.Model
	IP string `json:"ip"`
}

type ReqIpModel struct {
	IP string `json:"ip"`
}

type ResIpModel struct {
	IP        string    `json:"ip"`
	CreateAt  time.Time `json:"create_at"`
	UpdatedAt time.Time `json:"update_at"`
}

func ReqFormatterIP(req ReqIpModel) IPModel {
	formatter := IPModel{
		IP: req.IP,
	}
	return formatter
}

func ResFormatterIP(res IPModel) ResIpModel {
	formatter := ResIpModel{
		IP:        res.IP,
		CreateAt:  res.CreatedAt,
		UpdatedAt: res.UpdatedAt,
	}

	return formatter
}

func GetAllFormatterIP(results []*IPModel) []*ResIpModel {
	var response []*ResIpModel
	for _, result := range results {
		data := ResFormatterIP(*result)
		response = append(response, &data)
	}
	return response
}