FROM golang:latest AS build
WORKDIR /build
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM alpine:latest  
ENV PageTitle="Demo test page"
ENV ClusterName="localcluster"
ENV ListenPort="8080"
WORKDIR /root/
COPY --from=build /build/src/ /root/src/
COPY --from=build /build/app /root/
CMD ["./app"] 