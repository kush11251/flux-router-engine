package models

import (
	"encoding/json"
)

type Router struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	IP       string `json:"ip"`
}