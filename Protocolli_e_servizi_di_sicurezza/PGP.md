


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