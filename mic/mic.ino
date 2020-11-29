#define DATA_LEN 400
void setup() {
	Serial.begin(9600);
	pinMode(2, INPUT);
}

int i = 0;
bool data[DATA_LEN];
void loop() {
	if(i > DATA_LEN) {
		for (int i = 0; i < DATA_LEN; i++) {
			Serial.println(data[i]);
		}
		delay(10000);
	}
	else {
		data[i] = digitalRead(2);
	}
	i += 1;
}
