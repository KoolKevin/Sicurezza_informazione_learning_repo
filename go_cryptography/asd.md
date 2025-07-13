### considerazioni pratiche

1. Cryptography is about securing data. It provides a method for secure communication across an insecure medium such as the Internet

2. Symmetric encryption should typically just be used to encrypt **information for oneself**. If you're encrypting data **for someone else**, you should use asymmetric encryption.
    - There are a couple of reasons for this:
        1. You probably don't have a way to securely transport the key to another person
        2. When using asymmetric encryption, the decrypter will only be able to decrypt, so they can't modify the information in any way, giving you added protection

3. Keys are just strings of bytes used to secure data


