#include "DHT.h"

#define DHTPIN 2
#define DHTTYPE DHT22   // DHT 22  (AM2302)

#define VDIVPIN 0

DHT dht(DHTPIN, DHTTYPE);

void setup() {
  Serial.begin(9600);
  // Serial.println("");

  dht.begin();
}

void loop() {
  delay(6*1000);
  humiditySensorLoop();
  voltageDivider();
  Serial.println("");
}

void voltageDivider() {
  // This is an arbitrary value from two probes inside the sour dough
  // I might just be picking up radio waves...
  float h=analogRead(VDIVPIN);
  Serial.print(h);
}


void humiditySensorLoop() {
  float h = dht.readHumidity();
  float t = dht.readTemperature();
  float f = dht.readTemperature(true);

  if (isnan(h) || isnan(t) || isnan(f)) {
    Serial.println("Failed to read from DHT sensor!");
    return;
  }
  float hic = dht.computeHeatIndex(f, h,true);

// Humidity, Temp (in f), Humidity index
  Serial.print(h);
  Serial.print(',');
  Serial.print(f);
  Serial.print(',');
  Serial.print(hic);
  Serial.print(',');
}