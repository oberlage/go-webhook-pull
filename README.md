# go-webhook-pull

## Prerequisites
This (very basic) tool requires that you have a git repository already set up and authentication to the git server is handled (i.e. via SSH).
To test if this is the case you a simple `git pull` on the command line in the repository should work (the executable should run from within the main folder of the local git repo).

## Compiling
To compile this script use the `go build` command. You can build it for specific OS and architecture, please refer to [this blog post of Dave Cheney][1].

## Usage
Start on commandline by `TOKEN=abc123 PORT=1337 ./go-webhook-pull &`.
You can ommit the env variables (Defaults: no token auth, port 8080).

[1]: https://dave.cheney.net/2015/08/22/cross-compilation-with-go-1-5