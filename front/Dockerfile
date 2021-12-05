FROM --platform=linux/x86_64 node:14.4.0-alpine

RUN mkdir /app

WORKDIR /app

ADD . /app

RUN apk update && apk add python make g++ gcc

RUN npm install

RUN npm run build

ENV HOST 0.0.0.0

EXPOSE 3000

CMD ["npm", "run", "start"]