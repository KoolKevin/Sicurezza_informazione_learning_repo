### considerazioni pratiche

1. Cryptography is about securing data. It provides a method for secure communication across an insecure medium such as the Internet

2. Symmetric encryption should typically just be used to encrypt **information for oneself**. If you're encrypting data **for someone else**, you should use asymmetric encryption.
    - There are a couple of reasons for this:
        1. You probably don't have a way to securely transport the key to another person
        2. When using asymmetric encryption, the decrypter will only be able to decrypt, so they can't modify the information in any way, giving you added protection

3. Keys are just strings of bytes used to secure data

4. A cipher is just a set of algorithms for performing encryption or decryption.

5. A cipher has perfect security if an attacker who has access to only the ciphertext can infer absolutely nothing of interest about the plaintext.
    - one time pad is an example of a cipher with perfect security

6. Most production ciphers are not perfectly secure, but are "close enough". In short, trade-offs are made that add to the practical security of the system while sacrificing the perfect theoretical security of the cipher itself.
    - The big problem with the One Time Pad is that the key needs to be the same length as the message. That means to encrypt a 128 GB hard drive, I'd need a 128 GB key!! That's just not practical.