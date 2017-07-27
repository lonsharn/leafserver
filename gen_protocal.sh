#########################################################################
# File Name: gen_protocal.sh
# Author: lonsharn
# mail: 275089785@qq.com
# Created Time: Thu 27 Jul 2017 05:45:37 PM CST
#########################################################################
#!/bin/bash

protoc --go_out=src/ protocal/*.proto
