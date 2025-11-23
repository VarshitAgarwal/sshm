# sshm — Smart SSH Manager for Linux

`sshm` is a modern SSH productivity tool that helps developers, DevOps engineers,
and SREs manage multiple servers efficiently.

It automatically reads your `~/.ssh/config`, provides intelligent tab-completion,
and lets you connect or run commands remotely **without entering an interactive SSH session**.

---

## ✨ Features

- Auto-detect SSH hosts from `~/.ssh/config`
- Connect instantly using alias-based SSH profiles  
- Run commands remotely **without full SSH login**
- Interactive remote execution shell (`sshm shell`)
- Host auto-completion (Bash/Zsh/Fish)
- Inspired by `kubectx` + `kubectl exec`
- Fast, clean, developer-friendly UX

---

## 📦 Installation (Linux)

### **1. Install Go (if not installed)**

```bash
sudo apt install golang-go -y          # Ubuntu / Debian
sudo dnf install golang -y             # Fedora
sudo pacman -S go     
```
               

## Verify GO Version 
```bash
go version
```
## Clone The Git Repo

```bash
git clone https://github.com/VarshitAgarwal/sshm.git
cd sshm
```
##  Build the SSHM Package
```bash
go build -o sshm .
```

 ## Move to bin folder 
```bash
sudo mv sshm /usr/local/bin/
```

## Complete Auto Completion Setup

```bash
mkdir -p ~/.zsh/completions
sshm completion zsh > ~/.zsh/completions/_sshm
echo 'fpath=(~/.zsh/completions $fpath)' >> ~/.zshrc
source ~/.zshrc
```

## Example

```sh
sshm connect <TAB>


```sh
Host prod-db
    HostName 10.10.1.12
    User ubuntu
    IdentityFile ~/.ssh/id_rsa

Host staging
    HostName 10.10.2.5
    User ec2-user
    IdentityFile ~/.ssh/staging.pem

