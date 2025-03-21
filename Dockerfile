FROM debian:stable-slim

RUN apt-get update \
    && DEBIAN_FRONTEND=noninteractive apt-get install -y ca-certificates dumb-init \
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/*

ADD go-scaffold /
ADD config.yaml /
ADD resources /resources
ADD launcher /
ADD exec /exec
# For GCP Profiler
ENV GOOGLE_APPLICATION_CREDENTIALS=/gcp_profiler_credential.json
ENV GOGC=100

EXPOSE 8080

ENTRYPOINT ["dumb-init", "--"]
CMD ["/launcher", "-c", "/config.yaml"]
