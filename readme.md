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
sudo cp $(which humansd-manager) /usr/bin
sudo cp $(which processord) /usr/bin
```

## Initialize the validator, where "validator" is a moniker name
```
humansd init validator --chain-id testhuman

### Validator
### human17zc58s96rxj79jtqqsnzt3wtx3tern6areu43g
humansd keys add validator --keyring-backend test --recover

### Pool
### human14u53eghrurpeyx5cm47vm3qwugtmhcpnstfx9t
humansd keys add validator1 --keyring-backend test --recover

### Test 1
### human1dfjns5lk748pzrd79z4zp9k22mrchm2a5t2f6u
humansd keys add test1 --keyring-backend test --recover
```

## Add genesis accounts

```
humansd add-genesis-account $(humansd keys show validator -a --keyring-backend test) 90000000000000uheart
humansd add-genesis-account $(humansd keys show validator1 -a --keyring-backend test) 40000000000000uheart
humansd add-genesis-account $(humansd keys show test1 -a --keyring-backend test) 50000000000000uheart
```

## Generate CreateValidator signed transaction
```
humansd gentx validator 50000000000000uheart --keyring-backend test --chain-id testhuman
```

## Collect genesis transactions
```
humansd collect-gentxs
```

## replace stake to uheart

```
sed -i 's/stake/uheart/g' ~/.humans/config/genesis.json
```

## Create the service file "/etc/systemd/system/humansd.service" with the following content
```
sudo nano /etc/systemd/system/humansd.service

## Paste following content

[Unit]
Description=humansd

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

ExecStart=/usr/bin/humansd-manager start --pruning="nothing" --rpc.laddr "tcp://0.0.0.0:26657"

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

sudo systemctl enable humansd
sudo systemctl enable processord
sudo systemctl start humansd
sudo systemctl start processord

sudo systemctl status humansd
sudo systemctl status processord
```
