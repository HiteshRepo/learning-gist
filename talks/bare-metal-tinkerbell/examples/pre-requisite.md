## Install vagrant and virtualbox - linux/Ubuntu

Virtualbox Installation#1
----------------------
install: sudo apt install virtualbox
remove: sudo apt remove virtualbox

Virtualbox Installation#2
----------------------
go to https://www.virtualbox.org/
click on download virtualbox
click on linux distributions
click on your version of linux link
know your ubuntu version [command] : lsb_release -a
cd Downloads/
sudo dpkg -i virtualbox-.............deb
remove: sudo apt remove virtualbox
sudo apt remove virtualbox-dkms

Vagrant Installation
-------------------------------
sudo apt install vagrant
sudo apt remove vagrant


Virtualbox uninstall correct steps
-----------------------------------------------------
1. dpkg --get-selections | grep -i virtualbox -> check if virtualbox exists
2. sudo apt-get --purge remove virtualbox
3. sudo apt-get --purge remove virtualbox-6.1

Check version: vboxmanage --version


Libvirt Installation
-----------------------------
check if virtualisation is supported in your system: kvm-ok
sudo apt install qemu-kvm libvirt-daemon-system
Add user: sudo adduser $USER libvirt

systemctl start libvirtd

systemctl status libvirtd
libvirtd.service - Virtualization daemon
Loaded: loaded (/usr/lib/systemd/system/libvirtd.service; enabled)
Active: active (running) since Mon 2014-05-12 08:49:40 EDT; 2s ago
[...]

systemctl stop libvirtd

systemctl status libvirtd
[...]
Active: inactive (dead) since Mon 2014-05-12 08:51:11 EDT; 4s ago



Co-existence of KVM and Virtualbox:
-----------------------------------------------------------
https://www.dedoimedo.com/computers/kvm-virtualbox.html
