ARG ARCH_ARG=armhf
# For arm64 use:
# ARG ARCH_ARG=arm64
ARG IMAGE_TAG=2-bullseye

FROM torizon/debian-cross-toolchain-$ARCH_ARG:$IMAGE_TAG

# SSH
EXPOSE 22

# Install required packages
RUN apt-get -q -y update \
    && apt-get -q -y install --no-install-recommends \
    openssh-server \
    ssh \
    rsync \
    zip \
    && apt-get clean \
    && apt-get autoremove -y \
    && rm -rf /var/lib/apt/lists/*

# create folders needed for the different components
RUN mkdir /var/run/sshd && mkdir /home/build && mkdir /home/build/sysroot

# configures SSH access to the container and sets environment by default
RUN sed 's@session\s*required\s*pam_loginuid.so@session optional pam_loginuid.so@g' -i /etc/pam.d/sshd \
    && useradd -s /bin/bash -d /home/build -g root -G sudo,ssh -p $(echo "build" | openssl passwd -1 -stdin) build \
    && chown -R build /home/build

CMD ["/usr/sbin/sshd", "-D"]
