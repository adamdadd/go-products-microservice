FROM node as builder
WORKDIR /app
ADD . /app
RUN npm install && npm build

FROM nginx:alpine
WORKDIR /app
COPY --from=builder /app .
CMD ["nginx", "-g", "deamon off;"]