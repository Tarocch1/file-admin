FROM --platform=$BUILDPLATFORM alpine:latest as builder

WORKDIR /app
COPY dist .
COPY .github/actions/docker/file.sh .

ARG TARGETPLATFORM

RUN ash ./file.sh

FROM --platform=$TARGETPLATFORM alpine:latest

WORKDIR /app

COPY --from=builder /app/file-admin .

RUN set -ex && \
    apk --no-cache add ca-certificates && \
    chmod +x /app/file-admin && \
    /app/file-admin -v

ENV PATH /app:$PATH

CMD ["file-admin"]
