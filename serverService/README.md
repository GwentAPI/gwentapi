## socket activated http server

To use systemd socket activation feature for zero downtime upgrade, copy the service definition unit file ``gwentapi.service``
and the socket definition unit file ``gwentapi.socket`` to ``/etc/systemd/system/``.

Then, enable and start the service with systemd.

This is only supported on Linux with systemd.