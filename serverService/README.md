## Experimental socket activated http server

To use systemd socket activation feature for zero downtime upgrade follow this guide:

1. copy the service definition unit file ``gwentapi.service``
and the socket definition unit file ``gwentapi.socket`` to ``/etc/systemd/system/``.

2. Copy the executable and the config file at the location defined by the service unit file: ``/usr/local/bin/`` by default.

3. Rename the executable to gwentapi to match the filename in the service unit file.

4. Enable the socket and start it
    1. ``systemctl enable gwentapi.socket``
    2. ``systemctl start gwentapi.socket``

This experimental feature is only supported on Linux with systemd.

### To upgrade the service

1. Replace the executable
2. Restart the service with systemctl