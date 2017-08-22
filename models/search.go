package models

type PackageSearchQuery struct {
	Authors  []string `json:"authors"`
	Keywords []string `json:"keywords"`
}

/*
func SearchPackages(psq PackageSearchQuery, tx *pop.Connection) (Packages, error) {
	packages := Packages{}
	q := tx.Q()
	if len(psq.Authors) > 0 {
		authors := Authors{}
		aq := tx.Q()
		for _, n := range psq.Authors {
			n = strings.TrimSpace(n)
			if n != "" {
				aq = tx.Where("name ilike ?", fmt.Sprintf("%%%s%%", strings.ToLower(n)))
			}
		}
		err := aq.All(&authors)
		if err != nil {
			return packages, errors.WithStack(err)
		}
		if len(authors) > 0 {
			ids := []string{}
			for _, a := range authors {
				ids = append(ids, a.ID.String())
			}
			q = q.Where("(packages.author_ids)::uuid[] && ?::uuid[]", fmt.Sprintf("{%s}", strings.Join(ids, ",")))
		}
	}

	if len(psq.Keywords) > 0 {
		q = q.Where("keywords @> ?::varchar[]", fmt.Sprintf("{%s}", strings.Join(psq.Keywords, ",")))
	}

	err := q.All(&packages)
	if err != nil {
		return packages, errors.WithStack(err)
	}

	for i, p := range packages {
		authors := Authors{}
		v, err := p.AuthorIDs.Value()
		if err != nil {
			return packages, errors.WithStack(err)
		}
		err = tx.Where("authors.id = any(?)", v).All(&authors)
		if err != nil {
			return packages, errors.WithStack(err)
		}
		p.Authors = authors

		versions := Versions{}
		v, err = p.VersionIDs.Value()
		if err != nil {
			return packages, errors.WithStack(err)
		}
		err = tx.Where("versions.id = any(?)", v).All(&versions)
		if err != nil {
			return packages, errors.WithStack(err)
		}
		p.Versions = versions
		packages[i] = p
	}

	return packages, nil
}
*/
