Abbiamo visto come le chiavi dei cifrari simmetrici si:
- generano          -> prng
- memorizzano       -> fs_sicuro
- distribuiscono    -> KDC e DH

I cifrari simmetrici abbiamo visto essere impiegati più per la riservatezza che per l'autenticazione. I cifrari asimmetrici il contrario



### Le proprietà delle chiavi in un cifrario asimmetrico
Un certo utente che vuole rendersi disponibile per la comunicazione **genera una coppia di chiavi** impiegando un **algoritmo di generazione**
- Prima si genera la chiave privata 
    - numero casuale con delle proprietà particolari
        - per questo si uso un algoritmo di generazione apposito
        - primo molto grande se si usa esponenziazione modulare
- La chiave pubblica si può generare facilmente (algoritmo polinomiale) a partire dalla chiave privata. 
    - pensa a DH
- Recuperare la chiave privata a partire dalla chiave pubblica è computazionalmente infattibile
    - algoritmo che produce chiave pubblica a partire da quella privata è una **funzione one-way**

I meccanismi asimmetrici per la riservatezza e per l’autenticazione introducono in Crittografia tre aspetti di importanza fondamentale dal punto di vista della gestione delle chiavi:

1. Una volta scelto un Cifrario asimmetrico, solamente la chiave privata consente ad ogni utente di decifrare i messaggi riservati che chiunque altro gli invia cifrati con la corrispondente chiave pubblica;
    - chiave pubblica cifra per il possessore della chiave privata (che decifra)
    - chi non conosce la chiave privata, non riesce a decifrare (**pseudo-unidirezionalità** (unidirezionale se non si ha la chiave)). 

2. Una volta scelto uno Schema di firma, una sola chiave, quella privata, consente ad ogni utente di firmare messaggi di cui chiunque altro potrà verificare l’autenticità usando la corrispondente chiave pubblica;
    - chiave privata firma, chiave pubblica verifica la firma 
    - chi non conosce la chiave privata non riesce ad imitare la firma del suo proprietario (pseudo-unidirezionalità).

3. **La chiave pubblica di un utente può essere comunicata a tutti gli altri senza alcuna forma di riservatezza**
    - questo è d'altronde il grosso vantaggio della crittografia asimmetrica, non bisogna preoccuparsi di come comunicare le chiavi in maniera sicura

Aspetti a carico di ogni utente U di un meccanismo asimmetrico:
- **Disponibilità della chiave pubblica PU** (cifrario a chiave pubblica):
    - chiunque ne ha bisogno deve poterla facilmente reperire; a tal fine possono essere impiegati
        - o un incontro diretto tra gli interessati,
        - o un’e-mail inviata dal proprietario ai suoi potenziali corrispondenti,
        - o un file che questi ultimi possono scaricare dal sito del proprietario,
        - o, ed è questa la soluzione migliore, **un record che chiunque può ottenere da un database**.
- **Segretezza della chiave privata SU** (firma digitale a chiave pubblica):
    -  solo il proprietario deve poterne disporre; 
    - a tal fine deve conservarla cifrata (con la tecnica PBE) in una memoria elettronica con controllo d’accesso (hard disk, o floppy, o, molto meglio, una smart card).
    - analogo ai segreti del caso simmetrico

- **Autenticità della chiave pubblica**: il proprietario delle chiavi deve rendere di pubblico dominio anche informazioni atte a consentire a chiunque di verificare, in modo sicuro ed inequivocabile, **la sua identità**;






### Autenticità della chiave pubblica

**IMPORTANTISSIMO**:

La conoscenza di PU è una condizione necessaria, ma non sufficiente, per proteggere la riservatezza di un messaggio da inviare a U o per attribuire a U la paternità di un documento firmato con SU.
- Chi usa PU ad un’estremità del canale è, infatti, ben consapevole che all’altra estremità è stata o dovrà essere impiegata la chiave SU, ma **non ha alcuna garanzia che la coppia di numeri SU, PU sia proprio dell’utente di nome U** e non di un impostore che vuole spacciarsi per U. 
- In altre parole, l'utente che usa PU **non ha garanzie sulla sua autenticità**
- In queste condizioni il pericolo è grosso: si può, infatti, o inviare ad un intruso informazioni riservate destinate solo ad U (cifrario a chiave pubblica), o accettare come originate da U informazioni false predisposte dall’intruso (firma digitale a chiave pubblica). 

**Man-in-the-Middle attack**
Consideriamo il caso di un Cifrario asimmetrico e supponiamo che due utenti A e B abbiano inserito in un database DB le loro chiavi pubbliche PA e PB.

L’intruso I che vuole violare la riservatezza delle loro comunicazioni deve fare solo tre cose, neanche troppo difficili:
1. sostituisce la sua chiave pubblica PI al posto di quelle di A e di B
    - o modificando il contenuto di DB
    - o dirottando tutte le richieste di accesso sul suo calcolatore
2. quando A invia un c = E_PI(m) a B, lo intercetta, lo decifra, lo cifra con PB (pubblica) e lo inoltra a B;
3. quando B risponde ad A ripete lo stesso attacco attivo, completando così la violazione della riservatezza dei messaggi tra A e B. 

Discorso analogo per la firma digitale:
- dopo aver sostituito la chiave pubblica di B con la mia PI
- posso inviare un messaggio firmato con PI spacciandomi per B
- A verificherà la firma pensando che la paternità del messaggio sia di B e non si accorgerà di niente

