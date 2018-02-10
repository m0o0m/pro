# #
# Author: wrchen <cwr0401@gmail.com>
# Go version: 1.9.2

FROM reg.pk051.com/library/golang:v1.9.2

ARG GIT_BRANCH
ARG GIT_COMMIT_SHA
ARG GIT_COMMIT_LINK

ARG CI_BUILD_TIME
ARG CI_BUILD_NUMBER

ARG BUILD_ENV

ENV GIT_BRANCH ${GIT_BRANCH}
ENV GIT_COMMIT_SHA ${GIT_COMMIT_SHA}
ENV GIT_COMMIT_LINK ${GIT_COMMIT_LINK}

ENV CI_BUILD_TIME ${CI_BUILD_TIME}
ENV CI_BUILD_NUMBER ${CI_BUILD_NUMBER}

ENV BUILD_ENV ${BUILD_ENV}

LABEL git_branch=${GIT_BRANCH} git_commit_sha=${GIT_COMMIT_SHA} git_commit_link=${GIT_COMMIT_LINK}
LABEL ci_build_time=${CI_BUILD_TIME} ci_build_number=${CI_BUILD_NUMBER}
LABEL build_env=${BUILD_ENV}

RUN mkdir -p /go/bin/etc/fonts/
RUN mkdir -p /go/bin/template/
RUN mkdir -p /go/bin/log/server && chown -R golang.golang /go/bin/log/server
RUN touch /go/bin/log/server/sys.log && chown golang.golang /go/bin/log/server/sys.log

COPY release/server_linux /go/bin/server
ADD src/etc/fonts /go/bin/etc/fonts/


EXPOSE 9696 8989
HEALTHCHECK --interval=1m30s --timeout=5s --retries=3 CMD curl -f http://localhost:9696/version/test || exit 1
ENTRYPOINT ["/usr/sbin/gosu", "golang:golang", "/go/bin/server"]