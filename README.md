# Torizon Containers

This repository contains some "base" container images used by or with [Torizon OS](https://www.torizon.io/torizon-os),
the Easy-to-use Industrial Linux Software Platform.

These base images serve as a foundation for other open source projects such as the [VSCode Torizon Templates](https://github.com/toradex/vscode-torizon-templates) and the [Torizon Samples](https://github.com/toradex/torizon-samples) and will always be based on the current [Stable Debian](https://www.debian.org/releases/stable/) release.

As a rule of thumb, if a container needs hardware acceleration or some specific tricks to ensure smooth working conditions, we maintain it.

## Building

Although you can build the images locally, it's highly recommended to use the CI infrastructure instead (`.gitlab-ci.yml`).

Images that are platform-dependent (contain hardware specific software) have a `-<platform>` naming append, where `platform` is the name of the SoC such as iMX8 or AM62. If there's no `-<platform>`, that means we only use upstream software, and no third-party package feed.

Each platform has its own Debian Package Feed, which are hosted here:
`https://feeds.toradex.com/`.

For comprehensive documentation, check out our developer website:
`https://developer.toradex.com/torizon/provided-containers/debian-containers-for-torizon/`.

## Branch Scheme

Our release branch is always the name of the current Debian Stable.
Changes are rebased on top of this release branch from the `rc` (for `Release Candidate`) branch.

The reason why we don't do point-releases from `rc` directly is that we can re-run the release pipeline if a new version of Debian comes out. That includes a bump in the minor version of the tags defined in the [.gitlab-ci.yml](.gitlab-ci.yml).

We also have two branches based on the Stable branch but using the Unstable (`sid`) and Testing distributions of Debian instead of Stable. This allows us to easily release a new stable version from testing whenever the upstream promotes it. Whenever a change is done in the stable branch it should be cherry picked to the `sid` and the current testing branches!

## Where can I get the images?

All images are publish to Torizon's DockerHub. Images are by default build as multi-arch, so the same image name/tag can be used for different platforms (except where indicated).

Please take a look at the [Developer Documentation](https://developer.toradex.com/torizon/provided-containers/list-of-container-images-for-torizon/) for more comprehensive information on using the containers.

## SBOM

SBOM (Software Bill of Materials) is generated using [Anchore's Syft](https://github.com/anchore/syft) for every image and pushed as a OCI artifact to DockerHub.

To download the SBOM, use `imagetools` from `buildx` as follows:

```
docker buildx imagetools inspect torizon/wayland-base:rc --format '{{ json (index .SBOM "linux/arm64").SPDX.packages }}'
```

You can further parse it using Go statements, such as

```
docker buildx imagetools inspect torizon/wayland-base:rc --format '{{ range (index .SBOM "linux/arm64").SPDX.packages}}{{println .name .versionInfo}}{{end}}' | sort
```

Which will print all packages installed and their versions.

## CVE

CVE are generated in the Pipeline but can be also generated using tools such as [Trivy](https://github.com/aquasecurity/trivy) from the published images:

```
docker run --rm --privileged -v /var/run/docker.sock:/var/run/docker.sock bitnami/trivy image --no-progress --exit-code 0 --platform linux/arm/v7 torizon/debian:rc-bookworm
```

## Contributing

We're open for contributions! Please feel free to open a merge request and/or raise an issue!

## Todo

- [ ] Easily build a given image outside of CI for development (multiple context builds)
- [ ] RISC-V images
- [ ] GitHub Actions

## License

This project is licensed under the terms of MIT license - see `LICENSE`.
