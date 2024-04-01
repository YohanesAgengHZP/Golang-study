#!/bin/bash

echo "Logged into cqlsh - executing schema.cql"
cqlsh -u test -p testcassandra -e "COPY smr.number_routing_table(number,number_status,target_a2p,target_p2p) from '/data/number_routing_table.csv' WITH HEADER = TRUE;"
