TimeStampingServer

Kerberos, TSS e PGP sono servizi di livello applicativo e quindi non sono trasparenti all'utente (sono programmi da scaricare esplicitamente)


Firma digitale è un termine overloaded
- firma come servizio (con le sue proprietà: deve essere associata ad un solo documento, deve poter essere verificata da tutti, ...)
- firma come algoritmo implementato con cifrario (ad esempio con RSA)

La firma dipende da una marca temporale
- per validità
- o quantaltro (pensa a momenti di invio di un messaggio per una graduatoria)


Come al solito non dobbiamo studiare per filo e per segno i protocolli, ma dobbiamo capirne la logica e generalizzare le tecniche usate al suo interno per i nostri scopi e capire a cosa servono quest'ultime

### che proprietà deve avere un servizio di marcatura temporale? (questa è la parte utile che possiamo generalizzare)





### PGP
supporta sia cifratura simmetrica che asimmetrica

**OSS**: per la cifratura asimmetrica, PGP prevede un **portachiavi**, ovvero delle strutture dati che salvano la chiave privata dell'utente e le chiavi pubbliche degli altri utenti con cui si interagisce.


...

interessante notare le posizione inverse delle trasformazioni di compressione/decompressione in autenticazione e riservatezza

PGP supporta firma multipla ma anche firma a cipolla

**Formato dei messaggi pgp**




**chi mi da la garanzia che le chiavi pubbliche siano effettivamente del destinatario dichiarato?**
non c'è nessun certificatore


web of trust


ci chiariamo le idee guardando i portachiavi

PGP si basa sull'idea che le chiavi pubbliche possano essere distribuite in vari modi
- direttamente
- indirettamente
    - pubblicata su un sito
    - dentro ad un registro
    - ...

owner trust
- livello di fiducia che assegno a quella chiave
- gestita dall'utente

signature
- firma di chi mi ha fornito la chiave pubblica (e che ne certifica l'autenticità)
- potrebbe essere anche la firma di una CA, ma non necessariamente

signature trust
- fiducia che associo alla firma
- corrisponde all'owner trust della chiave di chiave pubblica associata a chi mi sta certificando la chiave

key legitimacy
- legittimitò finale della chiave viene determinata con una media pesata delle signature trust



questo sistema non ha validità legale, nessuno si assume responsabilità

la revoca ha una tempestività ridicola


**importante da sapere**: modello di fiducia del PGP come esempio di sistema decentralizzato
- owner trust: fiducia soggettiva
- fiducia verso intermediari

modello alla base dei trust management systems