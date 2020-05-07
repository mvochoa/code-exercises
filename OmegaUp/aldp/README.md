# Alicia y las llaves doradas de las puertas

Reference: [https://omegaup.com/arena/problem/aldp/#problems](https://omegaup.com/arena/problem/aldp/#problems)

|                                  |        |                          |        |
| -------------------------------- | ------ | ------------------------ | ------ |
| Puntos                           | 10.3   | Límite de memoria        | 32 MiB |
| Límite de tiempo (caso)          | 1s     | Límite de tiempo (total) | 1m0s   |
| Tamaño límite de entrada (bytes) | 10 KiB |

## Descripción

Después de su larga caída, Alicia quedó atrapada en el fondo del agujero del conejo. Afortunadamente no estaba totalmente oscuro y pudo ver que había muchas puertas (todas cerradas) y una mesa con varias llaves doradas. Seguramente alguna llave abriría alguna de estas puertas, así que Alicia lo intentó y pronto descubrió que las chapas de las puertas eran de diferentes tamaños y que estaban ordenadas en fila, de la más pequeña a la más grande. Cada llave sólo abriría la chapa del mismo tamaño. Como en verdad eran muchas puertas y muchas llaves, Alicia pensó que le tomaría mucho tiempo intentar abrir todas las puertas con las llaves correctas. Ayuda a Alicia a determinar qué llaves abren qué puertas.

## Entrada

Un entero \$`M`$ seguido de los tamaños de las chapas $`P_1,...P_M`$. Posteriormente un entero $`N`$ seguido de los tamaños de las llaves $`L_1,...L_N`$. Puedes suponer que $`1 \leq N, M \leq 100,000`$, que los tamaños de las chapas son distintos con $`1 \leq P_1 < P_2 < ... < P_M \leq 100,000`$ y que los tamaños de las llaves cumplen $`1 \leq L_i \leq 100,000`$ para toda $`1 \leq i \leq N`\$.

## Salida

Para cada una de las \$`N`$ llaves, el número de la puerta que puede ser abierta con esa llave o $`0`\$ si no corresponde con ninguna puerta.

## Ejemplo

| Entrada                                     | Salida         |
| ------------------------------------------- | -------------- |
| <p>5<br />1 3 4 5 9<br />4<br />2 5 1 8</p> | <p>0 4 1 0</p> |
