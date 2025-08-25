Abbiamo 3 servizi basati su chiave pubbliche.
- L'azienda ha quindi bisogno di una PKI per certificare questa pub_key.
- Dovendo pagare ogni certificato, quante coppie di chiavi devono essere usate (1, 2 o 3)?



Una coppia di chiavi non basta... perchè?
- vedi sotto

Se consideriamo due coppie di chiavi bisogna porsi tre domande:
- possiamo accorpare chiave di firma e chiave di cifratura?
    - no! 
    - Se io faccio firma digitale con validità legale devo avere garanzia di supporto al **non ripudio** e quindi la chiave privata è unicamente associata al dipendente.
        - -> se io uso la stessa chiave per cifrare, quando un dipendente **se ne va i messaggi sono irrecuperabili**
    - potrei pensare di star decifrando ma in realtà sto firmando qualcosa 
        - proprietà di reversibilità di RSA
    - (c'è anche un attacco con la proprietà moltiplicativa)
        -  (l’intrusore può aver intercettato i messaggi da Alice a Bob, che non riesce a decifrare, può mettere in piedi un server falso e, oscurato k cifrato con la chiave pubblica, mandarlo ad Alice, che firma e regala all’intrusore la chiave k)
- possiamo accorpare chiave di firma con chiave di identificazione?
    - no! se l'identification server NON fidato mi manda un nonce che non è un nonce, ma qualcosa che l'attacante voglia far firmare alla vittima, siccome la firma ha validità legale, e non è ripudiabile, io firmo qualcosa senza accorgermene e sarò tenuto a risponderne
- possiamo accorpare chiave di identificazione con chiave di cifratura?
    - no! posso far finta di essere un IdS e porre come nonce il cifrato di un messaggio che ho intercettato, siccome sto usando RSA nel rispondermi alla sfida mi viene mandato il testo decifrato
    - questa vulnerabilità può essere mitigata utilizzando un nonce di dimensione fissa piccola (128 bit)

Se nessuna può essere accoppiata, allora necessariamente bisogna usare 3 coppie



Se non facciamo altre ipotesi abbiamo bisogno di tre coppie

Con l'ipotesi sul nonce, al più possiamo accoppiare chiave di identificazione con chiave di cifratura






Abbiamo **identificazione unilaterale!**
- questa è una ipotesi forte
- l'identification server potrebbe essere malevolo
- **tom non ha nessuna garanzia che l'IS che lo sfida sia quello della propria azienda**





Negli appunti sul drive, a pagina 110 hai altra roba