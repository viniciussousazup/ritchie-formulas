FROM cimg/base:stable-20.04

USER root

RUN mkdir /app
COPY . /app
RUN chmod +x /app/set_umask.sh

WORKDIR /app
ENTRYPOINT ["/app/set_umask.sh"]
CMD ["/app/run.sh"]
