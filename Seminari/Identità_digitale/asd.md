identità digitale federata e decentralizzata


situazione tradizionale: server centrale che gestisce le identità dei vari utenti
- problema: tante identità tante quali sono i servizi di cui si vuole usufruire


## Identità federata
federeta significa che ci sono relazioni di fiducia


stessa identità across services

drawback:
- relazione di fiducia non immediata da stabilire
- l'identity provider sa che tu stai cercando di accedere ad un servizio (privacy lesa)
    - google sa che sto accedendo ad un servizio
- si tende a creare un ecosistema chiuso di dipendenze
    - io tendo a fare il login con google e pochi altri
    - se google domani smette di darmi la possibilità io non riesco più ad accedere al servizio
    - e mi diventa difficile staccarmi e a migrare su un altro identity provider




### OAuth
è un modo per un servizio di accedere a dati presenti su un altro servizio

mi da un accesso granulare rispetto alle risorse richieste dal servizio



...


ResourceOwner e client (detti anche Identity provider e Service Provider) si fidano l'uno dell'altro clientId e secret sono robe già scambiate (è così che viene implementata la relazione di fiducia)
- se tu dici di essere chi sei utilizzerai il segreto preconcordato

...


Authorization Code è un qualcosa che mi dice che i passi 1-4 sono stati già fatti quando vengo redirezionato a terrible pun tramite il redirect uri


**NB**: se lui mi desse direttamente l'access token, quest'ultimo sarebbe intercettabile
- l'access token serve solamente al servizio terrible pun, non serve ne al resource owner ne al browser, ed infatti solo lui lo ottiene








Nota: Qua siamo ancora in un contesto federato (e quindi centralizzato)
- nel contesto decentralizzato non esiste un entità centrale con cui ho una relazione di fiducia, 






### OpenID Connect
è un layer che sta sopra ad Oauth in cui invece di accedere a dati generici si accede all'identità digitale

oltre ad un access token viene dato al client un id token

perchè mi servono entrambi? Mi serve un anche un id token in quanto il client non sa chi è il resource owner. Con l'id token (che è un JWT) google mi certifica la mia identità firmando il JWT
- mi fa la parte di identificazione

Con queste informazioni presenti nell'id token, terrible pun of the day può registrare l'utente con queste informazioni certificate da google
- se le avessi chieste con l'access token e basta non sarebbero state certificate con la firma di google










## Self-sovereign identity (SSI)
...



decentralized identifiers
- non eliminabili in quanto replicati sui peer del sistema decentralizzato

...


non c'è un entità centralizzata prendo did e chiavi pubbliche dal verifiable data registry (blockchain) con cui posso verificare le VP e VC



selective disclosure

