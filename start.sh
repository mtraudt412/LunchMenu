#!/bin/bash
go get
go build
dev_appserver.py --port=8686 --admin_port=8787 app.yaml

