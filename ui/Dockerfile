FROM node:19.5.0-alpine

WORKDIR /app

COPY password-manager-ui/public/ /app/public
COPY password-manager-ui/src/ /app/src
COPY password-manager-ui/package.json /app/

EXPOSE 3000

RUN npm install

CMD ["npm", "start"]