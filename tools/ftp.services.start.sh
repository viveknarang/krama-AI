sudo ufw allow 20
sudo ufw allow 21
sudo ufw allow 20/tcp
sudo ufw allow 21/tcp
sudo ufw allow 990/tcp
sudo ufw allow 40000:50000/tcp
echo "Starting VSFTP ...." 
sudo systemctl start vsftpd
sudo ufw status
sudo nmap localhost