La contromisura è **comunicare chiavi pubbliche di cui è possibile verificare l’origine** (autenticità)
- autenticità concetto non chiarissimo, una definizione che mi piace è: "associare una identità ad un dato, in questo caso la chiave pubblica" 

```R28: ”prima d’impiegare una chiave pubblica bisogna o essere personalmente certi dell’identità del suo proprietario o disporre di un documento (certificato) che l’attesta, redatto da un’Entità di cui ci si fida”.```






## Certificati
Due che **si conoscono** non hanno particolari problemi ad autenticare le loro chiavi pubbliche. 
- Una soluzione ovvia è che se le scambino durante un incontro diretto.
- Una seconda soluzione è che se le invino per posta ed usino poi una telefonata per confermarne a voce il valore. 

**In generale, però, chi deve usare una chiave pubblica non conosce il suo proprietario**.

In questo caso la certificazione della associazione chiave-proprietario(identità) deve essere:
- fatta da **qualcuno di cui ci si può fidare** che conosoco,
- comunicata con un messaggio di cui sia possibile **verificare integrità ed origine**.
    - altrimenti l'attaccante può presentare la sua chiave pubblica e un certifiato falsato

**Il messaggio che attesta l’identità del proprietario di una chiave pubblica è detto certificato**. 
- L’Ente T che produce certificati è detto terza parte fidata.
- La **firma digitale di T** apposta sul certificato tramite un algoritmo asimmetrico è l’accorgimento efficiente che consente a tutti gli interessati di verificare l’integrità e l’origine dell’attestato.
- **NB**: Il presupposto è che tutti abbiano una copia autentica di PT (magari preinstallata nel browser, PC, ...)
    - altrimenti non si saprebbe come verificare la firma di T.

In questo caso **l'attacco dell'uomo in mezzo non è più possibile** in quanto l'attaccante non è in grado di forgiare un certificato riproducendo una firma della CA valida.



### Come è fatto un certificato?
Nella sua forma più ridotta, il certificato della chiave pubblica di un generico utente X da parte di T è formato da:
- un testo in chiaro m che forma il certificato e che contiene:
    - l'identità X del proprietario della chiave pubblica
    - la chiave pubblica associata 
    - svariati altri metadati
- affiancato dal hash firmato con ST del certificato
- cert(PX, T) = X || PX ||S_ST(H(X || PX))

**OSS**: schema di un tipico messaggio firmato

Oltre a questo, un certificato contiene anche molti altri campi:
- metadati identificativi sul possessore del certificato
- metadati sull'utilizzo della chiave pubblica
    - algoritmo asimmetrico da utilizzare con la chiave pubblica contenuta nel certificato
    - **periodo di validità**
- metadati su T
    - dati identificativi
    - algoritmi di hash e firma utilizzati
- estensioni utilizzate per varie cose (vedi dopo)

Una volta predisposto il certificato di X, T può consegnarlo
- ad un altro utente Y che gli ha direttamente richiesto l’autenticazione di PX,
- allo stesso utente X, che provvederà poi ad inoltrarlo a chi glielo richiederà,
- **ad un database a libero accesso (directory)**, da cui chiunque potrà estrarlo.

Chi riceve un certificato emesso da T
- si procura PT in modo sicuro (se non l’ha già)
- calcola u = V_PT(cert(PX,T)) = vero/falso
- e solo in caso di verifica positiva **trasferisce su PX la “fiducia” che ha su PT**. 


**Molto importante**
quello che c'è da capire è che il fatto di verificare la firma di un messaggio (integro) allegedly proveniente da X, firmato con una chiave privata (allegedly di X), mi da solo la certezza che il mittente di quel messaggio abbia la chiave privata e non che sia effettivamente X (POP). 
- **senza certificato manca l'identificazione del mittente, ho solo la prova di possesso**
- quando **l'ente fidato verifica sia prova di possesso, che identificazione**, allora rilascia il certificato (che ha la forma mostrata nella slide)
- anche i certificati restituiti dalla CA sono firmati e resi integri, per cui anche i CA distribuiscono una chiave pubblica

Allo stesso modo, se uno presenta un certificato (che ricordiamo essere pubblico) non è detto che sia il proprietario della chiave pubblica al suo interno
- bisogna verificare che abbia anche la relativa chiave privata (POP) altrimenti, di nuovo, manca l'identificazione del mittente

**In conclusione**: per identificare correttamente un interlocutore che non si conosce a priori, è necessario:
- sia il certificato -> che associa una identità ad una chiave pubblica
- che la POP         -> che dimostra che che il mittente è il proprietario della chiave pubblica presentata (e per transitività, identifica il mittente con l'identità presente nel certificato)


Occorre anche decidere chi può ricoprire il ruolo di T; due sono i modelli in uso.
1. T è un Ente ufficialmente riconosciuto come fidato ed istituzionalmente incaricato di emettere certificati.
2. T è un qualsiasi utente che dichiara spontaneamente di conoscerne bene un altro; chi si fida di lui, si fida anche della sua dichiarazione.

La prima soluzione, detta della **Autorità di certificazione (o CA)**, è formalizzata nello **standard ISO X.509** ed è impiegata sia in applicazioni sicure, sia nelle reti sicure; la seconda soluzione, propugnata dal **PGP**, è il contesto preferito da tutti quelli che non vogliono un Grande Fratello sopra di loro. 
