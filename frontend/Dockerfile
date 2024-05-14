FROM node:20-alpine
LABEL authors="Nifie"

WORKDIR /app
COPY package.json ./
RUN npm install
RUN npm i sharp

COPY . .

RUN npm run build

CMD ["npm", "start"]