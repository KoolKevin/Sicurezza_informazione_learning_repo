## PKI
Formata da:
- Certification authority
- Registration authority: identifica gli utenti con delle credenziali
- DB (directory): registro ad accesso pubblico in cui vengono mantenuti i certificati

![alt text](img/PKI_entities.png)



### Richiesta ed emissione del certificato (attività 1, 2, 3, 4) 
![alt text](img/CSR.png)

Supponiamo che X voglia farsi certificare una chiave di firma. 

Come prima cosa **genera un autocertificato della sua PX** (Certificate Signing Request), **per dimostrare di essere il proprietario di SX**.
- **NB**:notare che la generazione della coppia di chiavi avviene **lato utente**
    - la PKI non vede mai la chiave privata -> l'utente sarà **l'unico possessore della chiave privata**
    - questo è importante per garantire la proprietà di **non ripudio** quando questa coppia di chiavi verrà poi utilizzata
    - la CA potrà verificare che l'utente possegga la chiave privata, verificando la firma presente nel CSR con la chiave pubblica presentata al suo interno
    - in questa maniera e con le altre informazioni identificative presentate alla RA, l'utente viene identificato 

A questo punto X fornisce a RA i suoi dati (tramite un’e-mail o una form, quando è richiesto un basso livello di sicurezza, presentandosi invece con documenti autenticati, quando il livello di sicurezza deve essere alto). 

RA verifica la firma e, in caso positivo, passa i dati a CA. CA organizza secondo lo standard i dati raccolti da RA, aggiunge quelli di sua competenza, calcola l’hash del tutto e lo firma. **Costruendo in questo modo il certificato**.

Una prima copia del certificato è consegnata a X ed una seconda al repository DB accessibile dalla rete: si noti che un database di questo tipo non richiede né controllo degli accessi, né alcuna particolare forma di protezione dei dati in memoria (contiene solo informazioni pubbliche read-only)

L’intervallo di validità del certificato è tipicamente è un anno: prima della scadenza, l’utente ne deve richiedere il rinnovo. 



### Ripudio della chiave pubblica e revoca del suo certificato (attività 5, 6, 7, 8)
Su tutti i meccanismi basati su un segreto incombe il pericolo che qualcuno riesca o ad indovinarlo o a rubarlo.

La difesa che gli utenti devono adottare è espressa nella seguente regola.

R29: `quando uno ha il sospetto che la sua chiave privata sia stata violata, deve rinunciare immediatamente ad impiegarla, generarsene una nuova e farsi certificare la corrispondente chiave pubblica`


Il ripudio è l’azione con cui un utente, prima della scadenza, **spezza l’associazione chiave pubblica/proprietario** sancita da un certificato. E’ indispensabile che questo evento sia tempestivamente notificato a tutti gli altri utenti.
- La raccolta delle notifiche di ripudio è fatta dalla RA
- Alla CA compete la successiva azione di revoca del certificato.
    - La CA revoca i certificati anche spontaneamente quando scadono i periodi di validità

**come si distribuisce agli utenti l'informazione sullo stato di revoca di un certificato?**

Per rendere noto a tutti gli altri utenti che una chiave non deve essere più usata, **CA mantiene on line una lista autenticata dei certificati revocati (CRL).**
- Chiunque è dunque così in grado di sapere se è ancora valida la chiave pubblica di chiunque altro. 
- Fino a poco tempo fa occorreva reperire la CRL (dove è indicato nel certificato), scaricarla e poi interrogarla; ora alcune CA supportano query dirette.

## Revoca di un certificato
...

i certificati hanno una scadenza imposta dalla CA al momento del suo rilascio
- tuttavia, i certificati hanno anche uno stato di revoca (i.e. relativa chiave privata compromessa)

**modello pull/push**
modelli con cui si distribuisce agli utenti l'informazione sullo stato di revoca di un certificato
- modello pull è quello effettivamente utilizzato 
- modello push più complicato (complessità di un sistema di gestione degli eventi)

### CRL
una lista di revoca è una struttura dati analaga ad un certificato in quanto firmata dalla CA
- contiene l'indicazione dei certificati revocati
- tutto quello firmato prima della data di revoca deve essere considerato valido, e viceversa
- this update date e next update date identificano un intervallo temporale di validità della lista di revoca
- CRL extensions: informazioni aggiuntive utili a livello applicativo, non necessariamente bisogna garantirne autenticità e integrità
    - e.g. ragione di revoca

vantaggio: CRL funziona anche offline data la validità temporale
svantaggio: non garantiscono la freschezza delle informazioni! Se avviene una revoca in mezzo all'intervallo di tempo, non me ne accorgo fino al prossimo
- limite: non garantisce mai freschezza in tempo reale per quanto io diminuisca l'intervallo di validità


...


Nel caso di CRL a sottoliste, posso mettere in una estensione del certificato la sottolista in cui sarà contenuta l'informazione sulla sua informazione di revoca. 

il campo estensioni della CRL mi dice se una determinata CRL è una delta CRL, base CRL, sotto lista, ecc...

