FROM node:20-alpine3.17 AS build

WORKDIR /app/
COPY package.json yarn.lock /app/
RUN yarn

COPY . /app/
COPY .env /app/.env
ENV prod=true
RUN yarn build

ENTRYPOINT node .output/server/index.mjs

