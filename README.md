# go-bargain
a bargaining plugin for e-commerce platforms


baseDockerFile:
To setup a server with toolkit to build, lint and test the app.

dockerFile:
Spins up a container from scratch (minus the go compiler and the tool).
This is where the compiled application is copied and runs.

github workflow:
Runs the python script to setup the env config.
Sets up aws access keys.
Builds and runs the baseDockerFile.


# build the image from the baseDockerfile (follow semantic versioning for tagging)
$ DOCKER_BUILDKIT=1 docker build -f scripts/baseDockerfile -t <path in cloud>/builder:1.1.0 .
