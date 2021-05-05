#!/bin/bash

function create_slave_replication_user_on_master() {
    MASTER_HOST_IP=$1
    MASTER_ROOT_USER=$2
    MASTER_ROOT_PASSWORD=$3

    SLAVE_HOST_IP=$4
    REPL_USER=$5
    REPL_PASSWORD=$6

    MYSQL_QUERY="CREATE USER '$REPL_USER'@'$SLAVE_HOST_IP' IDENTIFIED BY '$REPL_PASSWORD';"
    mariadb --user="$MASTER_ROOT_USER" --password="$MASTER_ROOT_PASSWORD" --host="$MASTER_HOST_IP" -e "$MYSQL_QUERY"
}


function grant_replication_slave_permission_on_master() {
    MASTER_HOST_IP=$1
    MASTER_ROOT_USER=$2
    MASTER_ROOT_PASSWORD=$3
    REPL_USER=$4
    SLAVE_HOST_IP=$5

    MYSQL_QUERY="GRANT REPLICATION SLAVE ON *.* TO '$REPL_USER'@'$SLAVE_HOST_IP';"
    mariadb --user="$MASTER_ROOT_USER" --password="$MASTER_ROOT_PASSWORD" --host="$MASTER_HOST_IP" -e "$MYSQL_QUERY"
}

# Reload privileges for replication to take effect
function flush_privileges_for_replication_to_take_effect() {
    MASTER_HOST_IP=$1
    MASTER_ROOT_USER=$2
    MASTER_ROOT_PASSWORD=$3

    MYSQL_QUERY="FLUSH PRIVILEGES;"
    mariadb --user="$MASTER_ROOT_USER" --password="$MASTER_ROOT_PASSWORD" --host="$MASTER_HOST_IP" -e "$MYSQL_QUERY"
}

function get_master_replication_file_and_position_and_update_slave () {
    MASTER_HOST_IP=$1
    MASTER_ROOT_USER=$2
    MASTER_ROOT_PASSWORD=$3

    SLAVE_HOST_IP=$4
    SLAVE_ROOT_USER=$5
    SLAVE_ROOT_PASSWORD=$6

    REPL_USER=$7
    REPL_PASSWORD=$8

    REPL_LOG_FILE=$(eval "mariadb --user=$MASTER_ROOT_USER --password=$MASTER_ROOT_PASSWORD --host=$MASTER_HOST_IP -e 'SHOW MASTER STATUS\G' | grep File | sed -n -e 's/^.*: //p'")
    REPL_LOG_POSITION=$(eval "mariadb --user=$MASTER_ROOT_USER --password=$MASTER_ROOT_PASSWORD --host=$MASTER_HOST_IP -e 'SHOW MASTER STATUS\G' | grep Position | sed -n -e 's/^.*: //p'")

    MYSQL_QUERY="CHANGE MASTER TO MASTER_HOST = '$MASTER_HOST_IP', MASTER_USER = '$REPL_USER', MASTER_PASSWORD = '$REPL_PASSWORD', MASTER_LOG_FILE = '$REPL_LOG_FILE', MASTER_LOG_POS = $REPL_LOG_POSITION;"
    mariadb --user="$SLAVE_ROOT_USER" --password="$SLAVE_ROOT_PASSWORD" --host="$SLAVE_HOST_IP" -e "$MYSQL_QUERY"
}

function start_slave () {
    SLAVE_HOST_IP=$1
    SLAVE_ROOT_USER=$2
    SLAVE_ROOT_PASSWORD=$3

    MYSQL_QUERY="START SLAVE;"
    mariadb --user="$SLAVE_ROOT_USER" --password="$SLAVE_ROOT_PASSWORD" --host="$SLAVE_HOST_IP" -e "$MYSQL_QUERY"
}

