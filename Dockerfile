FROM       alpine:3.15
MAINTAINER Maxim Pogozhiy <foxdalas@gmail.com>

ARG TARGETARCH

COPY dependabot-terragrunt-generator /app/dependabot-terragrunt-generator
COPY templates /app/

ENTRYPOINT ["/app/dependabot-terragrunt-generator"]
