package testdata

import (
	"time"

	"github.com/bketelsen/rocksweb/models"
	"github.com/markbates/pop"
	"github.com/markbates/pop/slices"
	"github.com/pkg/errors"
)

func CreateBuffalo(tx *pop.Connection) error {
	a := models.Author{Name: "Mark Bates"}
	err := tx.Create(&a)
	if err != nil {
		return errors.WithStack(err)
	}

	tm, err := time.Parse(time.RubyDate, "Mon Aug 14 10:46:37 -0400 2017")
	if err != nil {
		return errors.WithStack(err)
	}
	v := models.Version{Tag: "v0.9.3", ReleasedOn: tm}

	err = tx.Create(&v)
	if err != nil {
		return errors.WithStack(err)
	}

	p := models.Package{
		AuthorIDs:   slices.String{a.ID.String()},
		VersionIDs:  slices.String{v.ID.String()},
		ImportPath:  "github.com/gobuffalo/buffalo",
		SourceUrl:   "https://github.com/gobuffalo/buffalo",
		License:     buffaloLicense,
		Keywords:    slices.String{"go", "buffalo", "golang", "rails", "gobuffalo", "web", "framework"},
		Description: buffaloDescription,
	}
	err = tx.Create(&p)
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}

const buffaloDescription = `Buffalo is a Go web development eco-system. Designed to make the life of a Go web developer easier.

Buffalo starts by generating a web project for you that already has everything from front-end (JavaScript, SCSS, etc...) to back-end (database, routing, etc...) already hooked up and ready to run. From there it provides easy APIs to build your web application quickly in Go.

Buffalo isn't just a framework, it's a holistic web development environment and project structure that lets developers get straight to the business of, well, building their business.`

const buffaloLicense = `The MIT License (MIT)
Copyright (c) 2016 Mark Bates

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.`
