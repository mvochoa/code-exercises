iniciar-programa
    define-nueva-instruccion accion como
    inicio
        si no-junto-a-zumbador entonces
        inicio
            si izquierda-libre entonces
            inicio
                gira-izquierda;
                avanza;
                accion;
                deja-zumbador;
                deja-zumbador;
            fin
            sino
            inicio
                si derecha-libre entonces
                inicio
                   gira-izquierda;
                   gira-izquierda;
                   gira-izquierda;
                   avanza;
                   accion;
                   deja-zumbador;
                   deja-zumbador;
                   deja-zumbador;
                   deja-zumbador;
                fin
                sino
                inicio
                    avanza;
                    accion;
                    deja-zumbador;
                fin;
            fin;
        fin;
    fin;

    inicia-ejecucion
        accion;
        coge-zumbador;
        apagate;
    termina-ejecucion
finalizar-programa
