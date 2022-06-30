# Install Golang:

// Install latest go version https://golang.org/doc/install
wget -q -O - https://raw.githubusercontent.com/canha/golang-tools-install-script/master/goinstall.sh | bash -s -- --version 1.18
source ~/.profile

// to verify that Golang installed
go version
// Should return go version go1.18 linux/amd64

#Running a Validator node

Install the executables

sudo rm -rf ~/.poc-human
make install

clear

mkdir -p ~/.poc-human/upgrade_manager/upgrades

mkdir -p ~/.poc-human/upgrade_manager/genesis/bin

# symlink genesis binary to upgrade
cp $(which poc-humand) ~/.poc-human/upgrade_manager/genesis/bin

sudo cp $(which pochumand-manager) /usr/bin

sudo cp $(which processord) /usr/bin

# Initialize the validator, where "validator" is a moniker name
poc-humand init validator --chain-id test
 
# Validator
# mun17zc58s96rxj79jtqqsnzt3wtx3tern6areu43g
echo "pet apart myth reflect stuff force attract taste caught fit exact ice slide sheriff state since unusual gaze practice course mesh magnet ozone purchase" | poc-humand keys add validator --keyring-backend test --recover

# Validator1
# mun14u53eghrurpeyx5cm47vm3qwugtmhcpnstfx9t
echo "bottom soccer blue sniff use improve rough use amateur senior transfer quarter" | poc-humand keys add validator1 --keyring-backend test --recover

# Test 1
# mun1dfjns5lk748pzrd79z4zp9k22mrchm2a5t2f6u
echo "betray theory cargo way left cricket doll room donkey wire reunion fall left surprise hamster corn village happy bulb token artist twelve whisper expire" | poc-humand keys add test1 --keyring-backend test --recover

# Add genesis accounts
poc-humand add-genesis-account $(poc-humand keys show validator -a --keyring-backend test) 90000000000000uhmn

poc-humand add-genesis-account $(poc-humand keys show validator1 -a --keyring-backend test) 40000000000000uhmn

poc-humand add-genesis-account $(poc-humand keys show test1 -a --keyring-backend test) 50000000000000uhmn

# Generate CreateValidator signed transaction
poc-humand gentx validator 50000000000000uhmn --keyring-backend test --chain-id test

# Collect genesis transactions
poc-humand collect-gentxs

# replace stake to TMUN
sed -i 's/stake/uhmn/g' ~/.poc-human/config/genesis.json


# Create the service file "/etc/systemd/system/pochumand.service" with the following content
sudo nano /etc/systemd/system/pochumand.service
# paste following content
[Unit]
Description=pochumand

Requires=network-online.target

After=network-online.target

[Service]
Restart=on-failure

RestartSec=3

User=venus

Group=venus

Environment=DAEMON_NAME=poc-humand

Environment=DAEMON_HOME=/home/venus/.poc-human

Environment=DAEMON_ALLOW_DOWNLOAD_BINARIES=on

Environment=DAEMON_RESTART_AFTER_UPGRADE=on

PermissionsStartOnly=true

ExecStart=/usr/bin/pochumand-manager start --pruning="nothing" --rpc.laddr "tcp://0.0.0.0:26657"

StandardOutput=file:/var/log/poc-humand/poc-humand.log

StandardError=file:/var/log/poc-humand/poc-humand_error.log

ExecReload=/bin/kill -HUP $MAINPID

KillSignal=SIGTERM

LimitNOFILE=4096

[Install]
WantedBy=multi-user.target

# Create the service file "/etc/systemd/system/processord.service" with the following content
sudo nano /etc/systemd/system/processord.service

# paste following content
[Unit]
Description=processord

Requires=network-online.target

After=network-online.target

[Service]
Restart=on-failure

RestartSec=3

User=venus

Group=venus

Environment=DAEMON_NAME=processord

Environment=DAEMON_HOME=/home/venus/.poc-human

PermissionsStartOnly=true

ExecStart=/usr/bin/processord start validator

StandardOutput=file:/var/log/processord/processord.log

StandardError=file:/var/log/processord/processord_error.log

ExecReload=/bin/kill -HUP $MAINPID

KillSignal=SIGTERM

LimitNOFILE=4096

[Install]
WantedBy=multi-user.target


# Create log files for loand
make log-files

sudo systemctl enable pochumand

sudo systemctl enable processord

sudo systemctl start pochumand

sudo systemctl start processord

sudo systemctl status pochumand
sudo systemctl status processord
