sudo ufw deny 20
sudo ufw deny 21
sudo ufw deny 20/tcp
sudo ufw deny 21/tcp
sudo ufw deny 990/tcp
sudo ufw deny 40000:50000/tcp
echo "Stopping VSFTPD ..."
sudo systemctl stop vsftpd
sudo ufw status
sudo nmap localhost
