FROM node:20-alpine3.17 AS build

WORKDIR /app/
COPY package.json yarn.lock /app/
RUN yarn

COPY . /app/
COPY .env /app/.env
RUN yarn build
ENV prod=true

ENTRYPOINT node .output/server/index.mjs

