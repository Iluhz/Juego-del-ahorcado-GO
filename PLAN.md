# Plan de Desarrollo - Juego del Ahorcado en Go

## Idea inicial
Desarrollar el juego clásico del ahorcado como aplicación de terminal en Go,
sin el uso de inteligencia artificial, como ejercicio de la materia de estructura de datos.

## Proceso de desarrollo

### Fase 1 - Estructura base
- Definir el array de palabras
- Implementar la selección aleatoria de una palabra usando posiciones random

### Fase 2 - Lógica del juego
- Control de intentos fallidos (máximo 6)
- Validación de letras ingresadas por el jugador
- Detección de letras repetidas
- Condición de victoria y derrota

### Fase 3 - Interfaz en terminal
- Dibujo progresivo del ahorcado según los errores
- Mostrar la palabra con guiones y letras adivinadas
- Mostrar letras ya utilizadas

## Dificultades y aprendizajes
- La parte más retadora fue desarrollar la lógica del juego desde cero
- Fue necesario pensar bien la estructura antes de escribir el código
- Proyecto desarrollado íntegramente sin IA, como ejercicio de aprendizaje
- Experiencia interesante al desarrollar algo funcional en un lenguaje nuevo
