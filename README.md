# PMT (Project Management Tool)

Inspired by the functionality of node.js and it's script execution from the 'package.json' file this CLI tool serves to make the tedious task of writing console commands much easier and simpler.

## Setup

Download the release appropriate for your system. If you are on windows make sure to add the binary path to the environment path variable.
Run the command `pmt -h` to see if it has been successfully added.

## Use

PMT can help with a bunch of tasks such as adding or removing directories and files, executing scripts from the CLI and opening the project with VSCode (might add additional support for IDEs in the future).

PMT works with a JSON file `PMTConfig.json` at the root of the project, this can be auto-generated with the `pmt init [ flags ]` command.

### Commands

All commands uses the pattern `pmt [ command ] [ flags ]`.

#### Initialize

`pmt initialize [flags]`\
Initialize a project with custom settings

Aliases:\
`initialize`, `init`

Flags:\
`-a, --author` string Sets project author (default "\<Your name here>")\
`-b, --backend` string Sets project backend dir\
`-f, --frontend` string Sets project frontend dir\
`-h, --help` help for initialize\
`-n, --name` string Sets project name (default "<Project name here>")\
`-w, --with` stringArray Initializes project with chosen option

#### Execute command

`pmt exec [flags]`\
Execute a cli script from the project

Aliases:\
`exec`, `run`, `do`

Flags:\
`-h, --help` help for exec

Example:\
`pmt do compile`, this will compile the project if using Go\
`pmt run dev`, this will start a dev server for a node.js project

#### Add command

`pmt add [flags]`\
Adds a file or directory to the project

Example:\
`pmt add example/`, adds a new directory, with the name 'example', at the root location.\
`pmt add example/file.txt`, adds a new file in the example directory. If directory does not exist then it will create that as well.

#### Remove commands

`pmt remove [flags]`

removes a file or directory to the project

Aliases:\
`rem`, `delete`, `del`

Example:\
`pmt rem example/`, removes a directory, with the name 'example', and all of its contents.\
`pmt delete example/file.txt`, removes a file in the example directory.\
`pmt del example/*.txt`, removes all files in the example directory with the extension '.txt'.
