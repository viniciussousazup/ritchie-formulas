FROM cimg/openjdk:8.0

USER root

RUN mkdir /app
COPY . /app
RUN chmod +x /app/set_umask.sh

WORKDIR /app
ENTRYPOINT ["/app/set_umask.sh"]
CMD ["/app/run.sh"]
