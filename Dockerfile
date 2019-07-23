FROM node:10.16.0-jessie-slim
RUN npm install -g @angular/cli
RUN mkdir -p /pinbox
WORKDIR /pinbox
COPY . /pinbox/
RUN npm install
RUN ng build --prod

FROM nginx:stable
COPY --from=0 /pinbox/dist /usr/share/nginx/html
CMD ["nginx", "-g", "daemon off;"]