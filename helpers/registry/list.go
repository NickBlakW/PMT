package registry

import "fmt"

func ListRegisteredProjects(verbose bool) (string, error) {
	projects, err := GetRegisteredProjects()
	if err != nil {
		return "", err
	}

	var output string

	for k, v := range projects {
		if k == "" { break }

		output += k

		if verbose { output += fmt.Sprintf("\t\t@  %s", v) }
	}

	return output, nil
}