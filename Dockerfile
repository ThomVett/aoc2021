# Build Stage
FROM lacion/alpine-golang-buildimage:1.13 AS build-stage

LABEL app="build-aoc2021"
LABEL REPO="https://github.com/ThomVett/aoc2021"

ENV PROJPATH=/go/src/github.com/ThomVett/aoc2021

# Because of https://github.com/docker/docker/issues/14914
ENV PATH=$PATH:$GOROOT/bin:$GOPATH/bin

ADD . /go/src/github.com/ThomVett/aoc2021
WORKDIR /go/src/github.com/ThomVett/aoc2021

RUN make build-alpine

# Final Stage
FROM lacion/alpine-base-image:latest

ARG GIT_COMMIT
ARG VERSION
LABEL REPO="https://github.com/ThomVett/aoc2021"
LABEL GIT_COMMIT=$GIT_COMMIT
LABEL VERSION=$VERSION

# Because of https://github.com/docker/docker/issues/14914
ENV PATH=$PATH:/opt/aoc2021/bin

WORKDIR /opt/aoc2021/bin

COPY --from=build-stage /go/src/github.com/ThomVett/aoc2021/bin/aoc2021 /opt/aoc2021/bin/
RUN chmod +x /opt/aoc2021/bin/aoc2021

# Create appuser
RUN adduser -D -g '' aoc2021
USER aoc2021

ENTRYPOINT ["/usr/bin/dumb-init", "--"]

CMD ["/opt/aoc2021/bin/aoc2021"]
