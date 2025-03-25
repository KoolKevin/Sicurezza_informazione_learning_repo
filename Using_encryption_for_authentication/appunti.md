### Encryption Algorythms
With a conventional encryption algorithm, such as DES, the same key is used for both encryption and decryption. Authentication depends upon the two participants in a conversation **being the only two principals who know the key** that is being used to encrypt the transmitted material. 
- so che sei effettivamente chi dici di essere dato che l'unico che può criptare in questo modo è l'unica altra persona con la chiave segreta

With a public-key encryption algorithm, a concept originated by Diffie and Hellman, two keys are necessary:
- one that is used in the conversion of cleartext to ciphertext,
- and another that is used in the conversion of ciphertext to cleartext.

Furthermore, knowledge of one key gives no help in finding the other, and the two keys will act as inverses for each other. 

Elegant systems may be devised in which each principal has one public key and one secret key.
- Anyone may encrypt a communication for A using his public key, but only A can decrypt the result using his secret key. 
- Likewise, only A can encrypt messages that will decrypt sensibly with A's public key. The first example of a public-key encryption algorithm 

In questa situazione l'autenticazione è data sempre dalla segretezza di una chiave (nell'esempio quella privata di A)
- so che sei effettivamente chi dici di essere dato che riesci a decifrare quello che dico, ed io riesco a decifrare quello che dici

### Authentication Servers
Quello che noi abbiamo chiamato KDC...

With both kinds of encryption the basis of authenticated communication is a secret key belonging to each principal using the network, and there is need for an authoritative source of information about these keys. We use the term authentication server for a server that can deliver identifying information computed from a requested principal's secret key.
- le chiavi segrete di ogni utente sono condivise con l'authentication server