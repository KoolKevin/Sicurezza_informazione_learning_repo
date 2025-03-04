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
stesse vulnerabilità e attacchi

non maschera in alcun modo la ridondanza del testo (caratteristiche statistiche del linguaggio)



**NB**: nella crittografia moderna si usa la sostituzione, con l'accorgimento di sostituire **blocchi di bit** per nascondere la ridondanza


### Sostituzione polialfabetica
battaglia navale, con la lettera della chiave seleziono la riga; con la lettera del testo in chiaro seleziono la colonna
- a questo punto una stessa lettera può essere cifrata in vari modi, uno per la lunghezza della chiave
- questo riduce la ridondanza del testo

posso dedurre però la lunghezza della chiave facendo ipotesi su ripetizioni che trovo nel testo cifrato
- trovata una lunghezza di chiave plausibile posso effettuare un attacco con statistica sui caratteri posti ad una determinata distanza l'uno dall'altro