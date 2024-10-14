
set -ue

wget https://github.com/PavelMilanov/opencd/releases/download/v0.2.4/opencd
chmod +x ./opencd
mv opencd /usr/bin/
sh -c "echo export PATH=$PATH:/usr/bin/opencd >> /etc/profile"

echo "opencd installed"
