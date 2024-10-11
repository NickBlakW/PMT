package utils

type PMTConfig struct {
	Name 		string					`json:"name,omitempty"`
	Author		string 					`json:"author,omitempty"`
	Scripts 	map[string]interface{}	`json:"scripts,omitempty"`
	Backend 	string					`json:"backend,omitempty"`
	Frontend	string					`json:"frontend,omitempty"`
}

var DefaultScripts = map[string]map[string]interface{}{
	"go": {
		"run": "go run %s",
		"compile": "go build %s && echo 'Compiled successfully'",
	},
	"npm": {
		"dev": "npm run dev --prefix %s",
		"build": "npm run build --prefix %s",
		"start": "npm run start --prefix %s",
	},
	"yarn": {
		"dev": "yarn -cwd %s dev",
		"build": "yarn -cwd %s build",
		"start": "yarn -cwd %s start",
	},
	"maven": {
		"ci": "mvn clean install -f %s/pom.xml",
		"package": "mvn package -f %s/pom.xml",
		"run": "mvn clean compile exec:java -f %s/pom.xml",
	},
}