# diutils
Diutils - A console utility for easily saving and connecting to servers via ssh

# Install
```
go install github.com/dinaxu-attack/diutils
```

# Usage

```
Save: diutils new -a <ALIAS> -ip <IP> -p <PORT (optional)>  -l <LOGIN> -pass <PASSWORD>
Example: diutils new -a myserver -ip 123.123.123 -p 22 -l root -pass qwerty

Connect: diutils -ssh -a <ALIAS>
Example: diutils -ssh -a myserver

List: diutils -list
```

# Showcase
![image](https://user-images.githubusercontent.com/102496559/221966606-4cc97871-e1d0-4edd-92bb-e5e9d081198f.png)
