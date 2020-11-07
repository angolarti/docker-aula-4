FROM node:14.11.0-alpine3.11
WORKDIR /avalista
COPY package*.json ./
RUN npm install
COPY . .
EXPOSE 9094
CMD [ "node", "avalista.js" ]