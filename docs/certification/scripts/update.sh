#!/bin/bash

source $HOME/.profile

go get -u -d github.com/tyaps/hydra
go get -d -u github.com/devopsfaith/krakend-examples/gin
(cd $HOME/hydra-login-consent-node; git pull -ff; npm i)
cd $HOME
go install github.com/tyaps/hydra
go install github.com/devopsfaith/krakend-examples/gin