Le feistel network sono un framework con cui implementare cifrari simmetrici. cifrari diversi si differenziano solo per la funzione F di cifratura usata in ogni round.
- esempio di cifrario che usa feistel networks è DES

really cool thing about feistel networks: 
- a causa della reversibilità dello xor, la decifratura è uguale alla cifratura! anche nel caso in cui la funzione di cifratura del round sia non reversibile
    - il lato destro viene preservato dato che viene passato così come allo xor del livello successivo
- la cifratura e la decifratura del messaggio sono implementate praticamente dallo stesso algoritmo (l'ordine delle chiavi è invertito). Implementare cifrari che usano feistel network è di conseguenza molto facile

