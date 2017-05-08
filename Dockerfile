FROM drhayt/ora

ADD . /go/src/github.com/divyag9/gooracle

RUN cd /go/src/github.com/divyag9/gooracle && \
	go install ./...
