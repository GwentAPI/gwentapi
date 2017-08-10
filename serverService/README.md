## Experimental socket activated http server

To use systemd socket activation feature for zero downtime upgrade follow this guide:

1. copy the service definition unit file ``gwentapi.service``
and the socket definition unit file ``gwentapi.socket`` to ``/etc/systemd/system/``.

2. Copy the executable at the location defined by the service unit file: ``/usr/local/bin/`` by default.

3. Rename the executable to gwentapi to match the filename in the service unit file.

4. Enable and start the service
    1. ``systemctl enable gwentapi``
    2. ``systemctl start gwentapi``

This experimental feature is only supported on Linux with systemd.