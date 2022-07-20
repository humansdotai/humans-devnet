# Install prerequisities
## Install Golang:

```
// Install latest 
go version https://golang.org/doc/install

wget -q -O - https://raw.githubusercontent.com/canha/golang-tools-install-script/master/goinstall.sh | bash -s -- --version 1.18

source ~/.profile

// to verify that Golang installed
go version

// Should return go version go1.18 linux/amd64
```

# Running a Validator node

## Install the executables
```
sudo rm -rf ~/.humans

make install

mkdir -p ~/.humans/upgrade_manager/upgrades

mkdir -p ~/.humans/upgrade_manager/genesis/bin
```

## Symlink genesis binary to upgrade

```
cp $(which humansd) ~/.humans/upgrade_manager/genesis/bin
sudo cp $(which pochumand-manager) /usr/bin
sudo cp $(which processord) /usr/bin
```

## Initialize the validator, where "validator" is a moniker name
```
humansd init validator --chain-id test

### Validator
### mun17zc58s96rxj79jtqqsnzt3wtx3tern6areu43g
echo "pet apart myth reflect stuff force attract taste caught fit exact ice slide sheriff state since unusual gaze practice course mesh magnet ozone purchase" | humansd keys add validator --keyring-backend test --recover

### Pool
### mun14u53eghrurpeyx5cm47vm3qwugtmhcpnstfx9t
echo "bottom soccer blue sniff use improve rough use amateur senior transfer quarter" | humansd keys add validator1 --keyring-backend test --recover

### Test 1
### mun1dfjns5lk748pzrd79z4zp9k22mrchm2a5t2f6u
echo "betray theory cargo way left cricket doll room donkey wire reunion fall left surprise hamster corn village happy bulb token artist twelve whisper expire" | humansd keys add test1 --keyring-backend test --recover
```

## Add genesis accounts

```
humansd add-genesis-account $(humansd keys show validator -a --keyring-backend test) 90000000000000uhmn
humansd add-genesis-account $(humansd keys show validator1 -a --keyring-backend test) 40000000000000uhmn
humansd add-genesis-account $(humansd keys show test1 -a --keyring-backend test) 50000000000000uhmn
```

## Generate CreateValidator signed transaction
```
humansd gentx validator 50000000000000uhmn --keyring-backend test --chain-id test
```

## Collect genesis transactions
```
humansd collect-gentxs
```

## replace stake to uhmn

```
sed -i 's/stake/uhmn/g' ~/.humans/config/genesis.json
```

## Create the service file "/etc/systemd/system/pochumand.service" with the following content
```
sudo nano /etc/systemd/system/pochumand.service

## Paste following content

[Unit]
Description=pochumand

Requires=network-online.target

After=network-online.target

[Service]
Restart=on-failure

RestartSec=3

User=venus

Group=venus

Environment=DAEMON_NAME=humansd

Environment=DAEMON_HOME=/home/venus/.humans

Environment=DAEMON_ALLOW_DOWNLOAD_BINARIES=on

Environment=DAEMON_RESTART_AFTER_UPGRADE=on

PermissionsStartOnly=true

ExecStart=/usr/bin/pochumand-manager start --pruning="nothing" --rpc.laddr "tcp://0.0.0.0:26657"

StandardOutput=file:/var/log/humansd/humansd.log

StandardError=file:/var/log/humansd/humansd_error.log

ExecReload=/bin/kill -HUP $MAINPID

KillSignal=SIGTERM

LimitNOFILE=4096

[Install]
WantedBy=multi-user.target
```


## Create log files for loand

```
make log-files

sudo systemctl enable pochumand
sudo systemctl enable processord
sudo systemctl start pochumand
sudo systemctl start processord

sudo systemctl status pochumand
sudo systemctl status processord
```
