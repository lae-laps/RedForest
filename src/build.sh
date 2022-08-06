#!/bin/bash
gcc main.c modules/dns/query.c interface/interface.c interface/ascii.c -o main
