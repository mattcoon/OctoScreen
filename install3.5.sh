cd ~
git clone https://github.com/goodtft/LCD-show.git
chmod -R 755 LCD-show
cd LCD-show/
sudo ./LCD35-show
sudo sed -i 's/^hdmi_cvt.*/#\0\nhdmi_cvt=800 533 60 6 0 0 0/' /boot/config.txt  # simulate compatiblity with OctoScreen resolution

sudo apt-get install libgtk-3-0 xserver-xorg xinit x11-xserver-utils

wget https://github.com/Z-Bolt/OctoScreen/releases/download/v2.7.4/octoscreen_2.7.4_armhf.deb
sudo dpkg -i octoscreen_2.7.4_armhf.deb  # this will say there are errors about missing dependencies, but otherwise work fine
sudo apt-get install --fix-broken        # fix those missing dependencies
sudo apt-get install x11-xserver-utils   # and some more missing dependencies
sudo sed -i 's/OCTOSCREEN_RESOLUTION=.*/OCTOSCREEN_RESOLUTION=800x533/' /etc/octoscreen/config
sudo systemctl set-default graphical     # otherwise octoscreen does not start by default
sudo apt-get install update
sudo apt-get install upgrade

sudo reboot now


curl -fsSL get.docker.com -o get-docker.sh
sh get-docker.sh
git clone https://github.com/mattcoon/OctoScreen
cd OctoScreen
sudo make build
ls -1 build/
sudo dpkg -i octoscreen_2.7.4-1_armhf.deb  # this will say there are errors about missing dependencies, but otherwise work fine
