FROM node:lts-alpine as build-stage

RUN mkdir /blog-web-app
WORKDIR /blog-web-app
COPY package.json /blog-web-app
RUN npm install -g @vue/cli
RUN npm install -g babel babel-runtime babel-register mocha nodemon
RUN npm install -g http-server
RUN npm install


RUN mkdir -p /usr/src/app
WORKDIR /usr/src/app
RUN ln -s /blog-web-app/node_modules node_modules

COPY . /usr/src/app

RUN npm run build

EXPOSE 8080
CMD [ "http-server", "dist" ]