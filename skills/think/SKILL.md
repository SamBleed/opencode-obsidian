---
name: think
description: >
  Aplica el loop de 10 principios (OBSERVE-OBSERVE-LISTEN-THINK-CONNECT-CONNECT-FEEL-ACCEPT-CREATE-GROW)
  a problemas no triviales. Guía al agente a través de observación externa, metacognición,
  escucha activa, análisis de primeros principios, conexión lateral, orquestación sistémica,
  intuición, humildad intelectual, salida generativa y crecimiento iterativo.
  Triggers on: "/think", "think this through", "piensa esto", "razoná esto",
  "analicemos", "revisión estructurada", "auditar decisión", "qué me estoy perdiendo",
  "deep think", "systematic thinking", "structured reasoning".
---

# think: El Loop de 10 Principios

Una meditación, una disciplina y una checklist. Usá este skill cuando un problema
sea lo suficientemente no trivial como para que el pensamiento disciplinado
pague su costo: decisiones arquitectónicas, post-mortems, pedidos ambiguos del
usuario, auditorías, tradeoffs multi-stakeholder, momentos de "shippeamos o no",
"qué nos estamos perdiendo".

Los 10 principios no son una receta. Son **etapas de atención**. Movete a través
de ellas en orden en el primer pase, después volvé a las anteriores cuando surja
nueva información. La disciplina está en NO saltearte las incómodas
(OBSERVE-internal, ACCEPT, GROW) solo porque son incómodas.

---

## Los 10 Principios

### 1. OBSERVE (input externo)
Pensar empieza con recolectar datos. Mirá el entorno, el landscape actual,
los patrones, las ineficiencias, las oportunidades — sin tratar de resolverlos
inmediatamente. Leé los inputs crudos.

En práctica: leé el código antes de cambiarlo. Leé la wiki antes de responder.
Resistí el impulso de saltar a una solución ante el primer síntoma.

### 2. OBSERVE (metacognición interna)
Ahora observate a vos mismo. ¿Estás operando sobre suposiciones? ¿Tenés
sesgo en esta arquitectura? ¿Estás anclado en una decisión previa?

En práctica: escribí un párrafo de "bias log" antes de decidir. Notá sesgo
de pertenencia, sesgo de ship-it, sesgo de familiaridad, anclaje.
El sesgo no desaparece por notarlo — se contiene.

### 3. LISTEN (receptividad activa)
Observar es visual o analítico. Escuchar requiere apagar el ego para absorber
feedback externo. Prestá atención a la intención del usuario, mensajes de error,
señales sutiles en el ruido.

En práctica: leé el pedido exacto del usuario antes de parafrasearlo. Leé el
mensaje de error antes de adivinar el failure mode. La confusión del usuario
es datos.

### 4. THINK (procesamiento crítico)
El motor analítico. Una vez que tenés los inputs, rompé el problema en
primeros principios. Estructurá la lógica, mapeá los flujos, evaluá las
restricciones, sintetizá los datos crudos en una estrategia coherente.

En práctica: leé-antes-de-escribir. Nombrá como si el próximo lector fuera
hostil. La unidad más chica que funciona. Borrá más de lo que agregás.
Evidencia sobre intuición. El failure mode es el spec.

### 5. CONNECT (pensamiento lateral)
Las grandes ideas raramente ocurren en vacío — ocurren en intersecciones.
Tomá dos conceptos aparentemente no relacionados y conectalos.

En práctica: cuando audités un skill, preguntate "este bug existe en skills
vecinos?" Cuando diseñes una API, preguntate "a qué otra interfaz es isomorfa?"

### 6. CONNECT (orquestación sistémica)
El segundo CONNECT es sobre ejecución. Pasar de una idea aislada a un sistema
integrado. ¿Cómo se enchufan estas piezas para crear un todo funcional?

En práctica: cuando agregues un skill, auditá cómo se integra con hooks,
scripts, el Makefile, los agentes. El skill que funciona aislado pero rompe
el pipeline de tests no es un skill que funciona.

