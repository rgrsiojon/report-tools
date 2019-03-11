# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang

# Copy the local package files to the container's workspace.
ADD . /go/src/github.com/rgrsiojon/report-tools

# Build the outyet command inside the container.
# (You may fetch or manage dependencies here,
# either manually or with a tool like "godep".)
# RUN go install github.com/rgrsiojon/report-tools
RUN go get github.com/adlio/trello github.com/gin-gonic/gin github.com/jinzhu/configor gopkg.in/mgo.v2 gopkg.in/mgo.v2/bson && go install github.com/rgrsiojon/report-tools 

# Run the outyet command by default when the container starts.
ENTRYPOINT /go/bin/report-tools

# Document that the service listens on port 8080.
EXPOSE 8080