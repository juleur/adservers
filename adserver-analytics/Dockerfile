FROM busybox:glibc

WORKDIR /home

COPY adserver-analytics .
COPY ./adserver-analytics/data/analytics.txt .

CMD ["./adserver-analytics"]