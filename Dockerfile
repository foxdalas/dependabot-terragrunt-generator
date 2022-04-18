FROM       alpine:3.15
MAINTAINER Maxim Pogozhiy <foxdalas@gmail.com>

ARG TARGETARCH

COPY dependabot-terragrunt-generator /bin/dependabot-terragrunt-generator

ENTRYPOINT ["/bin/dependabot-terragrunt-generator"]
