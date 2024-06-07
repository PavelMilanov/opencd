#!/usr/bin/bash

set -e

wget https://github.com/PavelMilanov/opencd/releases/download/v0.1.0/opencd
sudo mv opencd /usr/bin/
sudo sh -c "echo export PATH=$PATH:/usr/bin/opencd >> /etc/profile"
source /etc/profile
