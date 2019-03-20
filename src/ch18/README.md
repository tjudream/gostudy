#24 CSP并发机制

## CSP vs Actor
* 和Actor的直接通讯不同，CSP模式则是通过Channel进行通讯的，更松耦合一些
* Go中的channel是有容量限制并且独立于处理Groutine，而如Erlang，
Actor模式中的mailbox容量是无限的，接收进程也总是被动地处理消息