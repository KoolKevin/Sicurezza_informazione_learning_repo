Macchine a funzionamento segreto
- manutenibilità? scalabilità? ispezione?

Algoritmi segreti
- come sopra

**Algoritmi noti con parametri segreti**
- capisci meglio cosa significa robustezza
- se viene scoperto il parametro lo cambio e lascio il parametro

```
d'ora in poi con trasformazioni segrete intenderemento algoritmo noto con parametro segreto (segretezza del parametro della trasformazione)
```

### Cifrari simmetrici e asimmetrici
Algoritmi (un cifrario è un algoritmo di crittografia) a chiave
- simmetrica:   k_s == k_d (o facilmente derivabile)
    - segretezza totale della chiave
    - come le distribuisco? entrambi le devono avere
    - tipico per riservatezza

- asimmetrica:  k_s != k_d 
    - k_d (chiave pubblica può essere nota)
    - ma non si può derivare k_s da k_d (mentre il contrario è facile)
    - tipico per autenticazione


**proprietà delle chiavi**


**memorizzazione del parametro segreto**

il segreto serve a cifrare, ma anch'esso dev'essere cifrato... come fare?


Output cifrato con presenza di pattern ripetitivi da informazioni sfruttabili dall'intrusore
- per questo motivo è importante che gli algoritmi crittografici producano un'uscita il più possibile aleatoria











### Complessità computazionale
...


Nel contesto della sicurezza informatica devo considerare anche il caso migliore
- bisogna adottare degli accorgimenti per evitare che l'intrusore si trovi di fronte ad istanze facili del problema
- e.g. segreto con delle parti che sia ripetono


**unità di misura di robustezza**:
- Anno MIPS -> obsoleto
- Livello di sicurezza (che nome è?) -> numero di operazioni per un'attacco di forza bruta

128 bit minimo livello di sicurezza






### Schema a 3 livelli di segreto
se cifro tutto sono contento? No! Cifrare tutto con la stessa chiave da all'attancante un numero elevato di file cifrato con la stessa chiave -> facilito l'attacco

- passphrase
- master key, chiave con cui vado a cifrare le chiavi di cifratura dei file
    - memorizzato in maniera sicura nel FS
- chiavi di cifratura dei file
    - memorizzato in maniera sicura nel FS

RNG generano chiavi master e chiavi dei file