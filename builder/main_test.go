package main

import (
	"os"
	"reflect"
	"testing"
)

func TestParseDockerfileLabels(t *testing.T) {
	tempDockerfile := "testDockerfile"
	defer func() {
		// Clean up the temporary Dockerfile after the test
		if err := os.Remove(tempDockerfile); err != nil {
			t.Errorf("Error cleaning up temporary Dockerfile: %v", err)
		}
	}()

	dockerfileContent := `FROM ubuntu
LABEL org.toradex.image.base.registry="docker.io" \
      org.toradex.image.base.namespace="library" \
      org.toradex.image.base.name="debian" \
      org.toradex.image.base.tag.major="12" \
      org.toradex.image.base.tag.minor="4" \
      org.toradex.image.base.tag.patch="" \
      org.toradex.image.variant="slim" \
      org.toradex.image.registry="docker.io" \
      org.toradex.image.namespace="torizon" \
      org.toradex.image.name="debian" \
      org.toradex.image.tag.major="rc" \
      org.toradex.image.tag.minor="" \
      org.toradex.image.tag.patch="" \
      org.toradex.image.tag.variant="bookworm" \
      org.toradex.image.license="MIT" \
      org.toradex.image.arch="linux/amd64,linux/arm64/v8,linux/arm/v7"
COPY kms-setup.sh /usr/bin/kms-setup.sh

RUN apt-get -y update && apt-get install -y --no-install-recommends \
    apt-utils \
    && apt-get -y upgrade \
    && apt-get clean && apt-get autoremove && rm -rf /var/lib/apt/lists/*

# Install libqt5gui5-gles before libqt5opengl5 which also has an alternate dependency on libqt5gui5(non-gles)
RUN apt-get -y update  && apt-get install -y --no-install-recommends libqt5gui5-gles \
    && apt-get clean && apt-get autoremove && rm -rf /var/lib/apt/lists/*

# Install libqt5opengl5.
# Under Bookworm, libqt5opengl5 is currently at version 5.15.8+dfsg-11 and contains a redundant dependency for libqt5gui5 (>= 5.1.0):
#
# This forbids installing qtbase5-examples with libqt5gui5-gles.
# Workaround the issue by mangling the package file and remove the leftover dependency for each of the architectures
RUN apt-get -y update \
    && ARCH=$(dpkg --print-architecture) \
    && if test "$ARCH" = 'arm64' ; \
    then \
        apt-get -y install --no-install-recommends binutils xz-utils \
        && WORK_DIR=$(mktemp -d) \
        && cd $WORK_DIR \
        && apt-get download libqt5opengl5:$ARCH \
        && ar x libqt5opengl5_*_$ARCH.deb \
        && tar -xJf control.tar.xz \
        && sed -i '/^Depends:/s/, libqt5gui5 (>= 5.1.0)//' control \
        && tar -cJf control.tar.xz control md5sums shlibs symbols triggers \
        && ar rcs libqt5opengl5.deb debian-binary control.tar.xz data.tar.xz \
        && apt-get -y install --no-install-recommends ./libqt5opengl5.deb \
        && cd ~ \
        && rm -rf $WORK_DIR \
        && apt-get -y remove binutils xz-utils \
        && apt-mark hold libqt5opengl5 ; \
    elif test "$ARCH" = 'x86_64' ; \
    then \
        apt-get -y install --no-install-recommends binutils xz-utils \
        && WORK_DIR=$(mktemp -d) \
        && cd $WORK_DIR \
        && apt-get download libqt5opengl5:$ARCH \
        && ar x libqt5opengl5_*_$ARCH.deb \
        && tar -xJf control.tar.xz \
        && sed -i '/^Depends:/s/, libqt5gui5 (>= 5.1.0)//' control \
        && tar -cJf control.tar.xz control md5sums shlibs symbols triggers \
        && ar rcs libqt5opengl5.deb debian-binary control.tar.xz data.tar.xz \
        && apt-get -y install --no-install-recommends ./libqt5opengl5.deb \
        && cd ~ \
        && rm -rf $WORK_DIR \
        && apt-get -y remove binutils xz-utils \
        && apt-mark hold libqt5opengl5 ; \
    else \
        apt-get -y install --no-install-recommends binutils xz-utils \
        && WORK_DIR=$(mktemp -d) \
        && cd $WORK_DIR \
        && apt-get download libqt5opengl5:$ARCH \
        && ar x libqt5opengl5_*_$ARCH.deb \
        && tar -xJf control.tar.xz \
        && sed -i '/^Depends:/s/, libqt5gui5 (>= 5.1.0)//' control \
        && tar -cJf control.tar.xz control md5sums shlibs symbols triggers \
        && ar rcs libqt5opengl5.deb debian-binary control.tar.xz data.tar.xz \
        && apt-get -y install --no-install-recommends ./libqt5opengl5.deb \
        && cd ~ \
        && rm -rf $WORK_DIR \
        && apt-get -y remove binutils xz-utils \
        && apt-mark hold libqt5opengl5 ; \
    fi \
    && apt-get clean && apt-get autoremove && rm -rf /var/lib/apt/lists/*

# Install remaining dependencies required to run qtbase and qtdeclarative examples
RUN apt-get -y update && apt-get install -y --no-install-recommends \
        libfontconfig1-dev \
        libqt5quick5-gles \
        libqt5quickparticles5-gles \
        libqt5concurrent5 \
        libqt5dbus5 \
        libqt5network5 \
        libqt5printsupport5 \
        libqt5sql5 \
        libqt5test5 \
        libqt5widgets5 \
        libqt5xml5 \
        libqt5qml5 \
        libqt5quicktest5 \
        libqt5quickwidgets5 \
        qml-module-qt-labs-qmlmodels \
        qml-module-qtqml-models2 \
        qml-module-qtquick-layouts \
        qml-module-qtquick-localstorage \
        qml-module-qtquick-particles2 \
        qml-module-qtquick-shapes \
        qml-module-qttest \
    && apt-get clean && apt-get autoremove && rm -rf /var/lib/apt/lists/*

# Install Wayland Qt module
RUN apt-get -y update && apt-get install -y --no-install-recommends \
    qtwayland5 \
    && apt-get clean && apt-get autoremove && rm -rf /var/lib/apt/lists/*

ENV QT_QPA_PLATFORM="wayland"

# EGLFS configuration
ENV QT_QPA_EGLFS_INTEGRATION="eglfs_kms"
ENV QT_QPA_EGLFS_KMS_ATOMIC="1"
ENV QT_QPA_EGLFS_KMS_CONFIG="/etc/kms.conf"

`

	if err := os.WriteFile(tempDockerfile, []byte(dockerfileContent), 0644); err != nil {
		t.Fatalf("Error creating temporary Dockerfile: %v", err)
	}

	labels, err := parseDockerfileLabels(tempDockerfile)
	if err != nil {
		t.Errorf("Error parsing Dockerfile labels: %v", err)
	}

	expectedLabels := map[string]string{
		"org.toradex.image.base.registry":  "docker.io",
		"org.toradex.image.base.namespace": "library",
		"org.toradex.image.base.name":      "debian",
		"org.toradex.image.base.tag.major": "12",
		"org.toradex.image.base.tag.minor": "4",
		"org.toradex.image.base.tag.patch": "",
		"org.toradex.image.variant":        "slim",
		"org.toradex.image.registry":       "docker.io",
		"org.toradex.image.namespace":      "torizon",
		"org.toradex.image.name":           "debian",
		"org.toradex.image.tag.major":      "rc",
		"org.toradex.image.tag.minor":      "",
		"org.toradex.image.tag.patch":      "",
		"org.toradex.image.tag.variant":    "bookworm",
		"org.toradex.image.license":        "MIT",
		"org.toradex.image.arch":           "linux/amd64,linux/arm64/v8,linux/arm/v7",
	}

	if !reflect.DeepEqual(labels, expectedLabels) {
		t.Errorf("Parsed labels do not match expected labels. Got: %v, Expected: %v", labels, expectedLabels)
	}
}
