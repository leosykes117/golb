# etapa de compilación
FROM node:14.17.0-alpine3.12 as build-stage
WORKDIR /app
COPY blog-cms/package*.json ./
RUN npm install
COPY blog-cms/ ./
RUN npm run build

# etapa de producción
FROM nginx:1.19.10-alpine as production-stage
RUN rm /etc/nginx/conf.d/default.conf
COPY nginx/nginx.conf /etc/nginx/conf.d
COPY --from=build-stage /app/dist /usr/share/nginx/html
EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]
