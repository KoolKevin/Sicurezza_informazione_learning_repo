## Funzioni unidirezionali e pseudo-unidirezionali
Per ottenere sicurezza, la Crittografia chiama in causa **funzioni facili da calcolare e difficili da invertire**, per cui ha coniato la denominazione di funzioni unidirezionali (one-way functions).

Abbiamo già detto ad esempio:
- che la funzione hash sicura deve poter essere calcolata da un algoritmo con tempo polinomiale e che non deve esistere un algoritmo con tempo polinomiale in grado di risalire da un’impronta H(x) ad uno degli x che la generano;

- che la funzione di cifratura di un cifrario deve poter essere calcolata da un algoritmo con tempo polinomiale e che chi non conosce la funzione inversa non deve poter disporre di un altro algoritmo con tempo polinomiale in grado di risalire da E(m) a m;


OSS: Non si può escludere a priori che esistano anche alcuni risultati facili da invertire: l’importante è che siano pochissimi, per rendere assolutamente trascurabile la probabilità che l’attaccante si trovi proprio in quelle situazioni. 

**Funzione unidirezionale**: Una funzione f è detta unidirezionale se
- è invertibile,
- facile da calcolare
- e se per quasi tutti gli x appartenenti al dominio di f è difficile risolvere per x il problema y = f (x).

Due notizie, una cattiva ed una buona.
- La cattiva è che finora la Teoria della complessità computazionale non è riuscita a dimostrare l’esistenza delle funzioni unidirezionali.
- La buona. I Crittografi sono riusciti ad individuare funzioni di compressione, di cifratura e di firma calcolabili con algoritmi polinomiali ed hanno altresì verificato che tali funzioni sono invertibili solo con algoritmi o esponenziali, o sub-esponenziali. 

Un’ultima precisazione: nei casi della decifrazione e della firma, l’utente deve poter fare in tempo polinomiale quello che l’intruso non deve invece poter fare. Occorre dunque un ulteriore modello.

**Funzioni pseudo unidirezionali**: “una funzione F è detta pseudo-unidirezionale (trapdoor one-way function) se appare come unidirezionale per chiunque non sia in possesso di una particolare informazione sulla sua costruzione (chiave)"