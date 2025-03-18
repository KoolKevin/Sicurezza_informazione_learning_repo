### Master key
- può cifrare anche tante chiavi
- le chiavi cifrate sono casuali, imprevedibili e generate di volta in volta
- la modalità ECB va bene siccome il determinismo è annullato dalle proprietà delle chiavi

Chiavi di sessione
- usa e getta

Lo schema master key è complicato dato che richiede dei canali dedicati costosi (out of band)
- non scala quando c'è una comunintà di utenti (N^2 canali per N utenti)



 
### Centro distribuzione delle chiavi
ogni utente preconcorda con il centro una chiave con cui cifrare le chiavi da usare

chiaramente il centro di distribuzione, che è una terza parte, bisogna che sia fidato da parte degli utenti. Non in tutti contesti gli utenti si fidano 


...


RA è una sorta di identificatore di sessione di richiesta di una chiave
- associa una richiesta ad una chiave di sessione


Se non completo il protocollo posso replicare l'invio della sfida a B e indurlo a cifrare messaggi più volte con la stessa chiave
- come si evita?
- si cambia il protocollo per far si che B sia sicuro dell'identità di A (chiaramente non mi fai solo un echo)




ho problemi se implemento in modalità ECB?
- determinismo
    - se non si modifica di molto R_B l'attaccante può provare  


- malleabilità

attacco di replicabilità
- con ECB e_k(A||B) diventa e_k(A)||e_k(B) se i blocchi sono allineati
- se ho ottenuto la stessa chiave k da una sessione precedente posso spacciarmi per A


**confronto con protocolo 4-way**
- 4 way invia un messaggio in meno tra i due interlocutori
- 4 way effettua due operazioni di decifratura in più rispetto alle 0 si 5-way
- A di 4 way fa una decifratura in meno
- B di 4 way fa una decifratura in meno

con ECB cosa succede?
- se non si tiene traccia delle chiavi già usate I può indurre B a riusare una chiave che lui conosce da una sessione precedente




### Accordo DH
non preconcordiamo prima un segreto (master key), lo concordiamo online sul momento