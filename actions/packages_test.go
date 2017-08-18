package actions

import (
	"github.com/bketelsen/rocksweb/models"
	"github.com/bketelsen/rocksweb/testdata"
)

func (a *ActionSuite) Test_PackageIndex() {
	res := a.HTML("/").Get()
	a.Equal(200, res.Code)
}

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

func (a *ActionSuite) Test_PackagesSearch_NoArgs() {
	err := testdata.CreateBuffalo(a.DB)
	a.NoError(err)

	err = testdata.CreateVice(a.DB)

	psq := models.PackageSearchQuery{}
	res := a.JSON("/packages/search").Post(&psq)
	a.Equal(200, res.Code)

	pkgs := models.Packages{}
	res.Bind(&pkgs)
	a.Len(pkgs, 2)
}
