FROM alpine AS config
COPY pwgen .
RUN ./pwgen &>/dev/null

FROM gcr.io/distroless/static:nonroot
COPY --from=config /root/.config/pwgen-go .config/pwgen-go
WORKDIR /
COPY pwgen /
ENTRYPOINT ["/pwgen"]
