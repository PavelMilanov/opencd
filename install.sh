#!/usr/bin/bash

set -e

wget https://github.com/PavelMilanov/opencd/releases/download/v0.1.1/opencd
chmod +x ./opencd
mv opencd /usr/bin/
sh -c "echo export PATH=$PATH:/usr/bin/opencd >> /etc/profile"
source /etc/profile