FROM node:16.2.0-alpine as build
WORKDIR /app
COPY . /app
COPY package.json /app/package.json
ENV PATH /app/node_modules/.bin:$PATH

RUN ls 
RUN npm install --silent
# RUN npm install @vue/cli -g
RUN npm run build
RUN apk add nano
# production environment
FROM nginx:1.16.0-alpine
COPY --from=build /app/dist /usr/share/nginx/html
COPY nginx.conf /etc/nginx/conf.d/default.conf

EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]
