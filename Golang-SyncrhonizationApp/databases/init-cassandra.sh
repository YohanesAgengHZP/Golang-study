#!/bin/bash

echo "Waiting for Cassandra to start..."

# # Wait until Cassandra is ready to accept connections
# until cqlsh -e "DESCRIBE KEYSPACES" 2> /dev/null; do
#     >&2 echo "Cassandra is unavailable - sleeping"
#     sleep 2
# done

# echo "Cassandra is up - logging into cqlsh"

# Log into cqlsh (if authentication is required)
# Replace 'username' and 'password' with actual credentials
cqlsh -u test -p testcassandra -f schema.cql

echo "Logged into cqlsh - executing schema.cql"

echo "Schema initialization complete"
