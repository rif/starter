# Starter App

This is a basic Go web app suitable for deployment on [heroku][] or[openshift].
It uses the
[heroku-buildpack-go][] buildpack to compile and deploy on the Cedar stack.

## Usage

Clone this app into some directory and remove the git directory.

## Local compilation and testing

    $ ./init.sh
    $ ./build.sh
    $ PORT=9022 bin/srv

Your app is now serving requests on http://localhost:9022

[heroku]: http://heroku.com
[heroku-buildpack-go]: http://github.com/zeebo/heroku-buildpack-go
[openshift]: http://openshift.redhat.com

