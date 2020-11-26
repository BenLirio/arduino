#define NUM_MOTOR_CONTROL 4
int motorControl[NUM_MOTOR_CONTROL] = {2,3,4,5};

void setup() {
	for(int i = 0; i < NUM_MOTOR_CONTROL; i++) {
		pinMode(motorControl[i], OUTPUT);
	}
}

int tick = 0;
void loop() {
	tick += 1;
	int active = tick % 4;
	for(int i = 0; i < NUM_MOTOR_CONTROL; i++) {
		if(active == i) {
			digitalWrite(motorControl[i], HIGH);
		}
		else {
			digitalWrite(motorControl[i], LOW);	
		}
	}
	delay(5);
}
