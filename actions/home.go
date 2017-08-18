package actions

import (
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/buffalo/worker"
)

// HomeHandler is a default handler to serve up
// a home page.
func HomeHandler(c buffalo.Context) error {
	doWork()
	return c.Render(200, r.HTML("index.html"))
}

func doWork() {
	App().Worker.Perform(worker.Job{
		Queue:   "default",
		Handler: "index_package",
		Args: worker.Args{
			"repository": "https://github.com/bketelsen/captainhook",
		},
	})
}
