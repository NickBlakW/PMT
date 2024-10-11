package helpers

import (
	"errors"
	"strings"

	"github.com/NickBlakW/pmt/helpers/utils"
)

func GetProjectAuthors() (string, error) {
	cfg, err := utils.LoadConfig()
	if err != nil {
		return "PMT not initialized", err
	}

	cfgAuthor := cfg.Author
	authors := strings.Split(cfgAuthor, ",")

	if len(authors) == 1 {
		e := `Author(s) not set, please add an author to the project.
Use command 'pmt cfg mod -a [name]'`
		return "", errors.New(e)
	}

	for i, a := range authors {
		if strings.HasPrefix(a, "<") && strings.HasSuffix(a, ">") {
			err = errors.New("author not set for project")
			return "Please set author name or use 'PMT init' with the '-a' flag", err
		}

		authors[i] = strings.TrimSpace(authors[i])
	}

	return strings.Join(authors, "\n"), nil
}
