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
Lo scambio di chiavi Diffie-Hellman è un protocollo che permette a due utenti di scambiarsi una chiave segreta comunicando pubbicamente, senza rivelare la chiave segreta a degli attaccanti. 
- non preconcordiamo prima un segreto (master key), lo concordiamo online sul momento
- non c'è più bisogno di un accordo con cui cifrare le chiavi con cui comunichiamo, possiamo scambiare la chiave di sessione in maniera intelligente direttamente

I due partecipanti:
- concordano un numero casuale iniziale (possibilmente grande e primo in modo da annullare gli attacchi di forza bruta -> conseguenza di esponenziazione modulare) 
    - Per concordare un numero devono comunicarselo e quindi renderlo pubblico
    - Questo numero non può quindi essere la chiave
- Tuttavia, i due partecipanti generano anche un numero casuale segreto (e quindi distinto) che sarà una sorta di "chiave segreta personale"
- Chiaramente, non possiamo comunicare direttamente questi due nuovi numeri casuali, altrimenti saremmo punto e a capo. Tuttavia, se comunicassimo non direttamente le chiavi segrete, ma il **risultato di un operazione non facilmente reversibile** di quest'ultime con il primo numero casuale condiviso:
    - le chiavi segrete non verrebbero comunicate in quanto mascherate dall'operazione non reversibile
    - se l'operazione ha delle proprietà particolari (esponenziazione modulare), i due interlocutori **possono generare la stessa identica chiave di sessione** a partire solamente dai risultati dell'operazione non reversibile e quindi **senza mai essersi scambiati pubblicamente i loro segreti!**
        - ripeto per la terza volta: è fondamentale che l'operazione sia non reversibile, altrimenti un attaccante in ascolto potrebbe recuperare la chiave segreta di un interlocutore a partire dal risultato dell'operazione tra segreto e numero casuale concordato. 

**OSS**: D-H classico è vulnerabile agli attacchi Man-in-the-Middle se non viene autenticato. Per questo spesso viene combinato con altri meccanismi di autenticazione, come certificati digitali o firme crittografiche.