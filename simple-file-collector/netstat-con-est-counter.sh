#!/bin/bash

# Adjust as needed.
TEXTFILE_COLLECTOR_DIR=/var/lib/node_exporter/textfile_collector/
# Note the start time of the script.
START="$(date +%s)"
TCP_OPEN_CONNECTION=$(netstat -anp | grep :8428 | grep ESTABLISHED | wc -l)
sleep 10

# Write out metrics to a temporary file.
END="$(date +%s)"
cat << EOF > "$TEXTFILE_COLLECTOR_DIR/myscript.prom.$$"
myscript_duration_seconds $(($END - $START))
myscript_last_run_seconds $END
naren_netstat_Tcp_CurrEstab{port=8428} $TCP_OPEN_CONNECTION
EOF

# Rename the temporary file atomically.
# This avoids the node exporter seeing half a file.
mv "$TEXTFILE_COLLECTOR_DIR/myscript.prom.$$" \
  "$TEXTFILE_COLLECTOR_DIR/myscript.prom"
