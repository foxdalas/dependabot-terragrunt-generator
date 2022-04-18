FROM       alpine:3.15
MAINTAINER Maxim Pogozhiy <foxdalas@gmail.com>

ARG TARGETARCH
WORKDIR /app
COPY templates/config.tmpl /app/templates/config.tmpl
COPY dependabot-terragrunt-generator /app/dependabot-terragrunt-generator

ENTRYPOINT ["/app/dependabot-terragrunt-generator"]
