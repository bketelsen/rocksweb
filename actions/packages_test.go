package actions

import (
	"github.com/bketelsen/rocksweb/models"
	"github.com/bketelsen/rocksweb/testdata"
)

func (a *ActionSuite) Test_PackagesSearch() {
	err := testdata.CreateBuffalo(a.DB)
	a.NoError(err)

	psq := models.PackageSearchQuery{Authors: []string{"Mark bates"}}
	res := a.JSON("/packages/search").Post(&psq)
	a.Equal(200, res.Code)

	pkgs := models.Packages{}
	res.Bind(&pkgs)
	a.Len(pkgs, 1)
}
