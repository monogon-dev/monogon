#!/usr/bin/expect -f

set timeout 30

spawn core/scripts/launch.sh

expect "Network service got IP" {} default {
  send_error "Failed while waiting for IP address"
  exit 1
}

expect "Initialized encrypted storage" {
  exit 0
} default {
  send_error "Failed while waiting for encrypted storage"
  exit 1
}
