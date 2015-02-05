package handlers

import (
	"github.com/ehudu/toropress/libs"
	"github.com/ehudu/toropress/models"
)

type DeleteReplyHandler struct {
	libs.RootAuthHandler
}

func (self *DeleteReplyHandler) Get() {
	rid, _ := self.GetInt64(":rid")
	models.DelReply(rid)
	self.Ctx.Redirect(302, "/")
}
