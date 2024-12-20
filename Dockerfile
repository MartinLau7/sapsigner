# syntax=docker.io/docker/dockerfile:1

FROM docker.io/library/golang:1.23-bookworm AS build
WORKDIR /usr/local/src/sapsigner
RUN apt-get update && \
    apt-get install --no-install-recommends -y curl make && \
    rm -Rf /var/lib/apt/lists/*
COPY . .
RUN make SHELL=/bin/bash BUILD_STATIC=1 sapsigner-emu.out

FROM gcr.io/distroless/static-debian12:nonroot AS run
COPY --from=build /usr/local/src/sapsigner/sapsigner-emu.out /usr/local/bin/sapsigner
ENTRYPOINT ["sapsigner"]
