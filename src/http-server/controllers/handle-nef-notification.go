package controllers

import "github.com/opoze/5gaf/http-server/commands"

// Handle External NEF notify here
func HandleNefNotification(data *commands.NotifyData) {
	commands.AF(data)
}
