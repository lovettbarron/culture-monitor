# Cultural Monitoring Station

Small prototype to try and make a sourdough starter sensing station. Part of a broader attempt to see what interesting things I can do between electronics, software, and the bacteria we use in food.

```arduino``` is just an initial sketch to use a heat sensor (DHT-22) and a pair of probes poking into the sourdough solution.

```main.go``` just writes to culture.csv

CSV rows:
date, humidity, temp, humidity index, random value from probes

```culture.csv``` is just the collected data w/ time