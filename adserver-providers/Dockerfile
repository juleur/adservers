FROM busybox:glibc

WORKDIR /home

COPY adserver-providers .
COPY ./adserver-providers/data/campaigns.json .

CMD ["./adserver-providers"]

