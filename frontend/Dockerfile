FROM node:20 AS builder

ADD ./frontend /app

WORKDIR /app

RUN rm -rf node_modules
RUN rm package-lock.json

RUN npm install

CMD ["npm", "run", "dev", "--", "--host"]