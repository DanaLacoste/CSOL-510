# Kali Linux in Docker

OK, setting up Virtual Box is horrible.  It doesn't work very well and, frankly, you don't need it for this course.

Running Docker on Mac and on Windows actually loads a virtual machine for linux containers: this means you get all the niceties of running a VM, but with the ease and simplicity of using docker.

## How to use

1. Check out this repo
2. Install docker from https://docs.docker.com/get-docker/
3. From your command line (in this directory), run `docker build -t csol-510-kali .`
4. Once it has built (this is building a brand new kali image from the current builds, so it can take a while), run the image with `docker run -it -v `pwd`:/CSOL-510/work csol-510-kali`

There you go: you now have a fully functioning kali linux install and it took like 3 minutes.
When you exit, it will terminate the container.