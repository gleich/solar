package clone

import (
	"fmt"
	"os"
	"os/exec"

	"code.cloudfoundry.org/bytefmt"
	"github.com/gleich/lumber"
	"github.com/gleich/solar/pkg/api"
)

// Clone a star repo
func Star(star api.Star, starCount int, clonedSoFar int) error {
	if _, err := os.Stat(star.Owner.Login); os.IsNotExist(err) {
		err := os.Mkdir(star.Owner.Login, 0755)
		if err != nil {
			return err
		}
	}
	err := os.Chdir(star.Owner.Login)
	if err != nil {
		return err
	}

	lumber.Info(fmt.Sprintf(`Cloning %[1]v/%[2]v:

	Owner:       %[1]v
	Name:        %[2]v
	Description: %v
	Disk Usage:  %v
	URL:         %v

	Progress:    %v out of %v cloned`,
		star.Owner.Login,
		star.Name,
		star.Description,
		bytefmt.ByteSize(uint64(star.DiskUsage)*bytefmt.KILOBYTE),
		star.URL,
		clonedSoFar,
		starCount,
	))

	cmd := exec.Command("git", "clone", "--depth=1", star.URL+".git")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	err = cmd.Run()
	if err != nil {
		return err
	}

	fmt.Println()
	lumber.Success("Cloned", star.Owner.Login+"/"+star.Name)

	// Go back up
	err = os.Chdir("..")
	if err != nil {
		return err
	}

	return nil
}
