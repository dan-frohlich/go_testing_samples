#!/bin/sh
~/hacking/go_workspace/bin/mockgen \
    -source=interfaces.go \
    -destination=mock_interfaces.go \
    -package=service