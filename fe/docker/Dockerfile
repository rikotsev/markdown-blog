FROM nginxinc/nginx-unprivileged:1.27-alpine

COPY build/ /usr/share/nginx/html

EXPOSE 8080

CMD ["nginx", "-g", "daemon off;"]