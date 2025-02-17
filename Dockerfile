# SPDX-FileCopyrightText: 2024 SAP SE or an SAP affiliate company and Gardener contributors
#
# SPDX-License-Identifier: Apache-2.0

############# builder
FROM golang:1.22.0 AS builder

WORKDIR /go/src/github.com/gardener/gardener-extension-registry-cache
COPY . .

ARG EFFECTIVE_VERSION

RUN make install EFFECTIVE_VERSION=$EFFECTIVE_VERSION

############# base
FROM gcr.io/distroless/static-debian12:nonroot AS base
WORKDIR /

############# gardener-extension-registry-cache
FROM base AS registry-cache

COPY --from=builder /go/bin/gardener-extension-registry-cache /gardener-extension-registry-cache
ENTRYPOINT ["/gardener-extension-registry-cache"]

############# gardener-extension-registry-cache-admission
FROM base AS registry-cache-admission

COPY --from=builder /go/bin/gardener-extension-registry-cache-admission /gardener-extension-registry-cache-admission
ENTRYPOINT ["/gardener-extension-registry-cache-admission"]
