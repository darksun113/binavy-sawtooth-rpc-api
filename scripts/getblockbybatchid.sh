#!/bin/zsh

curl -s -X POST http://localhost:16899/common/getblockbybatchid \
-H "content-type: application/json" -d "$(cat block.json)" \
> block-rsp.json

code=$(jq '.code' block-rsp.json)
if [ $code != 0 ]
then
	echo "Failed"
else
	echo "Successfully. See block-rsp.json"
fi