FROM cimg/base:stable-20.04

USER root

RUN mkdir /app
COPY . /app
RUN chmod +x /app/set_umask.sh

RUN  echo "[user]" > /root/.gitconfig
RUN  echo "	name = Your Name" >> /root/.gitconfig
RUN  echo "	email = you@example.com" >> /root/.gitconfig
RUN  echo "[url \"git@github.com:\"]" >> /root/.gitconfig
RUN  echo "  insteadOf = https://github.com/" >> /root/.gitconfig

WORKDIR /app
ENTRYPOINT ["/app/set_umask.sh"]
CMD ["/app/run.sh"]