### OCSP
protocollo client-server che data una richiesta mi risponde con lo stato di revoca di un certificato (tutto firmato)

il server può avvalersi di liste di revoca ...
- non cambia nulla dal punto di vista della freschezza 

... come non
- può (non è detto) supportare una verifica in tempo reale (massima freschezza)

timeliness == tempestività della notifica di revoca 
















### La Gerarchia delle autorità di certificazione

Sicuramente non può bastare un unico certificatore per tutti gli utenti del mondo. Lo standard X.509 prevede una **gerarchia di Autorità**.

Alla base della piramide si trovano le Autorità che certificano ciascuna le chiavi di un certo numero di utenti.
- un singolo **dominio di fiducia**

Al di sopra si trovano Autorità che certificano le chiavi di Autorità sottostanti e spesso anche di Autorità dello stesso livello.

Nel punto più alto si trovano le Autorità radice, che si autocertificano, essendoci le condizioni per considerarle fidate.

Risalendo la gerarchia e controllando la catena dei certificati, un utente di una CAi può verificare il certificato di un altro, anche se è stato emesso da una CAj diversa dalla sua. Una seconda possibilità è che ognuno metta a disposizione degli altri la catena di certificati da usare per avere fiducia sulla autenticità della sua chiave










### Domini di fiducia
per recuperare certificati di altre CA devo recuperare il cammino dei certificati





quando un utente si registri presso una CA
- sicuramente nel certificato che riceve è contenuta la chiave pubblica della CA in cui si è registrato
- inoltre, ottiene anche le chiavi pubbliche di altre CA delle quali la CA presso cui si è registrato
- in particolare, assumiamo che un utente oltre alla chiave pubblica della sua CA abbia anche la chiave pubblica della ROOT CA del suo dominio
    - in questa maniera, un utente può cercare certificati di altre CA partendo dalla radice

**Cross certificate**: certificato che contiene chiavi pubbliche di altre CA
- i cross-certficate seguono le stesse regole dei normali certificati (integrità, autenticità, validità temporale e stato di revoca)
    - AuthorityRevocationList al posto di CRL
    - una CA emette una CRL per i user-certificates, ed una ARL per i cross-certificate


### Conclusioni
alla base di un qualsiasi cifrario asimmetrico vi è per forza dietro una PKI
- fondamentale, altrimenti le chiavi pubbliche non sono dotate di autenticità
- in realtà, esistono anche altre tecnologie alternative che permettono di fornire di autenticità le chiavi pubbliche

Quand'è che un prodotto commerciale si può considerare una infrastruttura a chiave pubblica?
- fa tutte le cose che ci sono in slide 35



Il certificato (l'intera struttura dati) è salvata all'interno del DB. Chi fa da ente certificatore si assume le responsabilità civili e penali sulla distribuzione e gestione delle chiavi: ciò che esso distribuisce ha validità legale.





















**DH varianti con autenticità**
piccola estensione di Diffie-Hellman in cui aggiungo un PRNG che mi permette di cambiare dinamicamente la chiave
- il problema riguardo l'autenticazione dell'Y però rimane
- la minaccia è che si concordi una chiave segreta con un Man in the middle

Per risolvere posso usare una PKI
- i certificati emessi contengono i parametri DH (p, g, Y) di cui voglio garantire l'autenticità al posto della chiave pubblica del mio corrispondente
    - questo meccanismo evita che alice concordi con un MIM la stessa chiave, tuttavia siccome manca identificazione c'è il rischio che io mandi dei campioni di testo cifrato ad una persona diversa da BOB (il certificato è pubblico)
    - poco male siccome l'attaccante non può ottenere la stessa chiave con cui decifrare i dati siccome non ha la chiave segreta (X_bob) 
   
- oppure, ephemeral DH
    - **NB**: il segreto X viene ricalcolato ogni volta e di conseguenza cambia anche la chiave DH pubblica Y
    - ho PERFECT-FORWARD-SECRECY ma sono più pesante

Fixed DH: “Ti mando un biglietto con la mia chiave DH stampata sopra, firmato da un notaio (la CA)”.

Ephemeral DH: “Ti mando una chiave DH generata al volo e ci metto la mia firma elettronica; allego il mio certificato per dimostrarti che quella firma è mia”.

in sostanza fixed autentica la chiave DH pubblica inserendola direttamente in un certificato, mentre ephemeral autentica la chiave pubblica DH firmandola con una sua chiave privata e condividendo un certificato che autentica la relativa chiave pubblica

...
Per evitare di avere sempre lo stesso segreto condiviso, anche se il pre_master_secret è statico, si aggiunge entropia nuova ad ogni sessione con:
- R_C: nonce (numero casuale) generato dal client
- R_S: nonce generato dal server
- Questi due valori sono scambiati in chiaro

**NOTA**: è importante distinguere il requisito di autenticazione da quelle di identificazione
- il primo richiede che i messaggi che arrivano siano effettivamente appartenenti a Bob
- il secondo mi richiede che i messaggi che arrivano siano stati mandati proprio da Bob 

