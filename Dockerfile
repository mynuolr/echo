FROM alpine
EXPOSE 1080/tcp
EXPOSE 1080/udp
ENTRYPOINT ["/echo"]
COPY echo /
