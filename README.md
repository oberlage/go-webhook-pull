# go-webhook-pull

go-webhook-pull is a very simple script that can be used to pull-on-push from a git repository.
You can configure your webhook so that this listener is called, and it wil `git pull` from the repository.

## Compiling
To compile this script use the `./build.sh` script that is included in the repo. You can append your specific OS and architecture, please refer to [this blog post of Dave Cheney][1] for more info.

## Setup
You have to setup your git repository as usual (by using `git clone` and checking out the branch you want). This script assumes that authentication is handled, i.e. via SSH keys. 

To test if your repository is ready for this script a simple `git pull` from the command line should work without any errors.

## Usage
Start on commandline by `TOKEN=abc123 PORT=1337 PULLPATH="public_html" ./go-webhook-pull &`.
You can ommit the env variables (Defaults: no token auth, port 8080, pullpath path of executable).

You can configure your webhooks to point to: `http://X.X.X.X:PORT/pull` (either GET or POST)
To check the current version of this listener: `http://X.X.X.X:PORT/info` (GET)

[1]: https://dave.cheney.net/2015/08/22/cross-compilation-with-go-1-5