### 7. FEEL (inteligencia emocional + intuición)
La lógica pura es frágil sin empatía. Diseñá con la experiencia del usuario
en mente. Confiá en la intuición ganada con esfuerzo cuando los datos son
ambiguos.

En práctica: un mensaje de error que dice "ERR: exit code 4" falla FEEL aunque
pase THINK. El usuario que usa el Bunker por primera vez experimenta tus
decisiones de una manera que `make test` no puede medir.

### 8. ACCEPT (humildad intelectual)
Ningún plan sobrevive el primer contacto con la realidad. Abrazá las
restricciones. Reconocé cuando una hipótesis falla. Soltá el costo hundido.

En práctica: calificá tus hallazgos honestamente. Si un test es 78/100, no
escribas 95/100. Si el veredicto es YELLOW, no lo llames GREEN. ACCEPT es
el firewall contra la sycophancy.

### 9. CREATE (salida generativa)
La parálisis por análisis es enemiga del progreso. En algún punto, parar de
estrategizar y empezar a producir. Escribí el código. Lanzá el sistema.
Shippeá el reporte.

En práctica: una auditoría que nunca se escribe es peor que una auditoría B+
que se shippea. Un fix que queda en el working tree para siempre es peor
que el mismo fix commiteado. CREATE es la respuesta cuando las etapas
anteriores te dieron suficiente.

### 10. GROW (el loop iterativo)
Pensar no es una línea recta — es un loop de feedback. Tomá lo que construiste
(CREATE), mirá cómo funciona en realidad, y usá esas lecciones para mejorar
en el próximo ciclo.

En práctica: toda auditoría debe terminar con una sección GROW. ¿Qué funcionó?
¿Qué mejorar para la próxima? ¿Qué inputs alimentan la v_next?
GROW es lo que convierte una buena decisión en un hábito que compone.

---

## Cuándo Invocarlo

Invocá `/think` cuando:

- Estés por tomar una decisión arquitectónica no trivial
- Estés auditando un sistema y necesites una columna vertebral metodológica
- El pedido del usuario sea ambiguo y necesites escuchar más antes de responder
- Encontrés un resultado sorprendente y necesites OBSERVE-internal antes de ajustar
- Estés por llamar algo "terminado" y necesites verificar ACCEPT antes de claimear el veredicto
- Un post-mortem después de que algo salió mal
- Estés cerrando una sesión y necesites un paso GROW antes del save

NO invocar `/think` para:

- Fixes de una línea (la disciplina es overkill)
- Búsquedas triviales (no hay decisión que tomar)
- Casos donde ya pasaste por los 10 implícitamente

El valor del framework escala con **novedad del problema + irreversibilidad**.
Para un fix de una línea fácil de revertir, el loop es peso muerto.
Para una decisión de auditoría que bloquea un release, saltear cualquier
etapa pierde calibración.

---

## Cómo Usar

```
/think <planteo del problema>
```

Camina a través de las 10 etapas en orden. Cada etapa alimenta la siguiente.
CREATE produce un artefacto o recomendación.

Las etapas 1, 4 y 9 suelen ser cortas. Las etapas 2, 7, 8 y 10 son donde
la mayoría saltea. Cuidate ahí.

---

## Prompts por Etapa

### 1. OBSERVE (externo)
- ¿Cuáles son los inputs crudos? (Código? Wiki? Intención del usuario? Logs?)
- ¿Qué leí completo vs. lo que asumí?
- ¿Cuál es el estado actual? (Working tree, tests, deploy)
- ¿Qué me sorprende, antes de empezar a interpretar?

### 2. OBSERVE (interno)
- ¿Hacia dónde estoy sesgado? (Pertenencia, ship-it, novedad, anclaje, familiaridad)
- ¿Qué resultado espero inconscientemente? ¿Por qué?
- Si un revisor fresh-context se uniera ahora, ¿qué cuestionaría que yo doy por sentado?
- Mi confianza está calibrada a la evidencia que realmente tengo?

