package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"

	"github.com/bketelsen/rocksweb/models"

	"golang.org/x/tools/go/vcs"
)

var packageDB models.PackagesDB
var versionDB models.VersionsDB
var authorDB models.AuthorsDB

type results struct {
	Packages []pack `json:"results"`
}

type pack struct {
	ImportCount int    `json:"import_count"`
	Path        string `json:"path"`
}

func main() {
	f, err := os.Open("/home/bketelsen/packages.json")
	if err != nil {
		log.Fatal(err)
	}
	b, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}

	var result results
	err = json.Unmarshal(b, &result)

	if err != nil {
		log.Fatal(err)
	}
	config := &models.Config{
		URL:       "",
		MasterKey: "",
	}
	packageDB = models.NewPackageDB("rocksweb", "packages", config)
	versionDB = models.NewVersionDB("rocksweb", "versions", config)
	authorDB = models.NewAuthorDB("rocksweb", "authors", config)

	x := 0

	for _, pkg := range result.Packages {
		fmt.Println("Processing: ", pkg.Path)

		// clone repository
		root, err := vcs.RepoRootForImportPath(pkg.Path, true)
		if err != nil {
			log.Println(err)
			continue
		}

		err = root.VCS.Create(filepath.Join(os.TempDir(), strconv.Itoa(x)), root.Repo)

		if err != nil {
			log.Println(err)
			continue
		}

		// git info

		//repo, err := git.OpenRepository(filepath.Join(os.TempDir(), strconv.Itoa(x), ".git"))
		//if err != nil {
		//	log.Fatal(err)
		//}

		pakg := models.Package{}
		pakg.ImportPath = pkg.Path
		pakg.SourceURL = root.Repo
		err = packageDB.Add(&pakg)

		if err != nil {
			log.Println(err)
			continue
		}
		x++
		fmt.Println(x, "completed:", pkg.Path)
	}
}
