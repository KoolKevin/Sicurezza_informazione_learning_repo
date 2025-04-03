noi facciamo solo RSA e non ci preoccupiamo degli algoritmi basati sulle curve ellittiche (algoritmi asimmetrici con le prestazioni migliori)

nota: di fatto DH è un cifrario asimmetrico


**OSS IMPORTANTE**: una cifratura asimmetrica è molto meno efficiente di una cifratura simmetrica?
- perchè? dobbiamo fare delle esponenziazioni modulari!
- nel caso simmetrico devo fare solo trasposizioni e sostituzioni


### Come facciamo ad ottimizzare l'esponenziazione modulare?
Algoritmi di teoria dei numeri non sono da sapere per filo e per segno. Tuttavia è importante sapere quali sono gli algoritmi che rendono più efficiente per l'eseponenziazione modulare. Perchè? Nelle implementazioni i vari provider possono utilizzare algoritmi diversi, è desiderabile utilizzare quello più efficiente.
- nel caso di RSA: il tempo di esponenziazione modulare scala con il cubo della dimensione dei bit (1024 -> 2048 -> 8 volte il tempo)

### Come facciamo a generare dei numeri primi molto grandi?
test deterministici -> esponenziali
test probabilistici -> polimoniali ma vanno ripetuti



### RSA
le modalità di cifratura ricompaiono anche in algoritmi asimmetrici

vogliamo rendere aleatoria un uscita aleatoria 

...

per decifrare non faccio la radice n-esima (che è un problema difficile) piuttosto faccio un'altra esponenziazione modulare con dei parametri scelti bene

**NB**: qua la cosa importante è la frammentazione. Siccome ad un unico testo in chiaro deve corrispondere uno e un solo testo cifrato, le operazioni modulari devono essere fatte su dei numeri minori di n -> frammentazione



NOTA: per la decifrazione esiste un particolare algoritmo (di Gartner) che è 4 volte più veloce rispetto a repeated square and multiply -> da tenere a mente nei provider 