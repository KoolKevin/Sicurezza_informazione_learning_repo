### Principio di Shannon
confusione & diffusione

principi alla base dei cifrari moderni
- e.g. cifrari a blocci DES e triple DES

sostituzione (con funzione non lineare) -> crea confusione
- il cifrario di cesare non crea confusione, non nasconde relazioni tra testo cifrato e testo in chiaro

trasposizione   -> crea diffusione (come?)

**NB**: per ottenere confusione e diffusione si combinano sostituzione e trasposizione in maniera iterata






## Cifrari simmetrici
si prestano bene per garantire riservatezza. In linea di principio possono anche garantire identificazione ed autenticazione



## Cifrari a flusso
si rifà al one-time pad. 

- genera un flusso di chiave (lungo quanto il testo) con un PRNG
- si fanno degli XOR con i bit di testo
- affinchè la decifrazione avvenga correttamente devo avere **perfetto sincronismo** e lo **stesso seed** del PRNG
    - **chiedi riproducibilità PRNG se seed è TRNG**

Due schemi di funzionamento:
- a flusso sincrono
    - più utilizzati perchè costano meno
- autosincronizzanti
    - più robusto alle perdite di sincronismo
    
**Quand'è che perdo sincronismo?**
attacchi attivi. In questo caso i cifrari a flusso autosincronizzante sono più robusti dato che ho perdita di sincronismo solamente fino a che nel registro di shift rimane presente il/i bit manomessi

**Modello di minaccia**:
- se l'intrusore vuole leggere e basta flusso sincrono va bene
- se l'intrusore ha interessa a dare fastidio invece l'altro


**Perchè le chiavi vanno usate una sola volta**
guarda proprietà dell'xor


**es cifrario a flusso RC4 in WEP**:
sostituito da WPA e WPA2


**attacchi legati alla malleabilità**
modifica al testo cifrato produce un effetto deterministico prevedibile al testo in chiaro

i cifrari a flusso sono malleabili dato che usano un xor

intrusore conosce messaggi predefiniti di una comunicazione strutturata

per le proprietà dell'xor la destinazione decifra m xor p

**il messaggio è**: se si sta utilizzando un cifrario a flusso sappiamo che è malleabile e quindi è opprtuno anche implementare meccanismi che garantiscano integrità



rompere un cifrario a flusso è più facile rispetto rompere un cifrario a blocchi **se non si fanno tutti gli opportuni accorgimenti**.




## Cifrari a blocco
cifrano blocchi di dimensione predefinita e non singoli bit alla volta
- padding di nuovo presente (e sempre presente con anche un intero blocco di padding)

la chiave è la sempre la stessa per ogni blocco





**rete di Feistel**
garantisce confusione e diffusione mediante vari round che applicano sostituzione e trasposizione



A noi interessa lo standard attuale: **AES**

non adotta il modello di Feistel ma adotto anche lui sostituzione e trasposizione


skip di roba.


## Modalità di cifratura | parte che interessa di più!
un cifrario a blocchi non può essere utilizzato solo come visto fino ad adesso (ECB)

Usare ECB non va bene (in generale):
- **determinismo di ECB**
- **malleabilità di ECB**:
    - se uso sempre la stessa chiave posso sostituire il blocco del destinatario con il blocco in cui sono io il destinatario
- se il messaggio è molto piccolo (sta dentro ad un blocco) ECB va vene
    - utilizzato per cifrare chiavi

Altra considerazione sulle modalità di cifratura: **come si propagano gli errori?**
- con ECB non si propaga (l'errore rimane confinato nel blocco modificato)

Le altre modalità diverse da ECB hanno come obiettivo rendere aleatoria l'uscita della cifrazione, ovvero, dato lo stesso messaggio/blocco in chiaro non voglio ottenere lo stesso messaggio cifrato


**LA MONTANARI CHIEDE GLI SCHEMI DELLE MODALITÀ DI CIFRATURA**

**Cypher Block Chaining**:
- pipeline
- che caratteristiche deve avere il v_0? deve essere usato una e una sola volta altrimenti blocchi uguali appartenenti a messaggi in generale diversi vengono cifrati allo stesso modo. Inoltre deve essere casuale e imprevedibile
    - per questi motivi i vettori di inizializzazione vengono generati con PRNG
    - questo requisito non viene sempre rispettato (TLS/SSL)

- efficienza
    - in cifratura non parallelizzabile
    - in decifatura è parallelizzabile

- propagazione degli errori: 
    - **in decifratura**, se un blocco viene alterato, si modifica solo il blocco seguente (vedi esempio)
    - si propaga a tutti i blocchi successivi




**Cypher Feedback e Output Feedback**:
ricordano i cifrari a flusso (sincrono e autosincronizzante)


nello schema OFB il vettore iniziale deve essere unico per DUE MOTIVI
- uno per aleatorietà del cifrato
- l'altro se no ricadiamo nel caso in cui si usa una chiave due volta (vedi xor dei cifrati) il che mina alla riservatezza del messaggio

c'è propagazione dell'errore


**counter**
molto efficiente dato che è parallelizzabile

non ho propagazione dell'errore