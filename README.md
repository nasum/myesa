# myesa

myesa is command line esa client

## How to install

Fetch from GitHub and install.

```zsh
$ go get github.com/nasum/myesa
$ go install
```

create `myesarc.json` in your `$HOME`

```json
{
    "ACCESS_TOKEN" : "access token",
    "TEAM" : "team name"
}
```

## How to use

please see help

```zsh
$ myesa help
esa client

Usage:
  myesa [command]

Available Commands:
  edit        edit articles
  help        Help about any command
  search      search articles
  show        show articles

Flags:
  -h, --help   help for myesa

Use "myesa [command] --help" for more information about a command.


```
