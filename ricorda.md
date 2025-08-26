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

Riassumendo, abbiamo che PRNG crittograficamente sicuro è caratterizzato dalla produzione di una sequenza di bit
- casuali
- imprevedibili in quanto scelti da un seme che seleziona una sottosequenza tra le molte presenti nel periodo grandissimo del PRNG
- seme deve essere tenuto segreto ed è anch'esso imprevedibile (i punti 2 e 3 sopra servono a questo) altrimenti la sequenza viene scelta viene svelata
    - per questo si impiega TRNG


4. Cifrari a flusso
    - sono in pratica uno xor con un flusso di chiave pseudocasuale
    - possono essere
        - a flusso sincrono             -> flusso di chiave generato con un PRNG (funzioni di F e G)
        - a flusso autosincronizzante   -> flusso di chiave generato con un registro a scorrimento (in cui ci finisce in retroazione l'uscita prodotta) e una funzione F 
    - a causa dello xor si ha
        - two-time-key vulnerability    -> c1 ^ c2 = m1^k ^ m2^k = m1 ^ m2
        - malleabilità: modifiche al testo cifrato hanno un effetto prevedibile sul testo in chiaro
            - c ^ m' -- viene decifrato come --> m* = m^k^m' ^k = m ^ m'
            - Se l’attaccante conoscesse il messaggio (o se lo immagina, perché m, magari, è fortemente strutturato) potrebbe cambiare parte del testo a suo piacimento.
            - Se l’intrusore pensa che i primi bit siano corrispondenti a “from Bob”, e vuole cambiare il mittente, deve trovare quel p che in xor con "from Bob ..."  produce “from Eve ...” (facile).
    - se si utilizza un cifrario a flusso sappiamo che è malleabile e quindi è fondamentale implementare meccanismi che garantiscano integrità


5. Cifrari a blocco
    - conclusioni ECB
        - Pro:
            - Alto parallelismo (efficente): i blocchi sono tra loro indipendenti e quindi possono venire cifrati in parallelo
                - ottimo per cifrare roba che sta in un blocco (chiavi) o dati totalmente non strutturati (per cui non si riesce a sostituire blocchi)
            - Non propagazione degli errori.
        - Contro:
            - **Deterministico**, a blocchi identici di testo in chiaro corrispondono blocchi identici di testo cifrato.
            - suscettibile ad attacchi di malleabilità (sostituzione di blocchi cifrati con la stessa chiave intercettati) che di chosen plaintext (controllo due cifrati e vedo se combaciano)
                - funzionano se la chiave è la stessa (nel secondo caso magari qualcuno cifra per me (voto elettronico) o magari posso ingannarlo a farlo)
                - motivo in più per cambiarla spesso
    - Un buon cifrario simmetrico dovrebbe:
        - produrre ciphertext diversi anche partendo da plaintext identici
            - altrimenti posso vedere se i cifrati combaciano (chose plaintext)
        - mascherare le regolarità del plaintext facendo **dipendere ogni blocco del ciphertext da qualcos'altro oltre corrispondente blocco del plaintext (e la chiave)**
            - altrimenti malleabilità attraverso sostituzione di blocchi
        - questo è quello che fanno le prossime modalità di cifratura
    - CBC
        - ogni blocco di cifrato si ottiene cifrano il relativo blocco di plaintext in xor con il cifrato precedente
            - Sorgente (cifratura): c0 = IV, ci = Ek(ci-1 ⊕ mi)
            - Destinazione (decifratura): IV = c0, mi = Dk(ci) ⊕ ci-1 = ci-1 ⊕ mi ⊕ ci-1
        - per il primo blocco si utilizza un IV che cambia sempre
            - in questa maniera messaggi uguali o che iniziano in maniera identica vengono cifrati in maniera diversa e non si è suscettibili al chosen plaintext attack citato sopra
        - In CBC **l'attacco di malleabilità è infattibile**
            - se si prova a sostituire un blocco cifrato tutti i successivi vengono decifrati male
            - in ogni caso è impossibile trovare un blocco cifrato nella stessa maniera dato che l'IV cambia sempre
        - modalità di cifratura abbastanza lenta in quanto non parallelizzabile
    - OFB e CFB
        - sono implementazioni di cifrari a flusso basati su cifrari a blocchi:
            - OFB realizza un Cifrario a flusso sincrono
            - CFB un Cifrario a flusso con auto-sincronizzazione
        - impiegano la sola trasformazione E per generare un flusso di chiave
            - identico lato sorgente e lato destionazione
        - il flusso di chiave è sempre diverso grazie ad un IV che fa da seed
        - malleabilità quindi prevenuta dato che non si riescono a trovare due messaggi cifrati con lo stesso flusso di chiave
    - Counter
        - simile ad OFB però opera su blocchi
            - il comportamento è quindi quello di un cifrario a flusso sincrono (con le relative considerazioni sulla sincronizzazine)
        - impiegano la sola trasformazione E per generare un flusso di chiave
            - identico lato sorgente e lato destionazione
        - di nuovo non malleabile perchè il seed (IV) non si ripete
    - Beast attack (block injection attack)
        - se conosco il CBC residue che verrà utilizzato (IV prevedibile) posso costruire un messaggio che annulla roba (xor) e viene cifrato come un messaggio precedente su cui sto facendo ipotesi
        - rompo il non determinismo di CBC e faccio un attacco simile a quello che facevo con ECB
    - paradosso del compleanno per cifrari a blocchi
        - La probabilità che un attaccante trovi due blocchi di testo cifrato uguali scala con 2^(n/2) e non con 2^n (paradosso del compleanno)
            - con n dimensione del blocco
        - dopo 2^(n/2) cifratura continuare a cifrare con la stessa chiave non è più consigliato perchè diventa probabile produrre cifrature uguali
        - se produce cifrature uguali posso sfruttare le proprietà dello xor (varia in base alla modalità di cifratura in questione) per attuare un **two time key attack**
        - un cifrario a blocchi è sicuro se e solo se i blocchi hanno dimensione >= 128 bit