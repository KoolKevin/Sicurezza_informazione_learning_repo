utile da sapere dato che nella crittografia moderna ci si ispira alla crittografia classica

### Sostituzione (monoalfabetica)
due alfabeti
- uno del messaggio
- uno cifrante

non maschera in alcun modo la ridondanza del testo (caratteristiche statistiche del linguaggio)

spazio degli alfabeti cifranti (chiavi) ampio
- vulnerabilità di ridondanza del testo
- attacco con statistiche


### Trasposizione
Substitution ciphers **preserve the order of the plaintext symbols but disguise them**.

Transposition ciphers, in contrast, **reorder the letters but do not disguise them**.

Le proprietà statistiche del testo cifrato non solo rimangono intatte ma sono le stesse del testo in chiaro
- stesse vulnerabilità e attacchi della sostituzione (attacco con statistiche)

non maschera in alcun modo la ridondanza del testo (caratteristiche statistiche del linguaggio)



**NB**: nella crittografia moderna si usa la sostituzione, con l'accorgimento di sostituire **blocchi di bit** per nascondere la ridondanza


### Sostituzione polialfabetica
la chiave definisce tanti alfabeti cifranti

battaglia navale, con la lettera della chiave seleziono la riga; con la lettera del testo in chiaro seleziono la colonna
- a questo punto una stessa lettera può essere cifrata in vari modi, uno per la lunghezza della chiave
- questo riduce la ridondanza del testo

posso dedurre però la lunghezza della chiave facendo ipotesi su ripetizioni che trovo nel testo cifrato
- trovata una lunghezza di chiave plausibile posso effettuare un attacco con statistica sui caratteri posti ad una determinata distanza l'uno dall'altro


### One-time pad
Chiave lunga quanto la frase


- Inviolabile con attacco passivo dato che più ipotesi di decifrazione possono produrre un messaggio sensato 
- problema: per trasmettere un messaggio riservato su un canale insicuro bisogna concordare una chiave altrettanto lunga su un canale sicuro; tantovale trasmettere direttamente il messaggio sul canale sicuro