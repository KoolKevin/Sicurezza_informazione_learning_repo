## TRASFORMAZIONI PER LA SICUREZZA
Le vulnerabilità del canale insicuro possono essere risolte da una contromisura in linea di principio molto semplice:
```
sorgente e destinazione, in perfetto accordo, attribuiscono alle informazioni comunicate sul canale insicuro una rappresentazione nota soltanto a loro ed atta ad impedire all’intruso di comprendere il significato della comunicazione (riservatezza) e/o di spacciare come inviati dalla sorgente messaggi che lui ha modificato (integrità) o forgiato (autenticità)
```

Sia nel caso di attacchi passivi che attivi, **i dati devono essere manipolati** (criptati nel caso di attacchi passivi, affiancati ad un qualcosa che ne garantisca l'integrità/autenticità nel caso di attacchi attivi). In altre parole, **i dati dalla sorgente alla destinazione devono essere trasformati**.

Le trasformazioni per la sicurezza possono essere implementate con:
- algoritmi: singola trasformazione
- protocolli: catena ben pensata di trasformazioni

Occorre, prevedere anche il caso che **il malintenzionato sia uno dei due corrispondenti**. Anticipiamo qui che per fronteggiare questa minaccia è necessario il coinvolgimento nel protocollo di una **terza parte fidata**, attribuendole il compito di intervenire nel corso (l’arbitro) o al termine (il giudice) dell'esecuzione del protocollo. 

Tutte le trasformazioni di sicurezza hanno come caratteristica comune una **codifica ridondante dei dati**
- o nella rappresentazione (uso più bit del caso senza trasformazione)
- o nel tempo (uso più tempo rispetto al caso senza trasformazione)

**crittologia, crittografia e crittanalisi**
La disciplina che studia gli algoritmi ed i protocolli da svolgere alle estremità di un canale insicuro è detta Crittologia.

Obiettivo fondamentale della Crittologia è l’individuazione di trasformazioni **efficienti, efficaci e sicure**; a tal fine la Crittologia è formata da due distinte e correlate discipline:
- la **Crittografia** individua le trasformazioni idonee a proteggere una o più proprietà critiche dell’informazione;
- la **Crittanalisi** ne valuta la robustezza, esaminando come, in quanto tempo e con quali risorse è possibile rompere le difese. 

**trasformazioni segrete e trasformazioni note con parametro segreto**
- algoritmo noto con parametro segreto (quello che considereremo d'ora in poi)
- algoritmo direttamente segreto (anche no)

vedi meglio dopo...

### I tre principi di difesa
Tutti gli algoritmi e i protocolli di sicurezza si basano su 3 principi, chiamati principi della difesa:
1. Impossibilità di sapere (Segretezza della trasformazione)
    - Descrizione: L'intruso non deve essere in grado di risalire al messaggio originale conoscendo il messaggio cifrato, senza conoscere la chiave.
    - Fondamento: La trasformazione del messaggio (algoritmo di cifratura) è progettata in modo tale da essere non reversibile senza chiave.
    - Obiettivo: Proteggere la riservatezza dell'informazione.
    - Concetto collegato: Cifratura segreta o simmetrica (ma vale anche per cifratura asimmetrica); uso di funzioni non invertibili senza chiave.
    - 📌 Esempio: Se intercetto un messaggio cifrato AES, non posso decifrarlo senza la chiave, anche se conosco l’algoritmo AES.

2. Impossibilità di indovinare (Resistenza alla probabilità)
    - Descrizione: L'intruso non deve riuscire a indovinare il messaggio originale basandosi su conoscenze statistiche, frequenze, o linguaggio naturale.
    - Fondamento: La cifratura deve "appiattire" ogni struttura del messaggio originale, rendendo i messaggi cifrati statisticamente uniformi.
    - Obiettivo: Evitare attacchi basati sull'analisi statistica.
    - Concetto collegato: Indistinguibilità semantica, distribuzione uniforme del messaggio cifrato.
    - 📌 Esempio: Un cifrario monoalfabetico (tipo Cesare) non rispetta questo principio perché si può analizzare la frequenza delle lettere. AES invece sì, perché produce output statisticamente casuali.

3. Impossibilità di dedurre (Resistenza computazionale)
    - Descrizione: Anche conoscendo il messaggio cifrato, non è computazionalmente possibile dedurre la chiave o il messaggio in chiaro in tempi ragionevoli.
    - Fondamento: Si basa su problemi computazionalmente difficili, come la fattorizzazione di numeri grandi (RSA) o il logaritmo discreto (Diffie-Hellman, ECC).
    - Obiettivo: Rendere impraticabile ogni tentativo di attacco brute-force o crittoanalisi, anche con risorse elevate.
    - Concetto collegato: Sicurezza computazionale, hard problems, complessità algoritmica.
    - 📌 Esempio: Rompere RSA a 2048 bit richiederebbe risorse computazionali enormi, al di là delle capacità pratiche attuali.

### Che proprietà interessano a noi?
La crittografia ci permette di ottenere: confidenzialità, integrità, autenticità (, non ripudio, identificazione), ma non availability. Il CIA triangle che consideriamo sostituisce la disponibilità con l'autenticità. 












### Trasformazioni per la riservatezza

**difende dall'attacco passivo di intercettazione!**

**Cifratura**: "la sorgente trasforma la rappresentazione originaria delle informazioni riservate in una rappresentazione che le renda
apparentemente incomprensibili; **la destinazione è l’unica a saper eseguire la trasformazione inversa**.


Una coppia di trasformazioni E, D è detta **cifrario**. 

Alcune osservazioni sull’uso di un cifrario:
- i due partecipanti non devono essere necessariamente on-line: in ricezione i messaggi cifrati possono, infatti, essere prima memorizzati e poi decifrati;
- uno stesso utente può essere sia sorgente, che destinazione: questo caso interessa chi vuole realizzare un proprio archivio di documenti riservati. 

La minaccia alla riservatezza è, come sappiamo, l’intercettazione. La difesa, **preventiva**, è mettere l’intruso I nella condizione di non riuscire a fare ciò che fa la destinazione B. A tal fine **la modalità di decifrazione deve essere tenuta segreta**, ma ciò non è sufficiente. L’intruso potrebbe, infatti, rimettere in chiaro il testo cifrato o, tanto peggio, individuare il segreto sulla trasformazione avvalendosi di conoscenze già in suo possesso (probabili motivi della comunicazione, probabile contenuto, ecc.) e di testi cifrati intercettati in precedenza. 

**Sicurezza perfetta**: dal cypher-text non si impara nulla.
- sicurezza perfetta non esiste; se il testo cifrato presenta ripetizioni (vedi linguaggio naturale) l'attaccante ottiene informazioni
- si parla solo di **sicurezza computazionale**: i calcoli per mettere in chiaro un testo cifrato devono essere computazionalemente impossibili (e.g. forza bruta con chiave da 512 bit)






### Trasformazioni per l'integrità

**difende dall'attacco attivo di modifica**

**Hashing**: “la sorgente affianca al messaggio un “riassunto” che ne rappresenti in modo univoco il contenuto e che funga da **attestato di integrità**; la destinazione calcola il riassunto del messaggio ricevuto e lo confronta con quello inviato dalla sorgente”.

La Crittografia impiega a tal fine **una particolare funzione, detta hash**, che comprime una stringa _m_ di lunghezza arbitraria in una stringa _h_ **di lunghezza piccola e prefissata**: h = H(m) 

#### Funzioni di Hash crittograficamente sicure
L’uscita di una funzione hash è detta riassunto o impronta (digest, fingerprint) del messaggio d’ingresso. Un’impronta di n bit (tipicamente nel range 128 ÷ 512) suddivide l’insieme di tutte le possibili stringhe d’ingresso in 2^n sottoinsiemi disgiunti, formati ciascuno da tutte e sole le stringhe (notare il plurale) che forniscono uguale h. 
- Due stringhe che hanno lo stesso hash sono dette essere in **collisione**.

Una funzione hash è detta crittograficamente sicura se:
- **il suo comportamento è apparentemente aleatorio**. Il modello a cui ci si fa riferimento è detto “oracolo casuale” e prevede che: fornendo in ingresso un
messaggio di cui non si conosce ancora l’impronta, **si riscontra sull’uscita, con uguale probabilità, uno qualsiasi dei 2^n valori possibili**.
- **NB**: il comportamento casuale delle funzioni di hash rende difficile trovare collisioni! Questo garantisce integrità!!

Gli algoritmi che approssimano tale comportamento impiegano un n “grande” per rendere estremamente improbabile che messaggi diversi abbiano la stessa impronta e rispettano la seguente regola:
```
(resistenza alle collisioni): “l’individuazione di due messaggi con la stessa impronta è un calcolo difficile”. 
```

In generale una funzione hash crittograficamente sicura ha queste proprietà:
- comportamento aleatorio (ogni hash ha la stessa probabilità di essere scelto)
    - è difficile dato un hash trovare un altro messaggio che mi produca lo stesso hash
    - questo mi da resistenza alle collisioni
- non invertibilie, è computazionalmente impossibile risalire ad *m* a partire da *H(m)*
    - in questo modo dato un hash non riesco a recuperare il messaggio originale 

**Basta l'hashing per garantire l'integrità?** trasmetto m||H(m)
- NOOO!!
- cosi come l'intrusore può modificare il messaggio, esso può modificare anche l'hash
- in particolare può modificare il messaggio, ricalcolarne l'hash e inviare il tutto al destinatario che non si accorgerà di niente

Come fare? 
- Una possibilità è trasmettere l'hash su un canale dedicato
    - l'intrusore non è in grado di trovare un m' che collide con m (resistenza alle collisioni)
    - perchè non trasmetto direttamente m sul canale dedicato?
        - Usare il canale dedicato ha un costo per messaggi di dimensione arbitraria (pensa a trasmettero un programma di 1MiB parlando al telefono)
- di fatto la soluzione adottata è usare un MAC

**Come possiamo combinare le trasformazioni E ed H?**
voglio garantire integrità e riservatezza

- E(m) || H(m) non va bene dato che non garantisce riservatezza
    - hash del messaggio in chiaro
    - un attaccante può provare tutti i messaggi finche non ne trova uno che ha un hash corrispondente
        - attacco di forza bruta per scoprire m partendo da H(m) 
        - funziona bene se ho poche possibilità ( |spazio dei messaggi| << |spazio degli hash| )
    - ancora più facile se l'attaccante ha informazioni di contesto e il dominio del messaggio è limitato (e.g PIN a 4 cifre)

- Versione corretta:
    1. p = m||H(m)              -> A forma un nuovo testo in chiaro p “concatenando” m e H(m)
    2. c = E(p)                 -> A cifra p ed invia a B il risultato c
    3. c -> c*                  -> attraversando il canale un canale può modificare c 
    4. p* = D(c*) = m*, H*(m)   -> B decifra c*, ottiene p* e separa le due componenti
    5. H(m*) =? H*(m)           -> B calcola l’impronta di m* e la confronta con l’impronta ricevuta: se sono uguali il messaggio è giudicato integro
    - ricorda quello che succede in SSL solo che al posto dell'hash del messaggio e basta vi è un MAC
    - L’intruso può solo modificare a caso il dato cifrato, ma ne danneggerebbe l’integrità e verrebbe solo sgamato.




### Trasformazioni per autenticazione/autenticità

**difende dall'attacco attivo di fabbricazione (e anche dalla modifica)**

vogliamo verficare che l'origine del messaggio sia effettivamente quella del mittente. Quando ricevo un messaggio come faccio a capire che la provenienza del messaggio è quella giusta? 

**la sorgente** aggiunge al documento informazioni **non imitabili** atte ad attestare chi l’ha predisposto (firma/prova di autenticità); la destinazione verifica che il documento ricevuto sia stato originato proprio da chi dichiara di averlo fatto”
- nessuno può produrre le informazioni contenute nell'attestato se non la sorgente legittima in quanto quest'ultima utilizza un segreto

Nuove trasformazioni:
- S(ign)
    - trasformazione segreta, la sa fare solo il mittente
    - nessuno può imitare S(m) del mittente
- V(erify)
    - la sanno fare tutti (non stiamo proteggendo la riservatezza)
    - produce due output: si/no sull'autenticità del messaggio; il messaggio m


Spesso è necessario trasmettere in chiaro il messaggio m. Per consentire a B di **accertare contemporaneamente l’integrità e l’origine del messaggio ricevuto**, la Crittografia ha messo a punto due differenti soluzioni:

#### Importante
Per garantire autenticità un prerequisito necessario (ma non sufficiente) è l'integrità. Infatti se il messaggio viene modificato quest'ultimo sicuramente non è più autentico. 

**Solamente dopo aver ricevuto un messaggio integro ci si può preoccupare di verificare la sua origine!** Altrimenti, chiunque potrebbe averlo modificato nel canale insicuro. L'origine di un messaggio viene verificata aggiungedo gli attestati di autenticità nominati sopra. 

In seguito i due schemi di attestato principali.



**Schema Firma digitale**
![alt text](immagini/firma_digitale.png)

il messaggio in chiaro viene affiancato con l'autenticazione del suo hash (firma):
- La firma corrisponde ad una cifratura con chiave privata di un cifrario asimmetrico
- La destinazione, con il meccanismo V, noto a tutti (decifrazione con chiave pubblica), mette in chiaro il l'hash cifrato e lo confronta con l'hash del messaggio ricevuto. In questa maniera rileva ogni alterazione apportata a *m* e/o a *c* e se i due hash combaciano, viene anche dimostrata l'autenticità del messaggio in quanto solamente il possessore della chiave privata può produrre un messaggio decifrabile con la chiave pubblica. 
- S(ign) è un algoritmo segreto! Per un intruso è impossibile modificare la firma *c* e/o *m* mantenendo l'integrità del messagggio e l'identità della sorgente (gli hash non combacerebbero)

Vantaggi:
- non ripudiabile in quanto solo S conosce il suo segreto con cui è stato firmato il messaggio 
- non c'è uno scambio di chiavi 

Svantaggi:
- devo impiegare crittografia asimmetrica che è meno efficente

**Schema con PARAMETRO segreto (MAC)**
![alt text](immagini/segreto_concordato.png)

La trasformazione S non è più una cifratura con chiave privata, ma si trasforma in H(m||S)

I due corrispondenti A e B concordano in segreto un dato S. Per consentire la verifica dell’integrità e dell’origine di un m, A e B eseguono il seguente protocollo:
1. A:        calcola H(m||s)
2. A → B:    c = m||H(m||s)
3. B:        riceve c* ed ottiene m*, H*(m||s)
4. B:        calcola H(m*||s)
5. Se e solo se H(m*||s)=H*(m||s), B considera m integro ed originato da A.

ripudiabile: sia bob che alice conoscono il segreto e quindi possono imbrogliare 
- la generazione di m ed il calcolo di H(m||s) possono essere fatti anche da B all’insaputa di A (entrambi conoscono il segreto) ed un’eventuale disputa tra i due non può essere risolta.
- La soluzione è quindi utile solo se i due corrispondenti **si fidano uno dell’altro**. 


A questo punto, se al posto di mandare il messaggio in chiaro, si manda E(m), e si usa E(m) al posto di m per costruire l'attestato di autenticità, **il canale insicuro viene reso sicuro!** 









## Trasformazioni per identificazione:
Riconoscimento dell’identità di chi sta effettuando una certa operazione o accede a delle risorse. 

Regole per l’identificazione: **l’utente che vuole farsi riconoscere deve fornire informazioni non imitabili**, atte ad individuarlo univocamente in quel preciso istante di tempo.

Esistono applicazioni che ammettono l’anonimato (informativa su servizi pubblici, posta elettronica) ed altre che lo presuppongono (voto elettronico, moneta elettronica). Nella maggioranza dei casi è però richiesto che l’erogatore di un servizio conosca l’identità di chi ne usufruisce (tipicamente per **delimitare** (autorizzazione) e per fatturare l’uso che ne viene fatto) e che il fruitore sia certo dell’identità dell’erogatore (per non cadere ingenuamente in una truffa).
- fondamentale per il controllo dell'accesso
- fondamentale per autenticazione

Al contrario delle trasformazioni precedenti, il processo d’identificazione è **necessariamente real-time**: si deve, infatti, riferire ad un ben preciso istante di tempo ed entrambi gli interessati (verificando e verificatore) devono quindi essere on line (sincrono).
- vale per un preciso istante di tempo
- la validità prolungata si ottiene affiancando al meccanismo di identificazione un altro (di autenticazione)

L’identificazione può essere:
- Singola: solo il client si deve identificare al server;
- Mutua: sia il client che il server devono identificarsi mutuamente.

**Come si identifica?**
“tramite il solo scambio di messaggi, l’Entità (macchina o uomo) che vuole farsi riconoscere deve fornire **informazioni non imitabili**, atte ad individuarla univocamente in quel preciso istante di tempo, e l’Entità che effettua il riconoscimento deve potersi convincere della loro **genuinità**”.

Tre sono in generale le dimostrazioni d’identità che una persona può fornire:
1. La **conoscenza** di un’informazione che ha in precedenza concordata (password/PIN)
2. Il **possesso** di un oggetto datogli in consegna (token, una smart card, ecc)
3. La conformità di una **misura biometrica** raccolta in quel momento con una misura registrata in precedenza;
    - oggetto di misura è una caratteristica individuale, o fisiologica (l’impronta digitale, la forma della mano, ...)

**NB**: L’identificazione di una macchina si basa ovviamente solo sulla conoscenza.

L'identificazione è il classico esempio in cui non basta una singola trasformazione. Due sono i momenti topici:
![alt text](immagini/identificazione.png)

- fase iniziale di registrazione
    - l’identificando I ed il verificatore V concordano e memorizzano **rispettivamente** il dato segreto S con cui I si farà riconoscere ed il termine di
    paragone T=f(S) che consentirà a V di accertare che I conosce S.
- fase a regime in cui si avvia una sessione di riconoscimento tramite un protocollo di identificazione a sua volta con 3 fasi:
    1. l’identificando dichiara chi è;
    2. il verificatore gli chiede di dimostrarlo;
    3. l’identificando fornisce la **prova di identità**; se è quella concordata, il verificatore lo identifica. 
    - 3 fasi -> 3 traformazioni
    - trasformazione di dimostrazione non imitabile 

**NB**: il tutto si base sul fatto che l'identificando mantenga il suo segreto segreto  