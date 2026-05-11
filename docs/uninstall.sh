#!/bin/bash

# check for root privileges
if [ "$EUID" -ne 0 ]; then
  echo "Please run as root to uninstall tfimport"
  exit 1
fi

if ! rm -f /usr/local/bin/tfimport; then
  echo "Failed to remove tfimport from /usr/local/bin"
  exit 1
fi

echo "tfimport uninstalled successfully. Thank you for using it!"
