package actions

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/gobuffalo/buffalo/worker"
	"github.com/pkg/errors"
)

func init() {
	fmt.Println("Registered")
	App().Worker.Register("index_package", func(args worker.Args) error {
		repository, ok := args["repository"].(string)
		if !ok {
			return errors.New("No repository specified")
		}

		// clone the repository in a tmp dir
		tempdir := os.TempDir()

		cmd := exec.Command("git", "clone", repository)
		cmd.Dir = tempdir
		fmt.Println(tempdir)
		stdoutStderr, err := cmd.CombinedOutput()
		err = cmd.Run()
		if err != nil {
			fmt.Println(err, stdoutStderr)
			return errors.Wrap(err, "executing git clone")
		}

		fmt.Println(err, stdoutStderr)
		fmt.Println("Git clone done")
		// get lots of information about it

		// create information to upload to the database

		// insert it

		/*
			repo, err := git.OpenRepository("./.git")
			if err != nil {
				fmt.Println("The current directory doesn't appear to be a git repository.")
				return err
			}
			tags, err := repo.GetTags()
			if err != nil {
				fmt.Println("Error Getting Tags: ", err)
				return err
			}
			fmt.Println(tags)
		*/
		return nil
	})

}
