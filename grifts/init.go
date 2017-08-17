package grifts

import (
	"github.com/bketelsen/rocksweb/actions"
	"github.com/gobuffalo/buffalo"
)

func init() {
	buffalo.Grifts(actions.App())
}
