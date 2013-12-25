package controllers

import "github.com/robfig/revel"

func init() {
	revel.OnAppStart(Init)
	revel.InterceptMethod((*ModelController).Begin, revel.BEFORE)
}