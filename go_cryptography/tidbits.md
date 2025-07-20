### considerazioni pratiche

1. Cryptography is about securing data. It provides a method for secure communication across an insecure medium such as the Internet

2. Symmetric encryption should typically just be used to encrypt **information for oneself**. If you're encrypting data **for someone else**, or more in general comunicating with another party, you should use asymmetric encryption.
    - There are a couple of reasons for this:
        1. You probably don't have a way to securely transport the key to another person
        2. When using asymmetric encryption, the decrypter will only be able to decrypt, so they can't modify the information in any way, giving you added protection

3. Keys are just strings of bytes used to secure data

4. A cipher is just a set of algorithms for performing encryption or decryption.

5. A cipher has perfect security if an attacker who has access to only the ciphertext can infer absolutely nothing of interest about the plaintext.
    - one time pad is an example of a cipher with perfect security

6. Most production ciphers are not perfectly secure, but are "close enough". In short, trade-offs are made that add to the practical security of the system while sacrificing the perfect theoretical security of the cipher itself.
    - The big problem with the One Time Pad is that the key needs to be the same length as the message. That means to encrypt a 128 GB hard drive, I'd need a 128 GB key!! That's just not practical.

7. A stream cipher is a symmetric key cipher where plaintext digits are combined with a key stream. In a stream cipher, each plaintext digit is encrypted one at a time with the corresponding digit of the keystream, to give a digit of the ciphertext stream.
    - operating one digit at a time reduces memory usage, since we don't need to hold all of the plaintext or key in memory
    - examples of stream ciphers are: RC4, Salsa20

8. A block cipher is a deterministic algorithm that operates on fixed-length groups of data, called blocks. Like stream ciphers, block ciphers are a kind of symmetric encryption. Block ciphers use fixed-size blocks and fixed-size keys to encrypt/decrypt variable size data. Block ciphers can operate on messages of any length (thanks to chunking and padding), and the key doesn't even need to be the same length as the message or the block.
    - examples of block ciphers are: AES, DES

9. Messages in a block cipher are broken up into blocks, and each block is encrypted separately. For example, let's say we are using a cipher that requires 256-bit blocks. We have a message of 650 bits that we want to encrypt. The algorithm would break that message up into three blocks:

    block1 = first 256 bits
    block2 = next 256 bits
    block3 = last 138 bits

The last block is then padded with extra garbage bits so that it also has 256 bits of data. That padding is stripped off when the message is decrypted.

10. Which Is Best? One isn't necessarily better than the other. Stream ciphers are typically used when a stream of data must be encrypted bit by bit. For example, when encrypting a stream of video data in transit. Block ciphers are more typically used on static data, things like passwords in a password manager


11. What's an IV? An IV, or initialization vector, is a random value that is used to initialize a block cipher. It is used to ensure that the same plaintext always encrypts to a different ciphertext. Without an IV, the same plaintext would always encrypt to the same ciphertext which is a big security vulnerability.

12. RSA is relatively slow. Because of this, it is not commonly used to directly encrypt user data. More often, RSA is used to **transmit shared keys for symmetric-key cryptography**, as it does in TLS/HTTPS.

13. At the end of the day, p, q, tot and d should all be kept secret, there's no reason to share anything other than (n, e).