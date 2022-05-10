package model

import "context"

type Connection struct {
	Type      string `json:"type"`
	ConnctStr string `json:"connctstr"`
	User      string `json:"user"`
	Password  string `json:"password"`
	PassKey   string `json:"passkey"`
	Path      string `json:"path"`
	Context   context.Context
}
