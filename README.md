# UnOfficial TPLink Smart Switches SG10X series command line

I created this utility initially to retrieve information in an scriptted way of my TPLink switches, and create scheduled backups.

## Usage

```raw
Interacts with your TLSG switch from command line

Usage:
  tlsg-cli [command]

Available Commands:
  backup      Creates a configuration backup
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  info        Prints Switch System information
  ports       Display Port Settings
  stats       Retrieve Port Statistics

Flags:
  -h, --help              help for tlsg-cli
  -z, --host string       Switch hostname or IP (default "localhost")
  -p, --password string   Switch login password (default "admin")
  -u, --username string   Switch login username (default "admin")
  -v, --version           version for tlsg-cli

Use "tlsg-cli [command] --help" for more information about a command.
```
