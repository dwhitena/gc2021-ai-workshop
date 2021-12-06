FROM alpine

RUN apk update && \
    apk add ca-certificates wget unzip && \
    rm -rf /root/.[acpw]* ipaexg00301*

ADD gcqa .

CMD ["./gcqa"]
