abbiamo visto con i cifrari simmetrici come le chiavi si:
- generano -> prng
- memorizzano -> fs_sicuro
- distribuiscono -> KDC e D.H.

i cifrari simmetrici abbiamo visto essere impiegati più per la riservatezza che per l'autenticazione. I cifrari asimmetrici il contrario

**importante**: 
La conoscenza di PU è una condizione necessaria, ma non sufficiente, per proteggere la riservatezza di un messaggio da inviare a U o per attribuire a U la paternità di un documento firmato con SU.
- Chi usa PU ad un’estremità del canale è, infatti, ben consapevole che all’altra estremità è stata o dovrà essere impiegata la chiave SU, ma **non ha alcuna garanzia che la coppia di numeri SU, PU sia proprio dell’utente di nome U** e non di un impostore che vuole spacciarsi per U. 
- In queste condizioni il pericolo è grosso: si può, infatti, o inviare ad un intruso informazioni riservate destinate solo ad U, o accettare come originate da U informazioni false predisposte dall’intruso. 


### prima cosa: come distribuiamo le chiavi pubbliche? di che cosa ci dobbiamo preoccuare
come abbiamo visto, la chiave pubblica serve a chi vuole comunicare con colui che ha la chiave privata
- cifrando i messaggi che vuole inviargli
- verificando la sua firma

La chiave pubblica deve essere autentica, ovvero essere distribuita in maniera tale da garantire che la sequenza di bit da cui è formata la chiave pubblica è univocamente associata ad un'identità.

R28: “prima d’impiegare una chiave pubblica bisogna o essere certi dell’identità del suo proprietario o poterla verificare”
- come faccio?
- una lookup table non mi basta
    - posso inserire la mia chiave pubblica
    - posso cambiare solo il nome e verificarmi come Corradi

## Infrastruttura a chiave pubblica (PKI)
basato su autorità di certificazione

come faccio io utente che genero una coppia di chiavi ad autenticarle? L'unico modo in questo modello a infrastruttura a chiave pubblica è inserire un terzo ente che fà da autorità fidata che fa da garante per questa autenticazione: **l'autorità di certificazione!**

L'autorità di certificazione:
- autentica le chiavi
- e le distribuisce all'interno di una struttura dati chiamata **certificato** (non come stringa di bit)

slide 8 ambigua, quello che c'è da capire è che il fatto di ricevere un messaggio (integro) firmato con la chiave privata di X ho solo la certezza che il mittente di quel messaggio abbia la chiave privata di X e non che sia effettivamente X. 
- qua manca l'identificazione, ho solo la prova di possesso
- quando l'ente fidato verifica sia prova di possesso che identificazione allora rilascia il certificato (che ha la forma mostrata nella slide)
- anche i certificati restituiti dalla CA sono firmati e resi integri, per cui anche i CA distribuiscono una chiave pubblicay


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


**DH anonimo**
piccola estensione di Diffie-Hellman in cui aggiungo un PRNG che mi permette di cambiare dinamicamente la chiave
- il problema riguardo l'autenticazione dell'Y però rimane
- la minaccia è che si concordi una chiave segreta con un Man in the middle

Per risolvere posso usare una PKI
- i certificati emessi contengono i parametri DH (p, g, Y) di cui voglio garantire l'autenticità al posto della chiave pubblica del mio corrispondente
    - questo meccanismo evita che alice concordi con un MIM la stessa chiave, tuttavia siccome manca identificazione c'è il rischio che io mandi dei campioni di testo cifrato ad una persona diversa da BOB (il certificato è pubblico)
    - poco male siccome l'attaccante non può concordare la stessa chiave con cui decifrare i dati (capisci meglio perchè)
- oppure, ephemeral DH

**NOTA**: è importante distinguere il requisito di autenticazione da quelle di identificazione
- il primo richiede che i messaggi che arrivano siano effettivamente appartenenti a Bob
- il secondo mi richiede che i messaggi che arrivano siano stati mandati proprio da Bob 

