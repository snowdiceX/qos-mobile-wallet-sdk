#!/bin/bash

rm -rf vendor
env GO111MODULE=on go mod vendor
env GO111MODULE=off gomobile bind -target android -o qosMobileWalletSDK.aar github.com/QOSGroup/qos-mobile-wallet-sdk


rm -rf vendor