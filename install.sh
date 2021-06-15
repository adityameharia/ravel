#!/usr/bin/bash

RED="\e[31m"
YELLOW="\e[33m"
BLUE="\e[34m"
ENDCOLOR="\e[0m"

echo -e "${RED}"
cat << "EOF"

                    | |
 _ __ __ ___   _____| |
| '__/ _` \ \ / / _ \ |
| | | (_| |\ V /  __/ |
|_|  \__,_| \_/ \___|_|
             
EOF
echo -n -e "${ENDCOLOR}"
echo "A fault-tolerant, sharded key-value store"
echo "-----------------------------------------"
echo "Downloading ravel_node and ravel_cluster_admin from github: "
echo ""

ravel_node_url=https://github.com/adityameharia/ravel/releases/download/0.1/ravel_node
ravel_cluster_admin_url=https://github.com/adityameharia/ravel/releases/download/0.1/ravel_cluster_admin

curl -LJO $ravel_node_url  && sudo mv ./ravel_node /usr/local/bin
curl -LJO $ravel_cluster_admin_url && sudo mv ./ravel_cluster_admin /usr/local/bin
chmod +x /usr/local/bin/ravel_node
chmod +x /usr/local/bin/ravel_cluster_admin

echo ""
echo -e "${YELLOW}ravel_node and ravel_cluster_admin were downloaded and moved to /usr/local/bin${ENDCOLOR}"
echo ""
echo "You can now run the following commands:"
echo -e "${BLUE}ravel_node --help${ENDCOLOR}"
echo -e "${BLUE}ravel_cluster_admin --help${ENDCOLOR}"