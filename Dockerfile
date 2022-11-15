# syntax=docker/dockerfile:1
FROM node:12-alpine
RUN apk add --no-cache python2 g++ make
WORKDIR /furyhubchain
COPY . .
RUN yarn install-testing
CMD ["node", "src/index.js"]
EXPOSE 3000
