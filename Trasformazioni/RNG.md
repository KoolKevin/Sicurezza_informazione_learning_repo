**Come si generano le chiavi?**
- i valori dei bit devono essere casuali e imprevedibili
- casualità
    - valori equiprobabili
    - indipendenza statisica (probabilità di una sequenza di bit == prodotto delle probabilità)
- come si ottiene casualità?
    - TRNG
        - imprevedibile
        - problemi:
            - frequenza di generazione del rumore
            - riproducibilità: A e B hanno ognuno il proprio TRNG che generà numeri diversi, ma se loro devono pordurre la stessa chiave?
    - PRNG
        - algoritmi che generano sequenze di bit casuali in maniera deterministica
            - configurati con un seme
            - pseudo casuale perchè ad un certo punto la sequenza si ripete
        - il valore in uscita è **prevedibile** se non si adottano opportuni accorgimenti 
        - chi genera il seme? un TRNG

**Come si memorizzano?**
**Come si distribuiscono?**
**Come vengono utilizzate?**

Un sistema è robusto solo se tutte le componenti sono sicure (anello più debole)



### Algoritmi di Hash 

...

- unidirezionalità è importante perchè permette di fare una sorta di cifratura senza chiave

h1 = imporonta calcolata con h_0 e la finestra di r bit corrente del messaggio

- m || S funziona; S || m no