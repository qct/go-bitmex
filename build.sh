#!/bin/bash

PROJECT_HOME=$(cd $(dirname "$0") && pwd)
GO_SWAGGER_URL="https://github.com/go-swagger/go-swagger/releases/download"
OS="linux"
ARCH="amd64"

get_latest_release() {
    curl --silent "https://api.github.com/repos/$1/releases/latest" | # Get latest release from GitHub api
      grep '"tag_name":' |                                            # Get tag line
      sed -E 's/.*"([^"]+)".*/\1/'                                    # Pluck JSON value
}

### Download go-swagger
if [[ ! -f $PROJECT_HOME/go-swagger ]]; then
  LATAEST_TAG=`get_latest_release "go-swagger/go-swagger"`
  GO_SWAGGER_DOWNLOAD_URL="$GO_SWAGGER_URL/$LATAEST_TAG/swagger_${OS}_$ARCH"
  echo "Downlaoding go-swagger binary into $PROJECT_HOME"
  wget $GO_SWAGGER_DOWNLOAD_URL -O $PROJECT_HOME/go-swagger
fi
chmod +x $PROJECT_HOME/go-swagger

### Download swagger.json
RESOURCES="https://www.bitmex.com/api/explorer/swagger.json"
wget $RESOURCES -O $PROJECT_HOME/swagger.json
cat $PROJECT_HOME/swagger.json|jq . > $PROJECT_HOME/swagger.json
sed -i 's/.*"default": {},/"default": null,/' $PROJECT_HOME/swagger.json

### Validate & format swagger.json
echo "Validating swagger.json ..."
$PROJECT_HOME/go-swagger validate $PROJECT_HOME/swagger.json
if [[ $? -ne 0 ]]; then echo "Failed validating swagger.json"; exit 1; fi;
echo "Validated successfully"

### Generate go client
$PROJECT_HOME/go-swagger generate client -f $PROJECT_HOME/swagger.json -A go-bitmex