function check_slave_status() {
    SLAVE_HOST_IP=$1
    SLAVE_ROOT_USER=$2
    SLAVE_ROOT_PASSWORD=$3

    MYSQL_QUERY="SHOW SLAVE STATUS \G;"
    mariadb --user="$SLAVE_ROOT_USER" --password="$SLAVE_ROOT_PASSWORD" --host="$SLAVE_HOST_IP" -e "$MYSQL_QUERY"
}


function setup_master_2_slave_replication() {
    MASTER_HOST_IP=$1
    MASTER_ROOT_USER=$2
    MASTER_ROOT_PASSWORD=$3

    SLAVE_HOST_IP=$4
    SLAVE_ROOT_USER=$5
    SLAVE_ROOT_PASSWORD=$6

    REPL_USER=$7
    REPL_PASSWORD=$8

    sleep 1
    create_slave_replication_user_on_master $MASTER_HOST_IP $MASTER_ROOT_USER $MASTER_ROOT_PASSWORD $SLAVE_HOST_IP $REPL_USER $REPL_PASSWORD

    sleep 1
    grant_replication_slave_permission_on_master $MASTER_HOST_IP $MASTER_ROOT_USER $MASTER_ROOT_PASSWORD $REPL_USER $SLAVE_HOST_IP

    sleep 1
    flush_privileges_for_replication_to_take_effect $MASTER_HOST_IP $MASTER_ROOT_USER $MASTER_ROOT_PASSWORD

    sleep 1
    get_master_replication_file_and_position_and_update_slave $MASTER_HOST_IP $MASTER_ROOT_USER $MASTER_ROOT_PASSWORD $SLAVE_HOST_IP $SLAVE_ROOT_USER $SLAVE_ROOT_PASSWORD $REPL_USER $REPL_PASSWORD

    sleep 1
    start_slave $SLAVE_HOST_IP $SLAVE_ROOT_USER $SLAVE_ROOT_PASSWORD

    sleep 1
    check_slave_status $SLAVE_HOST_IP $SLAVE_ROOT_USER $SLAVE_ROOT_PASSWORD
}

function create_tables() {
    USER=$1
    PASSWORD=$2
    HOST=$3
    DB=$4

    mariadb --user="$USER" --password="$PASSWORD" --host="$HOST" $DB < /opt/repl/scripts/create_tables.sql
}

apt-get -qq update
apt-get -qq install mariadb-client -y --no-install-recommends


until mariadb --user="root" --password="$PRIMARY_1_ROOT_PASSWORD" --host="$PRIMARY_1_HOST_IP" -e "SHOW TABLES;" golb; do
  >&2 echo "Mariadb 1 is unavailable - sleeping"
  sleep 1
done

until mariadb --user="root" --password="$PRIMARY_2_ROOT_PASSWORD" --host="$PRIMARY_2_HOST_IP" -e "SHOW TABLES;" golb; do
  >&2 echo "Mariadb 2 is unavailable - sleeping"
  sleep 1
done

#######################################################################################################################
# Setup replication from primary1 to primary2
#######################################################################################################################
setup_master_2_slave_replication $PRIMARY_1_HOST_IP $PRIMARY_1_ROOT_USER $PRIMARY_1_ROOT_PASSWORD $PRIMARY_2_HOST_IP $PRIMARY_2_ROOT_USER $PRIMARY_2_ROOT_PASSWORD $REPL_USER $REPL_PASSWORD

# Add some gap intentionally
# sleep 5

#######################################################################################################################
# Setup replication from primary2 to primary1
#######################################################################################################################
setup_master_2_slave_replication $PRIMARY_2_HOST_IP $PRIMARY_2_ROOT_USER $PRIMARY_2_ROOT_PASSWORD $PRIMARY_1_HOST_IP $PRIMARY_1_ROOT_USER $PRIMARY_1_ROOT_PASSWORD $REPL_USER $REPL_PASSWORD

create_tables golb_dev 100PerApi $PRIMARY_1_HOST_IP golb
