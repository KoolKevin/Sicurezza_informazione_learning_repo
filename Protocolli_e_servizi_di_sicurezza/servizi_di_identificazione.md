### BASI (molto importante)
obiettivo:
- avere la certezza che al termine del protocollo chi ha dichiarato una certa identità sia confermato che sia effettivamente chi dice di essere

requisiti:
- ...

proprietà da scegliere se garantire o meno in base all'implementazione
-

### Dimostrazione di conoscenza

**protocolli attivi e passivi**
- passivo (o debole): uso sempre lo stesso segreto come prova di identità (conoscenza)
- attivo (o forte): la prova di conoscenza cambia da sessione di identificazione all'altra
    - chiaramente più robusto rispetto a protocolli passivi


### Identificazione passiva
cringe

in locale anche accettabile

nei sistemi in rete inaccettabile, è sempre possibile l'attacco con replica



### Identificazione attiva
ci concentriamo su questa

a partire dalla prova di identità (intercettabile come la pwd) deve essere difficile risalire al segreto che l'ha generata

**one-time pwd**
...

si applica una funzione F che produce le prove di identità più volte per produrre N prove

l'attacco con replica non è più possibile dato che il verificatore cambia le prove di conoscenza ad ogni identificazione

robusto, ma ad un certo punto le prove finiscono -> necessità di rifare tutto il setup



**sfida/risposta**
...


se il nonce è prevedibile l'attaccante può sosituirsi a B e farsi mandare H(S || RB'). A questo punto l'intrusore può spacciarsi per A con B nella prossima sessione dato che conosce l'hash che verrebbe usato da A alla sfida con RB'. 


### Identificazione mutua
...

esempio non robusto
- tengo due sessioni attive

**attacco di interleaving**
apro due sessioni contemporanee con due entità diverse


**Attacco di reflection**
apro due sessioni contemporanee con la stessa entità





### qual'è il motivo alla base della non sicurezza del protocollo di identificazione mutua che abbiamo visto
1. c'è una forte simmetria tra i messaggi che A e B si mandano a vicenda
2. l'intrusore può rispondere con un ritardo rispetto a quando gli è stata fatta la sfida (aspetto di recuperare in qualche modo la risposta giusta)
    - non c'è un timestamping che limita la validità temporale dei messaggi
3. mancano dei legami tra i messaggi che mi consentono di collegare, ad esempio, una risposta ad una sfida



Contromisure qualitative:
- numeri random sicuri
- numeri di sequenza per collegare i messaggi
- timestamping

vengono usate in combinazione per eliminare i svantaggi


analogamente ai protocolli di autenticazione in cui si combina
- quello che sono
- quello che conosco
- ...






protocollo di identificazione = real time
- per mantenere l'identità dell'entità identificata nel tempo occorre **affiancare ai protocolli di identificazione, dei protocolli di autenticazione**
    - e.g. utilizzare degli HMAC dopo l'identificazione

