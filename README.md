# iep.ee

---
[iep.ee](http://iep.ee) (sounds like `a-i-pi-i-i`) is simple DNS wildcard service.
It was written in Go.


How to use iep.ee
---

iep.ee will parse subdomain part using `-` `.` as seperator.
Splited subdomain part will be parsed as IPv4. 
Parsed IP will be A record of that domain.

For example, `10.0.0.1.iep.ee` will be splited into `10, 0, 0, 1` which means `10.0.0.1`.
So `10.0.0.1.iep.ee` 's A Record is `10.0.0.1`.

Common Use-case

`10.0.0.1.iep.io` will be mapped to `10.0.0.1`

`a.b.c.10.0.0.1.iep.io` will be mapped to `10.0.0.1`

`10-0.0-1.iep.io` will be mapped to `10.0.0.1`

`a-b-c-10-0-0-1.iep.io` will be mapped to `10.0.0.1`

Contribution
---
Contributions, issues, pull requests are welcomed!
