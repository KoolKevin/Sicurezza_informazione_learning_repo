abbiamo visto con i cifrari simmetrici come le chiavi si:
- generano -> prng
- memorizzano -> fs_sicuro
- distribuiscono -> KDC e D.H.

i cifrari simmetrici abbiamo visto essere impiegati più per la riservatezza che per l'autenticazione. I cifrari asimmetrici il contrario

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

