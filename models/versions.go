package models

import (
	"fmt"
	"time"

	"github.com/a8m/documentdb-go"
)

// Version document
type Version struct {
	documentdb.Document
	Tag        string    `json:"tag" db:"tag"`
	ReleasedOn time.Time `json:"released_on" db:"released_on"`
}

// DB interface
type VersionDB interface {
	Get(id string) *Version
	GetAll() []*Version
	Add(a *Version) *Version
	Update(a *Version) *Version
	Remove(id string) error
}

// VersionsDB implement DB interface
type VersionsDB struct {
	Database   string
	Collection string
	db         *documentdb.Database
	coll       *documentdb.Collection
	client     *documentdb.DocumentDB
}

// Return new VersionDB
// Test if database and collection exist. if not, create them.
func NewVersionDB(db, coll string, config *Config) (versions VersionsDB) {
	versions.Database = db
	versions.Collection = coll
	versions.client = documentdb.New(config.URL, documentdb.Config{config.MasterKey})
	// Find or create `test` db and `versions` collection
	if err := versions.findOrDatabase(db); err != nil {
		panic(err)
	}
	if err := versions.findOrCreateCollection(coll); err != nil {
		panic(err)
	}
	return
}

// Get version by given id
func (v *VersionsDB) Get(id string) (version Version, err error) {
	var versions []Version
	query := fmt.Sprintf("SELECT * FROM ROOT r WHERE r.id='%s'", id)
	if err = v.client.QueryDocuments(v.coll.Self, query, &versions); err != nil || len(versions) == 0 {
		return version, err
	}
	version = versions[0]
	return version, nil
}

// Get all versions
func (v *VersionsDB) GetAll() (versions []Version, err error) {
	err = v.client.ReadDocuments(v.coll.Self, &versions)
	return
}

// Create version
func (v *VersionsDB) Add(version *Version) (err error) {
	return v.client.CreateDocument(v.coll.Self, version)
}

// Update version by id
func (v *VersionsDB) Update(id string, version *Version) (err error) {
	var versions []Version
	if err = v.client.QueryDocuments(v.coll.Self,
		fmt.Sprintf("SELECT * FROM ROOT r WHERE r.id='%s'", id), &versions); err != nil || len(versions) == 0 {
		return
	}
	nVersion := versions[0]
	version.Id = nVersion.Id
	if err = v.client.ReplaceDocument(nVersion.Self, &version); err != nil {
		return
	}
	return
}

// Remove version by id
func (v *VersionsDB) Remove(id string) (err error) {
	var versions []Version
	if err = v.client.QueryDocuments(v.coll.Self,
		fmt.Sprintf("SELECT * FROM ROOT r WHERE r.id='%s'", id), &versions); err != nil || len(versions) == 0 {
		return
	}
	nVersion := versions[0]
	if err = v.client.DeleteDocument(nVersion.Self); err != nil {
		return
	}
	return
}

// Find or create collection by id
func (v *VersionsDB) findOrCreateCollection(name string) (err error) {
	if colls, err := v.client.QueryCollections(v.db.Self, fmt.Sprintf("SELECT * FROM ROOT r WHERE r.id='%s'", name)); err != nil {
		return err
	} else if len(colls) == 0 {
		if coll, err := v.client.CreateCollection(v.db.Self, fmt.Sprintf(`{ "id": "%s" }`, name)); err != nil {
			return err
		} else {

			v.coll = coll
		}

	} else {
		v.coll = &colls[0]
	}
	return
}

// Find or create database by id
func (v *VersionsDB) findOrDatabase(name string) (err error) {
	if dbs, err := v.client.QueryDatabases(fmt.Sprintf("SELECT * FROM ROOT r WHERE r.id='%s'", name)); err != nil {
		return err
	} else if len(dbs) == 0 {
		if db, err := v.client.CreateDatabase(fmt.Sprintf(`{ "id": "%s" }`, name)); err != nil {
			return err
		} else {
			v.db = db
		}
	} else {
		v.db = &dbs[0]
	}
	return
}
