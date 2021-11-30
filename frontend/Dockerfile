# init stage
FROM node:14.17.0-alpine3.12 as init-stage
WORKDIR /usr/src/app/blog-cms
COPY blog-cms/package*.json ./
RUN npm install

# develop stage
FROM init-stage as develop-stage
EXPOSE 8080
ENV NODE_ENV=development
CMD ["npm", "run", "serve"]

# build stage
FROM init-stage as build-stage
ENV NODE_ENV=production
COPY blog-cms/ ./
RUN npm run build

# production stage
FROM nginx:1.19.10-alpine as production-stage
RUN rm /etc/nginx/conf.d/default.conf
COPY nginx/nginx.conf /etc/nginx/conf.d
COPY --from=build-stage /usr/src/app/blog-cms/dist /usr/share/nginx/html
EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]
