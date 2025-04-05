abbiamo visto con i cifrari simmetrici come le chiavi si:
- generano -> prng
- memorizzano -> fs_sicuro
- distribuiscono -> KDC e D.H.

i cifrari simmetrici abbiamo visto essere impiegati più per la riservatezza che per l'autenticazione. I cifrari asimmetrici il contrario

### Chiavi pubbliche 
I meccanismi asimmetrici per la riservatezza e per l’autenticazione introducono in Crittografia tre aspetti di importanza fondamentale dal punto di vista della gestione delle chiavi:
1. Una volta scelto un Cifrario asimmetrico, solamente la chiave privata consente ad ogni utente di decifrare i messaggi riservati che chiunque altro gli invia cifrati con la corrispondente chiave pubblica;
    - chi non conosce la chiave privata, non riesce a decifrare.

2. Una volta scelto uno Schema di firma, una sola chiave, quella privata, consente ad ogni utente di firmare messaggi di cui chiunque altro potrà verificare l’autenticità usando la corrispondente chiave pubblica;
    - chi non conosce la chiave privata non riesce ad imitare la firma del suo proprietario.

3. La chiave pubblica di un utente può essere comunicata a tutti gli altri senza alcuna forma di riservatezza
    - questo è d'altronde il grosso vantaggio della crittografia asimmetrica, non bisogna preoccuparsi di come comunicare le chiavi in maniera sicura

Aspetti a carico di ogni utente U di un meccanismo asimmetrico:
- Disponibilità della chiave pubblica PU: chiunque ne ha bisogno deve poterla facilmente reperire; a tal fine possono essere impiegati
    - o un incontro diretto tra gli interessati,
    - o un’e-mail inviata dal proprietario ai suoi potenziali corrispondenti,
    - o un file che questi ultimi possono scaricare dal sito del proprietario,
    - o, ed è questa la soluzione migliore, **un record che chiunque può ottenere da un database**.
- Segretezza della chiave privata SU: solo il proprietario deve poterne disporre; a tal fine deve conservarla cifrata (con la tecnica PBE) in una memoria elettronica con controllo d’accesso (hard disk, o floppy, o, molto meglio, una smart card).
    - analogo ai segreti del caso simmetrico












### Autenticità della chiave pubblicax  
**IMPORTANTISSIMO**: La conoscenza di PU è una condizione necessaria, ma non sufficiente, per proteggere la riservatezza di un messaggio da inviare a U o per attribuire a U la paternità di un documento firmato con SU.
- Chi usa PU ad un’estremità del canale è, infatti, ben consapevole che all’altra estremità è stata o dovrà essere impiegata la chiave SU, ma **non ha alcuna garanzia che la coppia di numeri SU, PU sia proprio dell’utente di nome U** e non di un impostore che vuole spacciarsi per U. 
- In altre parole, l'utente che usa PU **non ha garanzie sulla sua autenticità**
- In queste condizioni il pericolo è grosso: si può, infatti, o inviare ad un intruso informazioni riservate destinate solo ad U, o accettare come originate da U informazioni false predisposte dall’intruso. 

**Man-in-the-Middle attack**
Consideriamo il caso di un Cifrario asimmetrico e supponiamo che due utenti A e B abbiano inserito in un database DB le loro chiavi pubbliche PA e PB. L’intruso I che vuole violare la riservatezza delle loro comunicazioni deve fare solo tre cose, neanche troppo difficili:
1. sostituisce la sua chiave pubblica PI al posto di quelle di A e di B
    - o modificando il contenuto di DB
    - o dirottando tutte le richieste di accesso sul suo calcolatore
2. quando A invia un c = E_PI(m) a B, lo intercetta, lo decifra, lo cifra con PB (pubblica) e lo inoltra a B;
3. quando B risponde ad A ripete lo stesso attacco attivo, completando così la violazione della riservatezza dei messaggi tra A e B. 

Discorso analogo per la firma digitale:
- dopo aver sostituito la chiave pubblica di B con la mia PI
- posso inviare un messaggio firmato con PI spacciandomi per B
- A verificherà la firma pensando che la paternità del messaggio sia di B e non si accorgerà di niente

La contromisura è comunicare chiavi pubbliche di cui è possibile verificare l’origine (autenticità)
- autenticità concetto non chiarissimo, una definizione che mi piace è: "associare una identità ad un dato, in questo caso la chiave pubblica" 

```R28: ”prima d’impiegare una chiave pubblica bisogna o essere personalmente certi dell’identità del suo proprietario o disporre di un documento (certificato) che l’attesta, redatto da un’Entità di cui ci si fida”.```






## Certificati
Due che si conoscono non hanno particolari problemi ad autenticare le loro chiavi pubbliche. 
- Una soluzione ovvia è che se le scambino durante un incontro diretto.
- Una seconda soluzione è che se le invino per posta ed usino poi una telefonata per confermarne a voce il valore. 

**In generale, però, chi deve usare una chiave pubblica non conosce il suo proprietario**. In questo caso la certificazione della associazione chiave-proprietario(identità) deve essere:
- fatta da qualcuno di cui ci si può fidare che conosoco,
- comunicata con un messaggio di cui sia possibile verificare integrità ed origine.

Il messaggio che attesta l’identità del proprietario di una chiave pubblica è detto **certificato**. 
- L’Ente T che produce certificati è detto terza parte fidata.
- La firma digitale di T apposta sul certificato tramite un algoritmo asimmetrico è l’accorgimento efficiente che consente a tutti gli interessati di verificare l’integrità e l’origine dell’attestato.
- **NB**: Il presupposto è che tutti abbiano una copia autentica di PT (magari preinstallata nel browser, PC, ...)
    - altrimenti non si saprebbe come verificare la firma di T.

