FROM registry.access.redhat.com/ubi7/ubi-minimal:latest

ENV OPERATOR=/usr/local/bin/controllerv2 \
    USER_UID=1001 \
    USER_NAME=controllerv2

# install operator binary
COPY build/_output/bin/controllerv2 ${OPERATOR}

COPY build/bin /usr/local/bin
RUN  /usr/local/bin/user_setup

ENTRYPOINT ["/usr/local/bin/entrypoint"]

USER ${USER_UID}
