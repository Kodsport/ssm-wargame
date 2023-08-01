FROM node:20-alpine3.17 AS build

WORKDIR /app/
COPY package.json yarn.lock /app/
RUN yarn

COPY . /app/
COPY .env /app/.env
RUN yarn generate

FROM caddy:2.6.4-alpine

COPY --from=build /app/.output/public /var/www
COPY Caddyfile /etc/caddy/Caddyfile
