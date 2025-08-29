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
    - il flusso di chiave prodotto si basa su un seed. Questo deve essere scambiato in segreto da sorgente e destinazione altrimenti un attaccante può riprodurre lo stesso flusso di chiave
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
            - notare in queste implementazini, un attaccante che conosce il seed non può comunque riprodurre lo stesso flusso di chiave dato che non ha la chiave k
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

6. Master Key e KDC vai prima a riguardarti gli appunti
    - interessante l'utilizzo della cifratura con un segreto come strumento di identificazione
        - solo se sei chi dici di essere saprai decifrare questo messaggio contenente la chiave di sessione

7. Scambi DH
    - DH anonimo
        - L’obiettivo dello scambio DH è far si che due utenti qualsiasi A e B, **senza aver preso alcun precedente accordo segreto** (no master key sia essa in KDC o meno), riescano a **condividere un dato segreto K (chiave di sessione)** dopo aver calcolato ed essersi **scambiati senza alcuna segretezza due dati YA e YB (pubblici)**. 
            - Non abbiamo più una terza parte fidata (KDC) e nemmeno una Master Key pre-concordata.  
        - A tal fine, è necessario che ciascuno dei due utenti **scelga a caso un numero X (che terrà segreto)** ed usi poi una **funzione unidirezionale F** per calcolare il numero **Y = F(X) da inviare al corrispondente**
            - in questo modo, l’intruso che riesce ad intercettare Y non dispone di algoritmi efficienti per calcolare il segreto X = F^-1(Y). 
        - Una volta avvenuto lo scambio (A manda YA e B YB), il metodo prevede che A e B dispongano di una **particolare funzione G** che garantisca ad entrambi di **ottenere lo stesso risultato K (chiave di sessione)** a partire dai dati in loro possesso:
            - __G(XA,YB) = G(XB, YA) = K__
            - **NB**: ho comunicato solo roba pubblica YA, YB, ma **HO CONCORDATO UN SEGRETO CONDIVISO**
        - Il protocollo, nella sua versione più semplice detta DH anonimo, prevede quindi solo due passi:
            1. **Generazione delle chiavi segrete XA, XB e pubbliche YA, YB**
                - A e B, dopo aver generato a caso rispettivamente i numeri 
                    - 1 < XA < p-1
                    - 1 < XB < p-1
                - (che tengono segreti), calcolano
                    - YA = g^XA mod p
                    - YB = g^XB mod p
                - e si scambiano i risultati in modo non riservato. 
            2. **Calcolo del segreto K** 
                -  A e B calcolano rispettivamente:
                    - KA = YB^XA mod p = (g^XB mod p)^XA mod p = (g^XB)^XA mod p
                    - KB = YA^XB mod p = (g^XA mod p)^XB mod p = (g^XA)^XB mod p
                - **KA = KB = K**
        - I dati **p e g, NON sono segreti**, devono essere **noti ed uguali per entrambi** e quindi vengono comunicati in chiaro. 
            - Ad esempio chi inizia il protocollo può deciderli e comunicarli al corrispondente insieme al suo dato Y
            - in fixed DH sono presenti nel certificato
        - **NB**: Se si vuole generare ogni volta una chiave di sessione diversa, bisogna rigenerare ad ogni sessione i dati X e Y
            - in alternativa, è possibile generare solo una volta i dati X e Y concordando così un pre-master-secret sempre uguale
            - sarà sufficente che i due interlocutori si scambino ognuno un proprio nonce.
            - Questi due nonce verranno concatenati al pre-master-secret e infine verrà fatto l'hash di tutto per generare il master-secret che cambia sempre
                - master_secret = H( pre_master_secret || RC || RS )
    - DH non anonimi
        - ci si avvale dei certificati per rendere lo scambio non anonimo (vogliamo sapere chi è che ci sta mandando il dato Y)
        - **NB**: TLS sfrutta questi scambi per autenticare gli interlocutori
        - **Fixed DH**
            - il dato pubblico <p, g, Y> di ciascun utente è comunicato all’altro tramite un certificato X.509v3.
            - il resto dello scambio è analogo alla modalità anonima
            - notare che dato che i dati pubblici sono in un certificato, essi non cambiano mai, per non concordare ogni volta la stessa chiave di sessione si usa il trucco coi nonce per generare un master-secret a partire dal pre-master secret
            - **NB**: la CA autentica la tripla, ma nello scambio non c'è una PoP che mi identifica il mittente; è dunque possible che un attaccante presenti un certificato che non è il suo
                - poco male, non avendo il dato segreto X non può concordare la stessa chiave di sessione 
                - non è comunque il massimo dato che (se non si controlla di avere concordato la stessa chiave come in TLS) si potrebbe mandare del testo cifrato ad un attaccante
        - **Anonymous DH**
            - analogo a scambio anonimo, solamente che gli interlocutori inviano il loro dato pubblico Y firmato, il ricevente può autenticare verificando la firma
            - qua X e Y cambiano sempre e quindi la chiave di sessione concordata è sempre diversa (per compatibilità si usano comunque i nonce)
            - anche qua non c'è identificazione, dato che un attaccante potrebbe replicare un Y firmato intercettato in comunicazioni precedenti
                - di nuovo poco male dato che non avendo X non riesce a concordare la stessa chiave
            - **NB**: è importante distinguere il requisito di autenticazione da quelle di identificazione
                - il primo richiede che i messaggi che arrivano siano effettivamente appartenenti a Bob
                    - associazione paternità del dato
                - il secondo mi richiede che i messaggi che arrivano siano stati mandati proprio da Bob
                    - le repliche sono per definizione autentiche, ma non vanno bene dato che chi me le manda non è chi ha creato il messaggio
