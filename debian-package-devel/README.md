# Debian package development container

This is an container for the *arm* architecture with preinstalled packages for building TorizonCore Debian packages.

This container is prepared to build Debian packages for i.MX devices to
integrate the NXP downstream graphics stack. It is an `aarch64` container which
allows to build the Debian packages natively (either via Qemu user emulation on
`x86_64` or on an actual native device).

## 1. Setup

### 1.1 Assumptions

The following instructions assume:

 * Your user id on your host is **1000** (check using `id -u`).
   (You can adapt the Dockerfile otherwise.)
 * You have
   [created your own Toradex key pair](https://toradex.atlassian.net/wiki/spaces/TOR/pages/529432605/How+to+build+a+Debian+package+for+the+Torizon+platform#Create-your-Toradex-GPG-key-pair)
   for signing packages,
   and your default key ring is stored in `~/.gnupg/`
 * There is a directory `~/debian-pkg` in your host
   where packages will be built.

### 1.2 Building the container

```
$ docker build -t debian-package-devel .
```

## 2. Running the container

```
$ docker run -it --rm \
    --user debian \
    --workdir /home/debian/debian-pkg  \
    -v ~/.gitconfig:/home/debian/.gitconfig \
    -v ~/.gnupg/:/home/debian/.gnupg/ \
    -v ~/debian-pkg/:/home/debian/debian-pkg/ \
    debian-package-devel /bin/bash -l
```

## 3. Workflows

### 3.1 Building packages

Clone a package project from the
[TorizonCore Debian projects](https://gitlab.int.toradex.com/rd/torizon-core/debian)
and follow its build instructions,
or peek the project's `.gitlab-ci.yml` file and the project's CI pipeline
to hint on how the package is built by CI.

Also,
[How to build a Debian package for the Torizon platform](https://toradex.atlassian.net/wiki/spaces/TOR/pages/529432605/How+to+build+a+Debian+package+for+the+Torizon+platform)
offers introductory information regarding building Debian packages for the Torizon platform.

### 3.2 Test the package feed with new packages

After building a
[new version](https://toradex.atlassian.net/wiki/spaces/TOR/pages/364445852/Debian+packages+naming+and+versioning+convention)
of the package,
you may want to confirm that the package artifacts integrate well
to the [Todarex Debian package feed](https://feeds.toradex.com/debian/).

The [debian-package-feed](https://gitlab.int.toradex.com/rd/torizon-core/debian/debian-package-feed)
project can regenerate the feed with your updates.
Clone this project and follow its instructions to generate a test feed in the debian-package-devel container
that exports your new packages.

### 3.3 Access the test package feed from another container

If the test feed was started inside the debian-package-devel container as instructed,
by default it's not reachable outside it.
In order to reach the test feed instance (TCP port 8080 by default) of the debian-package-devel container
from another container (so you can test it),
you can run the test container using the `--link` option:

```
$ docker run -it --link=<ID> arm64v8/debian:bookworm /bin/bash
```
where `<ID>` is the id of the debian-package-devel container that is serving the feed locally.
The test feed can then be accessed by `http://<ID>:8080` in the test container.

Once accessing the feed, download its README file
and follow the instructions to enable the feed in your test container.
(Note that this file assumes the official feed,
so adapt the instructions for accessing your testing feed instead.)
