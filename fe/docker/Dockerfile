FROM nginx:1.27-alpine-otel
WORKDIR /usr/share/nginx/html
COPY build/ .
EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]