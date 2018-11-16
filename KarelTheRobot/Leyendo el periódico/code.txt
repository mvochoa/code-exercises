iniciar-programa
    define-nueva-instruccion gira-derecha como
    inicio
        repetir 3 veces
        inicio
           gira-izquierda;
        fin;
    fin;

    define-nueva-instruccion orientar como
    inicio
        si orientado-al-sur entonces
        inicio
           gira-izquierda;
           gira-izquierda;
        fin;
        si orientado-al-oeste entonces
        inicio
           gira-izquierda;
           gira-izquierda;
           gira-izquierda;
        fin;
        si orientado-al-este entonces
        inicio
           gira-izquierda;
        fin;
    fin;

    define-nueva-instruccion accion2(n) como
    inicio
        si frente-libre entonces
        inicio
           avanza;
           accion2(sucede(n));
           avanza;
        fin
        sino
        inicio
            coge-zumbador;
            gira-izquierda;
            gira-izquierda;
        fin;
    fin;

    define-nueva-instruccion accion(n) como
    inicio
        si frente-libre entonces
        inicio
           avanza;
           accion(sucede(n));
           avanza;
        fin
        sino
        inicio
            gira-izquierda;
            accion2(sucede(0));
            gira-derecha;
        fin;
    fin;

    inicia-ejecucion
        orientar;
        avanza;
        gira-izquierda;
        accion(0);
        gira-derecha;
        avanza;
        deja-zumbador;
        apagate;
    termina-ejecucion
finalizar-programa