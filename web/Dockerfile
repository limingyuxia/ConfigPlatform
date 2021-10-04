FROM node:latest

ENV WORKPATH /home
WORKDIR $WORKPATH

COPY vue .

CMD cd $WORKPATH
CMD npm install
CMD npm run dev
