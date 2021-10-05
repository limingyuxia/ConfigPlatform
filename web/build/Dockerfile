FROM node:latest

ENV WORKPATH /home
WORKDIR $WORKPATH

COPY web/vue .

RUN cd $WORKPATH
RUN npm install
RUN npm run build:prod
