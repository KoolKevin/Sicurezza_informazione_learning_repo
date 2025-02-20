Abbiamo visto che sia nel caso di attacchi passivi che attivi i dati devono essere manipolati (criptati nel caso di attacchi passivi, affiancati ad un qualcosa che ne garantisca l'integrità nel caso di attacchi attivi). In altre parole, **i dati dalla sorgente alla destinazione devono essere trasformati**.

Le trasformazioni per la sicurezza possono essere implementate con:
- algoritmi: singola trasformazione
- protocolli: catena ben pensata di trasformazioni

A volte nei protocolli può essere coinvolta una **terza parte fidata** che fa da arbitro
- spesso sorgente/destinazione non sono fidate

Tutte le trasformazioni di sicurezza hanno come caratteristica comune una **codifica ridondante dei dati**
- o nella rappresentazione (uso più bit del caso senza trasformazione)
- o nel tempo (uso più tempo rispetto al caso senza trasformazione)


### Trasformazioni segrete e note
- algoritmo noto con parametro segreto (quello che considereremo d'ora in poi)
- algoritmo direttamente segreto (anche no)

### Trasformazioni per la riservatezza
cifratura

sicurezza perfetta non esiste; si parla solo di sicurezza computazionale


### Trasformazioni per l'integrità
non esiste una trasformazione preventiva, esiste una trasformazioni che rileva modifiche del contenuto del messaggio.
- generazione di un attestato di integrità prodotto tramite hashing crittografico. Ha queste due proprietà
    - comportamento aleatorio (ogni hash ha la stessa probabilità di essere scelto)
    - resistenza alle collisioni
    - non invertibilie, è computazionalmente impossibile risalire ad *m* a partire da *H(m)*

basta l'hashing per garantire l'integrità? NOOO!!
- cosi come l'intrusore può modificare il messaggio, esso può modificare anche l'hash
- in particolare può modificare il messaggi e ricalcolare l'hash

Come fare? Bisogna trasmettere l'hash su un canale dedicato
- l'intrusore non è in grado di trovare un m' privo che collide con m
- perchè non trasmetto direttamente m sul canale dedicato? Usare il canale dedicato ha un costo per messaggi di dimensione arbitraria


### Come possiamo combinare le trasformazioni E ed H?
voglio garantire integrità e riservatezza

- E(m) || H(m) non va bene dato che non garantisce integrità
    - un attaccante può provare tutti i messaggi finche non ne trova uno che ha un hash corrispondente
    - ancora più facile per messaggi di lunghezza arbitraria se l'attaccante ha informazioni di contesto

- 2^ esempio corretto
    - ricorda quello che succede in SSL


### Trasformazioni per autenticazione
vogliamo verficare che l'origine del messaggio sia effettivamente quella del mittente

abbiamo bisogno di un attestato di autenticità. 
- nessuno può produrre le informazioni contenute nell'attestato se non la sorgente legittima

Nuove trasformazioni
- S(ign)
    - segreta, la sa fare solo il mittente
    - nessuno può imitare S(m) del mittente
    
- V(erify)
    - la sanno fare tutti (non stiamo proteggendo la riservatezza)
    - produce due output: si/no; il messaggio m

**Firma digitale**
firmo l'hash del messaggio che ha dimensione ridotta
- la destinazione non deve fare V(m)
- non devo firmare un messaggio lungo