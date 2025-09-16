# Homework 3

## Password Scheme: Image Steganography

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

## Contribution Table

| Student ID | Works |
| - | - |
| 314581015 | register login procedure |
| 313581047 | image steganography idea |
| 313581038 | security and usability analysis |
| 313581055 | advanced protection of image leakage |
| 412581005 | storage precedure |
