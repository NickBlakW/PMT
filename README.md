# PMT (Project Management Tool)

<p align="center">
    <a href="https://github.com/NickBlakW/PMT/blob/main/LICENSE">
        <img src="https://img.shields.io/github/license/NickBlakW/PMT" alt="license" />
    </a>
    <a href="https://github.com/NickBlakW/PMT/actions/workflows/build.yml" rel="nofollow">
        <img src="https://img.shields.io/github/actions/workflow/status/NickBlakW/PMT/build.yml?branch=main&logo=Github" alt="Build" />
    </a>
    <a href="https://github.com/NickBlakW/PMT/releases" rel="nofollow">
        <img alt="release" src="https://img.shields.io/github/v/tag/NickBlakW/PMT?include_prereleases&label=version"/>
    </a>
</p>

Inspired by the functionality of node.js and it's script execution from the 'package.json' file this CLI tool serves to make the tedious task of writing console commands much easier and simpler.

## Setup

Download the [release](https://github.com/NickBlakW/PMT/releases) appropriate for your system. If you are on windows make sure to add the binary path to the environment path variable.
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

#### Registry commands

Access the internal project registry in PMT\
The registry is a text file that will be created at a path like this `~/users/<user>/.pmt/registry.pmt`, the file can be modified but it is not recommended to do so.

Aliases:\
`registry`, `reg`

Args:\
`add` registers current initialized project in the registry\
`list` lists all registered projects in the registry\
`remove` removes a registered project from the registry

Flags:\
`-a, --alt` register project with an alternative key\
`-h, --help` help for registry
`-v, --verbose` Shows verbose messages (only for list command)

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

For now it is **NOT** recommended to use PMT for dev servers as there is no live printouts of errors and logs in the commandline.

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

## Future

-   Adding support for a YAML generator for docker-compose
-   More options for registry commands