**Come è fatto un certificato?**
Nella sua forma più ridotta, il certificato della chiave pubblica di un generico utente X da parte di T è formato da:
- un testo in chiaro m che contiene
    - l'identità X del proprietario della chiave pubblica
    - la chiave pubblica associata 
    - svariati altri metadati
- affiancato dal hash firmato con ST
- cert(PX, T) = m||S_ST(H(m))

**OSS**: schema di un tipico messaggio firmato

Oltre a questo, un certificato contiene anche molti altri campi:
- metadati identificativi sul possessore del certificato
- metadati sull'utilizzo della chiave pubblica
    - algoritmo asimmetrico da utilizzare
    - periodo di validità
- metadati su T
    - dati identificativi
    - algoritmi di hash e firma utilizzati
- estensioni utilizzate per varie cose (vedi dopo)

Una volta predisposto il certificato di X, T può consegnarlo
- ad un altro utente Y che gli ha direttamente richiesto l’autenticazione di PX,
- allo stesso utente X, che provvederà poi ad inoltrarlo a chi glielo richiederà,
- ad un database a libero accesso (directory), da cui chiunque potrà estrarlo.


Occorre anche decidere chi può ricoprire il ruolo di T; due sono i modelli in uso.
1. T è un Ente ufficialmente riconosciuto come fidato ed istituzionalmente incaricato di emettere certificati.
2. T è un qualsiasi utente che dichiara spontaneamente di conoscerne bene un altro; chi si fida di lui, si fida anche della sua dichiarazione.

La prima soluzione, detta della **Autorità di certificazione (o CA)**, è formalizzata nello standard ISO X.509 ed è impiegata sia in applicazioni sicure, sia nelle reti sicure; la seconda soluzione, propugnata dal **PGP**, è il contesto preferito da tutti quelli che non vogliono un Grande Fratello46 sopra di loro. 








## Infrastruttura a chiave pubblica (PKI)
basato su autorità di certificazione

come faccio io utente che genero una coppia di chiavi ad autenticarle? L'unico modo in questo modello a infrastruttura a chiave pubblica è inserire un terzo ente che fà da autorità fidata che fa da garante per questa autenticazione: **l'autorità di certificazione!**

L'autorità di certificazione:
- autentica le chiavi, firmandole con la propria chiave privata
- e le distribuisce all'interno di una struttura dati chiamata **certificato** (non come stringa di bit)
- gli utenti, che hanno già la mia chiave pubblica e che si fidano di me, possono verificare la mia firma quando ricevono un certificato



slide 8: quello che c'è da capire è che il fatto di ricevere un messaggio (integro) firmato con una chiave privata, mi da solo la certezza che il mittente di quel messaggio abbia la chiave privata e non che sia effettivamente X. 
- qua manca l'identificazione del mittente, ho solo la prova di possesso
- quando l'ente fidato verifica sia prova di possesso, che identificazione, allora rilascia il certificato (che ha la forma mostrata nella slide)
- anche i certificati restituiti dalla CA sono firmati e resi integri, per cui anche i CA distribuiscono una chiave pubblica

Lo standard X.509 definisce le informazioni contenute in un certificato. I certificati sono autentici ed integri dato che sono firmati dalla CA; se quindi ottengo un certificato che non passa la verifica di autenticazione, allora non userò mai la chiave pubblica al suo interno.

### Entità della PKI
- Certification authority
- Registration authority: identifica gli utenti con delle credenziali
- DB: registro ad accesso pubblico in cui vengono mantenuti i certificati

...


chi generà la coppia di chiavi?
come viene trasmessa la chiave pubblica alla CA?

...

schema centralizzato non garantisce il non ripudio dato che genera le chiavi a bordo!
- per questo motivo viene utilizzato uno schema a tre parti

**schema a tre parti**
...

terza alternativa è un aggravio inutile

La POP serve ad ottenere non ripudiabilità!








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















**DH varianti con autenticità**
piccola estensione di Diffie-Hellman in cui aggiungo un PRNG che mi permette di cambiare dinamicamente la chiave
- il problema riguardo l'autenticazione dell'Y però rimane
- la minaccia è che si concordi una chiave segreta con un Man in the middle

Per risolvere posso usare una PKI
- i certificati emessi contengono i parametri DH (p, g, Y) di cui voglio garantire l'autenticità al posto della chiave pubblica del mio corrispondente
    - questo meccanismo evita che alice concordi con un MIM la stessa chiave, tuttavia siccome manca identificazione c'è il rischio che io mandi dei campioni di testo cifrato ad una persona diversa da BOB (il certificato è pubblico)
    - poco male siccome l'attaccante non può ottenere la stessa chiave con cui decifrare i dati siccome non ha la chiave segreta (X_bob) 
- oppure, ephemeral DH

...
Per evitare di avere sempre lo stesso segreto condiviso, anche se il pre_master_secret è statico, si aggiunge entropia nuova ad ogni sessione con:
- R_C: nonce (numero casuale) generato dal client
- R_S: nonce generato dal server

Questi due valori sono scambiati in chiaro, ma vengono usati in una funzione hash, come parte di un algoritmo tipo:

**NOTA**: è importante distinguere il requisito di autenticazione da quelle di identificazione
- il primo richiede che i messaggi che arrivano siano effettivamente appartenenti a Bob
- il secondo mi richiede che i messaggi che arrivano siano stati mandati proprio da Bob 

