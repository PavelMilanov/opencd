
set -ue

version=0.2.5

wget https://github.com/PavelMilanov/opencd/releases/download/v`${version}`/opencd
chmod +x ./opencd
mv opencd /usr/bin/
sh -c "echo export PATH=$PATH:/usr/bin/opencd >> /etc/profile"

echo "opencd version `${version}` installed"
