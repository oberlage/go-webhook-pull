# go-webhook-pull

go-webhook-pull is a very simple script that can be used to pull-on-push from a git repository.
If you want easy uncomplicated deployment from your git repository, this will serve you well.

You can configure your webhook so that this listener is called upon a push to your repository,  this program will then perform `git pull` to load the latest version of your repository.

## Compiling
You need the Go tools for compiling to compile. To compile this script use the `./build.sh` script that is included in the repo. You can append your specific OS and architecture, please refer to [this blog post of Dave Cheney][1] for more info.

## Setup
You have to setup your git repository as usual (by using `git clone` and checking out the branch you want). This script assumes that authentication is handled, i.e. via SSH keys. 

To test if your repository is ready for this script a simple `git pull` from the command line should work without any errors.

## Running the program
You can run this program in several ways of which two are shown below:
### As backgroundprocess
Start on commandline by `TOKEN=abc123 PORT=1337 PULLPATH="public_html" ./go-webhook-pull &`.
You can ommit the env variables (Defaults: no token auth, port 8080, pullpath path of executable).


### As System Service
An example `systemd` service file is included (see `gowebhookpull.service` and `run_gowebhookpull.sh`). Setting it up as a service also handles automatic start of the program when the server is rebooted e.d. See [this blog post from @benmorel on Medium][1] for more info.

## Endpoints

| Method | Endpoint | Description |
| --- | --- | --- |
| `GET` `POST` | `http://X.X.X.X:PORT/pull` | Trigger to perform a git pull action |
| `GET` | `http://X.X.X.X:PORT/info` | Show info about this program (such as the current version) |

When a token is configured, you need to provide the token for all endpoints to access them.


[1]: https://dave.cheney.net/2015/08/22/cross-compilation-with-go-1-5
[2]: https://medium.com/@benmorel/creating-a-linux-service-with-systemd-611b5c8b91d6