### 3. LISTEN
- ¿Qué pidió exactamente el usuario? (Cita textual.)
- ¿Qué señales hay en el ruido? (Lo que NO dijo, correcciones previas en el hilo)
- ¿Hay señales del dominio que debería consultar? (GitHub issues, docs, patrones vecinos)

### 4. THINK
- ¿Cuáles son los primeros principios? (Restricciones, invariantes, radio de explosión)
- Alternativas que NO consideré?
- ¿Cuál es el experimento más barato que me probaría que estoy equivocado?

### 5. CONNECT (lateral)
- ¿Dónde más aparece este patrón?
- ¿Qué dominio aparentemente no relacionado resolvió un problema estructuralmente similar?
- Si el mismo bug existe en este código, ¿existe en tres vecinos?

### 6. CONNECT (sistémico)
- ¿Cómo se enchufa esto al wiring existente? (Hooks, scripts, Makefile, skills)
- ¿Hay que actualizar algo aguas arriba / abajo?
- ¿Qué nuevos failure modes crea la integración que el componente aislado no tiene?

### 7. FEEL
- ¿Cómo CAE esto en el usuario? (UX, mensajes de error, naming, fricción de onboarding)
- ¿En qué estado emocional está el usuario cuando llega a este código? (Frustrado? Explorando? Contrarreloj?)
- ¿Mi intuición dice "algo está mal" aunque los datos digan "estamos bien"?

### 8. ACCEPT
- ¿Cuál es el tier honesto de este hallazgo? (Sin inflación.)
- ¿Qué restricción me piden aflojar que NO debería aflojar?
- ¿Qué costo hundido estoy protegiendo que debería soltar?

### 9. CREATE
- ¿Cuál es el artefacto más chico que shippea la decisión?
- ¿Cuál es el camino más limpio de acá a "terminado"?
- Shippealo.

### 10. GROW
- ¿Qué funcionó bien en este ciclo?
- ¿Qué haría diferente la próxima vez?
- ¿Dónde debería guardar esta lección para que el futuro-yo no tenga que re-derivarla? (Wiki? Memoria? Hot.md? Un ADR?)

---

## Anti-patrones

El loop falla cuando:

- **Saltear OBSERVE-internal.** Ir directo de observación externa a THINK sin auditar sesgos produce respuestas confiadamente incorrectas.
- **Saltear ACCEPT.** Inflar un score, amigar un veredicto, llamar YELLOW "GREEN con disclosure". El contrato anti-sycophancy del framework muere cuando ACCEPT se vuelve opcional.
- **Saltear GROW.** Producir el artefacto, shippearlo, y seguir sin feedback. El próximo ciclo empieza desde la misma línea de base; nada compone.
- **Parálisis en THINK.** Loopeando dentro de la etapa 4 para siempre, sin llegar a CREATE. El framework es una secuencia, no una regla de parada. ACEPTÁ lo que tenés, CREÁ el artefacto, CRECÉ de la respuesta.
- **Ceremonia.** Escribir las 10 etapas para un fix de una línea. El costo del framework escala con el stake del problema.

---

## Composición con otros skills del Bunker

- **`wiki-query`**: El OBSERVE externo — leer la wiki antes de decidir.
- **`autoresearch`**: Un amplificador de LISTEN — investigar señales externas antes de decidir.
- **`code-review`**: Un OBSERVE-internal delegado — revisor fresh-context que atrapa sesgos.
- **`Makefile test`**: Un GROW sistémico — el pipeline de tests te dice si tu CREATE funcionó.
- **`wiki/hot.md`**: Donde deberían ir los outputs de GROW para la próxima sesión.

---

## Referencia

- Este skill es la fuente canónica del framework de 10 principios en el Bunker OS.
- No modifica archivos ni ejecuta mutaciones. Carga estructura y disciplina;
  lo que hagas con eso es la próxima decisión.

Los 10 principios son una meditación. Sin la postura, el framework se vuelve
ceremonia. Con la postura, cada ciclo compone.
