#!/bin/bash
set -eu

# ==================================================================================== #
# VARIABLES
# ==================================================================================== #

# Set the timezone for the server.
TIMEZONE=Europe/Stockholm

# Set the name of the new user to create.
USERNAME=sersophane

# Prompt to enter a password for the PostgreSQL sersophane user.
read -p "Enter password for sersophane DB user: " DB_PASSWORD

# Force all output to be presented in en_US for the duration of this script.
export LC_ALL=en_US.UTF-8


# ==================================================================================== #
# SCRIPT LOGIC
# ==================================================================================== #

# Enable the "universe" repository.
add-apt-repository --yes universe

# Update all software packages.
apt update
apt --yes -o Dpkg::Options::="--force-confnew" upgrade

# Set the system timezone and install all locales.
timedatectl set-timezone ${TIMEZONE}
apt --yes install locales-all

# Add the new user (and give them sudo privileges).
useradd --create-home --shell "/bin/bash" --groups sudo "${USERNAME}"

# Force a password to be set for the new user the first time they log in.
passwd --delete "${USERNAME}"
chage --lastday 0 "${USERNAME}"

# Copy the SSH keys from the root user to the new user.
rsync --archive --chown=${USERNAME}:${USERNAME} /root/.ssh /home/${USERNAME}

# Configure the firewall to allow SSH, HTTP and HTTPS traffic.
ufw allow 22
ufw allow 80/tcp
ufw allow 443/tcp
ufw --force enable

# Install fail2ban.
apt --yes install fail2ban

# Install the migrate CLI tool.
curl -L https://github.com/golang-migrate/migrate/releases/download/v4.14.1/migrate.linux-amd64.tar.gz | tar xvz
mv migrate.linux-amd64 /usr/local/bin/migrate

# Install PostgreSQL.
echo "deb http://apt.postgresql.org/pub/repos/apt $(lsb_release -cs)-pgdg main" > /etc/apt/sources.list.d/pgdg.list
wget --quiet -O - https://www.postgresql.org/media/keys/ACCC4CF8.asc | apt-key add -
apt-get update
apt --yes install postgresql-14

# Set up the sersophane DB and create a user account with the password entered earlier.
sudo -i -u postgres psql -c "CREATE DATABASE sersophane"
sudo -i -u postgres psql -d sersophane -c "CREATE EXTENSION IF NOT EXISTS citext"
sudo -i -u postgres psql -d sersophane -c "CREATE ROLE sersophane WITH LOGIN PASSWORD '${DB_PASSWORD}'"

# Add a DSN for connecting to the sersophane database to the system-wide environment variables in the /etc/environment file.
echo "SERSOPHANE_DB_DSN='postgres://sersophane:${DB_PASSWORD}@localhost/sersophane'" >> /etc/environment

# Install Caddy
apt --yes install -y debian-keyring debian-archive-keyring apt-transport-https
curl -L https://dl.cloudsmith.io/public/caddy/stable/gpg.key | sudo apt-key add -
curl -L https://dl.cloudsmith.io/public/caddy/stable/debian.deb.txt | sudo tee -a /etc/apt/sources.list.d/caddy-stable.list
apt update
apt --yes install caddy

echo "Script complete! Rebooting..."
reboot
