FROM node:22-alpine AS build
WORKDIR /app
COPY package.json package-lock.json ./
RUN npm i
COPY . ./
RUN npm run build

FROM node:22-alpine AS release
WORKDIR /app
COPY --from=build /app/.output ./
ENV HOST=0.0.0.0
EXPOSE 80
CMD [ "node", "/app/server/index.mjs" ]

FROM node:22-alpine AS debug
WORKDIR /debug
COPY package.json package-lock.json ./
RUN npm i
COPY . ./
ENV HOST=0.0.0.0
CMD [ "npm", "run", "dev" ]
