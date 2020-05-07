# El camino subterráneo de Shamash

Reference: [https://omegaup.com/arena/problem/camino-subterraneo-shamash/#problems](https://omegaup.com/arena/problem/camino-subterraneo-shamash/#problems)

|                                  |        |                          |        |
| -------------------------------- | ------ | ------------------------ | ------ |
| Puntos                           | 12.7   | Límite de memoria        | 32 MiB |
| Límite de tiempo (caso)          | 500ms  | Límite de tiempo (total) | 1m0s   |
| Tamaño límite de entrada (bytes) | 10 KiB |

## Descripción

En su búsqueda de Utanapíshtim, Gilgamesh llegó a las montañas Mashu. Al llegar, los hombres escorpión le informaron que tendría que tomar el camino de Shamash, un larguísimo camino subterráneo que medía PP pasos. Gilgamesh decidió anotar esa cantidad, por supuesto, en numerales babilónicos. Estos forman un sistema posicional base 60 que sólo usa dos símbolos para denotar cualquier cantidad de 0 a 59: una barra (I) para denotar el uno y una cuña (L) para denotar el diez. Por ejemplo, el 42 se representa como LLLLII. Adicionalmente, para los números más grandes se usa un separador (.). Por ejemplo, como 2014 = 34+33\*60, entonces se representa como LLLIIII.LLLIII en babilonio. Escribe un programa que lea P y que lo escriba en este sistema babilonio.

## Entrada

Un entero \$`P`$. Puedes suponer que $`1 \leq P \leq 2,000,000,000`\$.

## Salida

El entero $`P`$ escrito en el sistema babilonio.

## Ejemplo

| Entrada     | Salida                |
| ----------- | --------------------- |
| <p>2014</p> | <p>LLLIIII.LLLIII</p> |
