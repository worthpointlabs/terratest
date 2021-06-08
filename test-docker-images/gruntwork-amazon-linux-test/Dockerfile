# TODO: Is it worth referencing a specific tag instead of latest?
FROM amazonlinux:2

# Reduce Docker image size per https://blog.replicated.com/refactoring-a-dockerfile-for-image-size/
# - perl-Digest-SHA: installs shasum
RUN yum update -y && \
    yum upgrade -y && \
    yum install -y \
        hostname \
        jq \
        perl-Digest-SHA \
        rsyslog \
        sudo \
        tar \
        vim \
        python3 \
        pip3 \
        wget && \
        yum clean all && rm -rf /var/cache/yum

# Upgrade pip and install awscli
RUN pip3 install --upgrade pip && pip3 install awscli

# Ideally, we'd install the latest version of Docker to avoid a conflict between the Docker client in this container
# and the Docker API on your local host, but installing the latest version of Docker yields the error "Requires:
# container-selinux >= 2.9", whch indicates that a newer Linux kernel version is required than what comes with Amazon Linux.
# So we settle for the Amazon Linux supported version for now.
RUN amazon-linux-extras install docker
