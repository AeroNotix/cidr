cidr
----

Little tool to do CIDR calculations. I often need to look at CIDR
ranges and quite honestly it's not an immediate ezpz mental arithmetic
I can do. So I just wrote a tool to do the most common things I need.

Uses the helpful `cidr` [library](https://github.com/3th1nk/cidr)

Examples:

```
; cidr subnet 100.64.0.0/10 --num 2
100.64.0.0/11
100.96.0.0/11
; cidr range 100.64.0.0/11
100.64.0.0 100.95.255.255
; cidr count 100.64.0.0/11
2097152
; cidr print 100.64.0.0/30
100.64.0.0
100.64.0.1
100.64.0.2
100.64.0.3
```
