# Homework 4

## Password Dict

- [rockyou.txt](https://github.com/brannondorsey/naive-hashcat/releases/download/data/rockyou.txt)

    ```bash
    cd 2025-NetworkSecurity/homework4
    curl -L -o rockyou.txt.gz "https://github.com/brannondorsey/naive-hashcat/releases/download/data/rockyou.txt"
    ```

## P3HTTP

- ApacheMD5

    | User | Cracked | Method |
    | - | - | - |
    | admin | KKKKKKK | rockyou |
    | tomcat | tivoli | rockyou |
    | xampp | | |
    | server_admin | | |
    | demo | microsoft | rockyou |
    | QCC | iforget | rockyou |
    | cxsdk | electric | rockyou |
    | ovwebusr | 77777777 | rockyou |
    | role1 | | |
    | j2deployer | | |

## P3Secure

- MD5

    | Hash | Password | Method |
    | - | - | - |
    | e49201c3a8f548902b9ae9f16638f879 | 0890003871 | rockyou |
    | 19cf9dda4107b300d3218702df95c76d | nailz07 | rockyou |
    | c6281df39e8ade06c6cc9e0095fd5c0f | rksmbffs | rockyou |
    | a54034981409ed58d584dc9051853ddb | hidalgo212 | rockyou |
    | f58291f81868320f11235d9b9d416115 | aq12wsxz | rockyou |
    | ce1c96461fbb2ad92fffcafafe85d0d1 | CAROLIAN | [Hashes.com](https://hashes.com/zh/decrypt/hash) |
    | c6177167ebb2c37352c3a63f6fa0c39d | 19821983 | rockyou |
    | 5993428babd2cb253834e06de1800916 | netopia | rockyou |
    | bebc51b6f0bbd5da67950200a89026f6 | Autumn2018 | [Hashes.com](https://hashes.com/zh/decrypt/hash) |
    | 456c5a41af2eb09ac0ba0eb64f614887 | eeyore | rockyou |

- SHA3-256

    | Hash | Password | Method |
    | - | - | - |
    | 1074f17769cc2dfc0d65f713a7d8c4fd97fc78c69cfa13263b07b0e40b3cf83a | sweetlove |  rockyou |
    | 94f72dc2ea6bfae657b0ee3d5adb992aa669f6c4141717344e24e873dc09be04 | shunkoko | rockyou |
    | 19c743dc300d52fc93b5ee8c6d224f3beb8a05079e6439855cdae7e55bf16ef0 | mrzdale08 | rockyou |
    | 20e5b0556c431db9a147c3f73a0ae03d12f5ef391d277cd59ff0f2dd98198ec5 | minot24601 | rockyou |
    | a44cf105063b06bbb160c22058e9c3137c8ef424ae72f981d73b10fdc743026f | loverboydj242 | rockyou |
    | 74151544815c4a0153c2e7dfabcfd066d510d6996148d6c02f246c9c497bd15c | bear1194 | rockyou |
    | 745af7302284f80ddadf6893f64e247334aa899bfe90512a59aa41ea2863f56a | | |
    | 9d34ebe967a790ada61cfa2b4e16671bfb18f0ff59296f24a0eec20dacc5ece3 | | |
    | 0ecd9ac47c8e4b059c2b97db9657f80f203454ac8fcb01976e1decdb30af2510 | bambam | rockyou |
    | 3b2918324171f88304baee77d71cc0abd40e12f16f9a22404736000f00a7c7b6 | maddie1 | rockyou |

## P3SSH

attack.waynechiu.cc:3022

- SHA512

    |User | Password | Method |
    | - | - | - |
    | postfix | password | rockyou |
    | oracle | whatever | rockyou |
    | webmaster| | |
    | tcpdump | | |
    | linaro | | |
    | hplip | love12 | |
    | unscd | !@#$%^& | |
    | zabbix | vassiliki | |
    | omsagent | | |
    | xpdb | | |

## P3Windows

attack.waynechiu.cc:3089

- CMD

    ```bash
    impacket-secretsdump -sam SAM -system SYSTEM LOCAL
    gunzip -c /usr/share/wordlists/rockyou.txt.gz > /tmp/rockyou.txt
    ```

    ```bash
    john full_hashes.txt --format=NT --wordlist=/tmp/rockyou.txt
    john full_hashes.txt --format=NT --wordlist=/usr/share/set/src/fasttrack/wordlist.txt
    john full_hashes.txt --format=NT --wordlist=/usr/share/metasploit-framework/data/wordlists/default_pass_for_services_unhash.txt
    john full_hashes.txt --format=NT --show
    ```

| User | Password | Dict |
| - | - | - |
| Administrator | 123123123 | rockyou |
| Guest | (no-pwd) | rockyou |
| DefaultAccount | (no-pwd) | rockyou |
| WDAGUtilityAccount | | |
| db2admin | Manager | rockyou |
| sa | sqlsqlsqlsql | fasttrack/wordlist |
| owner | | |
| secure | | |
| admin | trinity | rockyou |
| demo | | |
| IEUser | personales | |
| nmt | ironport | metasploit-framework/data/wordlists/default_pass_for_services_unhash |
| user | | |
| john | trabajador | |

## Contribution Table

| Student ID | Works | Percentage |
| - | - | - |
| 314581015 | hash cracking & hashes validate | 20% |
| 313581047 | golang coding & multi-routine architecture | 20% |
| 313581038 | hash cracking & test file writing | 20% |
| 313581055 | hash cracking & cracking idea checking & report writing | 20% |
| 412581005 | hash cracking & kali using guide | 20% |
