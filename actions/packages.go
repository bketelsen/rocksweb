package actions

import (
	"github.com/bketelsen/rocksweb/models"
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/x/responder"
	"github.com/markbates/pop"
	"github.com/pkg/errors"
)

func PackagesIndex(c buffalo.Context) error {
	psq := models.PackageSearchQuery{}
	c.Set("psq", psq)
	return c.Render(200, r.HTML("packages/index.html"))
}

func PackagesSearch(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)

	psq := models.PackageSearchQuery{}
	if err := c.Bind(&psq); err != nil {
		return errors.WithStack(err)
	}

	pkgs, err := models.SearchPackages(psq, tx)
	if err != nil {
		return errors.WithStack(err)
	}

	responder.Wants("json", func(c buffalo.Context) error {
		return c.Render(200, r.JSON(&pkgs))
	}).Wants("html", func(c buffalo.Context) error {
		c.Set("psq", psq)
		c.Set("packages", pkgs)
		return c.Render(200, r.HTML("packages/search.html"))
	}).Respond(c)
	return nil
}
