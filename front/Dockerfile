FROM node:16-alpine3.11

ENV LANG ja_JP.UTF-8
ENV TZ Asia/Tokyo

RUN mkdir /front

WORKDIR /front

RUN apk add python make g++

COPY ../../front/package.json ./
COPY ../../front/yarn.lock ./

RUN yarn install