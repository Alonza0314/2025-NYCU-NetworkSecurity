# 2025 Network Security Midterm Report

- [Homework 1 -  Kali Linux and NYCU Security Policy](#homework-1----kali-linux-and-nycu-security-policy)
- [Homework 2 - Encrypt and Decrypt A File](#homework-2---encrypt-and-decrypt-a-file)

    - [Build](#build)
    - [Encrypt](#encrypt)
    - [Decrypt](#decrypt)
    - [Description](#description)
    - [Feature](#feature)
    - [Demo](#demo)
    - [HW2 Contribution Table](#hw2-contribution-table)

- [Homework 3 - Design a Pasword Scheme](#homework-3---design-a-pasword-scheme)

    - [Password Scheme: Image Steganography](#password-scheme-image-steganography)
    - [HW3 Contribution Table](#hw3-contribution-table)

- [Homework 4 - Password Recovery](#homework-4---password-recovery)

    - [Password Dicts](#password-dicts)
    - [P3HTTP](#p3http)
    - [P3Secure](#p3secure)
    - [P3SSH](#p3ssh)
    - [P3Windows](#p3windows)
    - [HW4 Contribution Table](#hw4-contribution-table)

## Homework 1 -  Kali Linux and NYCU Security Policy

In this homework, we installed a Kali Linux system, and learn how to configure VM with virtual box for windows users and VMware for Mac users, like this:

![kali](kali.png)

## Homework 2 - Encrypt and Decrypt A File

In thie homework, we learned how to use crypto package to encrypt and decrypt files.

Code is at: [homework2 directory](../homework2/).

### Build

```bash
cd homework2
make
```

### Encrypt

```bash
./build/fileCrypt encrypt -i <inputFile> -o <outFile> [-k <64-hex-key>]
```

### Decrypt

```bash
./build/fileCrypt decrypt -i <inputFile> -o <outFile> -k <64-hex-key>
```

### Description

- Encryption

    - Validates the input file exists and reads its bytes into memory.
    - Uses AES-256 in GCM (AEAD) mode for confidentiality and integrity.
    - Key handling:
        - If no key is provided, generates a random 32-byte key with crypto/rand and prints it as 64 hex characters. Save this key for decryption.
        - If a key is provided, it must be 64 hex characters (32 bytes).
    - Generates a fresh random nonce of size gcm.NonceSize() for every encryption.
    - Produces output as nonce || ciphertext (the nonce is prefixed to the ciphertext bytes).
    - Writes the result to the specified output file and logs helpful messages.

- Decryption

    - Validates the ciphertext file exists and requires a key (64 hex characters).
    - Reads the file bytes and splits them into: nonce = first gcm.NonceSize() bytes, ciphertext = remaining bytes.
    - Reconstructs AES-256-GCM with the provided key and decrypts with gcm.Open, which also verifies authenticity.
    - Decryption fails if the key, nonce, or data is invalid/tampered (authentication error).
    - Writes the recovered plaintext to the specified output file and logs helpful messages.

### Feature

- AES-256 key (32 bytes) for symmetric encryption
- GCM mode (AEAD) for confidentiality and integrity (auth tag included)
- Random per-file nonce, prefixed to output for stateless decryption
- Optional random key generation when no key is provided (printed as hex)
- Simple CLI with input/output paths and hex key handling

### Demo

1. Prepare a plaintext.txt

    ![plaintext](../homework2/images/plaintext.png)

2. Encrypt as ciphertext.txt

    ![ciphertext](../homework2/images/ciphertext.png)

3. Decrypt from ciphertext.txt as newPlaintext.txt

    ![newPlaintext](../homework2/images/newPlaintext.png)

### HW2 Contribution Table

| Student ID | Works | Percentage |
| - | - | - |
| 314581015 | Code section & Readme document | 20% |
| 313581047 | Code section & Github repo establishment | 20% |
| 313581038 | Execution & Build test | 20% |
| 313581055 | Execute procedure | 20% |
| 412581005 | Kali env setup & Kali usage help | 20% |

## Homework 3 - Design a Pasword Scheme

In this homework, our team designed a creative password scheme: Image Steganography.

### Password Scheme: Image Steganography

1. Concept

    This scheme uses **image steganography** to hide a randomly generated login code inside an image, which acts as the user’s “password card.”

    During registration, the system embeds the code into the image, and the user saves it.

    During login, the user uploads the image, the system extracts the code, and verifies it for authentication.

2. How it works?

    - Registration:

        - The system generates a random login token (password) for each user.
        - The token is optionally encrypted and embedded into a user-provided or system-generated image (PNG or JPEG).
        - Then, the user downloads and save this image.

    - Login

        - The user uploads the image.
        - The system extracts the token and compares it with stored hash.
        - If the comparison succeeds, authentication is granted.

    - Storage

        - The server stores only the hash of the token or image fingerprint, not the plaintext token, reducing leakage risk.

3. Security and Usability Analysis

    | Perspective | Security | Usability |
    | - | - | - |
    | System | Does not store plaintext password; only hashes or fingerprints are stored, reducing leakage risk. | Requires additional program to embed/extract the token from the image. |
    | User | Image is hard to guess, more secure than plain text passwords. | No need to remember complex passwords; must safely keep the image. |
    | Attacker | Without the image, login is impossible; pairing with PIN or periodic updates mitigates leaked image risks. | If the image leaks and no other protection is applied, it may be misused. |

4. Protection Against Image Leakage

    As with traditional passwords, leaked images can be misused. This scheme adds simple protective measures to reduce the risk:

    - **Two-factor authentication**: Require a PIN or gesture in addition to the image.
    - **Periodic updates**: System generates new images regularly; old images expire.
    - **Hash verification**: The server only stores image fingerprints to prevent direct misuse if the database leaks.

5. Conclusion

    This scheme hides the “password” within an image, making it intuitive and portable for users while being more creative than traditional text passwords.

    Although image leakage cannot be completely prevented, combining PINs, periodic updates, and server-side protections significantly reduces potential misuse.

### HW3 Contribution Table

| Student ID | Works | Percentage |
| - | - | - |
| 314581015 | register login procedure | 20% |
| 313581047 | image steganography idea | 20% |
| 313581038 | security and usability analysis | 20% |
| 313581055 | advanced protection of image leakage | 20% |
| 412581005 | storage precedure | 20% |

## Homework 4 - Password Recovery

In this homework, we experienced hash cracking. We used not only Kali's internal tools but also wrote some small programs to crack the hashes.

Code is at: [homework4 directory](../homework4/).

### Password Dicts

- [rockyou.txt](https://github.com/brannondorsey/naive-hashcat/releases/download/data/rockyou.txt)
- [weakpass_4.txt](https://weakpass.com/download)

### P3HTTP

- ApacheMD5

    | User | Hash | Cracked | Method |
    | - | - | - | - |
    | admin | $apr1$GToZ6eNr$wn/MOgnx/f4Y0diEa7BNa1 | KKKKKKK | rockyou |
    | tomcat | $apr1$no3kBUxm$pk411ceCDJpW2oNs0WB3U0 | tivoli | rockyou |
    | xampp | $apr1$8WlYVzJL$C8ZhuDziOtni3HFtvKzGe0 | backupexec | weakpass_4 |
    | server_admin | $apr1$y758YD1Y$R1lKrNG6SKSjd0Nmb3v5U1 | SQLSQLSQLSQL | weakpass_4 |
    | demo | $apr1$JwOVMfoS$zMOzDn/xOMZ3VuEQO6IIj/ | microsoft | rockyou |
    | QCC | $apr1$IKjpEtXZ$UsaHuXDT3z8OZaHtBab0X. | iforget | rockyou |
    | cxsdk | $apr1$zMVw.RLi$kzXVI2SkGcgmioh1/R71L1 | electric | rockyou |
    | ovwebusr | $apr1$rOLMqcAU$DiIw.q5MBQOHDsgtE7k7u. | 77777777 | rockyou |
    | role1 | $apr1$.MLLcn6R$izw9nnviyDmEO3t0lrBpB0 | sestosant | weakpass_4 |
    | j2deployer | $apr1$qfumuNjF$sH4FzC3E7yHoPjj2PTU5I1 | pantiled | weakpass_4 |

### P3Secure

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

### P3SSH

attack.waynechiu.cc:3022

- SHA512

    |User | Password | Method |
    | - | - | - |
    | postfix | password | rockyou |
    | oracle | whatever | rockyou |
    | webmaster| | |
    | tcpdump | | |
    | linaro | | |
    | hplip | love12 | rockyou |
    | unscd | !@#$%^& | rockyou |
    | zabbix | vassiliki | rockyou |
    | omsagent | | |
    | xpdb | | |

### P3Windows

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
| owner | init1234 | rockyou |
| secure | | |
| admin | trinity | rockyou |
| demo | | |
| IEUser | personales | |
| nmt | ironport | metasploit-framework/data/wordlists/default_pass_for_services_unhash |
| user | | |
| john | trabajador | |

### HW4 Contribution Table

| Student ID | Works | Percentage |
| - | - | - |
| 314581015 | hash cracking & hashes validate | 20% |
| 313581047 | golang coding & multi-routine architecture | 20% |
| 313581038 | hash cracking & test file writing | 20% |
| 313581055 | hash cracking & cracking idea checking & report writing | 20% |
| 412581005 | hash cracking & kali using guide | 20% |