8. Certificati
    - i certificati vengono firmati da una terza parte fidata, questo garantisce l'autenticità tra l'associazione di identità e chiave pubblica presente nel certificato
    - per identificare correttamente un interlocutore che non si conosce a priori, è necessario:
        - sia il certificato -> che associa una identità ad una chiave pubblica
        - che la POP         -> che dimostra che che il mittente è il proprietario della chiave pubblica presentata (e per transitività, identifica il mittente con l'identità presente nel certificato)

9. CRL
    - listone di tutti i certificati revocati
    - viene emessa ed aggiornata periodicamente
    - struttura simile ad un certificato nel senso che è firmata da una CA ed ha un periodo di validità (fino alla prossima emissione)
    - Può diventare molto grande è quindi è importante adottare delle tecniche che consentono di ridurre la dimensione/la quantità di dati da trasferire all'utente
        - elimino le entry relative a certificati scaduti nell CRL successive alla loro data di scadenza 
            - sono scadute non c'è bisogno di manterle 
                - se il certificato era stato revocato elimino la entry
                - se non era stato revocato non la aggiungo neanche
        - pubblico i diff -> aiuta solo la bandwidth
        - sotto-liste -> si grazie
    - problema della freschezza, tra un aggiornamento e l'altro potrebbero essere stati revocati dei certificati -> c'è anche OCSP

10. Gerarchia delle CA
    - non tutti gli utenti vengono certificati dalla stessa CA
    - le CA sono in relazione gerarchica, questo permette di trasferire la fiducia
        - se io mi fido della mia CA e la mia CA si fida di CA_B, allora anche io mi fido di CA_B
    - bisogna trovare le catene di fiducia (catene di certificati) fino ad una root-CA e verificarle
    - cross-certificates sono certificati che certificano la chiave pubblica di un altra CA
        - certificano i nodi intermedi del percorso di fiducia


11. Identificazione passiva
    - ripeto sempre lo stesso dato come prova di identità
    - sempre suscettibile ad attacco di intercettazione e replica -> assolutamente inadatto a sistemi in rete
    - per lo stesso motivo è anche impossibile fare identificazione reciproca
        - un MIM può aspettare che io gli invii le mie credenziali, non rispondermi, replicare le mie credenziali verso il destinatario legittimo spacciandosi per me

12. Identificazione attiva
    - l'identificando manda una prova di identità sempre diversa
    - abbiamo visto 3 varianti:
        - one-time pwd
        - sfida/risposta
        - zero-knowledge proofs

13. Identificazione attiva con sfida/risposta
    - esistono 3 varianti
        - hash con segreto preconcordato (pwd modem)
        - firma asimmetrica
        - cifratura asimmetrica
    - tutte e 3 le varianti usano come sfida un nonce in modo da evitare gli attacchi con replica
        - **NB**: il nonce è generato da un PRNG e come al solito occorre che sia **imprevedibile**. Se il nonce fosse prevedibile l'attaccante potrebbe spacciarsi per l'identificatore B e farsi mandare da A _H(S || RB')_; con RB' = nonce previsto dall'attaccante che B vero produrrebbe nella prossima sfida. A questo punto l'intrusore può spacciarsi per A con B nella prossima sessione dato che conosce l'hash che verrebbe usato da A alla sfida con RB'.

14. Identificazione mutua con sfida e risposta
    - se si usa la variante con hash e segreto precondiviso non si può identificare prima uno e poi l'altro come ci si aspetta
    - ho il problema che si possono tenere aperte due sessioni di identificazione
        - **attacco di interleaving (gran maestro scacchista)**: apro due sessioni contemporanee con due entità diverse
        - **Attacco di reflection**: apro due sessioni contemporanee con la stessa entità
        - in entrambi gli attacchi l'attaccante può rigirare messaggi ottenuti da sessioni diverse in cui svolge sia il ruolo di identificatore che di identificando
        - con un protocollo ingenguo, si lascia il tempo all'attaccante di svolgere più sessioni (in particolare, può aspettare che uno tenti di identificarsi presso di lui)
            - primo accorgimento utile è quindi limitare il periodo di validità delle sfide
        - un ulteriore accorgimento è rompere la simmetria dei messaggi nel protocollo e linkarli aggiungendo dentro all'hash nelle risposte
            - entrambi i nonce della sessione
            - e anche il destinatario della risposta
        - un ultimo accorgimento è quello di numerare le identificazioni (seq number) tra i due interlocutori (sia per inteleaving che per reflection).
            - in questa maniera se un attaccante ripropone una sfida/risposta il seq number non combacia

15. Cifrari asimmetrici
    - se sono a blocchi, è importante che il numero corrispondente al blocco binario, sia minore del monulo n. Questo perchè ad un testo in chiaro deve corrispondere uno e un solo testo cifrato; se se si eccede il modulo due blocchi di plaintext possono venire cifrato nello stesso modo dato che il modulo fa il giro
    - **NB**: un Cifrario asimmetrico è almeno **mille volte più lento** di un Cifrario simmetrico ed è quindi fortemente auspicabile impiegarlo con messaggi “corti”.
        - perchè?
            - dobbiamo fare delle esponenziazioni modulari e generare numeri primi grandi!
            - nel caso simmetrico devo fare solo trasposizioni e sostituzioni (con xor)
        - Ciò non ne limita l’utilità, dato che il suo uso tipico è la comunicazione della chiave di un Cifrario simmetrico (vedi cifrario ibrido) oppure la firma di messaggio per appendice.
            - il plaintext sta dentro ad un blocco
    - **NB**: è ancora più importante che qui i cifrari siano probabilistici dato che non c'è neanche bisogno di fare intercettazione per attuare attacchi di chosen-plaintext / malleabilità con sostituzione di blocchi (ricorda ECB)
        - la chiave pubblica è a disposizione di tutti e quindi tutti possono cifrare 
        - ricordiamo che:
            - un cifrario è deterministico se plaintext identici producono ciphertext identici
            - è probabilistico se quanto detto sopra non è vero dato che viene impiegato un numero casuale per ottenere cifrati distinti
    - Per eliminare questi punti deboli, il Cifrario deterministico deve essere “randomizzato”.
        - Se il **messaggio è più lungo del modulo**, si impiega tipicamente la **modalità CBC ed un IV casuale**.
        - Se il **messaggio è più corto del modulo**, ed è questo il caso di maggiore interesse (chiavi), il mittente, seguendo uno standard ben preciso indicatogli dal destinatario, **gli aggiunge in testa un padding contenente un numero a caso e poi cifra il tutto**.

16. RSA
    - cifratura     -> c = m^e mod n
        - _e_ ed _n_ formano la chiave pubblica
    - decifratura   -> m = c^d mod n = m^(e*d) mod n = m^1 mod n
        - _d_ ed _n_ forma la chiave privata
        - d inverso moltiplicativo di e -> e*d = 1
    - **è importante nascondere anche p e q oltre a d!**
        - altrimenti si riesce a calcolare prima _e_ e poi _d_
        - Non a caso dopo la generazione della coppia di chiavi solitamente vengono distrutti 
    - RSA è deterministico, per renderlo probabilistico si aggiunge del padding con lo standard OAEP
    - il tempo di esponenziazione modulare scala con il cubo della dimensione in bit della chiave (1024 -> 2048 -> 8 volte il tempo)
        - è desiderabile usare l'algoritmo più efficente possibile per questo calcolo
            - Repeated square and multiply per cifratura
            - per la decifratura, anche Algoritmo di Garner
        - per l’inviolabilità del Cifrario asimmetrico le chiavi devono essere molto lunghe, ma questo rende poi molto onerosi i calcoli di cifratura e di decifrazione; **bisogna quindi non esagerare mai nel dimensionamento delle chiavi**.  
    - RSA possiede la proprietà moltiplicativa
        - Consideriamo un messaggio m = m1 × m2.
        - La sua cifratura è: c = m^e mod n = ((m1^e mod n) × (m2^e mod n)) mod n
        - La cifratura del prodotto di due messaggi è uguale al prodotto dei due testi cifrati.
        - Analogamente, si ha che la decifrazione di un testo cifrato ottenuto moltiplicando due testi cifrati è uguale al prodotto dei due corrispondenti testi in chiaro
    - RSA è reversibile, posso scambiare il ruolo delle chiavi per effettuare delle firme
        - Nota: come per la chiave di sessione in un cifrario ibrido, l'hash ha un numero di bit molto minore del modulo n ed è quindi ritenuto insicuro trasformarla direttamente. Di nuovo si usa del paddding.

17. Attacco con proprietà moltiplicativa
    - si suppone che l’intruso possa richiedere al proprietario della chiave privata la decifrazione di qualsiasi messaggio di sua scelta, ad esclusione del messaggio c di suo interesse.
    - Un intruso che ha intercettato un cifrato c di cui vuole scoprire il plaintext m, può 
        - costruire un cifrato c' ottenuto come prodotto di c e qualcos'altro che gli pare: c' = c * r^e
        - far decifrare questo c'
        - annullare il termine in eccesso ed ottenere la decifratura che gli interessa 
    - Contromisure:
        - difficile però che l'intrusore riesca a convincere il proprietario a decifrare qualsiasi messaggio di sua scelta sopratutto se si ha identificazione
            - **NB**: se però una stessa coppia di chiave viene usata per cifrare e firmare uno potrebbe presentare il cifrato come messaggio da firmare, e la firma lo decifrerebbe a causa della reversibilità di RSA
        - se si usa padding (OAEP) la proprietà moltiplicativa non vale più

18. Cifrario ibrido
    - posso decidere una chiave di sessione e comunicarla al mio interlocutore cifrando la chiave di sessione con la sua chiave pubblica
    - chiaramente la sua chiave pubblica deve essere autentica e quindi devo avere il relativo certificato 
    - cifrare messaggi molto più piccoli della dimensione delle chiavi in RSA non è sicuro, per questo motivo, quando si cifrano chiavi simmetrica in un cifrario ibrido è importante adottare il padding (OAEP)

19. Firma digitale
    - Autentica un dato
        - ovvero, crea dunque un attendibile e verificabile **collegamento tra il dato e la persona/identità che l’ha creato**.
    - ha validità legale (quella con appendice)
    - due schemi:
        - con recupero  -> cifro con la chiave privata, suddivedendo in blocchi se necessario
        - con appendice -> hash del documento e poi firmo solo l'hash
    - molto più efficiente lo schema con appendice dato che devo firmare (cifrare con RSA) un solo blocco
    - la firma con recupero poi non ha validità legale dato che posso intercettare firme di vari documenti e relativi documenti e cominciare a fare sostituzione di blocchi mantenendo la validità della firma -> posso presentare come autentico un documento che ho modificato
        - la firma con recupero va anche bene se firmo un solo blocco -> questo è il caso di RSA

20. **Firma cieca**
    - possible grazie alla proprietà moltiplicativa di RSA
    - Siccome la firma è una cifratura, vale come prima che **la firma del prodotto di due messaggi è uguale al prodotto delle due firme**
    - Tale tecnica risulta particolarmente utile nella situazione in cui un utente X può **richiedere ad un’Autorità T di autenticargli un messaggio m, di cui, però, non vuole rivelare il contenuto**.
        - vedi voto elettronico in cui una terza parte mi firma il mio voto senza che sappia quale sia
    - un utente può quindi 
        - oscurare un messaggio moltiplicandolo con qulcosa che sa 'r' lui cifrato con la chiave pubblica della terza parte
        - presentare il prodotto a T per farselo firmare
            - la firma decifrerà r 
        - eliminare r ed ottenere il messaggio firmato senza che la terza parte abbia mai visto il contenuto
    - **NB**: siccome la terza parte firma, e quindi si assume una responsabilità, solitamente la terza parte impiega anche un protocollo di identificazione per capire chi è che mi sta chiedendo di firmare della roba. 
        - Questo elimina l'attacco con proprietà moltiplicativa
    - **A che cosa serve la firma cieca?**
        - in sostanza, ad autenticare i dati mantenendo però l'anonimato
        - La firma dell’autorità dice:
            - Questo messaggio è valido
            - Proviene da una persona autorizzata (perchè T ha identificato)
            - Non è stato modificato
        - Perché serve oscurare il messaggio nella fase di firma?
            - Se l’autorità vedesse il contenuto del voto quando lo firmi, saprebbe per chi voti
                - violazione della segretezza del voto.
            - Con la firma cieca l’autorità non sa cosa ha firmato
            - Tu rimuovi l’offuscamento e ottieni il voto firmato, pronto da inserire nell’urna
            - Chiunque controlla la firma sa che è un voto valido, **ma non sa chi è il proprietario del voto**
        - Conclusione: se firmassi io, tutti saprebbero per chi ho votato. La firma cieca mi permette di mantenere il mio anonimato

21. **Attacchi alla firma con RSA**
    - **attacco alla firma basato su proprietà moltiplicativa**
        - Supponiamo che l’intruso **possa richiedere al firmatario X di firmare qualsiasi messaggio di sua scelta, ad esclusione del messaggio m di suo interesse** (questo scenario è detto chosen message attack).
        - Per ottenere la firma di m è sufficiente che
            - l’intruso generi un numero a caso r
            - costruisca m1 = m × r
            - calcoli m2 = r^-1
            - chieda ad X di firmare entrambi i messaggi
            - e moltiplichi infine le firme di m1 e di m2 che gli vengono restituite.
                - Ricordiamo: **La firma del prodotto di due messaggi è uguale al prodotto delle due firme**
                - Abbiamo quindi che S(m1) * S(m2) = S(m1 * m2) = S(m * r * r^-1) = **S(m)**
                - **X ha firmato inconsapevolmente un messaggio arbitrario dell'attacante**

    - **attacco alla riservatezza facendo firmare un cifrato**
        - L’intruso sfrutta la proprietà moltiplicativa
            - oscura c in modo che il firmatore non si accorga che firmando decifra qualcosa,
            - se lo fa firmare da X (il che corrisponde alla decifrazione con la chiave privata),
            - elimina il numero a caso ed ottiene il testo in chiaro che non doveva poter conoscere.
        - Vale dunque la seguente regola: `chi impiega RSA per decifrare e per firmare deve utilizzare due differenti coppie di chiavi`
            - in questa maniera, firmare un cifrato non lo decifra dato che le coppie di chiavi sono diverse