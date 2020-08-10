# Sherlock
Sherlock is a tool for incident response teams to quickly gather log files and find system resources and binaries that may have been tampered with by attackers. This tool can also be used to establish baselines before an attack.

## Information
- Blue teaming and incident response can be stressful, hopefully this tool makes it a little easier!
- Has only been tested on Debian based systems so far.

## Usage
```bash
make # uses makefile and builds in /src/sherlock
sudo src/sherlock help
```

## Services supported
```
Nginx
Apache2
OpenSSH-Server
MySQL Server
```

## Dev environment 
```bash
git clone https://github.com/Akshay-Rohatgi/Sherlock
cd Sherlock && bash misc/install.sh
```