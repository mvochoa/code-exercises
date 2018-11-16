iniciar-programa
    define-nueva-instruccion gira-derecha como
    inicio
        repetir 3 veces
        inicio
            gira-izquierda;
        fin;
    fin;

    define-nueva-instruccion poner como
    inicio
        si no-junto-a-zumbador entonces
        inicio
            deja-zumbador;
        fin;
    fin;

    define-nueva-instruccion x como
    inicio
        si frente-libre entonces
        inicio
            poner;
            avanza;
            gira-derecha;
            avanza;
            gira-izquierda;
            x;
            poner;
            avanza;
            gira-izquierda;
            avanza;
            gira-derecha;
        fin
        sino
        inicio
            poner;
            gira-izquierda;
            mientras frente-libre hacer
            inicio
                avanza;
            fin;
            gira-izquierda;
        fin;
    fin;

    inicia-ejecucion
        mientras no-orientado-al-sur hacer
        inicio
            gira-izquierda;
        fin;
        mientras frente-libre hacer
        inicio
            avanza;
        fin;
        si derecha-libre entonces
        inicio
            gira-derecha;
        fin;
        mientras frente-libre hacer
        inicio
            avanza;
        fin;
        mientras no-orientado-al-norte hacer
        inicio
            gira-izquierda;
        fin;
        x;
        poner;

        apagate;
    termina-ejecucion
finalizar-programa