#!/bin/bash
snmpset -v 2c -c YOU_COMMUNITY  192.168.0.100 ifAdminStatus.49154 i 2
snmpset -v 2c -c YOU_COMMUNITY  192.168.0.100 ifAdminStatus.49155 i 2
snmpset -v 2c -c YOU_COMMUNITY  192.168.0.101 ifAdminStatus.49154 i 2
snmpset -v 2c -c YOU_COMMUNITY  192.168.0.101 ifAdminStatus.49155 i 2
