FROM node:10.16.0-jessie-slim
RUN npm install -g @angular/cli
RUN mkdir -p /pinbox
WORKDIR /pinbox
COPY . /pinbox/
RUN npm install
RUN ng build --prod

FROM nginx:stable
RUN rm -rf /usr/share/nginx/html/*
COPY nginx.conf /etc/nginx/nginx.conf
COPY server.conf /etc/nginx/conf.d/default.conf
COPY --from=0 /pinbox/dist/pinbox/ /usr/share/nginx/html/
CMD ["nginx", "-g", "daemon off;"]