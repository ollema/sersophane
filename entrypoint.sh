#!/bin/bash

# Start the first process
/pb/pocketbase serve &

# Start the second process
node build &

# Wait for any process to exit
wait -n

# Exit with status of process that exited first
exit $?