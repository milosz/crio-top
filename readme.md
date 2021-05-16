# crio-top
Minimalistic **customizable remote top** like application.

The idea came from [Warewulf Top (wwtop)](https://www.limulus-computing.com/Limulus-Manual/doku.php?id=monitoring_system_resources).

This is a playground to learn Go language.


# sample output
Sample output using provided configuration.

```
$ crio-top --configuration configuration.yaml 
```

```
Date: Sun May 16 21:21:02 CEST 2021

              Server machine  lavg1  lavg5 lavg15       memory         swap  conn temp  timestamp 
             desktop  x86_64   2.24   4.30   4.37    11Gi/15Gi  889Mi/979Mi   388   68 1621192861 
      rpi-archivebox  armv7l   0.13   0.12   0.11  110Mi/973Mi   24Mi/1.0Gi    74   56 1621192826 
       rpi-nextcloud  armv7l   0.52   0.66   0.63  307Mi/7.7Gi        0B/0B    83   52 1621192861 
```

# sample configuration

Sample configuration.

```
application:
  header:
    width:
      server: 20
    status: true
  refresh:
    window: 2
    data: 2

  servers:
    - name: desktop
      server: localhost
      user: milosz
      port: 22
    - name: rpi-archivebox
      server: 192.168.8.33
      user: dietpi
      port: 22
    - name: rpi-nextcloud
      server: 192.168.8.32
      user: dietpi
      port: 22

  commands:
    - name: machine
      command: uname -m
      width: 7
    - name: lavg1
      command: awk '{print $1}' /proc/loadavg
      width: 6
    - name: lavg5
      command: awk '{print $2}' /proc/loadavg
      width: 6
    - name: lavg15
      command: awk '{print $3}' /proc/loadavg
      width: 6
    - name: memory
      command: free -h | awk '/^Mem/ {print $3"/"$2}'
      width: 12
    - name: swap
      command: free -h | awk '/^Swap/ {print $3"/"$2}'
      width: 12
    - name: conn
      command: ss -s | awk '/^INET/ {print $2}'
      width: 5
    - name: temp
      command: >
        if [ "$(uname -m)" = "armv7l" ]; then
          awk '{ print int($1/1000)}' /sys/class/thermal/thermal_zone0/temp
        else
          grep -l x86_pkg_temp /sys/class/thermal/thermal_zone*/type | sed 's|/type$|/temp|' | xargs -I{} awk '{ print int($1/1000)}' {}
        fi
      width: 4
    - name: timestamp
      command: date +%s
      width: 10
```

