### Funzioni di Hash e algoritmi 
Una seconda primitiva crittografica è la funzione hash sicura, chiamata in causa in moltissimi meccanismi e servizi. Gli algoritmi che la realizzano devono in generale presentare le seguenti quattro proprietà:
- **efficienza**: “il calcolo di H(x) è computazionalmente facile per ogni x”.
    - implementazione con schema di **compressione iterata**
        - il messaggio viene suddiviso in tanti blocchi di dimensione pari a r bit
        - vari stadi di compressione in cui una funzione di hash *f* calcola l'hash prendendo gli r bit del blocco e gli *n* bit dell'hash dello stadio precedente
        - **NB**: m||S funziona; S||m no
- **robustezza debole alle collisioni**: “per ogni x è infattibile trovare un y ≠ x tale che H(y) = H(x)”.
    - l'input non lo decide l'attaccante
- **resistenza forte alle collisioni**: “è infattibile trovare una qualsiasi coppia x, y tale che H(y) = H(x)”.
    - l'input lo decide l'attaccante
- **unidirezionalità**: “per ogni h è infattibile trovare un x tale che H(x) = h”.
    - importante perchè **permette di fare una sorta di cifratura senza chiave**




### Quanti bit deve avere un hash per essere sicuro?
(i passaggi possono essere saltati)

in questo contesto essere sicuro significa essere resistenti alle collisioni, sia debolmente che fortemente.

- la probabilità di collisione scala con O(2^n) e quindi per garantire la resistenza debole mi bastano 128 bit
    - bastano anche per la sicurezza forte?


paradosso del compleanno (non sto fissando il giorno) <-> resistenza forte (non sto fissando l'hash)
- è molto più facile attaccare la resistenza forte!

**NB**: per la resistenza forte il numero di tentativi scala come O(2^(n/2)) -> l'impronta deve essere almeno 256 bit 