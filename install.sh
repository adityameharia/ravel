#!/usr/bin/bash

ravel_node_url=https://github.com/adityameharia/ravel/releases/download/0.1/ravel_node
ravel_cluster_admin_url=https://github.com/adityameharia/ravel/releases/download/0.1/ravel_cluster_admin

curl -LJO $ravel_node_url  && sudo mv ./ravel_node /usr/local/bin
curl -LJO $ravel_cluster_admin_url && sudo mv ./ravel_cluster_admin /usr/local/bin

echo "ravel_node and ravel_cluster_admin were downloaded and moved to /usr/local/bin"