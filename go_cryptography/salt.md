### Rainbow Tables
Rainbow tables are a way for attackers to get around slow KDFs. Attackers can **pre-hash all of the common passwords** once, then **compare this list (called a "rainbow table") to the list of hashes in a compromised database** and see if any match.

If the hashes match, then the attacker will know the original password, which they might then be able to use to log in to the user's bank account (or any other place they use the same password).

### Salts to the Rescue
A salt is a random chunk of data added to a password before it is hashed so that its **output hash will differ from the hash of the same password with a different salt**.
- Even if two users have the same password, their hashes will be different due to different salts.

For example:

    digest := hash(password+salt)
    // save the digest AND salt
    // checking the hash now involves the salt, hash, and the password to check

A rainbow table is ineffective against a salted database, as long as the salt generated for each password is unique.



### Salts and dictionary attacks
Since the salt is stored, an attacker can still perform a dictionary attack, but they must:
- Recompute hashes for **every salt+word combination**.
- Do it individually per user.

This slows down the attack massively, especially at scale.

Example | Let’s say an attacker has a dictionary of 1 million words and a database of 1 million users:
- Without salt: Only 1 million hashes to compute.
- With unique salt per user: 1 million × 1 million = 1 trillion hashes needed.