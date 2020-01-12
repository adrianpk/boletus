# Supervisord and Gulp

Supervisord and Gulp are not not required dependencies but they help to automate some tasks during development.

Any change in project source files triggers project compilation and server restart.

You can avoid all these steps if you want to perform these tasks manually (`make build` and `make run` after changing something.)

*Install*

[Supervisord official guide](http://supervisord.org/installing.html)

[Gulp official guide](https://gulpjs.com/docs/en/getting-started/quick-start)

```shell
$ sudo vim /etc/supervisor/conf.d
```

```shell
[supervisord]
logfile = /tmp/supervisord.log

[program:granica]
user=your-username
command=/path/to/foundation/backend/foundation/bin/granica
autostart=true
autorestart=true
redirect_stderr=true
```

To launch the server

```shell
$ cd /path/to/foundation/backend/foundation/bin/granica
$ gulp
```


If you have problems launching supervisord or running gulpfile tasks with a
a user other than root this supervisord config excerpt from the top of the file can be helpful as a reference.

```shell
$ sudo vim /etc/supervisor/supervisord.conf
```

```shell
; supervisor config file

[unix_http_server]
file=/var/run/supervisor.sock   ; (the path to the socket file)
chmod=0770                      ; sockef file mode (default 0700)
chown=root:supervisor

; Rest of the file omitted for clarity.
```

Restart it
```shell
$ sudo service supervisor restart
```

You will also need to append your user to supervisor group and restart your system or do a logout-login sequence after executing these commands to get the new rights.

```shell
$ sudo groupadd supervisor
$ sudo usermod -a -G supervisor your-username
```
