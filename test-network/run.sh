#!/bin/bash

rm -rf ../population/application-go/keystore
rm -rf ../population/application-go/wallet

/bin/bash ./network.sh down 

/bin/bash ./network.sh up

/bin/bash ./network.sh createChannel -c population

/bin/bash ./network.sh deployCC -c population -ccn basic -ccp ../population/chaincode-go -ccl go
