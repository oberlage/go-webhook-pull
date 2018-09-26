# go-webhook-pull

go-webhook-pull is a basic application that can be used to pull-on-push your code from a git repository.
If you want easy uncomplicated deployment from your git repository when somebody pushes code to it, this will serve you well.

You can configure the webhook of your git provider so that the HTTP endpoint is called upon a push to your repository,  this program will then perform `git pull` on your system to load the latest version of your code from the repository.

### Example Use Case
This program was written to pull PHP/JS/HTML/CSS code from a repository once it is updated. As these are all interpreted languages, just replacing the code with the latest version in the repository will do the job. You can easily keep the code of your website or web application in sync with the code in your repository with this script.

**Please note:** this application is very basic, and therefore (currently) not suited for performing build steps or running a (unit-)test framework before deployment. If you need this to be automated, you might want to look at something like [Jenkins][3] or contribute to this project to make it possible.

## Compiling
See the releases page in this repository for pre-compiled version. In case you want to compile yourself or for other platforms you need the Go tools for compiling to compile. To compile this script use the `./build.sh` script that is included in the repo. You can append your specific OS and architecture, please refer to [this blog post of Dave Cheney][1] for more info.

## Setup
You have to setup your git repository as usual (by using `git clone` and checking out the branch you want). This script assumes that authentication is handled, i.e. via SSH keys. 

To test if your repository is ready for this script a simple `git pull` from the command line should work without any errors.

## Running the program
You can run this program in several ways of which two are shown below:
### As backgroundprocess
Start on commandline by `TOKEN=abc123 PORT=1337 PULLPATH="public_html" ./go-webhook-pull &`.
You can ommit the env variables (Defaults: no token auth, port 8080, pullpath path of executable).


### As System Service
An example `systemd` service file is included (see `gowebhookpull.service` and `run_gowebhookpull.sh`). Setting it up as a service also handles automatic start of the program when the server is rebooted e.d. See [this blog post from @benmorel on Medium][2] for more info.

## Endpoints

| Method | Endpoint | Description |
| --- | --- | --- |
| `GET` `POST` | `http://X.X.X.X:PORT/pull` | Trigger to perform a git pull action |
| `GET` | `http://X.X.X.X:PORT/info` | Show info about this program (such as the current version) |

When a token is configured, you need to provide the token for all endpoints to access them.

# (Not so) Frequently Asked Questions
**Q:** Is HTTPS supported?
> Answer: No. If you want this, you can put the application behind a [nginx reverse-proxy with SSL][4]. Or contribute to the code to make `gin-gonic` work with HTTPS connections. 

[1]: https://dave.cheney.net/2015/08/22/cross-compilation-with-go-1-5
[2]: https://medium.com/@benmorel/creating-a-linux-service-with-systemd-611b5c8b91d6
[3]: https://jenkins.io
[4]: https://docs.nginx.com/nginx/admin-guide/security-controls/securing-http-traffic-upstream/
