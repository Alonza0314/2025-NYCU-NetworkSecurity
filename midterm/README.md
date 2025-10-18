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
