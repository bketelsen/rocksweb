package models

import (
	"fmt"

	"github.com/a8m/documentdb-go"
	"github.com/markbates/pop/slices"
)

// pkg document
type Package struct {
	documentdb.Document
	AuthorIDs   slices.String `json:"author_ids" db:"author_ids"`
	ImportPath  string        `json:"import_path" db:"import_path"`
	SourceURL   string        `json:"source_url" db:"source_url"`
	License     string        `json:"license" db:"license"`
	Keywords    slices.String `json:"keywords" db:"keywords"`
	Description string        `json:"description" db:"description"`
	VersionIDs  slices.String `json:"version_ids" db:"version_ids"`
	Authors     Authors       `json:"authors" db:"-"`
	Versions    Versions      `json:"versions" db:"-"`
}

// DB interface
type PackageDB interface {
	Get(id string) *Package
	GetAll() []*Package
	Add(a *Package) *Package
	Update(a *Package) *Package
	Remove(id string) error
}

// PackagesDB implement DB interface
type PackagesDB struct {
	Database   string
	Collection string
	db         *documentdb.Database
	coll       *documentdb.Collection
	client     *documentdb.DocumentDB
}

// Return new pkgDB
// Test if database and collection exist. if not, create them.
func NewPackageDB(db, coll string, config *Config) (pkgs PackagesDB) {
	pkgs.Database = db
	pkgs.Collection = coll
	pkgs.client = documentdb.New(config.URL, documentdb.Config{config.MasterKey})
	// Find or create `test` db and `pkgs` collection
	if err := pkgs.findOrDatabase(db); err != nil {
		panic(err)
	}
	if err := pkgs.findOrCreateCollection(coll); err != nil {
		panic(err)
	}
	return
}

// Get pkg by given id
func (p *PackagesDB) Get(id string) (pkg Package, err error) {
	var pkgs []Package
	query := fmt.Sprintf("SELECT * FROM ROOT r WHERE r.id='%s'", id)
	if err = p.client.QueryDocuments(p.coll.Self, query, &pkgs); err != nil || len(pkgs) == 0 {
		return pkg, err
	}
	pkg = pkgs[0]
	return pkg, nil
}

// Get all pkgs
func (p *PackagesDB) GetAll() (pkgs []Package, err error) {
	err = p.client.ReadDocuments(p.coll.Self, &pkgs)
	return
}

// Create pkg
func (p *PackagesDB) Add(pkg *Package) (err error) {
	return p.client.CreateDocument(p.coll.Self, pkg)
}

// Update pkg by id
func (p *PackagesDB) Update(id string, pkg *Package) (err error) {
	var pkgs []Package
	if err = p.client.QueryDocuments(p.coll.Self,
		fmt.Sprintf("SELECT * FROM ROOT r WHERE r.id='%s'", id), &pkgs); err != nil || len(pkgs) == 0 {
		return
	}
	npkg := pkgs[0]
	pkg.Id = npkg.Id
	if err = p.client.ReplaceDocument(npkg.Self, &pkg); err != nil {
		return
	}
	return
}

// Remove pkg by id
func (p *PackagesDB) Remove(id string) (err error) {
	var pkgs []Package
	if err = p.client.QueryDocuments(p.coll.Self,
		fmt.Sprintf("SELECT * FROM ROOT r WHERE r.id='%s'", id), &pkgs); err != nil || len(pkgs) == 0 {
		return
	}
	npkg := pkgs[0]
	if err = p.client.DeleteDocument(npkg.Self); err != nil {
		return
	}
	return
}

// Find or create collection by id
func (p *PackagesDB) findOrCreateCollection(name string) (err error) {
	if colls, err := p.client.QueryCollections(p.db.Self, fmt.Sprintf("SELECT * FROM ROOT r WHERE r.id='%s'", name)); err != nil {
		return err
	} else if len(colls) == 0 {
		if coll, err := p.client.CreateCollection(p.db.Self, fmt.Sprintf(`{ "id": "%s" }`, name)); err != nil {
			return err
		} else {

			p.coll = coll
		}

	} else {
		p.coll = &colls[0]
	}
	return
}

// Find or create database by id
func (p *PackagesDB) findOrDatabase(name string) (err error) {
	if dbs, err := p.client.QueryDatabases(fmt.Sprintf("SELECT * FROM ROOT r WHERE r.id='%s'", name)); err != nil {
		return err
	} else if len(dbs) == 0 {
		if db, err := p.client.CreateDatabase(fmt.Sprintf(`{ "id": "%s" }`, name)); err != nil {
			return err
		} else {
			p.db = db
		}
	} else {
		p.db = &dbs[0]
	}
	return
}
