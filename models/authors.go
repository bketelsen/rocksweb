package models

import (
	"fmt"

	"github.com/a8m/documentdb-go"
)

// Author document
type Author struct {
	documentdb.Document
	Name  string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
}

// DB interface
type AuthorDB interface {
	Get(id string) *Author
	GetAll() []*Author
	Add(a *Author) *Author
	Update(a *Author) *Author
	Remove(id string) error
}

// AuthorsDB implement DB interface
type AuthorsDB struct {
	Database   string
	Collection string
	db         *documentdb.Database
	coll       *documentdb.Collection
	client     *documentdb.DocumentDB
}

// Return new AuthorDB
// Test if database and collection exist. if not, create them.
func NewAuthorDB(db, coll string, config *Config) (authors AuthorsDB) {
	authors.Database = db
	authors.Collection = coll
	authors.client = documentdb.New(config.URL, documentdb.Config{config.MasterKey})
	// Find or create `test` db and `authors` collection
	if err := authors.findOrDatabase(db); err != nil {
		panic(err)
	}
	if err := authors.findOrCreateCollection(coll); err != nil {
		panic(err)
	}
	return
}

// Get user by given id
func (a *AuthorsDB) Get(id string) (author Author, err error) {
	var authors []Author
	query := fmt.Sprintf("SELECT * FROM ROOT r WHERE r.id='%s'", id)
	if err = a.client.QueryDocuments(a.coll.Self, query, &authors); err != nil || len(authors) == 0 {
		return author, err
	}
	author = authors[0]
	return author, nil
}

// Get all authors
func (a *AuthorsDB) GetAll() (authors []Author, err error) {
	err = a.client.ReadDocuments(a.coll.Self, &authors)
	return
}

// Create user
func (a *AuthorsDB) Add(author *Author) (err error) {
	return a.client.CreateDocument(a.coll.Self, author)
}

// Update user by id
func (a *AuthorsDB) Update(id string, author *Author) (err error) {
	var authors []Author
	if err = a.client.QueryDocuments(a.coll.Self,
		fmt.Sprintf("SELECT * FROM ROOT r WHERE r.id='%s'", id), &authors); err != nil || len(authors) == 0 {
		return
	}
	nAuthor := authors[0]
	author.Id = nAuthor.Id
	if err = a.client.ReplaceDocument(nAuthor.Self, &author); err != nil {
		return
	}
	return
}

// Remove user by id
func (a *AuthorsDB) Remove(id string) (err error) {
	var authors []Author
	if err = a.client.QueryDocuments(a.coll.Self,
		fmt.Sprintf("SELECT * FROM ROOT r WHERE r.id='%s'", id), &authors); err != nil || len(authors) == 0 {
		return
	}
	nAuthor := authors[0]
	if err = a.client.DeleteDocument(nAuthor.Self); err != nil {
		return
	}
	return
}

// Find or create collection by id
func (a *AuthorsDB) findOrCreateCollection(name string) (err error) {
	if colls, err := a.client.QueryCollections(a.db.Self, fmt.Sprintf("SELECT * FROM ROOT r WHERE r.id='%s'", name)); err != nil {
		return err
	} else if len(colls) == 0 {
		if coll, err := a.client.CreateCollection(a.db.Self, fmt.Sprintf(`{ "id": "%s" }`, name)); err != nil {
			return err
		} else {
			a.coll = coll
		}

	} else {
		a.coll = &colls[0]
	}
	return
}

// Find or create database by id
func (a *AuthorsDB) findOrDatabase(name string) (err error) {
	if dbs, err := a.client.QueryDatabases(fmt.Sprintf("SELECT * FROM ROOT r WHERE r.id='%s'", name)); err != nil {
		return err
	} else if len(dbs) == 0 {
		if db, err := a.client.CreateDatabase(fmt.Sprintf(`{ "id": "%s" }`, name)); err != nil {
			return err
		} else {
			a.db = db
		}
	} else {
		a.db = &dbs[0]
	}
	return
}
