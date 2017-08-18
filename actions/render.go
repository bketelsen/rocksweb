package actions

import (
	"strings"
	"time"

	"github.com/gobuffalo/buffalo/render"
	"github.com/gobuffalo/packr"
)

var r *render.Engine

func init() {
	r = render.New(render.Options{
		// HTML layout to be used for all HTML requests:
		HTMLLayout: "application.html",

		// Box containing all of the templates:
		TemplatesBox: packr.NewBox("../templates"),

		// Add template helpers here:
		Helpers: render.Helpers{
			"keywordFmt": func(s []string) string {
				return strings.Join(s, ", ")
			},
			"fmtTime": func(t time.Time) string {
				return t.Format("January 02, 2006 @ 15:04:05")
			},
		},
	})
}
