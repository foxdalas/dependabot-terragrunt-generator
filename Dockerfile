FROM       alpine:3.15
MAINTAINER Maxim Pogozhiy <foxdalas@gmail.com>

ARG TARGETARCH

COPY templates /app/
COPY dependabot-terragrunt-generator /app/dependabot-terragrunt-generator

ENTRYPOINT ["/app/dependabot-terragrunt-generator"]
