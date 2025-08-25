1. 
La sicurezza dipende da come è stato scelto il segreto (chiave). Se l’utente, per potersela facilmente ricordare a memoria, ha usato un’informazione che l’intruso può prevedere (ad esempio il suo nome, o il suo anno di nascita, o i dati di un suo familiare, ecc.), il tirare ad indovinare ha una notevole probabilità di successo.

La prima difesa preventiva è quindi quella di usare dati casuali.
- R12: ”I bit della stringa che rappresenta un dato segreto devono essere equiprobabili ed indipendenti ”
- Ciò richiede la disponibilità di un nuovo meccanismo per la sicurezza, il generatore di numeri casuali o RNG (Random Number Generator)

Abbiamo già anche commentato il fatto che la complessità computazione non mette in conto il caso più favorevole per l’intruso; utile è dunque la contemporanea adozione di un secondo provvedimento.
- R13: ”un dato segreto deve essere frequentemente modificato”. 
    - In questa maniera non concediamo all’intruso il tempo di svolgere molte prove.
    - per lo stesso motivo, usare tante volte lo stesso segreto per molti messaggi non è raccomandabile

2. 
tutto il file su funzioni di hash


3. 
Nelle applicazioni crittografiche la sola casualità non è sufficiente. Occorre, infatti, anche l’imprevedibilità: ```un intruso che è riuscito ad intercettare l’uscita o ad individuare, in tutto o in parte, lo stato del generatore non deve poter dedurre da quale seme sono partiti i calcoli e/o quali saranno i prossimi valori generati.```

Un generatore pseudocasuale che ha anche la proprietà di imprevedibilità è detto **PRNG crittograficamente sicuro** o CSPRBG (Cryptographically Secure PseudoRandom Bit Generator).

Per conseguire imprevedibilità occorre che:
- il periodo sia grandissimo (10^50, 10^100), per poterlo suddividere con il seed in moltissime sottosequenze;
- il seme sia imprevedibile e tenuto segreto
    - tipicamente il seme viene generato da un TRNG
    - **non siamo tornati punto e a capo con la riproducibilità?** Penso di no dato che il seme diventa un segreto da scambiare (di dimensione molto più limitata rispetto alla chiave di un one-time pad ad esempio)
- sia unidirezionale o la funzione di stato futuro, o la funzione d’uscita
    - per rendere impossibile ad un avversario che ha individuato uno stato il risalire agli stati precedenti ed al seme; 