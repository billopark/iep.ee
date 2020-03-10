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

[Try it now](https://toolbox.googleapps.com/apps/dig/#A/10.0.0.1.iep.ee)


#### Common Use-case

`10.0.0.1.iep.ee` will be mapped to `10.0.0.1`

`a.b.c.10.0.0.1.iep.ee` will be mapped to `10.0.0.1`

`10-0.0-1.iep.ee` will be mapped to `10.0.0.1`

`a-b-c-10-0-0-1.iep.ee` will be mapped to `10.0.0.1`


#### Why it is useful?

This service is needed when you need domains for local server. (For cookie or etc.) 
Traditionally, you need to modify your hosts file. But with this service, you don't need it.


Contribution
---
Contributions, issues, pull requests are welcomed!
