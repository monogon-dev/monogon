#!/bin/bash
exec "$1" -initrd-path "$2" -kernel-path "$3" -cmdline "$4"