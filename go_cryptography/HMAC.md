Here's a naive implementation:

    secretKey = 'thisIsASecretKey1234'
    childKey1 = 'thisIsASe'
    childKey2 = 'cretKey1234'
    hash(childKey1 + hash(childKey2 + 'the message we want to send'))

This is a simplified version of the function given in RFC-2104.

### Why Use HMAC? Why Do We Need to Hash Twice?
With some MACs, depending on the hash function, it is possible to change the message (without knowing the key) and obtain another valid MAC. This is called a **length extension attack**. There are no known extension attacks against the current HMAC specification, so you should prefer HMACs over MACs.