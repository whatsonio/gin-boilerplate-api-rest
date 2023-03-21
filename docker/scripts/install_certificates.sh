#!/bin/bash

apt-get update
apt-get install -y ca-certificates
update-ca-certificates
apt-get clean
