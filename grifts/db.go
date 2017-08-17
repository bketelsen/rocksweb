package grifts

import (
	"github.com/bketelsen/rocksweb/models"
	"github.com/bketelsen/rocksweb/testdata"
	"github.com/markbates/grift/grift"
	"github.com/markbates/pop"
	"github.com/pkg/errors"
)

var _ = grift.Namespace("db", func() {

	grift.Desc("seed", "Seeds a database")
	grift.Add("seed", func(c *grift.Context) error {
		return models.DB.Transaction(func(tx *pop.Connection) error {
			err := tx.TruncateAll()
			if err != nil {
				return errors.WithStack(err)
			}

			err = testdata.CreateBuffalo(tx)
			if err != nil {
				return errors.WithStack(err)
			}

			err = testdata.CreateVice(tx)
			if err != nil {
				return errors.WithStack(err)
			}

			return nil

		})
	})

})
